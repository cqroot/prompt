package list

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/cqroot/prompt/perrors"
)

type ListBaseModel struct {
	quitting      bool
	err           error
	message       string
	finishMessage string
	cursor        int
	choice        string
	choices       []string
	style         *ListStyle
	direction     listDirection
}

func (m *ListBaseModel) setQuitting(quitting bool) {
	m.quitting = quitting
}

func (m *ListBaseModel) setErr(err error) {
	m.err = err
}

func (m ListBaseModel) getErr() error {
	return m.err
}

func (m *ListBaseModel) setMessage(message string) {
	m.message = message
}

func (m ListBaseModel) getMessage() string {
	return m.message
}

func (m *ListBaseModel) setFinishMessage(finishMessage string) {
	m.finishMessage = finishMessage
}

func (m ListBaseModel) getFinishMessage() string {
	return m.finishMessage
}

func (m *ListBaseModel) setChoice(choice string) {
	m.choice = choice
}

func (m *ListBaseModel) getChoice() string {
	return m.choice
}

func (m *ListBaseModel) setCursor(cursor int) {
	m.cursor = cursor
}

func (m ListBaseModel) getCursor() int {
	return m.cursor
}

func (m *ListBaseModel) setChoices(choices []string) {
	m.choices = choices
}

func (m *ListBaseModel) setStyle(style *ListStyle) {
	m.style = style
}

func (m ListBaseModel) getStyle() *ListStyle {
	return m.style
}

func (m ListBaseModel) Init() tea.Cmd {
	return nil
}

func (m ListBaseModel) update(model IListBaseModel, msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			model.setQuitting(true)
			model.setErr(perrors.ErrUserQuit)
			return tea.Quit

		case "enter":
			model.setChoice(m.choices[model.getCursor()])
			return tea.Quit

		case "up", "k", "left", "h":
			if m.direction == directionVertical &&
				msg.String() != "up" && msg.String() != "k" {
				break
			}
			if m.direction == directionHorizontal &&
				msg.String() != "left" && msg.String() != "h" {
				break
			}
			model.setCursor(model.getCursor() - 1)
			if model.getCursor() < 0 {
				model.setCursor(len(m.choices) - 1)
			}

		case "down", "j", "right", "l", "tab", " ":
			if m.direction == directionVertical &&
				(msg.String() == "right" || msg.String() == "l") {
				break
			}
			if m.direction == directionHorizontal &&
				(msg.String() == "down" || msg.String() == "j") {
				break
			}
			model.setCursor(model.getCursor() + 1)
			if model.getCursor() >= len(m.choices) {
				model.setCursor(0)
			}
		}
	}

	return nil
}

func (m ListBaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m ListBaseModel) View() string {
	return ""
}

func listWithStyle(model IListBaseModel, choices []string, style *ListStyle, message string, finishMessage string) (tea.Model, error) {
	model.setMessage(message)
	model.setFinishMessage(finishMessage)
	model.setChoices(choices)
	model.setStyle(style)

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	return tm, err
}

// func List(model IListBaseModel, choices []string) (tea.Model, error) {
//     return ListWithStyle(model, choices, NewListStyle())
// }

func result(model IListBaseModel) (string, error) {
	if err := model.getErr(); err != nil {
		return "", err
	} else {
		return model.getChoice(), nil
	}
}
