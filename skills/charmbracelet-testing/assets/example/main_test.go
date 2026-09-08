package main

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/x/exp/teatest/v2"
)

func TestModelTransition(t *testing.T) {
	next, cmd := (model{}).Update(tea.KeyPressMsg{Code: '+', Text: "+"})
	if next.(model).count != 1 || cmd != nil {
		t.Fatal("increment transition failed")
	}
	if !strings.Contains(next.(model).View().Content, "Count: 1") {
		t.Fatal("view did not reflect state")
	}
}
func TestTeatestWithReleasedV2(t *testing.T) {
	tm := teatest.NewTestModel(t, model{}, teatest.WithInitialTermSize(60, 10))
	t.Cleanup(func() { _ = tm.Quit() })
	tm.Send(tea.KeyPressMsg{Code: '+', Text: "+"})
	tm.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	final := tm.FinalModel(t, teatest.WithFinalTimeout(3*time.Second)).(model)
	if final.count != 1 {
		t.Fatalf("runtime dropped input: %d", final.count)
	}
}
func TestHeadlessRuntimeQuit(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	p := tea.NewProgram(model{}, tea.WithContext(ctx), tea.WithInput(nil), tea.WithOutput(io.Discard), tea.WithoutSignals(), tea.WithoutRenderer(), tea.WithWindowSize(60, 10))
	finished := make(chan error, 1)
	go func() { _, err := p.Run(); finished <- err }()
	p.Send(tea.KeyPressMsg{Code: 'q', Text: "q"})
	select {
	case err := <-finished:
		if err != nil {
			t.Fatal(err)
		}
	case <-ctx.Done():
		t.Fatal("program failed to finish")
	}
}
func TestHeadlessRuntimeCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	p := tea.NewProgram(model{}, tea.WithContext(ctx), tea.WithInput(nil), tea.WithOutput(io.Discard), tea.WithoutSignals(), tea.WithoutRenderer())
	finished := make(chan error, 1)
	go func() { _, err := p.Run(); finished <- err }()
	cancel()
	select {
	case err := <-finished:
		if !errors.Is(err, tea.ErrProgramKilled) || !errors.Is(err, context.Canceled) {
			t.Fatalf("lost cancellation result: %v", err)
		}
	case <-time.After(3 * time.Second):
		p.Kill()
		t.Fatal("context cancellation did not stop the program")
	}
}
