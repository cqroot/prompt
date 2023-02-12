package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type TextAreaModel struct {
	textarea textarea.Model
}

func (m TextAreaModel) Data() any {
	if m.textarea.Value() == "" {
		return m.textarea.Placeholder
	} else {
		return m.textarea.Value()
	}
}

func (m TextAreaModel) DataString() string {
	data := m.Data().(string)
	if strings.Contains(data, "\n") {
		return fmt.Sprintf("...(%d bytes)", len(m.Data().(string)))
	} else {
		return data
	}
}

func (m TextAreaModel) KeyBindings() []key.Binding {
	return nil
}

func (m TextAreaModel) UseKeyQ() bool {
	return true
}

func (m TextAreaModel) UseKeyEnter() bool {
	return true
}

func (m TextAreaModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m TextAreaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m TextAreaModel) View() string {
	return "\n" + m.textarea.View()
}

func NewTextAreaModel(defaultValue string) *TextAreaModel {
	ti := textarea.New()
	ti.Placeholder = defaultValue
	ti.Focus()

	return &TextAreaModel{
		textarea: ti,
	}
}

func (p Prompt) TextArea(defaultValue string) (string, error) {
	pm := NewTextAreaModel(defaultValue)

	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}
