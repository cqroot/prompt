package prompt

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MultiSelectModel struct {
	quitting bool
	err      error
	prompt   Prompt
	result   []string

	cursor  int
	choice  uint64
	Choices []string

	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
}

func (m *MultiSelectModel) toggleChoice(index int) {
	i := uint64(index)
	if m.isSelected(index) {
		m.deselectItem(i)
	} else {
		m.selectItem(i)
	}
}

func (m MultiSelectModel) isSelected(index int) bool {
	i := uint64(index)
	curr := uint64(1) << i
	if (m.choice & curr) != 0 {
		return true
	} else {
		return false
	}
}

func (m *MultiSelectModel) selectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice | curr
}

func (m *MultiSelectModel) deselectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice & (^curr)
}

func (m MultiSelectModel) Init() tea.Cmd {
	return nil
}

func (m *MultiSelectModel) quit() {
	m.quitting = true
	m.result = make([]string, 0)

	for i := 0; i < len(m.Choices); i++ {
		if m.isSelected(i) {
			m.result = append(m.result, m.Choices[i])
		}
	}
}

func (m MultiSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
            m.quit()
			m.err = ErrUserQuit
			return m, tea.Quit

		case " ":
			m.toggleChoice(m.cursor)

		case "enter":
            m.quit()
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.Choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.Choices) - 1
			}
		}
	}

	return m, nil
}

func (m MultiSelectModel) View() string {
	if m.quitting {
		return fmt.Sprintf("%s %s\n",
			m.prompt.finishView(),
			m.ChoiceStyle.Render(strings.Join(m.result, ", ")),
		)
	}

	s := strings.Builder{}
	s.WriteString(m.prompt.view())
	s.WriteString("\n")

	for i := 0; i < len(m.Choices); i++ {
		if m.cursor == i {
			if m.isSelected(i) {
				s.WriteString(m.SelectedItemStyle.Render(fmt.Sprintf("[x] %s", m.Choices[i])))
			} else {
				s.WriteString(m.SelectedItemStyle.Render(fmt.Sprintf("[•] %s", m.Choices[i])))
			}
		} else {
			if m.isSelected(i) {
				s.WriteString(m.ItemStyle.Render(fmt.Sprintf("[x] %s", m.Choices[i])))
			} else {
				s.WriteString(m.ItemStyle.Render(fmt.Sprintf("[ ] %s", m.Choices[i])))
			}
		}
		s.WriteString("\n")
	}

	return s.String()
}

func NewMultiSelectModel(choices []string) *MultiSelectModel {
	m := MultiSelectModel{
		Choices:           choices,
		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

func (p Prompt) MultiSelectWithModel(model *MultiSelectModel) ([]string, error) {
	model.err = nil
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	m, ok := tm.(MultiSelectModel)
	if !ok {
		return nil, ErrModelConversion
	}

	if m.err != nil {
		return nil, m.err
	} else {
		return m.result, nil
	}
}

func (p Prompt) MultiSelect(choices []string) ([]string, error) {
	return p.MultiSelectWithModel(NewMultiSelectModel(choices))
}
