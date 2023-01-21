package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ToggleModel struct {
	quitting bool

	choice      bool
	TrueString  string
	FalseString string

	Prompt             string
	NormalPromptPrefix string
	DonePromptPrefix   string
	NormalPromptSuffix string
	DonePromptSuffix   string

	ItemStyle               lipgloss.Style
	SelectedItemStyle       lipgloss.Style
	ChoiceStyle             lipgloss.Style
	NormalPromptPrefixStyle lipgloss.Style
	DonePromptPrefixStyle   lipgloss.Style
	NormalPromptSuffixStyle lipgloss.Style
	DonePromptSuffixStyle   lipgloss.Style
}

func (m ToggleModel) Init() tea.Cmd {
	return nil
}

func (m ToggleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc", "enter":
			m.quitting = true
			return m, tea.Quit

		case "up", "down", "left", "right", "j", "k", "h", "l", "tab", "space":
			m.choice = !m.choice
		}
	}

	return m, nil
}

func (m ToggleModel) choiceToString() string {
	if m.choice {
		return m.TrueString
	} else {
		return m.FalseString
	}
}

func (m ToggleModel) View() string {
	if m.quitting {
		return fmt.Sprintf("%s %s %s %s\n",
			m.DonePromptPrefixStyle.Render(m.DonePromptPrefix),
			m.Prompt,
			m.DonePromptSuffixStyle.Render(m.DonePromptSuffix),
			m.ChoiceStyle.Render(m.choiceToString()),
		)
	}

	var toggleString string

	if m.choice {
		toggleString = fmt.Sprintf("%s / %s",
			m.SelectedItemStyle.Render(m.TrueString),
			m.ItemStyle.Render(m.FalseString),
		)
	} else {
		toggleString = fmt.Sprintf("%s / %s",
			m.ItemStyle.Render(m.TrueString),
			m.SelectedItemStyle.Render(m.FalseString),
		)
	}

	return fmt.Sprintf("%s %s %s %s",
		m.NormalPromptPrefixStyle.Render(m.NormalPromptPrefix),
		m.Prompt,
		m.NormalPromptSuffixStyle.Render(m.NormalPromptSuffix),
		toggleString,
	)
}

func ToggleWithModel(defaultValue bool, m ToggleModel) (bool, error) {
	m.choice = defaultValue
	p := tea.NewProgram(m)

	tm, err := p.Run()
	if err != nil {
		return defaultValue, err
	}

	m, ok := tm.(ToggleModel)
	if ok {
		return m.choice, nil
	} else {
		return defaultValue, nil
	}
}

func Toggle(prompt string, defaultValue bool) (bool, error) {
	m := ToggleModel{
		TrueString:              "Yes",
		FalseString:             "No",
		Prompt:                  prompt,
		NormalPromptPrefix:      "?",
		DonePromptPrefix:        "✔",
		NormalPromptSuffix:      "›",
		DonePromptSuffix:        "…",
		ItemStyle:               lipgloss.NewStyle(),
		SelectedItemStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color("14")),
		ChoiceStyle:             lipgloss.NewStyle().Foreground(lipgloss.Color("14")),
		NormalPromptPrefixStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("14")),
		DonePromptPrefixStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("10")),
		NormalPromptSuffixStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
		DonePromptSuffixStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("6")),
	}

	return ToggleWithModel(defaultValue, m)
}
