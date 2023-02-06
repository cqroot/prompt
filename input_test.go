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

type InputModelWithIntegerLimitTest struct {
	InputModelTest
}

func (_ InputModelWithIntegerLimitTest) Model() prompt.PromptModel {
	defaultVal := "default value"
	return prompt.NewInputModel(defaultVal).SetInputLimit(prompt.InputInteger)
}

func (mt InputModelWithIntegerLimitTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	pm := mt.Model()
	return pm, []KVPair{
		{[]byte("test-123.321.test.123"), "123321123"},
	}
}

func TestInputWithIntegerLimit(t *testing.T) {
	testPromptModel(t, InputModelWithIntegerLimitTest{})
}

type InputModelWithNumberLimitTest struct {
	InputModelTest
}

func (_ InputModelWithNumberLimitTest) Model() prompt.PromptModel {
	defaultVal := "default value"
	return prompt.NewInputModel(defaultVal).SetInputLimit(prompt.InputNumber)
}

func (mt InputModelWithNumberLimitTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	pm := mt.Model()
	return pm, []KVPair{
		{[]byte("test-123.321.test.123"), "123.321123"},
	}
}

func TestInputWithNumberLimit(t *testing.T) {
	testPromptModel(t, InputModelWithNumberLimitTest{})
}
