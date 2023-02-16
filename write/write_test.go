package write_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cqroot/prompt/constants"
	"github.com/cqroot/prompt/tester"
	"github.com/cqroot/prompt/write"
)

type Testcase struct {
	model write.Model
	keys  []byte
	data  string
	view  string
}

func testcases() []Testcase {
	testcases := make([]Testcase, 0, 20)

	defaultVal := "default value"
	val := `abcdefghijklmnopqrstuvwxyz1.2.3.4.5.6.7.8.9.0-=~!@#$%^&*()_+[]\{}|;':",./<>?`

	testcases = append(testcases, Testcase{
		model: *write.New(defaultVal),
		view: "\n┃ \x1b[7md\x1b[0mefault value                      \n" +
			`┃                                    
┃                                    
┃                                    
┃                                    
┃                                    `,
		keys: []byte{tester.KeyCtrlD},
		data: defaultVal,
	})
	testcases = append(testcases, Testcase{
		model: *write.New(defaultVal, write.WithHelp(true)),
		view: "\n┃ \x1b[7md\x1b[0mefault value                      \n" +
			`┃                                    
┃                                    
┃                                    
┃                                    
┃                                    

ctrl+d confirm • esc quit`,
		keys: append([]byte(val), tester.KeyCtrlD),
		data: val,
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

		m, ok := tm.(write.Model)
		require.Equal(t, true, ok)

		require.Equal(t, tc.data, m.Data(), "keys: %s", tc.keys)
		require.Equal(t, tc.data, m.DataString(), "keys: %s", tc.keys)
		require.Equal(t, true, m.Quitting())
		require.Nil(t, m.Error())
	}

	for _, quitKey := range []byte{tester.KeyEsc, tester.KeyCtrlC} {
		tm := tester.Exec(t, write.New(""), []byte{quitKey}, "\n┃ \x1b[7m \x1b[0m                                  "+`
┃                                    
┃                                    
┃                                    
┃                                    
┃                                    `)
		m, ok := tm.(write.Model)
		require.Equal(t, true, ok)
		require.Equal(t, constants.ErrUserQuit, m.Error())
	}
}
