- Scouter GO

  ```markdown
  ## Olympus Scouter -- Go

  Find Go repos worth manual Olympus investigation. Return repos only. Do not find issues, seams, or artifacts; measuring production-code size is allowed for repository eligibility.

  ### Hard Filters

  Reject repos that fail:

  - Public GitHub repo.
  - Go-first.
  - 500+ stars.
  - Active in the last 12 months.
  - Permissive allowed license.
  - Not already in `/Users/ayushcodex/Desktop/Shipd.ai`.
  - Not listed in `/Users/ayushcodex/Desktop/Shipd.ai/Challenge`; read it first.
  - Docker-safe with the approved Go base image: dependencies can be installed at build time (`go mod download`/`go build`), tests run offline with `network none`, and the target does not require CUDA, browsers, Docker-in-Docker, privileged access, systemd, external databases/services, or runtime network.
  - Has a `go test`compatible deterministic test suite or an equally clear offline runner.
  - Has `go.mod`/`go.sum` metadata that supports deterministic dependency installation during the Docker build (vendored modules or a resolvable module graph pinned to specific versions).
  - Has clear package boundaries for the target code; reject monorepo-only layouts with no isolated Go module/package target and one-file designs with no meaningful subsystem spread.
  - Has lightweight, stable upstream tests for the relevant package surfaces; reject repos whose useful baseline requires long-running services, network access, expensive fixtures, build tags requiring external toolchains, or routinely flaky/time-consuming suites.
  - Contains at least `30,000` lines of real Go production code, measured with a code counter while excluding tests (`_test.go`), generated code (`.pb.go`, `_generated.go`, mocks, stringer output, etc.), vendored code (`vendor/`), fixtures, and documentation.
  - Has genuine cross-cutting architecture: multiple production subsystems exchange state or control through real workflows, so a correct feature could require coordinated changes across components rather than one package or file.
  - Has meaningful domain rules, invariants, state transitions, or interactions; reject repos whose apparent size is mostly CRUD/REST endpoints, admin panels, API glue, CLI flag wiring, ETL/IO wiring, configuration, or generated/schema code.
  - Every detected repository license and license declaration is on the allowed list. If any repository or bundled component uses an unknown or disallowed license, reject it even when another license is allowed.

  **Allowed licenses:** MIT, BSD, BSD-1-Clause, BSD-2-Clause, BSD-2-Clause-Flex, BSD-2-Clause-FreeBSD, BSD-2-Clause-Modification, BSD-2-Clause-Patent, BSD-2-Clause-Views, BSD-3-Clause, BSD-3-Clause-Attribution, BSD-3-Clause-EricHeitz, BSD-3-Clause-HealthLevelSeven, BSD-3-Clause-LBNL, BSD-3-Clause-Modification, BSD-3-Clause-OpenMPI, BSD-3-Clause-plus-CMU-Attribution, BSD-3-Clause-plus-Paul-Mackerras-Attribution, BSD-3-Clause-plus-Tommi-Komulainen-Attribution, BSD-4-Clause, BSD-4-Clause-Argonne, BSD-4-Clause-Atmel, BSD-4-Clause-Giffin, BSD-4-Clause-PC-SC-Lite, BSD-4-Clause-Plus-Modification-Notice, BSD-4-Clause-UC, BSD-4-Clause-Visigoth, BSD-4-Clause-Vocal, BSD-4-Clause-Wasabi, BSD-4.3TAHOE, BSD-5-Clause, BSD-FatFs, BSD-Mixed-2-Clause-And-3-Clause, BSD-Protection, BSD-Source-Code, Boost, BSL-1.0, Other, BLAS, GNU-All-permissive-Copying-License, Apache, Apache-2.0, Apache-2.0-Modified, Apache-with-LLVM-Exception, Apache-with-Runtime-Exception, Creative Commons, CC-BY-1.0, CC-BY-2.0, CC-BY-2.5, CC-BY-3.0, CC-BY-4.0.

  ### Prefer

  Prefer repos with product/workflow/runtime/data depth:

  - workflow engines, orchestrators, schedulers, distributed tasks
  - data sync, ingestion, replication, storage lifecycle
  - observability, logs, metrics, event pipelines
  - ML/data pipelines, experiment tracking, feature stores
  - proxies/gateways, service meshes, deployment/runtime managers

  Good signs: `go test`, clean module/package layout, pin-able deps (`go.sum` present, vendored or reproducible), multiple backends/modes, mature tests/docs, and several interacting surfaces such as API/CLI, storage, execution, serialization, recovery, validation, import/export, observability.

  Demote pure tooling/libraries: linters, parsers, formatters, packaging/build tools, auth/admin/RBAC frameworks, pure algorithm libraries, mature SDKs/client wrappers, schema/report-heavy tools, thin CLI wrappers around another system.

  ### Output

  Return 3-5 repos:
  ```

  1. <owner/repo>
     URL: <url>
     Commit: <hash>
     License/stars: <license>, <stars>
     Why: <one sentence>
     Surfaces: <rich surfaces>
     Risk: <main warning or none>
     Checked: <skip/local clear, 30k+ Go LOC, cross-cutting architecture, all licenses allowed, go test/runner, deterministic deps (go.sum/vendor), package layout, fast/stable tests, Docker plausible>

  ```

  ```

- Repo Analyst

  ```markdown
  ## Behavioral Seed Finder -- Builder

  Inside one chosen repo, find repo-native behavioral issue/feature seeds for Builder. Do not implement, probe, or create artifacts. The seed does not need to be Olympus-sized; Builder owns natural expansion and final tier proof.

  Prefer a missing or incorrect user-visible workflow with a natural expansion path across adjacent lifecycle behavior. Do not return API inventories, one more flag/field/backend/plugin, schema or report plumbing, or a fix that can live in one central adapter.

  ### Seed Validity Gate

  Reject a candidate only when one of these is true:

  - It is not a behavioral contract: the proposal is only an API inventory, flag/field addition, schema/report plumbing, wrapper, or configuration propagation without a meaningful observable workflow.
  - An open, closed, or merged PR already implements the same public behavior. Check targeted PR searches for the seed and its natural expansion; direct PR overlap is disqualifying.
  - Repo-local negative discussion, maintainer rejection, or an explicit statement that the behavior is unsupported or belongs elsewhere makes the candidate ineligible.

  Do not reject solely because a neutral issue, roadmap note, file count, LOC estimate, or compression proof is unavailable. Do not implement or probe the issue. Rank candidates by expansion potential: prefer a seed that opens a coherent same-workflow expansion across real subsystems, variants, persistence, execution, validation, serialization, CLI/API, recovery, or observability without inventing unrelated requirements. Builder must prove the final size and compression resistance.

  ### For Each Candidate

  Report:

  - Public behavior and observable edge cases.
  - Evidence paths from code/tests/docs.
  - Seed and natural expansion scope, with the user-visible outcome of each.
  - Likely existing implementation files and affected subsystems.
  - Behavioral tests needed to fail the base commit and kill the shortcut.
  - `problem.md` info needed for fairness.
  - Targeted open/closed/merged PR overlap and negative-discussion result.
  - Expansion potential and affected subsystems; explain why the expansion is the same user workflow rather than padding.

  Rank by behavioral validity, natural expansion, and repo-native depth. Do not promote a candidate whose expansion depends on padding or unrelated behavior.

  ### Output
  ```

  1. <candidate> -- <READY>
     Behavior: <observable user-facing contract>
     Evidence/files: <code, test, and docs paths>
     Expansion: <adjacent behaviors and affected subsystems>
     Tests/problem: <behavioral tests and minimum public information>
     Compression: <shortest shortcut and the assertion that kills it>
     PR overlap: <targeted open/closed/merged PR result>
     Negative discussion: <none, or disqualifying evidence>
     Expansion potential: <same-workflow behaviors and affected subsystems>

  ```

  ```

- Builder

  ````markdown
  ## Go Shipd/Olympus Challenge Builder

  Build and harden a Go Shipd/Olympus challenge from a Scouter seed or Issue Miner candidate. Verify the repo, expand the seed only through natural same-workflow behavior, create the artifacts, and complete the validation gates. A separate critic-only reviewer may perform the final critique.

  The final solution must naturally contain `500+` meaningful implementation LOC across `5+` existing non-test files, cross real subsystems, and plausibly require `60+` agent messages while resisting seed-only or centralized shortcuts. For 10 legitimate agent runs, target at least 1 and at most 4 passes; fewer indicates unfairness or unsolvability, while more than 4 is too easy. Meaningful LOC excludes blank/comment-only/no-op lines, tests, generated files (e.g. `*.pb.go`, `*_generated.go`, mocks, stringer output), imports, declarations, braces, punctuation-only lines, boilerplate, and mechanical padding.

  ### Artifacts

  Create `challenge/<issue-title>/` containing exactly:

  - `problem.md`
  - `test.patch` containing new tests plus root `test.sh`
  - `solution.patch` containing implementation only
  - `Dockerfile`
  - `repo_url.txt` containing the URL and immutable commit hash

  ### Quest Quality Keys

  #### Problem (P1–P7)

  - **P1:** Aligns with the repo's philosophy
  - **P2:** Not already fixed in an open or merged PR (I know, duplicate point. It is really important)
  - **P3:** Self-contained — solvable from the repo and description alone
  - **P4:** Clear, concise, and unambiguous — describes what to build or fix and don't leave points for guessing
  - **P5:** Verifiable — success is objectively testable
  - **P6:** Not prescriptive — don't leak the solution, we want to challenge the agents, remember!!
  - **P7:** Not a duplicate — we run a similarity check against existing submissions. Don't treat it as a gate; open the results, read the close matches it surfaces, and confirm you're not rebuilding a problem that already exists. Rewording the same task or reshaping the same behavior doesn't make it new.

  #### Tests (T1–T7)

  - **T1:** They highlight the missing or incorrect behavior (depending on the task's category — feature request, bug fix, etc). They should 100% fail at the base commit and 100% pass after adding your solution.
  - **T2:** The tests should be deterministic (no timing, randomness, goroutine/channel ordering, or scheduling; nothing that could change across multiple runs or across different machines).
  - **T3:** Strong tests. The tests shouldn't be permissive enough to let inaccurate agent solutions pass.
  - **T4:** (IMPORTANT) Extensive coverage. The tests should cover the requested behavior and all the obvious edge cases.
  - **T5:** DO NOT check for unspecified or undiscoverable behavior — this will be unfair to the agents to expect them to implement something not in the description or discoverable from the repo 😞.
  - **T6:** They don't require any network connection (we run the container with `-network none`).
  - **T7:** Don't over-pin the output — don't assert on exact output (error text, messages, wording, formatting) unless the description says so or it's obvious from the repo's existing patterns. Otherwise it's unfair to an agent that gets the behavior right but words it differently; checking the behavior holds is enough.

  #### Solution (S1–S4)

  - **S1:** The solution should meet all the requirements. (If it's missing a requirement and it manages to pass your tests, you're in a bad position 😠.)
  - **S2:** No regressions and follow existing code patterns. Don't break existing working code by mistake (we'll still run the repo's existing tests).
  - **S3:** No irrelevant changes to the code — if something is unrelated to your task, keep it as it is.
  - **S4:** No AI slop (weird comments, unexplained defensive code, new coding patterns, etc).

  ### Gate 1: Repo And Seed

  Before editing, verify the pinned repo, clean/known worktree, Go-first scope, dependencies, tests, and Docker feasibility. Read repo policy files before editing. Work only in `/home/shrutikcs/Documents/shipd/<repo>` or an explicitly declared challenge workspace; never edit source in `/tmp`.

  Use authenticated `gh` CLI to inspect open/closed/merged PRs, issues, Discussions, full comments/reviews, and linked context. Read `README.md`, `CONTRIBUTING.md`, `AGENTS.md`, docs, examples, design notes, and relevant code comments. Stop on direct PR overlap, negative maintainer direction, unsupported architecture, or non-Go product scope. Submission-similarity results are evidence to inspect, not an automatic gate; reject only material duplication.

  The Issue Miner seed may be small. Builder owns expansion: enumerate adjacent public outcomes in the same workflow, such as persistence, execution, validation, serialization, CLI/API, backend variants, recovery, import/export, observability, or lifecycle behavior. Do not add unrelated parity, API inventory, schema/report plumbing, wrappers, convenience helpers, or test-only scope.

  Before artifacts, record the seed, expanded behaviors, affected subsystems/files, overlap and design evidence, Docker/base-test feasibility, nearest analog, and a measured size proof. Measure comparable non-test files or PR stats, remove excluded lines, separate reusable from new work, and estimate the smallest complete expanded implementation. Do not reject the seed for being small before attempting natural expansion; reject after expansion if it is artificial, unsupported, centralized, or below the final tier.

  For multiplicity/family work, identify sibling differences and generic-dispatch risk, implement 1-2 siblings as calibration, and re-scope if the family centralizes or lands materially below estimate. The smallest plausible seed-only fix must fail expanded user-visible behaviors under fair tests.

  ### Gate 2: Problem Contract

  Write `problem.md` last and keep it as short as possible, normally `80-160` ASCII words. It must be a natural maintainer issue: the first sentence is the request, followed by concise prose with no headings, bullet requirement list, code block, motivation, or "currently missing" preamble.

  Describe observable inputs, outputs, workflow effects, persistence/readback, ordering, errors, and edge cases. Use public names only when required for fairness. Do not name private helpers, internal files, storage keys, fixtures, assertion strategy, implementation types, or mechanisms. Every clause must be tested or clearly repo-derivable.

  Audit `P1-P7`: repo philosophy, no direct PR solution, self-contained, clear, verifiable, non-prescriptive, and materially distinct. Similarity matches require manual comparison, not automatic rejection. The solver sees only the pinned repo and `problem.md`, not challenge tests or the reference solution.

  ### Gate 3: Tests And Harness

  `test.patch` must contain only new tests and root `test.sh`; do not edit upstream tests or shared helpers. Use public package/user-facing entry points, self-contained fixtures, and random 3-byte hex suffixes from `openssl rand -hex 3` for new test filenames (ending in `_test.go` per Go convention). Do not include platform branding such as `Olympus`, `Shipd`, `MARS`, or `challenge` in patches.

  Tests must be deterministic and offline: no network, unseeded randomness, real sleeps in assertions, host state, locale/timezone/CPU/filesystem-order dependence, goroutine-scheduling dependence, or fail-fast flags. Every new test must fail on the fresh commit for its own missing behavior and pass with the solution; collection-only failure is insufficient. Assertions must be precise public outcomes, not private structure, internal storage, exact incidental formatting, or unspecified errors/order. Make sure There are tight behavioural Tests & new Tests.

  Negative behavior is first-class coverage. For every documented invalid, missing, duplicate, malformed, unsupported, unauthorized, conflicting, or partial-failure path, add a direct public behavioral test. The test must fail on the base commit, pass with the reference, and fail an implementation that accepts the input, returns/panics for the wrong reason, mutates state before rejecting, or leaves partial state after failure. Assert the documented rejection, error, or rollback outcome; an arbitrary error check or non-crash check is insufficient. Cover boundary values and failure ordering where they change the public result.

  Build a coverage matrix for every claimed topology, dialect/backend, input form, operation, and lifecycle state. Test meaningful cross-products in `new`, including non-default implementations and alternate input forms; remove untested claims from `problem.md` and `solution.patch`. Cover all named variants, backends, resource kinds, persistence/restart paths, boundaries, failures, idempotency, and obvious shortcuts. Almost all meaningful feature code must execute under `new`; base-only coverage is insufficient except for narrow compatibility preservation.

  Reference-to-test closure is mandatory: after writing `solution.patch`, map every meaningful hunk, branch, fallback, error path, state transition, variant, and public effect to a direct behavioral `new` test. Revert or disable each hunk and rerun `new`; if tests still pass, add the missing behavioral case or remove the unnecessary code. Do not write white-box tests for code with no observable effect; remove that code instead. Repeat until no meaningful reference behavior is untested.

  `test.sh` must accept exactly one `base` or `new` mode and exactly one nonempty `--output_path PATH` or `--output_path=PATH`; reject missing, duplicate, unknown, mixed-mode, and extra arguments. It must run only challenge tests in `new`, run real touched-surface upstream tests in `base`, emit non-empty JUnit XML at the requested path (e.g. via `gotestsum --junitfile` or `go-junit-report`), use `go test ./...` when `CHALLENGE_DOCKER=1`, return the runner status, and never install/fetch/mutate outside the repo.

  If the full upstream suite passes offline, run it excluding only new challenge tests. Otherwise include every touched surface and adjacent suite, recording only individually proven flaky or environment-bound exclusions. Never hide regressions with `-run` filters, build-tag exclusion, `t.Skip` conversion, fake XML, or a token smoke test.

  ### Gate 4: Solution Patch

  `solution.patch` contains implementation changes only: no tests, Dockerfile, `test.sh`, docs, generated files, unrelated formatting/import churn, dead code, padding, or whitespace-only hunks. Keep the implementation as short as possible while meeting the tier naturally and follow repo conventions, ownership boundaries, types, errors, serialization, and command wiring.

  Every new public symbol, field, output, persisted shape, or configuration maps to one problem sentence and one `new` test. Remove untested stats, fingerprints, reports, extra constructor fields, aliases, convenience APIs, duplicate helpers, unused fields, broad error suppression, and visible-test workarounds. Run applicable `go vet`, `gofmt`/`goimports`, and lint checks, then apply `S1-S4`.

  **GATE 5 - task quality (holistic AI review)**

  Read description + test.patch + solution.patch together: authentic engineering (maps to a real issue/ticket, not contrived/pedagogical/toy); readable even if dense; every required behavior specified; fair (tests align with the description, no gotchas); ONE coherent repo-wide feature (not an unfocused rewrite or padded pile); not trivial (no one-liner / single-file CRUD); objectively verifiable (deterministic, precise assertions); long-horizon and system-level (crosses API + impl + state + signals + error paths, or equivalent). Target ~1 passing strong rollout in 10 - an existence/fairness proof, NEVER permission to shrink an enormous fair task; do NOT run ten local agents to manufacture a rate (`tests-and fairness.md`,`SKILL.md`). Bounded local check:
  A prose claim without those durable `eval/` artifacts is UNVERIFIED. An edit to description/tests/solution stales the dependent blind-solve evidence for the changed surface, and editing test.patch to apply an on-site-audit recommendation means the site checks must re-run against the new tests - re-check that surface, never a full proactive re-run of everything.

  **GATE 6 - solvability + difficulty + verifier completeness**

  - 6.1 Blind reachability/fairness probe: one fresh description-only agent (no reference, no hidden tests). A passing solution is strongest evidence; a FAIR incomplete failure must NOT trigger scope reduction. Preserve its patch + base/new JUnit + report under `eval/`.
  - 6.2 Difficulty ~1/10 by DESIGN - judge sequential dependency, blast radius, invariants, horizon.
  - 6.3 Confirm Gate 4's scope thresholds hold on the reference + blind run: >= 4 files, >= 2 subsystems, 250-600 true logic LOC (within the band, 300-500 preferred, never over the 600 cap), >= 120 meaningful tool calls (never count tests or padding).
  - 6.4 No cheating: < 20% of runs hardcode test values / stub the tests / peek at test.patch.
  - 6.5 Verifier completeness (`tests-and-fairness.md` "Verifier completeness = a LOCAL agent check + the on-site audit"): checked TWO ways, run BOTH. (a) DISPATCH a LEAN local completeness-adversary agent - ONE bounded `general-purpose` agent that enumerates the promised behaviors + set members, designs the least-work partial/wrong solutions, and reports any behavior still passing the current suite (NO mutants, NO docker; NOT the old token-heavy 3-lens panel). (b) The USER runs the platform's on-site verifier completeness audit and COPIES its recommendations back - do NOT let the site auto-apply them. For BOTH sources: YOU vet each candidate gap for fairness, edit `test.patch` yourself for the fair ones (Bypass the unfair site recs on the site), cheap-verify each (base FAIL individually / solved PASS), and RE-VALIDATE the changed surface once. No baseline snapshot or diff is kept.
  - 6.6 PRE-SUBMISSION LOCAL REVIEW SIMULATION (`tests-and-fairness.md` "Pre-submission local review
    SIMULATION") - the mandatory LAST step before packaging, simulating the site's review axes so advisories are caught locally in ONE pass instead of many site round-trips: the 8.5 local completeness-adversary agent (the test-comprehensiveness pass, backstopped by the on-site audit) + ONE fairness-gate agent (vet each test you wrote, the local agent flagged, or the audit proposes FAIR/UNFAIR, keep only fair ones + the cheap local f2p check) + ONE solution-comprehensiveness / code-quality agent (fix any un-implemented behavior, dead code, non-idiomatic code, latent bug - keep the golden lean and the description
    behavioral). Preserve the completeness gap report + fairness rulings + solution-quality fixes under`eval/`; re-run the docker matrix + preflight ONCE after applying (do NOT re-run 8.1 for test-only or description-only edits - reason about the specific change directly).

  ### Alignment And Agent Feedback

  Create a three-column map: `problem.md` behavior, `new`/base test coverage, and solution hunks. Every behavior needs direct `new` coverage and implementation; every assertion is `problem-stated` or `repo-derivable`; every meaningful feature hunk maps to tested behavior; every touched file has base coverage or an honest exception. Remove anything unmapped.

  When agent runs or trajectories are provided, inspect decision points and attempted patches, not only the verdict. If a failure came from an ambiguous or undiscoverable observable behavior, add the smallest fair requirement to `problem.md` and a matching behavioral test, never an implementation hint. If it was a valid shortcut or genuine difficulty, strengthen behavior coverage or leave the contract unchanged. Re-run alignment and Docker validation after changes.

  Before counting an agent as a legitimate pass, adjudicate it independently: derive fair probes from `problem.md` and the pinned repo, isolate each public behavior and meaningful cross-product, and run the same probes in Docker against the candidate and reference patches. A candidate pass is invalid only when the candidate fails a prompt-stated or repo-discoverable probe that the reference passes. If both fail, treat it as a shared test/spec/reference gap and repair that gap before using the result.

  For every false-positive report, first reproduce why the incorrect solution passes and identify the exact missing assertion or weak boundary. Then choose one repair path: clarify the smallest ambiguous observable behavior in `problem.md`, or add the smallest tight public behavioral `new` test that kills the shortcut. Cover all meaningful edge partitions without arbitrary Cartesian expansion. Compare the proposed test with Shipd's Test Fairness advisory/coverage suggestions; matching suggestions support the gap, but the test still must fail on base and the inaccurate agent, pass with the reference, and satisfy T1-T7. Keep `problem.md` limited to behavior covered by tests. Reject implementation-coupled, undiscoverable, already-covered, or reference-failing probes. Repeat the audit after every repair.

  ### FP-Readiness Audit

  Shipd runs the actual FP Check after agent runs; Builder cannot run that website check. Before submission, loop this local gap audit until a complete pass finds no new gap, so passing agents cannot satisfy only the visible tests:

  - Build the behavior matrix from `problem.md` and create adversarial probes for every claimed edge, variant, backend, input form, boundary, failure, concurrency, and restart path.
  - Require every probe to fail on the base and pass with the reference in Docker using the actual patches.
  - Remove or rewrite any probe that is not prompt-stated or repo-discoverable, and repair any reference failure before spending agent runs.
  - Add a matching behavioral `new` test for every accepted probe, then rerun the alignment and clean-room matrix.
  - For every available passing agent, compare its patch against the same probes and add a test for each candidate-only failure before submission. Do not claim that the website FP Check was run.
  - After every test, problem, solution, or scope change, rebuild the matrix and repeat the loop; stop only when all probes are prompt-stated, base-failing, reference-passing, directly tested, and no available passing agent has a candidate-only failure.

  ### Final Validation

  Build the Docker image from the pinned commit and apply actual patches in a clean sequence:

  1. Apply `test.patch`; verify `base` passes and `new` fails.
  2. Apply `solution.patch`; verify both modes pass with `-network none`.
  3. Repeat suspicious tests, inspect JUnit XML, run vet/lint checks, measure meaningful LOC/file spread, and confirm no patch conflicts or stale hunks.

  For every Verify Tests or Verify Solution claim, use Docker from the pinned commit with the actual patches applied before changing artifacts or reporting success. If validation cannot run, report the exact command, blocker, and risk; do not claim readiness from partial evidence.

  ### Docker And Patches

  Build is in amd64, & not arm64.

  Use:

  dockerfile

  ```docker
  FROM public.ecr.aws/d3j8x8q7/olympus-base-go:latest
  WORKDIR /app
  COPY . .
  ENV CHALLENGE_DOCKER=1
  RUN go mod download
  CMD ["/bin/bash"]
  ```
  ````

  Install dependencies at build time. Do not test, fetch, clone, fabricate `.git`, or install at runtime. Validate `.dockerignore`, real pinned metadata (`go.mod`/`go.sum`), and offline execution. Generate both patches from staged source with `git diff --cached`; never hand-edit patches. Confirm `git apply --check`, executable `test.sh`, ASCII `problem.md`, clean patch boundaries, and no platform branding before submission.

  ```

  ```

- check Submission

  ```markdown
  ## Go Submission Reviewer -- Critic Only

  Review the actual `problem.md`, `test.patch`, `solution.patch`, `Dockerfile`, `repo_url.txt`, extracted `test.sh`, changed source, relevant upstream tests, and available agent runs. Do not fix files or soften findings. Decide `ACCEPTED` or `REVISION REQUESTED` with concrete evidence and repair requests.

  ### Review Order

  Do not stop at a shallow checklist. Run these gates in order and report incomplete gates:

  1. **Alignment and uniqueness preflight:** use authenticated `gh` CLI, not title-only search. Read README, CONTRIBUTING/AGENTS, docs, examples, design notes, code comments, and full PR/issue bodies, comments, reviews, linked context, and Discussions via `gh api`/GraphQL. Search semantic variants and prior-submission similarity results. Reject immediately for repo-design mismatch, negative maintainer direction, material duplicate, or open/closed/merged PR overlap. If `gh` is unavailable or unauthenticated, report the preflight incomplete.
  2. **Artifact and patch pass:** verify pinned immutable commit, repo eligibility, canonical checkout, clean patch application, patch boundaries, executable root `test.sh`, no conflicts, no platform branding, and correct Docker shape.
  3. **Alignment pass:** map every `problem.md` behavior to `new` tests and solution hunks; classify every assertion as `problem-stated` or `repo-derivable`. Anything unsupported or unmapped is a finding.
  4. **Shortcut and coverage pass:** inspect agent trajectories/patches, try plausible wrong implementations, and check every claimed topology, dialect/backend, input form, operation, lifecycle state, variant, resource kind, boundary, restart path, and sync/async path. Every T3/T4 finding must say: "A solution that <wrong behavior> would still pass because <missing assertion>."
  5. **FP-readiness audit:** Shipd runs the actual FP Check after agent runs; this prompt cannot run it. Before submission, loop until a complete pass finds no new gap: derive fair adversarial probes from `problem.md` and the pinned repo, vary one public behavior dimension at a time and then meaningful cross-products, run them in Docker against base, reference, and available passing-agent patches, and add a `new` test for every candidate-only failure. Also map every meaningful reference hunk, branch, fallback, error path, state transition, variant, and public effect to behavioral coverage; revert each hunk and require `new` to fail, or remove the unnecessary hunk. Fix any shared spec/test/reference gap, rebuild the matrix after every change, and stop only when all probes are prompt-stated, base-failing, reference-passing, directly tested, and no available passing agent has a candidate-only failure. Do not claim the website FP Check was run.
  6. **Solution pass:** inspect every hunk for bugs, regressions, dead/duplicate code, public-surface creep, broad error handling, collateral edits, style/type/import failures, stale hunks, and visible-test workarounds. Ask whether removing each meaningful hunk still lets `new` pass.
  7. **Race and regression pass:** repeat suspicious tests; inspect polling, timers, goroutines, filesystem moves, retries, caches, subprocesses, partial failures, and concurrent reads/writes. Audit adjacent readers after persistence/type/serializer changes and equivalent sync/async or backend paths after partial fixes.
  8. **Difficulty and matrix pass:** measure meaningful LOC, inspect passing agents, identify seed-only/read-time/centralized shortcuts, and verify base commit + `test.patch` gives base pass/new fail while both patches give base/new pass in Docker offline. A skipped matrix step is residual risk.
  9. **Final adversarial pass:** search once more for a missing edge test, removable hunk, race, implementation leak, collateral edit, and too-small passing shortcut. Do not accept unless all gates complete.

  For every Verify Tests or Verify Solution claim, run Docker from the pinned commit with the actual patches applied before changing artifacts or reporting success. Re-check every prior fairness advisory with evidence or explicitly mark it out of scope.

  For a false-positive report, reproduce why the incorrect solution passes and name the missing assertion or weak boundary before changing artifacts. Choose either the smallest observable clarification or the smallest tight public behavioral test that kills the shortcut, while covering all meaningful edge partitions. Compare the test with Shipd's Test Fairness advisory, require base failure/reference success/candidate failure, keep `problem.md` limited to tested behavior, and reject any undiscoverable, implementation-coupled, already-covered, or reference-failing probe.

  ### Output

  text
  ```

  Verdict: <ACCEPTED / REVISION REQUESTED>
  Problem Description: <1/3, 2/3, or 3/3> <Weak/Minor/Good>
  1. <P finding with evidence and requested repair>
     Tests: <1/3, 2/3, or 3/3> <Weak/Minor/Good> 2. <T finding with wrong solution, missing assertion, and repair>
     Solution & Code: <1/3, 2/3, or 3/3> <Weak/Minor/Good> 3. <S finding with path/hunk, impact, and repair>
     Other notes
     <completed gates, reference/passing-agent LOC, solve rate/effort, shortcuts, and residual risk>

  ```

  Findings come first, ordered by reproduced failures, repeat-run flakes, concrete regressions, coverage/spec gaps, then wording concerns. Do not claim confidence from environment timeouts or missing XML.

  ### Official Criteria

  #### Problem (P1-P7)

  - **P1:** Aligns with the repo's philosophy
  - **P2:** Not already fixed in an open or merged PR (I know, duplicate point. It is really important)
  - **P3:** Self-contained — solvable from the repo and description alone
  - **P4:** Clear, concise, and unambiguous — describes what to build or fix and don't leave points for guessing
  - **P5:** Verifiable — success is objectively testable
  - **P6:** Not prescriptive — don't leak the solution, we want to challenge the agents, remember!!
  - **P7:** Not a duplicate — we run a similarity check against existing submissions. Don't treat it as a gate; open the results, read the close matches it surfaces, and confirm you're not rebuilding a problem that already exists. Rewording the same task or reshaping the same behavior doesn't make it new.

  #### Tests (T1-T7)

  - **T1:** They highlight the missing or incorrect behavior (depending on the task's category — feature request, bug fix, etc). They should 100% fail at the base commit and 100% pass after adding your solution.
  - **T2:** The tests should be deterministic (no timing, randomness, or ordering; nothing that could change across multiple runs or across different machines).
  - **T3:** Strong tests. The tests shouldn't be permissive enough to let inaccurate agent solutions pass.
  - **T4:** Extensive coverage. The tests should cover the requested behavior and all the obvious edge cases.
  - **T5:** DO NOT check for unspecified or undiscoverable behavior — this will be unfair to the agents to expect them to implement something not in the description or discoverable from the repo 😞.
  - **T6:** They don't require any network connection (we run the container with --network none).
  - **T7:** Don't over-pin the output — don't assert on exact output (error text, messages, wording, formatting) unless the description says so or it's obvious from the repo's existing patterns. Otherwise it's unfair to an agent that gets the behavior right but words it differently; checking the behavior holds is enough.

  #### Solution (S1-S4)

  - **S1:** The solution should meet all the requirements. (If it's missing a requirement and it manages to pass your tests, you're in a bad position.)
  - **S2:** No regressions and follow existing code patterns. Don't break existing working code by mistake (we'll still run the repo's existing tests).
  - **S3:** No irrelevant changes to the code — if something is unrelated to your task, keep it as it is.
  - **S4:** No AI slop (weird comments, unexplained defensive code, new coding patterns, etc).

  ### Problem Description

  Use these keys: `P1` repo philosophy; `P2` no direct PR solution; `P3` self-contained; `P4` clear and concise; `P5` objectively verifiable; `P6` non-prescriptive and free of implementation leaks; `P7` materially distinct after similarity review.

  Request revision when `problem.md` uses internal types/files/helpers, constructor wiring, XML/context tags, exact private errors, storage/snapshot details, fixture names, assertion strategy, or mechanism instead of observable behavior; contradicts repo design; omits behavior needed for fairness; reads as an API inventory; or contains headings, bullets, code blocks, motivation, current-state filler, or non-ASCII. State observable precedence when tests depend on operation order, slot consumption, backfill, empty-key behavior, rollback, exact values, precision, or error distinctions; do not let tests choose between reasonable interpretations. The solver sees only the pinned repo and `problem.md`, not hidden tests, solution, or reviewer notes.

  ### Tests And Harness

  Use these keys: `T1` base/new/JUnit matrix; `T2` deterministic and non-flaky; `T3` shortcuts rejected; `T4` complete behavior/edge coverage; `T5` only stated or repo-discoverable behavior; `T6` no network/runtime fetching; `T7` no incidental output/format/order over-pinning.

  Request revision when:

  - Any new test passes on base, fails only through compilation/build failure, is not run by `new`, or does not fail for its own missing behavior.
  - Tests use private packages/helpers, undisclosed constructor shapes, internal storage, exact plumbing, weak non-crash checks, incidental ordering, exact command flags, or unspecified output/error text.
  - A named behavior, variant, backend, resource kind, topology, input form, lifecycle state, or expanded subsystem lacks direct behavioral `new` coverage.
  - A meaningful reference hunk, branch, fallback, error path, state transition, or public effect has no direct behavioral `new` test, or still survives when reverted while `new` passes. Add the missing behavior test or remove the hunk; do not add white-box tests solely to preserve unobservable code.
  - A coverage matrix is missing meaningful cross-products across topology, dialect/backend, input form, operation, and lifecycle state. Non-default implementations and alternate input forms need coverage when claimed.
  - Graph/DAG claims test only linear paths; TTL/time-window claims omit inside/outside/boundary cases; persistence omits empty, legacy, version, reuse, nested, restart, or metadata/history behavior; idempotent pairs test only one side.
  - Tests cover only a happy-path value subset: add negative/zero/maximum/precision and both operation directions, reflected or chained forms, exact boundaries, and semantic operand/unit/parameter order where the contract exposes them.
  - Invalid, duplicate, malformed, or partial inputs are untested; status/error branches, strict schema shapes, scope boundaries, and all advertised public entry points must be exercised without helpers that silently accept multiple shapes.
  - Negative behavior is first-class coverage. For every documented invalid, missing, duplicate, malformed, unsupported, unauthorized, conflicting, or partial-failure path, require a direct public behavioral test that fails on base, passes with the reference, and rejects acceptance, wrong-reason errors, mutation-before-rejection, or partial state. Require the documented rejection, error, or rollback outcome; arbitrary errors and non-crash checks are insufficient. Check boundary values and failure ordering when they affect the public result.
  - Failure injection does not cover partial creation/registration or a valid item before an invalid one, so rollback/continue-after-failure behavior and stale metadata remain untested.
  - Shared persistence or manifest behavior lacks a true concurrent multi-instance test for lost updates, write collisions, and readback integrity.
  - Recursive, chained, promoted, nested, or transitive behavior lacks alternate ordering, non-converging/depth-limit, or re-check coverage when the contract spans those cases.
  - Assertions check only a type, count, prefix/suffix, generic token, or broad error instead of semantic values, retained items, operand order, units, parameters, nested paths, or the exact public branch being promised. Test helpers must enforce one documented shape, and each test name must exercise the behavior it claims.
  - A feature advertises multiple CLI, stdio, server, parser, or integration entry points but tests only one route, or an integration test can pass through unrelated setup without exercising the named boundary.
  - Async/background changes are tested only after polling; first-read/restart semantics, sync equivalents, partial failures, or reads during concurrent moves can remain stale or fail.
  - `test.sh` accepts anything other than exactly one mode and one nonempty output path, allows duplicate/unknown/mixed/extra arguments, uses fail-fast behavior, omits native JUnit XML, silently falls back when the report path cannot be written, narrows base despite broader passing coverage, or hides failures with fake XML/deselection.
  - Base mode does not run the full offline upstream suite excluding only new tests and individually proven flaky/environment-bound tests, or omits adjacent touched-surface suites.
  - `test.patch` edits upstream tests/helpers or unrelated config, uses predictable filenames, network, unseeded randomness, real assertion sleeps, host state, or platform branding. New test filenames need `openssl rand -hex 3` suffixes (ending in `_test.go`).
  - Concurrency is a first-class contract but the repository's race detector (e.g. `go test -race`), sanitizer, or equivalent concurrency check is not run for the relevant new tests.
  - Every material test-fairness advisory must be closed by a direct test or explicitly documented as invalid/out of scope before acceptance.

  ### Solution And Difficulty

  Use these keys: `S1` all stated requirements; `S2` no regressions and repo conventions; `S3` no unrelated changes; `S4` no AI slop, dead code, unexplained defensive logic, or strange patterns.

  Request revision when the solution is incomplete, breaks legacy behavior, fails the repository's exact CI vet/lint/format/import checks (e.g. `go vet`, `gofmt`, `goimports`, `golangci-lint`) or adds new errors over baseline, adds untested public stats/fingerprints/reports/fields/aliases/convenience APIs, contains dead or duplicate helpers/fields, duplicated semantic logic that can drift instead of a shared repo-native path, broad error suppression, over-decomposition, visible-test workarounds, unrelated formatting/docs, whitespace padding, stale hunks, or changes subsystem ownership. Reversible transformations must preserve exact user-visible values, durable state, active-policy scope, nested/strict restoration, and all public consumers. Dialect/backend output must be valid for the selected implementation, accepted values must reject undocumented aliases, full-domain precision must not be silently truncated, and partial failures must leave no stale registrations or global state.

  Nearly all meaningful feature code must execute under `new`; base-only coverage is insufficient except narrow compatibility preservation. The reference must contain `500+` honest meaningful added implementation LOC across `5+` existing non-test files and plausibly require `60+` agent messages. If a legitimate passing agent is below `500`, touches fewer than `5` files, or completes in substantially fewer than `60` messages, treat it as compression evidence. Meaningful LOC excludes blanks, comments, no-ops, tests, generated files (e.g. `*.pb.go`, `*_generated.go`, mocks, stringer output), imports, declarations, braces, punctuation-only lines, boilerplate, wrappers, metadata, and mechanical propagation.

  Inspect passing agents when available. Report their LOC, shortcut style, legitimate solve rate, effort/messages, and whether timeouts or killed wrappers distorted results. For 10 legitimate runs, the target is at least 1 and at most 4 passes, keeping the rate below 40%; fewer indicates unfairness or unsolvability, while more than 4 is too easy. Use the current assignment panel if it differs. A large reference does not compensate for tests that permit a small seed-only or centralized solution.

  ### Docker And Fairness

  Request revision if Docker does not use the required language base image, `/app`, `COPY . .`, build-time dependencies, offline execution, no tests during build, patch-independent startup, and `CMD ["/bin/bash"]`. A solver must derive all tested behavior from `problem.md` and the pinned repo; Docker or environment failures are not difficulty.

  Use `REVISION REQUESTED` if any category is `Weak`, any material shortcut passes, any problem/test mismatch or regression exists, or artifact hygiene is invalid. Do not fix files; report the exact behavior, path/test area, passing shortcut, evidence, and requested repair.
  ```

- if not 3/3 ( run it 2-3 times max)

  ```markdown
  i need u to go though this each and everyh points given out ther e,
  i need u to prepare my test patch and solution patchc as the end result
  should give me all parameters passing it farely with no false positives
  and 3/3 ,which is 100% of the expeceted result .

  Give me pointer to archeive 3/3 on all the sections.and implement it and try to make it 3/3
  ```

- Test COverage

  ```markdown
  ## Go Test-Coverage Auditor -- Critic Only

  You are a critic-only test-coverage auditor. Do not edit files or invent requirements. Given a pinned repo, `problem.md`, `test.patch`, `solution.patch`, `Dockerfile`, `test.sh`, and optional agent runs or coverage suggestions, find every valid missing behavioral test and return a Shipd-style report for another agent to implement.

  ### Ground Rules

  Use the Shipd test criteria:

  - **T1:** Every new test fails on the base commit and passes with the solution.
  - **T2:** Tests are deterministic across runs and machines; no timing, randomness, ordering, or host-state dependence.
  - **T3:** Tests reject inaccurate or shortcut implementations.
  - **T4:** Tests extensively cover the requested behavior and obvious edge cases.
  - **T5:** Tests check only behavior stated in `problem.md` or clearly derivable from the pinned repo.
  - **T6:** Tests run offline without network or runtime dependency fetching.
  - **T7:** Tests assert behavior, not incidental formatting, exact errors, internal structure, or unspecified ordering.

  Negative behavior is first-class coverage. For every documented invalid, missing, duplicate, malformed, unsupported, unauthorized, conflicting, or partial-failure path, add a direct public behavioral test that fails on base, passes with the reference, and rejects acceptance, wrong-reason errors, mutation-before-rejection, or partial state. Assert the documented rejection, error, or rollback outcome; arbitrary errors and non-crash checks are insufficient. Cover boundary values and failure ordering when they affect the public result.

  The solver sees only the pinned repo and `problem.md`, not hidden tests, `solution.patch`, reviewer notes, or coverage suggestions. Treat any undiscoverable assertion as invalid until the problem is clarified. Do not force a particular file, helper, type, algorithm, constructor shape, schema layout, or implementation path.

  ### Audit Loop

  Repeat the following loop until a complete pass finds no new valid gap:

  1. Read `problem.md`, the complete `test.patch`, `test.sh`, `solution.patch`, relevant source/tests/docs, and all available agent patches or trajectories. Verify the pinned commit and patch boundaries.
  2. Extract every observable contract sentence and build a matrix of its meaningful dimensions: public entry point, input representation, variant, backend/dialect, topology, operation, lifecycle state, persistence/restart state, failure mode, boundary, ordering, precision, and concurrency.
  3. Map every matrix cell to existing `new` coverage. Check combinations, not only independent samples, when the contract says the dimensions interact.
  4. Trace every meaningful reference-solution hunk, branch, fallback, error path, state transition, variant, and public effect to a behavioral test. If code has no observable effect, recommend removing the code rather than adding a white-box test.
  5. For each suspected gap, design the smallest public behavioral test that isolates it. Verify it should fail on the base and pass with the reference. If an agent patch is available, verify the candidate fails and the reference passes; otherwise classify the issue as a reference/spec gap, not a candidate gap.
  6. Rebuild the matrix after every accepted finding. Stop only when all prompt-stated behaviors are covered, all valid cross-products are tested, and no available passing agent has a candidate-only failure.

  ### Coverage Checklist

  Check every applicable category:

  - Public APIs, CLI commands, stdio/server routes, and advertised integration entry points.
  - All named backends, dialects, adapters, serializers, resource kinds, generic/async variants, and non-default implementations.
  - Alternate input representations, including text/query input versus in-memory objects when both are promised.
  - Linear, branching, fan-in, multi-input, nested, chained, promoted, recursive, and transitive workflows when claimed.
  - Success, invalid, duplicate, malformed, empty, missing, partial, unknown, and unsupported inputs.
  - Lower/upper boundaries, inclusive/exclusive limits, negative/zero/maximum values, precision, units, reverse/reflected/chained operations, and full-domain values.
  - Operation precedence, filtering/paging/backfill, slot consumption, ordering, retained items, exact semantic operands, and parameter direction when observable.
  - Persistence, legacy readback, versioning, restart, reload, reuse, nested/strict restoration, and metadata/history consumers.
  - Failure injection before registration and after partial creation, rollback atomicity, continue-after-failure behavior, stale registrations, and global/shared state cleanup.
  - Stop/start and idempotency pairs, immediate first-read behavior, sync/async parity, concurrent multi-instance writes, race-detector or sanitizer coverage when concurrency is promised.
  - Actual output semantics rather than type/count/prefix/suffix/token checks; helpers must assert one documented shape and test names must exercise what they claim.

  ### Harness Audit

  Report harness gaps separately:

  - `test.sh` must accept exactly one `base` or `new` mode and exactly one nonempty `-output_path` value, rejecting missing, duplicate, unknown, mixed, or extra arguments with usage and exit code 2.
  - Valid runs must execute the real tests, return the runner status, and write native non-empty JUnit XML to the requested path. No hand-rolled, synthetic, fallback, or silently discarded reports.
  - `new` must run every new test and fail on the base commit. `base` must cover every touched upstream surface and use the full offline suite when feasible, excluding only individually proven flaky or environment-bound tests.
  - `test.patch` must contain only repo-native tests and root `test.sh`; reject unrelated config, platform branding, upstream helper edits, network access, predictable filenames, and fail-fast behavior.

  ### Finding Rules

  Classify each suggestion before reporting it:

  - **VALID GAP:** public behavior is problem-stated or repo-discoverable, current tests do not cover it, base should fail, reference should pass, and the test uses a public behavioral path.
  - **CANDIDATE GAP:** the reference passes but an available agent fails; add a discriminating `new` test and explain the shortcut.
  - **REFERENCE/SPEC BLOCKER:** the reference also fails, the contract is ambiguous, or the behavior is not discoverable. Do not recommend a test until the problem or reference is repaired.
  - **ALREADY COVERED:** existing base/new tests or an equivalent public path already pin the behavior.
  - **UNFAIR:** requires private structure, undisclosed API shape, implementation detail, or unspecified output. Reject it rather than expanding the prompt.

  Every T3/T4 finding must include: "A solution that <wrong behavior> would still pass because <missing assertion>." Do not report a gap merely because the reference has code, if that code has no distinct observable effect or is already covered indirectly.

  ### False-Positive Remediation

  When a judge or passing-agent review reports a false positive, first reproduce why the incorrect solution passes. Identify the exact missing assertion or weak test boundary before proposing changes; a wrong solution passing is evidence that coverage is not tight enough, not a reason to copy the reference implementation.

  1. Use a problem clarification when the required observable behavior is ambiguous or undiscoverable. Add only the smallest user-facing rule; never describe the implementation to avoid.
  2. Use a behavioral test when the contract is already clear. Design the smallest tight public-path test that fails on the base and inaccurate agent, passes with the reference, and proves the missing outcome rather than an internal shape. Cover every meaningful edge partition, not arbitrary Cartesian combinations.
  3. Compare the proposed test with Shipd's Test Fairness advisory/coverage suggestions. Matching independent suggestions are strong evidence that the test closes the intended gap; disagreement requires base/reference execution and fairness classification.
  4. Reject the test as unfair when it is unspecified, implementation-coupled, already covered, or also fails the reference. Repair the problem or reference instead of adding it.
  5. Re-run the complete audit loop after every clarification or test and stop only when no candidate-only false-positive gap remains.

  Keep `problem.md` limited to behavioral details that have corresponding tests. Remove untested description clauses instead of adding tests for undiscoverable or incidental behavior.

  ### Output

  text
  ```

  Coverage Audit: <GAPS FOUND / CLEAN / REFERENCE-SPEC BLOCKED>
  Summary
  <short result, matrix dimensions checked, and validation commands/evidence>
  Valid Test Gaps
  1. [T1/T3/T4] <short behavioral title>
     Behavior: <problem sentence or repo evidence>
     Current gap: <what existing tests miss>
     Shortcut: A solution that <wrong behavior> would still pass because <missing assertion>.
     Test: <public setup, action, and precise observable assertion>
     Evidence: <paths, base/reference result, and candidate result if available>
     Harness Gaps
  1. [T1/T6/T7] <specific test.sh or patch problem and repair>
     Reference/Spec Blockers
  1. [P3/P4/T5/S1] <ambiguous, undiscoverable, or reference-failing behavior and required resolution>
     Rejected Suggestions
  1. <suggestion> -- <already covered, unfair, unspecified, or non-discriminating reason>
     Final Matrix
     Covered: <dimensions and cross-products>
     Unresolved: <none or exact cells>
     Next action: <implement listed tests, repair problem/reference, or CLEAN>

  ```

  Do not claim the website FP Check was run. Return findings only; do not implement tests.
  ```

- work on action needed
- Gap Finder

  ```markdown
  Do a deep and complete analysis on finding critical gaps that could allow an incorrect, incomplete, hardcoded, overly simplified, or otherwise buggy implementation to pass all tests.
  Review the following points if missing from the test and fix them :

  - Missing behavioral scenarios that are required by the problem description.
  - Edge cases that are not validated.
  - Boundary conditions that are not exercised.
  - Error handling or failure cases that are not covered.
  - State transitions, interactions, or combinations of behaviors that are not tested.
  - Cases where a partially correct implementation could still pass.
  - Cases where a hardcoded or shortcut solution could pass.
  - Cases where an implementation that handles only the tested examples, rather than the general behavior, could pass.
  - Cases where tests verify only common paths and miss important alternative paths.
  - Situations where different valid inputs should produce different outcomes but the tests do not verify those distinctions.
  ```

- work on action needed
- FP checks

  ```markdown
  # SHIPD FALSE POSITIVE MASTER PROMPT

  Act as a strict False Positive hunter, mutation tester, Verifier Completeness auditor, test coverage auditor, and benchmark fairness adjudicator.

  Your objective is to determine whether an incorrect, incomplete, hardcoded, overly simplified, partially correct, or otherwise buggy implementation can pass the complete test suite while violating the actual task contract.

  The False Positive check is one of the highest priority checks. The verifier must not allow a buggy or partially implemented solution to pass.

  Do not rely on static reasoning alone. Execute adversarial mutations and candidate implementations against the real test suite.

  ## 1. ESTABLISH AND ALIGN THE CONTRACT

  Review:

  - Repository at the target commit
    • Problem description
    • Existing tests and public behavior
    • test.patch
    • solution.patch
    • test.sh
    • Relevant public APIs/documentation available to the agent

  Treat ONLY these as requirements:

  1. Behavior explicitly required by the problem statement.
  2. Behavior clearly discoverable from existing public repository behavior.

  Do NOT treat additional behavior implemented only by the reference solution as required.

  Before refining tests, verify that the problem description, tests, and reference solution describe the same observable behavior.

  Look for:

  - Contradictions between prompt, tests, and solution.
    • Requirements present in the prompt but absent from tests.
    • Tested requirements absent from the prompt/repository contract.
    • Reference behavior not actually required.
    • Unnecessary information.
    • Requirement leaks.
    • Implementation details accidentally exposed as requirements.

  Remove unnecessary information and requirement leaks.

  Think of the task as an exam. If the prompt requires A, B, and C, the tests must meaningfully verify A, B, and C.

  A mismatch can mean:

  - Missing test coverage.
    • Missing/unclear prompt requirement.
    • Unfair test.
    • Intentionally flexible behavior.

  ## 2. ATOMIZE THE PROMPT

  Break every prompt clause into independently testable behavioral atoms.

  Do not assume one sentence equals one requirement.

  Example:

  "Tracks the most recent non-empty value."

  contains at least:

  1. Non-empty values update state.
  2. Empty values do not update state.
  3. A newer qualifying value replaces the previous one.

  Extract atoms for:

  - Positive behavior.
    • Negative behavior.
    • Defaults.
    • Every enumerated value.
    • Documented errors/failures.
    • Boundaries.
    • State transitions.
    • Repeated operations.
    • Ordering.
    • Interactions between behaviors.
    • Distinct valid input classes.

  Treat symmetric operations independently:

  - set / reset
    • apply / rollback
    • encode / decode
    • enable / disable
    • add / remove
    • start / stop
    • success / failure

  Testing one side does not prove the other.

  ## 3. DECOMPOSE THE REFERENCE SOLUTION INTO STATES

  Break the reference solution into logical states or milestones.

  A state represents a point where a specific task goal has been achieved.

  For every state identify all required functionality, including:

  - Public behavior.
    • Relevant helper-driven behavior.
    • Validation logic.
    • Error handling.
    • Edge cases.
    • Defaults.
    • State updates.
    • Branches.
    • Input distinctions.
    • Transitions into and out of the state.

  Helpers themselves do not need direct tests unless they are part of the public contract.

  Instead, verify their required effects through observable public behavior.

  For every state ask:

  "What functionality must work correctly for this state to be valid?"

  Then verify that the test suite exercises each contractual behavior.

  After testing individual states, inspect interactions between states.

  Look for scenarios that appear only when multiple components, operations, or transitions work together.

  Add tests for such interactions when they are part of the contract.

  ## 4. BUILD A REQUIREMENT TO TEST MAP

  For every behavioral atom and contractual state behavior:

  1. Identify the exact test/assertion verifying it.
  2. Explain what incorrect behavior that assertion rejects.
  3. Determine whether it verifies the actual observable contract.

  Classify every tested behavior as:

  - Explicitly required by the prompt.
    • Discoverable from public repository behavior.
    • Flexible or underspecified.
    • Outside scope.

  Flag atoms or state behaviors without meaningful verification.

  ## 5. USE TEST FAIRNESS EARLY

  Treat Test Fairness as an early coverage tool, not only as a fairness check.

  Review every 💡 suggestion carefully.

  For each suggestion determine:

  - Is this a genuine contract requirement?
    • Is it already covered?
    • Does it expose a real verifier gap?
    • Is it noise or an incorrect suggestion?
    • Would implementing it create an unfair hidden requirement?

  Address as many valid suggestions as possible.

  Ideally continue until:

  - No meaningful suggestions remain, OR
    • Every remaining suggestion has a clear reason for being incorrect, redundant, underspecified, or unfair.

  Never blindly implement Test Fairness suggestions.

  ## 6. USE VCA EARLY

  Use the Verifier Completeness Audit, VCA, during early iterations to identify major verifier gaps before spending tokens on smaller refinements.

  Prioritize large structural coverage problems first.

  Use VCA to answer:

  - Are major requirements completely untested?
    • Can entire pieces of functionality be omitted?
    • Can a substantially incomplete implementation pass?
    • Are important states or interactions absent?
    • Does the verifier actually distinguish correct from broadly incorrect solutions?

  Fix major VCA findings before spending significant effort on minor refinements.

  This is especially important when the iteration/token budget is limited.

  ## 7. VERIFY REFERENCE SOLUTION QUALITY

  Run the Solution Quality check during early refinement.

  The reference solution should consistently achieve the platform's maximum expected quality score, such as 3/3 in both categories when that scoring applies.

  If Solution Quality identifies a reference solution bug:

  1. Fix the reference solution.
  2. Determine whether a participant/agent could make the same mistake.
  3. Determine whether that mistake violates the actual contract.
  4. Check whether existing tests catch it.
  5. If they do not, add the smallest fair regression test.

  Do not treat reference solution bugs as solution-only problems when they reveal verifier gaps.

  ## 8. ACTIVELY SEARCH FOR FALSE POSITIVES

  Take the correct/reference implementation and inspect it hunk by hunk and branch by branch.

  For every behavioral atom, construct the smallest implementation that violates that atom while preserving as much other behavior as possible.

  Examples:

  - Flip a boolean.
    • Change a constant.
    • Delete a branch.
    • Stub a branch to a no-op.
    • Remove a match arm.
    • Ignore an argument/input.
    • Return a fixed value.
    • Always return/use a default.
    • Skip validation.
    • Ignore an error.
    • Skip a state update.
    • Preserve stale state.
    • Break one side of a symmetric operation.
    • Collapse distinct inputs into the same result.
    • Remove boundary handling.
    • Handle only examples appearing in tests.

  Actually RUN the relevant complete test suite after each mutation.

  Never conclude that a mutation "should fail" without executing it.

  Record:

  - Behavioral atom.
    • Mutation.
    • Violated behavior.
    • Test result.
    • Whether the mutation survived.

  If a mutation stays green, that behavior is not adequately verified.

  ## 9. TEST PARTIALLY CORRECT AND HARDCODED IMPLEMENTATIONS

  Explicitly attempt candidates that:

  - Implement A and B but ignore C.
    • Support only common/documented examples.
    • Handle success but not failure.
    • Handle one side of a symmetric pair.
    • Work only on the first invocation.
    • Ignore certain arguments.
    • Return expected constants.
    • Recognize only tested inputs.
    • Special-case fixtures.
    • Treat distinct inputs identically.
    • Implement only the happy path.
    • Avoid general logic while satisfying current examples.

  Vary valid inputs beyond existing fixtures.

  If different valid inputs should produce different observable outcomes, verify that tests enforce those distinctions.

  ## 10. EDGE, BOUNDARY, ERROR, STATE, AND INTERACTION COVERAGE

  Where required by the contract, check:

  - Empty inputs.
    • Single and multiple values.
    • Minimum/maximum boundaries.
    • Just below/at/above boundaries.
    • Duplicates.
    • Missing optional values.
    • Invalid values.
    • Different ordering.
    • Initialized/uninitialized state.
    • Repeated operations.
    • Success followed by failure.
    • Failure followed by success.
    • Forward and reverse state transitions.
    • Interactions between components/features.

  For documented failures verify:

  - Failure occurs when required.
    • Incorrect success is rejected.
    • Required public error behavior is preserved.
    • State remains correct after failure.
    • Subsequent behavior remains correct where specified.

  Do not require exact error text unless established by the contract.

  ## 11. AUDIT ASSERTION STRENGTH

  Identify assertions that execute code without proving the required behavior.

  Examples:

  - is_ok()
    • is_some()
    • length > 0
    • contains(...)
    • object exists
    • truthy/success checks

  If the contract requires an exact observable result, distinction, or state transition, verify that directly.

  For suspicious assertions, deliberately change the expected behavior/value to something incorrect and rerun the test.

  If it remains green, the assertion is decorative or insufficient.

  ## 12. STUB THE FEATURE

  Create an extremely incomplete implementation while preserving compilation.

  Examples:

  - Replace functionality with no-ops.
    • Return defaults.
    • Remove most new logic.
    • Stub major feature branches.

  Run the complete suite.

  Feature-specific tests should become heavily red.

  Investigate anything that stays green and classify it as:

  - Intentional backward compatibility/unrelated behavior.
    • Missing verification.

  ## 13. PROVE A FALSE POSITIVE

  A valid False Positive exists only when ALL of these hold:

  1. The candidate passes every relevant hidden/new test.
  2. It violates an explicit or repository-discoverable requirement.
  3. A fair public API probe exposes the violation.
  4. The probe fails on the incorrect candidate.
  5. The exact same probe passes on the reference/correct implementation.

  A surviving mutation alone is not sufficient justification for adding a test.

  For every potential FP:

  1. Identify the violated contract atom.
  2. Construct the smallest fair public API probe.
  3. Run it against the incorrect candidate.
  4. Confirm it fails.
  5. Run the exact probe against the reference.
  6. Confirm it passes.
  7. Add it only after proving the behavior is required.

  ## 14. FAIRNESS GATE

  Every existing or proposed assertion must answer:

  "Could an agent know this requirement from the problem statement or clearly discoverable public repository behavior?"

  If no:

  1. If the behavior genuinely belongs to the task, state it clearly in the problem description.
  2. Otherwise remove/reject the assertion.

  Never keep a hidden requirement merely because it kills a mutation.

  Flag tests requiring:

  - Unstated representations.
    • Unstated algorithms.
    • Implementation-specific strategies.
    • Internal helpers/private APIs.
    • Specific internal file layouts.
    • Exact error text not established by the contract.
    • External protocol semantics not named by the prompt.
    • Behavior existing only in the reference solution.

  Whenever possible, test observable public behavior rather than private implementation details.

  Do not infer requirements merely because the reference implements them.

  ## 15. PRESERVE FLEXIBILITY

  Do not interpret flexible wording more narrowly than the contract defines.

  Pay particular attention to terms such as:

  - canonical
    • normalized
    • opaque
    • stable
    • comma-separated
    • preserve
    • consistent
    • deterministic
    • valid
    • compatible

  If multiple implementations satisfy the stated contract, tests must permit them.

  ## 16. VERIFY TEST.SH

  Check:

  - Distinct base and new modes.
    • Correct base/new test selection.
    • No accidental fail-fast behavior.
    • Correct nonzero exit propagation.
    • Failures actually fail evaluation.
    • Valid per-test JUnit XML.
    • No missing testcases.
    • No duplicate testcases.
    • Correct testcase reporting.
    • No package installation during evaluation.
    • No network dependency.
    • Required determinism.

  Run and record exact base/new results before and after the correct solution where applicable.

  Do not hide individual failures behind aggregate summaries.

  ## 17. INSPECT PASSING AGENTS

  A passing agent is not automatically a correct agent.

  For every passing agent:

  1. Compare its implementation against every behavioral atom.
  2. Search for requirements it violated despite green tests.
  3. If found, verify the requirement is contractual.
  4. Create a fair public API probe.
  5. Confirm it fails on the agent.
  6. Confirm it passes on the reference.
  7. Add the smallest fair regression test.
  8. Rerun fairness and FP analysis.

  If FP analysis is available after one passing agent, analyze that pass immediately rather than waiting unnecessarily for additional rollouts.

  ## 18. FINAL ADVERSARIAL PASS

  After fixing discovered gaps, ask:

  "What is the simplest wrong implementation that can still pass everything?"

  Create it and RUN it.

  Try different behavioral atoms until you cannot construct a meaningfully incorrect implementation that simultaneously:

  1. Passes the complete suite.
  2. Violates an explicit or repository-discoverable requirement.
  3. Can be exposed by a fair public API probe.

  Do NOT declare PASS because:

  - Tests look comprehensive.
    • Coverage appears high.
    • Static analysis found nothing.
    • Existing examples pass.

  PASS requires executable adversarial testing.

  # TOKEN-EFFICIENT EXECUTION ORDER

  Do not blindly run every quality check on every iteration.

  Use this order:

  1. Pass required prechecks.
  2. Align prompt, tests, and reference solution.
  3. Atomize the contract and reference states.
  4. Review Test Fairness 💡 suggestions.
  5. Use VCA during early iterations to find major verifier gaps.
  6. Run Solution Quality until the reference solution is consistently at the expected maximum quality level.
  7. Fix major fair coverage gaps.
  8. Verify base/new tests.
  9. Perform mutation-based FP hunting.
  10. Test partially correct and hardcoded candidates.
  11. Audit edge/error/state/interaction coverage.
  12. Audit assertion strength.
  13. Stub the feature.
  14. Fairness-check every new assertion.
  15. Spend later refinement iterations primarily on Test Fairness and remaining verifier gaps.
  16. Run agent rollout(s).
  17. Inspect every passing agent for FP.
  18. Run platform FP.
  19. Convert valid discovered FPs into fair regression tests.
  20. Repeat until the verifier cannot be fairly defeated.

  Prioritize major structural gaps early. Avoid wasting limited tokens repeatedly running unrelated expensive checks before the verifier and reference solution are mature.

  # FINAL OUTPUT

  Return:

  STATUS: PASS or NEEDS CHANGES

  ## Contract Atoms

  Every independently testable requirement.

  ## State Model

  Reference solution states/milestones and required contractual behavior within each.

  ## Coverage Map

  Requirement/state behavior → exact test/assertion → contract classification.

  ## Test Fairness Review

  Valid suggestions addressed and rejected suggestions with reasons.

  ## VCA Findings

  Major verifier completeness gaps and their resolution.

  ## Solution Quality Findings

  Reference solution issues and any regression tests derived from them.

  ## Proven False Positives

  For each:

  - Violated requirement.
    • Incorrect mutation/candidate.
    • Existing suite result.
    • Public API probe.
    • Candidate probe result.
    • Reference probe result.

  ## Surviving Mutations

  Mutations that remained green.

  ## Missing Coverage

  Behavioral, edge, boundary, error, state, interaction, and generalization gaps.

  ## Weak Assertions

  Assertions that fail to prove required behavior.

  ## Hardcoding Risks

  Shortcut/example-specific candidates capable of passing.

  ## Unfair Tests

  Assertions enforcing behavior outside the contract.

  ## Underspecified Requirements

  Prompt wording that cannot fairly support strict verification.

  ## Rejected Probes

  Tests rejected because they would be unfair, with reasons.

  ## Stub Results

  Results from the deliberately incomplete implementation.

  ## test.sh Audit

  Base/new behavior, exit propagation, JUnit validity, duplicates, missing cases, network/install dependencies.

  ## Exact Test Results

  Base/new results before and after the correct solution.

  ## Alignment Issues

  Any mismatch or requirement leak between prompt, tests, and reference solution.

  ## Changes Made

  Each test/prompt/solution change and its contractual justification.

  ## Remaining FP Risks

  Anything not conclusively validated.

  ## Final Adversarial Result

  Last incorrect candidates attempted and whether the suite rejected them.

  Do not declare PASS until executable adversarial testing has been completed and no known fair contract violation can pass the verifier.
  ```

- Testing

  ```markdown
  I need you to prepare Isolated Environment, wherein you will Test aganist Test.
  patch. You are only Allowed to use the Repo & Problem.md file. No references
  from any other file. Run 5 gemini 3.5 Flash High model for this. Publish the result after you
  completed.
  ```
