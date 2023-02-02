package main

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("Input your name:").Input("Your Name")
	if err != nil {
		panic(err)
	}
	fmt.Println("Val:", val)
}
