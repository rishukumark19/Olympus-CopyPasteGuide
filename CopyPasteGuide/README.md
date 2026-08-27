# Olympus Challenge — Copy-Paste Pipeline

One file per step. Open it, paste into your LLM, read the output, act on it.

---

## ⚡ Lean Submit Path (Minimum Viable Submission)

Time-pressured? This is the shortest path to a valid submission:

```
01 Find Repo  →  02 Pick Seed  →  03A Build (Docker BASE=PASS, NEW=FAIL)
→  03B Check Problem (PASS)  →  04 Build Solution (new chat)
→  04B Docker verify (BASE=PASS, NEW=PASS)  →  11 Submit Checklist  →  SUBMIT
```

Steps 05, 06, 07, 08, 09, 10 improve quality but are **not required to submit**.
Run them when you have time or if the challenge feels weak.

> **Platform auto-reviewer:** ShipD AI runs an automated score on your submission after you submit. Do NOT chase that score — it is feedback, not a gate. `Docker NEW=PASS` is the real requirement. Focus on getting that right, then submit.

---

## Complete Walkthrough — How to Do a Full Challenge From Scratch

> Read this once before starting. Everything else in this file is reference.

### Session 1 — Find a repo

1. Open a new LLM chat (any model works here — this is cheap)
2. Set your working directory to wherever you store repos
3. Paste `1_Build/01_Find_Repo/PROMPT.md`
4. Get back 3-5 repos in the fixed format
5. If you got 0 repos → search a different Go category and re-run

### Session 2 — Pick a seed (same or new chat)

1. Pick one repo from Session 1 that looks good (or let the LLM take the top candidate)
2. Paste `1_Build/02_Pick_Seed/PROMPT.md` — the LLM inspects PRs, issues, discussions, and architecture remotely (no local clone needed)
3. Get back a list of seeds ranked READY / DISQUALIFIED
4. If all DISQUALIFIED → pick another candidate from Session 1 (zero cleanup needed)
5. Note down your top READY seed — you'll use it next

### Session 3 — Build the problem + tests (03A)

> This is the most important session. Take your time.

1. Open a new LLM chat (use your best model — Gemini/Claude/GPT)
2. Paste `1_Build/03A_Build/PROMPT.md` with your seed and repo details
3. The LLM will: clone repo → check PR overlap → write `problem.md` → write `test.patch` → write `Dockerfile` → run Docker TEST CHECK automatically
4. If Docker RESULT is FAIL → tell the LLM what failed, it fixes and re-runs
5. If Docker RESULT is PASS → continue to Session 4

### Session 4 — Check problem description (03B)

1. Paste `1_Build/03B_Check_Problem/PROMPT.md` — the LLM auto-navigates to the challenge folder and reads `problem.md`
2. If PASS → go to Session 5
3. If ISSUES FOUND → fix `problem.md` in the same session, then re-run 03B

### Session 5 — Build the solution (04) ← NEW CHAT, STRICT RULE

> **MUST be a brand new chat session.** The solution agent cannot have seen test.patch.

1. Open a **brand new chat window** (no memory of Sessions 3-4)
2. Paste `1_Build/04_Build_Solution/PROMPT.md` — the LLM auto-navigates to the challenge folder and reads `problem.md` + repo info
3. The LLM writes `solution.patch` from `problem.md` + repo only

**After Session 5 ends — run Docker TEST CHECK + SOLUTION CHECK:**
```bash
git apply challenge/<slug>/solution.patch
docker build -t challenge-test .
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base   # MUST exit 0
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new    # MUST exit 0
```
If new still fails → paste `2_Review/07_Fix/4_Docker_Matrix.md` to diagnose → fix with `1_Fix_Solution.md`

### Session 6 — Scope Check (05) ← CATCH SCOPE ISSUES BEFORE REVIEW

> Run this immediately after Docker passes. If LOC is too thin or P/T/S has fundamental failures, fix them now — before investing time in a full review cycle.

1. Open a new chat
2. Paste `1_Build/05_Scope_Check/1_Check.md` with the challenge folder path
3. Read the verdict:
   - **✅ READY** → go to Session 7 (full review)
   - **⚠️ NEEDS REVISION** → route findings to `2_Review/07_Fix/` (S → T → P order) → re-run Docker → re-run 05
   - **🔴 NEEDS EXPANSION** → `4_Extras/11_Tune/Expand.md` → re-run Step 04 in a fresh chat → re-run Docker → re-run 05

### Session 7 — Full review (06) — loop up to 3 times *(optional quality step)*

> Skip this if time-pressured. Run it when you want the highest quality score.

1. Open a **NEW chat window/session**.
2. Paste `2_Review/06_Review/PROMPT.md`. (It will review and secretly save `review_findings.md`).
3. Read the verdict:
   - **ACCEPTED (3/3 all)** → go to Session 8
   - **REVISION REQUESTED** →
     1. Open another **NEW chat window/session**.
     2. Run the relevant `07_Fix/1, 2, 3` prompts (they auto-read `review_findings.md`).
     3. Run `07_Fix/4_Docker_Matrix.md` to verify fixes compile and pass.
     4. Loop back to step 1 (New window → `06_Review`).
   - Still failing after **3 loops** → Stop automated prompts. Manually intervene, or go back to Session 2 to pick a different seed.

### Session 8 — Coverage + Anti-Shortcut *(optional quality step)*

> Run these to harden tests before blind runs. Skip if time-pressured.

1. Paste `2_Review/08_Coverage/PROMPT.md` → must be CLEAN
2. If GAPS → fix with `2_Review/07_Fix/2_Fix_Tests.md` → re-run Docker → re-run 08
3. Paste `2_Review/09_Anti_Shortcut/PROMPT.md` → must PASS
4. If NEEDS CHANGES → fix → re-run Docker → re-run 09

### Session 9 — Blind testing (10) ← 5 separate fresh chats *(optional, most expensive)*

> Each agent gets ONLY `problem.md` and the repo. Nothing else.
> **This is the most token-expensive step.** Run it only after 03A-04B Docker passes are solid and quality steps are done. Do NOT run while any check is still failing.

1. Open a fresh chat × 5 (or use 5 different models)
2. For each: paste `problem.md` content + repo URL + commit hash
3. Each agent writes its own solution
4. Run each solution through Docker with your `test.patch`
5. Count how many pass legitimately (no hardcoding, no shortcutting)
6. See the Routing table below for what to do based on the pass rate

### Session 10 — Final checklist + submit

1. Paste `3_Validate/11_Submit/CHECKLIST.md`
2. All GREEN → submit
3. Any RED → fix it first

---

## Folder Structure

```
CopyPasteGuide/
│
├── 1_Build/                      ← steps 01–05: build, verify, scope check
│   ├── 01_Find_Repo/PROMPT.md
│   ├── 02_Pick_Seed/PROMPT.md
│   ├── 03A_Build/PROMPT.md           ← builds problem.md + test.patch + Dockerfile (Docker verify built in)
│   ├── 03B_Check_Problem/PROMPT.md
│   ├── 04_Build_Solution/PROMPT.md   ← brand new chat — solution.patch only
│   ├── 04B_Docker_Matrix/README.md   ← verify both patches pass
│   └── 05_Scope_Check/               ← scope + quality gate (run before review)
│       ├── 1_Check.md
│       └── 2_Expand.md
│
├── 2_Review/                     ← steps 06–09: full audit and fix
│   ├── 06_Review/PROMPT.md
│   ├── 07_Fix/1_Fix_Solution.md  2_Fix_Tests.md  3_Fix_Problem.md  4_Docker_Matrix.md
│   ├── 08_Coverage/PROMPT.md
│   └── 09_Anti_Shortcut/PROMPT.md
│
├── 3_Validate/                   ← steps 10–11: blind test and submit
│   ├── 10_Blind_Test/PROMPT.md       ← 5 fresh isolated solving agents
│   └── 11_Submit/CHECKLIST.md
│
└── 4_Extras/                     ← situational, use only if needed
    ├── 11_Tune/Harder.md  Easier.md  After_Runs.md  Expand.md  Reduce.md
    └── 12_Re_Check/Check_Problem.md  Check_Tests.md  Check_Solution.md  Full.md  Local.md  Quick_PTS.md
```

---

## Pipeline — Step by Step Decision Map

Follow top to bottom. Each row tells you what to run and what to do with the result.

| Step | What you paste | ✅ Good output → do this next | ❌ Bad output → do this instead |
|:----:|:--------------|:------------------------------|:--------------------------------|
| **01** | `1_Build/01_Find_Repo/PROMPT.md` | 3–5 repo candidates → pick one, go to **02** | 0 repos → search different Go category, re-run **01** |
| **02** | `1_Build/02_Pick_Seed/PROMPT.md` | 1+ seed `READY` → go to **03A** | All DISQUALIFIED → try next candidate or back to **01** |
| **03A** ⭐ | `1_Build/03A_Build/PROMPT.md` + seed *(new chat, best model)* | Docker `RESULT: PASS` → go to **03B** | FAIL → tell LLM what failed, fix in same session, re-run Docker |
| **03B** | `1_Build/03B_Check_Problem/PROMPT.md` | `PASS` → go to **04** | `ISSUES FOUND` → fix `problem.md` in same session, re-run **03B** |
| **04** ⚠️ | `1_Build/04_Build_Solution/PROMPT.md` *(brand new chat — never seen test.patch)* | `solution.patch` written → go to **04B** | — |
| **04B** | `1_Build/04B_Docker_Matrix/README.md` | All checks PASS → go to **05** | Test compile error → `07_Fix/5_Fix_Compile.md` · Solution fails tests → `07_Fix/1_Fix_Solution.md` · re-run **04B** |
| **05** ⭐ *(optional)* | `1_Build/05_Scope_Check/1_Check.md` *(new chat)* | `READY` → go to **06** (or skip to **11**) | `REVISION` → `07_Fix/` + Docker → re-run **05** · `EXPANSION` → `11_Tune/Expand.md` → new chat **04** → re-run **05** |
| **06** *(optional)* | `2_Review/06_Review/PROMPT.md` *(max 3 cycles)* | `ACCEPTED 3/3` → go to **08** (or skip to **11**) | P → `3_Fix_Problem` · T → `2_Fix_Tests` + Docker · S → `1_Fix_Solution` + Docker → re-run **06** · Still failing after 3× → back to **02** |
| **08** *(optional)* | `2_Review/08_Coverage/PROMPT.md` | `CLEAN` → go to **09** (or skip to **11**) | `GAPS FOUND` → `2_Fix_Tests` → Docker → re-run **08** |
| **09** *(optional)* | `2_Review/09_Anti_Shortcut/PROMPT.md` | `PASS` → go to **10** (or skip to **11**) | `NEEDS CHANGES` → `2_Fix_Tests` or `3_Fix_Problem` → Docker → re-run **09** |
| **10** *(optional, expensive)* | `3_Validate/10_Blind_Test/PROMPT.md` × 5 fresh chats | 1–4/5 legitimate → go to **11** | 0/5 confused → `Easier.md` · 0/5 understood → `After_Runs.md` · 5/5 → `Harder.md` or `Expand.md` · 3–4/5 shortcuts → `Harder.md` → Docker → **09** → re-run **10** |
| **11** | `3_Validate/11_Submit/CHECKLIST.md` | All GREEN → **submit** | Any RED → fix it, re-run checklist |

> ⭐ Step 03A now runs Docker verification automatically at the end — no separate 03C step needed.
> ⭐ Step 05 (Scope Check) runs immediately after 04B — before review. Catches scope/LOC issues before you waste review cycles.
> ⚠️ Step 04 MUST be a brand new chat window with no memory of Sessions 03A/03B.

---

## 🔧 Fix File Quick Reference

When something breaks, pick the right fix file from `2_Review/07_Fix/`:

| Symptom | File to paste |
|:--------|:-------------|
| Tests don't compile on base (`test.patch` compile error) | `5_Fix_Compile.md` |
| Solution wrong / incomplete (Solution check fails new tests) | `1_Fix_Solution.md` |
| Tests missing / too weak / wrong (post-review) | `2_Fix_Tests.md` |
| problem.md has quality issues | `3_Fix_Problem.md` |
| Docker fails unexpectedly, need to diagnose | `4_Docker_Matrix.md` |

> **After fixing `test.patch` or `solution.patch` → always re-run Docker before re-running 05/06/08/09.**

---

## Staleness Rules — Know What to Re-Run After Any Edit

| If you change... | Steps now STALE | Minimum re-run |
|------------------|-----------------|----------------|
| `problem.md` only | 03B, 05, 06 | → 03B → 05 → 06 |
| `test.patch` | 05, 06, 08, 09, 10 | → Docker (test only) → 05 → 06 → 08 → 09 |
| `solution.patch` | 05, 06, 10 | → Docker (both patches) → 05 → 06 |
| Both patches | 05, 06, 08, 09, 10 | → Docker → 05 → 06 → 08 → 09 |
| Any artifact after Step 10 | Step 10 rollouts | → re-run 10 before submitting |

> **STALENESS RULE:** Once a platform rollout finishes, editing ANY submission content marks it stale — it stops counting and those tokens are gone. Read ALL results first, decide ALL fixes, then edit ONCE.

> **Quick re-check without rebuilding?** Go to `4_Extras/12_Re_Check/` — run `Local.md` first (terminal, no LLM), then `Check_Problem` / `Check_Tests` / `Check_Solution`.

---

## Difficulty Tuning — When Step 10 Gives a Bad Result

| Step 10 result | What it means | File to use |
|---|---|---|
| 5/5 pass | Too easy — shortcuts exist | `4_Extras/11_Tune/Harder.md` |
| 3-4/5, shortcuts found | Partial shortcuts — tighten tests | `4_Extras/11_Tune/Harder.md` |
| 0/5 pass, confused agents | Too hard or unfair | `4_Extras/11_Tune/Easier.md` |
| 0/5 pass, agents understood task | Missing behavioral clarity | `4_Extras/11_Tune/After_Runs.md` |
| Some pass, others miss behaviors | Gaps in description | `4_Extras/11_Tune/After_Runs.md` |
| Critic flags LOC < 500 | Scope too thin | `4_Extras/11_Tune/Expand.md` |
| Critic flags LOC > 700 | Padding in solution | `4_Extras/11_Tune/Reduce.md` |
| Want a second opinion | Re-audit existing artifacts | `4_Extras/12_Re_Check/Full.md` |
| Fast mid-build sanity check | Recursive P/T/S fix loop | `4_Extras/12_Re_Check/Quick_PTS.md` |
| Using cheap model | Atomic checks, no hallucination risk | `4_Extras/12_Re_Check/Check_*.md` |

---

## The Iron Rules

1. **The lean path is always valid.** Steps 05, 06, 07, 08, 09, 10 improve quality but you can submit after 03A + 03B + 04 + 04B pass Docker.
2. **test.patch is LOCKED after Step 04B Docker confirms base=PASS, new=FAIL.** Any test change after this restarts from Step 04B Docker re-verification.
3. **Run Docker locally after Steps 03A and 04.** Do not skip even once — the LLM cannot run Docker, only you can verify it.
4. **Step 06 max 3 cycles.** Still failing after 3 → back to Step 02 with a different seed.
5. **Step 07_Fix order when multiple axes fail: S first → T → P.** Solution bugs can mask test failures.
6. **Fix ALL findings in one pass before re-running.** Running the same prompt after every small tweak wastes tokens. Read everything, fix everything, then re-run once.
7. **At least 1 agent must solve before you can submit (if running Step 09).** 0/5 is not submittable — iterate before spending more runs.
8. **Read failures before editing.** An agent failing because the task is hard is what you want. An agent failing because of ambiguity or an unfair test is a problem to fix.
9. **Platform auto-reviewer ≠ submission gate.** The ShipD AI auto-reviewer runs after you submit and provides score/feedback. It is not a blocker. `Docker NEW=PASS` is the only hard gate.
10. **Iterate inside each Docker gate before moving on.** Do not skip to the next step while current Docker checks are failing. Fix, re-verify, then advance.

---

## Common Rejection Reasons (from platform review)

- **Existing PR already implements it** — the #1 rejection reason. Check open, merged, AND closed PRs. Check GitHub Discussions too — that's where design rulings and declined features live.
- **Problem description is AI-written** — em-dashes, hard line-wrapping, "Additionally", lists of requirements instead of prose.
- **Tests are too permissive** — a partial or incorrect implementation passes. Failing because of ambiguity or a missing edge case, not because the task is hard.
- **Solution has no test** — every hunk in solution.patch must have a test in test.patch that forces it.
- **LOC is below the bar** — CRUD wiring, config plumbing, and single-layer changes rarely hit 500+ meaningful LOC. Think cross-cutting.
- **Framed as an outsider request** — "Library X currently lacks Y" — write as a maintainer, not an external reporter.
- **Not self-contained** — solver needs information not in problem.md or the pinned repo.

---

## Key Principle: Tests Before Solution

**`solution.patch` must NEVER be written before `test.patch` is locked and Docker-verified.**
This is why 03A and 04 are separate steps with mandatory local Docker runs between them.

---

## Key Principle: Test-Solution Isolation (Step 04 must be a fresh session)

**The agent writing `solution.patch` must NEVER have seen `test.patch`.**

Run Step 04 in a brand new chat session. The solution agent gets only `problem.md` and the repo — exactly what a competing agent on the platform will see. This is not optional.

**Why this matters:** If the same agent writes both tests and solution in the same session, it knows what the tests check. It will silently write the solution to pass *its own* tests rather than solve the actual problem. The result: your reference solution passes, but real competing agents fail on valid implementations — a broken benchmark.

**The signal:** When Step 04 runs in isolation and doesn't pass all tests on the first try, that gap tells you whether `problem.md` is clear enough. If the isolated agent can't figure out what to implement, neither can a real competing agent. Use `Fix_Problem.md` to improve clarity, not `Fix_Tests.md` to lower the bar.

**The enforcement:**
- **Reliable:** Start a new chat window for Step 04. No memory of what tests were written.
- **Fallback:** If you run Step 04 in the same session, the prompt will tell the LLM to warn you and stop.

---

## Token Budget

| Phase | Step | Cost |
|-------|------|------|
| Build | 01–02 | Low — scouting only |
| Build | 03A | Medium — most important creative step |
| Build | 03B | Very Low — quality check only |
| Build | 04 | Medium — implementation only |
| Build | 05 | Low — scope audit only |
| Review | 06 | Medium — thorough critic |
| Review | 07 | Low — targeted fix only |
| Review | 08 | Medium — coverage matrix |
| Review | 09 | HIGH — runs actual mutations |
| Validate | 10 | VERY HIGH — 5 full agent runs |
| Validate | 11 | Very Low — checklist only |
| Extras | 11_Tune/12_Re_Check | Varies — only if needed |

> Clear checks before agent runs. Work through 01–09 until everything passes, then run 10. Starting agent runs while a check is failing wastes the most expensive tokens in the pipeline.

---

## Local Docker Commands (copy-ready)

```bash
# After 03A — TEST CHECK (tests only, no solution yet)
git checkout <COMMIT_HASH>
git apply test.patch
docker build -t challenge-test .
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base   # MUST PASS (exit 0)
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new    # MUST FAIL (exit non-zero)

# After 04 — SOLUTION CHECK (both patches applied)
git apply solution.patch
docker build -t challenge-test .
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base   # MUST PASS (exit 0)
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new    # MUST PASS (exit 0)

# Clean start (use this before any final re-verification)
git stash
git checkout <COMMIT_HASH>
git apply test.patch && git apply solution.patch
docker build -t challenge-test .
```
