package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestChoose(t *testing.T) {
	testModel(t,
		prompt.NewChooseModel([]string{"Item 1", "Item 2", "Item 3"}),
		"jjjj", "Item 1", "Item 2",
	)
}
