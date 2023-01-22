package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
)

type data struct {
	InputVal       string
	SelectVal      string
	MultiSelectVal []string
	ToggleVal      bool
}

func main() {
	var err error
	d := data{}
	p := prompt.Default()

	d.InputVal, err = p.Ask("Input type").Input("Anonymous")
	checkErr(err)

	d.SelectVal, err = p.Ask("Select type:").Select([]string{"Taro", "Coffee", "Lychee", ""})
	checkErr(err)

	d.ToggleVal, err = p.Ask("Toggle type:").Toggle(true)
	checkErr(err)

	fmt.Printf("\ndata: %+v\n", d)
}

func checkErr(err error) {
	if err != nil {
		if errors.Is(err, prompt.ErrUserQuit) {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		} else {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}
}
