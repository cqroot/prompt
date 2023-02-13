package prompt_test

import (
	"bytes"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt"
)

const (
	KeyCtrlC byte = 3
	KeyCtrlD byte = 4
	KeyTab   byte = 9
	KeyCtrlS byte = 19
)

func expectedFinalView(result string) string {
	return "✔  … " + result + "\n"
}

type StringModelTestcase struct {
	Keys   []byte
	Result string
	View   string
}

type StringModelResultFunc func(*prompt.Prompt) (string, error)

func testStringModelResult(t *testing.T,
	resultFunc StringModelResultFunc,
	testcases []StringModelTestcase, confirmKey []byte,
) {
	for _, testcase := range testcases {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write(testcase.Keys)
		in.Write(confirmKey)
		p := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out))
		result, err := resultFunc(p)

		require.Nil(t, err)
		require.Equal(t, testcase.Result, result)
	}
}

func testStringModelView(t *testing.T,
	resultFunc StringModelResultFunc, initView string, confirmKey []byte,
) {
	var in bytes.Buffer
	var out bytes.Buffer
	var _initView, _finalView string

	in.Write(confirmKey)
	p := prompt.New().
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		WithTestView(&_initView, &_finalView)
	result, err := resultFunc(p)

	require.Nil(t, err)
	require.Equal(t, initView, _initView)
	require.Equal(t, expectedFinalView(result), _finalView)
}

func testStringModelViewWithHelp(t *testing.T,
	resultFunc StringModelResultFunc, initViewWithHelp string, confirmKey []byte,
) {
	var in bytes.Buffer
	var out bytes.Buffer
	var _initView, _finalView string

	in.Write(confirmKey)
	p := prompt.New().WithHelp(true).
		WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out)).
		WithTestView(&_initView, &_finalView)
	result, err := resultFunc(p)

	require.Nil(t, err)
	require.Equal(t, initViewWithHelp, _initView)
	require.Equal(t, expectedFinalView(result), _finalView)
}

func testStringModelUserQuitError(t *testing.T,
	resultFunc StringModelResultFunc, quitKeys []byte,
) {
	for _, quitKey := range quitKeys {
		var in bytes.Buffer
		var out bytes.Buffer

		in.Write([]byte{quitKey})
		p := prompt.New().
			WithProgramOptions(tea.WithInput(&in), tea.WithOutput(&out))
		_, err := resultFunc(p)

		require.Equal(t, prompt.ErrUserQuit, err)
	}
}

func testStringModel(t *testing.T,
	testcases []StringModelTestcase,
	resultFunc StringModelResultFunc,
	initView string, initViewWithHelp string,
	quitKeys []byte, confirmKey []byte,
) {
	testStringModelResult(t, resultFunc, testcases, confirmKey)
	testStringModelView(t, resultFunc, initView, confirmKey)
	testStringModelViewWithHelp(t, resultFunc, initViewWithHelp, confirmKey)
	testStringModelUserQuitError(t, resultFunc, quitKeys)
}
