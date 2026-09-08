# Contribute

Start with a concrete request or failure an agent should handle better. Reproduce it against the application's exact module versions before changing guidance. Add the smallest instruction, reference, or runnable example that addresses it.

Authoring follows [Anthropic’s Agent Skills best practices](https://platform.claude.com/docs/en/agents-and-tools/agent-skills/best-practices) and the [portable Agent Skills specification](https://agentskills.io/specification): concise entrypoints, descriptive activation metadata, progressive references, runnable resources when useful, and evaluation against concrete tasks.

## Content

- Keep automatic descriptions specific enough to select the right skill. Add positive and negative discovery prompts for new responsibilities.
- Prefer primary release source and official examples. Cite exact revisions and record them in the source manifest.
- Keep every skill independently installable. References link directly from SKILL.md; shared repository maintenance files are not runtime dependencies.
- Explain consequential version differences and integration boundaries. Avoid generic Go tutorials, exhaustive copied API catalogs, and unsupported universal rules.
- Add scripts only for useful deterministic work. Document their inputs, outputs, dependencies, and failure behavior.
- Keep original examples small and executable. Preserve upstream license text and attribution within the installed skill if a contribution incorporates substantial upstream code.

## Checks

From this checkout, install the working content into a disposable project with `npx skills add /path/to/skills-charmbracelet`. Use `.` as the source when running inside the checkout. Public users should use the GitHub source in the README.

Run `npm ci` and `npm run check`. Run `npm run check:upstream` when changing references, source pins, or dependency baselines. Use `npm run check:examples -- --race` for changes involving async state or sessions.

Every example is an independent Go module. Run dependency changes and tests from its own directory. Do not substitute the example module's pins for an unrelated application's dependency policy.

Describe the concrete regression and checks in the change description. Update docs/validation.md with the actual model, harness, toolchain, and result. Do not label unavailable provider, platform, or hosted checks as passed.
