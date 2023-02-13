package prompt_test

import (
	"strings"
	"testing"

	"github.com/cqroot/prompt"
)

func TestMultiChoose(t *testing.T) {
	testcases := []StringModelTestcase{
		{Keys: []byte{}, Result: ""},
		{Keys: []byte("kk jj "), Result: "Item 1, Item 2"},
		{Keys: []byte("kk  jj "), Result: "Item 1"},
		{Keys: []byte("kk jj  "), Result: "Item 2"},
		{Keys: []byte("kk  jj  "), Result: ""},
		{Keys: []byte{'k', 'k', ' ', KeyTab, KeyTab, ' '}, Result: "Item 1, Item 2"},
	}

	testStringModel(t,
		testcases,
		func(p *prompt.Prompt) (string, error) {
			result, err := p.MultiChoose([]string{"Item 1", "Item 2", "Item 3"})
			if err != nil {
				return "", err
			}
			return strings.Join(result, ", "), err
		},
		`?  › 
[•] Item 1
[ ] Item 2
[ ] Item 3
`,
		`?  › 
[•] Item 1
[ ] Item 2
[ ] Item 3

↑/k move up • ↓/j/tab move down • space choose • enter confirm • q/ctrl+c quit`,
		[]byte{'q', KeyCtrlC, KeyCtrlD},
		[]byte("\r\n"),
	)
}
