package TCP

import (
	"encoding/binary"
	"log"
	"os"
	"strings"
	"sync"
)

var compKeys []string //Used by Client

var compTemplates = map[string]Form{} //Used by Host

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

func SaveTemplates() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.ReadDir("save"); err != nil {
		os.Mkdir("save", os.ModeDir)
	}

	var keys []string
	var comps []Form
	for k, v := range compTemplates {
		keys = append(keys, k)
		comps = append(comps, v)
	}

	//Write to a file of templates
	keys_save, err := os.OpenFile(path+"/save/template_names.txt", os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer keys_save.Close()

	form_save, err := os.OpenFile(path+"/save/template_forms.MetalJacket", os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer form_save.Close()

	for i := 0; i < len(comps); i++ {
		keys_save.Write([]byte(keys[i] + "µ"))
		form_save.Write(append(ToBytes(comps[i]), []byte("µ")...))
	}
}

func loadTemplates() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	data, err := os.ReadFile(path + "/save/template_names.txt")
	if err != nil {
		log.Fatal(err)
	}

	keys := strings.Split(string(data), "µ")

	raw, err := os.ReadFile(path + "/save/template_forms.MetalJacket")
	if err != nil {
		log.Fatal(err)
	}

	// forms := strings.Split(string(raw), "µ")
	forms := SeperateBy(raw, 'µ')

	for i := 0; i < len(keys); i++ {
		compTemplates[keys[i]] = FromBytes([]byte(forms[i]), true)
	}
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
