package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DefaultItemStyle         = lipgloss.NewStyle()
	DefaultSelectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	DefaultChoiceStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
)
