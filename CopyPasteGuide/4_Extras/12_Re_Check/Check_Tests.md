# Check_Tests.md — 3 separate atomic audits for test.patch

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.
# Run each pass independently. Paste one block at a time to a fresh LLM session.
# All 3 must return PASS before test.patch is considered clean.

---

## PASS 1 of 3 — Hygiene Check

You are doing a single-focus hygiene audit of a test.patch file. Execute immediately. Answer PASS or list every violation. Do not explain what is fine — only report violations. Do not ask questions.

### test.patch:
(read test.patch from your current working directory)

Check ONLY these hygiene rules (report each violation separately):
1. test.patch modifies any existing upstream test file or shared test helper → FAIL (T1 risk)
2. New test filenames do NOT use a random hex suffix (pattern: _XXXXXX_test.go) → FAIL
3. Any filename or comment contains: "olympus", "shipd", "mars", "challenge" (any case) → FAIL
4. test.sh uses fail-fast flags (--exitfirst, -x, -failfast, set -e without recovery) → FAIL
5. test.sh does NOT accept both "base" and "new" as modes → FAIL
6. test.sh does NOT accept --output_path argument → FAIL
7. test.sh executes go test without producing JUnit XML (no gotestsum/go-junit-report) → FAIL
8. test.sh base mode uses a narrow -run filter that excludes valid upstream tests → FAIL
9. test.patch includes any Dockerfile, generated file, or unrelated config change → FAIL
10. Any test makes a network call (http, external fetch, git clone at runtime) → FAIL
11. Any test uses real time.Sleep for assertion (not for setup) → FAIL
12. Any test uses random unseeded values in assertions → FAIL

Output format:
PASS
-- or --
FAIL:
1. [rule number] exact line or pattern from the patch that violates the rule

---

## PASS 2 of 3 — Coverage and Strength Check

You are doing a single-focus coverage audit of a test.patch against its problem.md. Execute immediately. Answer PASS or list every gap. Do not explain what is fine — only report violations. Do not ask questions.

### problem.md:
(read problem.md from your current working directory)

### test.patch:
(read test.patch from your current working directory)

Check ONLY these coverage rules (report each violation separately):
1. A behavior explicitly stated in problem.md has NO direct test → FAIL (T4 gap)
   For each missing behavior: "Behavior X is stated in problem.md but has no test."
2. A test asserts behavior NOT stated in problem.md and NOT clearly discoverable from the repo → FAIL (T5 violation — unfair to agents)
   For each: "Test X checks behavior that is not in the description."
3. An error/rejection behavior described in problem.md has no test for it → FAIL (T4 gap — negative paths are first-class)
4. Tests check only the happy path — no boundary values, empty inputs, or edge cases → FAIL (T4 gap)
5. A test assertion uses only: is_ok, is_some, len > 0, contains, or generic error check — without verifying the specific observable value → FAIL (T3 weak)
   For each: "Test X asserts existence but not the actual required value."
6. Tests assert exact error message text when the contract does not specify it → FAIL (T7 over-pinned)

Output format:
PASS
-- or --
FAIL:
1. [rule number] specific test name or behavior and what is wrong

---

## PASS 3 of 3 — Shortcut Resistance Check

You are doing a single-focus shortcut resistance audit of tests against a problem.md. Execute immediately. For each test, try to construct the simplest WRONG implementation that would still pass it. Answer PASS if no shortcuts survive, or list each surviving shortcut. Do not ask questions.

### problem.md:
(read problem.md from your current working directory)

### test.patch:
(read test.patch from your current working directory)

For each test in the patch, answer:
"A solution that <specific wrong behavior> would still pass this test because <missing assertion>."

If no shortcut exists for a test, write: "[test name]: No shortcut — assertions are sufficient."

Focus especially on:
- Tests that only check return type or existence (not value)
- Tests that use a single input value (hardcoding that value would pass)
- Tests that assert A → B but not the inverse B ≠ A (symmetric gaps)
- Tests that check success but not rejection of invalid inputs
- Tests that cover only the first call, not repeated/idempotent calls
- Tests that use only examples that appear in the test itself (no input variation)
- Tests for a success path with no corresponding failure/rejection test for the same operation
- Tests that verify a state change but not state-after-failure (partial failure leaves stale state)
- Tests that check only one backend/variant/dialect when multiple are promised by problem.md
- Cross-product gaps: A passes + B passes but A-then-B combination is never tested
- Ordering-dependent behaviors that are only tested in one order

Additionally, try these adversarial implementations and check if all tests still pass:
1. Return hardcoded expected values from the test fixtures
2. Implement only the happy path (no error handling, no edge cases)
3. Implement only the first operation in a multi-operation workflow
4. Ignore all input beyond the first argument
5. Implement A and B correctly but skip C entirely

Output format:
PASS — all tests resist shortcuts
-- or --
SHORTCUTS FOUND:
1. [test name]: A solution that <wrong behavior> would still pass because <missing assertion>
2. ...

For each shortcut found: suggest the minimum assertion to add to kill it.
