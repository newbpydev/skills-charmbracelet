---
name: charmbracelet-huh
description: Builds and embeds Go forms using Charm Huh. Covers validated inputs, selects, dynamic forms, themes, accessible prompts, completion and cancellation, and Bubble Tea integration across Huh versions.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Huh

Use Huh for coherent form workflows while preserving form state and terminal ownership.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Choose standalone `Run`/`RunWithContext` or embedding before integration. Do not call blocking Form.Run inside a parent Update.
- For embedding, call form.Init, retain the form returned by Update and its command, and wrap form.View's string in the parent's `tea.View`.
- In Huh 2.0.3, Form.Update returns `huh.Model` (a compatibility model), not the root v2 `tea.Model` interface. Treat the released form source as authoritative over stale README type claims.
- Handle completion and abortion separately; do not act on partially entered values as if submission succeeded.
- Accessible mode is configured at form level. Make it selectable for the user's workflow and validate it with actual prompts.

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
