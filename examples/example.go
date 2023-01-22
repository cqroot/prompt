package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/cqroot/prompt"
)

type data struct {
	InputVal  string
	ToggleVal bool
	SelectVal string
}

func main() {
	var err error
	d := data{}
	p := prompt.Default()

	d.InputVal, err = p.Ask("Input type").Input("Anonymous")
	checkErr(err)

	d.ToggleVal, err = p.Ask("Toggle type:").Toggle(true)
	checkErr(err)

	d.SelectVal, err = p.Ask("Select type:").Select(
		[]string{"Option 1", "Option 2", "Option 3"},
	)
	checkErr(err)

    fmt.Println()
	fmt.Printf("  Input  result:  %+v\n", d.InputVal)
	fmt.Printf("  Toggle result:  %+v\n", d.ToggleVal)
	fmt.Printf("  Select result:  %+v\n", d.SelectVal)
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
