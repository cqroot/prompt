package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ChooseModel struct {
	choices []string
	keys    []key.Binding
	ListHandler
}

func (m ChooseModel) Data() any {
	return m.choices[m.Cursor()]
}

func (m ChooseModel) DataString() string {
	return m.choices[m.Cursor()]
}

func (m ChooseModel) KeyBindings() []key.Binding {
	return m.keys
}

func (m ChooseModel) Init() tea.Cmd {
	return nil
}

func (m ChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys[0]):
			m.MovePrev()

		case key.Matches(msg, m.keys[1]):
			m.MoveNext()
		}
	}

	return m, nil
}

func (m ChooseModel) View() string {
	s := strings.Builder{}
	s.WriteString("\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString(m.Style().SelectedItemStyle.Render(fmt.Sprintf("• %s", m.choices[i])))
		} else {
			s.WriteString(m.Style().ItemStyle.Render(fmt.Sprintf("  %s", m.choices[i])))
		}
		s.WriteString("\n")
	}

	return s.String()
}

func NewChooseModelWithStyle(choices []string, style *ListStyle) *ChooseModel {
	chooseKeys := []key.Binding{
		key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		key.NewBinding(
			key.WithKeys("down", "j", "tab"),
			key.WithHelp("↓/j/tab", "move down"),
		),
	}

	model := ChooseModel{
		choices:     choices,
		keys:        chooseKeys,
		ListHandler: *NewListHandler(len(choices), style),
	}

	return &model
}

func NewChooseModel(choices []string) *ChooseModel {
	return NewChooseModelWithStyle(choices, NewListStyle())
}

// ChooseWithStyle lets the user choose one of the given choices. Appearance
// uses the given style.
func (p Prompt) ChooseWithStyle(choices []string, style *ListStyle) (string, error) {
	pm := NewChooseModelWithStyle(choices, style)
	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.DataString(), nil
}

// Choose lets the user choose one of the given choices. Appearance uses the
// default style.
func (p Prompt) Choose(choices []string) (string, error) {
	return p.ChooseWithStyle(choices, NewListStyle())
}
