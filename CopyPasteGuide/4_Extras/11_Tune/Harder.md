You are tightening the tests for a Go Olympus challenge that is too easy. Execute immediately. Do not explain, do not ask questions. Add the minimum fair behavioral tests that eliminate the identified shortcuts. Do NOT invent new requirements. Do NOT change problem.md or solution.patch.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

### Current problem.md:
(read problem.md from your current working directory)

### Current test.patch:
(read test.patch from your current working directory)

### Current solution.patch:
(read solution.patch from your current working directory)

### Repository:
URL: (read from repo_url.txt in your current working directory)
Commit: (read from commit.txt in your current working directory)

### Blind Test Results (from Step 09):
[PASTE which agents passed, what they implemented, LOC count, files touched, shortcuts noticed]

### Surviving Mutations (from Step 08, if any):
[PASTE or write NONE]

---

### Hardening Rules

Goal: make the challenge require a complete, correct implementation — not just a shortcut that passes visible tests.

For each identified shortcut or surviving mutation:
1. Identify the exact behavioral atom the shortcut exploits.
2. Design the smallest public behavioral test that kills this specific shortcut.
3. Verify the new test:
   - Fails on base commit (for its own missing behavior, not just compilation)
   - Passes with the reference solution
   - Fails the shortcut/incorrect implementation
   - Does NOT fail a complete correct implementation
4. Check against fairness:
   - Only tests behavior stated in problem.md or clearly discoverable from the pinned repo
   - Does not require private APIs, undisclosed internal shapes, or exact error text unless contractually required
   - Does not make the task unfair for an agent that implements everything correctly

Do NOT:
- Add tests for behavior not covered by problem.md (unfair T5 violation)
- Add tests that require implementation-specific internals
- Tighten assertions to exact incidental output unless the contract requires it
- Add tests just to reduce pass rate — only add tests for real behavioral gaps

Test rules (T1-T8 must hold for every new test):
- T1: Fails on base, passes with solution
- T2: Deterministic
- T3: Rejects the specific shortcut
- T4: Covers the edge the shortcut exploits
- T5: Only stated or discoverable behavior
- T6: Offline
- T7: No over-pinning of incidental output
- T8: Failure diagnostics intact

New test filenames: use openssl rand -hex 3 suffix. No "challenge"/"olympus"/"shipd"/"mars".

### Output

Revised artifacts (only the ones that changed):

**problem.md** (if Type A):
Save the revised text directly to the file. Do not print it in chat.

**test.patch** (if Type B or C):
Save the revised patch directly to the file. Do not print it in chat.

New tests added:
- Test <name>: kills shortcut <description> — fair because <reason>
