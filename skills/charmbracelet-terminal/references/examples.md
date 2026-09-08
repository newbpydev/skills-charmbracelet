# Terminal example map

| Need | Start here | Adaptation boundary |
| --- | --- | --- |
| Direct terminal lifecycle | [Bundled example](../assets/example/main.go) | Uses the pinned Ultraviolet revision and restores its own terminal. |
| Offscreen drawing | [Bundled tests](../assets/example/main_test.go) | Use ScreenBuffer for the Screen interface at this revision. |
| Low-level API overview | [Pinned Ultraviolet README](https://github.com/charmbracelet/ultraviolet/blob/0277a179edd9d8f94d2a2c950edde2e398012525/README.md) | Verify prose against declarations and compiled calls. |
| Exact cells and screen contracts | [Ultraviolet declarations](https://github.com/charmbracelet/ultraviolet/blob/0277a179edd9d8f94d2a2c950edde2e398012525/uv.go) and [buffer implementation](https://github.com/charmbracelet/ultraviolet/blob/0277a179edd9d8f94d2a2c950edde2e398012525/buffer.go) | Do not substitute a plain Buffer where Screen is required. |
| Capabilities inside Bubble Tea | [Color profile example](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/colorprofile/main.go) | Let the application runtime own terminal queries. |
| Layers without owning raw I/O | [Lip Gloss canvas](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/canvas/main.go) | Check whether existing rendering primitives already cover the task. |

Consult [version selection](versions.md) before adapting current-main terminal examples to an older pseudo-version.
