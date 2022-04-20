package main

import (
	"2068_Scouting/TCP"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var status uint8

type blank struct{}

func main() {
	input := bufio.NewScanner(os.Stdin)
	var header string
	for {
		switch status {
		case 2:
			header = "Host Options | IP Address: " + TCP.FindIP() + " | -q = stop hosting | -n = new competition"
		case 1:
			header = "Client Options: -q = disconnect\n | -n = new submission"
		default:
			header = "Options: -q = close app | -h = host | -c = connect to a host"
		}
		header += "\n" + TCP.ListCompetitions()
		fmt.Println(header)

		input.Scan()

		switch input.Text() {
		case "-q", "quit":
			switch status {
			case 2:
				TCP.ShutDown()
			case 1:
				TCP.DisconnectTCP()
			case 0:
				return
			}
			status = 0
		case "-h", "host":
			go TCP.StartTCP()
			status = 2
		}

		cmd := strings.Split(input.Text(), " ")

		switch cmd[0] {
		case "-n", "new":
			if status == 2 {
				NewCompetition(input, blank{})
			} else if status == 1 {
				NewCompetition(input, struct{}{})
			}
		case "-c", "connect":
			TCP.ConnectToTCP(cmd[1])
			status = 1
		}
		fmt.Println(cmd)
	}
}

func NewCompetition(input *bufio.Scanner, submission struct{}) {
	var stop struct{}
	var newForm TCP.Form
	var currentIdx uint8

	fmt.Println("Commands: -c or close + name (if empty it will not save) = to stop editing")
	for {
		input.Scan()

		fmt.Println(newForm.Questions[currentIdx])

		switch input.Text() {
		case "-c", "close":
			stop = struct{}{}
			fmt.Println("Closing...")
		case ">":
			currentIdx++
		case "<":
			currentIdx--
		}

		if stop == struct{}{} {
			fmt.Println("Enter the Competition Name...")
			input.Scan()
			if submission == struct{}{} {
				TCP.SubmitForm(newForm)
			} else {
				TCP.AddCompetition(input.Text(), newForm) //Save
			}
			return
		}

		if int(currentIdx) < len(newForm.Questions) {
			newForm.Questions[currentIdx] = input.Text()
			continue
		}

		newForm.Questions = append(newForm.Questions, input.Text())

		currentIdx++
	}
}
