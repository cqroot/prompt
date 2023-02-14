package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("Choose value:").Choose(
		[]string{"Item 1", "Item 2", "Item 3"},
		prompt.WithTheme(prompt.ChooseThemeArrow),
	)
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println(val)
}
