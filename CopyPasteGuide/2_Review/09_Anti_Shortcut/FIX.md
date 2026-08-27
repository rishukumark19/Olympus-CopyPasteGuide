You are fixing behavioral test gaps and false positives found by the Anti-Shortcut audit for a Go Olympus challenge. Execute immediately. Do not ask questions.

### Working Context
<!-- Fill in BOTH fields before pasting -->
REPO_LOCAL_PATH: [ABSOLUTE PATH TO CLONED REPO — e.g. C:\Users\you\repos\myrepo]
CHALLENGE_SLUG:  [SLUG — the folder name inside challenge/ — e.g. least-latency-selection-policy]

**Navigation:** Use your shell/terminal to run:
```
cd <REPO_LOCAL_PATH>/challenge/<CHALLENGE_SLUG>
```
Do NOT proceed if `problem.md`, `test.patch`, and `solution.patch` are not all present.

### Proven False Positives / Surviving Mutations to fix:
<!-- Paste the "Proven False Positives" and "Surviving Mutations" sections from the 09_Anti_Shortcut output here -->

### Fix Rules

**For each Proven False Positive (a wrong implementation that passes):**

Add the SMALLEST discriminating probe that:
- Fails on the incorrect candidate implementation
- Passes on the reference solution
- Uses only public APIs discoverable from problem.md or the repo
- Is NOT a private/internal check — if behavior is only discoverable via internal APIs, fix problem.md to state it instead

Fairness gate: before adding any assertion, confirm — "Could an agent know this requirement from problem.md or clearly discoverable public repo behavior?" If no, either add it to problem.md or drop the assertion.

**For each Surviving Mutation (a behavioral atom with no meaningful test):**

Add the smallest test that:
- Makes the mutation fail
- Passes on the reference implementation
- Is deterministic and offline

**For weak assertions (is_ok, length > 0, contains checks):**

Replace with a direct check of the required observable value.

New test filenames must use `openssl rand -hex 3` suffix. No "challenge", "olympus", "shipd", "mars" in filenames or comments.

**If problem.md needs a clause added (for fairness):**

Apply the MINIMUM addition. Keep within 80-160 word limit. No bullets, no headings, no em-dashes.

### Output

Use your file writing tool to overwrite `test.patch` (and `problem.md` if needed) in the current working directory.
Do NOT output diffs in chat. Write to disk, then confirm:

Files changed: <list>
False positives killed:
- FP #N: [mutation] → [probe added]
Mutations now failing:
- Mutation #N: [one line]
