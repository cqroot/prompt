package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ChooseModel struct {
	choice  string
	choices []string
	keys    listKeyMap
	ListHandler
}

func (m ChooseModel) Init() tea.Cmd {
	return nil
}

func (m ChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.Quit()
			return m, tea.Quit

		case key.Matches(msg, m.keys.Choose):
			m.choice = m.choices[m.Cursor()]
			return m, tea.Quit

		case key.Matches(msg, m.keys.Prev):
			m.MovePrev()

		case key.Matches(msg, m.keys.Next):
			m.MoveNext()
		}
	}

	return m, nil
}

func (m ChooseModel) View() string {
	if m.choice != "" {
		return m.finishView(m.Style().ChoiceStyle.Render(m.choice))
	}
	if m.Quitting() {
		return ""
	}

	s := strings.Builder{}
	s.WriteString("\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString(m.Style().SelectedItemStyle.Render(fmt.Sprintf("• %s", m.choices[i])))
		} else {
			s.WriteString(m.Style().ItemStyle.Render(fmt.Sprintf("  %s", m.choices[i])))
		}
		s.WriteString("\n")
	}

	return m.view(s.String())
}

func NewChooseModel(choices []string, style *ListStyle, message string, finishMessage string) *ChooseModel {
	model := ChooseModel{
		choices:     choices,
		ListHandler: *NewListHandler(len(choices), style, message, finishMessage),
	}

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
		return "", ErrModelConversion
	}

	if err := m.err; err != nil {
		return "", err
	} else {
		return m.choice, nil
	}
}

func (p *Prompt) NewChooseModel(choices []string, style *ListStyle) *ChooseModel {
	chooseKeys := listKeyMap{
		Prev: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Next: key.NewBinding(
			key.WithKeys("down", "j", "tab"),
			key.WithHelp("↓/j/tab", "move down"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
		),
	}

	model := ChooseModel{
		choices:     choices,
		keys:        chooseKeys,
		ListHandler: *p.NewListHandler(len(choices), style),
	}
	model.setKeyMap(model.keys)

	return &model
}

func (p *Prompt) ChooseWithStyle(choices []string, style *ListStyle) (string, error) {
	model := p.NewChooseModel(choices, style)

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(ChooseModel)
	if !ok {
		return "", ErrModelConversion
	}

	if err := m.err; err != nil {
		return "", err
	} else {
		return m.choice, nil
	}
}

func (p *Prompt) Choose(choices []string) (string, error) {
	return p.ChooseWithStyle(choices, NewListStyle())
}
