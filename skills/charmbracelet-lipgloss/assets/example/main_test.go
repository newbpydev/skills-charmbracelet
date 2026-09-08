package main

import (
	"strings"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

func TestCellBounds(t *testing.T) {
	for _, w := range []int{0, 1, 4, 12, 32} {
		for _, dark := range []bool{false, true} {
			out := render("日本語 é 👋", w, dark)
			if w >= 6 && lipgloss.Width(out) != w {
				t.Fatalf("outer width: want %d, got %d", w, lipgloss.Width(out))
			}
			for _, line := range strings.Split(out, "\n") {
				if lipgloss.Width(line) > w {
					t.Fatalf("width %d overflow: %q", w, line)
				}
			}
		}
	}
}
func TestThemeKeepsContent(t *testing.T) {
	a := ansi.Strip(render("Hello", 24, true))
	b := ansi.Strip(render("Hello", 24, false))
	if a != b || !strings.Contains(a, "Hello") {
		t.Fatal("theme changed content")
	}
}
