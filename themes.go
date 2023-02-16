package prompt

import (
	"strings"

	"github.com/cqroot/prompt/styles"
)

type State int

const (
	StateNormal State = iota
	StateFinish
	StateError
)

type Theme func(string, State, string) string

func ThemeDefault(msg string, state State, model string) string {
	s := strings.Builder{}

	switch state {
	case StateNormal:
		s.WriteString(styles.DefaultFinishPromptPrefixStyle.Render("?"))
	case StateFinish:
		s.WriteString(styles.DefaultFinishPromptPrefixStyle.Render("✔"))
	case StateError:
		s.WriteString(styles.DefaultErrorPromptPrefixStyle.Render("✖"))
	}

	s.WriteString(" ")
	s.WriteString(msg)
	s.WriteString(" ")

	if state == StateNormal {
		s.WriteString(styles.DefaultNormalPromptSuffixStyle.Render("›"))
		s.WriteString(" ")
		s.WriteString(model)
	} else {
		s.WriteString(styles.DefaultFinishPromptSuffixStyle.Render("…"))
		s.WriteString(" ")
		s.WriteString(model)
		s.WriteString("\n")
	}

	return s.String()
}
