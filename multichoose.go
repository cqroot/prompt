package prompt

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MultiChooseModel struct {
	quitting bool
	err      error
	prompt   Prompt
	result   []string

	cursor  int
	choice  uint64
	Choices []string

	ItemStyle         lipgloss.Style
	ChooseedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
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

func (m *MultiChooseModel) quit() {
	m.quitting = true
	m.result = make([]string, 0)

	for i := 0; i < len(m.Choices); i++ {
		if m.isChooseed(i) {
			m.result = append(m.result, m.Choices[i])
		}
	}
}

func (m MultiChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m MultiChooseModel) View() string {
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
			if m.isChooseed(i) {
				s.WriteString(m.ChooseedItemStyle.Render(fmt.Sprintf("[x] %s", m.Choices[i])))
			} else {
				s.WriteString(m.ChooseedItemStyle.Render(fmt.Sprintf("[•] %s", m.Choices[i])))
			}
		} else {
			if m.isChooseed(i) {
				s.WriteString(m.ItemStyle.Render(fmt.Sprintf("[x] %s", m.Choices[i])))
			} else {
				s.WriteString(m.ItemStyle.Render(fmt.Sprintf("[ ] %s", m.Choices[i])))
			}
		}
		s.WriteString("\n")
	}

	return s.String()
}

func NewMultiChooseModel(choices []string) *MultiChooseModel {
	m := MultiChooseModel{
		Choices:           choices,
		ItemStyle:         DefaultItemStyle,
		ChooseedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

func (p Prompt) MultiChooseWithModel(model *MultiChooseModel) ([]string, error) {
	model.err = nil
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	m, ok := tm.(MultiChooseModel)
	if !ok {
		return nil, ErrModelConversion
	}

	if m.err != nil {
		return nil, m.err
	} else {
		return m.result, nil
	}
}

func (p Prompt) MultiChoose(choices []string) ([]string, error) {
	return p.MultiChooseWithModel(NewMultiChooseModel(choices))
}
