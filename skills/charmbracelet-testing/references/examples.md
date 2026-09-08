# Test examples and implementation contracts

| Need | Start here | Adaptation boundary |
| --- | --- | --- |
| Direct model assertions and released-v2 teatest integration | [Bundled tests](../assets/example/main_test.go) | These run against this skill’s exact module pins. |
| Native helper contracts | [Pinned teatest implementation](https://github.com/charmbracelet/x/blob/3986e9119cf98efcf5809969e11ad369fddb5522/exp/teatest/v2/teatest.go) | Inspect final-model, timeout and output behavior before writing assertions. |
| Multiple component ownership | [Composable views](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/composable-views/main.go) | Turn component transitions into direct message tests. |
| External producers | [Send messages](https://github.com/charmbracelet/bubbletea/blob/73b6d91ac1c3854dd4af046ab5f9e51d3b3b4290/examples/send-msg/main.go) | Add bounded producer cleanup and race checks. |
| SSH shutdown | [Wish graceful shutdown](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/graceful-shutdown/main.go) | Use ephemeral keys, actual clients, and a deadline. |

A demonstration is not a complete test oracle. Assert the requested state or terminal behavior and distinguish a raw incremental stream from an emulated final screen. See [testing patterns](patterns.md).
