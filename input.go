package prompt

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InputModel struct {
	quitting bool
	err      error
	prompt   Prompt
	df       string
	result   string

	textInput         textinput.Model
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *InputModel) quit() {
	m.quitting = true
	m.result = m.textInput.Value()
	if m.result == "" {
		m.result = m.df
	}
}

func (m InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.quit()
			return m, tea.Quit

		case tea.KeyCtrlC, tea.KeyEsc:
			m.quit()
			m.err = ErrUserQuit
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m InputModel) View() string {
	if m.quitting {
		return fmt.Sprintf("%s %s\n",
			m.prompt.finishView(),
			m.ChoiceStyle.Render(m.result),
		)

	}

	return fmt.Sprintf(
		"%s %s",
		m.prompt.view(),
		m.textInput.View(),
	)
}

func NewInputModel(defaultValue string) *InputModel {
	ti := textinput.New()
	ti.Placeholder = defaultValue
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.Prompt = ""

	m := InputModel{
		textInput: ti,
		err:       nil,
		df:        defaultValue,

		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

func (p Prompt) InputWithModel(model *InputModel) (string, error) {
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return "", err
	}

	m, ok := tm.(InputModel)
	if !ok {
		return "", ErrModelConversion
	}

	if m.err != nil {
		return "", m.err
	} else {
		return m.result, nil
	}
}

func (p Prompt) Input(defaultValue string) (string, error) {
	return p.InputWithModel(NewInputModel(defaultValue))
}
