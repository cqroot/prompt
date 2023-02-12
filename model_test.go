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
	KeyCtrlC byte = 3
	KeyTab   byte = 9
	KeyCtrlS byte = 19
)

type KVPair struct {
	Key  []byte
	Val  string
	View string
}

type PromptModelTest interface {
	Model() prompt.PromptModel
	DataTestcases() []KVPair
	InitViewTestcase() string
	InitViewWithHelpTestcase() string
}

// testPromptModelData tests whether the returned result is as expected after
// the specified key input.
func testPromptModelData(t *testing.T, model prompt.PromptModel, input []byte, val string, view string) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(input)
	if model.UseKeyEnter() {
		in.Write([]byte{KeyCtrlS})
	} else {
		in.Write([]byte("\r\n"))
	}

	pm, err := prompt.New().Ask("").
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		Run(model)
	require.Nil(t, err)

	dataString, ok := pm.Data().(string)
	if ok {
		require.Equal(t, val, dataString)
	} else {
		require.Equal(t, val, strings.Join(pm.Data().([]string), ", "))
	}

	require.Equal(t, view, pm.DataString())
}

// testPromptModelError tests whether the corresponding error is returned after
// the user quits.
func testPromptModelError(t *testing.T, model prompt.PromptModel) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	_, err := prompt.New().Ask("").
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		Run(model)
	require.Equal(t, prompt.ErrUserQuit, err)

	if model.UseKeyQ() {
		return
	}

	in.Reset()
	in.Write([]byte{'q'})

	_, err = prompt.New().Ask("").
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		Run(model)
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
	p := prompt.New().Ask("").SetModel(model).WithHelp(true)
	require.Equal(t, view, p.View())
}

func testPromptModel(t *testing.T, pmt PromptModelTest) {
	pairs := pmt.DataTestcases()
	for _, pair := range pairs {
		testPromptModelData(t, pmt.Model(), pair.Key, pair.Val, pair.View)
	}

	testPromptModelError(t, pmt.Model())

	view := pmt.InitViewTestcase()
	testPromptModelView(t, pmt.Model(), view)

	view = pmt.InitViewWithHelpTestcase()
	testPromptModel_ViewWithHelp(t, pmt.Model(), view)
}
