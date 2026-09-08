package main

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestInput(t *testing.T) {
	m := newModel()
	if !m.input.Focused() || m.Init() == nil {
		t.Fatal("focus lost")
	}
	n, _ := m.Update(tea.KeyPressMsg{Code: 'q', Text: "q"})
	m = n.(model)
	if m.input.Value() != "q" {
		t.Fatal("q lost")
	}
	n, _ = m.Update(tea.WindowSizeMsg{Width: 1})
	if n.(model).input.Width() < 1 {
		t.Fatal("invalid width")
	}
}
