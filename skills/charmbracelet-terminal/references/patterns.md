# Terminal primitives

## Package boundaries

`github.com/charmbracelet/x/ansi` provides ANSI-aware string handling and terminal-sequence utilities. `github.com/charmbracelet/x/term` provides terminal detection and terminal operations. `github.com/charmbracelet/colorprofile` adapts colors for a writer/environment. These are independent modules; do not replace their imports with a made-up charm.land/v2 path.

Ultraviolet provides cells, screens, input events, and terminal rendering primitives and powers modern Charm libraries. Direct usage is appropriate for a custom renderer or primitive not covered by the application's existing framework. Avoid depending on an internal package or upgrading a transitive Ultraviolet dependency just to match the newest online example.

## Source and lifecycle

Use the exact pinned revision's README, tutorial, and package declarations together. Current Ultraviolet exposes a terminal lifecycle and screen operations, but earlier pseudo-versions can have substantially different APIs. A module path without a major suffix is not evidence of API stability.

Start and stop the terminal deliberately; ensure stop/restore runs on normal exit and error paths. Handle window-size events and keep screen bounds and rendered cells synchronized. Flush through the owning renderer. Do not call os.Exit before cleanup or leave background input readers running after returning control to the shell.

## Width, truncation, and control sequences

Use `ansi.StringWidth` and appropriate ANSI-aware truncate/wrap functions for styled text. Byte slicing can cut a UTF-8 sequence or terminal escape; rune counts do not represent grapheme/cell width. Follow the renderer's width method when exact agreement matters. Include combining text, CJK, and emoji in tests.

A width helper is not a sanitizer. If external input can contain OSC/CSI/control sequences, decide which sequences the application allows and strip or escape the rest at that boundary. Do not emit arbitrary input as a terminal command.

## Capabilities and output

Colorprofile uses the destination and environment. A server's process environment is not the SSH client's environment. Keep plain output and structured CLI responses usable when no terminal is attached. Features such as enhanced keyboard protocols, clipboard requests, mouse tracking, and background queries require capability-aware fallbacks.

In Bubble Tea applications use its public view fields, request commands, and returned messages; do not introduce a second probe that reads stdin while its event loop owns input.

Consult the exact pinned APIs and examples in [sources](sources.md).

## Offscreen screen interface at the pinned revision

Use `uv.NewScreenBuffer` when a drawing function requires `uv.Screen`. A plain `*uv.Buffer` at this revision lacks WidthMethod and does not satisfy Screen, despite broader README language about buffers. The bundled offscreen tests compile this boundary.
