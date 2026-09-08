# Maintain the references

## Source records

`upstream-sources.json` contains one record per documented module baseline: repository, module path, version, commit, minimum Go declaration, license path, source paths, and consumer references. The installed source references explicitly identify experimental modules. Installed skills carry their relevant citations locally; the root manifest is maintainer tooling.

Run `npm run check:upstream` to query the Go proxy and verify pinned source paths. The command reports changes without editing files. A network error is an unavailable check, not proof of a deleted source. A newer module release is a review signal, not evidence the current guidance is broken.

## Review an update

1. Read the release notes, module declaration, API changes, and affected official examples at the candidate revision.
2. Check typed boundaries with companions. Confirm the Go toolchain declaration and resolved dependencies.
3. Update only affected skill references, examples, and source records. Preserve legacy guidance when that family remains part of the collection's supported workflows.
4. Compile and test examples, including meaningful behavior and installation checks. For experimental modules, verify the exact pseudo-version against the selected released core.
5. Run relevant evaluations and record results. Keep improvements proportional to observed failures.

Use source code and executable tests to resolve documentation contradictions. Record a concrete discrepancy with its version so later maintainers can remove the caveat if upstream changes. Do not perpetuate a workaround as a universal rule.

## Releasing the collection

Keep skill names stable because users install them by name. Use collection versioning for the content release and exact source versions for the library baselines; these are different version numbers.

The public source is [newbpydev/skills-charmbracelet](https://github.com/newbpydev/skills-charmbracelet). Users install through the Skills CLI; `package.json` stays private because it describes development tooling, not an npm-distributed skill package.

1. Run `npm ci` and `npm run check`. Inspect the staged publication tree and attribution; never publish `.cache`, credentials, or raw evaluation transcripts.
2. Push the candidate to `main` and observe every Linux/macOS/Windows job. Fix failures before releasing.
3. From a checkout with matching skill contents, verify the exact public revision:

   ```sh
   npm run check:install -- https://github.com/newbpydev/skills-charmbracelet/tree/COMMIT_SHA
   ```

   Replace `COMMIT_SHA` with the published revision. The script compares every installed resource with this checkout, across nine targets and standalone selections. It uses disposable projects, suppresses telemetry, and never installs globally.
4. Record the observations in [validation evidence](validation.md). Commit and push evidence changes, then require successful CI on that final commit, matching local/remote heads, and a clean tree before tagging.
5. Create an immutable content tag and GitHub release with the exact commit, successful run URL, tested source versions, and relevant limitations. Do not move a published tag. Revert harmful content on `main`, validate, and publish a correction as a new version. Users can update named skills with `npx skills update charmbracelet-bubbletea`.

## skills.sh visibility

The [official FAQ](https://skills.sh/docs/faq) describes automatic listing from public GitHub installations through the CLI's anonymous telemetry. Use one genuine public-source installation to verify the user journey; respect telemetry opt-outs. Automated regression checks keep telemetry disabled. Do not generate repeated installs or synthetic telemetry to influence rankings.

Read the [collection page](https://skills.sh/newbpydev/skills-charmbracelet), individual skill pages, and search results to verify visibility. Match actual names and content, not just HTTP status. Record delayed indexing or audit results as pending; public installation, indexed pages, search ranking, and audit outcomes are different observations.

Root [skills.sh.json](../skills.sh.json) groups the catalog for display, as described in the [customization documentation](https://skills.sh/docs/customize). It does not control installation or register the skills. The collection validator checks complete, unique membership; when changing the configuration shape, validate against the linked official JSON schema too.

Review upstream versions monthly and when relevant releases or issues are noticed. Use the source checker's report to decide what needs inspection. Maintain the published legacy guidance while it remains part of the collection's scope. See the [initial publication plan](plans/2026-09-07-public-distribution.md) and [review](publication-review.md) for release decisions.
