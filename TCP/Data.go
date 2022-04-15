package TCP

import (
	"encoding/binary"
	"fmt"
	"strings"
	"sync"
)

var compList []string

var compTemplates map[string]Form

var competitions = netComp{
	competitions: make(map[string][]Form),
	RWMutex:      &sync.RWMutex{},
}

type netComp struct {
	competitions map[string][]Form
	*sync.RWMutex
}

type Form struct {
	Team               uint16
	Questions, Answers []string
	// Questions []struct{ Question, Answer string }
}

func ToBytes(form Form) []byte {
	teamByte := make([]byte, 2)
	binary.BigEndian.PutUint16(teamByte, uint16(form.Team))

	var bodyString string

	for i := range form.Questions {
		bodyString += form.Questions[i] + "¶" + form.Answers[i] + "¶"
	}

	return append(teamByte, []byte(bodyString)...)
}

func FromBytes(data []byte) Form {
	teamByte := binary.BigEndian.Uint16(data[:2])
	newForm := Form{Team: teamByte}

	body := strings.Split(string(data[2:]), "¶")

	for i := range body {
		if i%2 == 0 {
			newForm.Questions = append(newForm.Questions, body[i])
		} else {
			newForm.Answers = append(newForm.Answers, body[i])
		}
	}

	return newForm
}

func AddCompetition(compName string, newForm Form) {
	compTemplates[compName] = newForm
}

func ListCompetitions() {
	fmt.Println("Competitions:", compList)
}
