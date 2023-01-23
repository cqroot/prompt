package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cqroot/prompt"
)

type data struct {
	InputVal       string
	ToggleVal      bool
	SelectVal      string
	MultiSelectVal []string
}

func main() {
	var err error
	d := data{}
	p := prompt.New()

	d.InputVal, err = p.Ask("Input type:").Input("Default string")
	checkErr(err)

	d.ToggleVal, err = p.Ask("Toggle type:").Toggle(true)
	checkErr(err)

	d.SelectVal, err = p.Ask("Select type:").Select(
		[]string{"Option 1", "Option 2", "Option 3"},
	)
	checkErr(err)

	d.MultiSelectVal, err = p.Ask("MultiSelect type:").MultiSelect(
		[]string{"Option 1", "Option 2", "Option 3"},
	)

	fmt.Println()
	fmt.Printf("  Input        result:  %+v\n", d.InputVal)
	fmt.Printf("  Toggle       result:  %+v\n", d.ToggleVal)
	fmt.Printf("  Select       result:  %+v\n", d.SelectVal)
	fmt.Printf("  MultiSelect  result:  %s\n", strings.Join(d.MultiSelectVal, ", "))
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
