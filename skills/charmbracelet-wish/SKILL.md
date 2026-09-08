---
name: charmbracelet-wish
description: Builds Go SSH applications with Charm Wish and integrates Bubble Tea per SSH session. Covers middleware, authentication, PTYs, session environment, terminal capabilities, concurrent sessions, and graceful shutdown.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Wish

Serve a terminal application through SSH with independent session state and explicit server behavior.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Create a fresh model per SSH session. Use the session's input, output, dimensions, environment, and capability messages rather than process-global terminal state.
- Use Wish's Bubble Tea middleware to manage terminal integration and resize forwarding. Confirm the selected release's handler and SSH type imports.
- Inspect middleware ordering; the last supplied middleware is the outermost wrapper and is reached first for incoming sessions.
- Treat authentication as an application decision. Demo servers may accept any connection; do not present that as a secured deployment.
- Handle listener errors and graceful shutdown. End session-owned producers when the session ends, and preserve host keys between production restarts.

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
