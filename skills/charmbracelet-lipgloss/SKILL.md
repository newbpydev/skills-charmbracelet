---
name: charmbracelet-lipgloss
description: Styles and lays out Go terminal output with Charm Lip Gloss. Covers borders, spacing, alignment, terminal widths, themes, colors, tables, lists, trees, and legacy-to-v2 styling differences.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Lip Gloss

Render readable terminal layouts with correct cell geometry and output-specific color handling.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Measure rendered cell widths with Lip Gloss or ANSI-aware utilities; `len` and rune counts do not measure visible terminal columns.
- Account for borders, margins, and padding before sizing children. Provide a compact layout when the available screen cannot fit decoration.
- In v2, `Color` is a function returning `image/color.Color`; `Renderer` and root adaptive-color types from v1 are not the same API.
- Separate deterministic styling from output capability adaptation. Bubble Tea v2 handles downsampling for its own view; standalone output needs an appropriate writer.
- Choose light/dark state from the actual terminal/session. Avoid process-global terminal probing in an embedded TUI or shared SSH server.

## References

- [Version selection](references/versions.md): read before editing imports or using an unfamiliar example.
- [Implementation patterns](references/patterns.md): read the sections relevant to the requested behavior.
- [Example index](references/examples.md): find the closest released integration before adapting code.
- [Pinned sources](references/sources.md): verify exact signatures, compatibility, and provenance.
- [Legacy maintenance](references/legacy.md): read when the project uses GitHub-era core imports.

## Runnable example

[Example module](assets/example/go.mod) and [source](assets/example/main.go) provide a small pinned implementation. Copy the whole module to a scratch directory before experimenting; run `go test ./...` from that copied module. See its [tests](assets/example/main_test.go) for behavioral assertions.

## Completion evidence

- Confirm the selected module family and the behavior changed.
- Check errors, initialization, teardown, and relevant boundaries rather than only a happy-path screenshot.
- Use concrete failures to justify new helper scripts or abstractions; standard Go tooling is sufficient for ordinary work.
- Distinguish compilation, automated behavior, and live terminal verification. State untested environments without implying their results.

Original instructions and bundled examples use the [MIT license](LICENSE.txt). Upstream links retain their own authorship and licenses.
