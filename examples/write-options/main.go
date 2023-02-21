package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/write"
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
	val, err := prompt.New().Ask("Write:").Write(
		"Blah blah blah...",
		write.WithHelp(true),
		write.WithCharLimit(800),
	)
	CheckErr(err)

	fmt.Println(val)
}
