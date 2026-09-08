# Working on this collection

Canonical distributable content lives in `skills/charmbracelet-*/`. Every skill must work when installed alone. Keep essential guidance local; do not require sibling skills or repository-root files at runtime.

Use released source, matching examples, and compilation to verify API claims. Record sources in `upstream-sources.json`; put useful citations beside the relevant guidance. Preserve the user's project dependency family unless migration is requested. Experimental modules need exact versions.

Keep SKILL.md concise, with discriminating descriptions and direct links to conditional references. Add scripts only for useful deterministic work. Do not copy entire upstream documentation trees. Keep upstream code attribution with distributed assets.

Run `npm ci` and `npm run check`. `npm run check:upstream` is a separate network check. Go modules in assets are independent; run commands from the containing module and preserve their go.sum files. Do not run a repository-wide `go mod tidy`.

Record actual validation in `docs/validation.md`. Installation, native discovery, generated-code evaluation, and terminal behavior are different evidence. Do not claim a platform or model passed a check that was not run.

Keep evaluations sequential. Do not publish the repository or install into real user-global skill directories as part of development checks.
