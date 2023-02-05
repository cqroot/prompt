package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestMultiChoose(t *testing.T) {
	testPromptModel(t,
		prompt.NewMultiChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		"jj k ", "", "Item 2, Item 3",
	)
}
