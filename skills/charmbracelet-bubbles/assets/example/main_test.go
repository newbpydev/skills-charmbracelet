package main

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestFocusAndCommandPropagation(t *testing.T) {
	m := newModel()
	if !m.inputs[0].Focused() || m.Init() == nil {
		t.Fatal("missing initial focus")
	}
	n, cmd := m.Update(tea.KeyPressMsg{Code: tea.KeyTab})
	m = n.(model)
	if cmd == nil || !m.inputs[1].Focused() || m.inputs[0].Focused() {
		t.Fatal("focus command/state lost")
	}
	n, _ = m.Update(tea.KeyPressMsg{Code: 'q', Text: "q"})
	m = n.(model)
	if m.inputs[1].Value() != "q" || m.inputs[0].Value() != "" {
		t.Fatal("key routed incorrectly")
	}
	n, _ = m.Update(tea.WindowSizeMsg{Width: 1})
	m = n.(model)
	if m.inputs[0].Width() < 1 || m.inputs[1].Width() < 1 {
		t.Fatal("invalid child width")
	}
}
