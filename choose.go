package prompt

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ChooseModel struct {
	choices []string
	cursor  int
	keys    []key.Binding
	theme   ChooseTheme
}

func (m ChooseModel) Data() any {
	return m.choices[m.cursor]
}

func (m ChooseModel) DataString() string {
	return m.Data().(string)
}

func (m *ChooseModel) initKeys() {
	chooseKeys := make([]key.Binding, 0, 4)
	if m.theme.Direction == ChooseDirectionH || m.theme.Direction == ChooseDirectionAll {
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
	if m.theme.Direction == ChooseDirectionV || m.theme.Direction == ChooseDirectionAll {
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

func (m ChooseModel) KeyBindings() []key.Binding {
	return m.keys
}

func (m ChooseModel) UseKeyQ() bool {
	return false
}

func (m ChooseModel) UseKeyEnter() bool {
	return false
}

func (m ChooseModel) Init() tea.Cmd {
	return nil
}

func (m ChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m ChooseModel) View() string {
	return m.theme.View(m.choices, m.cursor)
}

func NewChooseModel(choices []string, opts ...ChooseOption) *ChooseModel {
	model := &ChooseModel{
		choices: choices,
		theme:   ChooseThemeDefault,
	}

	for _, opt := range opts {
		opt(model)
	}
	model.initKeys()

	return model
}

// Choose lets the user choose one of the given choices.
func (p Prompt) Choose(choices []string, opts ...ChooseOption) (string, error) {
	pm := NewChooseModel(choices, opts...)

	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}
