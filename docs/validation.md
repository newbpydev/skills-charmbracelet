# Validation evidence

Local verification on **2026-09-07**, Linux x86-64. Node **24.19.0**, Skills CLI **1.5.24**, and Go **1.25.12** for the complete check, race checks, and final PTY run. Earlier development checks also used local Go 1.27.1-X:nodwarf5. These results describe this collection and its pinned examples, not every downstream application.

## Packaging and executable checks

| Check | Outcome | Scope |
| --- | --- | --- |
| Collection validator | 12 skills, zero errors | YAML, Markdown links, standalone resources, direct reference discovery, source records |
| Skill Creator quick_validate.py with PyYAML | All 12 valid | Additional frontmatter/scaffold check |
| Node tests | 11 passed after publication preparation | Missing/escaping resources, metadata, hidden references, malformed URLs, process failures/timeouts, and catalog membership |
| Go example tests, vet, gofmt | 12 independent modules passed | 26 ordinary tests plus one opt-in upstream-race reproducer, skipped by default; manifests remain unchanged |
| Go race-enabled suite | All 12 ordinary module suites passed | Includes authenticated SSH sessions and headless runtime cancellation; excludes the documented early-resize reproducer |
| Skills CLI installation | 228 installed skill/resource checks passed | Nine logical targets, two modes, plus 12 individual standalone installs |
| Source availability | 20 module baselines, 138 pinned source URLs checked | Zero unavailable sources; one reviewed legacy-version difference |
| npm audit | Zero reported vulnerabilities | Current locked development dependency tree |
| Linux PTYs | Bubble Tea and Ultraviolet passed | Actual input, resize, exit, termios restoration and alternate-screen entry/exit |

Reproduce ordinary checks with `npm ci` and `npm run check`. Use `GOTOOLCHAIN=go1.25.12 npm run check` on Unix to repeat the minimum-toolchain run. Additional commands are `npm run check:examples -- --race`, `npm run check:upstream`, and `npm run check:pty`. Python 3 is needed only for the Linux PTY helper. Quick validation used the locally available Skill Creator script through `uv run --with pyyaml`; it is supplementary, not a repository dependency.

The source check intentionally still reports the legacy Bubbles fixture at 0.21.1 versus the GitHub module's 1.0.0 release. That honorary release was inspected and recorded separately; it remains part of the legacy Bubble Tea v1 family. The fixture was not upgraded. See [source records](../upstream-sources.json).

## Installation versus native harness behavior

The real Skills CLI installed and preserved every reference, license, source file and go.sum in disposable project directories. The targets were Claude Code, Codex, Cursor, Gemini CLI, OpenCode, GitHub Copilot, Windsurf, Cline, and Crush. Each skill was also installed alone. No user-global skill directory was changed.

Six targets share `.agents/skills`; Claude uses `.claude/skills`, Windsurf `.windsurf/skills`, and Crush `.crush/skills`. In the default-mode test, the harness directories existed first. Claude, Windsurf, and Crush received actual symlinks on Linux. In the copy-mode test, a fresh project received real directories. CLI 1.5.24 otherwise skips some non-universal project symlinks when their harness directory is absent, even if explicitly selected. The README recommends `--copy` for multi-harness installs into fresh projects.

Codex CLI **0.153.4** discovered all **12** skills in a fresh project and then read the installed Huh entrypoint, references, example and cached dependency source. Its answer correctly explained the compatibility model, accessible value binding, discarded field errors and unsupported timeout. Discovery was machine-checked; navigation and tool use were inspected manually.

Claude CLI **2.1.143** could not perform a native evaluation because its existing OAuth session returned an expired-token 401. Its file installation passed. Other harnesses received installation checks only. Native activation and model behavior have not been demonstrated across all nine harnesses.

## Generated-code evaluations

The same three tasks were run sequentially in fresh Codex processes, first without supplied skill bodies and then with the selected skills and references. Final with-skill runs were repeated after source-review corrections. CLI default-model verification reported **gpt-6-astra**, provider OpenAI, reasoning effort `none`; no model override was used. [Results](../evaluations/results.json) record prompt/code hashes, durations, supplied files and event types. Exact prompts, code and logs remain in the ignored local `.cache/evaluations` directories.

| Task | Baseline | With skill content |
| --- | --- | --- |
| New v2 search/input/loading app | Compiled; basic behavior and one-cell bounds probes passed | Compiled; basic behavior and one-cell bounds probes passed |
| Preserve and fix a legacy app | Compiled; focus/typing/quit were correct, but tiny resize assigned width zero and the view overflowed | Compiled; positive editor width and a tiny-terminal fallback passed both probes |
| Migrate the legacy app to v2 | Compiled; basic behavior passed, but the one-cell view overflowed | Compiled; reserved label/prompt/cursor space and a tiny-terminal fallback passed both probes |

The extra probes were written after generation and run against copies of the saved output. They check persistent focus, retained initial commands, editable q, ctrl+c quit, a positive editor width, v2 cursor availability/full-screen state, and final cell bounds at 1×1. These supplement the predeclared case criteria; they are not a statistically controlled benchmark. The legacy baseline met the original nonnegative-width wording but zero has unbounded-width semantics for its editor. Probe source and results remain under `.cache/review`.

Manual inspection also checked version-family preservation, command propagation, and request-ID rejection of stale results. All six generation logs contained only model responses and the CLI's shortened-skill-description warning; no tool calls occurred. Existing globally installed skill descriptions could still be visible, so these are comparisons without versus with supplied Charm bodies, not proof of a completely empty baseline context. The with-skill runner supplies selected reference contents directly; it does not measure progressive file navigation.

Both versions of the generated loading app keep the spinner ticking while idle. The bundled Bubble Tea example stops idle ticks and changes the spinner owner on restart, but this instruction did not reliably transfer to generated code. Bubbles 2.2.1 enables virtual cursors by default: the absence of an explicit setter in generated code is valid and was checked, not counted as an improvement or failure. These small trials support specific observations, not a claim that skills make every model output correct.

## Description routing and real terminal checks

All **20** metadata-routing cases passed: 14 positive/boundary cases and six negatives (including unrelated Go work and ordinary meanings of bubble tea, harmonica and charm bracelet). This tests selection from supplied descriptions. Native Codex discovery and Huh file navigation are separate evidence. See the [evaluation protocol](../evaluations/README.md).

The PTY helper compiled and started the actual Bubble Tea and Ultraviolet examples on fresh Linux PTYs at 60×12, resized to 2×2 and 80×20, entered input, and quit. Bubble Tea also loaded a query and accepted q as editor text before ctrl+c. Both exited zero, restored terminal attributes, and emitted alternate-screen exit. These checks do not emulate a final screen or certify screen-reader support. Transcripts and results remain under `.cache/pty`.

## Known dependency limits and unexecuted checks

Wish 2.0.3's SSH 0.4.3 dependency has a reproduced race between early window-change writes and startup `Session.Pty()` reads. The opt-in reproducer exited nonzero with race-detector evidence. Ordinary SSH tests wait for the first rendered view before resizing; they prove established-session behavior only. The source explanation and reproduction command are in the [installed Wish reference](../skills/charmbracelet-wish/references/patterns.md#known-dependency-race-in-the-verified-snapshot). No upstream patch or issue was published as part of this work.

Huh's accessible-mode result/error limitations and Ultraviolet's Buffer/Screen mismatch are documented beside their pinned sources. Huh post-validation cannot recover discarded I/O or cancellation errors for arbitrary forms. A passing example test does not remove these dependency limitations.

At the end of the initial local-delivery phase, hosted CI, native macOS/Windows behavior, public Git-URL installation, and Skills directory listing had not run. The publication phase below records subsequent evidence. See the [implementation review](review.md), [publication review](publication-review.md), and [maintenance workflow](maintenance.md).

## Public distribution

Publication preparation on **2026-09-07** repeated `npm ci` and `GOTOOLCHAIN=go1.25.12 npm run check`: all 12 skills, 11 Node tests, 12 Go modules, and 228 installation/resource checks passed. npm reported zero known dependency vulnerabilities. The grouped catalog passed the current official skills.sh JSON schema using Python jsonschema, and the repository validator checked complete, unique membership. Both staged and unstaged `git diff --check` passed.

The repository [newbpydev/skills-charmbracelet](https://github.com/newbpydev/skills-charmbracelet) is public with default branch `main`. An anonymous HTTPS fetch matched the published Bubble Tea entrypoint byte-for-byte at [c9995a0](https://github.com/newbpydev/skills-charmbracelet/commit/c9995a05cd1cbf60910b17cbd1d62af6d7463f43). Running `check:install` against that exact GitHub revision passed all 228 checks, including nine targets in two modes and twelve standalone selections. Automated checks disabled telemetry.

One ordinary CLI installation from `newbpydev/skills-charmbracelet`, outside CI with existing telemetry preferences unchanged, installed all twelve skills and preserved all 135 distributed files. It used a disposable Codex project and was removed after verification; no real global directory was changed.

The [skills.sh collection page](https://skills.sh/newbpydev/skills-charmbracelet) and all twelve individual skill pages returned their matching names and source content on 2026-09-07 (00:19 UTC on September 8). The collection showed twelve skills. Keyword search and configured display groups were still absent at 00:24 UTC. Page visibility does not establish search inclusion or a leaderboard rank. The Bubble Tea page displayed three passing external audit labels; those labels were observed, not independently validated as part of this release.

Both `skills find charmbracelet --owner newbpydev` and the exact `charmbracelet-bubbletea` query returned no results. Similar page-versus-search behavior is reported in the still-open upstream [indexing issue #1402](https://github.com/vercel-labs/skills/issues/1402). That report corroborates the distinction, not the cause or resolution time for this collection. Installation and direct skill pages work; keyword discovery remains an external outstanding condition. No repeated installs or synthetic telemetry were used to increase counts.

The [initial hosted run](https://github.com/newbpydev/skills-charmbracelet/actions/runs/34172758009) passed all Go modules and Node tests on Linux, macOS, and Windows. Linux also passed race and PTY checks. The installer check failed on macOS and Windows because its temporary-root path was not canonicalized before comparison with resolved installed paths. A Linux symlink-alias probe reproduced the false escape; resolving the temporary root fixed it, and all 228 installation checks then passed through the alias.

The [corrected hosted run](https://github.com/newbpydev/skills-charmbracelet/actions/runs/34173022026) passed the complete Linux, macOS, and Windows matrix on [ea7052b](https://github.com/newbpydev/skills-charmbracelet/commit/ea7052b9717dbeeb56459b17a7cbb1095a714266), including Linux race and PTY steps. All skill contents are unchanged from the public-source installation revision. A separate Markdown-target check resolved all 155 local links across the public tree. The final release gate uses the exact commit and run linked from the GitHub release, rather than treating this earlier run as evidence for future changes.
