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

func testPromptModel(t *testing.T, model prompt.PromptModel, input []byte, val string) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(input)
	in.Write([]byte("\r\n"))

	pm, err := prompt.New().Ask("").Run(model, tea.WithInput(&in), tea.WithOutput(&out))
	require.Nil(t, err)
	require.Equal(t, val, pm.DataString())
	testPromptModel_Data(t, val, pm.Data())

	in.Reset()
	in.Write([]byte{'q'})
	_, err = prompt.New().Ask("").Run(model, tea.WithInput(&in), tea.WithOutput(&out))
	require.Equal(t, prompt.ErrUserQuit, err)
}

func testPromptModel_Data(t *testing.T, expected string, actual any) {
	dataString, ok := actual.(string)
	if ok {
		require.Equal(t, expected, dataString)
	} else {
		require.Equal(t, expected, strings.Join(actual.([]string), ", "))
	}
}
