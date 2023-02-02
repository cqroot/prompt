package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestToggle(t *testing.T) {
	testModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		"jjj", "Yes", "No",
	)
}
