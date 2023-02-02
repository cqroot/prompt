package main

import (
	"fmt"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("Toggle value:").Toggle([]string{"Yes", "No"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Val:", val)
}
