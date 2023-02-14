package prompt

import (
	"fmt"
	"strings"
)

type ChooseDirection int

const (
	ChooseDirectionAll ChooseDirection = iota
	ChooseDirectionH
	ChooseDirectionV
)

type ChooseTheme struct {
	Direction ChooseDirection
	View      func([]string, int) string
}

var ChooseThemeDefault = ChooseTheme{
	Direction: ChooseDirectionV,
	View: func(choices []string, cursor int) string {
		s := strings.Builder{}
		s.WriteString("\n")

		for i := 0; i < len(choices); i++ {
			if cursor == i {
				s.WriteString(DefaultSelectedItemStyle.Render(fmt.Sprintf("• %s", choices[i])))
			} else {
				s.WriteString(DefaultItemStyle.Render(fmt.Sprintf("  %s", choices[i])))
			}
			s.WriteString("\n")
		}

		return s.String()
	},
}

var ChooseThemeArrow = ChooseTheme{
	Direction: ChooseDirectionV,
	View: func(choices []string, cursor int) string {
		s := strings.Builder{}
		s.WriteString("\n")

		for i := 0; i < len(choices); i++ {
			if cursor == i {
				s.WriteString(DefaultSelectedItemStyle.Render(fmt.Sprintf("❯ %s", choices[i])))
			} else {
				s.WriteString(DefaultItemStyle.Render(fmt.Sprintf("  %s", choices[i])))
			}
			s.WriteString("\n")
		}

		return s.String()
	},
}

var ChooseThemeLine = ChooseTheme{
	Direction: ChooseDirectionH,
	View: func(choices []string, cursor int) string {
		s := strings.Builder{}

		result := make([]string, len(choices))
		for index, choice := range choices {
			if index == cursor {
				result[index] = DefaultSelectedItemStyle.Render(choice)
			} else {
				result[index] = DefaultItemStyle.Render(choice)
			}
		}
		s.WriteString(strings.Join(result, " / "))

		return s.String()
	},
}

type ChooseOption func(*ChooseModel)

func WithTheme(theme ChooseTheme) ChooseOption {
	return func(m *ChooseModel) {
		m.theme = theme
	}
}
