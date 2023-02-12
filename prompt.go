package prompt

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	quitting       bool
	err            error
	model          PromptModel
	isHelpVisible  bool
	help           help.Model
	programOptions []tea.ProgramOption
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

// New returns a *Prompt using the default style.
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

// SetHelpVisible sets whether the help of the keymap is visible
func (p *Prompt) SetHelpVisible(visible bool) *Prompt {
	p.isHelpVisible = visible
	return p
}

// SetModel sets the model used by the prompt. In most cases you won't need to
// use this.
func (p *Prompt) SetModel(pm PromptModel) *Prompt {
	p.model = pm
	return p
}

// WithProgramOptions sets the `tea.ProgramOption` passed when calling
// `tea.NewProgram`. This function is mainly used for testing, usually you
// don't need to use this function.
func (p *Prompt) WithProgramOptions(opts ...tea.ProgramOption) *Prompt {
	p.programOptions = append(p.programOptions, opts...)
	return p
}

// Run runs the program using the given model, blocking until the user chooses
// or exits.
func (p *Prompt) Run(pm PromptModel) (PromptModel, error) {
	p.model = pm

	tm, err := tea.NewProgram(p, p.programOptions...).Run()
	if err != nil {
		return nil, err
	}

	m, ok := tm.(Prompt)
	if !ok {
		return nil, ErrModelConversion
	}

	if m.err != nil {
		return nil, m.err
	}

	return m.model, nil
}
