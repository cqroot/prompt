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
	ChooseVal      string
	MultiChooseVal []string
}

func main() {
	var err error
	d := data{}
	p := prompt.New()

	d.InputVal, err = p.Ask("Input type:").Input("Default string")
	checkErr(err)

	d.ToggleVal, err = p.Ask("Toggle type:").Toggle(true)
	checkErr(err)

	d.ChooseVal, err = p.Ask("Choose type:").Choose(
		[]string{"Option 1", "Option 2", "Option 3"},
	)
	checkErr(err)

	d.MultiChooseVal, err = p.Ask("MultiChoose type:").MultiChoose(
		[]string{"Option 1", "Option 2", "Option 3"},
	)

	fmt.Println()
	fmt.Printf("  Input        result:  %+v\n", d.InputVal)
	fmt.Printf("  Toggle       result:  %+v\n", d.ToggleVal)
	fmt.Printf("  Choose       result:  %+v\n", d.ChooseVal)
	fmt.Printf("  MultiChoose  result:  %s\n", strings.Join(d.MultiChooseVal, ", "))
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
