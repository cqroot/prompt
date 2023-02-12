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
	UseKeyQ() bool
	UseKeyEnter() bool
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
		case "q":
			if p.model.UseKeyQ() {
				break
			}
			p.quitting = true
			p.err = ErrUserQuit
			return p, tea.Quit

		case "ctrl+c":
			p.quitting = true
			p.err = ErrUserQuit
			return p, tea.Quit

		case "enter":
			if p.model.UseKeyEnter() {
				break
			}
			p.quitting = true
			return p, tea.Quit

		case "ctrl+s":
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

	if p.enableHelp {
		if len(modelView) > 1 && modelView[len(modelView)-1] == '\n' {
			s.WriteString("\n")
		} else {
			s.WriteString("\n\n")
		}
		keyBindings := p.model.KeyBindings()

		var confirmKeyBinding key.Binding
		if p.model.UseKeyEnter() {
			confirmKeyBinding = key.NewBinding(
				key.WithKeys("ctrl+s"),
				key.WithHelp("ctrl+s", "confirm"),
			)
		} else {
			confirmKeyBinding = key.NewBinding(
				key.WithKeys("enter", "ctrl+s"),
				key.WithHelp("enter", "confirm"),
			)
		}

		var quitKeyBinding key.Binding
		if p.model.UseKeyQ() {
			quitKeyBinding = key.NewBinding(
				key.WithKeys("ctrl+c"),
				key.WithHelp("ctrl+c", "quit"),
			)
		} else {
			quitKeyBinding = key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q/ctrl+c", "quit"),
			)
		}

		keyBindings = append(
			keyBindings,
			confirmKeyBinding,
			quitKeyBinding,
		)
		s.WriteString(p.help.ShortHelpView(keyBindings))
	}

	return s.String()
}
