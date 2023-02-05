package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

type ChooseModelTest struct{}

func (_ ChooseModelTest) Model() prompt.PromptModel {
	return prompt.NewChooseModel([]string{"Item 1", "Item 2", "Item 3"})
}

func (mt ChooseModelTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	pm := mt.Model()
	return pm, []KVPair{
		{[]byte{}, "Item 1"},
		{[]byte("kkjjj"), "Item 2"},
		{[]byte{'k', 'k', KeyTab, KeyTab, KeyTab}, "Item 2"},
	}
}

func (mt ChooseModelTest) ViewTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, `?  › 
• Item 1
  Item 2
  Item 3
`
}

func TestChoose(t *testing.T) {
	testPromptModel(t, ChooseModelTest{})
}
