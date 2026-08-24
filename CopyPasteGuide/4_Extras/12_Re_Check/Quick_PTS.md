You are doing a fast recursive P/T/S alignment audit. Execute immediately. Do not ask questions. Read the repo docs, then fix violations in one pass.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

### Artifacts (read from your current working directory):
- problem.md
- test.patch
- solution.patch
- repo_url.txt (repo URL + commit)

---

### Step 1: Read the Docs First

Before auditing, read the repo's README, CONTRIBUTING, design docs, and any code comments relevant to the task area. This surfaces what a solver would actually find — and flags anything in problem.md that either leaks implementation or is already obvious from the repo.

---

### Step 2: Recursive P/T/S Violation Loop

Repeat until no new violations surface:

**P violations** (problem.md issues):
- Framed as external ("the repo lacks X") instead of maintainer prose
- Prescriptive: names internal helpers, file paths, exact method signatures the solver would find themselves
- Ambiguous: a solver could implement two different things and both look correct
- Undiscoverable: requires knowledge not in problem.md or the repo at the pinned commit
- Over-specified: lists discoverable details the solver will read in the repo anyway
- AI-written tone: em-dashes, hard wrapping, "Additionally", "Furthermore"

**T violations** (test issues):
- A test fails an agent for a reason not stated in problem.md or discoverable from the repo (unfair)
- A contract behavior in problem.md has no test (uncovered)
- A test passes even when a key behavior is wrong (weak assertion)
- A partial implementation — implements A and B but not C — still passes all tests

**S violations** (solution issues):
- Solution doesn't implement a behavior that problem.md requires
- Solution breaks existing tests (regression)
- Solution contains dead code, padding, or AI slop
- A meaningful hunk can be removed and all new tests still pass (hunk unverified)

**For each violation found:**
1. State the exact clause, test, or hunk
2. Apply the minimum fix
3. Re-check if the fix introduced a new P, T, or S issue
4. Continue until the loop completes with no new findings

---

### Output

List every violation found and the fix applied:

P findings:
1. [P#] <clause> → <fix applied>

T findings:
1. [T#] <test/behavior> → <fix applied>

S findings:
1. [S#] <hunk/path> → <fix applied>

STATUS: CLEAN (no violations remain) or NEEDS HUMAN REVIEW (edge case requiring human judgment — describe it)
