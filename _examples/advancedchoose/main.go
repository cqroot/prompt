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
		AdvancedChoose([]choose.Choice{
			{Text: "Item 1", Note: "The note for item 1"},
			{Text: "Another item", Note: "The note for item 2"},
			{Text: "Item 3", Note: "The note for item 3"},
		})
	CheckErr(err)

	val2, err := prompt.New().Ask("Choose:").
		AdvancedChooseIndex([]choose.Choice{
			{Text: "Item 1", Note: "The note for item 1"},
			{Text: "Another item", Note: "The note for item 2"},
			{Text: "Item 3", Note: "The note for item 3"},
		})
	CheckErr(err)

	val3, err := prompt.New().Ask("Choose:").
		AdvancedChoose(
			[]choose.Choice{
				{Text: "Item 1", Note: "The note for item 1"},
				{Text: "Another item", Note: "The note for item 2"},
				{Text: "Item 3", Note: "The note for item 3"},
			},
			choose.WithHelp(true),
		)
	CheckErr(err)

	fmt.Printf("{ %s }, { %d }, { %s }\n", val1, val2, val3)
}
