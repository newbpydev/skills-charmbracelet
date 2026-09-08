---
name: charmbracelet-migration
description: Migrates Go applications between legacy Charm libraries and released v2 APIs. Covers deliberate Bubble Tea, Bubbles, Lip Gloss, Huh, Glamour, Wish, or Log upgrades and mixed-version type errors.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Migration

Migrate a compatible set of typed library boundaries while preserving the application's visible behavior.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Inventory modules, workspaces, local replacements, direct imports, and test helpers before editing. Identify the user's migration target and minimum Go toolchain.
- Characterize the old application's input, focus, layout, asynchronous flow, and exit behavior. A compiling import rewrite is not sufficient evidence of preserved behavior.
- Upgrade coupled Bubble Tea/Bubbles/Lip Gloss types together, then adapt Huh, Wish, custom styles, and testing boundaries actually used by the app.
- Use explicit module mappings; keep GitHub imports for supporting modules that still declare them. Distinguish direct type conflicts from harmless indirect coexistence.
- Verify against released code and compile at each meaningful boundary. Do not implement automated textual API rewrites without a concrete safe transformation.

## References

- [Version selection](references/versions.md): read before editing imports or using an unfamiliar example.
- [Implementation patterns](references/patterns.md): read the sections relevant to the requested behavior.
- [Example index](references/examples.md): find the closest released integration before adapting code.
- [Pinned sources](references/sources.md): verify exact signatures, compatibility, and provenance.

## Runnable example

[Example module](assets/example/go.mod) and [source](assets/example/main.go) provide a small pinned implementation. Copy the whole module to a scratch directory before experimenting; run `go test ./...` from that copied module. See its [tests](assets/example/main_test.go) for behavioral assertions.

## Completion evidence

The [working legacy module](assets/legacy/go.mod), [legacy source](assets/legacy/main.go), and [legacy tests](assets/legacy/main_test.go) characterize the same input behavior before migration. Run each fixture from its own module directory.

- Confirm the selected module family and the behavior changed.
- Check errors, initialization, teardown, and relevant boundaries rather than only a happy-path screenshot.
- Use concrete failures to justify new helper scripts or abstractions; standard Go tooling is sufficient for ordinary work.
- Distinguish compilation, automated behavior, and live terminal verification. State untested environments without implying their results.

Original instructions and bundled examples use the [MIT license](LICENSE.txt). Upstream links retain their own authorship and licenses.
