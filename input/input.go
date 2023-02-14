package input

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/styles"
)

type Model struct {
	df           string
	textInput    textinput.Model
	validateFunc ValidateFunc
	inputMode    InputMode
}

func (m Model) Data() any {
	if m.textInput.Value() == "" {
		return m.textInput.Placeholder
	} else {
		return m.textInput.Value()
	}
}

func (m Model) DataString() string {
	if m.textInput.EchoMode == EchoNormal {
		return m.Data().(string)
	}
	m.textInput.Blur()
	str := m.textInput.View()
	m.textInput.Focus()
	return str
}

func (m *Model) WithInputMode(mode InputMode) *Model {
	m.inputMode = mode
	return m
}

func (m *Model) WithEchoMode(mode EchoMode) *Model {
	m.textInput.EchoMode = mode
	return m
}

func (m *Model) WithValidateFunc(vf ValidateFunc) *Model {
	m.validateFunc = vf
	return m
}

func (m Model) KeyBindings() []key.Binding {
	return nil
}

func (m Model) UseKeyQ() bool {
	return true
}

func (m Model) UseKeyEnter() bool {
	return false
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.inputMode == InputNumber || m.inputMode == InputInteger {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			keypress := msg.String()
			if len(keypress) == 1 {
				if keypress == "." {
					if m.inputMode != InputNumber ||
						strings.Contains(m.textInput.Value(), ".") {
						return m, nil
					}
				} else {
					if !unicode.IsNumber([]rune(keypress)[0]) {
						return m, nil
					}
				}
			}
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	view := m.textInput.View()

	if m.textInput.Value() != "" && m.validateFunc != nil {
		err := m.validateFunc(m.textInput.Value())
		if err != nil {
			view = view + styles.DefaultErrorPromptPrefixStyle.Render("\n✖  ") +
				styles.DefaultNoteStyle.Render(err.Error())
		} else {
			view = view + styles.DefaultFinishPromptPrefixStyle.Render("\n✔")
		}
	}

	return view
}

func New(defaultValue string, opts ...Option) *Model {
	ti := textinput.New()
	ti.Placeholder = defaultValue
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	ti.Prompt = ""

	m := &Model{
		textInput: ti,
		df:        defaultValue,
		inputMode: InputAll,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}
