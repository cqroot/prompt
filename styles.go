package prompt

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DefaultNormalPromptPrefix      = "?"
	DefaultFinishPromptPrefix      = "✔"
	DefaultNormalPromptSuffix      = "›"
	DefaultFinishPromptSuffix      = "…"
	DefaultItemStyle               = lipgloss.NewStyle()
	DefaultSelectedItemStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultChoiceStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultNormalPromptPrefixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultFinishPromptPrefixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	DefaultNormalPromptSuffixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	DefaultFinishPromptSuffixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
)

func New() Prompt {
	return Prompt{
		NormalPrefix:      DefaultNormalPromptPrefix,
		FinishPrefix:      DefaultFinishPromptPrefix,
		NormalSuffix:      DefaultNormalPromptSuffix,
		FinishSuffix:      DefaultFinishPromptSuffix,
		PrefixStyle:       DefaultNormalPromptPrefixStyle,
		FinishPrefixStyle: DefaultFinishPromptPrefixStyle,
		SuffixStyle:       DefaultNormalPromptSuffixStyle,
		FinishSuffixStyle: DefaultFinishPromptSuffixStyle,
	}
}
