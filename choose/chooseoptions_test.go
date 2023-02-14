package choose_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/choose"
)

func TestThemeDefaultView(t *testing.T) {
	require.Equal(t,
		"\n• item 1\n  item 2\n  item 3\n",
		choose.ThemeDefault.View([]string{"item 1", "item 2", "item 3"}, 0),
	)
	require.Equal(t,
		"\n  item 1\n• item 2\n  item 3\n",
		choose.ThemeDefault.View([]string{"item 1", "item 2", "item 3"}, 1),
	)
	require.Equal(t,
		"\n  item 1\n  item 2\n• item 3\n",
		choose.ThemeDefault.View([]string{"item 1", "item 2", "item 3"}, 2),
	)
}

func TestThemeArrowView(t *testing.T) {
	require.Equal(t,
		"\n❯ item 1\n  item 2\n  item 3\n",
		choose.ThemeArrow.View([]string{"item 1", "item 2", "item 3"}, 0),
	)
	require.Equal(t,
		"\n  item 1\n❯ item 2\n  item 3\n",
		choose.ThemeArrow.View([]string{"item 1", "item 2", "item 3"}, 1),
	)
	require.Equal(t,
		"\n  item 1\n  item 2\n❯ item 3\n",
		choose.ThemeArrow.View([]string{"item 1", "item 2", "item 3"}, 2),
	)
}

func TestThemeLineView(t *testing.T) {
	require.Equal(t,
		"item 1 / item 2 / item 3",
		choose.ThemeLine.View([]string{"item 1", "item 2", "item 3"}, 0),
	)
	require.Equal(t,
		"item 1 / item 2 / item 3",
		choose.ThemeLine.View([]string{"item 1", "item 2", "item 3"}, 1),
	)
	require.Equal(t,
		"item 1 / item 2 / item 3",
		choose.ThemeLine.View([]string{"item 1", "item 2", "item 3"}, 2),
	)
}
