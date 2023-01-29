package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/perrors"
)

type ToggleModel struct {
	ListBaseModel
}

func (m ToggleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.update(&m, msg)
	return m, cmd
}

func (m ToggleModel) View() string {
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
	s.WriteString(" ")

	choices := make([]string, len(m.choices))
	for index, choice := range m.choices {
		if index == m.cursor {
			choices[index] = m.getStyle().SelectedItemStyle.Render(choice)
		} else {
			choices[index] = m.getStyle().ItemStyle.Render(choice)
		}
	}
	s.WriteString(strings.Join(choices, " / "))

	return s.String()
}

func ToggleWithStyle(choices []string, style *ListStyle, message string, finishMessage string) (string, error) {
	model := ToggleModel{}
	model.direction = directionHorizontal
	tm, err := listWithStyle(&model, choices, style, message, finishMessage)
	if err != nil {
		return "", err
	}

	m, ok := tm.(ToggleModel)
	if !ok {
		return "", perrors.ErrModelConversion
	}

	return result(&m)
}
