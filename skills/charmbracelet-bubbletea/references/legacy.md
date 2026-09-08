# Legacy maintenance

For `github.com/charmbracelet/bubbletea` v1, the root model returns `View() string`, key presses commonly use the `tea.KeyMsg` struct, and full-screen mode can be set with `tea.WithAltScreen()`. Preserve these on a bug fix. Legacy Bubbles component APIs use fields where v2 may require methods.

For released v2, use `charm.land/bubbletea/v2`, `View() tea.View`, `tea.NewView`, and `view.AltScreen`. Keyboard structs and imperative terminal options have changed; adapt their behavior using the release upgrade guide rather than only changing imports.

Do not call a v1 pattern wrong simply because a newer release exists. Conversely, a snippet marked v2 that returns a string from the root model will not satisfy the released v2 interface. See the versioned guide in [sources](sources.md).

## Pinned legacy bubbletea source

`github.com/charmbracelet/bubbletea@v1.3.10` declares Go 1.24.0. [API](https://pkg.go.dev/github.com/charmbracelet/bubbletea@v1.3.10), [README](https://github.com/charmbracelet/bubbletea/blob/9edf69c677c7353eca5fae6d3ea3986af39717b7/README.md), [module](https://github.com/charmbracelet/bubbletea/blob/9edf69c677c7353eca5fae6d3ea3986af39717b7/go.mod), [MIT license](https://github.com/charmbracelet/bubbletea/blob/9edf69c677c7353eca5fae6d3ea3986af39717b7/LICENSE).
