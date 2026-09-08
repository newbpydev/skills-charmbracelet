# Released example map

Read the example matching the task and its selected release. Inspect the example’s own go.mod before copying: nested modules may use local replacements, and demos simplify errors, sizing, and authentication. The bundled example is an independent module.

| Task and adaptation boundary | Official example |
| --- | --- |
| Create a fresh model for each SSH session with session-specific I/O. | [wish: bubbletea](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/bubbletea/main.go) |
| Bound shutdown and handle the expected closed-server result. | [wish: graceful shutdown](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/graceful-shutdown/main.go) |
| Inspect client identity; identity display alone is not an authorization policy. | [wish: identity](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/identity/main.go) |
| Configure multiple SSH authentication methods explicitly. | [wish: multi auth](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/multi-auth/main.go) |
| Study concurrent clients and shared domain state separately from UI state. | [wish: multichat](https://github.com/charmbracelet/wish/blob/76f2bdd2ccf0f7e87f7b30bffa2558cad3df93f1/examples/multichat/main.go) |
| Create a form for each SSH session; supply application authentication. | [huh: ssh form](https://github.com/charmbracelet/huh/blob/3c0116c6fd1f578ef50a25d868d2db19bd985222/examples/ssh-form/main.go) |
