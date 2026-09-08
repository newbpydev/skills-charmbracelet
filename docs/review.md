# Local implementation review

Reviewed sequentially in the main agent session on 2026-09-07. This was source inspection and executable verification, not independent multi-reviewer approval. The review covered all skill entrypoints, conditional references, original examples, validation/installation tools, evaluation boundaries, and CI configuration.

| Finding | Resolution or explicit boundary |
| --- | --- |
| Bubbles legacy tags were described as exclusively v0.x | Verified honorary v1.0.0 source and corrected all 12 local version guides; retained the older characterization fixture intentionally. |
| An explicit virtual-cursor setter could be mistaken for a required fix | Verified New() enables it in Bubbles 2.2.1; corrected guidance and tested baseline cursor availability. |
| Huh README root-model claims do not match its released compatibility interface | Compiled a parent adapter and documented the exact return types. |
| Huh accessible values/errors differ from interactive completion | Bound stable values, tested Ada and EOF, checked required postconditions, and documented the limits of that workaround. |
| Plain Ultraviolet Buffer does not implement the pinned Screen interface | Used ScreenBuffer in offscreen tests and linked the matching declarations. |
| A completed spinner run can leave a tick in flight when restarted | Assigned a fresh spinner ID on restart and added a stale-owner regression test. |
| Lip Gloss example subtracted decoration twice | Applied outer Width once and asserted the final cell width. |
| Reduced-motion retargeting and runtime cancellation needed stronger coverage | Added assertions for retargeting without animation and actual program context cancellation. |
| Tool process failure/timeout could be misreported | Preserved diagnostics and nonzero status, added forced termination fallback and regression tests. |
| CLI success did not imply every requested harness path existed | Checked actual installed bytes, documented the existing-directory condition, tested default and copy modes plus standalone installs. |
| SSH startup resize races inside the pinned dependency | Kept an opt-in failing reproducer and limited ordinary integration claims to established sessions. |
| Generated code still ran idle spinner ticks | Recorded the failure to transfer that guidance; made no universal effectiveness claim. |

All actionable local findings were resolved and affected checks rerun. The upstream SSH race and Huh accessible error behavior remain explicit dependency limitations. The validation report separates local packaging, code generation, native discovery, and terminal results. Publication and hosted evidence remain outside this local delivery.
