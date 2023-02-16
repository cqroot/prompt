package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
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
	val1, err := prompt.New().Ask("Input:").Input("Blah blah")
	CheckErr(err)

	val2, err := prompt.New().Ask("Input with Help:").Input(
		"Blah blah",
		input.WithHelp(true),
	)
	CheckErr(err)

	fmt.Printf("{ %s }, { %s }\n", val1, val2)
}
