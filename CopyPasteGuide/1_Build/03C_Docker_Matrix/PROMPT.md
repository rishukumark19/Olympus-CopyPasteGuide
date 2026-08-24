You are running the Docker verification matrix for a Go Olympus challenge — Phase 1 only (test.patch, no solution). Execute immediately. Do not explain. Do not ask questions. Use your shell/terminal tool to run every command below.

### Setup

Read the pinned commit hash from `commit.txt` in your current working directory.
The challenge folder is your current working directory.

**CRITICAL RULES FOR THE LLM:**
1. DO NOT write your own scripts (`.ps1`, `.py`, `.sh`).
2. Execute the commands directly in your shell tool.
3. If you are using PowerShell, substitute `$(cat commit.txt)` with the actual hash you read from the file.
4. **DOCKER PREFERRED, NATIVE FALLBACK:** Try Docker first. If `docker` is not installed or fails, run `./test.sh` natively on the host.

### Phase 1 — test.patch only (no solution exists yet)

```bash
git stash
git checkout <INSERT_HASH_FROM_COMMIT_TXT>
git apply test.patch

# Try Docker. If missing: `./test.sh --output_path /tmp/base1.xml base`
docker build -t challenge-verify .
docker run --rm --network none challenge-verify ./test.sh --output_path /tmp/base1.xml base
echo "PHASE1_BASE_EXIT=$?"

# Try Docker. If missing: `./test.sh --output_path /tmp/new1.xml new`
docker run --rm --network none challenge-verify ./test.sh --output_path /tmp/new1.xml new
echo "PHASE1_NEW_EXIT=$?"
```

Expected:
- PHASE1_BASE_EXIT=0   (upstream tests pass — no regressions)
- PHASE1_NEW_EXIT≠0   (challenge tests fail — correct, no solution yet)

### Cleanup

```bash
git stash pop
```

### Output

Report exactly this format. Fill in every field:

---

PHASE1_BASE: <PASS | FAIL (exit code N)>
  git apply result: <"OK" | "FAILED — error: <exact error message>">
  Failing tests (if any): <list test names, or "none">
  Root cause: <"patch corrupt/conflict" | "upstream test broken" | "docker error" | "none">

PHASE1_NEW: <FAIL (exit≠0) = CORRECT | PASS (exit=0) = BROKEN — tests must fail without solution>
  Failing tests: <list test names — these are expected to fail>

---

RESULT: <PASS (both checks correct) | FAIL>

What to fix next:
- PHASE1_BASE FAIL + git apply error → test.patch is corrupt. Run: `git apply --check test.patch` to see the exact error. Fix with Fix_Tests.md.
- PHASE1_BASE FAIL + git apply OK → a test in test.patch broke an existing upstream test. Check which base test failed.
- PHASE1_NEW PASS (= BROKEN) → your challenge tests pass WITHOUT the solution. They are not testing the right thing. Fix with Fix_Tests.md.
- RESULT PASS → test.patch is verified. Proceed to Step 04 (in a new session).
