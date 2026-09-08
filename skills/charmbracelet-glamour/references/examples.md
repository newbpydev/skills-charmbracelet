# Released example map

Read the example matching the task and its selected release. Inspect the example’s own go.mod before copying: nested modules may use local replacements, and demos simplify errors, sizing, and authentication. The bundled example is an independent module.

| Task and adaptation boundary | Official example |
| --- | --- |
| Configure a renderer’s style explicitly. | [glamour: custom_renderer](https://github.com/charmbracelet/glamour/blob/95c93db04489cebf252ec856584445c7e85ee276/examples/custom_renderer/main.go) |
| Start with a minimal one-shot Markdown render. | [glamour: helloworld](https://github.com/charmbracelet/glamour/blob/95c93db04489cebf252ec856584445c7e85ee276/examples/helloworld/main.go) |
| Integrate streamed CLI input/output and check its output contract. | [glamour: stdin stdout](https://github.com/charmbracelet/glamour/blob/95c93db04489cebf252ec856584445c7e85ee276/examples/stdin-stdout/main.go) |
| Embed rendered Markdown in a Bubble Tea application. | [bubbletea: glamour](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/glamour/main.go) |
| Handle runtime color-profile changes; keep the destination’s capabilities explicit. | [bubbletea: colorprofile](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/colorprofile/main.go) |
