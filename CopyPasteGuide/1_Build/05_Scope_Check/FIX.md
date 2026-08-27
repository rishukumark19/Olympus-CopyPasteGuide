You are repairing challenge artifacts for a Go Olympus challenge based on findings from the Scope Check. Execute immediately. Do not ask questions.

### Working Context
<!-- Fill in BOTH fields before pasting -->
REPO_LOCAL_PATH: [ABSOLUTE PATH TO CLONED REPO — e.g. C:\Users\you\repos\myrepo]
CHALLENGE_SLUG:  [SLUG — the folder name inside challenge/ — e.g. least-latency-selection-policy]

**Navigation:** Use your shell/terminal to run:
```
cd <REPO_LOCAL_PATH>/challenge/<CHALLENGE_SLUG>
```
Do NOT proceed if the relevant files are not present.

### Findings to fix:
<!-- Paste the FINDINGS list from the 05_Scope_Check output here -->

---

### Fix Order (always apply in this order)

**If S* findings exist — fix solution.patch first:**

Apply ONLY the fixes needed for the S findings.
- S1: Implement every missing requirement from problem.md.
- S2: No regressions — follow existing code patterns, do not break upstream tests.
- S3: No unrelated changes — remove any hunk outside the feature scope.
- S4: No AI slop — remove weird comments, unexplained defensive code, alien patterns.

Use your file writing tool to overwrite `solution.patch`. Do NOT output the diff in chat.

---

**If T* findings exist — fix test.patch second:**

Apply ONLY the fixes needed for the T findings.

For each finding, add the SMALLEST public behavioral test that:
- Fails on the base commit for its own missing behavior (not just compilation)
- Passes with the reference solution
- Uses only public APIs — no private helpers, no undisclosed constructor shapes
- Is deterministic — no timing, randomness, goroutine ordering, host state
- Runs offline — no network access
- Does not over-pin exact output/error text unless contractually specified

New test filenames must use `openssl rand -hex 3` suffix. No "challenge", "olympus", "shipd", "mars" in filenames or comments.

Use your file writing tool to overwrite `test.patch`. Do NOT output the diff in chat.

---

**If P* findings exist — fix problem.md last:**

Apply the MINIMUM change that resolves each P finding.
- 80-160 ASCII words maximum.
- First sentence is the request itself — no preamble, no motivation.
- Natural maintainer prose — full sentences, no bullets, no headings, no code blocks.
- Every clause must correspond to a tested behavior.
- No em-dashes (—). No hard line-wrapping at ~70-85 chars.

Use your file writing tool to overwrite `problem.md`. Do NOT output the revised content in chat.

---

### Output

Confirm what was fixed:

Files changed: <list>
Findings addressed:
- Finding #N: [one line]
