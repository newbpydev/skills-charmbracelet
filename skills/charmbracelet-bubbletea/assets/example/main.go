// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

type loadedMsg struct {
	id   int
	text string
	err  error
}
type model struct {
	input                    textinput.Model
	spinner                  spinner.Model
	width, height, requestID int
	loading                  bool
	result                   string
	ctx                      context.Context
	initial                  tea.Cmd
}

func newModel(ctx context.Context) model {
	input := textinput.New()
	input.Placeholder = "Type a query; Enter loads"
	input.SetWidth(36)
	input.SetVirtualCursor(true)
	initial := input.Focus()
	return model{input: input, spinner: spinner.New(), width: 60, height: 12, ctx: ctx, initial: initial}
}
func (m model) Init() tea.Cmd { return m.initial }
func load(ctx context.Context, id int, query string) tea.Cmd {
	return func() tea.Msg {
		timer := time.NewTimer(100 * time.Millisecond)
		defer timer.Stop()
		select {
		case <-ctx.Done():
			return loadedMsg{id: id, err: ctx.Err()}
		case <-timer.C:
			return loadedMsg{id: id, text: "Result: " + query}
		}
	}
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = max(0, msg.Width), max(0, msg.Height)
		m.input.SetWidth(max(1, m.width-4))
	case loadedMsg:
		if msg.id != m.requestID {
			return m, nil
		}
		m.loading = false
		if msg.err != nil {
			m.result = "Error: " + msg.err.Error()
		} else {
			m.result = msg.text
		}
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.requestID++
			cmds = append(cmds, load(m.ctx, m.requestID, m.input.Value()))
			if !m.loading {
				// A previous run can still have a tick in flight. A fresh spinner
				// ID prevents that old tick from starting a second tick chain.
				m.spinner = spinner.New()
				cmds = append(cmds, m.spinner.Tick)
			}
			m.loading = true
		}
	case spinner.TickMsg:
		if m.loading {
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
func fit(text string, w, h int) string {
	if w <= 0 || h <= 0 {
		return ""
	}
	lines := strings.Split(text, "\n")
	lines = lines[:min(len(lines), h)]
	for i := range lines {
		lines[i] = ansi.Truncate(lines[i], w, "")
	}
	return strings.Join(lines, "\n")
}
func (m model) View() tea.View {
	content := m.input.View() + "\n" + m.result + "\nEnter: load · ctrl+c: quit"
	if m.loading {
		content = m.input.View() + "\n" + m.spinner.View() + " Loading\nctrl+c: quit"
	}
	if m.width < 12 || m.height < 3 {
		content = "Resize"
	}
	v := tea.NewView(fit(content, m.width, m.height))
	v.AltScreen = true
	return v
}
func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := tea.NewProgram(newModel(ctx), tea.WithContext(ctx)).Run()
	return err
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
