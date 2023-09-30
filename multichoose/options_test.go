package multichoose_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/multichoose"
)

func TestWithTeaProgramOpts(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	withInput := tea.WithInput(&in)
	withOutput := tea.WithOutput(&out)

	model := multichoose.New(
		[]string{"Item 1", "Item 2", "Item 3"},
		multichoose.WithTeaProgramOpts(withInput, withOutput),
	)

	require.True(t, reflect.ValueOf(withInput) == reflect.ValueOf(model.TeaProgramOpts()[0]))
	require.True(t, reflect.ValueOf(withOutput) == reflect.ValueOf(model.TeaProgramOpts()[1]))
}

func TestWithDefaultIndexes(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	for _, testcase := range []struct {
		model multichoose.Model
		keys  []byte
		data  []string
	}{
		{
			model: *multichoose.New(items, multichoose.WithDefaultIndexes(1, []int{1, 2})),
			keys:  []byte("\r\n"),
			data:  []string{"Item 2", "Item 3"},
		},
	} {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write(testcase.keys)
		tm, err := tea.NewProgram(testcase.model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
		require.Nil(t, err)

		m, ok := tm.(multichoose.Model)
		require.Equal(t, true, ok)

		require.Equal(t, testcase.data, m.Data())
		require.Equal(t, strings.Join(testcase.data, ", "), m.DataString())
		require.Equal(t, true, m.Quitting())
	}
}
