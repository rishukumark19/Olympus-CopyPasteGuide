You are fixing a Go Olympus challenge that returned 0/5 passes in blind testing. Execute immediately. Do not explain, do not ask questions. Identify the root cause and make the minimum fix to make it fair and solvable.

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
[PASTE what agents tried to implement, where they got stuck, error messages, what they misunderstood]

---

### Diagnosis Step (do this first)

Read the agent failure logs carefully. Classify each failure:

Type A — Misunderstood the task:
- Agent implemented something different from what was asked
- Agent asked clarifying questions or expressed uncertainty about what to build
- Agent built the right thing but in a completely wrong location or API
→ Root cause: problem.md is ambiguous. Fix: clarify the observable behavior minimum needed.

Type B — Understood but couldn't implement:
- Agent correctly identified what to build and where
- Agent ran out of time/messages or the scope was too large
- Agent made partial progress but couldn't finish
→ Root cause: scope too large. Fix: reduce behavioral scope or split the task.

Type C — Environment broken:
- Agent failed due to test.sh errors, Docker issues, patch conflicts, or missing dependencies
- Agent never got to implementation
→ Root cause: infrastructure problem. Fix: repair test.sh / Dockerfile / patch hygiene.

Report: which type(s) are present, with evidence from the failure logs.

---

### Fix Rules

For Type A (ambiguity) — Minimal problem.md clarification:
- Add ONLY the missing observable behavior detail that agents demonstrably misunderstood
- Do not name private helpers, internal files, or implementation approach
- Do not add implementation hints — add behavioral contract clarity
- Every added clause must correspond to an existing test (do not expand the contract)
- Word count must stay 80-160 words

For Type B (scope too large) — Scope reduction:
- Identify which behaviors are the "core" of the task vs. "expansion" behaviors
- Remove the expansion behaviors from test.patch (and the corresponding solution.patch hunks)
- Update problem.md to remove any clauses that are no longer tested
- The remaining scope must still cross 2+ subsystems and 250+ meaningful LOC
- If scope cannot be reduced to a fair level, recommend going back to 1_Build/02_Pick_Seed for a different seed

For Type C (environment) — Infrastructure fix:
- Fix test.sh argument handling, base/new modes, JUnit output
- Fix Dockerfile if needed (check go mod download, offline execution)
- Fix patch conflicts or stale hunks
- No problem.md or behavioral changes needed

Fairness check after any fix:
- Every test still fails on base and passes with solution (T1)
- Problem.md still has no implementation hints (P6)
- Agents can derive every tested behavior from problem.md + pinned repo alone (P3, T5)

### Output

Respond in this order — no other text before the diagnosis:

Diagnosis: <Type A/B/C and evidence>

Recommended fix: <description of minimum change>

Revised artifacts (only the ones that changed):

**Revised problem.md** (if Type A):
Save the revised text directly to the file. Do not print it in chat.

**Revised test.patch** (if Type B or C):
Save the revised patch directly to the file. Do not print it in chat.

**Revised solution.patch** (if Type B):
Save the revised patch directly to the file. Do not print it in chat.

Changes made:
- <change 1> — addresses <failure type>
- <change 2> — addresses <failure type>
