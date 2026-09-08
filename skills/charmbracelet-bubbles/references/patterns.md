# Component integration

## Shared contract

Store the child in the parent, invoke the child's initializer where provided, assign the child returned by Update, and retain its command. Use `tea.Batch` to combine independent commands. String-returning component views are wrapped into the parent's v2 `tea.View`; do not change every component signature to match the root interface.

Focus is application state. Call Focus/Blur when ownership changes and retain the resulting commands. Deliver only relevant keyboard events to the active editor, while keeping timers, spinners, and asynchronous filtering messages flowing to their owner. Avoid a root `q` shortcut that steals typed input or list filter text.

Initialize focus in the persistent model, not only on a value copy inside Init. In Bubbles 2.2.1, textinput.New enables its virtual cursor by default. If the application disables it, propagate its real Cursor to the parent's tea.View with layout offsets. Joining the child's View string alone does not position a real terminal cursor. An explicit SetVirtualCursor(true) is optional with this release, not a correctness requirement.

## Choose by task

| Need | Component and integration concern |
| --- | --- |
| Searchable selection | `list`: item filtering values, delegates, selected-item types, and filter focus |
| Structured rows | `table`: columns, row widths, selection, and focus; update size after resize |
| Single-line entry | `textinput`: Focus command, validation, suggestions, secret echo, width semantics |
| Multiline entry | `textarea`: focus/blur styles, wrapping, cursor, line/character limits |
| Long rendered content | `viewport`: content updates, available dimensions, scroll position |
| Indeterminate activity | `spinner`: initialize tick and forward tick messages |
| Measured completion | `progress`: retain animation commands returned when setting a target |
| Timed interaction | `timer` or `stopwatch`: interval options, initialization, lifecycle messages |
| File choice | `filepicker`: selection errors, initial location, height; selection is not authorization |
| Pagination/help | `paginator`, `help`, `key`: derive visible help from enabled bindings |

A list delegate controls item rendering and height; the list owns navigation and filtering. After replacing items, account for empty lists and invalidated selection. A table does not automatically provide all list filtering behaviors.

## Released v2 changes

Use `viewport.New(viewport.WithWidth(w), viewport.WithHeight(h))` or create then set dimensions. Several components expose `SetWidth`, `Width()`, `SetHeight`, and `Height()` instead of public dimension fields. Do not assume every component changed identically: inspect its API.

Textinput styles are accessed through the release's Styles/SetStyles API; textarea separates focused and blurred style state. Several default keymaps are functions returning fresh maps, and removed `NewModel` aliases are replaced by `New`. Progress blend/color options changed; older gradient option names are not valid v2 substitutes.

Adaptive themes need explicit background state with Lip Gloss v2. Handle a background message and rebuild the relevant style values. Cursor APIs differ by component and release; use the source for real-cursor and virtual-cursor integration instead of assuming identical fields across textinput and textarea.

## Layout and validation

Derive component content dimensions from the parent's usable area. Choose a tiny-terminal fallback instead of relying on negative or zero widths. Viewport content should be re-rendered when Markdown wrap width changes, not merely resized after rendering at the old width.

Validate user-entered data at the application boundary as well as providing editor feedback. Do not expose secrets in result views, logs, or test artifacts. Test focus changes, empty data, resize, and command propagation with the real component.

Consult the exact pinned APIs and examples in [sources](sources.md).
