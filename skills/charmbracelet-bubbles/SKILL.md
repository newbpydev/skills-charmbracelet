---
name: charmbracelet-bubbles
description: Selects, integrates, and customizes Charm Bubbles components in Go TUIs. Covers list, table, textinput, textarea, viewport, spinner, progress, timer, filepicker, focus, key maps, and component migration.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Bubbles

Compose Bubbles components without losing their state, commands, focus, or version-specific behavior.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Choose the component for its behavior: a filtered selectable list, a tabular selector, a scrollable document, or an editor are different abstractions.
- Return initial and update commands; focus and animation often require them. Process component-owned messages even while another component has keyboard focus.
- Inspect exact constructors and setters before copying old samples. Bubbles v2 requires the v2 Bubble Tea and Lip Gloss families.
- Distinguish interactive Bubbles table/list components from Lip Gloss static table/list renderers.
- Keep key bindings and displayed help synchronized, and preserve normal editing keys while an input or list filter owns focus.

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
