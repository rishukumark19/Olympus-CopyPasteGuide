You are repairing a problem.md for a Go Olympus challenge. Execute immediately. Read problem.md from your current working directory. Apply only the findings listed below. Do not ask questions.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

### P-axis findings to fix:
Use the P-axis findings from `review_findings.md` in the current working directory. Apply every finding listed under "Problem Description". If there are no P-axis findings in the file, confirm and stop.

### Repair Rules

Apply the MINIMUM change that resolves each finding. Do not rewrite unrelated parts.

- 80-160 ASCII words maximum.
- First sentence is the request itself — no preamble, no motivation.
- Natural maintainer prose — full sentences, no bullets, no headings, no code blocks.
- Describes observable inputs, outputs, workflow effects, persistence/readback, ordering, errors, edge cases.
- Public names only when required for fairness.
- No private helpers, internal files, storage keys, fixture names, implementation types, exact method signatures.
- No em-dashes (—). No hard line-wrapping at ~70-85 chars.
- Every clause must correspond to a tested behavior.

### Output

Use your file writing tool to overwrite `problem.md` in the current working directory with the revised text.
Do NOT output the revised content in chat. Write to disk, then confirm with one short line.
