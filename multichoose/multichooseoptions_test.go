package multichoose_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/multichoose"
)

func isSelected(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func TestThemeDefaultView(t *testing.T) {
	require.Equal(t,
		"\n[•] item 1\n[x] item 2\n[ ] item 3\n",
		multichoose.ThemeDefault([]string{"item 1", "item 2", "item 3"}, 0, isSelected),
	)

	require.Equal(t,
		"\n[ ] item 1\n[x] item 2\n[ ] item 3\n",
		multichoose.ThemeDefault([]string{"item 1", "item 2", "item 3"}, 1, isSelected),
	)

	require.Equal(t,
		"\n[ ] item 1\n[x] item 2\n[•] item 3\n",
		multichoose.ThemeDefault([]string{"item 1", "item 2", "item 3"}, 2, isSelected),
	)
}

func TestThemeDotView(t *testing.T) {
	require.Equal(t,
		"\n○ item 1\n◉ item 2\n○ item 3\n",
		multichoose.ThemeDot([]string{"item 1", "item 2", "item 3"}, 0, isSelected),
	)

	require.Equal(t,
		"\n○ item 1\n◉ item 2\n○ item 3\n",
		multichoose.ThemeDot([]string{"item 1", "item 2", "item 3"}, 1, isSelected),
	)

	require.Equal(t,
		"\n○ item 1\n◉ item 2\n○ item 3\n",
		multichoose.ThemeDot([]string{"item 1", "item 2", "item 3"}, 2, isSelected),
	)
}
