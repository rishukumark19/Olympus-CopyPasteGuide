# Check_Solution.md — 3 separate atomic audits for solution.patch
# Run each pass independently. Paste one block at a time to a fresh LLM session.
# All 3 must return PASS before solution.patch is considered clean.

---

## PASS 1 of 3 — Hygiene Check

You are doing a single-focus hygiene audit of a solution.patch file. Execute immediately. Answer PASS or list every violation. Do not explain what is fine — only report violations. Do not ask questions.

### solution.patch:
(read solution.patch from your current working directory)

Check ONLY these hygiene rules (report each violation separately):
1. solution.patch modifies any _test.go file → FAIL (S3 — tests must not be in solution)
2. solution.patch modifies test.sh → FAIL (S3)
3. solution.patch modifies the Dockerfile → FAIL (S3)
4. solution.patch contains whitespace-only hunks or formatting-only changes in untouched files → FAIL (S3)
5. solution.patch contains import reordering or gofmt changes in files unrelated to the feature → FAIL (S3)
6. solution.patch contains any generated file changes (*.pb.go, *_generated.go, mocks) → FAIL (S3)
7. solution.patch contains dead code: functions, branches, or variables that are never called/read → FAIL (S4)
8. solution.patch contains unexplained defensive code (broad error suppression, panic recovery with no test) → FAIL (S4)
9. solution.patch contains comments that explain what the code does rather than why (AI tell) → FAIL (S4)
10. solution.patch introduces a new coding pattern inconsistent with the rest of the repo → FAIL (S4)
11. Files in solution.patch contain: "olympus", "shipd", "mars", "challenge" (any case) → FAIL
12. solution.patch contains any network calls or runtime fetches → FAIL
13. solution.patch modifies Dockerfile or removes ENV CHALLENGE_DOCKER=1 → FAIL (this env var is required for go test ./... fallback in test.sh)

Output format:
PASS
-- or --
FAIL:
1. [rule number] exact file/hunk and what the violation is

---

## PASS 2 of 3 — Completeness and Correctness Check

You are doing a single-focus completeness audit of a solution.patch against a problem.md. Execute immediately. Answer PASS or list every gap. Do not explain what is fine — only report violations. Do not ask questions.

### problem.md:
(read problem.md from your current working directory)

### test.patch (summary of what is tested):
(read test.patch from your current working directory)

### solution.patch:
(read solution.patch from your current working directory)

Check ONLY these completeness rules (report each violation separately):
1. A behavior explicitly stated in problem.md is NOT implemented in solution.patch → FAIL (S1)
   For each: "Behavior X is required by problem.md but has no implementation hunk."
2. A test in test.patch would fail even with solution.patch applied (solution is incomplete) → FAIL (S1)
   For each: "Test X tests behavior that solution.patch does not implement."
3. A hunk in solution.patch has no corresponding test in test.patch → FAIL (S1 risk / dead code)
   For each: "Hunk in [file] at [location] has no behavioral test — either add a test or remove the hunk."
4. solution.patch duplicates logic that already exists in the repo with a slightly different implementation → FAIL (S2)
   For each: "Hunk X reimplements [existing repo function/pattern] instead of reusing it."
5. solution.patch changes the behavior of existing public APIs beyond what problem.md requires → FAIL (S2/S3)

Output format:
PASS
-- or --
FAIL:
1. [rule number] specific hunk/behavior and what is wrong

---

## PASS 3 of 3 — LOC and Quality Check

You are doing a single-focus quality and size audit of a solution.patch. Execute immediately. Answer PASS or list every issue. Do not explain what is fine — only report violations. Do not ask questions.

### solution.patch:
(read solution.patch from your current working directory)

Count and classify every added line (+) in solution.patch as one of:
- LOGIC: has actual branching, computation, state mutation, or behavioral effect
- BOILERPLATE: struct declaration, interface implementation, import, brace-only, blank
- MECHANICAL: propagation (wiring a value through without logic), getter/setter without validation

Check ONLY these quality rules (report each violation separately):
1. Meaningful LOC (LOGIC only) below 250 → FAIL (scope too thin — use 4_Extras/11_Tune/Expand.md)
2. Meaningful LOC above 600 → FAIL (scope too large — use 4_Extras/11_Tune/Reduce.md)
3. Fewer than 4 existing non-test files touched → FAIL (not cross-cutting enough)
4. Fewer than 2 distinct subsystems touched → FAIL (not cross-cutting enough)
5. Any hunk could be removed without any test failing (unverified dead code) → FAIL
   For each: "Hunk in [file] at [location] has no test coverage — can be safely deleted."
6. go vet / gofmt / goimports issues are present (misaligned, wrong formatting, unused imports) → FAIL

Output format:
LOC Summary:
- Total added: X
- Meaningful (logic): X
- Boilerplate: X
- Mechanical: X

PASS
-- or --
FAIL:
1. [rule number] specific issue

> **Note on LOC:** If agents are solving it in noticeably fewer lines than your solution, your real count is likely lower than it looks. Only LOGIC lines count. If meaningful LOC is near the boundary, do a manual hunk-by-hunk LOGIC/BOILERPLATE/MECHANICAL classification before accepting the count.
