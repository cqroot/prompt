package prompt

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type selectModel struct {
	cursor   int
	quitting bool

	Prompt             string
	Choice             string
	Choices            []string
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

var ()

func (m selectModel) Init() tea.Cmd {
	return nil
}

func (m selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			m.Choice = m.Choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.Choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.Choices) - 1
			}
		}
	}

	return m, nil
}

func (m selectModel) View() string {
	if m.Choice != "" {
		return fmt.Sprintf("%s %s %s %s\n",
			m.DonePromptPrefixStyle.Render(m.DonePromptPrefix),
			m.Prompt,
			m.DonePromptSuffixStyle.Render(m.DonePromptSuffix),
			m.ChoiceStyle.Render(m.Choice),
		)
	}
	if m.quitting {
		return ""
	}

	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%s %s %s\n",
		m.NormalPromptPrefixStyle.Render(m.NormalPromptPrefix),
		m.Prompt,
		m.NormalPromptSuffixStyle.Render(m.NormalPromptSuffix),
	))

	for i := 0; i < len(m.Choices); i++ {
		if m.cursor == i {
			s.WriteString(m.SelectedItemStyle.Render(fmt.Sprintf("❯ %s", m.Choices[i])))
		} else {
			s.WriteString(m.ItemStyle.Render(fmt.Sprintf("  %s", m.Choices[i])))
		}
		s.WriteString("\n")
	}

	return s.String()
}

func Select(prompt string, choices []string) (string, error) {
	p := tea.NewProgram(selectModel{
		Prompt:                  prompt,
		Choices:                 choices,
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
	})

	tm, err := p.Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(selectModel)
	if ok && m.Choice != "" {
		return m.Choice, nil
	} else {
		return "", nil
	}
}
