package prompt

import (
	"github.com/cqroot/prompt/choose"
)

var WithChooseTheme = choose.WithTheme

// Choose lets the user choose one of the given choices.
func (p Prompt) Choose(choices []string, opts ...choose.Option) (string, error) {
	pm := choose.New(choices, opts...)

	m, err := p.Run(*pm)
	if err != nil {
		return "", err
	}
	return m.Data().(string), nil
}
