# Layout and output

## Geometry

Treat terminal cells as the unit of layout. Measure styled output using `lipgloss.Width`, `Height`, or `Size`; use `GetHorizontalFrameSize` and `GetVerticalFrameSize` when allocating child content. CJK characters, combining marks, emoji, and ANSI sequences make byte counts unreliable.

Compute the outer usable area, subtract decoration once, allocate child widths, and then render. `Width` is not a universal clipping guarantee for arbitrary long content. Wrap or truncate explicitly as needed and test the final rendered size. Avoid nesting multiple components that each believe they own the full window.

Use `JoinHorizontal` and `JoinVertical` for simple compositions. Static `table`, `list`, and `tree` subpackages render structure but do not manage input. Canvas/layer APIs are advanced options; confirm their versioned signatures when stacking content or implementing overlays.

## V2 colors and writers

`lipgloss.Color("#..." )` returns a standard `color.Color`. Variables and custom APIs that used the old Lip Gloss color type need the standard interface. The v1 renderer object and its global configuration are removed. Styles are values; create reusable styles where useful, but creating a small style in View is not automatically a performance bug.

V2 Render emits full-fidelity ANSI. Bubble Tea adapts it for its terminal. For standalone output, use the supported Lip Gloss printing helpers or a `colorprofile.NewWriter` for the intended writer and environment. A `bytes.Buffer` render test does not exercise terminal capability detection.

Use `lipgloss.LightDark(isDark)` when terminal background state is already known. In a running Bubble Tea app, request/handle `tea.BackgroundColorMsg`. Outside the TUI, the current `HasDarkBackground` signature takes input and output files; some companion upgrade snippets omit those arguments and are stale.

The `compat` package provides transition types, including adaptive/complete colors, but can probe process stdin/stdout globally. That is a poor fit for multiple SSH sessions or a terminal already owned by Bubble Tea. Prefer explicit state there.

## Readability and output contracts

Use spacing, labels, focus indicators, and a limited hierarchy suited to the data. Do not communicate errors only through red coloring. Check light and dark backgrounds, lower-color output, and monochrome paths. Palette libraries such as charmtone are optional, not a mandatory dependency or universal design rule.

Respect the application's stdout contract. When structured output is requested, emit that format directly instead of wrapping it in terminal styling. When incorporating untrusted text into terminal output, consider control-sequence handling at that boundary; visible text width and safe terminal output are separate concerns.

Consult the exact pinned APIs and examples in [sources](sources.md).
