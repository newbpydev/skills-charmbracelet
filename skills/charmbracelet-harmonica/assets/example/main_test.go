package main

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func TestConvergesAndStops(t *testing.T) {
	m := newModel()
	for i := 0; i < 2000 && m.active; i++ {
		n, cmd := m.Update(tickMsg{m.generation})
		m = n.(model)
		if !m.active && cmd != nil {
			t.Fatal("settled animation still ticks")
		}
	}
	if m.active || m.position != m.target || m.velocity != 0 {
		t.Fatal("did not converge")
	}
}
func TestStaleTickAndReducedMotion(t *testing.T) {
	m := newModel()
	n, _ := m.Update(tea.KeyPressMsg{Code: ' ', Text: " "})
	m = n.(model)
	before := m.position
	n, cmd := m.Update(tickMsg{generation: m.generation - 1})
	if n.(model).position != before || cmd != nil {
		t.Fatal("stale tick changed state")
	}
	n, cmd = m.Update(tea.KeyPressMsg{Code: 'r', Text: "r"})
	m = n.(model)
	if m.active || cmd != nil || m.position != m.target {
		t.Fatal("reduced motion still animates")
	}
	n, cmd = m.Update(tea.KeyPressMsg{Code: ' ', Text: " "})
	m = n.(model)
	if m.active || cmd != nil || m.position != m.target || m.velocity != 0 {
		t.Fatal("retargeting reactivated reduced motion")
	}
}
