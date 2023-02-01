package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type multiChooseKeyMap struct {
	Prev    key.Binding
	Next    key.Binding
	Choose  key.Binding
	Confirm key.Binding
	Quit    key.Binding
}

func (k multiChooseKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Prev, k.Next, k.Choose, k.Confirm, k.Quit,
	}
}

func (k multiChooseKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Prev, k.Next, k.Choose, k.Confirm, k.Quit},
	}
}

type MultiChooseModel struct {
	choice  uint64
	choices []string
	result  []string
	keys    multiChooseKeyMap
	ListHandler
}

func (m *MultiChooseModel) toggleChoice(index int) {
	i := uint64(index)
	if m.isChooseed(index) {
		m.deselectItem(i)
	} else {
		m.selectItem(i)
	}
}

func (m MultiChooseModel) isChooseed(index int) bool {
	i := uint64(index)
	curr := uint64(1) << i
	if (m.choice & curr) != 0 {
		return true
	} else {
		return false
	}
}

func (m *MultiChooseModel) selectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice | curr
}

func (m *MultiChooseModel) deselectItem(index uint64) {
	curr := uint64(1) << index
	m.choice = m.choice & (^curr)
}

func (m *MultiChooseModel) quit() {
	m.quitting = true
	m.result = make([]string, 0)

	for i := 0; i < len(m.choices); i++ {
		if m.isChooseed(i) {
			m.result = append(m.result, m.choices[i])
		}
	}
}

func (m MultiChooseModel) Init() tea.Cmd {
	return nil
}

func (m MultiChooseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.Quit()
			return m, tea.Quit

		case " ":
			m.toggleChoice(m.cursor)

		case "enter":
			m.quit()
			return m, tea.Quit

		case "up", "k":
			m.MovePrev()

		case "down", "j", "tab":
			m.MoveNext()
		}
	}

	return m, nil
}

func (m MultiChooseModel) View() string {
	if m.quitting {
		return m.finishView(m.style.ChoiceStyle.Render(strings.Join(m.result, ", ")))
	}

	s := strings.Builder{}
	s.WriteString("\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			if m.isChooseed(i) {
				s.WriteString(m.Style().SelectedItemStyle.Render(fmt.Sprintf("[x] %s", m.choices[i])))
			} else {
				s.WriteString(m.Style().SelectedItemStyle.Render(fmt.Sprintf("[•] %s", m.choices[i])))
			}
		} else {
			if m.isChooseed(i) {
				s.WriteString(m.Style().ItemStyle.Render(fmt.Sprintf("[x] %s", m.choices[i])))
			} else {
				s.WriteString(m.Style().ItemStyle.Render(fmt.Sprintf("[ ] %s", m.choices[i])))
			}
		}
		s.WriteString("\n")
	}

	return m.view(s.String())
}

func (p *Prompt) NewMultiChooseModel(choices []string, style *ListStyle) *MultiChooseModel {
	multiChooseKeys := multiChooseKeyMap{
		Prev: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Next: key.NewBinding(
			key.WithKeys("down", "j", "tab"),
			key.WithHelp("↓/j/tab", "move down"),
		),
		Choose: key.NewBinding(
			key.WithKeys("space"),
			key.WithHelp("space", "choose"),
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

	model := MultiChooseModel{
		choices:     choices,
		keys:        multiChooseKeys,
		ListHandler: *p.NewListHandler(len(choices), style),
	}
	model.setKeyMap(model.keys)

	return &model
}

func (p *Prompt) MultiChooseWithStyle(choices []string, style *ListStyle) ([]string, error) {
	model := p.NewMultiChooseModel(choices, style)

	tm, err := tea.NewProgram(model).Run()
	if err != nil {
		return nil, err
	}

	m, ok := tm.(MultiChooseModel)
	if !ok {
		return nil, ErrModelConversion
	}

	if err := m.err; err != nil {
		return nil, err
	} else {
		return m.result, nil
	}
}

func (p *Prompt) MultiChoose(choices []string) ([]string, error) {
	return p.MultiChooseWithStyle(choices, NewListStyle())
}
