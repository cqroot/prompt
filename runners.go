package prompt

import (
	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/input"
	"github.com/cqroot/prompt/multichoose"
	"github.com/cqroot/prompt/write"
)

// Choose lets the user choose one of the given choices.
func (p Prompt) Choose(choices []string, opts ...choose.Option) (string, error) {
	pm := choose.New(choices, opts...)

	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}

// MultiChoose lets the user choose multiples from the given choices.
func (p Prompt) MultiChoose(choices []string, opts ...multichoose.Option) ([]string, error) {
	pm := multichoose.New(choices, opts...)
	m, err := p.Run(*pm)
	if err != nil {
		return nil, err
	}
	return m.Data().([]string), nil
}

// Input asks the user to enter a string.
func (p Prompt) Input(defaultValue string, opts ...input.Option) (string, error) {
	pm := input.New(defaultValue, opts...)

	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}

func (p Prompt) Write(defaultValue string) (string, error) {
	pm := write.New(defaultValue)

	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}
