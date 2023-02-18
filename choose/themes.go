package choose

import (
	"fmt"
	"strings"

	"github.com/cqroot/prompt/constants"
)

type Theme func([]string, int) string

var ThemeDefault = func(choices []string, cursor int) string {
	s := strings.Builder{}
	s.WriteString("\n")

	for i := 0; i < len(choices); i++ {
		if cursor == i {
			s.WriteString(constants.DefaultSelectedItemStyle.Render(fmt.Sprintf("• %s", choices[i])))
		} else {
			s.WriteString(constants.DefaultItemStyle.Render(fmt.Sprintf("  %s", choices[i])))
		}
		s.WriteString("\n")
	}

	return s.String()
}

var ThemeArrow = func(choices []string, cursor int) string {
	s := strings.Builder{}
	s.WriteString("\n")

	for i := 0; i < len(choices); i++ {
		if cursor == i {
			s.WriteString(constants.DefaultSelectedItemStyle.Render(fmt.Sprintf("❯ %s", choices[i])))
		} else {
			s.WriteString(constants.DefaultItemStyle.Render(fmt.Sprintf("  %s", choices[i])))
		}
		s.WriteString("\n")
	}

	return s.String()
}

var ThemeLine = func(choices []string, cursor int) string {
	s := strings.Builder{}

	result := make([]string, len(choices))
	for index, choice := range choices {
		if index == cursor {
			result[index] = constants.DefaultSelectedItemStyle.Render(choice)
		} else {
			result[index] = constants.DefaultItemStyle.Render(choice)
		}
	}
	s.WriteString(strings.Join(result, " / "))
	s.WriteString("\n")

	return s.String()
}
