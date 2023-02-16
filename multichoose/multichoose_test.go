package multichoose_test

import (
	"strings"
	"testing"

	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/multichoose"
	"github.com/cqroot/prompt/tester"
	"github.com/stretchr/testify/require"
)

type Testcase struct {
	model multichoose.Model
	keys  []byte
	data  []string
	view  string
}

func testcases() []Testcase {
	items := []string{"Item 1", "Item 2", "Item 3"}

	testcases := make([]Testcase, 0, 20)
	for _, tc := range []struct {
		keys []byte
		data []string
	}{
		{keys: []byte("\r\n"), data: []string{}},
		{keys: []byte("kk jj \r\n"), data: []string{"Item 1", "Item 2"}},
		{keys: []byte("kk  jj \r\n"), data: []string{"Item 1"}},
		{keys: []byte("kk jj  \r\n"), data: []string{"Item 2"}},
		{keys: []byte("kk  jj  \r\n"), data: []string{}},
		{keys: []byte{'k', 'k', ' ', tester.KeyTab, tester.KeyTab, ' ', '\r', '\n'}, data: []string{"Item 1", "Item 2"}},
	} {
		testcases = append(testcases, Testcase{
			model: *multichoose.New(items),
			view:  "\n[•] Item 1\n[ ] Item 2\n[ ] Item 3\n",
			keys:  tc.keys,
			data:  tc.data,
		})
		testcases = append(testcases, Testcase{
			model: *multichoose.New(items, multichoose.WithHelp(true)),
			view:  "\n[•] Item 1\n[ ] Item 2\n[ ] Item 3\n\n? toggle help • space choose • enter confirm • q quit",
			keys:  tc.keys,
			data:  tc.data,
		})

		testcases = append(testcases, Testcase{
			model: *multichoose.New(items, multichoose.WithTheme(multichoose.ThemeDot)),
			view:  "\n○ Item 1\n○ Item 2\n○ Item 3\n",
			keys:  tc.keys,
			data:  tc.data,
		})
		testcases = append(testcases, Testcase{
			model: *multichoose.New(items, multichoose.WithTheme(multichoose.ThemeDot), multichoose.WithHelp(true)),
			view:  "\n○ Item 1\n○ Item 2\n○ Item 3\n\n? toggle help • space choose • enter confirm • q quit",
			keys:  tc.keys,
			data:  tc.data,
		})
	}

	return testcases
}

func TestMultiChooseModel(t *testing.T) {
	for _, tc := range testcases() {
		tm := tester.Exec(t,
			tc.model,
			tc.keys,
			tc.view,
		)

		m, ok := tm.(multichoose.Model)
		require.Equal(t, true, ok)

		require.Equal(t, tc.data, m.Data(), "keys: %s", tc.keys)
		require.Equal(t, strings.Join(tc.data, ", "), m.DataString(), "keys: %s", tc.keys)
		require.Equal(t, true, m.Quitting())
		require.Nil(t, m.Error())
	}

	for _, quitKey := range []byte{'q', tester.KeyEsc, tester.KeyCtrlC} {
		tm := tester.Exec(t, multichoose.New([]string{}), []byte{quitKey}, "\n")
		m, ok := tm.(multichoose.Model)
		require.Equal(t, true, ok)
		require.Equal(t, constants.ErrUserQuit, m.Error())
	}
}
