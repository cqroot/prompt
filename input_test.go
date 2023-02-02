package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestInput(t *testing.T) {
	defaultVal := "default value"
	val := "test value"

	testModel(t, prompt.NewInputModel(defaultVal), val, defaultVal, val)
}
