# Charmbracelet skills for Go developers

[![Check skills](https://github.com/newbpydev/skills-charmbracelet/actions/workflows/check.yml/badge.svg)](https://github.com/newbpydev/skills-charmbracelet/actions/workflows/check.yml)

Version-aware Agent Skills for building and maintaining Go terminal applications with the Charm ecosystem. Install the collection or select individual skills through the [Skills CLI](https://github.com/vercel-labs/skills).

This is a community-maintained project, independent of Charm. It combines concise guidance, pinned official references, small runnable examples, and behavioral checks. Existing legacy applications stay on their dependency family unless you request migration.

Current content release: [v0.1.0](https://github.com/newbpydev/skills-charmbracelet/releases/tag/v0.1.0).

## Install

Run this from the Go project where you want to use the skills, then select skills and agents in the installer:

```sh
npx skills add newbpydev/skills-charmbracelet
```

Select one skill or specific harnesses:

```sh
npx skills add newbpydev/skills-charmbracelet --skill charmbracelet-bubbletea
npx skills add newbpydev/skills-charmbracelet --skill '*' --agent codex claude-code --copy
npx skills add newbpydev/skills-charmbracelet --skill '*' --agent cursor --copy
```

Add `--global` for a user-global installation. `--list` lists discoverable skills without installing them. `--copy` is available where symlinks are unsuitable. Browse the [skills.sh catalog](https://skills.sh/newbpydev/skills-charmbracelet) or [GitHub source](https://github.com/newbpydev/skills-charmbracelet).

For an explicit multi-harness install into a new project, use `--copy`. In CLI 1.5.24, default project symlinks for non-universal harnesses such as Windsurf and Crush require their `.windsurf` or `.crush` directory to exist first. The CLI can otherwise report overall success while skipping those links. The installation check verifies actual files and resources in both modes.

Node.js is required for the Skills CLI; the tested CLI version is **1.5.24**, requiring Node **22.20.0 or newer**. The installed skills use your agent's file tools and ordinary Go tooling. Their content follows the [Agent Skills specification](https://agentskills.io/specification), with no mandatory MCP server or harness-specific plugin.

Installation and resource preservation are checked for Claude Code, Codex, Cursor, Gemini CLI, OpenCode, GitHub Copilot, Windsurf, Cline, and Crush. Native discovery and file navigation have been demonstrated in Codex; see [validation evidence](docs/validation.md) for the remaining coverage.

Start a new agent session in the project and ask, for example: “Use charmbracelet-bubbletea to add a searchable list. Inspect go.mod first and preserve this project's dependency family.” For a deliberate upgrade, ask for `charmbracelet-migration`. If you are unsure which libraries fit, start with `charmbracelet-ecosystem`.

To update an installed skill, use the CLI's [update command](https://github.com/vercel-labs/skills#skills-update):

```sh
npx skills update charmbracelet-bubbletea
```

Skill updates refresh agent guidance; changes to your application's Go dependencies remain a separate decision.

## Choose a skill

| Skill | Use it for |
| --- | --- |
| [charmbracelet-ecosystem](skills/charmbracelet-ecosystem/SKILL.md) | Library selection and integrations spanning the ecosystem |
| [charmbracelet-bubbletea](skills/charmbracelet-bubbletea/SKILL.md) | Models, commands, async work, input, views, and lifecycle |
| [charmbracelet-bubbles](skills/charmbracelet-bubbles/SKILL.md) | Interactive components, focus, sizing, key maps, and commands |
| [charmbracelet-lipgloss](skills/charmbracelet-lipgloss/SKILL.md) | Styling, layout, color, Unicode widths, tables, and trees |
| [charmbracelet-huh](skills/charmbracelet-huh/SKILL.md) | Standalone/embedded forms and accessible prompts |
| [charmbracelet-glamour](skills/charmbracelet-glamour/SKILL.md) | Markdown rendering and viewport reflow |
| [charmbracelet-wish](skills/charmbracelet-wish/SKILL.md) | SSH, authentication, PTYs, sessions, and shutdown |
| [charmbracelet-log](skills/charmbracelet-log/SKILL.md) | Structured logging, writers, and slog integration |
| [charmbracelet-harmonica](skills/charmbracelet-harmonica/SKILL.md) | Spring animation, retargeting, and tick lifetime |
| [charmbracelet-terminal](skills/charmbracelet-terminal/SKILL.md) | Ultraviolet, ANSI, capabilities, and terminal primitives |
| [charmbracelet-migration](skills/charmbracelet-migration/SKILL.md) | Deliberate coordinated legacy-to-v2 upgrades |
| [charmbracelet-testing](skills/charmbracelet-testing/SKILL.md) | Model tests, teatest, headless runtime, and terminal checks |

Each skill works alone. References and examples stay inside its installed directory. A task might use several skills, but none requires an absent sibling skill.

## What the collection adds

- Explicit distinctions between legacy GitHub imports and released `charm.land/*/v2` APIs, including supporting libraries that retain GitHub paths.
- Guidance on component command propagation, persistent focus, real versus virtual cursors, stale async results, and terminal ownership.
- Corrections for demonstrated documentation/API mismatches, including Huh's compatibility model and accessible result handling, and Ultraviolet's offscreen screen interface.
- Curated links to released official examples, with independent local modules that do not rely on upstream `replace ../` directives.
- Source provenance and repeatable checks. Read [validation evidence](docs/validation.md) for the distinction between installation, native discovery, generated code, and terminal behavior.

The collection emphasizes Bubble Tea, Bubbles, Lip Gloss, Huh, Glamour, Wish, Log, and Harmonica. Ultraviolet and selected supporting modules receive focused advanced guidance. It is not an exhaustive manual for every Charm repository or an end-user guide to Gum, Glow, VHS, Soft Serve, or Crush.

## Verified reference baselines

The 2026-09-07 snapshot uses Bubble Tea 2.0.9, Bubbles 2.2.1, Lip Gloss 2.0.6, Huh 2.0.3, Glamour 2.0.1, Wish 2.0.3, Log 2.0.1, and Harmonica 0.2.0. Legacy migration fixtures use Bubble Tea 1.3.10, Bubbles 0.21.1, and Lip Gloss 1.1.0. The honorary Bubbles 1.0.0 release also belongs to the legacy family; it is documented without upgrading that fixture. Experimental modules use exact pseudo-versions.

These are tested references, not a rule to upgrade every application. [The source manifest](upstream-sources.json) records module versions, commits, Go declarations, licenses, and consumers. Examples declare Go **1.25.12**; inspect the target project's own toolchain requirements when adapting them.

## Develop and verify

```sh
npm ci
npm run check
npm run check:upstream
```

`check` validates skills, runs validator tests, tests/vets/formats all independent example modules without rewriting dependencies, and exercises isolated Skills CLI installations. `check:upstream` is a separate network report; newer versions require review rather than automatic content changes. `npm run check:examples -- --race` adds Go race detection where supported.

`npm run check:pty` exercises actual terminals on Linux using Python 3. `npm run check:native` and `npm run check:routing` are optional authenticated Codex evaluations; they use the CLI's existing session and may consume model usage. They are excluded from the default check and hosted CI.

The [CI workflow](https://github.com/newbpydev/skills-charmbracelet/actions/workflows/check.yml) checks Linux, macOS, and Windows. The [validation record](docs/validation.md) identifies observed runs and their scope. The [contribution guide](CONTRIBUTING.md), [maintenance guide](docs/maintenance.md), and [evaluation guide](evaluations/README.md) describe how to update evidence.

## License

Original skill instructions and examples are [MIT licensed](LICENSE). Each distributed skill includes its license. Official upstream projects retain their authorship and licenses; this repository links to their source and does not redistribute their full manuals or example trees.
