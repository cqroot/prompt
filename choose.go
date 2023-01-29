package prompt

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ChooseModel struct {
	quitting bool
	err      error
	prompt   Prompt

	cursor  int
	choice  string
	Choices []string

	ItemStyle         lipgloss.Style
	ChooseedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
}

func (m ChooseModel) Init() tea.Cmd {
	return nil
}

func (m ChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			m.err = ErrUserQuit
			return m, tea.Quit

		case "enter":
			m.choice = m.Choices[m.cursor]
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

func (m ChooseModel) View() string {
	if m.choice != "" {
		return fmt.Sprintf("%s %s\n",
			m.prompt.finishView(),
			m.ChoiceStyle.Render(m.choice),
		)
	}
	if m.quitting {
		return ""
	}

	s := strings.Builder{}
	s.WriteString(m.prompt.view())
	s.WriteString("\n")

	for i := 0; i < len(m.Choices); i++ {
		if m.cursor == i {
			// s.WriteString(m.ChooseedItemStyle.Render(fmt.Sprintf("❯ %s", m.Choices[i])))
			s.WriteString(m.ChooseedItemStyle.Render(fmt.Sprintf("• %s", m.Choices[i])))
		} else {
			s.WriteString(m.ItemStyle.Render(fmt.Sprintf("  %s", m.Choices[i])))
		}
		s.WriteString("\n")
	}

	return s.String()
}

func NewChooseModel(choices []string) *ChooseModel {
	m := ChooseModel{
		Choices:           choices,
		ItemStyle:         DefaultItemStyle,
		ChooseedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

func (p Prompt) ChooseWithModel(model *ChooseModel) (string, error) {
	model.err = nil
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(ChooseModel)
	if !ok {
		return "", ErrModelConversion
	}

	if m.err != nil {
		return "", m.err
	} else {
		return m.choice, nil
	}
}

func (p Prompt) Choose(choices []string) (string, error) {
	return p.ChooseWithModel(NewChooseModel(choices))
}
