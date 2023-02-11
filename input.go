package prompt

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InputLimit int

const (
	InputAll     InputLimit = iota // allow any input.
	InputInteger                   // only integers can be entered.
	InputNumber                    // only integers and decimals can be entered.
)

type InputModel struct {
	df                string
	textInput         textinput.Model
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
	inputLimit        InputLimit
}

func (m InputModel) Data() any {
	return m.DataString()
}

func (m InputModel) DataString() string {
	if m.textInput.Value() == "" {
		return m.textInput.Placeholder
	} else {
		return m.textInput.Value()
	}
}

func (m *InputModel) SetInputLimit(inputLimit InputLimit) *InputModel {
	m.inputLimit = inputLimit
	return m
}

func (m InputModel) KeyBindings() []key.Binding {
	return nil
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.inputLimit == InputNumber || m.inputLimit == InputInteger {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			keypress := msg.String()
			if len(keypress) == 1 {
				if keypress == "." {
					if m.inputLimit != InputNumber ||
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
		textInput:  ti,
		df:         defaultValue,
		inputLimit: InputAll,

		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

// Input asks the user to enter a string. It restricts the types of characters
// the user can enter.
func (p Prompt) InputWithLimit(defaultValue string, inputLimit InputLimit) (string, error) {
	pm := NewInputModel(defaultValue)
	pm.inputLimit = inputLimit
	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.DataString(), nil
}

// Input asks the user to enter a string. You can use InputWithLimit to limit
// what the user can enter.
func (p Prompt) Input(defaultValue string) (string, error) {
	return p.InputWithLimit(defaultValue, InputAll)
}
