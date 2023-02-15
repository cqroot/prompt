package choose

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Prev    key.Binding
	Next    key.Binding
	Help    key.Binding
	Confirm key.Binding
	Quit    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Confirm, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Prev, k.Next},            // first column
		{k.Help, k.Confirm, k.Quit}, // second column
	}
}

func keys(direction Direction) keyMap {
	keys := keyMap{
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "toggle help"),
		),
		Confirm: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "confirm"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}

	if direction == DirectionH {
		keys.Prev = key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "move left"),
		)
		keys.Next = key.NewBinding(
			key.WithKeys("right", "l", "tab", " "),
			key.WithHelp("→/l/tab/space", "move right"),
		)
	}
	if direction == DirectionV {
		keys.Prev = key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		)
		keys.Next = key.NewBinding(
			key.WithKeys("down", "j", "tab"),
			key.WithHelp("↓/j/tab", "move down"),
		)
	}

	return keys
}

func WithHelp(show bool) Option {
	return func(m *Model) {
		m.showHelp = show
	}
}
