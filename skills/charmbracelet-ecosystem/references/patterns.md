# Integration decisions

## Library boundaries

Bubble Tea owns the event loop. Commands return messages; updates mutate or replace the model's state; the root view declares what the terminal should show. Components may return strings even though the root v2 model returns `tea.View`. Huh v2's form uses a compatibility model interface; wrap it in the application instead of assuming it directly implements `tea.Model`.

Keep independent commands concurrent with `tea.Batch`. Use result messages to connect operations that require each other's returned data. `tea.Sequence` orders command execution but does not magically refresh a model value captured earlier by a command closure.

## Dependency inspection

Locate all modules participating in the feature. A repository can include examples with a nested go.mod and a local replace; copying that file unchanged into another project can silently select the wrong code or fail. Check module declarations rather than assuming a GitHub repository URL is the import path.

The stable v2 core must exchange v2 model, command, component, and style types. Other libraries can have independent major versions. Use only public APIs; terminal internals used by Bubble Tea are not automatically an application extension API.

## Product defaults

Use inline output when preserving scrollback matters and full-screen mode for applications that own a screen. Do not make a global `q` shortcut consume normal text input. Provide discoverable help and visible focus. Choose a simple fallback for tiny terminals. Use state and labels as well as color to convey status.

Treat redirected output as a different output target. Preserve structured stdout when a CLI offers JSON or other machine-readable output; send diagnostics to an appropriate writer. A graphical TUI cannot by itself prove screen-reader compatibility; Huh offers an explicit accessible mode for form workflows.

## When to consult upstream

Use the pinned example index to find the closest integration and inspect its imports, dependencies, and lifecycle. Examples are demonstrations: adapt authentication, errors, output destinations, and shutdown to the actual task. Do not apply demo-only defaults universally.

Consult the exact pinned APIs and examples in [sources](sources.md).
