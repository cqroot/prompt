package prompt

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InputModel struct {
	textInput textinput.Model
	err       error
	quitting  bool

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

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.quitting = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			m.err = ErrUserQuit
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {
	if m.quitting {
		return fmt.Sprintf("%s %s %s %s\n",
			m.DonePromptPrefixStyle.Render(m.DonePromptPrefix),
			m.Prompt,
			m.DonePromptSuffixStyle.Render(m.DonePromptSuffix),
			m.ChoiceStyle.Render(m.textInput.Value()),
		)

	}

	return fmt.Sprintf(
		"%s %s %s %s",
		m.NormalPromptPrefixStyle.Render(m.NormalPromptPrefix),
		m.Prompt,
		m.NormalPromptSuffixStyle.Render(m.NormalPromptSuffix),
		m.textInput.View(),
	)
}

func InputWithModel(m InputModel, defaultValue string) (string, error) {
	ti := textinput.New()
	ti.Placeholder = defaultValue
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Prompt = ""

	m.textInput = ti
	m.err = nil

	p := tea.NewProgram(m)

	tm, err := p.Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(InputModel)
	if !ok {
		return "", ErrModelConversion
	}

	if m.err != nil {
		return "", m.err
	} else if m.textInput.Value() == "" {
		return defaultValue, nil
	} else {
		return m.textInput.Value(), nil
	}
}

func Input(prompt string, defaultValue string) (string, error) {
	m := InputModel{
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

	return InputWithModel(m, defaultValue)
}
