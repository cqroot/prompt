package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestChoose(t *testing.T) {
	testcases := []StringModelTestcase{
		{Keys: []byte{}, Result: "Item 1"},
		{Keys: []byte("kkjjj"), Result: "Item 2"},
		{Keys: []byte{'k', 'k', KeyTab, KeyTab, KeyTab}, Result: "Item 2"},
	}

	testStringModel(t,
		testcases,
		func(p *prompt.Prompt) (string, error) {
			return p.Choose([]string{"Item 1", "Item 2", "Item 3"})
		},
		`?  › 
• Item 1
  Item 2
  Item 3
`,
		`?  › 
• Item 1
  Item 2
  Item 3

↑/k move up • ↓/j/tab move down • enter confirm • q/ctrl+c quit`,
		[]byte{'q', KeyCtrlC, KeyCtrlD},
		[]byte("\r\n"),
	)
}
