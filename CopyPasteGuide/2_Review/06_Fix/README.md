# STEP 06 — Targeted Fixes
#
# You are here because Step 05 returned REVISION REQUESTED.
#
# WHICH FILE TO USE:
#   P score < 3/3 → Fix_Problem.md   (problem.md repair)
#   T score < 3/3 → Fix_Tests.md     (test.patch repair)
#   S score < 3/3 → Fix_Solution.md  (solution.patch repair)
#
# IF MULTIPLE AXES FAIL — fix in this order:
#   1. Fix_Solution.md first  (solution bugs can mask test failures)
#   2. Fix_Tests.md second    (test gaps need clean solution to re-verify)
#   3. Fix_Problem.md last    (description is fastest to fix)
#
# AFTER EACH FIX:
#   - For P fix only: no Docker needed — go directly back to Step 05.
#   - For S or T fix: run the FULL Docker matrix below before going back to Step 05.
#
#   HOW TO RUN DOCKER:
#     If your LLM has shell/terminal access → paste Docker_Matrix.md
#     If not → run the commands manually from the section below
#
# MAXIMUM CYCLES:
#   Run Step 05 → Step 06 loop a maximum of 3 times.
#   Still failing after 3 cycles → go back to Step 02 and pick a different seed.

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
