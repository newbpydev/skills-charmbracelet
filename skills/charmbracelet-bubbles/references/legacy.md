# Legacy Bubbles

Legacy `github.com/charmbracelet/bubbles` includes `v0.x` releases and the honorary `v1.0.0` release. Both the 0.21.1 fixture and the verified 1.0.0 module integrate with Bubble Tea v1 and Lip Gloss v1. Preserve the selected version for maintenance; the number alone does not imply the released v2 APIs. See [version selection](versions.md).

Legacy examples commonly assign `input.Width`, use `viewport.New(w, h)`, and configure root Bubble Tea program options imperatively. Keep those APIs for legacy maintenance. Use `charm.land/bubbles/v2` only as part of an intentional compatible migration. Read the per-component upgrade mapping in [sources](sources.md), then compile against the selected release.

## Pinned legacy bubbles source

`github.com/charmbracelet/bubbles@v0.21.1` declares Go 1.24.2. [API](https://pkg.go.dev/github.com/charmbracelet/bubbles@v0.21.1), [README](https://github.com/charmbracelet/bubbles/blob/9329772de61d80756b4f1ea3acea4000a499bf71/README.md), [module](https://github.com/charmbracelet/bubbles/blob/9329772de61d80756b4f1ea3acea4000a499bf71/go.mod), [MIT license](https://github.com/charmbracelet/bubbles/blob/9329772de61d80756b4f1ea3acea4000a499bf71/LICENSE).
