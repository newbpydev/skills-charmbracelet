# Released example map

Read the example matching the task and its selected release. Inspect the example’s own go.mod before copying: nested modules may use local replacements, and demos simplify errors, sizing, and authentication. The bundled example is an independent module.

| Task and adaptation boundary | Official example |
| --- | --- |
| Study layers and cell-based composition; verify the pinned canvas API. | [lipgloss: canvas](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/canvas/main.go) |
| Understand transition color types; review global terminal queries before SSH use. | [lipgloss: compat](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/compat/bubbletea/main.go) |
| Allocate borders, padding, and joined regions in terminal cells. | [lipgloss: layout](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/layout/main.go) |
| Render a static styled list; interactive selection needs a component. | [lipgloss: list](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/list/duckduckgoose/main.go) |
| Render ANSI-styled table cells and inspect their measured widths. | [lipgloss: table](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/table/ansi/main.go) |
| Style tree nodes and backgrounds without adding input handling. | [lipgloss: tree](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/tree/background/main.go) |
| Handle runtime color-profile changes; keep the destination’s capabilities explicit. | [bubbletea: colorprofile](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/colorprofile/main.go) |
