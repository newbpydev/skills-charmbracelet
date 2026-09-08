---
name: charmbracelet-terminal
description: "Integrates advanced Charm terminal primitives in Go: Ultraviolet, x/ansi, x/term, and colorprofile. Covers custom cell rendering, terminal capabilities, ANSI-aware text handling, and low-level lifecycle work beyond ordinary Bubble Tea components."
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Terminal

Use low-level terminal primitives only where the requested capability requires them, with version and lifecycle boundaries made explicit.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Prefer the existing Bubble Tea/Lip Gloss APIs for ordinary applications. Direct Ultraviolet use introduces explicit terminal lifecycle and rendering responsibilities.
- Pin Ultraviolet's pseudo-version and inspect that revision before using examples. Its API is actively changing; a current README may not match the version selected transitively by Bubble Tea.
- Use ANSI-aware width, truncation, and wrapping for styled text. Sanitizing terminal controls is a separate operation from measuring width.
- Direct terminal input has one owner. Do not mix independent readers, raw-mode managers, or output renderers on the same terminal without an explicit handoff.
- Detect capabilities for the actual output/session and make missing enhancements degrade to ordinary interaction.

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
