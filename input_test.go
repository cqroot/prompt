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

func (mt InputModelTest) DataTestcases() []KVPair {
	defaultVal := "default value"
	val := `abcdefghijklmnopqrstuvwxyz1234567890-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	return []KVPair{
		{Key: []byte{}, Val: defaultVal, View: defaultVal},
		{Key: []byte(val), Val: val, View: val},
	}
}

func (mt InputModelTest) InitViewTestcase() string {
	return "?  › \x1b[7md\x1b[0mefault value"
}

func (mt InputModelTest) InitViewWithHelpTestcase() string {
	return "?  › \x1b[7md\x1b[0mefault value" + `

enter confirm • ctrl+c quit`
}

func TestInputModel(t *testing.T) {
	testPromptModel(t, InputModelTest{})
}

func TestInput(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	_, err := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		Input("")
	require.Equal(t, prompt.ErrUserQuit, err)

	testcases := InputModelTest{}.DataTestcases()
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
