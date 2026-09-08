---
name: charmbracelet-bubbletea
description: Builds, debugs, and maintains Go TUIs with Charm Bubble Tea. Covers model/update/view, commands, async messages, input, rendering, and terminal lifecycle; distinguishes legacy Bubble Tea from v2.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Bubble Tea

Build responsive Bubble Tea applications using the API family already selected by the project.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Keep blocking I/O inside a returned `tea.Cmd`; return its result as a message and update state on the event loop. Capture immutable request data, not a pointer that another goroutine will mutate.
- Retain commands returned by child updates and initial focus/spinner/timer operations. Route keyboard input to the focused child; route background messages to their owners even when those children are not focused.
- Treat v2 `tea.View` as the complete declaration of terminal state, including `AltScreen` and mouse mode. Set persistent fields on every returned view where they are needed.
- Use `tea.KeyPressMsg` for press-only actions. In v2, `tea.KeyMsg` is an interface covering presses and releases; it is not universally invalid.
- Keep rendering free of I/O and model mutation. Use the application's existing value or pointer model convention; value receivers do not make slices, maps, or captured state immutable.

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
