package prompt

import (
	"github.com/cqroot/prompt/list"
)

func (p Prompt) Toggle(choices []string) (string, error) {
	return p.ToggleWithStyle(choices, list.NewListStyle())
}

func (p Prompt) ToggleWithStyle(choices []string, style *list.ListStyle) (string, error) {
	return list.ToggleWithStyle(choices, style, p.view(), p.finishView())
}

func (p Prompt) Choose(choices []string) (string, error) {
	return p.ChooseWithStyle(choices, list.NewListStyle())
}

func (p Prompt) ChooseWithStyle(choices []string, style *list.ListStyle) (string, error) {
	return list.ChooseWithStyle(choices, style, p.view(), p.finishView())
}
