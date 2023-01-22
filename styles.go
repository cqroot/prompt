package prompt

import (
	"github.com/charmbracelet/bubbles/textinput"
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

func Default() Prompt {
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

func NewToggleModel() *ToggleModel {
	m := ToggleModel{
		TrueString:        "Yes",
		FalseString:       "No",
		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

func NewInputModel(defaultValue string) *InputModel {
	ti := textinput.New()
	ti.Placeholder = defaultValue
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Prompt = ""

	m := InputModel{
		textInput: ti,
		err:       nil,
		df:        defaultValue,

		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

func NewSelectModel(choices []string) *SelectModel {
	m := SelectModel{
		Choices:           choices,
		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}
