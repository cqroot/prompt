package prompt_test

import (
	"bytes"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

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

↑/k move up • ↓/j/tab move down • space choose • enter confirm • q/ctrl+c quit`
}

func TestMultiChooseModel(t *testing.T) {
	testPromptModel(t, MultiChooseModelTest{})
}

func TestMultiChoose(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{'q'})

	_, err := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		MultiChoose([]string{"Item 1", "Item 2", "Item 3"})
	require.Equal(t, prompt.ErrUserQuit, err)

	_, testcases := MultiChooseModelTest{}.DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte("\r\n"))

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			MultiChoose([]string{"Item 1", "Item 2", "Item 3"})
		require.Nil(t, err)
		if testcase.Val == "" {
			require.Equal(t, []string{}, val)
		} else {
			require.Equal(t, strings.Split(testcase.Val, ", "), val)
		}
	}
}
