package main

import (
	"errors"
	"fmt"

	"github.com/cqroot/prompt"
)

type data struct {
    Name string
    BubbleTea string
    Confirm bool
}

func main() {
    var err error
    d := data{}

	d.Name, err = prompt.Input("What's your name?", "Anonymous")
	if err != nil {
		panic(err)
	}

	d.BubbleTea, err = prompt.Select(
		"What kind of Bubble Tea would you like to order?",
		[]string{"Taro", "Coffee", "Lychee", ""},
	)
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Println("You have no choice.")
			return
		} else {
			panic(err)
		}
	}

	d.Confirm, err = prompt.Toggle("Can you confirm?", true)
	if err != nil {
		panic(err)
	}

    fmt.Printf("\ndata: %+v\n", d)
}
