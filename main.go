package main

import (
	"2068_Scouting/TCP"
	"fmt"
)

func main() {
	//defer Save
	var input string
	var host bool

	for {
		fmt.Println("H O M E")
		if host {
			fmt.Println("Host Options: -q = quit hosting | -n = new competition")
		} else {
			fmt.Println("Options: -q = close app | -h = host | -c = connect to a host")
		}

		TCP.ListCompetitions()

		fmt.Scanln(&input)

		if host {
			switch input {
			case "-q":
				TCP.ShutDown()
				host = false
				fmt.Println("Stopped Hosting")
			case "-n":
				formMenu(&input)
			}
		} else {
			switch input {
			case "-c":
				fmt.Println("IP: ")
				fmt.Scanln(&input)
				TCP.ConnectToTCP(input)
				fmt.Println("You Now Connected!")
			case "-h":
				go TCP.StartTCP()
				fmt.Println("\nStarted Hosting")
				fmt.Println("Share this IP: " + TCP.FindIP() + "\n")
				host = true
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

	fmt.Println("-a = new QnA | -b = back a question | -f = show full form | -q = exits menu")
	fmt.Scanln(input)

	for {
		switch *input {
		case "-a":
		case "-b":
		case "-f":
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
