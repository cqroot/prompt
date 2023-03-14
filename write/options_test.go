package write_test

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/tester"
	"github.com/cqroot/prompt/write"
)

func TestWithTeaProgramOpts(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	withInput := tea.WithInput(&in)
	withOutput := tea.WithOutput(&out)

	model := write.New(
		"",
		write.WithTeaProgramOpts(withInput, withOutput),
	)

	require.True(t, reflect.ValueOf(withInput) == reflect.ValueOf(model.TeaProgramOpts()[0]))
	require.True(t, reflect.ValueOf(withOutput) == reflect.ValueOf(model.TeaProgramOpts()[1]))
}

func TestWithValidateFunc(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	validateErr := errors.New("validation error")
	validateFunc := func(string) error {
		return validateErr
	}

	in.Write([]byte{tester.KeyCtrlD})

	model := write.New("", write.WithValidateFunc(validateFunc))

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(write.Model)
	require.True(t, ok)

	require.Equal(t, m.Error(), validateErr)
}
