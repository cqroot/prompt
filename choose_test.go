package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

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

func (mt ChooseModelTest) ViewWithHelpTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, `?  › 
• Item 1
  Item 2
  Item 3

↑/k move up • ↓/j/tab move down • enter confirm • q/ctrl+c quit`
}

func TestChooseModel(t *testing.T) {
	testPromptModel(t, ChooseModelTest{})
}

func TestChoose(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{'q'})

	_, err := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		Choose([]string{"Item 1", "Item 2", "Item 3"})
	require.Equal(t, prompt.ErrUserQuit, err)

	_, testcases := ChooseModelTest{}.DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte("\r\n"))

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			Choose([]string{"Item 1", "Item 2", "Item 3"})
		require.Nil(t, err)
		require.Equal(t, testcase.Val, val)
	}
}
