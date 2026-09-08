package main

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/ansi"
)

func TestMarkdownReflow(t *testing.T) {
	wide, err := renderMarkdown(document, 60)
	if err != nil {
		t.Fatal(err)
	}
	narrow, err := renderMarkdown(document, 24)
	if err != nil {
		t.Fatal(err)
	}
	if wide == narrow {
		t.Fatal("width did not reflow Markdown")
	}
	if !strings.Contains(ansi.Strip(narrow), "small document") {
		t.Fatal("lost content")
	}
}
func TestViewportResize(t *testing.T) {
	m := newModel()
	n, _ := m.Update(tea.WindowSizeMsg{Width: 24, Height: 8})
	m = n.(model)
	if m.viewport.Width() != 24 || m.viewport.Height() != 7 {
		t.Fatal("viewport not resized")
	}
	for _, line := range strings.Split(m.View().Content, "\n") {
		if ansi.StringWidth(line) > 24 {
			t.Fatal("viewport overflow")
		}
	}
	n, _ = m.Update(tea.WindowSizeMsg{Width: 1, Height: 1})
	if !n.(model).View().AltScreen {
		t.Fatal("lost terminal state on small view")
	}
}
