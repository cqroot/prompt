package input_test

import (
	"bytes"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/input"
	"github.com/stretchr/testify/require"
)

func TestChoose(t *testing.T) {
	defaultVal := "default value"
	val := `abcdefghijklmnopqrstuvwxyz1.2.3.4.5.6.7.8.9.0-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	for _, testcase := range []struct {
		model input.Model
		keys  []byte
		data  string
	}{
		{
			model: *input.New(defaultVal),
			keys:  []byte("\r\n"),
			data:  defaultVal,
		},
		{
			model: *input.New(defaultVal),
			keys:  []byte(val + "\r\n"),
			data:  val,
		},
		{
			model: *input.New(defaultVal, input.WithInputMode(input.InputInteger)),
			keys:  []byte(val + "\r\n"),
			data:  "1234567890",
		},
		{
			model: *input.New(defaultVal, input.WithInputMode(input.InputNumber)),
			keys:  []byte(val + "\r\n"),
			data:  "1.234567890",
		},
	} {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write(testcase.keys)
		tm, err := tea.NewProgram(testcase.model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
		require.Nil(t, err)

		m, ok := tm.(input.Model)
		require.Equal(t, true, ok)

		require.Equal(t, testcase.data, m.Data())
		require.Equal(t, testcase.data, m.DataString())
		require.Equal(t, true, m.Quitting())
	}
}

func TestErrors(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	in.Write([]byte{byte(tea.KeyCtrlC)})
	tm, err := tea.NewProgram(*input.New(""), tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	m, ok := tm.(input.Model)
	require.Equal(t, true, ok)

	require.Equal(t, constants.ErrUserQuit, m.Error())
}

func TestThemes(t *testing.T) {
	defaultVal := "default value"

	pad := func(width int, s string) string {
		return s + strings.Repeat(" ", width-lipgloss.Width(s))
	}

	for _, testcase := range []struct {
		model input.Model
		view  string
	}{
		{
			model: *input.New(defaultVal),
			view:  pad(40, "default value"),
		},
		{
			model: *input.New(defaultVal, input.WithHelp(true)),
			view:  pad(40, "default value") + "\n\nenter confirm • esc quit",
		},
		{
			model: func() input.Model {
				var tm tea.Model
				tm = input.New(defaultVal)
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test")})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyEnter})
				return tm.(input.Model)
			}(),
			view: pad(41, "test"),
		},
		{
			model: func() input.Model {
				var tm tea.Model
				tm = input.New(defaultVal, input.WithEchoMode(input.EchoNone))
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test")})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyEnter})
				return tm.(input.Model)
			}(),
			view: pad(37, ""),
		},
		{
			model: func() input.Model {
				var tm tea.Model
				tm = input.New(defaultVal, input.WithEchoMode(input.EchoPassword))
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test")})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyEnter})
				return tm.(input.Model)
			}(),
			view: pad(41, "****"),
		},
	} {
		require.Equal(t, testcase.view, testcase.model.View())
	}
}

func TestBatchedRunesWithInputMode(t *testing.T) {
	var tm tea.Model = *input.New("default", input.WithInputMode(input.InputInteger))
	tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("ab12cd34")})
	m := tm.(input.Model)

	require.Equal(t, "1234", m.Data())
}
