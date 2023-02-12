package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("Input your story:").SetHelpVisible(true).
		TextArea("Bla bla bla...")
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println("Val:", val)
}
