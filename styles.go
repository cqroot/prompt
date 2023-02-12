package prompt

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DefaultNormalPromptPrefix      = "?"
	DefaultFinishPromptPrefix      = "✔"
	DefaultNormalPromptSuffix      = "›"
	DefaultFinishPromptSuffix      = "…"
	DefaultNormalPromptPrefixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultFinishPromptPrefixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	DefaultErrorPromptPrefixStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	DefaultNormalPromptSuffixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	DefaultFinishPromptSuffixStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))
	DefaultNoteStyle               = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#909090",
		Dark:  "#626262",
	})

	DefaultItemStyle         = lipgloss.NewStyle()
	DefaultSelectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultChoiceStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
)

type ListStyle struct {
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
}

func NewListStyle() *ListStyle {
	return &ListStyle{
		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
}
