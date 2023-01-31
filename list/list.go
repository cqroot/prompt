package list

import "github.com/cqroot/prompt/perrors"

type ListHandler struct {
	quitting      bool
	err           error
	cursor        int
	choiceCount   int
	message       string
	finishMessage string
	style         *ListStyle
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
	h.err = perrors.ErrUserQuit
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

func (h *ListHandler) SetMessage(message string) {
	h.message = message
}

func (h *ListHandler) SetFinishMessage(message string) {
	h.finishMessage = message
}

func (h *ListHandler) SetChoiceCount(choiceCount int) {
	h.choiceCount = choiceCount
}

func (h *ListHandler) SetStyle(style *ListStyle) {
	h.style = style
}
