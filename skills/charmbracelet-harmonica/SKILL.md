---
name: charmbracelet-harmonica
description: Animates Go terminal interfaces with Charm Harmonica springs. Covers spring motion, timestep and damping selection, Bubble Tea animation ticks, retargeting, and stopping animations when settled.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Harmonica

Add controlled spring motion without blocking the update loop or running idle animation forever.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Harmonica remains `github.com/charmbracelet/harmonica`; do not invent a charm.land/v2 import.
- Keep position and velocity in model state. Update both from each spring step.
- Match the spring's timestep to the update schedule. A rendering frame-rate setting is not itself an animation clock.
- Schedule one next tick while moving; stop when both displacement and velocity are below tolerances chosen for the visible scale.
- Handle retargeting without creating multiple tick chains. A reduced-motion mode should jump to the target and stop ticking.

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
