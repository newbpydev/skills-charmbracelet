# Logging integration

## Destination and format

Build a logger for the intended writer. Terminal-oriented text, JSON logs, and application result output are different contracts. Choose the relevant formatter through the release's documented API and test the result as structured data when using JSON.

During a Bubble Tea program, stderr commonly points to the same physical terminal as stdout. Merely choosing stderr does not prevent display corruption. Prefer a file or another destination when the renderer owns the terminal. Close files after the program has finished and propagate write/open errors where the application can handle them.

## Structure and scope

Use contextual fields for operation, request, and session identifiers. Use the actual error value rather than losing it through unhelpful string conversion. Add only useful details; omit credentials, private form values, and secret-bearing configuration. Keep test timestamps deterministic where they affect assertions.

Use independent logger instances or supported derived loggers for scoped configuration. Do not change global style/profile settings per concurrent SSH connection. When writing a library, return errors instead of using Fatal or process exit to handle the caller's policy.

## slog and v2

The Charm logger can be used as a `log/slog.Handler`. Construct `slog.New(handler)` and retain structured attributes and groups. Test the emitted representation instead of merely asserting that a method was called.

The v2 module is `charm.land/log/v2`. Color-profile configuration uses `github.com/charmbracelet/colorprofile`, and custom styles use `charm.land/lipgloss/v2`. Update those typed boundaries together when migrating. Keep a legacy project on its existing Log family for ordinary maintenance.

Consult the exact pinned APIs and examples in [sources](sources.md).
