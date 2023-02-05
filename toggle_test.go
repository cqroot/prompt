package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestToggle(t *testing.T) {
	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte("lhh"),
		"Yes", "No",
	)

	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte("kjj"),
		"Yes", "No",
	)

	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte{KeyTab},
		"Yes", "No",
	)

	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte{' '},
		"Yes", "No",
	)
}
