package multichoose_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/multichoose"
	"github.com/stretchr/testify/require"
)

func TestThemes(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}

	for _, testcase := range []struct {
		model multichoose.Model
		view  string
	}{
		{
			model: *multichoose.New(items),
			view:  "\n[•] Item 1\n[ ] Item 2\n[ ] Item 3\n",
		},
		{
			model: func() multichoose.Model {
				var tm tea.Model
				tm = *multichoose.New(items)
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeySpace})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeySpace})
				return tm.(multichoose.Model)
			}(),
			view: "\n[ ] Item 1\n[x] Item 2\n[x] Item 3\n",
		},
		{
			model: *multichoose.New(items, multichoose.WithHelp(true)),
			view: `
[•] Item 1
[ ] Item 2
[ ] Item 3

? toggle help • space choose • enter confirm • q quit`,
		},
		{
			model: *multichoose.New(items, multichoose.WithTheme(multichoose.ThemeDot)),
			view:  "\n○ Item 1\n○ Item 2\n○ Item 3\n",
		},
		{
			model: func() multichoose.Model {
				var tm tea.Model
				tm = *multichoose.New(items, multichoose.WithTheme(multichoose.ThemeDot))
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeySpace})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeySpace})
				return tm.(multichoose.Model)
			}(),
			view: "\n○ Item 1\n◉ Item 2\n◉ Item 3\n",
		},
	} {
		require.Equal(t, testcase.view, testcase.model.View())
	}
}
