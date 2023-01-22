package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Prompt struct {
	Text              string
	NormalPrefix      string
	FinishPrefix      string
	NormalSuffix      string
	FinishSuffix      string
	PrefixStyle       lipgloss.Style
	FinishPrefixStyle lipgloss.Style
	SuffixStyle       lipgloss.Style
	FinishSuffixStyle lipgloss.Style
}

func (p Prompt) view() string {
	return fmt.Sprintf(
		"%s %s %s",
		p.PrefixStyle.Render(p.NormalPrefix),
		p.Text,
		p.SuffixStyle.Render(p.NormalSuffix),
	)
}

func (p Prompt) finishView() string {
	return fmt.Sprintf("%s %s %s",
		p.FinishPrefixStyle.Render(p.FinishPrefix),
		p.Text,
		p.FinishSuffixStyle.Render(p.FinishSuffix),
	)
}

func (p *Prompt) Ask(text string) *Prompt {
	p.Text = text
	return p
}

func (p Prompt) ToggleWithModel(model *ToggleModel, defaultValue bool) (bool, error) {
	model.choice = defaultValue
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return defaultValue, err
	}

	m, ok := tm.(ToggleModel)
	if !ok {
		return defaultValue, ErrModelConversion
	}

	if m.err != nil {
		return defaultValue, m.err
	} else {
		return m.choice, nil
	}
}

func (p Prompt) Toggle(defaultValue bool) (bool, error) {
	return p.ToggleWithModel(NewToggleModel(), defaultValue)
}

func (p Prompt) InputWithModel(model *InputModel) (string, error) {
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(InputModel)
	if !ok {
		return "", ErrModelConversion
	}

	if m.err != nil {
		return "", m.err
	} else {
		return m.result, nil
	}
}

func (p Prompt) Input(defaultValue string) (string, error) {
	return p.InputWithModel(NewInputModel(defaultValue))
}

func (p Prompt) SelectWithModel(model *SelectModel) (string, error) {
	model.err = nil
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
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

func (p Prompt) Select(choices []string) (string, error) {
	return p.SelectWithModel(NewSelectModel(choices))
}
