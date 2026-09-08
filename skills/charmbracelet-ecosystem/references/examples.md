# Released example map

Read the example matching the task and its selected release. Inspect the example’s own go.mod before copying: nested modules may use local replacements, and demos simplify errors, sizing, and authentication. The bundled example is an independent module.

| Task and adaptation boundary | Official example |
| --- | --- |
| Compose child models and retain their returned commands. | [bubbletea: composable views](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/composable-views/main.go) |
| Embed rendered Markdown in a Bubble Tea application. | [bubbletea: glamour](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/glamour/main.go) |
| Embed the form’s compatibility model inside a root Bubble Tea model. | [huh: bubbletea](https://github.com/charmbracelet/huh/blob/3c0116c6fd1f578ef50a25d868d2db19bd985222/examples/bubbletea/main.go) |
| Create a fresh model for each SSH session with session-specific I/O. | [wish: bubbletea](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/bubbletea/main.go) |
| Create a form for each SSH session; supply application authentication. | [huh: ssh form](https://github.com/charmbracelet/huh/blob/3c0116c6fd1f578ef50a25d868d2db19bd985222/examples/ssh-form/main.go) |
| Integrate Charm Log through the standard library slog handler interface. | [log: slog](https://github.com/charmbracelet/log/blob/394fd9b15aba02cf7dd118b345248855174a90a4/examples/slog/main.go) |
| Allocate borders, padding, and joined regions in terminal cells. | [lipgloss: layout](https://github.com/charmbracelet/lipgloss/blob/733ce53541bb688fd735524b665e0d96e35433bb/examples/layout/main.go) |
| Study spring behavior. This upstream demo uses OpenGL; use the bundled Go TUI example for terminal animation. | [harmonica: spring](https://github.com/charmbracelet/harmonica/blob/5ae57d3f7a21232590cc173f92e2f27eeab6ec43/examples/spring/opengl/main.go) |
