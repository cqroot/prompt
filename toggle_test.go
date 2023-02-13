package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestToggle(t *testing.T) {
	testcases := []StringModelTestcase{
		{Keys: []byte{}, Result: "Yes"},
		{Keys: []byte{}, Result: "Yes"},
		{Keys: []byte("lhh"), Result: "No"},
		{Keys: []byte("kjj"), Result: "No"},
		{Keys: []byte{KeyTab}, Result: "No"},
		{Keys: []byte{' '}, Result: "No"},
	}

	testStringModel(t,
		testcases,
		func(p *prompt.Prompt) (string, error) {
			return p.Toggle([]string{"Yes", "No"})
		},
		"?  › Yes / No",
		`?  › Yes / No

←/h/j move left • →/l/k/tab/space move right • enter confirm • q/ctrl+c quit`,
		[]byte{'q', KeyCtrlC, KeyCtrlD},
		[]byte("\r\n"),
	)
}
