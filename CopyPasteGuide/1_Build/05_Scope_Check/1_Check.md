## Go Challenge Evaluator

Read all files in the challenge folder, audit every quality gate, measure scope, and give a verdict. Do NOT fix files — report only.

**CHALLENGE_FOLDER:** `[FILL IN — e.g. D:\SHIPD ai\challanges\challange2\challenge\pull-consumer-priority-groups]`

---

### 1. Read Files

In this order: `repo_url.txt` → `problem.md` → `test.patch` → `solution.patch` → `Dockerfile`

---

### 2. Problem Audit (P1–P7)

Report ✅ PASS / ⚠️ WEAK / ❌ FAIL + one-line evidence for each.

- **P1** Aligns with repo philosophy
- **P2** Not already fixed in an open/closed/merged PR
- **P3** Self-contained — solvable from repo + `problem.md` alone
- **P4** Clear and unambiguous — no guessing, no leaked hints
- **P5** Verifiable — success is objectively testable
- **P6** Non-prescriptive — no internal file names, types, or mechanisms leaked
- **P7** Not a duplicate of an existing submission

---

### 3. Test Audit (T1–T7)

Report ✅ PASS / ⚠️ WEAK / ❌ FAIL + one-line evidence for each.

- **T1** Every new test fails on base commit, passes with solution
- **T2** Deterministic — no timing, randomness, or goroutine-ordering dependence
- **T3** Strong — shortcut or inaccurate solutions cannot pass
- **T4** Extensive — covers all behaviors, edge cases, and negative paths (errors, rollbacks, rejections)
- **T5** No undiscoverable behavior — every assertion maps to `problem.md` or the repo
- **T6** Fully offline — no network, no runtime installs
- **T7** No over-pinning — checks behavior, not exact error text or formatting

---

### 4. Solution Audit (S1–S4)

- **S1** Meets ALL requirements — every clause in `problem.md` is implemented
- **S2** No regressions — follows existing code patterns, does not break upstream tests
- **S3** No unrelated changes — nothing outside the feature's scope is touched
- **S4** No AI slop — no weird comments, unexplained defensive code, or alien patterns

---

### 5. Scope Measurement

Count from `solution.patch` (meaningful LOC excludes blanks, comments, braces, imports, tests, generated files):

```
Meaningful LOC:   
Non-test files:   (list)
Subsystems:       (list)
problem.md words: 
```

| Threshold | Required | Measured | Status |
|---|---|---|---|
| Meaningful LOC | 250+ | | |
| Non-test files | 5+ | | |
| Subsystems | 2+ | | |
| problem.md words | 80–160 | | |
| Agent effort (est.) | 60+ messages | | |

---

### 6. Conflict Resolution Rule

Before routing, apply this rule to every finding:
**Shipd rules override finding suggestions.** If a finding says "add X to problem.md" but X is an internal type, file name, JSON schema field, or error code — that would violate P6 (non-prescriptive). The correct fix is to **relax the test** to check the observable behavior, not to add the internal detail to `problem.md`. Apply this judgment before routing.

---

### 7. Verdict

Give exactly one:

**✅ READY** — all gates pass, thresholds met.
> **Next:** `3_Validate/11_Submit/CHECKLIST.md`

**⚠️ NEEDS REVISION** — list each failing gate with a concrete repair instruction.
> **Next:** Route findings to `2_Review/07_Fix/` — use in this order:
> 1. `1_Fix_Solution.md` for any S* findings
> 2. `2_Fix_Tests.md` for any T* findings
> 3. `3_Fix_Problem.md` for any P* findings
>
> Then run `4_Docker_Matrix.md` to verify. Then re-run this prompt.
>
> **Note:** If LOC is still under 250 after revision, also run `05_Scope_Check/2_Expand.md` after.

**🔴 NEEDS EXPANSION** — valid but below scope. State the LOC shortfall and suggest natural same-workflow additions (persistence, validation, serialization, CLI variants, recovery, lifecycle behavior). Do NOT suggest API inventory, schema plumbing, or unrelated parity.
> **Next:** `1_Build/05_Scope_Check/2_Expand.md` → expand → re-run Step 04 in fresh chat → re-run this prompt.

---

### Output Format

```
VERDICT: <READY TO SUBMIT / NEEDS REVISION / NEEDS EXPANSION>

PROBLEM:   P1 ✅  P2 ✅  P3 ⚠️  P4 ✅  P5 ✅  P6 ✅  P7 ✅
TESTS:     T1 ✅  T2 ✅  T3 ⚠️  T4 ❌  T5 ✅  T6 ✅  T7 ✅
SOLUTION:  S1 ✅  S2 ✅  S3 ✅  S4 ✅

SCOPE:
  Meaningful LOC: X  (PASS/FAIL)
  Non-test files: X  (PASS/FAIL)
  Subsystems: [list]  (PASS/FAIL)
  problem.md words: X  (PASS/FAIL)

FINDINGS:
  1. <gate> — <evidence> — <exact repair or expansion suggestion>
  2. ...

NEXT STEP: <07_Fix + Docker → re-run / 05_Scope_Check/2_Expand → re-run / 11_Submit>
```
