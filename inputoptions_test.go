package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt"
	"github.com/stretchr/testify/require"
)

type InputModelWithIntegerLimitTest struct {
	InputModelTest
}

func (_ InputModelWithIntegerLimitTest) Model() prompt.PromptModel {
	defaultVal := "default value"
	return prompt.NewInputModel(defaultVal).WithInputMode(prompt.InputInteger)
}

func (mt InputModelWithIntegerLimitTest) DataTestcases() []KVPair {
	return []KVPair{
		{Key: []byte("test-123.321.test.123"), Val: "123321123", View: "123321123"},
	}
}

func TestInputModelWithIntegerLimit(t *testing.T) {
	testPromptModel(t, InputModelWithIntegerLimitTest{})
}

type InputModelWithNumberLimitTest struct {
	InputModelTest
}

func (_ InputModelWithNumberLimitTest) Model() prompt.PromptModel {
	defaultVal := "default value"
	return prompt.NewInputModel(defaultVal).WithInputMode(prompt.InputNumber)
}

func (mt InputModelWithNumberLimitTest) DataTestcases() []KVPair {
	return []KVPair{
		{Key: []byte("test-123.321.test.123"), Val: "123.321123", View: "123.321123"},
	}
}

func TestInputModelWithNumberLimit(t *testing.T) {
	testPromptModel(t, InputModelWithNumberLimitTest{})
}

func TestInputWithInputInteger(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	testcases := InputModelWithIntegerLimitTest{}.DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte("\r\n"))

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			Input("default value", prompt.WithInputMode(prompt.InputInteger))
		require.Nil(t, err)
		require.Equal(t, testcase.Val, val)
	}
}

func TestInputWithInputNumber(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	testcases := InputModelWithNumberLimitTest{}.DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte("\r\n"))

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			Input("default value", prompt.WithInputMode(prompt.InputNumber))
		require.Nil(t, err)
		require.Equal(t, testcase.Val, val)
	}
}
