package prompt

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DefaultNormalPromptPrefix      = "?"
	DefaultDonePromptPrefix        = "✔"
	DefaultNormalPromptSuffix      = "›"
	DefaultDonePromptSuffix        = "…"
	DefaultItemStyle               = lipgloss.NewStyle()
	DefaultSelectedItemStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultChoiceStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultNormalPromptPrefixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultDonePromptPrefixStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	DefaultNormalPromptSuffixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	DefaultDonePromptSuffixStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
)
