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
	val1, err := prompt.New().Ask("Write:").Write("Blah blah blah...")
	CheckErr(err)

	val2, err := prompt.New().Ask("Write with Help:").Write(
		"Blah blah blah...",
		write.WithHelp(true),
	)
	CheckErr(err)

	fmt.Println(val1)
	fmt.Println("====================")
	fmt.Println(val2)
}
