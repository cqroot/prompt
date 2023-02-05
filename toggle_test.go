package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestToggle(t *testing.T) {
	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte{},
		"Yes",
	)

	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte("lhh"),
		"No",
	)

	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte("kjj"),
		"No",
	)

	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte{KeyTab},
		"No",
	)

	testPromptModel(t,
		prompt.NewToggleModel([]string{"Yes", "No"}),
		[]byte{' '},
		"No",
	)
}
