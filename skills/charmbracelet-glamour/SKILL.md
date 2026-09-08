---
name: charmbracelet-glamour
description: Renders Markdown in Go terminals with Charm Glamour. Covers styled Markdown, wrapping, custom styles, viewport integration, output color adaptation, and Glamour version migration.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Glamour

Render Markdown at the correct content width and adapt output to the destination without repeated expensive work.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Create a renderer with an explicit style and wrap width, check its error, then check rendering errors.
- For an embedded document, wrap at viewport content width after subtracting margins and borders. Re-render when content, width, or theme changes.
- Cache rendered Markdown outside View; a resize should invalidate the relevant cache, not recreate the renderer on every frame.
- V2 removed WithAutoStyle and WithColorProfile. Select the style from application background state and adapt terminal output through Bubble Tea or an appropriate writer.
- Preserve original Markdown as the source for reflow; wrapping previously rendered ANSI output is not equivalent to rendering at the new width.

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
