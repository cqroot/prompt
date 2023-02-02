package main

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("Choose value:").
		Choose([]string{"Item 1", "Item 2", "Item 3"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Val:", val)
}
