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
	val1, err := prompt.New().Ask("MultiChoose:").
		MultiChoose([]string{"Item 1", "Item 2", "Item 3"})
	CheckErr(err)

	val2, err := prompt.New().Ask("MultiChoose:").
		MultiChooseIndex([]string{"Item 1", "Item 2", "Item 3"})
	CheckErr(err)

	val3, err := prompt.New().Ask("MultiChoose with Help:").
		MultiChoose(
			[]string{"Item 1", "Item 2", "Item 3"},
			multichoose.WithDefaultIndexes(1, []int{1, 2}),
			multichoose.WithHelp(true),
		)
	CheckErr(err)

	fmt.Printf("{ %s }, { %v }, { %s }\n", strings.Join(val1, ", "), val2, strings.Join(val3, ", "))
}
