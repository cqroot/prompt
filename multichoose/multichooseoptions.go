package multichoose

type Option func(*Model)

func WithTheme(theme Theme) Option {
	return func(m *Model) {
		m.theme = theme
	}
}
