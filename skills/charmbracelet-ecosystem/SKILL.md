---
name: charmbracelet-ecosystem
description: Chooses and integrates Charm libraries for Go terminal applications. Covers ecosystem selection, dependency compatibility, or features spanning multiple Charm libraries; focused library guidance suits a single-library change.
license: MIT
metadata:
  author: Xoomby
  version: "0.1.0"
---

# Charm Ecosystem

Choose the smallest set of Charm libraries that meets the requested terminal workflow, and keep their interfaces and versions compatible.

## Ground the change

1. Find the module and any workspace/replacement selecting the code being changed. Inspect imports and selected versions before choosing an API family.
2. Preserve the requested behavior and dependency major. Resolve uncertain signatures from local `go doc` or pinned source; do not upgrade simply because a newer example exists.
3. Read only the references relevant to this task. Resolve bundled paths relative to this skill's directory, not the user's project working directory.
4. Implement the smallest change that fits the application. Retain its established architecture unless the task requires a change.
5. Build and test the affected module, then exercise terminal behavior when input, layout, or lifecycle changed. Report what was actually checked.

## Decisions that matter

- Use Bubble Tea for a stateful event loop, Bubbles for reusable interactive components, and Lip Gloss for styled output and layout.
- Use Huh for forms, Glamour for Markdown, Wish for SSH delivery, Log for diagnostics, and Harmonica for motion only when those capabilities serve the task.
- Static output can use Lip Gloss without a TUI runtime. A straightforward form can use Huh standalone. Do not introduce an event-loop architecture when the requested output is a plain CLI response.
- Keep one owner of a terminal's input and rendering. Embedded components participate in their parent's update cycle. SSH sessions each need their own model and terminal context.
- Keep library integration decisions separate from application product decisions: inline/full-screen, keyboard shortcuts, empty/error/loading states, and accessible or noninteractive output must fit the user's workflow.

## References

- [Version selection](references/versions.md): read before editing imports or using an unfamiliar example.
- [Implementation patterns](references/patterns.md): read the sections relevant to the requested behavior.
- [Example index](references/examples.md): find the closest released integration before adapting code.
- [Pinned sources](references/sources.md): verify exact signatures, compatibility, and provenance.

## Completion evidence

- Confirm the selected module family and the behavior changed.
- Check errors, initialization, teardown, and relevant boundaries rather than only a happy-path screenshot.
- Use concrete failures to justify new helper scripts or abstractions; standard Go tooling is sufficient for ordinary work.
- Distinguish compilation, automated behavior, and live terminal verification. State untested environments without implying their results.

Original instructions and bundled examples use the [MIT license](LICENSE.txt). Upstream links retain their own authorship and licenses.
