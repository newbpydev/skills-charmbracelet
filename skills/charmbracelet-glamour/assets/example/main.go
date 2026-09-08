// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/glamour/v2"
)

const document = "# A small document\n\nMarkdown is rendered again when the available width changes. 日本語 and emoji 👋 remain ordinary content.\n\n- Keep the original source\n- Resize the viewport\n- Reflow the document\n"

func renderMarkdown(doc string, width int) (string, error) {
	r, err := glamour.NewTermRenderer(glamour.WithStandardStyle("dark"), glamour.WithWordWrap(max(12, width)))
	if err != nil {
		return "", err
	}
	return r.Render(doc)
}

type model struct {
	viewport      viewport.Model
	width, height int
	err           error
}

func newModel() model {
	m := model{viewport: viewport.New(), width: 60, height: 15}
	m.resize()
	return m
}
func (m *model) resize() {
	m.viewport.SetWidth(max(1, m.width))
	m.viewport.SetHeight(max(1, m.height-1))
	content, err := renderMarkdown(document, m.width)
	m.err = err
	if err == nil {
		m.viewport.SetContent(content)
	}
}
func (m model) Init() tea.Cmd { return nil }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if m.width != msg.Width || m.height != msg.Height {
			m.width = max(0, msg.Width)
			m.height = max(0, msg.Height)
			m.resize()
		}
	case tea.KeyPressMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
func (m model) View() tea.View {
	content := m.viewport.View()
	if m.err != nil {
		content = "Markdown render failed"
	}
	if m.width < 12 || m.height < 3 {
		content = ""
	}
	v := tea.NewView(content)
	v.AltScreen = true
	return v
}
func main() {
	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
