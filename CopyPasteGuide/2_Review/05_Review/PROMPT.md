You are performing an adversarial submission review for a Go Olympus challenge. Execute immediately. Do not explain what you are about to do. Do not ask for clarification. Read all artifacts from the challenge directory, run all review gates, and output the verdict now.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

Read the following files from the challenge folder:
- problem.md
- test.patch
- solution.patch
- Dockerfile
- repo_url.txt (contains repo URL and commit hash)
- Extract test.sh from test.patch (it is added as a new file in the patch)

---

### Review Gates (run all, report incomplete ones)

**1. Alignment and Uniqueness Preflight**
Use authenticated gh CLI, not title-only search. Read README, CONTRIBUTING/AGENTS, docs, examples, design notes, code comments, and full PR/issue bodies, comments, reviews, linked context, and Discussions via gh api/GraphQL. Search semantic variants. Reject immediately for repo-design mismatch, negative maintainer direction, material duplicate, or open/closed/merged PR overlap.

**2. Artifact and Patch Pass**
Verify pinned immutable commit, repo eligibility, canonical checkout, clean patch application, patch boundaries, executable root test.sh, no conflicts, correct Docker shape. No platform branding ("challenge", "olympus", "shipd", "mars") in test file names or test comments — **EXCEPTION: the Dockerfile MUST use `FROM public.ecr.aws/d3j8x8q7/olympus-base-go:latest` and MUST include `ENV CHALLENGE_DOCKER=1` — these are required by the platform and must NOT be flagged.**

**3. Alignment Pass**
Map every problem.md behavior to new tests and solution hunks. Classify every assertion as problem-stated or repo-derivable. Anything unsupported or unmapped is a finding.

**4. Shortcut and Coverage Pass**
Inspect agent trajectories/patches if available. Try plausible wrong implementations. Check every claimed topology, dialect/backend, input form, operation, lifecycle state, variant, resource kind, boundary, restart path. Every T3/T4 finding must state: "A solution that <wrong behavior> would still pass because <missing assertion>."

**5. FP-Readiness Audit**
Derive adversarial probes from problem.md and the pinned repo. Vary one public behavior dimension at a time. Run probes against base, reference, and any available passing-agent patches. Add a new test for every candidate-only failure. Map every meaningful reference hunk, branch, fallback, error path, state transition to behavioral coverage; revert each hunk and require new to fail, or remove the hunk. Stop when all probes are: prompt-stated, base-failing, reference-passing, and directly tested.

**6. Solution Pass**
Inspect every hunk for bugs, regressions, dead/duplicate code, public-surface creep, broad error handling, collateral edits, style/type/import failures, stale hunks, visible-test workarounds. Ask: "If I remove this hunk, does new still pass?" (if yes, the hunk is unverified — remove it or add a test for it)

**7. Race and Regression Pass**
Inspect: polling, timers, goroutines, filesystem moves, retries, caches, subprocesses, partial failures, concurrent reads/writes. Audit adjacent readers after persistence/type/serializer changes.

**8. Difficulty and Matrix Pass**
Measure meaningful LOC (excluding blanks, comments, no-ops, tests, generated files, imports, declarations, braces, punctuation-only lines, boilerplate, mechanical propagation).
Verify base commit + test.patch → base PASS, new FAIL.
Verify both patches → base PASS, new PASS (in Docker offline).
Check: >= 5 files touched, >= 2 subsystems, 300-500 meaningful implementation LOC (flag if below 250 or above 600).
If a legitimate passing agent completed in substantially fewer LOC than the reference, treat it as compression evidence.

**9. Final Adversarial Pass**
Search for: missing edge test, removable hunk, race, implementation leak, collateral edit, shortcut that still passes.

---

### Official Criteria

Problem (P1-P7):
- P1: Aligns with repo's philosophy
- P2: Not fixed in any open/merged/closed PR
- P3: Self-contained, solvable from repo + description alone
- P4: Clear, concise, unambiguous
- P5: Verifiable — objectively testable
- P6: Not prescriptive — no solution leaks
- P7: Materially distinct — not a duplicate

Tests (T1-T8):
- T1: 100% fail on base, 100% pass with solution
- T2: Deterministic — no timing, randomness, ordering
- T3: Strong — reject inaccurate solutions
- T4: Extensive — covers all behaviors and obvious edge cases
- T5: Only stated or repo-discoverable behavior
- T6: Offline — no network dependency
- T7: No over-pinning exact output/error text
- T8: Failure diagnostics intact — no catch-all messages hiding real failures

Solution (S1-S4):
- S1: Meets all requirements
- S2: No regressions, follows repo patterns
- S3: No irrelevant changes
- S4: No AI slop

---

### Output

Respond in exactly this format — no other text before or after:

Verdict: <ACCEPTED or REVISION REQUESTED>

Problem Description: <1/3, 2/3, or 3/3> <Weak/Minor/Good>
Findings:
1. [P#] <specific finding with evidence and exact repair instruction>
(write "None." if 3/3)

Tests: <1/3, 2/3, or 3/3> <Weak/Minor/Good>
Findings:
2. [T#] <wrong solution that would still pass> — <missing assertion> — <exact repair>
(write "None." if 3/3)

Solution & Code: <1/3, 2/3, or 3/3> <Weak/Minor/Good>
Findings:
3. [S#] <path/hunk, impact, exact repair>
(write "None." if 3/3)

Other Notes:
- Gates completed: <list>
- LOC count: <count> meaningful LOC
- File spread: <count> files across <count> subsystems
- Shortcuts found: <list or none>
- Residual risk: <risk>


Rules:
- Use REVISION REQUESTED if ANY category is below 3/3, any material shortcut passes, any problem/test mismatch, or artifact hygiene is invalid.
- Do not claim confidence from environment timeouts or missing XML.
- Findings ordered by: reproduced failures, repeat-run flakes, concrete regressions, coverage/spec gaps, wording concerns.
