package main

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInput(t *testing.T) {
	m := newModel()
	if !m.input.Focused() || m.Init() == nil {
		t.Fatal("focus lost")
	}
	n, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	m = n.(model)
	if m.input.Value() != "q" {
		t.Fatal("q lost")
	}
	n, _ = m.Update(tea.WindowSizeMsg{Width: 1})
	if n.(model).input.Width < 1 {
		t.Fatal("invalid width")
	}
}
