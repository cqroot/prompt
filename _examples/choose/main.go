package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
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
	val1, err := prompt.New().Ask("Choose:").
		Choose([]string{"Item 1", "Item 2", "Item 3"})
	CheckErr(err)

	val2, err := prompt.New().Ask("Choose:").
		ChooseIndex([]string{"Item 1", "Item 2", "Item 3"})
	CheckErr(err)

	val3, err := prompt.New().Ask("Choose with Help:").
		Choose(
			[]string{"Item 1", "Item 2", "Item 3"},
			choose.WithDefaultIndex(1),
			choose.WithHelp(true),
		)
	CheckErr(err)

	fmt.Printf("{ %s }, { %d }, { %s }\n", val1, val2, val3)
}
