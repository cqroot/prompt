package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestChoose(t *testing.T) {
	testPromptModel(t,
		prompt.NewChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		[]byte{},
		"Item 1",
	)

	testPromptModel(t,
		prompt.NewChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		[]byte("kkjjj"),
		"Item 2",
	)

	testPromptModel(t,
		prompt.NewChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		[]byte{'k', 'k', KeyTab, KeyTab, KeyTab},
		"Item 2",
	)
}
