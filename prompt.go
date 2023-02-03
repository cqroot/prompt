package prompt

import (
	"errors"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	quitting      bool
	err           error
	model         PromptModel
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

func (p *Prompt) SetHelpVisible(visible bool) *Prompt {
	p.isHelpVisible = visible
	return p
}

func (p Prompt) Model() PromptModel {
	return p.model
}

func (p *Prompt) SetModel(pm PromptModel) *Prompt {
	p.model = pm
	return p
}

func (p *Prompt) Error() error {
	return p.err
}

func (p *Prompt) Run() (PromptModel, error) {
	if p.model == nil {
		return nil, errors.New("prompt has no model")
	}

	tm, err := tea.NewProgram(p).Run()
	if err != nil {
		return nil, err
	}

	m, ok := tm.(Prompt)
	if !ok {
		return nil, ErrModelConversion
	}

	return m.model, nil
}
