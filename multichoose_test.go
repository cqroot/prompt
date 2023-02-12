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

func (mt MultiChooseModelTest) DataTestcases() []KVPair {
	return []KVPair{
		{Key: []byte{}, Val: "", View: ""},
		{Key: []byte("kk jj "), Val: "Item 1, Item 2", View: "Item 1, Item 2"},
		{Key: []byte("kk  jj "), Val: "Item 1", View: "Item 1"},
		{Key: []byte("kk jj  "), Val: "Item 2", View: "Item 2"},
		{Key: []byte("kk  jj  "), Val: "", View: ""},
		{Key: []byte{'k', 'k', ' ', KeyTab, KeyTab, ' '}, Val: "Item 1, Item 2", View: "Item 1, Item 2"},
	}
}

func (mt MultiChooseModelTest) InitViewTestcase() string {
	return `?  › 
[•] Item 1
[ ] Item 2
[ ] Item 3
`
}

func (mt MultiChooseModelTest) InitViewWithHelpTestcase() string {
	return `?  › 
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

	testcases := MultiChooseModelTest{}.DataTestcases()
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
