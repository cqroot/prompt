package write

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textarea textarea.Model
}

func (m Model) Data() any {
	if m.textarea.Value() == "" {
		return m.textarea.Placeholder
	} else {
		return m.textarea.Value()
	}
}

func (m Model) DataString() string {
	data := m.Data().(string)
	if strings.Contains(data, "\n") {
		return fmt.Sprintf("...(%d bytes)", len(m.Data().(string)))
	} else {
		return data
	}
}

func (m Model) KeyBindings() []key.Binding {
	return nil
}

func (m Model) UseKeyQ() bool {
	return true
}

func (m Model) UseKeyEnter() bool {
	return true
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return "\n" + m.textarea.View()
}

func New(defaultValue string) *Model {
	ti := textarea.New()
	ti.Placeholder = defaultValue
	ti.Focus()

	return &Model{
		textarea: ti,
	}
}
