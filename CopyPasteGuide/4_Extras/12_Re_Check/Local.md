# ╔══════════════════════════════════════════════════════════════╗
# ║  STEP 12 — Local (terminal checks, no LLM needed)      ║
# ╠══════════════════════════════════════════════════════════════╣
# ║  PURPOSE : Ground-truth verification with terminal commands.  ║
# ║            Docker never lies. Run this ALWAYS, before and    ║
# ║            after trusting any model output (cheap or strong). ║
# ║  NEEDS   : All artifacts on disk + Docker running locally.    ║
# ║            Replace <COMMIT_HASH> throughout with your hash.  ║
# ║  OUTPUT  : Each check prints PASS or FAIL to terminal.        ║
# ║            Docker run exit codes:                            ║
# ║              base run → must exit 0 (PASS)                   ║
# ║              new run (test only) → must exit non-zero (FAIL) ║
# ║              new run (both patches) → must exit 0 (PASS)     ║
# ║  DECISION: All checks PASS → artifacts are valid             ║
# ║            Any check FAIL → fix before any LLM audit         ║
# ╚══════════════════════════════════════════════════════════════╝
# Local.md — Run These Locally to Verify Any Model's Claims
#
# No LLM needed. These are terminal commands — they return objective truth.
# Run this BEFORE trusting any model output, cheap or expensive.
# Run this AFTER any artifact change to confirm nothing broke.
#
# Replace <COMMIT_HASH> and paths as needed throughout.

---

## 1. VERIFY: problem.md is valid ASCII, right word count, no forbidden patterns

```bash
# Word count (must be 80-160)
wc -w problem.md

# Must return "ASCII text" — no non-ASCII characters
file problem.md

# Must return nothing — no em-dashes
grep -P "\u2014" problem.md && echo "FAIL: em-dash found" || echo "PASS: no em-dashes"

# Must return nothing — no URLs
grep -oE 'https?://[^ ]+' problem.md && echo "FAIL: URL found" || echo "PASS: no URLs"

# Must return nothing — no headings
grep -E "^#{1,3} " problem.md && echo "FAIL: headings found" || echo "PASS: no headings"

# Must return nothing — no platform branding
grep -iE "olympus|shipd|mars|challenge" problem.md && echo "FAIL: branding found" || echo "PASS: clean"
```

---

## 2. VERIFY: Patches apply cleanly on the pinned commit

```bash
git stash
git checkout <COMMIT_HASH>

# test.patch applies cleanly
git apply --check test.patch && echo "PASS: test.patch applies" || echo "FAIL: test.patch conflict"

# solution.patch applies on top of test.patch
git apply test.patch
git apply --check solution.patch && echo "PASS: solution.patch applies" || echo "FAIL: solution.patch conflict"

git stash pop
```

---

## 3. VERIFY: No platform branding in patches

```bash
grep -iE "olympus|shipd|mars|challenge" test.patch && echo "FAIL: branding in test.patch" || echo "PASS"
grep -iE "olympus|shipd|mars|challenge" solution.patch && echo "FAIL: branding in solution.patch" || echo "PASS"
```

---

## 4. VERIFY: test.sh is executable and has correct argument handling

```bash
# Must show -rwxr-xr-x or similar (executable bit set)
git show HEAD:test.sh | head -1   # or check from patch
grep "mode 100755" test.patch && echo "PASS: executable bit set" || echo "FAIL: check chmod +x"

# Extract test.sh from patch and check it accepts --output_path
grep "output_path" test.patch | head -5

# Check it handles both base and new modes
grep -E "\bbase\b|\bnew\b" test.patch | head -10

# Must NOT contain --exitfirst or -x (fail-fast)
grep -E "\-\-exitfirst|\-x " test.patch && echo "FAIL: fail-fast flag found" || echo "PASS: no fail-fast"
```

---

## 5. VERIFY: Docker — base=PASS, new=FAIL (with test.patch only)

```bash
git stash
git checkout <COMMIT_HASH>
git apply test.patch
docker build -t re-audit-test .

# This MUST exit 0 (PASS)
docker run --rm --network none re-audit-test ./test.sh --output_path /tmp/base.xml base
echo "base exit code: $?"   # must be 0

# This MUST exit non-zero (FAIL — tests failing is correct here)
docker run --rm --network none re-audit-test ./test.sh --output_path /tmp/new.xml new
echo "new exit code: $?"    # must be non-zero (1 or 2)
```

---

## 6. VERIFY: Docker — base=PASS, new=PASS (with both patches)

```bash
# Continuing from above (test.patch already applied)
git apply solution.patch
docker build -t re-audit-test .

# Both MUST exit 0 (PASS)
docker run --rm --network none re-audit-test ./test.sh --output_path /tmp/base_full.xml base
echo "base exit code: $?"   # must be 0

docker run --rm --network none re-audit-test ./test.sh --output_path /tmp/new_full.xml new
echo "new exit code: $?"    # must be 0

git stash pop
```

---

## 7. VERIFY: JUnit XML is non-empty and contains real results

```bash
# After a Docker run, check the output file is real XML with test results
cat /tmp/new_full.xml | grep -E "<testsuite|<testcase" | head -10
# Must show actual test names, not an empty or fake file
wc -l /tmp/new_full.xml   # should be > 5 lines
```

---

## 8. VERIFY: solution.patch contains no test code

```bash
# Must return nothing (solution.patch must not touch _test.go files)
grep "_test\.go" solution.patch && echo "FAIL: test files in solution.patch" || echo "PASS: no test files"

# Must return nothing (no test.sh in solution)
grep "test\.sh" solution.patch && echo "FAIL: test.sh in solution.patch" || echo "PASS: clean"
```

---

## 9. VERIFY: New test filenames use random hex suffix

```bash
# All new _test.go files should end in _XXXXXX_test.go (6-char hex suffix before _test.go)
grep "^+++ b/.*_test\.go" test.patch | grep -vE "_[0-9a-f]{6}_test\.go$" \
  && echo "FAIL: test file missing hex suffix" || echo "PASS: all test files have hex suffix"
```

---

## 10. VERIFY: Dockerfile is correct shape

```bash
grep "FROM public.ecr.aws/d3j8x8q7/olympus-base-go:latest" Dockerfile \
  && echo "PASS: base image correct" || echo "FAIL: wrong base image"
grep "WORKDIR /app" Dockerfile && echo "PASS" || echo "FAIL: missing WORKDIR /app"
grep "COPY \. \." Dockerfile && echo "PASS" || echo "FAIL: missing COPY . ."
grep "go mod download" Dockerfile && echo "PASS" || echo "FAIL: missing go mod download"
grep 'CMD \["/bin/bash"\]' Dockerfile && echo "PASS" || echo "FAIL: wrong CMD"
grep "go test" Dockerfile && echo "FAIL: tests run during build" || echo "PASS: no tests in build"
grep "CHALLENGE_DOCKER=1" Dockerfile && echo "PASS: CHALLENGE_DOCKER set" || echo "FAIL: missing ENV CHALLENGE_DOCKER=1"
```

---

## QUICK FULL RUN (paste all at once)

```bash
echo "=== PROBLEM.MD ===" && wc -w problem.md && file problem.md
echo "=== PATCH APPLY ===" && git checkout <COMMIT_HASH> && git apply --check test.patch && git apply test.patch && git apply --check solution.patch
echo "=== BRANDING ===" && grep -iE "olympus|shipd|mars|challenge" test.patch solution.patch && echo "FOUND BRANDING" || echo "CLEAN"
echo "=== DOCKERFILE ===" && grep "olympus-base-go" Dockerfile && grep "WORKDIR /app" Dockerfile && grep 'CMD \["/bin/bash"\]' Dockerfile
echo "=== DOCKER BASE ===" && docker build -t re-audit . && docker run --rm --network none re-audit ./test.sh --output_path /tmp/b.xml base; echo "base=$?"
echo "=== DOCKER NEW ===" && docker run --rm --network none re-audit ./test.sh --output_path /tmp/n.xml new; echo "new=$? (should be non-zero)"
git apply solution.patch && docker build -t re-audit-full . && docker run --rm --network none re-audit-full ./test.sh --output_path /tmp/bf.xml base; echo "base+sol=$?" && docker run --rm --network none re-audit-full ./test.sh --output_path /tmp/nf.xml new; echo "new+sol=$? (should be 0)"
```
