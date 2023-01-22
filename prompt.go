package prompt

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	Text              string
	NormalPrefix      string
	FinishPrefix      string
	NormalSuffix      string
	FinishSuffix      string
	PrefixStyle       lipgloss.Style
	FinishPrefixStyle lipgloss.Style
	SuffixStyle       lipgloss.Style
	FinishSuffixStyle lipgloss.Style
}

func (p Prompt) view() string {
	return fmt.Sprintf(
		"%s %s %s",
		p.PrefixStyle.Render(p.NormalPrefix),
		p.Text,
		p.SuffixStyle.Render(p.NormalSuffix),
	)
}

func (p Prompt) finishView() string {
	return fmt.Sprintf("%s %s %s",
		p.FinishPrefixStyle.Render(p.FinishPrefix),
		p.Text,
		p.FinishSuffixStyle.Render(p.FinishSuffix),
	)
}

func (p *Prompt) Ask(text string) *Prompt {
	p.Text = text
	return p
}
