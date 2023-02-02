package main

import (
	"fmt"
	"strings"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("MultiChoose value:").
		MultiChoose([]string{"Item 1", "Item 2", "Item 3"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Val:", strings.Join(val, ", "))
}
