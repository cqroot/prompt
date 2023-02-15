package input

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Confirm key.Binding
	Quit    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Confirm, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Confirm, k.Quit},
	}
}

func keys() keyMap {
	keys := keyMap{
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
		),
		Quit: key.NewBinding(
			key.WithKeys("esc", "ctrl+c"),
			key.WithHelp("esc", "quit"),
		),
	}

	return keys
}

func WithHelp(show bool) Option {
	return func(m *Model) {
		m.showHelp = show
	}
}
