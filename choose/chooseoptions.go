package choose

import (
	"fmt"
	"strings"

	"github.com/cqroot/prompt/styles"
)

type Direction int

const (
	DirectionAll Direction = iota
	DirectionH
	DirectionV
)

type Theme struct {
	Direction Direction
	View      func([]string, int) string
}

var ThemeDefault = Theme{
	Direction: DirectionV,
	View: func(choices []string, cursor int) string {
		s := strings.Builder{}
		s.WriteString("\n")

		for i := 0; i < len(choices); i++ {
			if cursor == i {
				s.WriteString(styles.DefaultSelectedItemStyle.Render(fmt.Sprintf("• %s", choices[i])))
			} else {
				s.WriteString(styles.DefaultItemStyle.Render(fmt.Sprintf("  %s", choices[i])))
			}
			s.WriteString("\n")
		}

		return s.String()
	},
}

var ThemeArrow = Theme{
	Direction: DirectionV,
	View: func(choices []string, cursor int) string {
		s := strings.Builder{}
		s.WriteString("\n")

		for i := 0; i < len(choices); i++ {
			if cursor == i {
				s.WriteString(styles.DefaultSelectedItemStyle.Render(fmt.Sprintf("❯ %s", choices[i])))
			} else {
				s.WriteString(styles.DefaultItemStyle.Render(fmt.Sprintf("  %s", choices[i])))
			}
			s.WriteString("\n")
		}

		return s.String()
	},
}

var ThemeLine = Theme{
	Direction: DirectionH,
	View: func(choices []string, cursor int) string {
		s := strings.Builder{}

		result := make([]string, len(choices))
		for index, choice := range choices {
			if index == cursor {
				result[index] = styles.DefaultSelectedItemStyle.Render(choice)
			} else {
				result[index] = styles.DefaultItemStyle.Render(choice)
			}
		}
		s.WriteString(strings.Join(result, " / "))

		return s.String()
	},
}

type Option func(*Model)

func WithTheme(theme Theme) Option {
	return func(m *Model) {
		m.theme = theme
	}
}
