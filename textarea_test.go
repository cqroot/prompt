package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
)

func TestTextArea(t *testing.T) {
	val := `abcdefghijklmnopqrstuvwxyz1234567890-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	testcases := []StringModelTestcase{
		{Keys: []byte{}, Result: InputDefaultValue},
		{Keys: []byte(val), Result: val},
		{Keys: []byte("test\r\naaa"), Result: "test\naaa"},
	}

	testStringModel(t,
		testcases,
		func(p *prompt.Prompt) (string, error) {
			return p.TextArea(InputDefaultValue)
		},
		"?  › \n┃  1 \x1b[7md\x1b[0mefault value                      \n"+
			`┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    `,
		"?  › \n┃  1 \x1b[7md\x1b[0mefault value                      \n"+
			`┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    
┃  ~                                    

ctrl+s confirm • ctrl+c quit`,
		[]byte{KeyCtrlC, KeyCtrlD},
		[]byte{KeyCtrlS},
	)
}
