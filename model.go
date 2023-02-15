package prompt

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/styles"
)

type PromptModel interface {
	tea.Model
	DataString() string // Returns a string for display in the result position.
	Quitting() bool
	Error() error
}

func (p Prompt) Init() tea.Cmd {
	return nil
}

func (p Prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := p.model.Update(msg)
	p.model = model.(PromptModel)
	return p, cmd
}

func (p Prompt) View() string {
	s := strings.Builder{}

	if p.model.Error() != nil {
		s.WriteString(styles.DefaultErrorPromptPrefixStyle.Render("✖"))
	} else if p.model.Quitting() {
		s.WriteString(p.FinishPrefixStyle.Render(p.FinishPrefix))
	} else {
		s.WriteString(p.PrefixStyle.Render(p.NormalPrefix))
	}

	s.WriteString(" ")
	s.WriteString(p.Message)
	s.WriteString(" ")

	if p.model.Quitting() {
		s.WriteString(p.FinishSuffixStyle.Render(p.FinishSuffix))
		s.WriteString(" ")
		s.WriteString(p.model.DataString())
		s.WriteString("\n")
	} else {
		s.WriteString(p.SuffixStyle.Render(p.NormalSuffix))
		s.WriteString(" ")

		modelView := p.model.View()
		s.WriteString(modelView)
	}

	return s.String()
}
