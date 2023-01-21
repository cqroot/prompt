package main

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func main() {
	choice, err := prompt.Toggle("Can you confirm?", true)
	if err != nil {
		panic(err)
	}

	if choice {
		fmt.Println("You have confirmed")
	} else {
		fmt.Println("You have declined")
	}
}
