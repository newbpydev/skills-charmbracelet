# Public distribution review

Reviewed on 2026-09-07 using the Compound Engineering document-review lenses against the [publication plan](plans/2026-09-07-public-distribution.md). Classification: plan. No separate upstream requirements document was supplied; the user's publication request and approval of `newbpydev/skills-charmbracelet` are the decision authority.

## Findings resolved

| Finding | Severity | Resolution |
| --- | --- | --- |
| The existing installer check always uses the local checkout | P1 | Add an optional source argument and compare actual GitHub-installed resources with the checked-out release contents. |
| Evidence edits could leave the release ahead of its successful CI run | P1 | Commit evidence first, then require matching local/remote heads and green jobs on that exact commit before tagging. |
| Existing users have no documented path to receive corrections | P2 | Add a named-skill update command beside the public installation instructions. |

Three fixes applied to the plan under the user's approval to proceed with the recommended approach and publish. The first is a feasibility omission (confidence 100); the second is an adversarial sequencing omission (75); the third is a design-flow omission (75). Each has a concrete resolution. No catalog redesign, additional hosting, or npm publication is required.

## Coverage

Coherence, feasibility, product, design, security, scope, and adversarial lenses were applied sequentially in the main review context, following the user's AGENTS.md tool mapping. These are seven lenses, not seven independent reviewers. No confidence promotion from self-agreement was used. Coherence, product, security, and scope raised no additional actionable findings.

The optional cross-model product, security, adversarial, and whole-document passes requested Claude Opus 5 at high reasoning through Claude CLI. All four returned no review artifact: the installed CLI rejected the runner's required `--safe-mode` option. No served-model identity or independent corroboration was obtained. The runner's isolation restrictions were retained. This is a capability failure, not current evidence about account authentication.

## Verdict and limits

**Ready to execute the publication plan.** This verdict is not publication evidence; the [validation record](validation.md) records executed checks and external outcomes.

The public artifact inspection found no private-key/token patterns or personal machine paths in the initial 165 tracked files; caches and dependencies were ignored. That bounded check is not a comprehensive security audit. Known upstream Wish and Huh limitations remain documented. Native agent behavior is demonstrated only for Codex; nine-target resource installation is a separate claim. skills.sh indexing and automated audits are external service outcomes, verified after publication.
