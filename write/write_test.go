package write_test

import (
	"bytes"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/write"
	"github.com/stretchr/testify/require"
)

func TestMultiChoose(t *testing.T) {
	defaultVal := "default value"
	val := "abcdefghij\r\nklmnopqrst\r\nuvwxyz1.2.\r\n3.4.5.6.7.\r\n8.9.0-=~!@\r\n#$%^&*()_+\r\n[]\\{}|;':\",./<>?"

	for _, testcase := range []struct {
		model      write.Model
		keys       []byte
		data       string
		dataString string
	}{
		{
			model:      *write.New(defaultVal),
			keys:       []byte{byte(tea.KeyCtrlD)},
			data:       defaultVal,
			dataString: defaultVal,
		},
		{
			model:      *write.New(defaultVal),
			keys:       append([]byte(val), byte(tea.KeyCtrlD)),
			data:       val,
			dataString: "...(82 bytes)",
		},
	} {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write(testcase.keys)
		tm, err := tea.NewProgram(testcase.model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
		require.Nil(t, err)

		m, ok := tm.(write.Model)
		require.Equal(t, true, ok)

		require.Equal(t, strings.ReplaceAll(testcase.data, "\r", ""), m.Data())
		require.Equal(t, testcase.dataString, m.DataString())
		require.Equal(t, true, m.Quitting())
	}
}

func TestErrors(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	in.Write([]byte{byte(tea.KeyCtrlC)})
	tm, err := tea.NewProgram(*write.New(""), tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(write.Model)
	require.Equal(t, true, ok)

	require.Equal(t, constants.ErrUserQuit, m.Error())
}

func TestThemes(t *testing.T) {
	defaultVal := "default value"

	// The textarea renders each line padded to its full width (40 columns by
	// default), so build the expected filler explicitly instead of relying on
	// hard-coded trailing whitespace.
	pad := func(s string) string {
		return s + strings.Repeat(" ", 40-lipgloss.Width(s))
	}
	blank := pad("┃")

	for _, testcase := range []struct {
		model write.Model
		view  string
	}{
		{
			model: *write.New(defaultVal),
			view: "\n" + pad("┃ default value") + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank,
		},
		{
			model: *write.New(defaultVal, write.WithHelp(true)),
			view: "\n" + pad("┃ default value") + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				"\n" +
				"ctrl+d confirm • esc quit",
		},
		{
			model: *write.New(defaultVal, write.WithLineNumbers(true)),
			view: "\n" + pad("┃   1 default value") + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank,
		},
		{
			model: func() write.Model {
				var tm tea.Model
				tm = write.New(defaultVal, write.WithCharLimit(2))
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test")})
				return tm.(write.Model)
			}(),
			view: "\n" + pad("┃ te") + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank + "\n" +
				blank,
		},
		{
			model: *write.New(defaultVal, write.WithWidth(3)),
			view:  "\n┃ d\n┃ e\n┃ f\n┃ a\n┃ u\n┃ l",
		},
	} {
		require.Equal(t, testcase.view, testcase.model.View())
	}
}
