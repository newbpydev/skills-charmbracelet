// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"fmt"
	"math"
	"os"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/harmonica"
)

const fps = 30

// A hundredth of a cell and cell/second are visually settled for this demo.
const tolerance = 0.01

type tickMsg struct{ generation int }
type model struct {
	spring                     harmonica.Spring
	position, velocity, target float64
	generation                 int
	active, reduced            bool
}

func newModel() model {
	return model{spring: harmonica.NewSpring(harmonica.FPS(fps), 8, 1), target: 20, active: true}
}
func tick(generation int) tea.Cmd {
	return tea.Tick(time.Second/fps, func(time.Time) tea.Msg { return tickMsg{generation} })
}
func (m model) Init() tea.Cmd {
	if m.active {
		return tick(m.generation)
	}
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "space":
			m.target = 20 - m.target
			if m.reduced {
				m.position = m.target
				m.velocity = 0
				return m, nil
			}
			m.generation++
			m.active = true
			return m, tick(m.generation)
		case "r":
			m.reduced = true
			m.position = m.target
			m.velocity = 0
			m.active = false
			m.generation++
			return m, nil
		}
	case tickMsg:
		if msg.generation != m.generation || !m.active {
			return m, nil
		}
		m.position, m.velocity = m.spring.Update(m.position, m.velocity, m.target)
		if math.Abs(m.position-m.target) < tolerance && math.Abs(m.velocity) < tolerance {
			m.position = m.target
			m.velocity = 0
			m.active = false
			return m, nil
		}
		return m, tick(m.generation)
	}
	return m, nil
}
func (m model) View() tea.View {
	return tea.NewView(fmt.Sprintf("Position: %.2f\nspace: retarget · r: reduced motion · q: quit", m.position))
}
func main() {
	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
