package prompt

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type InputOption func(*InputModel)

type EchoMode = textinput.EchoMode

const (
	// EchoNormal displays text as is. This is the default behavior.
	EchoNormal EchoMode = textinput.EchoNormal

	// EchoPassword displays the EchoCharacter mask instead of actual
	// characters.  This is commonly used for password fields.
	EchoPassword EchoMode = textinput.EchoPassword

	// EchoNone displays nothing as characters are entered. This is commonly
	// seen for password fields on the command line.
	EchoNone EchoMode = textinput.EchoNone
)

func WithEchoMode(mode EchoMode) InputOption {
	return func(m *InputModel) {
		m.WithEchoMode(mode)
	}
}

type InputMode int

const (
	InputAll     InputMode = iota // allow any input.
	InputInteger                  // only integers can be entered.
	InputNumber                   // only integers and decimals can be entered.
)

func WithInputMode(mode InputMode) InputOption {
	return func(m *InputModel) {
		m.WithInputMode(mode)
	}
}

type ValidateFunc func(string) error

func WithValidateFunc(vf ValidateFunc) InputOption {
	return func(m *InputModel) {
		m.WithValidateFunc(vf)
	}
}
