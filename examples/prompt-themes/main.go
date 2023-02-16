package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/constants"
)

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

func EmojiTheme(msg string, state prompt.State, model string) string {
	s := strings.Builder{}

	switch state {
	case prompt.StateNormal:
		s.WriteString(constants.DefaultFinishPromptPrefixStyle.Render("🤔 "))
	case prompt.StateFinish:
		s.WriteString(constants.DefaultFinishPromptPrefixStyle.Render("😃 "))
	case prompt.StateError:
		s.WriteString(constants.DefaultErrorPromptPrefixStyle.Render("😡 "))
	}

	s.WriteString(" ")
	s.WriteString(msg)
	s.WriteString(" ")

	if state == prompt.StateNormal {
		s.WriteString(constants.DefaultNormalPromptSuffixStyle.Render("›"))
		s.WriteString(" ")
		s.WriteString(model)
	} else {
		s.WriteString(constants.DefaultFinishPromptSuffixStyle.Render("…"))
		s.WriteString(" ")
		s.WriteString(model)
		s.WriteString("\n")
	}

	return s.String()
}

func main() {
	val, err := prompt.New(prompt.WithTheme(EmojiTheme)).
		Ask("Emoji Theme:").Input("Blah blah")
	CheckErr(err)

	fmt.Printf("{ %s }\n", val)

	_, err = prompt.New(prompt.WithTheme(EmojiTheme)).
		Ask("Emoji Theme:").Input("Blah blah")
	CheckErr(err)
}
