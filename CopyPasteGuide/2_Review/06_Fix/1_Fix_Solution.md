You are repairing a solution.patch for a Go Olympus challenge. Execute immediately. Read problem.md, test.patch, and solution.patch from your current working directory. Apply only the findings listed below. Do not ask questions.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

### S-axis findings to fix:
Use the S-axis findings from `review_findings.md` in the current working directory. Apply every finding listed under "Solution & Code". If there are no S-axis findings in the file, confirm and stop.

### Repair Rules

Apply ONLY the fixes needed for the listed findings.

- S1: Solution must meet ALL requirements stated in problem.md. Implement any missing requirement now.
- S2: No regressions. Follow existing repo code patterns, types, error handling, serialization, and command wiring.
- S3: No irrelevant changes. Remove any hunk unrelated to the task — formatting churn, import reordering, doc changes, style fixes in untouched files.
- S4: No AI slop. Remove: weird comments, unexplained defensive code, new coding patterns inconsistent with the repo, dead branches, broad error suppression, over-decomposition.

For dead code findings: either add the minimum behavioral test that requires it (if it is a genuine contract requirement), or remove the hunk from solution.patch.

For regression findings: fix the specific hunk that breaks the upstream test. Do not change test files.

For LOC findings: if above 600, remove padding hunks that have no test. If below 500, check if natural same-workflow expansion exists before reporting.

### Output

Use your file writing tool to overwrite `solution.patch` in the current working directory with the revised patch.
Do NOT output the diff in chat. Write to disk, then list findings addressed:

Findings addressed:
- Finding #N: [one line]
