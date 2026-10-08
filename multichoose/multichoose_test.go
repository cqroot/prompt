package multichoose_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/multichoose"
	"github.com/stretchr/testify/require"
)

func manyItems(n int) []string {
	items := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		items = append(items, fmt.Sprintf("Item %d", i))
	}
	return items
}

// runKeys feeds the raw key bytes into the model the same way tea would, one
// KeyMsg at a time, without starting a real terminal program.
func runKeys(model multichoose.Model, height int, keys []byte) multichoose.Model {
	var tm tea.Model = model

	if height > 0 {
		tm, _ = tm.Update(tea.WindowSizeMsg{Width: 80, Height: height})
	}

	for _, k := range keys {
		var msg tea.KeyMsg
		switch k {
		case '\r', '\n':
			msg = tea.KeyMsg{Type: tea.KeyEnter}
		case ' ':
			msg = tea.KeyMsg{Type: tea.KeySpace}
		case byte(tea.KeyTab):
			msg = tea.KeyMsg{Type: tea.KeyTab}
		default:
			msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{rune(k)}}
		}
		tm, _ = tm.Update(msg)
	}

	return tm.(multichoose.Model)
}

func TestMultiChoose(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	for _, testcase := range []struct {
		model    multichoose.Model
		height   int
		keys     []byte
		view     string
		data     []string
		quitting bool
	}{
		{
			model:    *multichoose.New(items),
			keys:     []byte("\r\n"),
			data:     []string{},
			quitting: true,
		},
		{
			model:    *multichoose.New(items),
			keys:     []byte("kk jj \r\n"),
			data:     []string{"Item 1", "Item 2"},
			quitting: true,
		},
		{
			model:    *multichoose.New(items),
			keys:     []byte("kk  jj \r\n"),
			data:     []string{"Item 1"},
			quitting: true,
		},
		{
			model:    *multichoose.New(items),
			keys:     []byte("kk jj  \r\n"),
			data:     []string{"Item 2"},
			quitting: true,
		},
		{
			model:    *multichoose.New(items),
			keys:     []byte("kk  jj  \r\n"),
			data:     []string{},
			quitting: true,
		},
		{
			model:    *multichoose.New(items),
			keys:     []byte{'k', 'k', ' ', byte(tea.KeyTab), byte(tea.KeyTab), ' ', '\r', '\n'},
			data:     []string{"Item 1", "Item 2"},
			quitting: true,
		},

		// When all options fit in the terminal, the list is not scrolled.
		{
			model:  *multichoose.New(items),
			height: 10,
			view:   "\n[•] Item 1\n[ ] Item 2\n[ ] Item 3\n",
			data:   []string{},
		},

		// When the list is taller than the terminal, only a window is shown and
		// the cursor stays visible at the bottom edge.
		{
			model:  *multichoose.New(manyItems(20)),
			height: 6,
			view:   "\n[•] Item 1\n[ ] Item 2\n[ ] Item 3\n[ ] Item 4\n[ ] Item 5\n",
			data:   []string{},
		},
		{
			model:  *multichoose.New(manyItems(20)),
			height: 6,
			keys:   []byte("jjjj"),
			view:   "\n[ ] Item 1\n[ ] Item 2\n[ ] Item 3\n[ ] Item 4\n[•] Item 5\n",
			data:   []string{},
		},
		{
			model:  *multichoose.New(manyItems(20)),
			height: 6,
			keys:   []byte("jjjjj"),
			view:   "\n[ ] Item 2\n[ ] Item 3\n[ ] Item 4\n[ ] Item 5\n[•] Item 6\n",
			data:   []string{},
		},
		{
			model:  *multichoose.New(manyItems(20)),
			height: 6,
			keys:   []byte(strings.Repeat("j", 19)),
			view:   "\n[ ] Item 16\n[ ] Item 17\n[ ] Item 18\n[ ] Item 19\n[•] Item 20\n",
			data:   []string{},
		},
		{
			model:  *multichoose.New(manyItems(20)),
			height: 6,
			keys:   []byte(strings.Repeat("j", 20)),
			view:   "\n[•] Item 1\n[ ] Item 2\n[ ] Item 3\n[ ] Item 4\n[ ] Item 5\n",
			data:   []string{},
		},

		// Selected state is preserved across the visible window.
		{
			model:    *multichoose.New(manyItems(10)),
			height:   4,
			keys:     []byte(" jjjjj\r\n"),
			data:     []string{"Item 1"},
			quitting: true,
		},
	} {
		m := runKeys(testcase.model, testcase.height, testcase.keys)

		if testcase.view != "" {
			require.Equal(t, testcase.view, m.View())
		}
		require.Equal(t, testcase.data, m.Data())
		require.Equal(t, strings.Join(testcase.data, ", "), m.DataString())
		require.Equal(t, testcase.quitting, m.Quitting())
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
