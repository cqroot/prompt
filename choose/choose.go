package choose

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices []string
	cursor  int
	keys    []key.Binding
	theme   Theme
}

func (m Model) Data() any {
	return m.choices[m.cursor]
}

func (m Model) DataString() string {
	return m.Data().(string)
}

func (m *Model) initKeys() {
	chooseKeys := make([]key.Binding, 0, 4)
	if m.theme.Direction == DirectionH || m.theme.Direction == DirectionAll {
		chooseKeys = append(chooseKeys,
			key.NewBinding(
				key.WithKeys("left", "h"),
				key.WithHelp("←/h", "move left"),
			),
			key.NewBinding(
				key.WithKeys("right", "l", "tab", " "),
				key.WithHelp("→/l/tab/space", "move right"),
			),
		)
	}
	if m.theme.Direction == DirectionV || m.theme.Direction == DirectionAll {
		chooseKeys = append(chooseKeys,
			key.NewBinding(
				key.WithKeys("up", "k"),
				key.WithHelp("↑/k", "move up"),
			),
			key.NewBinding(
				key.WithKeys("down", "j", "tab"),
				key.WithHelp("↓/j/tab", "move down"),
			),
		)
	}
	m.keys = chooseKeys
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

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys[0]):
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}

		case key.Matches(msg, m.keys[1]):
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	return m.theme.View(m.choices, m.cursor)
}

func NewChooseModel(choices []string, opts ...Option) *Model {
	model := &Model{
		choices: choices,
		theme:   ThemeDefault,
	}

	for _, opt := range opts {
		opt(model)
	}
	model.initKeys()

	return model
}
