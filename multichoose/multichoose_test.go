package multichoose_test

import (
	"bytes"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/multichoose"
	"github.com/stretchr/testify/require"
)

func TestMultiChoose(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	for _, testcase := range []struct {
		model multichoose.Model
		keys  []byte
		index []int
		data  []string
	}{
		{
			model: *multichoose.New(items),
			keys:  []byte("\r\n"),
			index: []int{},
			data:  []string{},
		},
		{
			model: *multichoose.New(items),
			keys:  []byte("kk jj \r\n"),
			index: []int{0, 1},
			data:  []string{"Item 1", "Item 2"},
		},
		{
			model: *multichoose.New(items),
			keys:  []byte("kk  jj \r\n"),
			index: []int{0},
			data:  []string{"Item 1"},
		},
		{
			model: *multichoose.New(items),
			keys:  []byte("kk jj  \r\n"),
			index: []int{1},
			data:  []string{"Item 2"},
		},
		{
			model: *multichoose.New(items),
			keys:  []byte("kk  jj  \r\n"),
			index: []int{},
			data:  []string{},
		},
		{
			model: *multichoose.New(items),
			keys:  []byte{'k', 'k', ' ', byte(tea.KeyTab), byte(tea.KeyTab), ' ', '\r', '\n'},
			index: []int{0, 1},
			data:  []string{"Item 1", "Item 2"},
		},
	} {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write(testcase.keys)
		tm, err := tea.NewProgram(testcase.model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
		require.Nil(t, err)

		m, ok := tm.(multichoose.Model)
		require.Equal(t, true, ok)

		require.Equal(t, testcase.index, m.Index())
		require.Equal(t, testcase.data, m.Data())
		require.Equal(t, strings.Join(testcase.data, ", "), m.DataString())
		require.Equal(t, true, m.Quitting())
	}
}

func TestErrors(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	in.Write([]byte("q"))
	tm, err := tea.NewProgram(*multichoose.New([]string{"item"}), tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(multichoose.Model)
	require.Equal(t, true, ok)

	require.Equal(t, constants.ErrUserQuit, m.Error())
}
