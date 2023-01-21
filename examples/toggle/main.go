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
		fmt.Println("you have confirmed")
	} else {
		fmt.Println("you have declined")
	}
}
