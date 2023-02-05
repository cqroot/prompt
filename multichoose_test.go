package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestMultiChoose(t *testing.T) {
	testPromptModel(t,
		prompt.NewMultiChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		[]byte{},
		"",
	)

	testPromptModel(t,
		prompt.NewMultiChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		[]byte("kk jj "),
		"Item 1, Item 2",
	)

	testPromptModel(t,
		prompt.NewMultiChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		[]byte{'k', 'k', ' ', KeyTab, KeyTab, ' '},
		"Item 1, Item 2",
	)
}
