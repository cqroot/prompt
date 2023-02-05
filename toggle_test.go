package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

type ToggleModelTest struct{}

func (_ ToggleModelTest) Model() prompt.PromptModel {
	return prompt.NewToggleModel([]string{"Yes", "No"})
}

func (mt ToggleModelTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	pm := mt.Model()
	return pm, []KVPair{
		{[]byte{}, "Yes"},
		{[]byte("lhh"), "No"},
		{[]byte("kjj"), "No"},
		{[]byte{KeyTab}, "No"},
		{[]byte{' '}, "No"},
	}
}

func (mt ToggleModelTest) ViewTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, "?  › Yes / No"
}

func TestToggle(t *testing.T) {
	testPromptModel(t, ToggleModelTest{})
}
