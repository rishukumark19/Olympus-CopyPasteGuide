# STEP 07 — Targeted Fixes
#
# You are here because either:
#   (A) 05_Scope_Check returned NEEDS REVISION  (scope/quality gate)
#   (B) 06_Review returned REVISION REQUESTED   (full critic review)
#
# WHICH FILE TO USE:
#   P findings → 3_Fix_Problem.md   (problem.md repair)
#   T findings → 2_Fix_Tests.md     (test.patch repair)
#   S findings → 1_Fix_Solution.md  (solution.patch repair)
#
# IF MULTIPLE AXES FAIL — fix in this order:
#   1. 1_Fix_Solution.md first  (solution bugs can mask test failures)
#   2. 2_Fix_Tests.md second    (test gaps need clean solution to re-verify)
#   3. 3_Fix_Problem.md last    (description is fastest to fix)
#
# THE WORKFLOW LOOP:
#   After 05_Scope_Check NEEDS REVISION:
#     → Run the relevant fix file(s) in a NEW chat
#     → Run 4_Docker_Matrix.md to verify
#     → Re-run 1_Build/05_Scope_Check/1_Check.md
#
#   After 06_Review REVISION REQUESTED:
#     → Run `06_Review/PROMPT.md` in a NEW chat. It saves findings to `review_findings.md`.
#     → Run `07_Fix/1, 2, 3` in a NEW chat. They auto-read `review_findings.md`.
#     → Run `4_Docker_Matrix.md` to verify.
#     → Re-run `06_Review/PROMPT.md` in a NEW chat.
#
# MAXIMUM CYCLES (06_Review loop):
#   Run the Review -> Fix -> Docker loop a maximum of 3 times.
#   If stuck after 3 cycles, manually intervene or go back to Step 02.

---

## FULL Docker Matrix (run this after any S or T fix)

Run all 4 checks in order. Replace <COMMIT_HASH> with your pinned commit.

```bash
# STEP 1 — clean state, apply test.patch only (no solution)
git stash
git checkout <COMMIT_HASH>
git apply test.patch
docker build -t challenge-test .

# STEP 2 — base must exit 0 (PASS)
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base.xml base
echo "base exit code: $?"    # must be 0

# STEP 3 — new must exit non-zero (FAIL — correct, no solution yet)
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new.xml new
echo "new exit code: $?"     # must be non-zero (1 or 2)

# STEP 4 — apply solution.patch on top
git apply solution.patch
docker build -t challenge-test .

# STEP 5 — base must still exit 0 (PASS)
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/base2.xml base
echo "base exit code: $?"    # must be 0

# STEP 6 — new must now exit 0 (PASS — solution fixes the tests)
docker run --rm --network none challenge-test ./test.sh --output_path /tmp/new2.xml new
echo "new exit code: $?"     # must be 0
```

All 4 checks pass (base=0, new≠0, base=0, new=0)?
→ Go back to Step 05.

Any check fails?
→ Fix the relevant patch and re-run this matrix before going back to Step 05.
