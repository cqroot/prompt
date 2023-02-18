package choose_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/choose"
	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/tester"
)

type Testcase struct {
	model choose.Model
	keys  []byte
	data  string
	view  string
}

func testcases() []Testcase {
	items := []string{"Item 1", "Item 2", "Item 3"}

	// Vertical Test Cases
	vTcs := []struct {
		keys []byte
		data string
	}{
		{keys: []byte("\r\n"), data: "Item 1"},
		{keys: []byte("kkjjj\r\n"), data: "Item 2"},
		{keys: []byte{'k', 'k', tester.KeyTab, tester.KeyTab, tester.KeyTab, '\r', '\n'}, data: "Item 2"},
	}

	// Horizontal Test Cases
	hTcs := []struct {
		keys []byte
		data string
	}{
		{keys: []byte("\r\n"), data: "Item 1"},
		{keys: []byte("hhlll\r\n"), data: "Item 2"},
		{keys: []byte{'h', 'h', tester.KeyTab, tester.KeyTab, tester.KeyTab, '\r', '\n'}, data: "Item 2"},
	}

	testcases := make([]Testcase, 0, 20)

	for _, vTc := range vTcs {
		testcases = append(testcases, Testcase{
			model: *choose.New(items),
			view:  "\n• Item 1\n  Item 2\n  Item 3\n",
			keys:  vTc.keys,
			data:  vTc.data,
		})
		testcases = append(testcases, Testcase{
			model: *choose.New(items, choose.WithHelp(true)),
			view:  "\n• Item 1\n  Item 2\n  Item 3\n\n? toggle help • enter confirm • q quit",
			keys:  vTc.keys,
			data:  vTc.data,
		})

		testcases = append(testcases, Testcase{
			model: *choose.New(items, choose.WithTheme(choose.ThemeArrow)),
			view:  "\n❯ Item 1\n  Item 2\n  Item 3\n",
			keys:  vTc.keys,
			data:  vTc.data,
		})
		testcases = append(testcases, Testcase{
			model: *choose.New(items, choose.WithTheme(choose.ThemeArrow), choose.WithHelp(true)),
			view:  "\n❯ Item 1\n  Item 2\n  Item 3\n\n? toggle help • enter confirm • q quit",
			keys:  vTc.keys,
			data:  vTc.data,
		})
	}

	for _, hTc := range hTcs {
		testcases = append(testcases, Testcase{
			model: *choose.New(items,
				choose.WithTheme(choose.ThemeLine),
				choose.WithKeyMap(choose.HorizontalKeyMap)),
			view: "Item 1 / Item 2 / Item 3\n",
			keys: hTc.keys,
			data: hTc.data,
		})
		testcases = append(testcases, Testcase{
			model: *choose.New(items,
				choose.WithTheme(choose.ThemeLine),
				choose.WithHelp(true),
				choose.WithKeyMap(choose.HorizontalKeyMap)),
			view: "Item 1 / Item 2 / Item 3\n\n? toggle help • enter confirm • q quit",
			keys: hTc.keys,
			data: hTc.data,
		})
	}

	return testcases
}

func TestModel(t *testing.T) {
	for _, tc := range testcases() {
		tm := tester.Exec(t,
			tc.model,
			tc.keys,
			tc.view,
		)

		m, ok := tm.(choose.Model)
		require.Equal(t, true, ok)

		require.Equal(t, tc.data, m.Data(), "keys: %s", tc.keys)
		require.Equal(t, tc.data, m.DataString(), "keys: %s", tc.keys)
		require.Equal(t, true, m.Quitting())
		require.Nil(t, m.Error())
	}

	for _, quitKey := range []byte{'q', tester.KeyEsc, tester.KeyCtrlC} {
		tm := tester.Exec(t, choose.New([]string{}), []byte{quitKey}, "\n")
		m, ok := tm.(choose.Model)
		require.Equal(t, true, ok)
		require.Equal(t, constants.ErrUserQuit, m.Error())
	}
}
