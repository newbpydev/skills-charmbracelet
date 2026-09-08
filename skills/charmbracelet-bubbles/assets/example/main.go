// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"fmt"
	"os"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

type model struct {
	inputs         [2]textinput.Model
	focused, width int
	initial        tea.Cmd
}

func newModel() model {
	m := model{width: 50}
	for i := range m.inputs {
		m.inputs[i] = textinput.New()
		m.inputs[i].SetVirtualCursor(true)
		m.inputs[i].SetWidth(40)
	}
	m.inputs[0].Placeholder = "Name"
	m.inputs[1].Placeholder = "Project"
	m.initial = m.inputs[0].Focus()
	return m
}
func (m model) Init() tea.Cmd { return m.initial }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = max(0, msg.Width)
		for i := range m.inputs {
			m.inputs[i].SetWidth(max(1, m.width-4))
		}
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab", "shift+tab":
			m.inputs[m.focused].Blur()
			m.focused = (m.focused + 1) % len(m.inputs)
			return m, m.inputs[m.focused].Focus()
		}
	}
	var cmds []tea.Cmd
	for i := range m.inputs {
		// Keys only go to the active editor; blink and other events still reach their owner.
		if _, key := msg.(tea.KeyMsg); key && i != m.focused {
			continue
		}
		var cmd tea.Cmd
		m.inputs[i], cmd = m.inputs[i].Update(msg)
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}
func (m model) View() tea.View {
	if m.width < 8 {
		return tea.NewView(ansi.Truncate("Resize", m.width, ""))
	}
	lines := []string{m.inputs[0].View(), m.inputs[1].View(), "tab: focus · ctrl+c: quit"}
	for i := range lines {
		lines[i] = ansi.Truncate(lines[i], m.width, "")
	}
	return tea.NewView(strings.Join(lines, "\n"))
}
func main() {
	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
