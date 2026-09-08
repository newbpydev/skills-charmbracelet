# Migration procedure

## Inventory and baseline

Read the relevant go.mod/go.work files and replacements, then inspect `go list -m -json all` and application imports. Nested example modules can select local upstream code with replace directives. Record the existing application's build/tests and terminal behavior before migration so unrelated failures are distinguishable.

Prefer a focused migration branch and preserve unrelated work. Determine all packages exchanging tea.Model, tea.Cmd, component models, Lip Gloss styles, Huh models, SSH sessions, and test harness values. Confirm the target dependency set's Go requirement before changing go directives.

## Core transformations

| Legacy pattern | Released v2 behavior |
| --- | --- |
| GitHub core imports | `charm.land/bubbletea/v2`, `charm.land/bubbles/v2`, `charm.land/lipgloss/v2` |
| Root `View() string` | Root `View() tea.View`; wrap text with `tea.NewView` |
| `tea.WithAltScreen()` and terminal toggle commands | Set corresponding View fields whenever that state applies |
| KeyMsg struct for presses | KeyPressMsg; KeyMsg is now an interface for press/release |
| Old key fields and space string | Verify v2 Key fields; space string is `space` |
| MouseMsg struct fields | Specific event messages or MouseMsg interface access |
| Direct width/height fields in affected components | Component-specific setters/getters |
| `viewport.New(w, h)` | Option-based constructor or setters |
| Lip Gloss Renderer/global profile | Styles plus terminal/output adaptation |
| Lip Gloss Color type | Color function and standard `image/color.Color` values |

Paste, cursor, keyboard enhancements, focus reporting, and terminal title behavior need attention only where used. Do not replace every valid `case tea.KeyMsg` with a blanket warning: interface matching can intentionally handle both press and release.

## Companion boundaries

Huh v2 keeps a string-returning compatibility model for Form; wrap it in a root v2 model rather than mechanically asserting that *Form implements tea.Model. Adapt themes and form-level accessibility. Glamour v2 needs explicit style choice and removes renderer color-profile options. Wish changes middleware helpers and session type dependencies; match `charm.land/ssh` for the pinned current Wish release. Log v2 styles/profile types must migrate with Lip Gloss/colorprofile.

Use `github.com/charmbracelet/x/exp/teatest/v2` for v2 test models, not the legacy teatest module. Both are experimental and require pinned versions. Harmonica and other independent GitHub modules are not covered by the core import rewrite.

## Verification and finish

Resolve dependencies intentionally, then run gofmt, build/tests, and go vet from every affected module. Do not use `go mod tidy` at an unrelated repository root. Review dependency changes and remove obsolete imports only after they are no longer used.

Exercise normal text entry including q/space/paste, focus changes, resize, background/theme behavior, external programs, concurrent results, and normal/interrupt exit as applicable. Verify terminal restoration using a real PTY. Characterization and migrated fixtures should express the same behavior across their different API families.

Consult the exact pinned APIs and examples in [sources](sources.md).
