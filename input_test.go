package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestInput(t *testing.T) {
	defaultVal := "default value"
	val := "test value"

	testPromptModel(t, prompt.NewInputModel(defaultVal),
		[]byte{}, defaultVal,
	)

	testPromptModel(t, prompt.NewInputModel(defaultVal),
		[]byte(val), val,
	)
}
