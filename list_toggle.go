package prompt

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ToggleModel struct {
	choice  string
	choices []string
	keys    listKeyMap
	ListHandler
}

func (m ToggleModel) Init() tea.Cmd {
	return nil
}

func (m ToggleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m ToggleModel) View() string {
	if m.choice != "" {
		return m.finishView(m.Style().ChoiceStyle.Render(m.choice))
	}
	if m.quitting {
		return ""
	}

	s := strings.Builder{}

	choices := make([]string, len(m.choices))
	for index, choice := range m.choices {
		if index == m.cursor {
			choices[index] = m.Style().SelectedItemStyle.Render(choice)
		} else {
			choices[index] = m.Style().ItemStyle.Render(choice)
		}
	}
	s.WriteString(strings.Join(choices, " / "))

	return m.view(s.String())
}

func (p *Prompt) NewToggleStyle(choices []string, style *ListStyle) *ToggleModel {
	toggleKeys := listKeyMap{
		Prev: key.NewBinding(
			key.WithKeys("left", "h", "j"),
			key.WithHelp("←/h/j", "move left"),
		),
		Next: key.NewBinding(
			key.WithKeys("right", "l", "k", "tab", " "),
			key.WithHelp("→/l/k/tab/space", "move right"),
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

	model := ToggleModel{
		choices:     choices,
		keys:        toggleKeys,
		ListHandler: *p.NewListHandler(len(choices), style),
	}
	model.setKeyMap(model.keys)

	return &model
}

func (p *Prompt) ToggleWithStyle(choices []string, style *ListStyle) (string, error) {
	model := p.NewToggleStyle(choices, style)

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(ToggleModel)
	if !ok {
		return "", ErrModelConversion
	}

	if err := m.err; err != nil {
		return "", err
	} else {
		return m.choice, nil
	}
}

func (p *Prompt) Toggle(choices []string) (string, error) {
	return p.ToggleWithStyle(choices, NewListStyle())
}
