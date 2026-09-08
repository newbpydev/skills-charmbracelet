# Lip Gloss version differences

Legacy Lip Gloss v1 uses `github.com/charmbracelet/lipgloss`, renderer objects, and string-based color types. The v2 module is `charm.land/lipgloss/v2`. V2 color adaptation happens when writing, not within Style.Render.

For legacy fixes, retain existing renderers and adaptive color behavior. For migration, update exported color types, renderer-dependent constructors, custom writers, and companion style structs together. Check whitespace, tree, underline, and color getter changes only where the application uses them. The versioned upgrade guide in [sources](sources.md) contains those conditional mappings.

## Pinned legacy lipgloss source

`github.com/charmbracelet/lipgloss@v1.1.0` declares Go 1.18. [API](https://pkg.go.dev/github.com/charmbracelet/lipgloss@v1.1.0), [README](https://github.com/charmbracelet/lipgloss/blob/f0e45475a64ee60d712b81145172d3739db36a93/README.md), [module](https://github.com/charmbracelet/lipgloss/blob/f0e45475a64ee60d712b81145172d3739db36a93/go.mod), [MIT license](https://github.com/charmbracelet/lipgloss/blob/f0e45475a64ee60d712b81145172d3739db36a93/LICENSE).
