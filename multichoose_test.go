package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

type MultiChooseModelTest struct{}

func (_ MultiChooseModelTest) Model() prompt.PromptModel {
	return prompt.NewMultiChooseModel([]string{"Item 1", "Item 2", "Item 3"})
}

func (mt MultiChooseModelTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	pm := mt.Model()
	return pm, []KVPair{
		{[]byte{}, ""},
		{[]byte("kk jj "), "Item 1, Item 2"},
		{[]byte("kk  jj "), "Item 1"},
		{[]byte("kk jj  "), "Item 2"},
		{[]byte("kk  jj  "), ""},
		{[]byte{'k', 'k', ' ', KeyTab, KeyTab, ' '}, "Item 1, Item 2"},
	}
}

func (mt MultiChooseModelTest) ViewTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, `?  › 
[•] Item 1
[ ] Item 2
[ ] Item 3
`
}

func (mt MultiChooseModelTest) ViewWithHelpTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, `?  › 
[•] Item 1
[ ] Item 2
[ ] Item 3

↑/k move up • ↓/j/tab move down • space choose • enter confirm • q quit`
}

func TestMultiChoose(t *testing.T) {
	testPromptModel(t, MultiChooseModelTest{})
}
