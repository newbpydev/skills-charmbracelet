// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/x/ansi"
)

func render(label string, width int, dark bool) string {
	if width <= 0 {
		return ""
	}
	if width < 6 {
		return ansi.Truncate(label, width, "")
	}
	c := lipgloss.LightDark(dark)(lipgloss.Color("#243047"), lipgloss.Color("#CED8ED"))
	frame := lipgloss.NewStyle().Foreground(c).Border(lipgloss.RoundedBorder()).Padding(0, 1)
	inner := max(1, width-frame.GetHorizontalFrameSize())
	return frame.Width(width).Render(ansi.Truncate(label, inner, ""))
}
func main() {
	w := colorprofile.NewWriter(os.Stdout, os.Environ())
	if _, err := fmt.Fprintln(w, render("Go · 日本語 · é · 👋", 32, true)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
