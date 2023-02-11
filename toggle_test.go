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

func (mt ToggleModelTest) ViewWithHelpTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, `?  › Yes / No

←/h/j move left • →/l/k/tab/space move right • enter confirm • q quit`
}

func TestToggleModel(t *testing.T) {
	testPromptModel(t, ToggleModelTest{})
}

func TestToggle(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{'q'})

	_, err := prompt.New().Toggle([]string{"Yes", "No"}, tea.WithInput(&in), tea.WithOutput(&out))
	require.Equal(t, prompt.ErrUserQuit, err)
}
