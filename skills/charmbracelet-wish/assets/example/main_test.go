package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"flag"
	"net"
	"path/filepath"
	"sync"
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"
	gossh "golang.org/x/crypto/ssh"
)

var reproduceEarlyResize = flag.Bool("repro-early-resize", false, "exercise the known ssh v0.4.3 startup resize race")

type observedOutput struct {
	mu      sync.Mutex
	data    bytes.Buffer
	changed chan struct{}
}

func (o *observedOutput) Write(p []byte) (int, error) {
	o.mu.Lock()
	n, err := o.data.Write(p)
	o.mu.Unlock()
	select {
	case o.changed <- struct{}{}:
	default:
	}
	return n, err
}
func (o *observedOutput) waitFor(t *testing.T, text string) {
	t.Helper()
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for {
		o.mu.Lock()
		found := bytes.Contains(o.data.Bytes(), []byte(text))
		o.mu.Unlock()
		if found {
			return
		}
		select {
		case <-o.changed:
		case <-timer.C:
			t.Fatalf("terminal did not render %q", text)
		}
	}
}
func signer(t *testing.T) gossh.Signer {
	t.Helper()
	_, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	s, err := gossh.NewSignerFromKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
func TestSessionModelsAreIndependent(t *testing.T) {
	a, b := model{width: 40}, model{width: 80}
	n, _ := a.Update(tea.KeyPressMsg{Code: '+', Text: "+"})
	a = n.(model)
	if a.count != 1 || b.count != 0 || b.width != 80 {
		t.Fatal("session state leaked")
	}
	n, _ = a.Update(tea.WindowSizeMsg{Width: 51, Height: 14})
	if n.(model).width != 51 || b.width != 80 {
		t.Fatal("resize did not remain session-local")
	}
}
func TestSSHAuthenticationPTYAndShutdown(t *testing.T) { exerciseSSH(t, false) }

// Opt in explicitly: this reproduces a dependency defect, not a passing gate.
func TestEarlyResizeReproducer(t *testing.T) {
	if !*reproduceEarlyResize {
		t.Skip("known ssh v0.4.3 race; pass -repro-early-resize to reproduce")
	}
	exerciseSSH(t, true)
}
func exerciseSSH(t *testing.T, earlyResize bool) {
	t.Helper()
	clientKey := signer(t)
	server, err := newServer("127.0.0.1:0", filepath.Join(t.TempDir(), "host_key"), clientKey.PublicKey())
	if err != nil {
		t.Fatal(err)
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan error, 1)
	go func() { done <- server.Serve(listener) }()
	t.Cleanup(func() {
		_ = server.Close()
		_ = listener.Close()
		select {
		case <-done:
		case <-time.After(3 * time.Second):
			t.Error("server did not stop")
		}
	})
	config := func(key gossh.Signer) *gossh.ClientConfig {
		return &gossh.ClientConfig{User: "test", Auth: []gossh.AuthMethod{gossh.PublicKeys(key)}, HostKeyCallback: gossh.FixedHostKey(server.HostSigners[0].PublicKey()), Timeout: 3 * time.Second}
	}
	if c, err := gossh.Dial("tcp", listener.Addr().String(), config(signer(t))); err == nil {
		c.Close()
		t.Fatal("unauthorized key accepted")
	}
	client, err := gossh.Dial("tcp", listener.Addr().String(), config(clientKey))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	plain, err := client.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	if err := plain.Run("ignored"); err == nil {
		t.Fatal("non-PTY client accepted")
	}
	plain.Close()

	// Both sessions are alive at the same time, with separate models and I/O.
	type activeSession struct {
		session *gossh.Session
		out     *observedOutput
		send    func(string)
	}
	var sessions []activeSession
	for i := 0; i < 2; i++ {
		session, err := client.NewSession()
		if err != nil {
			t.Fatal(err)
		}
		defer session.Close()
		out := &observedOutput{changed: make(chan struct{}, 1)}
		session.Stdout = out
		input, err := session.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		if err := session.RequestPty("xterm-256color", 12, 40+i*20, gossh.TerminalModes{}); err != nil {
			t.Fatal(err)
		}
		if err := session.Shell(); err != nil {
			t.Fatal(err)
		}
		if earlyResize {
			if err := session.WindowChange(14, 50+i*20); err != nil {
				t.Fatal(err)
			}
		}
		// SSH v0.4.3 has unsynchronized Pty reads during startup. This
		// normal lifecycle test resizes only after the first rendered view.
		out.waitFor(t, "count=0")
		sessions = append(sessions, activeSession{session, out, func(text string) {
			if _, err := input.Write([]byte(text)); err != nil {
				t.Fatal(err)
			}
		}})
	}
	for i, active := range sessions {
		if err := active.session.WindowChange(14, 50+i*20); err != nil {
			t.Fatal(err)
		}
		// A complete width value may be emitted as a differential update.
		// The unit test covers exact dimensions; real input must update state.
		active.send("+")
		active.out.waitFor(t, "Increment accepted")
		active.send("q")
		finished := make(chan error, 1)
		go func() { finished <- active.session.Wait() }()
		select {
		case err := <-finished:
			if err != nil {
				t.Fatal(err)
			}
		case <-time.After(5 * time.Second):
			t.Fatal("PTY session did not quit")
		}
	}
	if err := client.Close(); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		t.Fatal(err)
	}
}
