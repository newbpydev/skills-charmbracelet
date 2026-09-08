# Wish integration

## Server and middleware

Use the selected release's `wish.NewServer`, server options, and `bubbletea.Middleware` handler contract. Wish 2.0.3 uses session types from `charm.land/ssh`; some upgrade snippets still import the earlier GitHub SSH path. Match the actual server dependency.

Middleware composition wraps the preceding handler: the last middleware passed to WithMiddleware is reached first. Place logging, access checks, PTY requirements, recovery, and the application in an order that enforces the intended behavior. Authentication must be configured using actual SSH authentication options; a middleware name alone does not establish authenticated identity.

Use the activeterm middleware for a TUI that requires a PTY, and give non-PTY clients an intentional response. Restrict binding to loopback in local examples. A publicly reachable app needs its intended authentication and authorization rules before deployment.

## Per-session state

The handler constructs a fresh model and returns it with program options. Each client has its own window, background, color capabilities, and environment. Query capabilities through Bubble Tea and use session environment information, including `tea.EnvMsg` where applicable. Do not install a shared global Lip Gloss renderer or query the server's stdin to infer the client's background.

Sharing immutable application configuration is fine. Sharing mutable UI models, focused editors, or session-specific styles across connections is not. Shared domain state needs its own concurrency design, independent of each TUI's model loop.

Wish's middleware already connects session I/O and resize updates. Do not launch a second competing terminal loop or global input reader. When introducing an external process, confirm how its session context and terminal ownership are handled.

## Authentication and lifetime

Host keys authenticate the server to clients; they do not authorize clients. Configure client authentication and application authorization explicitly. Avoid silently generating a fresh production host key on every start. Keep example keys ephemeral and private.

Return immediately if server construction fails. Handle ListenAndServe errors and ignore only the expected closed-server condition. Use a bounded context for Shutdown. Cancel session-owned requests and producer goroutines on disconnect; do not wait indefinitely for a departed client's input.

## Tests

Exercise real loopback SSH sessions with ephemeral test keys. Test authorized and rejected keys, a non-PTY session, PTY input, resize forwarding, and two simultaneous independent models. Use timeouts and cleanly close clients, listeners, and programs. A server that compiles has not proven remote authentication or terminal behavior.

Consult the exact pinned APIs and examples in [sources](sources.md).

## Known dependency race in the verified snapshot

With Wish 2.0.3 and `charm.land/ssh` 0.4.3, sending an SSH `window-change` immediately after starting a shell can race a startup call to `Session.Pty()`. The race detector identified the write to `sess.pty.Window` and the unlocked read of `*sess.pty` in the pinned [SSH session source](https://github.com/charmbracelet/ssh/blob/7273fef53b8bd95059231bc60c0fc1efc0c2cde4/session.go). This is a dependency limitation, not fixed by this example.

The ordinary integration test waits for the first rendered view before resizing, so it proves established-session behavior only. The opt-in `TestEarlyResizeReproducer` preserves the failing startup scenario. From the bundled example module, run:

```sh
go test -race -run TestEarlyResizeReproducer -count=20 -args -repro-early-resize
```

Race reproduction is timing-dependent. A clean run does not prove the race is fixed. For a public service, verify an upstream fix or a reviewed dependency patch against this case before relying on the startup path. Do not describe waiting in a test as a server-side mitigation.
