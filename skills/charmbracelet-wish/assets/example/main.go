// Original example: Copyright (c) 2026 Xoomby contributors. MIT licensed.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/ssh"
	"charm.land/wish/v2"
	"charm.land/wish/v2/activeterm"
	"charm.land/wish/v2/bubbletea"
	gossh "golang.org/x/crypto/ssh"
)

type model struct{ count, width int }

func (m model) Init() tea.Cmd { return nil }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = max(0, msg.Width)
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "+":
			m.count++
		}
	}
	return m, nil
}
func (m model) View() tea.View {
	status := "Waiting for input"
	if m.count > 0 {
		status = "Increment accepted"
	}
	v := tea.NewView(fmt.Sprintf("count=%d width=%d\n%s\n+: increment · q: quit", m.count, m.width, status))
	v.AltScreen = true
	return v
}
func newServer(address, hostKeyPath string, allowed ssh.PublicKey) (*ssh.Server, error) {
	if allowed == nil {
		return nil, errors.New("an authorized client public key is required")
	}
	return wish.NewServer(
		wish.WithAddress(address),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithPublicKeyAuth(func(_ ssh.Context, key ssh.PublicKey) bool { return ssh.KeysEqual(key, allowed) }),
		wish.WithMiddleware(bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
			pty, _, _ := s.Pty()
			return model{width: pty.Window.Width}, nil
		}), activeterm.Middleware()),
	)
}
func run() error {
	keyPath := flag.String("authorized-key", "", "path to one authorized client public key")
	hostKey := flag.String("host-key", ".ssh/example_ed25519", "persistent server host key path")
	address := flag.String("address", "127.0.0.1:23234", "listen address")
	flag.Parse()
	if *keyPath == "" {
		return errors.New("provide -authorized-key with a client public key file")
	}
	data, err := os.ReadFile(*keyPath)
	if err != nil {
		return err
	}
	key, _, _, _, err := gossh.ParseAuthorizedKey(data)
	if err != nil {
		return err
	}
	server, err := newServer(*address, *hostKey, key)
	if err != nil {
		return err
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	serve := make(chan error, 1)
	go func() { serve <- server.ListenAndServe() }()
	select {
	case err := <-serve:
		if errors.Is(err, ssh.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		shutdown, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdown); err != nil {
			_ = server.Close()
			return err
		}
		return nil
	}
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
