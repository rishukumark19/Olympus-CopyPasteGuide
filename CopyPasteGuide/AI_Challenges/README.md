# AI Challenge — Solving Loop

This folder is the entry point when you have a challenge and want to test or iterate on it.

---

## Where Each Step Lives in the Main Workflow

```
challenge/<slug>/
├── problem.md       ← written by 1_Build/03A_Build/PROMPT.md
├── repo_url.txt     ← written by 1_Build/03A_Build/PROMPT.md
├── Dockerfile       ← written by 1_Build/03A_Build/PROMPT.md
├── test.patch       ← written by 1_Build/03A_Build/PROMPT.md
└── solution.patch   ← written by 1_Build/04_Build_Solution/PROMPT.md
```

---

## The Full Loop

```
1. BUILD problem + tests
   → 1_Build/03A_Build/PROMPT.md          (new chat, best model)
   → 1_Build/03B_Check_Problem/PROMPT.md  (quality check)

2. SOLVE — reference solution
   → 1_Build/04_Build_Solution/PROMPT.md  (brand new chat — must NOT have seen test.patch)
   → 1_Build/04B_Docker_Matrix/README.md  (verify both patches pass in Docker)

3. SCOPE CHECK — run BEFORE full review
   → 1_Build/05_Scope_Check/1_Check.md
        ✅ READY      → go to step 4 (Review)
        ⚠️ REVISION   → 2_Review/07_Fix/ → Docker → re-run 1_Check.md
        🔴 EXPANSION  → 4_Extras/11_Tune/Expand.md → re-run step 2 (new chat) → re-run 1_Check.md

4. REVIEW
   → 2_Review/06_Review/PROMPT.md         (full P/T/S audit — up to 3 cycles)
   → 2_Review/07_Fix/ (1, 2, 3)           (fix S → T → P in that order)
   → 2_Review/08_Coverage/PROMPT.md
   → 2_Review/09_Anti_Shortcut/PROMPT.md

5. BLIND TEST
   → 3_Validate/10_Blind_Test/PROMPT.md   (5 fresh isolated chats)

6. SUBMIT
   → 3_Validate/11_Submit/CHECKLIST.md
```

---

## Key Rules

| Rule | Why |
|---|---|
| Step 2 (solve) MUST be a brand-new chat | The agent must not have seen `test.patch` |
| Blind test agents get ONLY `problem.md` + repo | Mirrors what real platform agents see |
| Fix order when multiple axes fail: S → T → P | Solution bugs can mask test failures |
| Always run Docker after any patch change | LLM cannot verify Docker — only you can |
| `solution.patch` = implementation only | No tests, no Dockerfile, no docs, no unrelated changes |

---

## Staleness — What to Re-Run After Any Edit

| If you change... | Minimum re-run |
|---|---|
| `problem.md` only | → 03B → 05 → 06 |
| `test.patch` | → Docker (test only) → 05 → 06 → 08 → 09 |
| `solution.patch` | → Docker (both patches) → 05 → 06 |
| Both patches | → Docker → 05 → 06 → 08 → 09 |

---

## Quick Fix Reference

| What's broken | File to use |
|---|---|
| Tests don't compile | `2_Review/07_Fix/5_Fix_Compile.md` |
| Solution wrong / incomplete | `2_Review/07_Fix/1_Fix_Solution.md` |
| Tests too weak / missing | `2_Review/07_Fix/2_Fix_Tests.md` |
| problem.md has issues | `2_Review/07_Fix/3_Fix_Problem.md` |
| Docker failing unexpectedly | `2_Review/07_Fix/4_Docker_Matrix.md` |
| Challenge too small (LOC) | `4_Extras/11_Tune/Expand.md` |
| Challenge too hard / unfair | `4_Extras/11_Tune/Easier.md` |
| Challenge too easy (shortcuts) | `4_Extras/11_Tune/Harder.md` |
