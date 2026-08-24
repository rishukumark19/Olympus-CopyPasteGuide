# Olympus Challenge — Copy-Paste Pipeline

One file per step. Open it, paste into your LLM, read the output, act on it.

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
3. The LLM will: automatically clone the repo & checkout commit → check for PR overlap → write `problem.md` → write `test.patch` → write `Dockerfile`
4. It creates a `challenge/<slug>/` folder with all files

**After Session 3 ends — run Docker yourself:**
```bash
git checkout <COMMIT_HASH>
git apply challenge/<slug>/test.patch
docker build -t challenge-test .
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base   # MUST exit 0
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new    # MUST exit non-zero
```
If base fails or new passes → the LLM made a mistake. Tell it what happened and fix it.
If both correct → continue to Session 4.

### Session 4 — Check problem description (03B)

1. Paste `1_Build/03B_Check_Problem/PROMPT.md` — the LLM auto-navigates to the challenge folder and reads `problem.md`
2. If PASS → go to Session 5
3. If ISSUES FOUND → fix `problem.md` in the same session, then re-run 03B

### Session 5 — Build the solution (04) ← NEW CHAT, STRICT RULE

> **MUST be a brand new chat session.** The solution agent cannot have seen test.patch.

1. Open a **brand new chat window** (no memory of Sessions 3-4)
2. Paste `1_Build/04_Build_Solution/PROMPT.md` — the LLM auto-navigates to the challenge folder and reads `problem.md` + repo info
3. The LLM writes `solution.patch` from `problem.md` + repo only

**After Session 5 ends — run Docker again:**
```bash
git apply challenge/<slug>/solution.patch
docker build -t challenge-test .
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base   # MUST exit 0
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new    # MUST exit 0
```
If new still fails → paste `2_Review/06_Fix/4_Docker_Matrix.md` to diagnose → fix with `1_Fix_Solution.md`

### Session 6 — Full review (05) — loop up to 3 times

1. Paste `2_Review/05_Review/PROMPT.md`
2. Read the verdict:
   - **ACCEPTED (3/3 all)** → go to Session 7
   - **REVISION REQUESTED** → fix with the relevant `06_Fix/` file → re-run Docker if T or S changed → re-run 05
   - Still failing after **3 loops** → back to Session 2, pick a different seed

### Session 7 — Coverage + Anti-Shortcut

1. Paste `2_Review/07_Coverage/PROMPT.md` → must be CLEAN
2. If GAPS → fix with `2_Review/06_Fix/2_Fix_Tests.md` → re-run Docker → re-run 07
3. Paste `2_Review/08_Anti_Shortcut/PROMPT.md` → must PASS
4. If NEEDS CHANGES → fix → re-run Docker → re-run 08

### Session 8 — Blind testing (09) ← 5 separate fresh chats

> Each agent gets ONLY `problem.md` and the repo. Nothing else.

1. Open a fresh chat × 5 (or use 5 different models)
2. For each: paste `problem.md` content + repo URL + commit hash
3. Each agent writes its own solution
4. Run each solution through Docker with your `test.patch`
5. Count how many pass legitimately (no hardcoding, no shortcutting)
6. See the Routing table below for what to do based on the pass rate

### Session 9 — Final checklist + submit

1. Paste `3_Validate/10_Submit/CHECKLIST.md`
2. All GREEN → submit
3. Any RED → fix it first

---

## Folder Structure

```
CopyPasteGuide/
│
├── 1_Build/                      ← steps 01–04B: creating and verifying things
│   ├── 01_Find_Repo/PROMPT.md
│   ├── 02_Pick_Seed/PROMPT.md
│   ├── 03A_Build/PROMPT.md
│   ├── 03B_Check_Problem/PROMPT.md
│   ├── 03C_Docker_Matrix/PROMPT.md   ← Docker check after 03A (test.patch only)
│   ├── 04_Build_Solution/PROMPT.md
│   └── 04B_Docker_Matrix/README.md  ← Docker check after 04 (both patches) — paste 06_Fix/Docker_Matrix.md
│
├── 2_Review/                     ← steps 05–08: auditing things
│   ├── 05_Review/PROMPT.md
│   ├── 06_Fix/1_Fix_Solution.md  2_Fix_Tests.md  3_Fix_Problem.md  4_Docker_Matrix.md
│   ├── 07_Coverage/PROMPT.md
│   └── 08_Anti_Shortcut/PROMPT.md
│
├── 3_Validate/                   ← steps 09–10: proving it works
│   ├── 09_Blind_Test/PROMPT.md
│   └── 10_Submit/CHECKLIST.md
│
└── 4_Extras/                     ← situational, use only if needed
    ├── 11_Tune/Harder.md  Easier.md  After_Runs.md  Expand.md  Reduce.md
    └── 12_Re_Check/Check_Problem.md  Check_Tests.md  Check_Solution.md  Full.md  Local.md  Quick_PTS.md
```

---

## Pipeline Flow

| # | Path | What It Does |
|---|------|-------------|
| 01 | `1_Build/01_Find_Repo/` | Find 3-5 Go repo candidates |
| 02 | `1_Build/02_Pick_Seed/` | Find the best behavioral seed in the chosen repo |
| 03A | `1_Build/03A_Build/` | Write problem.md + test.patch (NO solution yet) |
| 03B | `1_Build/03B_Check_Problem/` | Audit problem.md quality — must PASS before Step 03C |
| 03C | `1_Build/03C_Docker_Matrix/` | Docker check — test.patch only: base=PASS, new=FAIL |
| 04 | `1_Build/04_Build_Solution/` | Write solution.patch in a **fresh session** |
| 04B | `1_Build/04B_Docker_Matrix/` | Docker check — both patches: all 4 must PASS |
| 05 | `2_Review/05_Review/` | Full review — must get ACCEPTED 3/3 before moving on |
| 06 | `2_Review/06_Fix/` | Repair for P / T / S axis failures from Step 05 |
| 07 | `2_Review/07_Coverage/` | Find missing behavioral tests — must be CLEAN |
| 08 | `2_Review/08_Anti_Shortcut/` | Hunt false positives / surviving mutations — must PASS |
| 09 | `3_Validate/09_Blind_Test/` | 5 blind agent runs — target 1-4/5 legitimate passes |
| 10 | `3_Validate/10_Submit/` | All-green checklist before submitting |
| 11 | `4_Extras/11_Tune/` | Adjust difficulty / LOC after Step 09 reveals issues |
| 12 | `4_Extras/12_Re_Check/` | Re-check existing artifacts or audit with a cheap/small model |

---

## Routing — What to Do After Each Output

| Step output | Action |
|-------------|--------|
| 01 → got 3-5 repos | Pick 1 → run 02 |
| 01 → 0 repos | Search different category → re-run 01 |
| 02 → 1+ READY | Take top READY seed → run 03A |
| 02 → all DISQUALIFIED | New repo → back to 01 |
| 03A → files created | → run 03B |
| 03B → PASS | → run 03C (Docker Matrix) |
| 03B → ISSUES FOUND | Fix problem.md → re-run 03B |
| 03C → base=PASS, new=FAIL | → run 04 (new session!) |
| 03C → base=FAIL or new=PASS | Fix test.patch → re-run 03C |
| 03C → 03A aborted (seed too small) | New seed → back to 02 |
| 04 → solution.patch created | → run 04B (Docker Matrix, both patches) |
| 04B → all PASS | → run 05 |
| 04B → any FAIL | Fix solution.patch → re-run 04B |
| 05 → ACCEPTED (all 3/3) | → run 07 |
| 05 → P findings | 3_Fix_Problem → re-run 05 |
| 05 → T findings | 2_Fix_Tests → 4_Docker_Matrix → re-run 05 |
| 05 → S findings | 1_Fix_Solution → 4_Docker_Matrix → re-run 05 |
| 05 → still failing after 3 cycles | New seed → back to 02 |
| 07 → CLEAN | → run 08 |
| 07 → GAPS FOUND | Fix_Tests → re-run 07 |
| 08 → PASS | → run 09 |
| 08 → NEEDS CHANGES | Fix_Tests or Fix_Problem → re-run 08 |
| 09 → 0/5, confused agents | → Easier.md |
| 09 → 0/5, agents understood task | → After_Runs.md |
| 09 → 1-2/5 legitimate | → 10_Submit |
| 09 → 3-4/5 with shortcuts | → Harder.md → Docker → 08 → re-run 09 |
| 09 → 3-4/5 legitimate | → 10_Submit (acceptable) |
| 09 → 5/5 | TOO EASY → Harder.md or Expand.md |

---

## Staleness Rules — Know What to Re-Run After Any Edit

| If you change... | Steps now STALE | Minimum re-run |
|------------------|-----------------|----------------|
| `problem.md` only | 03B, 05 | → 03B → 05 |
| `test.patch` | 05, 07, 08, 09 | → Docker (test only) → 05 → 07 → 08 |
| `solution.patch` | 05, 09 | → Docker (both patches) → 05 |
| Both patches | 05, 07, 08, 09 | → Docker → 05 → 07 → 08 |
| Any artifact after Step 09 | Step 09 rollouts | → re-run 09 before submitting |

> **STALENESS RULE:** Once a platform rollout finishes, editing ANY submission content marks it stale — it stops counting and those tokens are gone. Read ALL results first, decide ALL fixes, then edit ONCE.

> **Quick re-check without rebuilding?** Go to `4_Extras/12_Re_Check/` — run `Local.md` first (terminal, no LLM), then `Check_Problem` / `Check_Tests` / `Check_Solution`.

---

## Difficulty Tuning — When Step 09 Gives a Bad Result

| Step 09 result | What it means | File to use |
|---|---|---|
| 5/5 pass | Too easy — shortcuts exist | `4_Extras/11_Tune/Harder.md` |
| 3-4/5, shortcuts found | Partial shortcuts — tighten tests | `4_Extras/11_Tune/Harder.md` |
| 0/5 pass, confused agents | Too hard or unfair | `4_Extras/11_Tune/Easier.md` |
| 0/5 pass, agents understood task | Missing behavioral clarity | `4_Extras/11_Tune/After_Runs.md` |
| Some pass, others miss behaviors | Gaps in description | `4_Extras/11_Tune/After_Runs.md` |
| Critic flags LOC < 250 | Scope too thin | `4_Extras/11_Tune/Expand.md` |
| Critic flags LOC > 600 | Padding in solution | `4_Extras/11_Tune/Reduce.md` |
| Want a second opinion | Re-audit existing artifacts | `4_Extras/12_Re_Check/Full.md` |
| Fast mid-build sanity check | Recursive P/T/S fix loop | `4_Extras/12_Re_Check/Quick_PTS.md` |
| Using cheap model | Atomic checks, no hallucination risk | `4_Extras/12_Re_Check/Check_*.md` |

---

## The Iron Rules

1. **Never skip a step gate.** Each gate is a hard requirement.
2. **test.patch is LOCKED after Step 04 Docker confirms base=PASS, new=PASS.** Any test change after this restarts from Step 04 Docker re-verification.
3. **Run Docker locally after Steps 03A and 04.** Do not skip even once — the LLM cannot run Docker, only you can verify it.
4. **Step 05 max 3 cycles.** Still failing after 3 → back to Step 02 with a different seed.
5. **Step 06 fix order when multiple axes fail: S first → T → P.** Solution bugs can mask test failures.
6. **Fix ALL findings in one pass before re-running.** Running the same prompt after every small tweak wastes tokens. Read everything, fix everything, then re-run once.
7. **At least 1 agent must solve before you can submit.** 0/5 is not submittable — iterate before spending more runs.
8. **Read failures before editing.** An agent failing because the task is hard is what you want. An agent failing because of ambiguity or an unfair test is a problem to fix.

---

## Common Rejection Reasons (from platform review)

- **Existing PR already implements it** — the #1 rejection reason. Check open, merged, AND closed PRs. Check GitHub Discussions too — that's where design rulings and declined features live.
- **Problem description is AI-written** — em-dashes, hard line-wrapping, "Additionally", lists of requirements instead of prose.
- **Tests are too permissive** — a partial or incorrect implementation passes. Failing because of ambiguity or a missing edge case, not because the task is hard.
- **Solution has no test** — every hunk in solution.patch must have a test in test.patch that forces it.
- **LOC is below the bar** — CRUD wiring, config plumbing, and single-layer changes rarely hit 300+ meaningful LOC. Think cross-cutting.
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
| Review | 05 | Medium — thorough critic |
| Review | 06 | Low — targeted fix only |
| Review | 07 | Medium — coverage matrix |
| Review | 08 | HIGH — runs actual mutations |
| Validate | 09 | VERY HIGH — 5 full agent runs |
| Validate | 10 | Very Low — checklist only |
| Extras | 11–12 | Varies — only if needed |

> Clear checks before agent runs. Work through 01–08 until everything passes, then run 09. Starting agent runs while a check is failing wastes the most expensive tokens in the pipeline.

---

## Local Docker Commands (copy-ready)

```bash
# After 03A — tests only, no solution yet
git checkout <COMMIT_HASH>
git apply test.patch
docker build -t challenge-test .
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base   # MUST PASS
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new    # MUST FAIL

# After 04 — both patches applied
git apply solution.patch
docker build -t challenge-test .
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base   # MUST PASS
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new    # MUST PASS

# Clean start (use this before any final re-verification)
git stash
git checkout <COMMIT_HASH>
git apply test.patch && git apply solution.patch
docker build -t challenge-test .
```
