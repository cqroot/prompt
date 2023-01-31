package prompt

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/cqroot/prompt/styles"
)

type Prompt struct {
	Message           string
	NormalPrefix      string
	FinishPrefix      string
	NormalSuffix      string
	FinishSuffix      string
	PrefixStyle       lipgloss.Style
	FinishPrefixStyle lipgloss.Style
	SuffixStyle       lipgloss.Style
	FinishSuffixStyle lipgloss.Style
}

func New() *Prompt {
	return &Prompt{
		NormalPrefix:      styles.DefaultNormalPromptPrefix,
		FinishPrefix:      styles.DefaultFinishPromptPrefix,
		NormalSuffix:      styles.DefaultNormalPromptSuffix,
		FinishSuffix:      styles.DefaultFinishPromptSuffix,
		PrefixStyle:       styles.DefaultNormalPromptPrefixStyle,
		FinishPrefixStyle: styles.DefaultFinishPromptPrefixStyle,
		SuffixStyle:       styles.DefaultNormalPromptSuffixStyle,
		FinishSuffixStyle: styles.DefaultFinishPromptSuffixStyle,
	}
}

func (p *Prompt) Ask(message string) *Prompt {
	p.Message = message
	return p
}

func (p Prompt) view() string {
	return fmt.Sprintf(
		"%s %s %s",
		p.PrefixStyle.Render(p.NormalPrefix),
		p.Message,
		p.SuffixStyle.Render(p.NormalSuffix),
	)
}

func (p Prompt) finishView() string {
	return fmt.Sprintf("%s %s %s",
		p.FinishPrefixStyle.Render(p.FinishPrefix),
		p.Message,
		p.FinishSuffixStyle.Render(p.FinishSuffix),
	)
}
