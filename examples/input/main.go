package main

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func main() {
	name, err := prompt.Input("What's your name?", "Anonymous")
	if err != nil {
		panic(err)
	}
    fmt.Println("Your name is", name)
}
