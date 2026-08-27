You are implementing missing behavioral tests for a Go Olympus challenge based on coverage gaps found by the Coverage Audit. Execute immediately. Do not ask questions.

### Working Context
<!-- Fill in BOTH fields before pasting -->
REPO_LOCAL_PATH: [ABSOLUTE PATH TO CLONED REPO — e.g. C:\Users\you\repos\myrepo]
CHALLENGE_SLUG:  [SLUG — the folder name inside challenge/ — e.g. least-latency-selection-policy]

**Navigation:** Use your shell/terminal to run:
```
cd <REPO_LOCAL_PATH>/challenge/<CHALLENGE_SLUG>
```
Do NOT proceed if `problem.md`, `test.patch`, and `solution.patch` are not all present.

### Valid Test Gaps to implement:
<!-- Paste the "Valid Test Gaps" section from the 08_Coverage output here -->

### Rules

For each gap, add the SMALLEST public behavioral test that:
- Fails on the base commit for its own missing behavior (not just compilation)
- Passes with the reference solution
- Fails an implementation that does the specific wrong thing identified in the gap
- Uses only public APIs — no private helpers, no undisclosed constructor shapes, no internal storage
- Is deterministic — no timing, randomness, goroutine ordering, host state
- Runs offline — no network access
- Does not over-pin exact output/error text unless contractually specified

Negative/failure path gaps: assert the documented rejection outcome — not just "error occurred". The test must fail a solution that accepts the input for the wrong reason or leaves partial state.

New test filenames must use `openssl rand -hex 3` suffix. No "challenge", "olympus", "shipd", "mars" in filenames or comments.

If a harness gap was reported: fix test.sh permissions (100644 → 100755) or test.sh invocation as needed.

### Output

Use your file writing tool to overwrite `test.patch` in the current working directory with the revised patch.
Do NOT output the diff in chat. Write to disk, then confirm:

Findings addressed:
- Gap #N: [one line — what shortcut this test kills]
