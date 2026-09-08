// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/colorprofile"
	uv "github.com/charmbracelet/ultraviolet"
	"github.com/charmbracelet/ultraviolet/screen"
	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
)

func draw(scr uv.Screen, text string) {
	screen.Clear(scr)
	b := scr.Bounds()
	if b.Dx() <= 0 || b.Dy() <= 0 {
		return
	}
	text = ansi.Truncate(text, b.Dx(), "")
	x := max(0, (b.Dx()-ansi.StringWidth(text))/2)
	screen.NewContext(scr).DrawString(text, x, b.Dy()/2)
}
func run() (err error) {
	if !term.IsTerminal(os.Stdout.Fd()) || !term.IsTerminal(os.Stdin.Fd()) {
		_, err := fmt.Fprintln(colorprofile.NewWriter(os.Stdout, os.Environ()), "Charm terminal example")
		return err
	}
	t := uv.DefaultTerminal()
	scr := t.Screen()
	scr.EnterAltScreen()
	if err := t.Start(); err != nil {
		return err
	}
	defer func() {
		if stopErr := t.Stop(); err == nil {
			err = stopErr
		}
	}()
	for event := range t.Events() {
		switch event := event.(type) {
		case uv.WindowSizeEvent:
			scr.Resize(event.Width, event.Height)
			draw(scr, "Hello 日本語 · q to quit")
			scr.Render()
			if err := scr.Flush(); err != nil {
				return err
			}
		case uv.KeyPressEvent:
			if event.MatchString("q", "ctrl+c") {
				return nil
			}
		}
	}
	return nil
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
