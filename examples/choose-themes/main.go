package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
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
	val1, err := prompt.New().Ask("Theme Arrow:").Choose(
		[]string{"Item 1", "Item 2", "Item 3"},
		choose.WithTheme(choose.ThemeArrow),
	)
	CheckErr(err)

	val2, err := prompt.New().Ask("Theme Line:").Choose(
		[]string{"Item 1", "Item 2", "Item 3"},
		choose.WithTheme(choose.ThemeLine),
	)
	CheckErr(err)

	fmt.Printf("{ %s }, { %s }\n", val1, val2)
}
