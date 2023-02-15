package multichoose

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/merrors"
)

type Model struct {
	choice  uint64
	choices []string
	cursor  int

	theme    Theme
	quitting bool
	err      error
	keys     keyMap
	showHelp bool
	help     help.Model
}

func New(choices []string, opts ...Option) *Model {
	m := &Model{
		choice:   0,
		choices:  choices,
		cursor:   0,
		theme:    ThemeDefault,
		quitting: false,
		err:      nil,
		keys:     keys(),
		showHelp: false,
		help:     help.New(),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

func (m Model) Data() any {
	result := make([]string, 0)

	for i := 0; i < len(m.choices); i++ {
		if m.isSelected(i) {
			result = append(result, m.choices[i])
		}
	}
	return result
}

func (m Model) DataString() string {
	return strings.Join(m.Data().([]string), ", ")
}

func (m Model) Quitting() bool {
	return m.quitting
}

func (m Model) Error() error {
	return m.err
}

func (m *Model) toggleChoice(index int) {
	i := uint64(index)
	if m.isSelected(index) {
		m.deselectItem(i)
	} else {
		m.selectItem(i)
	}
}

func (m Model) isSelected(index int) bool {
	i := uint64(index)
	curr := uint64(1) << i
	if (m.choice & curr) != 0 {
		return true
	} else {
		return false
	}
}

func (m *Model) selectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice | curr
}

func (m *Model) deselectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice & (^curr)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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

		case key.Matches(msg, m.keys.Choose):
			m.toggleChoice(m.cursor)

		case key.Matches(msg, m.keys.Confirm):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Help):
			if m.showHelp {
				m.help.ShowAll = !m.help.ShowAll
			}

		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			m.err = merrors.ErrUserQuit
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	view := m.theme(m.choices, m.cursor, m.isSelected)
	if m.showHelp {
		view += "\n"
		view += m.help.View(m.keys)
	}
	return view
}
