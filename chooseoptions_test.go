package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
	"github.com/stretchr/testify/require"
)

func TestChooseThemeDefaultView(t *testing.T) {
	require.Equal(t,
		"\n• item 1\n  item 2\n  item 3\n",
		prompt.ChooseThemeDefault.View([]string{"item 1", "item 2", "item 3"}, 0),
	)
	require.Equal(t,
		"\n  item 1\n• item 2\n  item 3\n",
		prompt.ChooseThemeDefault.View([]string{"item 1", "item 2", "item 3"}, 1),
	)
	require.Equal(t,
		"\n  item 1\n  item 2\n• item 3\n",
		prompt.ChooseThemeDefault.View([]string{"item 1", "item 2", "item 3"}, 2),
	)
}

func TestChooseThemeArrowView(t *testing.T) {
	require.Equal(t,
		"\n❯ item 1\n  item 2\n  item 3\n",
		prompt.ChooseThemeArrow.View([]string{"item 1", "item 2", "item 3"}, 0),
	)
	require.Equal(t,
		"\n  item 1\n❯ item 2\n  item 3\n",
		prompt.ChooseThemeArrow.View([]string{"item 1", "item 2", "item 3"}, 1),
	)
	require.Equal(t,
		"\n  item 1\n  item 2\n❯ item 3\n",
		prompt.ChooseThemeArrow.View([]string{"item 1", "item 2", "item 3"}, 2),
	)
}

func TestChooseThemeLineView(t *testing.T) {
	require.Equal(t,
		"item 1 / item 2 / item 3",
		prompt.ChooseThemeLine.View([]string{"item 1", "item 2", "item 3"}, 0),
	)
	require.Equal(t,
		"item 1 / item 2 / item 3",
		prompt.ChooseThemeLine.View([]string{"item 1", "item 2", "item 3"}, 1),
	)
	require.Equal(t,
		"item 1 / item 2 / item 3",
		prompt.ChooseThemeLine.View([]string{"item 1", "item 2", "item 3"}, 2),
	)
}
