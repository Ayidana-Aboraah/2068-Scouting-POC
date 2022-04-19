package main

import (
	"2068_Scouting/TCP"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var status uint8

func main() {
	input := bufio.NewScanner(os.Stdin)
	var header string
	for {
		input.Scan()

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
		}

		cmd := strings.Split(input.Text(), " ")

		switch cmd[0] {
		case "-n", "new":
			if status == 2 {
				NewCompetition(input)
			} else if status == 1 {
				NewSubmission(input)
			}
		case "-c", "connect":
			TCP.ConnectToTCP(cmd[1])
		}
		fmt.Println(cmd)
	}
}

func NewCompetition(input *bufio.Scanner) {
	for {
		input.Scan()

		switch input.Text() {
		case "":
		}
	}
}

func NewSubmission(input *bufio.Scanner) {
	for {
		input.Scan()
	}
}
