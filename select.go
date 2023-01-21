package prompt

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SelectModel struct {
	cursor   int
	quitting bool
	err      error

	choice  string
	Choices []string

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

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			m.err = ErrUserQuit
			return m, tea.Quit

		case "enter":
			m.choice = m.Choices[m.cursor]
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

func (m SelectModel) View() string {
	if m.choice != "" {
		return fmt.Sprintf("%s %s %s %s\n",
			m.DonePromptPrefixStyle.Render(m.DonePromptPrefix),
			m.Prompt,
			m.DonePromptSuffixStyle.Render(m.DonePromptSuffix),
			m.ChoiceStyle.Render(m.choice),
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

func SelectWithModel(m SelectModel) (string, error) {
	p := tea.NewProgram(m)

	tm, err := p.Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(SelectModel)
	if !ok {
		return "", ErrModelConversion
	}

	if m.err != nil {
		return "", m.err
	} else {
		return m.choice, nil
	}
}

func Select(prompt string, choices []string) (string, error) {
	m := SelectModel{
		Choices:                 choices,
		Prompt:                  prompt,
		NormalPromptPrefix:      DefaultNormalPromptPrefix,
		DonePromptPrefix:        DefaultDonePromptPrefix,
		NormalPromptSuffix:      DefaultNormalPromptSuffix,
		DonePromptSuffix:        DefaultDonePromptSuffix,
		ItemStyle:               DefaultItemStyle,
		SelectedItemStyle:       DefaultSelectedItemStyle,
		ChoiceStyle:             DefaultChoiceStyle,
		NormalPromptPrefixStyle: DefaultNormalPromptPrefixStyle,
		DonePromptPrefixStyle:   DefaultDonePromptPrefixStyle,
		NormalPromptSuffixStyle: DefaultNormalPromptSuffixStyle,
		DonePromptSuffixStyle:   DefaultDonePromptSuffixStyle,
	}

	return SelectWithModel(m)
}
