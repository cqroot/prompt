package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("Toggle value:").Toggle([]string{"Yes", "No"})
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println("Val:", val)
}
