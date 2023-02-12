package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cqroot/prompt"
)

func main() {
	val, err := prompt.New().Ask("MultiChoose value:").WithHelp(true).
		MultiChoose([]string{"Item 1", "Item 2", "Item 3"})
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
	fmt.Println("Val:", strings.Join(val, ", "))
}
