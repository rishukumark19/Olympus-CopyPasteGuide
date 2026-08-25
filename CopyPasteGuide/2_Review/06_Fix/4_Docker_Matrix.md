You are running the Docker verification matrix for a Go Olympus challenge. Execute immediately. Do not explain. Do not ask questions. Use your shell/terminal tool to run every command below.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

### Setup

Read the pinned commit hash from `commit.txt` in your current working directory.
The challenge folder is your current working directory.

**CRITICAL RULES FOR THE LLM:**
1. DO NOT write your own scripts (`.ps1`, `.py`, `.sh`).
2. Execute the commands directly in your shell tool.
3. If you are using PowerShell, substitute `$(cat commit.txt)` with the actual hash you read from the file.
4. **DOCKER PREFERRED, NATIVE FALLBACK:** You MUST try to use Docker first. If `docker` is not installed or fails, skip the `docker build` and `docker run` commands and just run `./test.sh` natively on the host (e.g., `./test.sh --output_path /tmp/base1.xml base`).

### Phase 1 — test.patch only (no solution)

Run these commands in order (adjust syntax if using PowerShell):

```bash
git stash
git checkout <INSERT_HASH_FROM_COMMIT_TXT>
git apply test.patch

# Try Docker. If docker is missing, skip to the native fallback: `./test.sh --output_path /tmp/base1.xml base`
docker build -t challenge-verify .
docker run --rm --network none challenge-verify ./test.sh --output_path /tmp/base1.xml base
echo "PHASE1_BASE_EXIT=$?"

# Try Docker. If missing: `./test.sh --output_path /tmp/new1.xml new`
docker run --rm --network none challenge-verify ./test.sh --output_path /tmp/new1.xml new
echo "PHASE1_NEW_EXIT=$?"
```

Expected:
- PHASE1_BASE_EXIT=0   (upstream tests pass — no regressions)
- PHASE1_NEW_EXIT≠0   (challenge tests fail — correct, solution not applied yet)

### Phase 2 — both patches applied

Run these commands in order:

```bash
git apply solution.patch

# Try Docker. If missing: `./test.sh --output_path /tmp/base2.xml base`
docker build -t challenge-verify-full .
docker run --rm --network none challenge-verify-full ./test.sh --output_path /tmp/base2.xml base
echo "PHASE2_BASE_EXIT=$?"

# Try Docker. If missing: `./test.sh --output_path /tmp/new2.xml new`
docker run --rm --network none challenge-verify-full ./test.sh --output_path /tmp/new2.xml new
echo "PHASE2_NEW_EXIT=$?"
```

Expected:
- PHASE2_BASE_EXIT=0   (upstream tests still pass)
- PHASE2_NEW_EXIT=0    (challenge tests now pass with solution)

### Cleanup

```bash
git stash pop
```

### Output

Report EXACTLY this format after all commands complete. Fill in every field — do not skip any:

---

PHASE1_BASE: <PASS | FAIL (exit code N)>
  git apply result: <"OK" | "FAILED — error: <exact error message from git>">
  Failing tests (if any): <list test names, or "none">
  Root cause: <"patch corrupt/conflict" | "upstream test broken" | "docker error" | "none">

PHASE1_NEW: <FAIL (exit≠0) = CORRECT | PASS (exit=0) = BROKEN — tests must fail without solution>
  Failing tests: <list test names — these are expected to fail>

PHASE2_BASE: <PASS | FAIL (exit code N)>
  Failing tests (if any): <list test names, or "none">
  Root cause: <"solution introduced regression" | "docker error" | "none">

PHASE2_NEW: <PASS | FAIL (exit code N)>
  Failing tests (if any): <list each test name on its own line>
  Passing tests count: <N>
  Root cause: <"solution incomplete — missing implementation for these tests" | "test is incorrect" | "none">

---

MATRIX RESULT: <ALL PASS | FAIL>

What to fix next:
- PHASE1_BASE FAIL + git apply error → test.patch is corrupt or has wrong paths. Fix the patch formatting (line endings, diff headers, file paths). Run: `git apply --check test.patch` to see the exact error.
- PHASE1_BASE FAIL + git apply OK → a test in test.patch broke an existing upstream test. Check which base test failed — it is likely in a file your test.patch touches.
- PHASE1_NEW PASS (= BROKEN) → Your tests are broken (they pass without the solution). Open a NEW window and run `06_Fix/2_Fix_Tests.md` to rewrite them so they fail on the base repo.
- PHASE2_BASE FAIL → Your solution caused a regression. Open a NEW window and run `06_Fix/1_Fix_Solution.md` to fix the hunk that broke the old tests.
- PHASE2_NEW FAIL → Your solution is incomplete (or your test is wrong). Open a NEW window and run `06_Fix/1_Fix_Solution.md` to add the missing logic, or `2_Fix_Tests.md` if the test itself is bugged.
- ALL PASS → You did it! Open a NEW window and run `05_Review/PROMPT.md` for your final review.
