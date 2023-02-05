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

func (p Prompt) Init() tea.Cmd {
	return nil
}

func (p Prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.help.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			p.quitting = true
			p.err = ErrUserQuit
			return p, tea.Quit

		case "enter":
			p.quitting = true
			return p, tea.Quit
		}
	}

	model, cmd := p.model.Update(msg)
	p.model = model.(PromptModel)
	return p, cmd
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
		s.WriteString(p.model.DataString())
		s.WriteString("\n")
		return s.String()
	}

	s.WriteString(p.PrefixStyle.Render(p.NormalPrefix))
	s.WriteString(" ")
	s.WriteString(p.Message)
	s.WriteString(" ")
	s.WriteString(p.SuffixStyle.Render(p.NormalSuffix))
	s.WriteString(" ")

	modelView := p.model.View()
	s.WriteString(modelView)

	if p.isHelpVisible {
		if len(modelView) > 1 && modelView[len(modelView)-1] == '\n' {
			s.WriteString("\n")
		} else {
			s.WriteString("\n\n")
		}
		keyBindings := p.model.KeyBindings()
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
