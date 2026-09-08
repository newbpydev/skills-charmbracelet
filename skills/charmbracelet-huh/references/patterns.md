# Huh workflows

## Standalone forms

Build groups and fields, attach validation, then call Run or RunWithContext at the application boundary. Bind values with Value pointers only when the pointed-to storage remains stable for the form's lifetime. An alternative is Key plus typed GetString/GetInt/GetBool after completion. Do not bind to a short-lived copy of a parent model and assume later copies keep independent backing values.

Use validation errors that explain how to correct the value. Keep expensive remote validation out of synchronous interaction where it would freeze the UI; coordinate it through asynchronous application state when required. Secret input values must not appear in completion banners or logs.

## Embedding in Bubble Tea v2

Store `*huh.Form`. Init returns its initialization command. On Update, call `next, cmd := m.form.Update(msg)`, retain the `*huh.Form` result, and return cmd. Root View returns `tea.NewView(m.form.View())`, with persistent terminal fields set by the parent.

Huh's `Model` alias uses a string-returning compatibility view. The form is not directly interchangeable with the released Bubble Tea v2 root model. Some official prose and older snippets describe it as tea.Model; inspect `form.go` and the full released example instead.

Read the form's State to distinguish normal interaction, completion, and abortion. Process submission once. If moving to another screen, preserve commands needed for that transition. Avoid a parent global `q` shortcut while a field is accepting text. Decide explicitly whether esc cancels the form, navigates back, or quits the app.

## Dynamic fields and layout

Dynamic titles/options/visibility must depend on the correct values and bindings; use the documented dynamic API for the selected release. Recompute dimensions on parent resize. Account for the outer app's header/footer rather than giving the form the entire screen.

Do not infer an API from another field: a select, multi-select, file picker, and text editor expose different options and value types. Handle filesystem and selection errors rather than treating the selected path as validated access permission.

## Accessibility and theming

Use `Form.WithAccessible(true)` for supported standalone prompt workflows; it changes the interaction mode rather than making a visual TUI automatically accessible. WithInput/WithOutput allow explicit streams. Offer a CLI option or documented environment setting when the application needs user control.

V2 removed field-level WithAccessible and the old accessibility package. Theme functions and style types changed with explicit background selection. Inspect the current Theme/Styles signatures rather than using an old no-argument theme function. When embedding, terminal queries belong to the parent runtime.

Consult the exact pinned APIs and examples in [sources](sources.md).

## Accessible result handling in 2.0.3

The released accessible runner delegates to fields without populating the same result map as interactive completion, and discards field-level RunAccessible errors. Bind output with Value to stable storage, validate required postconditions after Run, and do not infer successful submission solely from a nil Run error or GetString. Accessible timeouts are explicitly unsupported by this release. The bundled tests exercise value retention and EOF. See the released form.go linked from sources.

Post-validation checks the data; it cannot recover discarded cancellation or I/O errors, especially when optional or prefilled values are already valid. If reliable submission/error semantics are required, use a verified upstream fix or a prompt workflow that propagates those errors. Do not present the bundled empty-name postcondition as a general cancellation fix.
