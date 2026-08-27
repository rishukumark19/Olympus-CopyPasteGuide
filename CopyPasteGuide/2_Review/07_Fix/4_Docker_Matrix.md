You are running the Docker verification matrix for a Go Olympus challenge. Execute immediately. Do not explain. Do not ask questions. Use your shell/terminal tool to run every command below.

### Working Context
<!-- Fill in BOTH fields before pasting -->
REPO_LOCAL_PATH: [ABSOLUTE PATH TO CLONED REPO — e.g. C:\Users\you\repos\myrepo]
CHALLENGE_SLUG:  [SLUG — the folder name inside challenge/ — e.g. least-latency-selection-policy]

**Navigation:** Use your shell/terminal to run:
```
cd <REPO_LOCAL_PATH>
```
The challenge folder is `challenge/<CHALLENGE_SLUG>/` inside the repo root. All git and docker commands run from the repo root.

**PREFLIGHT — run this first and STOP if any file is missing:**
```
ls challenge/<CHALLENGE_SLUG>/commit.txt
ls challenge/<CHALLENGE_SLUG>/test.patch
ls challenge/<CHALLENGE_SLUG>/solution.patch
ls challenge/<CHALLENGE_SLUG>/Dockerfile
```
If `Dockerfile` is missing: STOP. Go back to `1_Build/03A_Build/PROMPT.md` to regenerate it.
If `commit.txt` is missing: create it by running: `git rev-parse HEAD > challenge/<CHALLENGE_SLUG>/commit.txt`

**CRITICAL RULES FOR THE LLM:**
1. DO NOT write your own scripts (`.ps1`, `.py`, `.sh`).
2. DO NOT invent or overwrite `test.sh` — it must come from `test.patch` via `git apply`.
3. Execute the commands directly in your shell tool, one at a time.
4. If you are using PowerShell, substitute `$(cat ...)` with the actual hash you read from the file.
5. **DOCKER PREFERRED, NATIVE FALLBACK:** You MUST try to use Docker first. If `docker` is not installed or fails, skip the `docker build` and `docker run` commands and run `bash challenge/<CHALLENGE_SLUG>/test.sh` natively instead.

---

### Pre-Docker Gate (MUST run before any docker build)

Before building Docker, verify patches apply cleanly from a CLEAN checkout:

```bash
git stash
git reset --hard HEAD
git clean -fd -e challenge/
git checkout <INSERT_HASH_FROM_COMMIT_TXT>

# 1. Verify test.patch applies cleanly
git apply --check challenge/<CHALLENGE_SLUG>/test.patch
# If this fails: fix test.patch — do NOT proceed to Docker

git apply challenge/<CHALLENGE_SLUG>/test.patch

# 2. Verify test.sh has LF line endings (critical for Linux Docker)
if (Select-String -Path "challenge/<CHALLENGE_SLUG>/test.sh" -Pattern "`r" -Quiet) {
    Write-Host "CRLF DETECTED in test.sh — converting to LF now"
    (Get-Content challenge/<CHALLENGE_SLUG>/test.sh -Raw) -replace "`r","" | Set-Content challenge/<CHALLENGE_SLUG>/test.sh -NoNewline
    git add challenge/<CHALLENGE_SLUG>/test.sh
}

# If also verifying solution (Step 2):
git apply --check challenge/<CHALLENGE_SLUG>/solution.patch
# If this fails: fix solution.patch — do NOT proceed to Docker
```

**Only proceed to Docker builds once BOTH `--check` commands pass without errors.**

---

### Step 1: TEST CHECK (`test.patch` only — no solution)

Run these commands in order (adjust syntax if using PowerShell):

```bash
# STEP 1 — clean state, apply test.patch only (no solution)
git stash
git reset --hard HEAD
git clean -fd -e challenge/
git checkout <INSERT_HASH_FROM_COMMIT_TXT>
git apply challenge/<CHALLENGE_SLUG>/test.patch

# docker build uses the Dockerfile from the challenge folder
docker build -t challenge-verify -f challenge/<CHALLENGE_SLUG>/Dockerfile .
docker run --rm --network none challenge-verify bash ./challenge/<CHALLENGE_SLUG>/test.sh --output_path /tmp/base1.xml base
echo "TEST_CHECK_BASE_EXIT=$?"

docker run --rm --network none challenge-verify bash ./challenge/<CHALLENGE_SLUG>/test.sh --output_path /tmp/new1.xml new
echo "TEST_CHECK_NEW_EXIT=$?"
```

**Expected outcomes:**
- `TEST_CHECK_BASE_EXIT=0`   (Upstream tests pass — tests don't break existing features)
- `TEST_CHECK_NEW_EXIT≠0`   (New tests fail — correct, because solution is not implemented yet)

---

### Step 2: SOLUTION CHECK (`solution.patch` applied)

Run these commands in order:

```bash
git apply challenge/<CHALLENGE_SLUG>/solution.patch

docker build -t challenge-verify-full -f challenge/<CHALLENGE_SLUG>/Dockerfile .
docker run --rm --network none challenge-verify-full bash ./challenge/<CHALLENGE_SLUG>/test.sh --output_path /tmp/base2.xml base
echo "SOLUTION_CHECK_BASE_EXIT=$?"

docker run --rm --network none challenge-verify-full bash ./challenge/<CHALLENGE_SLUG>/test.sh --output_path /tmp/new2.xml new
echo "SOLUTION_CHECK_NEW_EXIT=$?"
```

**Expected outcomes:**
- `SOLUTION_CHECK_BASE_EXIT=0`   (Upstream tests still pass — no regressions)
- `SOLUTION_CHECK_NEW_EXIT=0`    (New challenge tests now pass — solution is complete)

---

### Cleanup

```bash
git stash pop
```

---

### Output Report

Report EXACTLY this format after all commands complete. Fill in every field — do not skip any:

---

```
========================= DOCKER MATRIX REPORT =========================

[1] TEST CHECK (test.patch only):
  • Base upstream tests : <PASS (exit 0) | FAIL (exit N)>
    - git apply result  : <"OK" | "FAILED — error: <exact error message>">
    - Failing tests     : <list names or "none">
    - Root cause        : <"patch corrupt/conflict" | "upstream test broken" | "docker error" | "none">
  • New challenge tests : <FAIL (exit≠0) = CORRECT | PASS (exit=0) = BROKEN — tests must fail without solution>
    - Failing tests     : <list test names — these are expected to fail>

[2] SOLUTION CHECK (solution.patch applied):
  • Base upstream tests : <PASS (exit 0) | FAIL (exit N) = REGRESSION>
    - Failing tests     : <list test names or "none">
    - Root cause        : <"solution introduced regression" | "docker error" | "none">
  • New challenge tests : <PASS (exit 0) = COMPLETE | FAIL (exit N) = INCOMPLETE>
    - Failing tests     : <list failing test names or "none">
    - Passing count     : <N>
    - Root cause        : <"solution incomplete" | "test is incorrect" | "none">

========================================================================
MATRIX RESULT: <ALL PASS | FAIL>
```

---

### What to Fix Next if Failed:

- **Dockerfile is missing** ➔ STOP. Go back to `1_Build/03A_Build/PROMPT.md` and regenerate it.
- **TEST CHECK Base Fails (git apply error)** ➔ `test.patch` has wrong format/paths. Run `git apply --check challenge/<CHALLENGE_SLUG>/test.patch` to see the exact error.
- **TEST CHECK Base Fails (compile or upstream test error)** ➔ A test in `test.patch` broke existing code or uses symbols that don't exist on base. Run `2_Review/07_Fix/5_Fix_Compile.md` or `2_Review/07_Fix/2_Fix_Tests.md`.
- **TEST CHECK New Passes (`exit=0`)** ➔ **BROKEN:** Tests pass without a solution! Run `2_Review/07_Fix/2_Fix_Tests.md` to make tests properly assert the new behavior.
- **SOLUTION CHECK Base Fails (`exit!=0`)** ➔ **REGRESSION:** Solution broke existing code. Run `2_Review/07_Fix/1_Fix_Solution.md`.
- **SOLUTION CHECK New Fails (`exit!=0`)** ➔ **INCOMPLETE:** Solution is missing logic. Run `2_Review/07_Fix/1_Fix_Solution.md` (or `2_Fix_Tests.md` if the test is wrong).
- **ALL PASS** ➔ Ready! Re-run `1_Build/05_Scope_Check/1_Check.md`.
