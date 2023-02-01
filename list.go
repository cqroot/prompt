package prompt

type ListHandler struct {
	cursor      int
	choiceCount int
	style       *ListStyle
}

func NewListHandler(choiceCount int, style *ListStyle) *ListHandler {
	lh := ListHandler{
		cursor:      0,
		choiceCount: choiceCount,
		style:       style,
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

func (h ListHandler) Cursor() int {
	return h.cursor
}

func (h ListHandler) Style() *ListStyle {
	return h.style
}
