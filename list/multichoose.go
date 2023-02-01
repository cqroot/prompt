package list

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/perrors"
)

type MultiChooseModel struct {
	choice  uint64
	choices []string
	result  []string
	ListHandler
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

func (m *MultiChooseModel) quit() {
	m.quitting = true
	m.result = make([]string, 0)

	for i := 0; i < len(m.choices); i++ {
		if m.isChooseed(i) {
			m.result = append(m.result, m.choices[i])
		}
	}
}

func (m MultiChooseModel) Init() tea.Cmd {
	return nil
}

func (m MultiChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.Quit()
			return m, tea.Quit

		case " ":
			m.toggleChoice(m.cursor)

		case "enter":
			m.quit()
			return m, tea.Quit

		case "up", "k":
			m.MovePrev()

		case "down", "j", "tab":
			m.MoveNext()
		}
	}

	return m, nil
}

func (m MultiChooseModel) View() string {
	if m.quitting {
		return fmt.Sprintf("%s %s\n",
			m.FinishMessage(),
			m.style.ChoiceStyle.Render(strings.Join(m.result, ", ")),
		)
	}

	s := strings.Builder{}
	s.WriteString(m.Message())
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

func NewMultiChooseModel(choices []string, style *ListStyle, message string, finishMessage string) *MultiChooseModel {
	model := MultiChooseModel{
		choices: choices,
	}
	model.SetChoiceCount(len(choices))
	model.SetStyle(style)
	model.SetMessage(message)
	model.SetFinishMessage(message)

	return &model
}

func MultiChooseWithStyle(choices []string, style *ListStyle, message string, finishMessage string) ([]string, error) {
	model := NewMultiChooseModel(choices, style, message, finishMessage)

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	m, ok := tm.(MultiChooseModel)
	if !ok {
		return nil, perrors.ErrModelConversion
	}

	if err := m.err; err != nil {
		return nil, err
	} else {
		return m.result, nil
	}
}
