package list

import tea "github.com/charmbracelet/bubbletea"

type IListBaseModel interface {
	setQuitting(bool)
	setErr(error)
	getErr() error
	setChoice(string)
	getChoice() string
	setCursor(int)
	getCursor() int
	setMessage(string)
	getMessage() string
	setFinishMessage(string)
	getFinishMessage() string
	setChoices([]string)
	setStyle(*ListStyle)
	getStyle() *ListStyle

	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
}
