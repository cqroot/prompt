package tester

import (
	"bytes"
	"strconv"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"
)

const (
	KeyCtrlC byte = 3
	KeyCtrlD byte = 4
	KeyTab   byte = 9
	// KeyCtrlS byte = 19
	KeyEsc byte = 27
)

func Exec(t *testing.T, model tea.Model, keys []byte, initView string) tea.Model {
	require.Equal(t, initView, model.View(), "init view with keys: %s", strconv.Quote(string(keys)))

	var in bytes.Buffer
	var out bytes.Buffer

	in.Write(keys)
	m, err := tea.NewProgram(model, tea.WithInput(&in), tea.WithOutput(&out)).Run()
	require.Nil(t, err)

	return m
}
