package test

import (
	"2068_Scouting/TCP"
	"os"
	"strings"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	err = os.Mkdir("save", os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	//Saving
	save, err := os.OpenFile(path+"/save/comps.MetalJacket", os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	test_forms := []TCP.Form{
		{},
		{},
		{},
	}

	for i := range test_forms {
		save.Write(append(TCP.ToBytes(test_forms[i]), []byte("µ")...))
	}

	save.Close()

	//Loading

	data, err := os.ReadFile(path + "/save/comps.MetalJacket")
	if err != nil {
		t.Error(err)
	}

	raw := strings.Split(string(data), "µ")
	var newForms []TCP.Form

	for i := range raw {
		newForms = append(newForms, TCP.FromBytes([]byte(raw[i]), false))
	}

	//Testing
	if len(newForms) != len(test_forms) {
		t.Errorf("One is greater than the other |Test: %v, New: %v", len(test_forms), len(newForms))
	}

	for i := range newForms {
		if newForms[i].Team != test_forms[i].Team {
			t.Errorf("Test Team: %v and New Team %v |don't match up", test_forms[i].Team, newForms[i].Team)
		}

		for x := range newForms[i].Questions {
			if newForms[i].Questions[x] != test_forms[i].Questions[x] {
				t.Errorf("Question %v in form %v don't match up", x, i)
			}

			if newForms[i].Answers[x] != test_forms[i].Answers[x] {
				t.Errorf("Question %v in form %v don't match up", x, i)
			}
		}
	}
}
