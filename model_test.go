package prompt_test

import (
	"bytes"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt"
)

const (
	KeyTab   byte = 9
	KeyCtrlC byte = 3
)

type KVPair struct {
	Key []byte
	Val string
}

type PromptModelTest interface {
	Model() prompt.PromptModel
	DataTestcases() (prompt.PromptModel, []KVPair)
	ViewTestcases() (prompt.PromptModel, string)
	ViewWithHelpTestcases() (prompt.PromptModel, string)
}

// testPromptModelData tests whether the returned result is as expected after
// the specified key input.
func testPromptModelData(t *testing.T, model prompt.PromptModel, input []byte, val string) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(input)
	in.Write([]byte("\r\n"))

	pm, err := prompt.New().Ask("").Run(model, tea.WithInput(&in), tea.WithOutput(&out))
	require.Nil(t, err)
	require.Equal(t, val, pm.DataString())

	dataString, ok := pm.Data().(string)
	if ok {
		require.Equal(t, val, dataString)
	} else {
		require.Equal(t, val, strings.Join(pm.Data().([]string), ", "))
	}
}

// testPromptModelError tests whether the corresponding error is returned after
// the user quits.
func testPromptModelError(t *testing.T, model prompt.PromptModel) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	_, err := prompt.New().Ask("").Run(model, tea.WithInput(&in), tea.WithOutput(&out))
	require.Equal(t, prompt.ErrUserQuit, err)

	if model.UseKeyQ() {
		return
	}

	in.Reset()
	in.Write([]byte{'q'})

	_, err = prompt.New().Ask("").Run(model, tea.WithInput(&in), tea.WithOutput(&out))
	require.Equal(t, prompt.ErrUserQuit, err)
}

// testPromptModelView tests that the model's interface displays as expected.
func testPromptModelView(t *testing.T, model prompt.PromptModel, view string) {
	p := prompt.New().Ask("").SetModel(model)
	require.Equal(t, view, p.View())
}

// testPromptModel_ViewWithHelp tests that the model interface with the help
// message displays as expected
func testPromptModel_ViewWithHelp(t *testing.T, model prompt.PromptModel, view string) {
	p := prompt.New().Ask("").SetModel(model).SetHelpVisible(true)
	require.Equal(t, view, p.View())
}

func testPromptModel(t *testing.T, pmt PromptModelTest) {
	model, pairs := pmt.DataTestcases()
	for _, pair := range pairs {
		testPromptModelData(t, model, pair.Key, pair.Val)
	}

	testPromptModelError(t, pmt.Model())

	model, view := pmt.ViewTestcases()
	testPromptModelView(t, model, view)

	model, view = pmt.ViewWithHelpTestcases()
	testPromptModel_ViewWithHelp(t, model, view)
}
