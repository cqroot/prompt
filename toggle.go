package prompt

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type ToggleModel struct {
	choices []string
	keys    []key.Binding
	ListHandler
}

func (m ToggleModel) Data() any {
	return m.choices[m.Cursor()]
}

func (m ToggleModel) DataString() string {
	return m.choices[m.Cursor()]
}

func (m ToggleModel) KeyBindings() []key.Binding {
	return m.keys
}

func (m ToggleModel) Init() tea.Cmd {
	return nil
}

func (m ToggleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys[0]):
			m.MovePrev()

		case key.Matches(msg, m.keys[1]):
			m.MoveNext()
		}
	}

	return m, nil
}

func (m ToggleModel) View() string {
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

	return s.String()
}

func NewToggleModelWithStyle(choices []string, style *ListStyle) *ToggleModel {
	toggleKeys := []key.Binding{
		key.NewBinding(
			key.WithKeys("left", "h", "j"),
			key.WithHelp("←/h/j", "move left"),
		),
		key.NewBinding(
			key.WithKeys("right", "l", "k", "tab", " "),
			key.WithHelp("→/l/k/tab/space", "move right"),
		),
	}

	model := ToggleModel{
		choices:     choices,
		keys:        toggleKeys,
		ListHandler: *NewListHandler(len(choices), style),
	}

	return &model
}

func NewToggleModel(choices []string) *ToggleModel {
	return NewToggleModelWithStyle(choices, NewListStyle())
}

func (p Prompt) ToggleWithStyle(choices []string, style *ListStyle) (string, error) {
	pm := NewToggleModelWithStyle(choices, style)
	p.SetModel(*pm)
	m, err := p.Run()
	return m.DataString(), err
}

func (p Prompt) Toggle(choices []string) (string, error) {
	return p.ToggleWithStyle(choices, NewListStyle())
}
