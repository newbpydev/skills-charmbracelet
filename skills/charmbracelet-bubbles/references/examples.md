# Released example map

Read the example matching the task and its selected release. Inspect the example’s own go.mod before copying: nested modules may use local replacements, and demos simplify errors, sizing, and authentication. The bundled example is an independent module.

| Task and adaptation boundary | Official example |
| --- | --- |
| Manage several editors, their focus, and their cursor commands. | [bubbletea: textinputs](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/textinputs/main.go) |
| Compose child models and retain their returned commands. | [bubbletea: composable views](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/composable-views/main.go) |
| Recalculate table dimensions after window-size messages. | [bubbletea: table resize](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/table-resize/main.go) |
| Embed rendered Markdown in a Bubble Tea application. | [bubbletea: glamour](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/glamour/main.go) |

For components without a matching example above, use the pinned component API: [list](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/list), [table](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/table), [textinput](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/textinput), [textarea](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/textarea), [viewport](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/viewport), [spinner](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/spinner), [progress](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/progress), and [filepicker](https://pkg.go.dev/charm.land/bubbles/v2@v2.2.1/filepicker). Read the component’s constructor, Update result, dimensions, and key map together.
