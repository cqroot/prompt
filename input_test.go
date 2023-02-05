package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

type InputModelTest struct{}

func (_ InputModelTest) Model() prompt.PromptModel {
	defaultVal := "default value"
	return prompt.NewInputModel(defaultVal)
}

func (mt InputModelTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	defaultVal := "default value"
	val := "test value"

	pm := mt.Model()
	return pm, []KVPair{
		{[]byte{}, defaultVal},
		{[]byte(val), val},
	}
}

func (mt InputModelTest) ViewTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, "?  › \x1b[7md\x1b[0mefault value"
}

func (mt InputModelTest) ViewWithHelpTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, "?  › \x1b[7md\x1b[0mefault value" + `

enter confirm • q quit`
}

func TestInput(t *testing.T) {
	testPromptModel(t, InputModelTest{})
}
