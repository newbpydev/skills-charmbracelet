// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/huh/v2"
)

func validateName(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("enter a name")
	}
	return nil
}

func newForm(name *string) *huh.Form {
	return huh.NewForm(huh.NewGroup(huh.NewInput().Key("name").Value(name).Title("Name").Validate(validateName))).WithWidth(40)
}

type model struct {
	form     *huh.Form
	canceled bool
}

func (m model) Init() tea.Cmd { return m.form.Init() }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "ctrl+c" {
		m.canceled = true
		return m, tea.Quit
	}
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.form.WithWidth(max(1, size.Width-2))
	}
	next, cmd := m.form.Update(msg)
	if f, ok := next.(*huh.Form); ok {
		m.form = f
	}
	if m.form.State == huh.StateCompleted || m.form.State == huh.StateAborted {
		return m, tea.Batch(cmd, tea.Quit)
	}
	return m, cmd
}
func (m model) View() tea.View {
	if m.canceled || m.form.State == huh.StateAborted {
		return tea.NewView("Canceled")
	}
	if m.form.State == huh.StateCompleted {
		return tea.NewView("Hello, " + m.form.GetString("name"))
	}
	return tea.NewView(m.form.View())
}
func run() error {
	accessible := flag.Bool("accessible", false, "use ordinary accessible prompts")
	flag.Parse()
	var name string
	f := newForm(&name)
	if *accessible {
		if err := f.WithAccessible(true).Run(); err != nil {
			return err
		}
		if err := validateName(name); err != nil {
			return err
		}
		fmt.Println("Hello,", name)
		return nil
	}
	_, err := tea.NewProgram(model{form: f}).Run()
	return err
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
