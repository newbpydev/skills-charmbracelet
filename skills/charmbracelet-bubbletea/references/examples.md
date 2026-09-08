# Released example map

Read the example matching the task and its selected release. Inspect the example’s own go.mod before copying: nested modules may use local replacements, and demos simplify errors, sizing, and authentication. The bundled example is an independent module.

| Task and adaptation boundary | Official example |
| --- | --- |
| Compose child models and retain their returned commands. | [bubbletea: composable views](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/composable-views/main.go) |
| Manage several editors, their focus, and their cursor commands. | [bubbletea: textinputs](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/textinputs/main.go) |
| Reject obsolete delayed events using message ownership or identifiers. | [bubbletea: debounce](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/debounce/main.go) |
| Yield terminal ownership to an interactive child process and handle its result. | [bubbletea: exec](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/exec/main.go) |
| Deliver messages from an external producer through Program.Send. | [bubbletea: send msg](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/send-msg/main.go) |
| Accept piped data without assuming stdin is the interactive keyboard. | [bubbletea: pipe](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/pipe/main.go) |
| Handle terminal focus events; distinguish them from editor focus. | [bubbletea: focus blur](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/focus-blur/main.go) |
| Request enhanced keys while retaining ordinary-key fallbacks. | [bubbletea: keyboard enhancements](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/keyboard-enhancements/main.go) |
| Handle version-specific mouse events and terminal mouse mode. | [bubbletea: mouse](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/mouse/main.go) |
