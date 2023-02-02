package prompt

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type PromptModel interface {
	tea.Model
	Data() any
	DataString() string
	KeyBindings() []key.Binding
}

func (m Prompt) Init() tea.Cmd {
	return nil
}

func (m Prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			m.err = ErrUserQuit
			return m, tea.Quit

		case "enter":
			m.quitting = true
			return m, tea.Quit
		}
	}

	model, cmd := m.subModel.Update(msg)
	m.subModel = model.(PromptModel)
	return m, cmd
}

func (p Prompt) View() string {
	if p.err != nil {
		return ""
	}

	s := strings.Builder{}

	if p.quitting {
		s.WriteString(p.FinishPrefixStyle.Render(p.FinishPrefix))
		s.WriteString(" ")
		s.WriteString(p.Message)
		s.WriteString(" ")
		s.WriteString(p.FinishSuffixStyle.Render(p.FinishSuffix))
		s.WriteString(" ")
		s.WriteString(p.subModel.DataString())
		s.WriteString("\n")
		return s.String()
	}

	s.WriteString(p.PrefixStyle.Render(p.NormalPrefix))
	s.WriteString(" ")
	s.WriteString(p.Message)
	s.WriteString(" ")
	s.WriteString(p.SuffixStyle.Render(p.NormalSuffix))
	s.WriteString(" ")
	s.WriteString(p.subModel.View())

	if p.isHelpVisible {
		s.WriteString("\n\n")
		keyBindings := p.subModel.KeyBindings()
		keyBindings = append(
			keyBindings,
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "confirm"),
			),
			key.NewBinding(
				key.WithKeys("q", "esc", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
		)
		s.WriteString(p.help.ShortHelpView(keyBindings))
	}

	return s.String()
}
