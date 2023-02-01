package prompt

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	quitting      bool
	err           error
	subModel      PromptModel
	isHelpVisible bool
	help          help.Model
	// Style
	Message           string
	NormalPrefix      string
	FinishPrefix      string
	NormalSuffix      string
	FinishSuffix      string
	PrefixStyle       lipgloss.Style
	FinishPrefixStyle lipgloss.Style
	SuffixStyle       lipgloss.Style
	FinishSuffixStyle lipgloss.Style
}

// New returns a default style *Prompt.
func New() *Prompt {
	return &Prompt{
		quitting:      false,
		err:           nil,
		isHelpVisible: false,
		help:          help.New(),
		// Style
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

// Ask set prompt message
func (p *Prompt) Ask(message string) *Prompt {
	p.Message = message
	return p
}

func (p *Prompt) SetHelpVisible(visible bool) {
	p.isHelpVisible = visible
}

func (p *Prompt) RunModel(pm PromptModel) (PromptModel, error) {
	p.subModel = pm

	tm, err := tea.NewProgram(p).Run()
	if err != nil {
		return nil, err
	}

	m, ok := tm.(Prompt)
	if !ok {
		return nil, ErrModelConversion
	}

	return m.subModel, nil
}
