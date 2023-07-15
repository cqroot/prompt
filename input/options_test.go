package input_test

import (
	"bytes"
	"errors"
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/input"
)

func TestWithTeaProgramOpts(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	withInput := tea.WithInput(&in)
	withOutput := tea.WithOutput(&out)

	model := input.New(
		"",
		input.WithTeaProgramOpts(withInput, withOutput),
	)

	require.True(t, reflect.ValueOf(withInput) == reflect.ValueOf(model.TeaProgramOpts()[0]))
	require.True(t, reflect.ValueOf(withOutput) == reflect.ValueOf(model.TeaProgramOpts()[1]))
}

func TestWithValidateFunc(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	validateErr := errors.New("validation error")
	validateFunc := func(s string) error {
		if s != "test" {
			return validateErr
		}
		return nil
	}

	in.Write([]byte("\r\n"))

	model := input.New("", input.WithValidateFunc(validateFunc))

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(input.Model)
	require.True(t, ok)

	require.Equal(t, m.Error(), validateErr)
}

func TestDefaultValueWithValidateFunc(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	validateErr := errors.New("validation error")
	validateFunc := func(s string) error {
		if s != "test" {
			return validateErr
		}
		return nil
	}

	in.Write([]byte("\r\n"))

	model := input.New("test", input.WithValidateFunc(validateFunc))

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(input.Model)
	require.True(t, ok)

	require.Nil(t, m.Error())
}
