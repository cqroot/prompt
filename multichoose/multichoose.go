package multichoose

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/multichoose"

	"github.com/cqroot/prompt/constants"
)

type Model struct {
	mc      *multichoose.MultiChoose
	choices []string
	cursor  int

	// scrolling viewport
	height int
	offset int

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
		mc:             multichoose.New(len(choices)),
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

func (m Model) Data() []string {
	result := make([]string, 0)

	for i := 0; i < len(m.choices); i++ {
		if m.mc.IsSelected(i) {
			result = append(result, m.choices[i])
		}
	}
	return result
}

func (m Model) DataString() string {
	return strings.Join(m.Data(), ", ")
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

// pageSize returns how many options fit in the terminal at once. It is derived
// from the terminal height, leaving room for the leading blank line and, when
// enabled, the help line. When the height is unknown it does not scroll.
func (m Model) pageSize() int {
	if m.height <= 0 {
		return len(m.choices)
	}

	size := m.height - 1 // leading blank line
	if m.showHelp {
		size-- // help line
	}
	if size < 1 {
		size = 1
	}
	return size
}

// clampOffset moves the offset so that the cursor stays inside the visible
// window. When every option fits, the offset is always zero.
func (m *Model) clampOffset() {
	size := m.pageSize()
	total := len(m.choices)

	if size >= total || size <= 0 {
		m.offset = 0
		return
	}

	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+size {
		m.offset = m.cursor - size + 1
	}

	max := total - size
	if m.offset > max {
		m.offset = max
	}
	if m.offset < 0 {
		m.offset = 0
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
		m.height = msg.Height
		m.clampOffset()

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Prev):
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
			m.clampOffset()

		case key.Matches(msg, m.keys.Next):
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
			m.clampOffset()

		case key.Matches(msg, m.keys.Choose):
			m.mc.Toggle(m.cursor)

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
	size := m.pageSize()
	total := len(m.choices)

	var view string
	if size >= total || size <= 0 {
		view = m.theme(m.choices, m.cursor, m.mc.IsSelected)
	} else {
		end := m.offset + size
		if end > total {
			end = total
		}
		window := m.choices[m.offset:end]

		isSelected := func(i int) bool {
			return m.mc.IsSelected(m.offset + i)
		}
		view = m.theme(window, m.cursor-m.offset, isSelected)
	}

	if m.showHelp {
		view += "\n"
		view += m.help.View(m.keys)
	}
	return view
}
