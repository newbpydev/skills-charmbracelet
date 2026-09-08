# Markdown rendering

## Renderer lifecycle

Use `glamour.NewTermRenderer` with `WithStandardStyle("dark")` or another deliberate supported style and `WithWordWrap(contentWidth)`. Check construction and Render errors. Retain the original document so a resize can render from the source again.

For static CLI output, obtain an appropriate width or use a documented default when no TTY size is available. For Bubble Tea, use its window-size message; avoid competing reads from stdin. For SSH, use the client's size/background rather than the server process's terminal.

## Viewport integration

Allocate the viewport's usable area after subtracting fixed UI elements and style frames. Cache rendering with a key containing document identity/content, wrap width, and theme. When width changes, render again and update viewport content. Decide whether the app preserves a meaningful scroll position or scrolls to new content; do not unexpectedly reset on every tick.

Small terminals may not fit configured Markdown margins, lists, or code blocks. Use a compact style or fallback and verify the actual rendered dimensions. Unicode display widths and long unbroken text deserve tests with the selected Glamour release.

## Output and styles

V2 emits deterministic styled content; it does not automatically select a light/dark style or downsample to a terminal profile during rendering. Bubble Tea v2 handles its view's output. For a standalone writer, apply colorprofile/Lip Gloss output handling as needed.

V2 removed WithAutoStyle, WithColorProfile, and the Overlined style property. Custom style configuration still requires checking the release's actual field types. The upgrade guide contains a no-argument Lip Gloss background probe that does not match Lip Gloss 2.0.6; outside a TUI the current probe takes input and output, and inside a TUI use background messages.

When rendering content from other users, define a control-sequence policy at the output boundary. Markdown rendering is not a guarantee that every terminal control sequence in arbitrary input is safe. For machine-readable or raw-Markdown output modes, return the requested format rather than terminal decoration.

Consult the exact pinned APIs and examples in [sources](sources.md).
