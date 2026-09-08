package main

import (
	"context"
	"strings"
	"testing"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

func TestFocusAndTyping(t *testing.T) {
	m := newModel(context.Background())
	cmd := m.Init()
	if cmd == nil {
		t.Fatal("missing initial focus command")
	}
	// Init uses a value receiver; focus must already live in the returned model.
	if !m.input.Focused() {
		t.Fatal("initial model is not focused")
	}
	next, cmd := m.Update(tea.KeyPressMsg{Code: 'q', Text: "q"})
	m = next.(model)
	if m.input.Value() != "q" {
		t.Fatal("q was not entered")
	}
	if cmd != nil {
		if _, quit := cmd().(tea.QuitMsg); quit {
			t.Fatal("q quit")
		}
	}
	_, cmd = m.Update(tea.KeyPressMsg{Code: 'c', Mod: tea.ModCtrl})
	if _, ok := cmd().(tea.QuitMsg); !ok {
		t.Fatal("ctrl+c must quit")
	}
}
func TestRefreshRejectsStaleResults(t *testing.T) {
	m := newModel(context.Background())
	m.requestID = 2
	m.loading = true
	n, _ := m.Update(loadedMsg{id: 1, text: "old"})
	m = n.(model)
	if !m.loading || m.result != "" {
		t.Fatal("stale result changed state")
	}
	n, _ = m.Update(loadedMsg{id: 2, text: "new"})
	m = n.(model)
	if m.loading || m.result != "new" {
		t.Fatal("current result missing")
	}
}
func TestTinyResizeAndViewState(t *testing.T) {
	for _, size := range [][2]int{{0, 0}, {1, 1}, {8, 2}, {40, 8}} {
		m := newModel(context.Background())
		n, _ := m.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})
		m = n.(model)
		if m.input.Width() < 1 {
			t.Fatal("invalid input width")
		}
		v := m.View()
		if !v.AltScreen {
			t.Fatal("lost full-screen state")
		}
		for _, line := range strings.Split(v.Content, "\n") {
			if ansi.StringWidth(line) > size[0] {
				t.Fatalf("overflow %v: %q", size, line)
			}
		}
	}
}
func TestCanceledLoad(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	m := load(ctx, 3, "q")().(loadedMsg)
	if m.id != 3 || m.err == nil {
		t.Fatal("cancellation lost")
	}
}
func TestOldSpinnerTickCannotRestartAfterRefresh(t *testing.T) {
	m := newModel(context.Background())
	oldTick := spinner.TickMsg{ID: m.spinner.ID()}
	next, _ := m.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	m = next.(model)
	if m.spinner.ID() == oldTick.ID {
		t.Fatal("refresh reused the previous tick owner")
	}
	before := m.spinner.View()
	next, cmd := m.Update(oldTick)
	if next.(model).spinner.View() != before || cmd != nil {
		t.Fatal("old tick restarted animation")
	}
}
