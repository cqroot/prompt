package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt"
)

type ToggleModelTest struct{}

func (_ ToggleModelTest) Model() prompt.PromptModel {
	return prompt.NewToggleModel([]string{"Yes", "No"})
}

func (mt ToggleModelTest) DataTestcases() []KVPair {
	return []KVPair{
		{Key: []byte{}, Val: "Yes", View: "Yes"},
		{Key: []byte("lhh"), Val: "No", View: "No"},
		{Key: []byte("kjj"), Val: "No", View: "No"},
		{Key: []byte{KeyTab}, Val: "No", View: "No"},
		{Key: []byte{' '}, Val: "No", View: "No"},
	}
}

func (mt ToggleModelTest) InitViewTestcase() string {
	return "?  › Yes / No"
}

func (mt ToggleModelTest) InitViewWithHelpTestcase() string {
	return `?  › Yes / No

←/h/j move left • →/l/k/tab/space move right • enter confirm • q/ctrl+c quit`
}

func TestToggleModel(t *testing.T) {
	testPromptModel(t, ToggleModelTest{})
}

func TestToggle(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{'q'})

	_, err := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		Toggle([]string{"Yes", "No"})
	require.Equal(t, prompt.ErrUserQuit, err)

	testcases := ToggleModelTest{}.DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte("\r\n"))

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			Toggle([]string{"Yes", "No"})
		require.Nil(t, err)
		require.Equal(t, testcase.Val, val)
	}
}
