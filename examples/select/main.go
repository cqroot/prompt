package main

import (
	"errors"
	"fmt"

	"github.com/cqroot/prompt"
)

func main() {
	choice, err := prompt.Select(
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

	fmt.Printf("You chose %s!\n", choice)
}
