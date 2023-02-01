package prompt

import "github.com/charmbracelet/bubbles/key"

type listKeyMap struct {
	Prev   key.Binding
	Next   key.Binding
	Quit   key.Binding
	Choose key.Binding
}

func (k listKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Prev, k.Next, k.Choose, k.Quit,
	}
}

func (k listKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Prev, k.Next, k.Choose, k.Quit},
	}
}

type ListHandler struct {
	quitting      bool
	err           error
	cursor        int
	choiceCount   int
	message       string
	finishMessage string
	style         *ListStyle
	Prompt
}

func NewListHandler(choiceCount int, style *ListStyle, message string, finishMessage string) *ListHandler {
	lh := ListHandler{
		quitting:      false,
		err:           nil,
		cursor:        0,
		choiceCount:   choiceCount,
		message:       message,
		finishMessage: finishMessage,
		style:         style,
	}

	return &lh
}

func (p *Prompt) NewListHandler(choiceCount int, style *ListStyle) *ListHandler {
	lh := ListHandler{
		quitting:    false,
		err:         nil,
		cursor:      0,
		choiceCount: choiceCount,
		style:       style,
		Prompt:      *p,
	}

	return &lh
}

func (h *ListHandler) MoveNext() {
	h.cursor++
	if h.cursor >= h.choiceCount {
		h.cursor = 0
	}
}

func (h *ListHandler) MovePrev() {
	h.cursor--
	if h.cursor < 0 {
		h.cursor = h.choiceCount - 1
	}
}

func (h *ListHandler) Quit() {
	h.quitting = true
	h.err = ErrUserQuit
}

func (h ListHandler) Quitting() bool {
	return h.quitting
}

func (h ListHandler) Cursor() int {
	return h.cursor
}

func (h ListHandler) Message() string {
	return h.message
}

func (h ListHandler) FinishMessage() string {
	return h.finishMessage
}

func (h ListHandler) Style() *ListStyle {
	return h.style
}
