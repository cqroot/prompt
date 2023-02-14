package prompt

import (
	"github.com/cqroot/prompt/multichoose"
)

// MultiChoose lets the user choose multiples from the given choices.
func (p Prompt) MultiChoose(choices []string, opts ...multichoose.Option) ([]string, error) {
	pm := multichoose.New(choices, opts...)
	m, err := p.Run(*pm)
	if err != nil {
		return nil, err
	}
	return m.Data().([]string), nil
}
