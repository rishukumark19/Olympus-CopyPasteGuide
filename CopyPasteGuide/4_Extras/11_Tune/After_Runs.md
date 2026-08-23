You are analyzing blind-test agent solutions to add surgical clarity to problem.md. Execute immediately. Do not ask questions.

**Use this ONLY when:** Step 09 returned 0 legitimate passes and agents clearly understood the task but missed specific behaviors (not when they were confused about what to build — use Easier.md for that). Goal: add the minimum fair behavioral clarity so that 1-2 out of 5 agents can produce a complete solution.

**Hard constraint:** Every hint added must be behavioral (observable input/output/effect), not implementation (no file paths, helper names, or algorithmic hints). The solver must still figure out HOW — you are only clarifying WHAT.

### Artifacts (read from your current working directory):
- problem.md
- test.patch (you will NOT expose test specifics to agents — read to understand what's missing)
- solution.patch (read to understand the full expected behavior)
- repo_url.txt

### Agent Run Results:
[PASTE the agent summaries here: which tests each agent passed/failed, what they implemented, their LOC, which subsystems they touched, and any error output or trajectories you have]

---

### Step 1: Triage Agent Failures

For each agent that failed, classify the failure:

**Type A — Missing behavior:** Agent built the right core but didn't implement one or more required behaviors. The behavior IS testable but the agent didn't attempt it.
→ Root cause: problem.md didn't make that behavior visible enough as a requirement.

**Type B — Wrong implementation path:** Agent implemented something different — a different API, a different abstraction, a different output format.
→ Root cause: problem.md was ambiguous about what the observable outcome should be.

**Type C — Incorrect partial:** Agent partially implemented but their partial solution is wrong in a way that fails tests (not just incomplete).
→ Root cause: problem.md left a behavioral edge case unspecified that the agent guessed wrong.

**Type D — Scope confusion:** Agent didn't know the full scope of what was required.
→ Root cause: problem.md didn't describe all the required behaviors clearly.

Report: for each agent, which type(s) of failure and the specific gap.

---

### Step 2: Find the Shared Gaps

Across all agents that partially passed, find:
1. Which specific behaviors did MOST agents miss?
2. Which behaviors did the 1-2 strongest agents implement correctly?
3. What did agents that came closest do right? What did they still miss?

The hint target is: behaviors that agents demonstrably tried but got wrong, or behaviors they didn't attempt because nothing in problem.md pointed at them.

Do NOT target behaviors that agents simply implemented differently but correctly — those are flexible behaviors and adding a hint would unfairly constrain valid implementations.

---

### Step 3: Design the Surgical Hints

For each identified gap, design the minimum behavioral clause to add to problem.md:

Rules for every hint:
- Describes an observable behavior: input → expected output/effect/error. Never an algorithm or internal file.
- The solver could have known this FROM a well-written problem.md — this is clarity, not a new requirement.
- Does not disturb agents that are already passing (if any): a passing agent's implementation must still satisfy the new clause.
- Does not reveal test-specific values, fixture names, or test structure.
- Does not leak the solution approach (P6).
- Fits within the 80-160 word limit — trim elsewhere if needed.

For each proposed hint:
1. State the gap: which behavior was missing/ambiguous
2. Write the proposed clause (plain English, maintainer prose)
3. Verify: would an agent that read only this clause implement the right thing?
4. Verify: does this clause rule out the wrong implementations agents produced?
5. Verify: would a currently-passing agent (if any) still pass?

---

### Step 4: Check Fairness Before Applying

Before touching problem.md, run this fairness gate on each proposed clause:
- Is this behavior discoverable from the pinned repo? (If yes — the description was the gap, this is fair to add)
- Is this implementation-neutral? (If it implies a specific file, function, or data structure — remove that part)
- Is it already implied by problem.md and the agents just missed it? (If yes — rephrase the existing clause instead of adding a new one)

Remove any clause that fails these checks.

---

### Output

Agent failure analysis:
| Agent | Type | Specific gap |
|-------|------|-------------|
| Agent 1 | A/B/C/D | <what they missed> |
...

Shared gaps (most agents missed):
1. <behavior> — missed by <N> agents

Proposed hints:
1. Clause: "<exact prose to add>"
   Gap closed: <which failure type this addresses>
   Fairness: discoverable from repo? yes/no — implementation-neutral? yes/no
   Effect on passing agents: none / <describe if any>

Revised problem.md:
Save the updated text directly to problem.md in your current working directory. Do NOT print it in chat.

Confirm:
- Word count: ___
- Clauses added: ___
- Clauses rephrased: ___
- Net effect: agents should now have enough to build a complete solution; 1-2 out of 5 expected to pass
