package prompt

type Option func(*Prompt)

func WithTheme(theme Theme) Option {
	return func(p *Prompt) {
		p.theme = theme
	}
}
