package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/multichoose"
)

func main() {
	val, err := prompt.New().Ask("MultiChoose value:").MultiChoose(
		[]string{"Item 1", "Item 2", "Item 3"},
		multichoose.WithTheme(multichoose.ThemeDot),
	)
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println(strings.Join(val, ", "))
}
