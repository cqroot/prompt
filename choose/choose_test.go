package choose_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/constants"
	"github.com/stretchr/testify/require"
)

func TestChoose(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	for _, testcase := range []struct {
		model choose.Model
		keys  []byte
		index int
	}{
		{
			model: *choose.NewWithStrings(items),
			keys:  []byte("\r\n"),
			index: 0,
		},
		{
			model: *choose.NewWithStrings(items),
			keys:  []byte("kkjjj\r\n"),
			index: 1,
		},
		{
			model: *choose.NewWithStrings(items),
			keys:  []byte{'k', 'k', byte(tea.KeyTab), byte(tea.KeyTab), byte(tea.KeyTab), '\r', '\n'},
			index: 1,
		},
		{
			model: *choose.NewWithStrings(items, choose.WithKeyMap(choose.HorizontalKeyMap)),
			keys:  []byte("\r\n"),
			index: 0,
		},
		{
			model: *choose.NewWithStrings(items, choose.WithKeyMap(choose.HorizontalKeyMap)),
			keys:  []byte("hhlll\r\n"),
			index: 1,
		},
		{
			model: *choose.NewWithStrings(items, choose.WithKeyMap(choose.HorizontalKeyMap)),
			keys:  []byte{'h', 'h', byte(tea.KeyTab), byte(tea.KeyTab), byte(tea.KeyTab), '\r', '\n'},
			index: 1,
		},
	} {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write(testcase.keys)
		tm, err := tea.NewProgram(testcase.model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
		require.Nil(t, err)

		m, ok := tm.(choose.Model)
		require.Equal(t, true, ok)

		require.Equal(t, testcase.index, m.Index())
		require.Equal(t, items[testcase.index], m.Data())
		require.Equal(t, items[testcase.index], m.DataString())
		require.Equal(t, true, m.Quitting())
	}
}

func TestErrors(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	in.Write([]byte("q"))
	tm, err := tea.NewProgram(*choose.NewWithStrings([]string{"item"}), tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(choose.Model)
	require.Equal(t, true, ok)

	require.Equal(t, constants.ErrUserQuit, m.Error())
}
