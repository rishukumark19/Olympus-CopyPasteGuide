You are solving a real GitHub issue for a production Go repository. Execute immediately. Do not ask for clarification. Implement the solution and output your diff.

You have access to:
1. The repository at the pinned commit (checkout and explore freely)
2. This problem description only

You do NOT have access to any tests or reference solution. Write your solution based ONLY on the repository code and this description.

### Repository
URL: (read from repo_url.txt in your current working directory)
Commit: (read from commit.txt in your current working directory)

Checkout: git checkout (read from commit.txt in your current working directory)

### Problem Description

(read problem.md from your current working directory)

### Instructions

1. Explore the repository to understand the codebase, existing patterns, types, and conventions.
2. Implement the required behavior as described.
3. Follow existing code patterns exactly — no new patterns, no unnecessary abstractions.
4. Do not add tests. Do not modify test files.
5. Do not add comments explaining what you did.
6. Your implementation should be the minimum code needed to correctly implement the described behavior.

### Output

Save your implementation patch directly into `solution.patch` in the current working directory. Do NOT just print the diff in the chat.
After saving the file, write a short summary message.

---
> **STALENESS RULE (read before editing anything):** Once an agent run finishes, editing ANY submission content (problem.md, test.patch, solution.patch, Dockerfile) marks those runs stale — they stop counting and the tokens are gone. Read ALL run results first, decide ALL fixes, then edit once. Never trigger new runs while a batch is still in flight.

EVALUATION SHEET — fill this in after all 5 runs complete (human note, not for LLM):

Agent 1: PASS / FAIL / PARTIAL — LOC: ___ — Files: ___ — Notes: ___
Agent 2: PASS / FAIL / PARTIAL — LOC: ___ — Files: ___ — Notes: ___
Agent 3: PASS / FAIL / PARTIAL — LOC: ___ — Files: ___ — Notes: ___
Agent 4: PASS / FAIL / PARTIAL — LOC: ___ — Files: ___ — Notes: ___
Agent 5: PASS / FAIL / PARTIAL — LOC: ___ — Files: ___ — Notes: ___

Pass rate: ___ / 5

For each PASSING solution, verify it is legitimate:
  - Did not hardcode test values? YES/NO
  - Did not read/stub the test.patch? YES/NO
  - LOC is comparable to reference? YES/NO
  - Touches the right subsystems? YES/NO
  → Legitimate pass? YES/NO

Decision:
  0/5, agents were confused about what to build → `4_Extras/11_Tune/Easier.md`
  0/5, agents understood the task but missed behaviors → `4_Extras/11_Tune/After_Runs.md`
  1-2/5 legitimate passes → Go to `3_Validate/10_Submit/CHECKLIST.md`
  3-4/5 with shortcuts found → `4_Extras/11_Tune/Harder.md` → re-run Docker → re-run Step 08 → re-run Step 09
  3-4/5 legitimate passes → Go to `3_Validate/10_Submit/CHECKLIST.md` (acceptable)
  5/5 → TOO EASY → `4_Extras/11_Tune/Harder.md` or expand scope via `4_Extras/11_Tune/Expand.md`
