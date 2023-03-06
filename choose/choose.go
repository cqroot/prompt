package choose

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/cqroot/prompt/constants"
)

type Model struct {
	choices []string
	cursor  int

	theme          Theme
	quitting       bool
	err            error
	keys           KeyMap
	showHelp       bool
	help           help.Model
	teaProgramOpts []tea.ProgramOption
}

func New(choices []string, opts ...Option) *Model {
	m := &Model{
		choices:        choices,
		cursor:         0,
		theme:          ThemeDefault,
		quitting:       false,
		err:            nil,
		keys:           DefaultKeyMap,
		showHelp:       false,
		help:           help.New(),
		teaProgramOpts: make([]tea.ProgramOption, 0),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m Model) Data() string {
	return m.choices[m.cursor]
}

func (m Model) DataString() string {
	return m.Data()
}

func (m Model) Quitting() bool {
	return m.quitting
}

func (m Model) Error() error {
	return m.err
}

func (m Model) TeaProgramOpts() []tea.ProgramOption {
	return m.teaProgramOpts
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Prev):
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}

		case key.Matches(msg, m.keys.Next):
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case key.Matches(msg, m.keys.Confirm):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			if m.showHelp {
				m.help.ShowAll = !m.help.ShowAll
			}

		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			m.err = constants.ErrUserQuit
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	view := m.theme(m.choices, m.cursor)
	if m.showHelp {
		view += "\n"
		view += m.help.View(m.keys)
	}
	return view
}
