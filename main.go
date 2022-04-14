package main

import (
	"2068_Scouting/Database"
	"fmt"
)

var competitions []struct {
	Name      string
	Questions []Database.Question
}

// var form struct {
// 	Team  uint16
// 	Questions []Database.Question
// }

//

func altmain() {
	//defer Save
	var input string
	fmt.Println("WELCOME!!!")
	for {
		fmt.Println("H O M E")
		fmt.Println("Options: -q = close app | -n = new competiton | -h = host | -c = connect to a host")
		//List all current competitions
		fmt.Scanln(&input)

		switch input {
		case "-n":
			Competition(&input)
		case "-h":
		case "-l":
		case "-q":
			return
		}
	}
}

func Competition(input *string) {
	// var active_questions []Database.Question
	fmt.Scanln(input)
	for _, v := range competitions {
		if v.Name == *input {
			fmt.Println(v.Name)
			// active_questions = v.Questions
			break
		}
	}

	fmt.Println("Exiting Competitions...\n")
}
