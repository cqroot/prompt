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
	val1, err := prompt.New().Ask("Choose:").Choose(
		[]string{"Item 1", "Item 2", "Item 3"},
		choose.WithTeaProgramOpts(tea.WithInput(&in), tea.WithOutput(&out)),
	)
	CheckErr(err)

	// If the prompt and submodule set the same tea program option, the submodule's option will override the prompt's.
	val2, err := prompt.New(prompt.WithTeaProgramOpts(tea.WithOutput(&out))).Ask("Choose:").Choose(
		[]string{"Item 1", "Item 2", "Item 3"},
		choose.WithTeaProgramOpts(tea.WithOutput(os.Stderr)),
	)
	CheckErr(err)

	fmt.Printf("{ %s, %s }\n", val1, val2)
}
