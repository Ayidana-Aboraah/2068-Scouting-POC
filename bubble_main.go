package main

import (
	"2068_Scouting/TCP"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

// var status uint8

type model struct {
	current  [][]string
	cursor   int8
	selected map[int8]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if status == 2 {
				TCP.ShutDown()
			}
			TCP.DisconnectTCP()
			return m, tea.Quit
		case "up", "w":
			if m.cursor-1 > -1 {
				m.cursor--
			}
		case "down", "s":
			if m.cursor+1 < int8(len(m.current)) {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	if _, ok := m.selected[m.cursor]; ok {
		switch status {
		case 0:
			switch m.cursor {
			case 1:
				TCP.ShutDown()
				fmt.Println("Shutting Down...")
			case 0:
				go TCP.StartTCP()
			}

			status = 2

		case 1:
			switch m.cursor {
			case 1:
				TCP.ShutDown()
			case 0:
				go TCP.StartTCP()
			}

			status = 2
		case 2:
			switch m.cursor {
			case 2:
				return m, tea.Quit
			case 1:
				TCP.ConnectToTCP(TCP.FindIP())
			case 0:
				go TCP.StartTCP()
			}

			status = uint8(m.cursor)
		}

		m.cursor = 0
		delete(m.selected, m.cursor)
	}

	return m, nil
}

func (m model) View() string {
	var result string
	switch status {
	case 2:
		result = "Host | Share IP Address: " + TCP.FindIP()
	case 1:
		result = "Client | Comps: " + TCP.ListCompetitions()
	default:
		result = "Home"
	}
	result += "\n"

	for i, choice := range m.current[status] {

		cursor := " " // no cursor
		if m.cursor == int8(i) {
			cursor = ">" // cursor!
		}

		checked := "○" // not selected
		if _, ok := m.selected[int8(i)]; ok {
			checked = "●" // selected!
		}

		// Render the row
		result += fmt.Sprintf("%s %s %s\n", cursor, checked, choice)
	}
	//Options

	return result
}

func bubble_main() {
	program := tea.NewProgram(model{
		cursor: 0,
		current: [][]string{
			{"Host", "Connect", "Quit"},
			{"New Submission", "Disconnect"},
			{"New Comp/Form", "Stop Hosting"},
		},
		selected: make(map[int8]struct{}),
	})

	if err := program.Start(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
