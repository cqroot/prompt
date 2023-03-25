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
	val, err := prompt.New().Ask("Theme Dot:").MultiChoose(
		[]string{"Item 1", "Item 2", "Item 3"},
		multichoose.WithTheme(multichoose.ThemeDot),
	)
	CheckErr(err)

	fmt.Printf("{ %s }\n", strings.Join(val, ", "))
}
