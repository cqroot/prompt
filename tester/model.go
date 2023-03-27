package tester

import (
	"bytes"
	"strconv"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"
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
