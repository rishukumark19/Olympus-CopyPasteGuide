You are building a Go Olympus challenge. Execute immediately. Do not explain. Do not ask questions. This session produces ONLY problem.md and test.patch. Do NOT produce solution.patch.

**Size target (non-negotiable):** The final challenge must require 500+ meaningful implementation LOC across 5+ existing non-test files touching 2+ subsystems. Meaningful LOC excludes blank lines, comment-only lines, imports, declarations without logic, braces, generated files, boilerplate, and mechanical propagation. Do NOT pad to hit numbers — scope must emerge from real requirements. If natural expansion of this seed cannot reach 500 LOC across 5 files, report that and stop.

**Cross-cutting is your strongest lever.** A solution that cuts across several layers or subsystems (e.g. parser → planner → executor, or config → CLI → storage) is harder for agents and carries more effective LOC than one confined to a single spot. If your idea keeps landing short on difficulty or LOC, expand it to the next adjacent layer in the same workflow — not a disconnected feature.

### Seed
[PASTE THE SEED SUMMARY FROM STEP 02 — include behavior, expansion scope, affected subsystems, and compression risk]

### Repository
URL: [REPO URL from Step 01/02]
Commit: [PINNED COMMIT HASH from Step 01/02]

**Automatic Clone & Checkout:** If the repository is not already cloned locally in your working directory, use your shell/terminal tool to clone it and check out the pinned commit automatically:
`git clone <URL> && cd <repo_name> && git checkout <COMMIT>`
Perform all subsequent build steps inside this cloned repository. Do not ask the user to clone it manually.

### Pre-Work (Do First)

**STEP 0 — Blocked Repo Check:** Before anything else, confirm the target repo is NOT on the project's blocked-repo list. If it is blocked, STOP immediately and tell the user to pick a different repo. This is the single most expensive mistake to make late.

Before writing any artifact:
1. Use authenticated gh CLI to inspect open/closed/merged PRs, issues, Discussions, full comments/reviews. Stop on direct PR overlap or negative maintainer direction.
2. Read README.md, CONTRIBUTING.md, AGENTS.md, docs, examples, design notes, and relevant code comments.
3. Record: seed behaviors, affected subsystems/files, overlap evidence, Docker/base-test feasibility, and measured size proof.
4. For multiplicity/family work, implement 1-2 siblings as calibration first. If the family centralizes into one adapter/file, re-scope — do not proceed with a seed that shrinks under expansion.
5. Expand the seed only through **natural same-workflow behaviors**: persistence, execution, validation, serialization, CLI/API, backend variants, recovery, import/export, observability, lifecycle behavior. Do NOT add unrelated parity, API inventory, schema/report plumbing, wrappers, or test-only scope.

### Gate 1: Repo and Seed Verification

Verify:
- Pinned repo is clean, Go-first, Docker feasible with public.ecr.aws/d3j8x8q7/olympus-base-go:latest
- go.mod/go.sum present, go mod download works offline
- Upstream test suite is stable and fast for the target package surfaces
- No direct PR overlap for any part of the seed + expansion
- The smallest plausible seed-only fix would fail the expanded user-visible behaviors under fair tests

### Gate 2: Tests and Harness (write these BEFORE problem.md)

test.patch must contain ONLY new test files + root test.sh. Do NOT edit upstream tests or shared helpers.

Test rules (T1-T8):
- **T0 — COMPILE ON BASE (most common failure):** Every test file added by test.patch MUST compile cleanly on the base commit with NO solution applied. Never reference struct fields, functions, types, or imports that solution.patch introduces. If you need to assert a new field exists, do it via a public round-trip (e.g. marshal → parse → check output) — not by accessing the field directly. Violating T0 causes Phase 1 of the Docker matrix to fail with a compile error. Verify this mentally before writing each test: "Does this test file compile if I delete solution.patch entirely?"
- T1: 100% fail at base commit, 100% pass with solution
- T2: Deterministic — no timing, randomness, goroutine scheduling, host state. **Never use bare `time.Sleep` in tests.** Instead, use a polling helper:
  ```go
  // Standard pattern — copy into every test file that needs it:
  func waitFor(t *testing.T, timeout, interval time.Duration, condition func() bool) {
      t.Helper()
      deadline := time.Now().Add(timeout)
      for time.Now().Before(deadline) {
          if condition() { return }
          time.Sleep(interval)
      }
      t.Fatalf("condition not met within %s", timeout)
  }
  ```
- T3: Strong — reject inaccurate/shortcut implementations
- T4: Extensive — all behaviors + obvious edge cases covered
- T5: Only stated or repo-discoverable behavior
- T6: Offline — no network (container runs with --network none)
- T7: No over-pinning exact output/error text unless contractually required
- T8: Failure diagnostics intact — no catch-all messages hiding real failures

Test file naming: use `openssl rand -hex 3` suffix in filename (ending in _test.go).
Do NOT include "challenge", "olympus", "shipd", "mars" anywhere in patches.
No fail-fast flags. Every test result must be visible.

Negative behavior is first-class: For every documented invalid/missing/duplicate/malformed/partial-failure path, add a direct public behavioral test that fails on base, passes with reference, and fails a wrong/partial implementation.

Build a coverage matrix for every: topology, dialect/backend, input form, operation, lifecycle state. Test meaningful cross-products in `new`, including non-default implementations.

Reference-to-test closure: After writing tests, map every meaningful behavior → test → assertion. Any behavior without a test: add the test or remove that behavior from problem.md.

test.sh requirements:
- Accepts exactly one mode: base or new
- Accepts exactly one nonempty: --output_path PATH or --output_path=PATH
- Rejects: missing, duplicate, unknown, mixed-mode, extra arguments
- Emits non-empty JUnit XML at the requested path (via gotestsum --junitfile or go-junit-report)
- **base mode: scope TIGHTLY to only the specific packages and test patterns that the solution touches** (e.g. `-run TestJetStream` not `./...`). Do NOT run the entire upstream suite — this causes 10–30 minute timeouts from unrelated tests. NEVER hides regressions within scope.
- new mode: runs ONLY the new challenge tests
- Uses go test ./... when CHALLENGE_DOCKER=1
- Returns the runner exit status
- Never installs, fetches, or mutates outside the repo
- **MUST use Unix LF line endings (not Windows CRLF).** After writing test.sh, always run: `git add challenge/<slug>/test.sh` then verify with `file test.sh` that it is not CRLF. On Windows, set: `git config core.autocrlf false` before any git operations.

### Gate 3: Problem Description (write LAST, after tests are locked)

Rules for problem.md:
- 80-160 ASCII words maximum.
- Natural maintainer issue prose: first sentence IS the request. No preamble.
- No headings, no bullet requirement lists, no code blocks, no "currently missing" preamble.
- Describe observable inputs, outputs, workflow effects, persistence/readback, ordering, errors, and edge cases.
- Use public names ONLY when required for fairness. Do NOT name private helpers, internal files, storage keys, fixtures, or implementation types.
- Every clause must be tested or clearly repo-derivable.
- No em-dashes (—). No hard-wrapped lines at ~70-85 chars. No wall-of-text AI tells.
- Audit P1-P7 before finalizing:
  - P1: Aligns with repo philosophy
  - P2: Not already fixed in any PR (open, merged, or closed)
  - P3: Self-contained — solvable from repo + description alone
  - P4: Clear, concise, unambiguous
  - P5: Verifiable — success is objectively testable
  - P6: Not prescriptive — no solution leaks
  - P7: Materially distinct from existing submissions

Common writing mistakes — any of these is a rejection:
- **Framed as external:** "Library X currently lacks Y" — write as a maintainer issue, not an outsider report
- **List of requests:** bullet points or numbered requirements instead of flowing prose
- **Code snippets when prose works:** write "model updates and deletes should..." not "`model.update()` and `model.delete()` should..." — use code only when the exact shape or signature is required for fairness
- **Discoverable repo details listed:** do not describe internal behavior the solver will find in the repo anyway
- **AI-written tone:** em-dashes, hard line-wrapping, "Additionally,", "Furthermore," — write it yourself

### Dockerfile

```dockerfile
FROM public.ecr.aws/d3j8x8q7/olympus-base-go:latest
WORKDIR /app
COPY . .
ENV CHALLENGE_DOCKER=1
RUN go mod download
CMD ["/bin/bash"]
```

Build is amd64, NOT arm64. Install dependencies at build time. No tests in Dockerfile.

Generate test.patch with: `git diff --cached` (never hand-edit patches).

**Pre-Docker gate — run these BEFORE building Docker:**
```bash
# From a CLEAN checkout of the pinned commit:
git stash && git checkout <HASH> && git clean -fd -e challenge/
git apply --check challenge/<slug>/test.patch   # MUST print no errors
git apply challenge/<slug>/test.patch
git apply --check challenge/<slug>/solution.patch  # Only after test.patch applied
```
If either `--check` fails, fix the patch on disk before running Docker. Never debug inside a running container.

Confirm: git apply --check passes, test.sh is executable (chmod +x in patch), no platform branding.

### Folder Setup (do this first)

Create: challenge/<short-slug-of-the-issue-title>/

> **This slug becomes CHALLENGE_SLUG** — write it down. Every downstream prompt (03B, 04, 05, 06_Fix/*, 07, 08) requires you to fill in `REPO_LOCAL_PATH` and `CHALLENGE_SLUG` before pasting.

Inside it, create:
- problem.md
- test.patch
- solution.patch (empty for now — filled in Step 04)
- Dockerfile
- repo_url.txt  — contains: <repo URL>\n<commit hash>
- commit.txt    — contains the pinned commit hash only
- title.txt     — contains the short issue title slug

### Output

After completing all gates and writing files, confirm:

Challenge folder created: challenge/<slug>/
CHALLENGE_SLUG: <slug>      ← **save this for all downstream prompts**
REPO_LOCAL_PATH: <absolute path to repo root>   ← **save this too**
Files written:
- problem.md ✓
- test.patch ✓
- Dockerfile ✓
- repo_url.txt ✓
- commit.txt ✓
- title.txt ✓

Then print the full content of problem.md for visual verification.

### Docker Verification (TEST CHECK — run immediately after writing files)

After writing all files, use your shell/terminal to run the Docker verification. Do NOT skip this. Execute every command directly in your shell tool.

```bash
# From the repo root (parent of the challenge folder):
HASH=$(cat challenge/<slug>/commit.txt)
git stash
git checkout $HASH
git apply challenge/<slug>/test.patch

# Build and run — try Docker first; if unavailable, run ./test.sh natively
docker build -t challenge-verify .
docker run --rm --network none challenge-verify ./test.sh --output_path /tmp/base1.xml base
echo "TEST_CHECK_BASE_EXIT=$?"
docker run --rm --network none challenge-verify ./test.sh --output_path /tmp/new1.xml new
echo "TEST_CHECK_NEW_EXIT=$?"

git stash pop
```

Expected results:
- TEST_CHECK_BASE_EXIT=0   → upstream tests pass (no regressions)
- TEST_CHECK_NEW_EXIT≠0   → challenge tests fail correctly (no solution yet)

Report this before finishing:

TEST_CHECK_BASE: <PASS | FAIL (exit code N)>
  git apply result: <"OK" | "FAILED — error: <exact error>">
  Failing tests (if any): <list or "none">
  Root cause: <"patch corrupt/conflict" | "upstream test broken" | "docker error" | "none">

TEST_CHECK_NEW: <FAIL (exit≠0) = CORRECT | PASS (exit=0) = BROKEN — tests must fail without solution>
  Failing tests: <list — these are expected to fail>

RESULT: <PASS | FAIL>

If RESULT is FAIL:
- TEST_CHECK_BASE FAIL + git apply error → test.patch is corrupt. Fix it before proceeding.
- TEST_CHECK_BASE FAIL + git apply OK → a test broke an upstream test. Identify and fix.
- TEST_CHECK_NEW PASS → challenge tests pass without a solution. Fix tests before proceeding.
If RESULT is PASS → proceed to Step 03B (problem description audit) in the same or new chat.

DO NOT produce solution.patch. This session ends after Docker verification passes.

