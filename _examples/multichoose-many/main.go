package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/multichoose"
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

func main() {
	choices := make([]string, 0, 256)
	for i := 1; i <= 256; i++ {
		choices = append(choices, fmt.Sprintf("Item %d", i))
	}

	val1, err := prompt.New().Ask("MultiChoose (256 items):").
		MultiChoose(choices)
	CheckErr(err)

	val2, err := prompt.New().Ask("MultiChoose with Help (256 items):").
		MultiChoose(choices, multichoose.WithHelp(true), multichoose.WithTheme(multichoose.ThemeDot))
	CheckErr(err)

	fmt.Printf("{ %s }\n{ %s }\n", strings.Join(val1, ", "), strings.Join(val2, ", "))
}
