package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt"
)

type InputModelTest struct{}

func (_ InputModelTest) Model() prompt.PromptModel {
	defaultVal := "default value"
	return prompt.NewInputModel(defaultVal)
}

func (mt InputModelTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	defaultVal := "default value"
	val := `abcdefghijklmnopqrstuvwxyz1234567890-=~!@#$%^&*()_+[]\{}|;':",./<>?`

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

enter confirm • ctrl+c quit`
}

func TestInputModel(t *testing.T) {
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

func TestInputModelWithIntegerLimit(t *testing.T) {
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

func TestInputModelWithNumberLimit(t *testing.T) {
	testPromptModel(t, InputModelWithNumberLimitTest{})
}

func TestInput(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	_, err := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		Input("")
	require.Equal(t, prompt.ErrUserQuit, err)

	_, testcases := InputModelTest{}.DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte("\r\n"))

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			Input("default value")
		require.Nil(t, err)
		require.Equal(t, testcase.Val, val)
	}
}
