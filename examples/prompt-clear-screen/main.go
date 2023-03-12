package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
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
	val, err := prompt.New(
		prompt.WithTheme(prompt.ThemeDefaultClear),
	).Ask("Clear screen after selection:").Choose([]string{"Item 1", "Item 2", "Item 3"})
	CheckErr(err)

	fmt.Printf("{ %s }\n", val)
}
