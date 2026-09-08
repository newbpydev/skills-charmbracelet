# Deterministic tests

## Model and command tests

Instantiate the real model and send typed messages directly through Update. Assert selection, focus, state transitions, stale result rejection, and exit behavior. When a child must continue blinking or ticking, verify its command is retained. Test legitimate value receivers without requiring artificial deep copies of every slice.

Execute a command only when its dependencies are controlled and its work is bounded. Inject an HTTP client, clock, loader, or producer at an existing application boundary when needed. A returned BatchMsg contains multiple commands; do not mistake the container for every command having already run.

## Runtime and teatest

For simple integration, use tea.NewProgram with explicit input/output, no OS signal handlers, controlled dimensions, and a bounded context. Use synchronized buffers when the renderer writes concurrently with assertions. A regular bytes.Buffer is not inherently safe for concurrent reads and writes.

The legacy helper is `github.com/charmbracelet/x/exp/teatest`; the v2 helper adds `/v2`. Pin the pseudo-version, then compile against the application's selected Bubble Tea release. The v2 helper's own go.mod can still name an older release candidate; that alone does not establish incompatibility or compatibility with the final release.

Use explicit input messages, condition waits, final timeouts, and cleanup. Test the program's final model separately from output. teatest's RequireEqualOutput uses a system `diff` executable; that dependency may not exist in a Windows environment. Ordinary Go assertions can avoid it when full golden comparisons are unnecessary.

## Rendering assertions

For deterministic content, compare the model's View content or normalized semantic text at fixed dimensions/theme. For styled widths use an ANSI-aware cell measurement. Include CJK, combining marks, and emoji. Do not strip escape sequences and then treat the resulting incremental renderer output as a faithful final-screen snapshot; use a terminal emulator for that claim.

Update golden files only after inspecting intended output changes. Add a regression test that would fail for a real behavior error, not a test asserting the wording of this skill.

## Terminal and SSH checks

A PTY smoke test should start the actual app, enter keys, resize, quit, and verify restoration. Use timeouts and clean up child processes. Unix PTY evidence does not prove Windows Console behavior.

For Wish, use loopback listeners, ephemeral test keys, explicit authentication, and multiple clients. Check rejection, non-PTY behavior, session independence, and disconnect cleanup. Run race detection on asynchronous/session code on supported platforms.

## Evidence

Record command, environment, module pins, and outcome. Compiling an example, testing a model, installing a skill, and observing a model use a skill are separate results. Do not call a skill universally effective because its frontmatter validator passed.

Consult the exact pinned APIs and examples in [sources](sources.md).
