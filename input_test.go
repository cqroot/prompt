package prompt_test

import (
	"testing"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

var InputDefaultValue = "default value"

func TestInput(t *testing.T) {
	val := `abcdefghijklmnopqrstuvwxyz1234567890-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	testcases := []StringModelTestcase{
		{Keys: []byte{}, Result: InputDefaultValue},
		{Keys: []byte(val), Result: val},
	}

	inputOptions := []input.Option{
		input.WithEchoMode(input.EchoNormal),
		input.WithEchoMode(input.EchoPassword),
		input.WithEchoMode(input.EchoNone),
	}

	for _, inputOption := range inputOptions {
		testStringModel(t,
			testcases,
			func(p *prompt.Prompt) (string, error) {
				return p.Input(InputDefaultValue, inputOption)
			},
			"?  › \x1b[7md\x1b[0mefault value",
			"?  › \x1b[7md\x1b[0mefault value"+`

enter confirm • ctrl+c quit`,
			[]byte{KeyCtrlC, KeyCtrlD},
			[]byte("\r\n"),
		)
	}
}

func TestInputWithIntegerOnly(t *testing.T) {
	testcases := []StringModelTestcase{
		{Keys: []byte("test-123.321.test.123"), Result: "123321123"},
	}

	testStringModel(t,
		testcases,
		func(p *prompt.Prompt) (string, error) {
			return p.Input(InputDefaultValue, input.WithInputMode(input.InputInteger))
		},
		"?  › \x1b[7md\x1b[0mefault value",
		"?  › \x1b[7md\x1b[0mefault value"+`

enter confirm • ctrl+c quit`,
		[]byte{KeyCtrlC, KeyCtrlD},
		[]byte("\r\n"),
	)
}

func TestInputWithNumberOnly(t *testing.T) {
	testcases := []StringModelTestcase{
		{Keys: []byte("test-123.321.test.123"), Result: "123.321123"},
	}

	testStringModel(t,
		testcases,
		func(p *prompt.Prompt) (string, error) {
			return p.Input(InputDefaultValue, input.WithInputMode(input.InputNumber))
		},
		"?  › \x1b[7md\x1b[0mefault value",
		"?  › \x1b[7md\x1b[0mefault value"+`

enter confirm • ctrl+c quit`,
		[]byte{KeyCtrlC, KeyCtrlD},
		[]byte("\r\n"),
	)
}
