package list

import (
	"github.com/charmbracelet/lipgloss"
)

var (
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
