package input_test

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
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

func TestWithCharLimit(t *testing.T) {
	// Initialize input and output buffers
	var in bytes.Buffer
	var out bytes.Buffer

	inputString := []byte(strings.Repeat("a", 400) + "\r\n")

	in.Write(inputString)

	// Create a new model with a custom char limit using WithCharLimit
	model := input.New(
		"test",
		input.WithCharLimit(400),
	)

	// Run the model using tea program
	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	// Retrieve the final model after running the tea program
	m, ok := tm.(input.Model)
	require.True(t, ok)

	// Check if the CharLimit is correctly set
	require.Equal(t, 400, len(m.Data()))
}

func TestWithSmallCharLimitError(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	inputString := []byte(strings.Repeat("a", 50) + "\r\n")
	in.Write(inputString)

	model := input.New(
		"test",
		input.WithCharLimit(10),
	)

	tm, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(input.Model)
	require.True(t, ok)

	require.Equal(t, 10, len(m.Data()))
}
