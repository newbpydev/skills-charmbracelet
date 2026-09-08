# Choose the dependency family first

Inspect the nearest module, any active `go.work`, `replace` directives, and the actual imports before writing code. `go list -m -json all` exposes selected module versions and replacements; run it from the relevant module. If it cannot resolve offline, read the local manifests and cached source and state the uncertainty. Do not use a failed lookup as permission to upgrade.

| Library | Current family | Legacy family |
| --- | --- | --- |
| Bubble Tea | `charm.land/bubbletea/v2` | `github.com/charmbracelet/bubbletea` |
| Bubbles | `charm.land/bubbles/v2` | `github.com/charmbracelet/bubbles` (`v0.x` and `v1.x`) |
| Lip Gloss | `charm.land/lipgloss/v2` | `github.com/charmbracelet/lipgloss` |
| Huh | `charm.land/huh/v2` | `github.com/charmbracelet/huh` |
| Glamour | `charm.land/glamour/v2` | `github.com/charmbracelet/glamour` |
| Wish | `charm.land/wish/v2` | `github.com/charmbracelet/wish` |
| Log | `charm.land/log/v2` | `github.com/charmbracelet/log` |

Harmonica, Ultraviolet, colorprofile, and the `x` modules do not follow this blanket mapping. Inspect each module's declared path. Wish's SSH session type also depends on its selected `ssh` module; the current release uses `charm.land/ssh`.

For existing work, use the installed major and verify APIs against that exact version. A maintenance request does not imply migration. For a new application, use a compatible stable release set and check its minimum Go version. The examples in this skill are pinned baselines, not an instruction to force those patch versions into every project.

Pre-release imports and examples from older `/v2` development branches can differ from released v2. Read the versioned package source when `go doc` and a tutorial disagree. Avoid rewriting all `github.com/charmbracelet` imports. A dependency graph may legitimately contain both families indirectly; the error is passing incompatible types across an application boundary.

Bubbles also has an honorary `v1.0.0` release on its GitHub module path. It still depends on Bubble Tea 1.3.10 and Lip Gloss 1.1.0; it is not the v2 component family. The bundled legacy fixture deliberately retains 0.21.1. Verify the selected version rather than inferring compatibility from the label “v1 ecosystem.” See the [pinned v1 module](https://github.com/charmbracelet/bubbles/blob/4824effc3f91c9517c776d8200ef99a1207136e0/go.mod) and [release note](https://github.com/charmbracelet/bubbles/releases/tag/v1.0.0).
