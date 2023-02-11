package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type MultiChooseModel struct {
	choice  uint64
	choices []string
	keys    []key.Binding
	ListHandler
}

func (m MultiChooseModel) Data() any {
	result := make([]string, 0)

	for i := 0; i < len(m.choices); i++ {
		if m.isChooseed(i) {
			result = append(result, m.choices[i])
		}
	}
	return result
}

func (m MultiChooseModel) DataString() string {
	return strings.Join(m.Data().([]string), ", ")
}

func (m MultiChooseModel) KeyBindings() []key.Binding {
	return m.keys
}

func (m MultiChooseModel) UseKeyQ() bool {
	return false
}

func (m *MultiChooseModel) toggleChoice(index int) {
	i := uint64(index)
	if m.isChooseed(index) {
		m.deselectItem(i)
	} else {
		m.selectItem(i)
	}
}

func (m MultiChooseModel) isChooseed(index int) bool {
	i := uint64(index)
	curr := uint64(1) << i
	if (m.choice & curr) != 0 {
		return true
	} else {
		return false
	}
}

func (m *MultiChooseModel) selectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice | curr
}

func (m *MultiChooseModel) deselectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice & (^curr)
}

func (m MultiChooseModel) Init() tea.Cmd {
	return nil
}

func (m MultiChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			m.toggleChoice(m.cursor)

		case "up", "k":
			m.MovePrev()

		case "down", "j", "tab":
			m.MoveNext()
		}
	}

	return m, nil
}

func (m MultiChooseModel) View() string {
	s := strings.Builder{}
	s.WriteString("\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			if m.isChooseed(i) {
				s.WriteString(m.Style().SelectedItemStyle.Render(fmt.Sprintf("[x] %s", m.choices[i])))
			} else {
				s.WriteString(m.Style().SelectedItemStyle.Render(fmt.Sprintf("[•] %s", m.choices[i])))
			}
		} else {
			if m.isChooseed(i) {
				s.WriteString(m.Style().ItemStyle.Render(fmt.Sprintf("[x] %s", m.choices[i])))
			} else {
				s.WriteString(m.Style().ItemStyle.Render(fmt.Sprintf("[ ] %s", m.choices[i])))
			}
		}
		s.WriteString("\n")
	}

	return s.String()
}

func NewMultiChooseModelWithStyle(choices []string, style *ListStyle) *MultiChooseModel {
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

	model := MultiChooseModel{
		choices:     choices,
		keys:        multiChooseKeys,
		ListHandler: *NewListHandler(len(choices), style),
	}

	return &model
}

func NewMultiChooseModel(choices []string) *MultiChooseModel {
	return NewMultiChooseModelWithStyle(choices, NewListStyle())
}

// MultiChooseWithStyle lets the user choose one of the given choices.
// Appearance uses the given style.
func (p Prompt) MultiChooseWithStyle(choices []string, style *ListStyle, opts ...tea.ProgramOption) ([]string, error) {
	pm := NewMultiChooseModelWithStyle(choices, style)
	m, err := p.Run(*pm, opts...)
	if err != nil {
		return nil, err
	}
	return m.Data().([]string), nil
}

// MultiChoose lets the user choose multiples from the given choices.
// Appearance uses the default style.
func (p Prompt) MultiChoose(choices []string, opts ...tea.ProgramOption) ([]string, error) {
	return p.MultiChooseWithStyle(choices, NewListStyle(), opts...)
}
