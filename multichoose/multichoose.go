package multichoose

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choice  uint64
	choices []string
	cursor  int
	keys    []key.Binding
	theme   Theme
}

func (m Model) Data() any {
	result := make([]string, 0)

	for i := 0; i < len(m.choices); i++ {
		if m.isSelected(i) {
			result = append(result, m.choices[i])
		}
	}
	return result
}

func (m Model) DataString() string {
	return strings.Join(m.Data().([]string), ", ")
}

func (m Model) KeyBindings() []key.Binding {
	return m.keys
}

func (m Model) UseKeyQ() bool {
	return false
}

func (m Model) UseKeyEnter() bool {
	return false
}

func (m *Model) toggleChoice(index int) {
	i := uint64(index)
	if m.isSelected(index) {
		m.deselectItem(i)
	} else {
		m.selectItem(i)
	}
}

func (m Model) isSelected(index int) bool {
	i := uint64(index)
	curr := uint64(1) << i
	if (m.choice & curr) != 0 {
		return true
	} else {
		return false
	}
}

func (m *Model) selectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice | curr
}

func (m *Model) deselectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice & (^curr)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			m.toggleChoice(m.cursor)

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}

		case "down", "j", "tab":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.theme(m.choices, m.cursor, m.isSelected)
}

func New(choices []string, opts ...Option) *Model {
	multiChooseKeys := []key.Binding{
		key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		key.NewBinding(
			key.WithKeys("down", "j", "tab"),
			key.WithHelp("↓/j/tab", "move down"),
		),
		key.NewBinding(
			key.WithKeys("space"),
			key.WithHelp("space", "choose"),
		),
	}

	model := &Model{
		choices: choices,
		keys:    multiChooseKeys,
		theme:   ThemeDefault,
	}

	for _, opt := range opts {
		opt(model)
	}

	return model
}
