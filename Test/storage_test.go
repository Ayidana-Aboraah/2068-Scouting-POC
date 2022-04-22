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
		CompareForm(test_forms[i], newForms[i], t)
	}
}

func TestTemplateSave(t *testing.T) {
	test_form := TCP.Form{
		Team:      2068,
		Questions: []string{"Duck"},
		Answers:   []string{""},
	}
	TCP.AddCompetition("blue_man", test_form)

	TCP.SaveTemplates()

	TCP.LoadTemplates()

	form, found := TCP.CompTemplates["blue_man"]
	if !found {
		t.Error("Name not working")
	}

	t.Log(form)

	// CompareForm(TCP.CompTemplates["blue_man"], form, t)

	if form.Team != TCP.CompTemplates["blue_man"].Team {
		t.Errorf("Test Team: %v and New Team %v |don't match up", TCP.CompTemplates["blue_man"].Team, form.Team)
	}

	for x := range TCP.CompTemplates["blue_man"].Questions {
		if form.Questions[x] != TCP.CompTemplates["blue_man"].Questions[x] {
			t.Errorf("Questions %v don't match up", x)
		}
	}
}

func CompareForm(reference, newForm TCP.Form, t *testing.T) {
	if newForm.Team != reference.Team {
		t.Errorf("Test Team: %v and New Team %v |don't match up", reference.Team, newForm.Team)
	}

	for x := range newForm.Questions {
		if newForm.Questions[x] != reference.Questions[x] {
			t.Errorf("Questions %v don't match up", x)
		}

		if newForm.Answers[x] != reference.Answers[x] {
			t.Errorf("Answers %v don't match up", x)
		}
	}
}

func TestMom(t *testing.T) {
	TCP.LoadTemplates()
}
