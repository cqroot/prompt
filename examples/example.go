package main

import (
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
)

func main() {
	p := prompt.New()

	_, _ = p.Ask("What's your name?").Input("input your name")
	_, _ = p.Ask("Choose an item:").Choose([]string{"Item 1", "Item 2", "Item 3"})
	_, _ = p.Ask("Choose multiple items").MultiChoose([]string{"Item 1", "Item 2", "Item 3"})
	_, _ = p.Ask("Are you sure?").Choose([]string{"Yes", "No"}, choose.WithTheme(choose.ThemeLine))
}
