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
	ToggleVal_1    string
	ToggleVal_2    string
	ChooseVal      string
	MultiChooseVal []string
}

func main() {
	p := prompt.New()
	example(p)

    fmt.Println()
    fmt.Println("Example with help message")
    fmt.Println()
	p.SetHelpVisible(true)
	example(p)
}

func example(p *prompt.Prompt) {
	d := data{}
	var err error

	d.InputVal, err = p.Ask("Input example:").Input("Default string")
	checkErr(err)

	d.ToggleVal_1, err = p.Ask("Toggle example 1:").Toggle([]string{"Yes", "No"})
	checkErr(err)
	d.ToggleVal_2, err = p.Ask("Toggle example 2:").Toggle(
		[]string{"Option 1", "Option 2", "Option 3"},
	)
	checkErr(err)

	d.ChooseVal, err = p.Ask("Choose example:").Choose(
		[]string{"Option 1", "Option 2", "Option 3"},
	)
	checkErr(err)

	d.MultiChooseVal, err = p.Ask("MultiChoose example:").MultiChoose(
		[]string{"Option 1", "Option 2", "Option 3"},
	)

	fmt.Println()
	fmt.Printf("    Input        result:  %s\n", d.InputVal)
	fmt.Printf("    Toggle 1     result:  %s\n", d.ToggleVal_1)
	fmt.Printf("    Toggle 2     result:  %s\n", d.ToggleVal_2)
	fmt.Printf("    Choose       result:  %s\n", d.ChooseVal)
	fmt.Printf("    MultiChoose  result:  %s\n", strings.Join(d.MultiChooseVal, ", "))
	fmt.Println()
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
