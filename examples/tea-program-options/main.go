package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
)

func CheckErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}

func main() {
	var in bytes.Buffer
	var out bytes.Buffer

	in.Write([]byte("j\r\n"))

	val, err := prompt.New().Ask("Choose:").Choose(
		[]string{"Item 1", "Item 2", "Item 3"},
		choose.WithTeaProgramOpts(tea.WithInput(&in), tea.WithOutput(&out)),
	)
	CheckErr(err)

	fmt.Printf("{ %s }\n", val)
}
