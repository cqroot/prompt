package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/perrors"
)

type ChooseModel struct {
	ListBaseModel
}

func (m ChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.update(&m, msg)
	return m, cmd
}

func (m ChooseModel) View() string {
	if m.choice != "" {
		return fmt.Sprintf("%s %s\n",
			m.getFinishMessage(),
			m.getStyle().ChoiceStyle.Render(m.choice),
		)
	}
	if m.quitting {
		return ""
	}

	s := strings.Builder{}
	s.WriteString(m.getMessage())
	s.WriteString("\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString(m.getStyle().SelectedItemStyle.Render(fmt.Sprintf("• %s", m.choices[i])))
		} else {
			s.WriteString(m.getStyle().ItemStyle.Render(fmt.Sprintf("  %s", m.choices[i])))
		}
		s.WriteString("\n")
	}

	return s.String()
}

func ChooseWithStyle(choices []string, style *ListStyle, message string, finishMessage string) (string, error) {
	model := ChooseModel{}
	model.direction = directionVertical
	tm, err := listWithStyle(&model, choices, style, message, finishMessage)
	if err != nil {
		return "", err
	}

	m, ok := tm.(ChooseModel)
	if !ok {
		return "", perrors.ErrModelConversion
	}

	return result(&m)
}
