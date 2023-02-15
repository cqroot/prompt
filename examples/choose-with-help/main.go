package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
)

func main() {
	val, err := prompt.New().Ask("Choose value:").Choose(
		[]string{"Item 1", "Item 2", "Item 3"},
		choose.WithHelp(true),
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
