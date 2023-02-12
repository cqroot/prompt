package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt"
)

type TextAreaModelTest struct{}

func (_ TextAreaModelTest) Model() prompt.PromptModel {
	defaultVal := "default value"
	return prompt.NewTextAreaModel(defaultVal)
}

func (mt TextAreaModelTest) DataTestcases() (prompt.PromptModel, []KVPair) {
	defaultVal := "default value"
	val := `abcdefghijklmnopqrstuvwxyz1234567890-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	pm := mt.Model()
	return pm, []KVPair{
		{[]byte{}, defaultVal},
		{[]byte(val), val},
	}
}

func (mt TextAreaModelTest) ViewTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, "?  › \n┃  1 \x1b[7md\x1b[0mefault value                      \n" +
		`┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    `
}

func (mt TextAreaModelTest) ViewWithHelpTestcases() (prompt.PromptModel, string) {
	pm := mt.Model()
	return pm, "?  › \n┃  1 \x1b[7md\x1b[0mefault value                      \n" +
		`┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    

ctrl+s confirm • ctrl+c quit`
}

func TestTextAreaModel(t *testing.T) {
	testPromptModel(t, TextAreaModelTest{})
}

func TestTextArea(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	_, err := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		TextArea("")
	require.Equal(t, prompt.ErrUserQuit, err)

	_, testcases := TextAreaModelTest{}.DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte{KeyCtrlS})

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			TextArea("default value")
		require.Nil(t, err)
		require.Equal(t, testcase.Val, val)
	}
}
