package TCP

import (
	"encoding/binary"
	"log"
	"os"
	"strings"
	"sync"
)

var compKeys []string //Used by Client

var CompTemplates = map[string]Form{} //Used by Host

var Database = netComp{
	Competitions: make(map[string][]Form),
	RWMutex:      &sync.RWMutex{},
}

type netComp struct {
	Competitions map[string][]Form
	*sync.RWMutex
}

type Form struct {
	Team               uint16
	Questions, Answers []string
}

func AddCompetition(compName string, newForm Form) {
	if compName == "" {
		return
	}

	CompTemplates[compName] = newForm
}

func ListCompetitions() string {
	var output string
	for k := range CompTemplates {
		output += k + "\n"
	}

	return output
}

func SubmitForm(form Form) {
	conn.Write(append([]byte("Comp "), ToBytes(form)...))
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

func FromBytes(data []byte, template bool) Form {
	teamByte := binary.BigEndian.Uint16(data[:2])
	newForm := Form{Team: teamByte}

	body := strings.Split(string(data[2:]), "¶")

	for i := 0; i < len(body)-1; i++ {
		if template {
			newForm.Questions = append(newForm.Questions, body[i])

		} else {
			if i%2 == 0 {
				newForm.Questions = append(newForm.Questions, body[i])
			} else {
				newForm.Answers = append(newForm.Answers, body[i])
			}
		}
	}

	return newForm
}

func SeperateBy(data []byte, seperator byte) [][]byte {
	var raw [][]byte
	var previous int

	for i := 0; i < len(data); i++ {
		if data[i] == seperator {
			raw = append(raw, data[previous:i])
			i++
			previous = i
		}
	}

	return raw
}

func SaveTemplates() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.ReadDir("save"); err != nil {
		os.Mkdir("save", os.ModeDir)
	}

	var keys string
	var comps []byte
	for k, v := range CompTemplates {
		keys = k + "\n"
		comps = append(comps, ToBytes(v)...)
	}

	err = os.WriteFile(path+"/save/template_names.txt", []byte(keys), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path+"/save/template_forms.MetalJacket", comps, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadTemplates() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(path + "/save/template_names.txt")
	if err != nil {
		log.Fatal(err)
	}

	raw, err := os.ReadFile(path + "/save/template_forms.MetalJacket")
	if err != nil {
		log.Fatal(err)
	}

	keys := strings.Split(string(data), "\n")

	forms := SeperateBy(raw, '\n')

	newTemp := map[string]Form{}

	for i := 0; i < len(keys)-1; i++ {
		newTemp[keys[i]] = FromBytes([]byte(forms[i]), true)
	}

	CompTemplates = newTemp
}

func Save() {
	var keys string
	var data []byte

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range Database.Competitions {
		keys += k + "\n"
		for i := range v {
			data = append(data, ToBytes(v[i])...)
		}
		data = append(data, byte('µ'))
	}

	os.WriteFile(path+"/save/keys.txt", []byte(keys), os.ModePerm)
	os.WriteFile(path+"/save/forms.MetalJacket", data, os.ModePerm)
}

func Load() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(path + "/save/keys.txt")
	if err != nil {
		log.Fatal(err)
	}

	raw, err := os.ReadFile(path + "/save/forms.MetalJacket")
	if err != nil {
		log.Fatal(err)
	}

	keys := strings.Split(string(data), "\n")

	stack := SeperateBy(raw, 'µ')

	newTemp := map[string][]Form{}

	var compForms [][]Form

	for i := range stack {
		var subForm []Form
		bang := SeperateBy(stack[i], '\n')
		for o := range bang {
			subForm = append(subForm, FromBytes(bang[o], false))
		}
		compForms = append(compForms, subForm)
	}

	for i := 0; i < len(keys)-1; i++ {
		newTemp[keys[i]] = compForms[i]
	}

	Database.Competitions = newTemp
}
