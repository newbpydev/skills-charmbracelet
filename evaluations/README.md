# Behavioral evaluations

Define tasks before adding guidance. Compare the same prompt, model, and execution limits with and without selected skill content. Keep each run in a fresh process and a temporary workspace. Save model/version, prompt, supplied references, output, compiler/test results, and duration.

The three baseline cases cover a new v2 application, legacy maintenance, and migration. The bundled example tests separately exercise library boundaries. Compilation is necessary but not sufficient: score the observable behaviors in cases.json using tests and inspection of produced artifacts.

`scripts/evaluate.mjs` runs one case using the installed Codex CLI in noninteractive, read-only mode and records output under `.cache/evaluations/`. The prompt prohibits tool calls; inspect the event log to verify that boundary. This measures explicit-context guidance, not automatic skill discovery or the ability to navigate references. No API keys are read or written by the runner. Authentication uses the CLI's existing session. It ignores user configuration and uses the CLI's default model, a fresh temporary workspace, a timeout, and no persisted conversation. Installed skill descriptions may still be present; their bodies are not accessed without tools.

Run baseline and skill variants sequentially. Use `--check-only` to assess previously generated Go output. Record results in docs/validation.md; do not count an unavailable provider as a passed evaluation.

Native harness discovery and installation are separate checks. Unrelated Go requests should not activate Charm skills. An upgrade request should select migration; a request to preserve v1 must not silently upgrade dependencies.

## Commands and boundaries

```sh
node scripts/evaluate.mjs new-v2 baseline
node scripts/evaluate.mjs new-v2 skill
node scripts/evaluate.mjs new-v2 skill --check-only
npm run check:routing
npm run check:native
npm run check:pty
```

Use `legacy-fix` and `migrate-v2` for the other generation cases. Generation runs create a temporary workspace and save prompts, exact supplied content, JSON responses, generated Go, CLI execution metadata, and compiler output in `.cache/evaluations`. Preserve the prior directory before rerunning a case if comparing iterations. Check-only rechecks existing generated code without invoking a model. These are opt-in authenticated CLI calls and can consume model usage.

Routing evaluates only the name/description catalog against 14 positive or boundary requests and six negative requests in [routing.json](routing.json). Its exact-choice grader fails on missing, duplicate, incorrect, or unrelated selections. This is a metadata discrimination check, not evidence every harness automatically activates the skill.

Native checks install into a temporary Codex project, verify all available names, then ask Codex to read Huh's local references. Discovery is machine-checked; navigation output and tool use require human review. The evidence lives in `.cache/native`; the scratch project is removed afterward. Other global skill descriptions may remain visible to the CLI. A read-only sandbox is a tool restriction, not a claim of an empty system prompt or isolated credentials.

PTY checks compile the two original terminal examples, operate fresh Linux PTYs, test input/resize/exit, and compare terminal attributes and alternate-screen transitions. Binary terminal transcripts and JSON results are saved in `.cache/pty`. These are lifecycle checks, not screenshots, screen-reader audits, or Windows Console tests. `.cache` artifacts are local evidence and intentionally excluded from distribution; summarize reviewed outcomes in [validation evidence](../docs/validation.md).
