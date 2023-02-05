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
	KeyTab byte = 9
)

func testPromptModel(t *testing.T, model prompt.PromptModel, input []byte, defaultVal string, val string) {
	var out bytes.Buffer
	var in bytes.Buffer

	p := prompt.New().Ask("").SetModel(model)

	in.Write(input)
	in.Write([]byte("\r\n"))

	prog := tea.NewProgram(p, tea.WithInput(&in), tea.WithOutput(&out))
	tm, err := prog.Run()
	require.Nil(t, err)
	require.Nil(t, p.Error())

	require.Equal(t, defaultVal, p.Model().DataString())
	require.Equal(t, val, tm.(prompt.Prompt).Model().DataString())

	testPromptModel_Data(t, defaultVal, p.Model().Data())
	testPromptModel_Data(t, val, tm.(prompt.Prompt).Model().Data())
}

func testPromptModel_Data(t *testing.T, expected string, actual any) {
	dataString, ok := actual.(string)
	if ok {
		require.Equal(t, expected, dataString)
	} else {
		require.Equal(t, expected, strings.Join(actual.([]string), ", "))
	}
}
