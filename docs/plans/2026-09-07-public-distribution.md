---
title: Publish the Charmbracelet skill collection
type: feat
status: published
date: 2026-09-07
---

# Publish the Charmbracelet skill collection

## Purpose and success

Go developers must be able to discover this community collection on skills.sh and install all twelve skills, or one standalone skill, from a public GitHub repository. The local collection already contains version-aware guidance, pinned official references, independent examples, licenses, and executable checks. This work makes that existing content publicly usable and records the evidence for its distribution.

Publication, successful installation, hosted checks, and skills.sh indexing are separate outcomes. Completion requires an anonymous public source, successful installation from that source, a green Linux/macOS/Windows CI matrix, and observed skills.sh pages containing the actual collection and skill identities. Search visibility is checked separately; installation counts and ranking are not a release promise.

## Key technical decisions

- Publish the existing collection publicly, rather than retaining a local-only deliverable. (session-settled: user-directed; rejected alternative: local-only delivery)
- Use `newbpydev/skills-charmbracelet` with `main` as its default branch. The user approved creating the repository, pushing, and publishing. (session-settled: user-approved; rejected alternative: another owner or repository name)
- Preserve the twelve stable, standalone skill names and their canonical `skills/` directories. (session-settled: prior implementation approved; rejected alternative: redesigning the catalog during publication)
- Distribute through the Skills CLI's GitHub support. Keep the development package private on npm. The official skills.sh FAQ describes automatic listing from public-source installations; a separate npm package, hosted service, and unlisted pack are unnecessary for this release.
- Add root `skills.sh.json` with four groups: Core TUI, Forms and content, Runtime, and Maintenance. This changes the repository page display; it is not an indexing manifest or an installer prerequisite.

## U1 — Prepare the public artifact

Update the README to lead with `npx skills add newbpydev/skills-charmbracelet`, followed by standalone and explicit multi-agent examples. Retain local development instructions in the contribution guide. Recommend `--copy` for installations into fresh projects with several harnesses, based on the observed CLI 1.5.24 symlink behavior. Describe nine tested installation targets and Codex-only native behavior evidence accurately.

Include a short first-use prompt and the CLI's named-skill update command. Users must have a documented route to receive the corrections described in the maintenance policy.

Provide repository description, appropriate GitHub topics, MIT licensing, a link to the skills.sh collection page, and a clear community/independent-of-Charm statement. Preserve all pinned references and upstream attribution. Group every actual skill exactly once in `skills.sh.json`; validate the file against the official schema and check its names against the catalog.

Inspect the staged publication contents before the first commit. Exclude caches, node_modules, local credentials, generated model transcripts, and personal machine paths. Keep the existing evaluation report explicit that private scratch probes are local observations, not a publicly reproducible benchmark. Do not copy those transcripts into the release.

## U2 — Verify and publish

Run `npm ci`, `npm run check`, and `git diff --check` on the release contents. The check covers skill/resource integrity, validator/process tests, twelve Go modules, and isolated installation checks. Run changed behavior's additional checks when warranted; unchanged authenticated model evaluations need not be repeated.

Create the public GitHub repository after the artifact is reviewable, commit the reviewed contents, push `main`, and observe `.github/workflows/check.yml`. Keep Actions permissions at `contents: read`; no publishing credentials, package registry tokens, or `pull_request_target` workflow are needed. If hosted jobs reveal platform problems, fix them and observe the corrected commit's complete matrix before creating version `v0.1.0` and its GitHub release.

Verify the repository is anonymously readable rather than relying only on authenticated `gh` access. Install the public source into disposable project directories with the pinned Skills CLI. Compare installed resources byte-for-byte with the published commit, test all nine targets and standalone selection, and ensure checks cannot write to the user's real global skill directories. Repeated automated checks disable telemetry.

Extend `scripts/check-install.mjs` with one optional source argument, defaulting to this checkout. Reuse its existing resource comparison and temporary-project isolation. Run it against a GitHub URL pinned to the commit being verified, from a checkout with matching skill contents; print the source in its output.

## U3 — Verify skills.sh discovery

Run one ordinary, genuine installation of all twelve skills from the public repository outside CI, into a disposable project. Respect existing `DO_NOT_TRACK` and `DISABLE_TELEMETRY` preferences. Do not loop installs to increase counts or make direct synthetic telemetry requests. With telemetry enabled, the official documented discovery path is the CLI's public-source installation event.

Read the public skills.sh repository page and all twelve skill pages. Require matching repository/skill content, not merely HTTP 200, because a generic page is not indexing evidence. Check `npx skills find charmbracelet` or the public website's search for discoverability. Record the observed URLs and date, and distinguish repository-page visibility, individual-skill indexing, search visibility, and later automated audit results. An audit that is still pending is not a clean security assessment.

Pages and search can be cached. Poll reads at reasonable intervals without repeating installs. If the service has not indexed the collection during the publication session, keep indexing explicitly pending in the evidence record and report the public installation URL and the remaining external condition. Do not call the overall visibility objective complete while claiming unobserved indexing.

## U4 — Release and maintain

Record the public repository, reviewed commit, hosted run, remote-install results, skills.sh observations, and known limitations in `docs/validation.md`. Add a publication review record and work-log entry without rewriting the previous local implementation review. Publish `v0.1.0` release notes describing the twelve skills, reference snapshot, supported installation paths, and evidence limits.

Commit and push the evidence updates before the final release gate. Immediately before tagging, require clean local contents, local HEAD equal to remote `main`, and successful required jobs on that exact commit. Put that final commit and run URL in GitHub release notes, avoiding a self-referential commit hash in the source file. Earlier verification observations retain their own recorded revision and scope.

A documentation-only follow-up after publication may record the release's actual URL and completion status. It does not move the verified tag or change the released skill contents.

Keep collection versions distinct from library versions. Keep release tags immutable; correct published mistakes with a new commit and content release. For a harmful change, revert the change on `main`, validate, and publish the correction. Review upstream changes monthly and when a relevant release or issue is noticed; use the existing source checker as a report, not an automatic upgrade.

Known upstream limitations remain visible: Wish's pinned SSH dependency has a reproduced early-resize race, Huh's accessible mode discards some errors, and installation checks for nine harnesses do not prove native model behavior for nine harnesses. These limit downstream claims but do not prevent publishing useful, accurately qualified guidance.

## Verification checklist

- [x] Public README, licenses, and twelve catalog group entries agree; GitHub metadata is prepared.
- [x] Reviewed publication tree contains no private scratch artifacts.
- [x] Required local checks pass.
- [x] Public repository is anonymously readable and uses `main`.
- [x] Hosted Linux, macOS, and Windows jobs pass on the release candidate.
- [x] Public-source installs preserve all resources across nine targets and standalone selections.
- [x] One genuine public install follows existing telemetry preferences.
- [x] Collection and twelve skill pages show matching content on skills.sh.
- [x] Search visibility is separately recorded as pending.
- [x] Release `v0.1.0` and its evidence links are published.

## Observed outcome

[v0.1.0](https://github.com/newbpydev/skills-charmbracelet/releases/tag/v0.1.0) is public. Installation and resource preservation passed, the exact release commit passed all three hosted platforms, and twelve skills.sh pages display the collection's instructions. See [validation evidence](../validation.md). Keyword search inclusion and the configured display groups remain external follow-up items; they are not claimed complete.

## Sources checked on 2026-09-07

- [Skills CLI source and usage](https://github.com/vercel-labs/skills)
- [skills.sh FAQ: hosting, listing, and packs](https://skills.sh/docs/faq)
- [skills.sh repository-page customization](https://skills.sh/docs/customize)
- [skills.sh customization schema](https://skills.sh/schemas/skills.sh.schema.json)
- [Agent Skills specification](https://agentskills.io/specification)
- [Anthropic skill-authoring best practices](https://platform.claude.com/docs/en/agents-and-tools/agent-skills/best-practices)
