package test

import (
	"2068_Scouting/TCP"
	"os"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	if _, err := os.ReadDir("save"); err != nil {
		os.Mkdir("save", os.ModeDir)
	}

	//Saving
	save, err := os.OpenFile(path+"/save/comps.MetalJacket", os.O_CREATE, os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	test_forms := []TCP.Form{
		{Team: 2068, Questions: []string{"searching"}, Answers: []string{"blame"}},
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
	t.Log(data)

	raw := TCP.SeperateBy(data, byte('µ'))
	var newForms []TCP.Form

	t.Log(raw)

	if len(raw) == 0 {
		t.Error("Raw = 0; Something seems off...")
	}

	for i := range raw {
		if len(raw[i]) == 0 {
			t.Log(i)
			continue
		}
		t.Log(raw[i])
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
				t.Errorf("Questions %v in form %v don't match up", x, i)
			}

			if newForms[i].Answers[x] != test_forms[i].Answers[x] {
				t.Errorf("Answers %v in form %v don't match up", x, i)
			}
		}
	}
}

func TestTemplateSave(t *testing.T) {
	//Add to the comp templates
	//Save the comps and keys
	//Load the comps and keys
	//Check the comps and keys against what you originally sent
}
