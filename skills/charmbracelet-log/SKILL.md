---
name: charmbracelet-log
description: Adds and configures structured Go logging using Charm Log, including slog integration, writers, formats, styles, and version-aware color profiles. Supports logging that must coexist with a Bubble Tea TUI.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Log

Choose an appropriate log destination and preserve structured information without interfering with terminal rendering.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Use a logger instance with an explicit writer when a component or session needs its own destination or context.
- Keep machine-readable stdout and active TUI output separate from diagnostics. A file or explicitly separate stream is safer than printing over the renderer.
- Attach structured key/value fields; avoid logging secrets or dumping request objects that contain credentials.
- Charm Log implements a slog handler; use it with `slog.New` when integrating standard-library logging.
- In v2, style values use Lip Gloss v2 and explicit color profiles use colorprofile types, not the old termenv profile types.

## References

- [Version selection](references/versions.md): read before editing imports or using an unfamiliar example.
- [Implementation patterns](references/patterns.md): read the sections relevant to the requested behavior.
- [Example index](references/examples.md): find the closest released integration before adapting code.
- [Pinned sources](references/sources.md): verify exact signatures, compatibility, and provenance.

## Runnable example

[Example module](assets/example/go.mod) and [source](assets/example/main.go) provide a small pinned implementation. Copy the whole module to a scratch directory before experimenting; run `go test ./...` from that copied module. See its [tests](assets/example/main_test.go) for behavioral assertions.

## Completion evidence

- Confirm the selected module family and the behavior changed.
- Check errors, initialization, teardown, and relevant boundaries rather than only a happy-path screenshot.
- Use concrete failures to justify new helper scripts or abstractions; standard Go tooling is sufficient for ordinary work.
- Distinguish compilation, automated behavior, and live terminal verification. State untested environments without implying their results.

Original instructions and bundled examples use the [MIT license](LICENSE.txt). Upstream links retain their own authorship and licenses.
