You are fixing a test.patch for a Go Olympus challenge that fails Phase 1 of the Docker matrix with compilation errors. Execute immediately. Do not explain. Do not ask questions.

### Working Context
<!-- Fill in BOTH fields before pasting -->
REPO_LOCAL_PATH: [ABSOLUTE PATH TO CLONED REPO — e.g. C:\Users\you\repos\myrepo]
CHALLENGE_SLUG:  [SLUG — the folder name inside challenge/ — e.g. least-latency-selection-policy]

**Navigation:** Use your shell/terminal to run:
```
cd <REPO_LOCAL_PATH>/challenge/<CHALLENGE_SLUG>
```
Do NOT proceed if `problem.md`, `test.patch`, and `solution.patch` are not all present.

Read `problem.md`, `test.patch`, and `solution.patch` from this directory.

### The Problem

Phase 1 applies ONLY `test.patch` to the base commit (no solution). It is failing because the test file references struct fields, functions, types, or methods that do NOT exist until `solution.patch` is applied.

**Compilation errors from Phase 1:**
(paste the exact error lines here before running this prompt)

### Fix Rules — test.patch

Rewrite every test that causes a compile error so that:
1. It compiles cleanly on the BASE commit — zero reference to any symbol introduced by solution.patch
2. It FAILS on base with a logical assertion failure (e.g., expected X, got Y, or feature missing), NOT a compile error
3. It PASSES once solution.patch is applied
4. It uses ONLY public APIs that already exist in the base commit

To find which symbols are safe to use:
- Read the base repo source in the affected package
- Only call functions, access fields, and use types that exist BEFORE your solution is applied
- If you need to assert that a new field exists, do it via a round-trip (marshal → unmarshal → check output), not by directly referencing the field

Do NOT:
- Call functions that solution.patch introduces
- Reference struct fields that solution.patch adds
- Import packages that solution.patch adds
- Use any interface that only solution.patch satisfies

### Fix Rules — solution.patch (if Phase 2 new also fails)

If Phase 2 new also fails (some tests still fail even with both patches applied):
- Read the failing test names from the Docker output
- Identify which behaviors they require
- Add the missing implementation to solution.patch
- Every hunk in solution.patch must be traceable to a clause in problem.md

### Output

Overwrite `test.patch` (and `solution.patch` if needed) in the current working directory.
Then confirm:

test.patch fixed ✓  (compile-on-base issue resolved)
solution.patch fixed ✓  (only if it was also changed)
Symbols removed from test: <list the solution-only symbols you removed or worked around>
