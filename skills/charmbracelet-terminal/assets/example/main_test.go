package main

import (
	"strings"
	"testing"

	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/x/ansi"
)

func TestOffscreenRenderingAndUnicode(t *testing.T) {
	for _, width := range []int{0, 1, 4, 30} {
		buffer := uv.NewScreenBuffer(width, 3)
		draw(buffer, "日本語 é 👋")
		for _, line := range strings.Split(buffer.Render(), "\n") {
			if ansi.StringWidth(line) > width {
				t.Fatalf("offscreen buffer overflow at %d: %q", width, line)
			}
		}
	}
}
func TestCenterUsesCells(t *testing.T) {
	buffer := uv.NewScreenBuffer(10, 3)
	draw(buffer, "日本")
	cell := buffer.CellAt(3, 1)
	if cell == nil || cell.Content != "日" {
		t.Fatalf("CJK content not centered: %v", cell)
	}
}
