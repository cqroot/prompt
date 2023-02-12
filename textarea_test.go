package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt"
)

type TextAreaModelTest struct {
	defaultVal string
}

func NewTextAreaModelTest() *TextAreaModelTest {
	return &TextAreaModelTest{
		defaultVal: "default value",
	}
}

func (mt TextAreaModelTest) Model() prompt.PromptModel {
	return prompt.NewTextAreaModel(mt.defaultVal)
}

func (mt TextAreaModelTest) DataTestcases() []KVPair {
	val := `abcdefghijklmnopqrstuvwxyz1234567890-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	return []KVPair{
		{Key: []byte{}, Val: mt.defaultVal, View: mt.defaultVal},
		{Key: []byte(val), Val: val, View: val},
		{Key: []byte("test\r\naaa"), Val: "test\naaa", View: "...(8 bytes)"},
	}
}

func (mt TextAreaModelTest) InitViewTestcase() string {
	return "?  › \n┃  1 \x1b[7md\x1b[0mefault value                      \n" +
		`┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    `
}

func (mt TextAreaModelTest) InitViewWithHelpTestcase() string {
	return "?  › \n┃  1 \x1b[7md\x1b[0mefault value                      \n" +
		`┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    

ctrl+s confirm • ctrl+c quit`
}

func TestTextAreaModel(t *testing.T) {
	testPromptModel(t, NewTextAreaModelTest())
}

func TestTextArea(t *testing.T) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write([]byte{KeyCtrlC})

	_, err := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		TextArea("")
	require.Equal(t, prompt.ErrUserQuit, err)

	testcases := NewTextAreaModelTest().DataTestcases()
	for _, testcase := range testcases {
		in.Reset()
		in.Write(testcase.Key)
		in.Write([]byte{KeyCtrlS})

		val, err := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
			TextArea(NewTextAreaModelTest().defaultVal)
		require.Nil(t, err)
		require.Equal(t, testcase.Val, val)
	}
}
