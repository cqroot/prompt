package input_test

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/input"
	"github.com/cqroot/prompt/tester"
	"github.com/stretchr/testify/require"
)

type Testcase struct {
	model input.Model
	keys  []byte
	data  string
	view  string
}

func testcases() []Testcase {
	testcases := make([]Testcase, 0, 20)

	defaultVal := "default value"
	val := `abcdefghijklmnopqrstuvwxyz1.2.3.4.5.6.7.8.9.0-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	testcases = append(testcases, Testcase{
		model: *input.New(defaultVal),
		view:  "\x1b[7md\x1b[0mefault value",
		keys:  []byte("\r\n"),
		data:  defaultVal,
	})
	testcases = append(testcases, Testcase{
		model: *input.New(defaultVal),
		view:  "\x1b[7md\x1b[0mefault value",
		keys:  []byte(val + "\r\n"),
		data:  val,
	})
	testcases = append(testcases, Testcase{
		model: *input.New(defaultVal, input.WithHelp(true)),
		view:  "\x1b[7md\x1b[0mefault value\n\nenter confirm • esc quit",
		keys:  []byte(val + "\r\n"),
		data:  val,
	})
	testcases = append(testcases, Testcase{
		model: *input.New(defaultVal, input.WithInputMode(input.InputInteger)),
		view:  "\x1b[7md\x1b[0mefault value",
		keys:  []byte(val + "\r\n"),
		data:  "1234567890",
	})
	testcases = append(testcases, Testcase{
		model: *input.New(defaultVal, input.WithInputMode(input.InputNumber)),
		view:  "\x1b[7md\x1b[0mefault value",
		keys:  []byte(val + "\r\n"),
		data:  "1.234567890",
	})

	return testcases
}

func TestModel(t *testing.T) {
	for _, tc := range testcases() {
		tm := tester.Exec(t,
			tc.model,
			tc.keys,
			tc.view,
		)

		m, ok := tm.(input.Model)
		require.Equal(t, true, ok)

		require.Equal(t, tc.data, m.Data(), "keys: %s", tc.keys)
		require.Equal(t, tc.data, m.DataString(), "keys: %s", tc.keys)
		require.Equal(t, true, m.Quitting())
		require.Nil(t, m.Error())
	}

	for _, quitKey := range []byte{byte(tea.KeyEsc), byte(tea.KeyCtrlC)} {
		tm := tester.Exec(t, input.New(""), []byte{quitKey}, "\x1b[7m \x1b[0m                                        ")
		m, ok := tm.(input.Model)
		require.Equal(t, true, ok)
		require.Equal(t, constants.ErrUserQuit, m.Error())
	}
}
