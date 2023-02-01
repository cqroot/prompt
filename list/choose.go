package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/perrors"
)

type ChooseModel struct {
	choice  string
	choices []string
	ListHandler
}

func (m ChooseModel) Init() tea.Cmd {
	return nil
}

func (m ChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.Quit()
			return m, tea.Quit

		case "enter":
			m.choice = m.choices[m.Cursor()]
			return m, tea.Quit

		case "up", "k":
			m.MovePrev()

		case "down", "j", "tab":
			m.MoveNext()
		}
	}

	return m, nil
}

func (m ChooseModel) View() string {
	if m.choice != "" {
		return fmt.Sprintf("%s %s\n",
			m.FinishMessage(),
			m.Style().ChoiceStyle.Render(m.choice),
		)
	}
	if m.Quitting() {
		return ""
	}

	s := strings.Builder{}
	s.WriteString(m.Message())
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

func NewChooseModel(choices []string, style *ListStyle, message string, finishMessage string) *ChooseModel {
	model := ChooseModel{
		choices: choices,
	}
	model.SetChoiceCount(len(choices))
	model.SetStyle(style)
	model.SetMessage(message)
	model.SetFinishMessage(message)

	return &model
}

func ChooseWithStyle(choices []string, style *ListStyle, message string, finishMessage string) (string, error) {
	model := NewChooseModel(choices, style, message, finishMessage)

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(ChooseModel)
	if !ok {
		return "", perrors.ErrModelConversion
	}

	if err := m.err; err != nil {
		return "", err
	} else {
		return m.choice, nil
	}
}
