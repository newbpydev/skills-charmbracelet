# Bubble Tea implementation patterns

## Initialization and composition

Create child state before starting the program. Collect initial commands with `tea.Batch`, including the command returned from input focus and each animated component's initializer. For each update, retain both the returned child model and its command. Components with a typed `Update` and string `View` are not necessarily root `tea.Model` implementations.

Focus state must persist in the model passed to the program. Calling a mutating Focus method inside a value-receiver Init changes a copy. Set focus while constructing the model and save its command for Init, or use an appropriate pointer model consistently.

For Bubbles 2.2.1 textinput, `New()` already enables the virtual cursor. Preserve that default, or make the choice explicit with `SetVirtualCursor(true)`. If the app disables it, attach `input.Cursor()` to the root View and adjust its coordinates for the input's rendered position. A string-only composition does not propagate a real cursor automatically. Do not report a missing SetVirtualCursor call as a bug without checking the selected release's default.

Route system changes such as window size and background color to the components that depend on them. Keyboard focus is not a filter for asynchronous results: a hidden spinner, download, or timer can still own a message. Avoid duplicate tick chains when a child is shown again.

## Commands and concurrency

A function returning `tea.Cmd` may be entirely nonblocking: inspect where the I/O actually executes before flagging it. A command's return value is a message; only the update loop changes shared model state. `tea.Batch` has no result ordering guarantee. `tea.Sequence` runs commands in order, but a subsequent command that needs new result data is usually constructed after receiving that result.

Assign each refresh a request identifier and include it in completion/error messages. Ignore superseded results. Cancellation belongs to the operation's context; `tea.WithContext` controls the program, but does not automatically cancel every arbitrary HTTP request. Use explicit request contexts and release resources when replacing work or exiting.

`tea.Tick` schedules one message. Return the next tick command from the handler if repetition is desired, and stop doing so when finished. For existing producers, `Program.Send` is a bridge; do not mutate the model from the producer. Consider producer shutdown and whether the program has started.

## Input and terminal state in v2

Handle `tea.KeyPressMsg` for shortcuts. Printable input belongs to the focused editor; reserve an application exit combination such as ctrl+c deliberately. The space key string is `space` in v2. Keyboard enhancement support varies by terminal; support ordinary key input without assuming releases or every modifier combination will be reported.

Paste and mouse events have their own message types. Do not model an arbitrary paste as one normal rune key. Use the selected release's public mouse interfaces and fields rather than old `MouseMsg.X` struct access. Derive hit areas from the same layout that renders them.

Build a `tea.NewView(content)` and set `AltScreen`, `MouseMode`, focus reporting, title, cursor, or keyboard enhancements only as required. Re-declare persistent state each time; switching loading/error/empty views must not accidentally leave full-screen mode. Request terminal background information through Bubble Tea and handle its message; do not launch a second stdin reader during a running program.

## Resize and rendering

Calculate available content space after padding, borders, and fixed headers/footers. Clamp sizes according to component semantics: zero can mean unlimited width in some components. For terminals too small for the layout, show a short fallback and omit oversized decorations. Use ANSI-aware cell measurements rather than byte or rune counts.

Keep expensive Markdown rendering and large data transformations out of every View call. Cache work by content, dimensions, and theme; invalidate when those inputs change. Profile before replacing the renderer or adding a custom framework.

## External programs, output, and exit

Use `tea.ExecProcess` for an interactive child process that needs terminal ownership, and handle its completion/error message. A background subprocess with captured output can run in a normal command. Printing to the same active terminal outside the rendering system can corrupt the display; log to a file or separate writer.

Return quit/interrupt commands deliberately and handle `Program.Run` errors at the application boundary. Normal shutdown must restore terminal state. Avoid `os.Exit` from inside Update because deferred cleanup will not run.

Consult the exact pinned APIs and examples in [sources](sources.md).
