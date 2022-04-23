package test

import (
	"2068_Scouting/TCP"
	"testing"
)

func TestSave(t *testing.T) {
	test_form := TCP.Form{
		Team:      2068,
		Questions: []string{},
		Answers:   []string{},
	}

	TCP.Database.Competitions["test"] = []TCP.Form{test_form}

	TCP.Save()

	TCP.Database.Competitions = map[string][]TCP.Form{}

	TCP.Load()

	forms, found := TCP.Database.Competitions["test"]
	if !found {
		t.Error("value not in map")
	}

	for i := range forms {
		CompareForm(test_form, forms[i], t)
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

	TCP.CompTemplates = map[string]TCP.Form{}

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
