package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/perrors"
)

type ToggleModel struct {
	choice  string
	choices []string
	ListHandler
}

func (m ToggleModel) Init() tea.Cmd {
	return nil
}

func (m ToggleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.Quit()
			return m, tea.Quit

		case "enter":
			m.choice = m.choices[m.Cursor()]
			return m, tea.Quit

		case "left", "up", "h", "j":
			m.MovePrev()

		case "right", "down", "l", "k", "tab":
			m.MoveNext()
		}
	}

	return m, nil
}

func (m ToggleModel) View() string {
	if m.choice != "" {
		return fmt.Sprintf("%s %s\n",
			m.FinishMessage(),
			m.Style().ChoiceStyle.Render(m.choice),
		)
	}
	if m.quitting {
		return ""
	}

	s := strings.Builder{}
	s.WriteString(m.Message())
	s.WriteString(" ")

	choices := make([]string, len(m.choices))
	for index, choice := range m.choices {
		if index == m.cursor {
			choices[index] = m.Style().SelectedItemStyle.Render(choice)
		} else {
			choices[index] = m.Style().ItemStyle.Render(choice)
		}
	}
	s.WriteString(strings.Join(choices, " / "))

	return s.String()
}

func NewToggleStyle(choices []string, style *ListStyle, message string, finishMessage string) *ToggleModel {
	model := ToggleModel{
		choices: choices,
	}
	model.SetChoiceCount(len(choices))
	model.SetStyle(style)
	model.SetMessage(message)
	model.SetFinishMessage(message)

	return &model
}

func ToggleWithStyle(choices []string, style *ListStyle, message string, finishMessage string) (string, error) {
	model := NewToggleStyle(choices, style, message, finishMessage)

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(ToggleModel)
	if !ok {
		return "", perrors.ErrModelConversion
	}

	if err := m.err; err != nil {
		return "", err
	} else {
		return m.choice, nil
	}
}
