package prompt

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
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
	isHelpVisible     bool
	help              help.Model
	keyMap            help.KeyMap
}

// New returns a default style *Prompt.
func New() *Prompt {
	return &Prompt{
		NormalPrefix:      DefaultNormalPromptPrefix,
		FinishPrefix:      DefaultFinishPromptPrefix,
		NormalSuffix:      DefaultNormalPromptSuffix,
		FinishSuffix:      DefaultFinishPromptSuffix,
		PrefixStyle:       DefaultNormalPromptPrefixStyle,
		FinishPrefixStyle: DefaultFinishPromptPrefixStyle,
		SuffixStyle:       DefaultNormalPromptSuffixStyle,
		FinishSuffixStyle: DefaultFinishPromptSuffixStyle,
		isHelpVisible:     false,
		help:              help.New(),
	}
}

// Ask set prompt message
func (p *Prompt) Ask(message string) *Prompt {
	p.Message = message
	return p
}

func (p *Prompt) SetHelpVisible(visible bool) {
	p.isHelpVisible = visible
}

func (p *Prompt) setKeyMap(keyMap help.KeyMap) {
	p.keyMap = keyMap
}

func (p Prompt) view(content string) string {
	s := strings.Builder{}
	s.WriteString(p.PrefixStyle.Render(p.NormalPrefix))
	s.WriteString(" ")
	s.WriteString(p.Message)
	s.WriteString(" ")
	s.WriteString(p.SuffixStyle.Render(p.NormalSuffix))
	s.WriteString(" ")
	s.WriteString(content)
	if p.isHelpVisible && p.keyMap != nil {
		s.WriteString("\n")
		s.WriteString(p.help.View(p.keyMap))
	}

	return s.String()
}

func (p Prompt) finishView(content string) string {
	return strings.Join([]string{
		p.FinishPrefixStyle.Render(p.FinishPrefix),
		p.Message,
		p.FinishSuffixStyle.Render(p.FinishSuffix),
		content,
	}, " ") + "\n"
}
