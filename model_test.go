package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt"
)

func testModel(t *testing.T, model prompt.PromptModel, input string, defaultVal string, val string) {
	var out bytes.Buffer
	var in bytes.Buffer

	p := prompt.New().Ask("").SetModel(model)

	in.Write([]byte(input))
	in.Write([]byte("\r\n"))

	prog := tea.NewProgram(p, tea.WithInput(&in), tea.WithOutput(&out))
	tm, err := prog.Run()
	require.Nil(t, err)
	require.Nil(t, p.Error())

	require.Equal(t, defaultVal, p.Model().DataString())
	require.Equal(t, val, tm.(prompt.Prompt).Model().DataString())
}
