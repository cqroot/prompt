package write

type Option func(*Model)

func WithHelp(show bool) Option {
	return func(m *Model) {
		m.showHelp = show
	}
}

func WithKeyMap(keyMap KeyMap) Option {
	return func(m *Model) {
		m.keys = keyMap
	}
}

func WithCharLimit(limit int) Option {
	return func(m *Model) {
		m.textarea.CharLimit = limit
	}
}
