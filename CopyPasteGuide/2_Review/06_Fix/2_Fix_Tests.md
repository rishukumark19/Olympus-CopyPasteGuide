You are repairing a test.patch for a Go Olympus challenge. Execute immediately. Read problem.md and test.patch from your current working directory. Apply only the findings listed below. Do not ask questions.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

### T-axis findings to fix:
Use the T-axis findings from `review_findings.md` in the current working directory. Apply every finding listed under "Tests". If there are no T-axis findings in the file, confirm and stop.

### Repair Rules

Apply ONLY the fixes needed for the listed findings. Do not add unrelated tests.

For each finding, add the SMALLEST public behavioral test that:
- Fails on the base commit for its own missing behavior (not just compilation)
- Passes with the reference solution
- Fails an implementation that does the specific wrong thing identified in the finding
- Uses only public APIs — no private helpers, no undisclosed constructor shapes, no internal storage
- Is deterministic — no timing, randomness, goroutine ordering, host state
- Runs offline — no network access
- Does not over-pin exact output/error text unless contractually specified

Negative/failure path findings: assert the documented rejection outcome — not just "error occurred". The test must fail a solution that accepts the input for the wrong reason or leaves partial state.

New test filenames must use openssl rand -hex 3 suffix. No "challenge", "olympus", "shipd", "mars" in filenames or comments.

If the finding is about test.sh permissions (100644 vs 100755): fix the file mode in the patch header.
If the finding is about test.sh not running specific tests: fix test.sh to include them.

### Output

Use your file writing tool to overwrite `test.patch` in the current working directory with the revised patch.
Do NOT output the diff in chat. Write to disk, then list findings addressed:

Findings addressed:
- Finding #N: [one line — which shortcut this test kills]
