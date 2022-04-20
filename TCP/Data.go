package TCP

import (
	"encoding/binary"
	"strings"
	"sync"
)

var compKeys []string //Used by Client

var compTemplates map[string]Form //Used by Host

var database = netComp{
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
}

func ToBytes(form Form) []byte {
	teamByte := make([]byte, 2)
	binary.BigEndian.PutUint16(teamByte, form.Team)

	var bodyString string

	for i := range form.Questions {
		bodyString += form.Questions[i] + "¶" + form.Answers[i] + "¶"
	}
	bodyString += "\n"

	return append(teamByte, []byte(bodyString)...)
}

func FromBytes(data []byte) Form {
	teamByte := binary.BigEndian.Uint16(data[:2])
	newForm := Form{Team: teamByte}

	body := strings.Split(string(data[2:]), "¶")

	for i := 0; i < len(body)-1; i++ {
		if i%2 == 0 {
			newForm.Questions = append(newForm.Questions, body[i])
		} else {
			newForm.Answers = append(newForm.Answers, body[i])
		}
	}

	return newForm
}

func AddCompetition(compName string, newForm Form) {
	if compName == "" {
		return
	}

	compTemplates[compName] = newForm
}

func ListCompetitions() string {
	var output string
	for k := range compTemplates {
		output += k + "\n"
	}

	return output
}

func SubmitForm(form Form) {
	conn.Write(append([]byte("Comp "), ToBytes(form)...))
}
