package prompt

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ToggleModel struct {
	quitting bool
	err      error
	prompt   Prompt

	choice      bool
	TrueString  string
	FalseString string

	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	ChoiceStyle       lipgloss.Style
}

func (m ToggleModel) Init() tea.Cmd {
	return nil
}

func (m ToggleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			m.err = ErrUserQuit
			return m, tea.Quit

		case "enter":
			m.quitting = true
			return m, tea.Quit

		case "up", "down", "left", "right", "j", "k", "h", "l", "tab", "space":
			m.choice = !m.choice
		}
	}

	return m, nil
}

func (m ToggleModel) choiceToString() string {
	if m.choice {
		return m.TrueString
	} else {
		return m.FalseString
	}
}

func (m ToggleModel) View() string {
	if m.quitting {
		return fmt.Sprintf("%s %s\n",
			m.prompt.finishView(),
			m.ChoiceStyle.Render(m.choiceToString()),
		)
	}

	var toggleString string

	if m.choice {
		toggleString = fmt.Sprintf("%s / %s",
			m.SelectedItemStyle.Render(m.TrueString),
			m.ItemStyle.Render(m.FalseString),
		)
	} else {
		toggleString = fmt.Sprintf("%s / %s",
			m.ItemStyle.Render(m.TrueString),
			m.SelectedItemStyle.Render(m.FalseString),
		)
	}

	return fmt.Sprintf("%s %s",
		m.prompt.view(),
		toggleString,
	)
}

func NewToggleModel() *ToggleModel {
	m := ToggleModel{
		TrueString:        "Yes",
		FalseString:       "No",
		ItemStyle:         DefaultItemStyle,
		SelectedItemStyle: DefaultSelectedItemStyle,
		ChoiceStyle:       DefaultChoiceStyle,
	}
	return &m
}

func (p Prompt) ToggleWithModel(model *ToggleModel, defaultValue bool) (bool, error) {
	model.choice = defaultValue
	model.prompt = p

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return defaultValue, err
	}

	m, ok := tm.(ToggleModel)
	if !ok {
		return defaultValue, ErrModelConversion
	}

	if m.err != nil {
		return defaultValue, m.err
	} else {
		return m.choice, nil
	}
}

func (p Prompt) Toggle(defaultValue bool) (bool, error) {
	return p.ToggleWithModel(NewToggleModel(), defaultValue)
}
