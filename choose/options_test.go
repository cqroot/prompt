package choose_test

import (
	"bytes"
	"reflect"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/choose"
)

func TestWithTeaProgramOpts(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	withInput := tea.WithInput(&in)
	withOutput := tea.WithOutput(&out)

	model := choose.NewWithStrings(
		[]string{"Item 1", "Item 2", "Item 3"},
		choose.WithTeaProgramOpts(withInput, withOutput),
	)

	require.True(t, reflect.ValueOf(withInput) == reflect.ValueOf(model.TeaProgramOpts()[0]))
	require.True(t, reflect.ValueOf(withOutput) == reflect.ValueOf(model.TeaProgramOpts()[1]))
}

func TestWithDefaultIndex(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	for _, testcase := range []struct {
		model choose.Model
		keys  []byte
		data  string
	}{
		{
			model: *choose.NewWithStrings(items, choose.WithDefaultIndex(1)),
			keys:  []byte("\r\n"),
			data:  "Item2",
		},
	} {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write(testcase.keys)
		tm, err := tea.NewProgram(testcase.model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
		require.Nil(t, err)

		m, ok := tm.(choose.Model)
		require.Equal(t, true, ok)

		require.Equal(t, testcase.data, m.Data())
		require.Equal(t, testcase.data, m.DataString())
		require.Equal(t, true, m.Quitting())
	}
}
