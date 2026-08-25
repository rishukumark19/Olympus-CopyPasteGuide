# STEP 06 — Targeted Fixes
#
# You are here because Step 05 returned REVISION REQUESTED.
#
# WHICH FILE TO USE:
#   P score < 3/3 → 3_Fix_Problem.md   (problem.md repair)
#   T score < 3/3 → 2_Fix_Tests.md     (test.patch repair)
#   S score < 3/3 → 1_Fix_Solution.md  (solution.patch repair)
#
# IF MULTIPLE AXES FAIL — fix in this order:
#   1. 1_Fix_Solution.md first  (solution bugs can mask test failures)
#   2. 2_Fix_Tests.md second    (test gaps need clean solution to re-verify)
#   3. 3_Fix_Problem.md last    (description is fastest to fix)
#
# THE WORKFLOW LOOP:
#   1. Run `05_Review/PROMPT.md` in a NEW chat window/session. It will secretly save its findings to `review_findings.md`.
#   2. Run `06_Fix/1, 2, 3` in a NEW chat window/session. They will automatically read `review_findings.md` and apply the fixes.
#   3. Run `06_Fix/4_Docker_Matrix.md`. You MUST do this to verify the fixes actually compile and pass.
#   4. If the Matrix fails, use `5_Fix_Compile.md` or manually fix the code.
#   5. If the Matrix passes, go back to Step 1 (`05_Review/PROMPT.md` in a NEW window/session) to get a fresh review!
#
# MAXIMUM CYCLES:
#   Run this Review -> Fix -> Matrix loop a maximum of 3 times.
#   If you are stuck in a loop after 3 cycles, stop the automated prompts and manually intervene, or go back to Step 02.

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
