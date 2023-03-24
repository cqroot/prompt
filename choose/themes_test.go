package choose_test

import (
	"testing"

	"github.com/cqroot/prompt/choose"
	"github.com/stretchr/testify/require"
)

func TestThemes(t *testing.T) {
	items := []string{"Item 1", "Item 2", "Item 3"}
	choices := []choose.Choice{
		{Text: "Item 1", Note: "The note for item 1"},
		{Text: "Another item", Note: "The note for item 2"},
		{Text: "Item 3", Note: "The note for item 3"},
	}

	for _, testcase := range []struct {
		model choose.Model
		view  string
	}{
		{
			model: *choose.NewWithStrings(items),
			view:  "\n• Item 1\n  Item 2\n  Item 3\n",
		},
		{
			model: *choose.NewWithStrings(items, choose.WithHelp(true)),
			view:  "\n• Item 1\n  Item 2\n  Item 3\n\n? toggle help • enter confirm • q quit",
		},
		{
			model: *choose.New(choices),
			view: `
• Item 1        The note for item 1
  Another item  The note for item 2
  Item 3        The note for item 3
`,
		},
		{
			model: *choose.NewWithStrings(items, choose.WithTheme(choose.ThemeArrow)),
			view:  "\n❯ Item 1\n  Item 2\n  Item 3\n",
		},
		{
			model: *choose.New(choices, choose.WithTheme(choose.ThemeArrow)),
			view: `
❯ Item 1        The note for item 1
  Another item  The note for item 2
  Item 3        The note for item 3
`,
		},
		{
			model: *choose.NewWithStrings(items, choose.WithTheme(choose.ThemeLine)),
			view:  "Item 1 / Item 2 / Item 3\n",
		},
		{
			model: *choose.NewWithStrings(items, choose.WithTheme(choose.ThemeLine), choose.WithHelp(true)),
			view:  "Item 1 / Item 2 / Item 3\n\n? toggle help • enter confirm • q quit",
		},
		{
			model: *choose.New(choices, choose.WithTheme(choose.ThemeLine)),
			view:  "Item 1 / Another item / Item 3\n",
		},
	} {
		require.Equal(t, testcase.view, testcase.model.View())
	}
}
