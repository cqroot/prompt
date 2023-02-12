package prompt

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InputModel struct {
	df                string
	textInput         textinput.Model
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
	inputMode         InputMode
}

func (m InputModel) Data() any {
	if m.textInput.Value() == "" {
		return m.textInput.Placeholder
	} else {
		return m.textInput.Value()
	}
}

func (m InputModel) DataString() string {
	if m.textInput.EchoMode == EchoNormal {
		return m.Data().(string)
	}
	m.textInput.Blur()
	str := m.textInput.View()
	m.textInput.Focus()
	return str
}

// Deprecated: use InputModel.WithInputMode instead.
func (m *InputModel) SetInputLimit(inputLimit InputMode) *InputModel {
	m.inputMode = inputLimit
	return m
}

func (m *InputModel) WithInputMode(mode InputMode) *InputModel {
	m.inputMode = mode
	return m
}

func (m *InputModel) WithEchoMode(mode EchoMode) *InputModel {
	m.textInput.EchoMode = mode
	return m
}

func (m InputModel) KeyBindings() []key.Binding {
	return nil
}

func (m InputModel) UseKeyQ() bool {
	return true
}

func (m InputModel) UseKeyEnter() bool {
	return false
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.inputMode == InputNumber || m.inputMode == InputInteger {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			keypress := msg.String()
			if len(keypress) == 1 {
				if keypress == "." {
					if m.inputMode != InputNumber ||
						strings.Contains(m.textInput.Value(), ".") {
						return m, nil
					}
				} else {
					if !unicode.IsNumber([]rune(keypress)[0]) {
						return m, nil
					}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {
	return m.textInput.View()
}

func NewInputModel(defaultValue string) *InputModel {
	ti := textinput.New()
	ti.Placeholder = defaultValue
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	ti.Prompt = ""

	m := InputModel{
		textInput: ti,
		df:        defaultValue,
		inputMode: InputAll,

		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

// Deprecated: use InputModel.Input("", prompt.WithInputMode()) instead.
// Input asks the user to enter a string. It restricts the types of characters
// the user can enter.
func (p Prompt) InputWithLimit(defaultValue string, inputLimit InputMode) (string, error) {
	pm := NewInputModel(defaultValue)
	pm.inputMode = inputLimit
	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}

// Input asks the user to enter a string.
func (p Prompt) Input(defaultValue string, opts ...InputOption) (string, error) {
	pm := NewInputModel(defaultValue)

	for _, opt := range opts {
		opt(pm)
	}

	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}
