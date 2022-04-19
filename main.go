package main

import (
	"2068_Scouting/TCP"
	"fmt"
)

func ALTmain() {
	var input string
	var status uint8

	for {
		fmt.Println("H O M E")
		switch status {
		case 2:
			fmt.Println("Host Options: -q = quit hosting | -n = new competition")
		case 1:
		default:
			fmt.Println("Options: -q = close app | -h = host | -c = connect to a host")
		}

		TCP.ListCompetitions()

		fmt.Scanln(&input)

		switch status {
		case 2:
			switch input {
			case "-q":
				TCP.ShutDown()
				//Save current form Submissions
				status = 0
				fmt.Println("Stopped Hosting")
			case "-n":
				formMenu(&input)
			}
		case 1:
		default:
			switch input {
			case "-c":
				QnA("IP: ", &input)
				TCP.ConnectToTCP(input)
				fmt.Println("You Now Connected!")
				//status = 1
			case "-h":
				go TCP.StartTCP()
				fmt.Println("\nStarted Hosting\nShare this IP: " + TCP.FindIP() + "\n")
				//Open Saved Competition Files
				status = 2
			case "-q":
				TCP.DisconnectTCP()
				return
			}
		}
	}
}

func formMenu(input *string) {
	fmt.Println("\nWELCOME TO COMPETITION MAKER\n")

	fmt.Println("What's the name of the Competiton?")
	fmt.Scanln(input)

	var newComp TCP.Form
	newName := *input

	for {
		QnA("-a = new QnA | -b = back a question | -f = show full form | -q = exits menu", input)

		switch *input {
		case "-a":
			QnA("What's the Question?", input)
			newComp.Questions = append(newComp.Questions, *input)
		case "-b":
			fmt.Scanln(input)
			for i := range newComp.Questions {
				if newComp.Questions[i] != *input {

				}
			}
		case "-f":
			fmt.Println(newComp)
		case "-exit":
			fmt.Println("Save? [Y/N]")
			fmt.Scanln(input)
			if *input == "Y" {
				TCP.AddCompetition(newName, newComp)
				//Add the form to the competiton list
			}
			return
		}
	}
}

func QnA(question string, answer *string) {
	fmt.Println(question)
	fmt.Scanln(answer)
}
