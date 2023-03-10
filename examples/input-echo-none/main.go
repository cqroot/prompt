package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func main() {
	val, err := prompt.New().Ask("Input your password:").
		Input("", input.WithEchoMode(input.EchoNone))
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
