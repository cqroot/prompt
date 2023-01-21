package main

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func main() {
	choice, err := prompt.Select(
		"What kind of Bubble Tea would you like to order?",
		[]string{"Taro", "Coffee", "Lychee", ""},
	)
	if err != nil {
		panic(err)
	}
	if choice == "" {
		fmt.Println("You have no choice.")
		return
	}

	fmt.Printf("You chose %s!\n", choice)
}
