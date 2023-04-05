package input_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
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

	for _, testcase := range []struct {
		model input.Model
		view  string
	}{
		{
			model: *input.New(defaultVal),
			view:  "default value",
		},
		{
			model: *input.New(defaultVal, input.WithHelp(true)),
			view:  "default value\n\nenter confirm • esc quit",
		},
		{
			model: func() input.Model {
				var tm tea.Model
				tm = input.New(defaultVal)
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test")})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyEnter})
				return tm.(input.Model)
			}(),
			view: "test                                     ",
		},
		{
			model: func() input.Model {
				var tm tea.Model
				tm = input.New(defaultVal, input.WithEchoMode(input.EchoNone))
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test")})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyEnter})
				return tm.(input.Model)
			}(),
			view: "                                     ",
		},
		{
			model: func() input.Model {
				var tm tea.Model
				tm = input.New(defaultVal, input.WithEchoMode(input.EchoPassword))
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("test")})
				tm, _ = tm.Update(tea.KeyMsg{Type: tea.KeyEnter})
				return tm.(input.Model)
			}(),
			view: "****                                     ",
		},
	} {
		require.Equal(t, testcase.view, testcase.model.View())
	}
}
