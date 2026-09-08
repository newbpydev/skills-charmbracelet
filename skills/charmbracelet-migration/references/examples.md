# Released example map

Read the example matching the task and its selected release. Inspect the example’s own go.mod before copying: nested modules may use local replacements, and demos simplify errors, sizing, and authentication. The bundled example is an independent module.

| Task and adaptation boundary | Official example |
| --- | --- |
| Manage several editors, their focus, and their cursor commands. | [bubbletea: textinputs](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/textinputs/main.go) |
| Request enhanced keys while retaining ordinary-key fallbacks. | [bubbletea: keyboard enhancements](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/keyboard-enhancements/main.go) |
| Handle version-specific mouse events and terminal mouse mode. | [bubbletea: mouse](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/mouse/main.go) |
| Understand transition color types; review global terminal queries before SSH use. | [lipgloss: compat](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/compat/bubbletea/main.go) |
| Embed the form’s compatibility model inside a root Bubble Tea model. | [huh: bubbletea](https://github.com/charmbracelet/huh/blob/3c0116c6fd1f578ef50a25d868d2db19bd985222/examples/bubbletea/main.go) |
| Create a fresh model for each SSH session with session-specific I/O. | [wish: bubbletea](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/bubbletea/main.go) |
| Integrate Charm Log through the standard library slog handler interface. | [log: slog](https://github.com/charmbracelet/log/blob/394fd9b15aba02cf7dd118b345248855174a90a4/examples/slog/main.go) |
