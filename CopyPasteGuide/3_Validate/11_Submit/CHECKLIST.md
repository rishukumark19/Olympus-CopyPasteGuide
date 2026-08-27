# ╔══════════════════════════════════════════════════════════════╗

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.
# ║  STEP 10 — Final Submission Checklist                        ║
# ╠══════════════════════════════════════════════════════════════╣
# ║  PURPOSE : Final sanity check before submitting. No LLM      ║
# ║            needed — this is all human + terminal verification.║
# ║  NEEDS   : All 5 artifacts + all pipeline gates logged in     ║
# ║            00_Progress/STATE.md as PASS.                ║
# ║  OUTPUT  : N/A — you fill in the checkboxes yourself.        ║
# ║            Sections: Artifacts, Docker, problem.md, tests,   ║
# ║            test.sh, solution, Dockerfile, pipeline gates,    ║
# ║            final verification commands.                      ║
# ║  DECISION: ALL boxes checked → SUBMIT                        ║
# ║            ANY box unchecked → fix it first, DO NOT submit   ║
# ╚══════════════════════════════════════════════════════════════╝
# STEP 10 — Final Submission Checklist
#
# WHAT TO DO:
#   Go through EVERY item below. ALL must be GREEN (checked) before submitting.
#   Do NOT submit if anything is RED or UNKNOWN.
#   Fix it, re-run the relevant Docker checks, and come back to this checklist.

---

## PRE-SUBMISSION FINAL CHECKLIST

### A. Artifacts Present
[ ] problem.md exists
[ ] test.patch exists (contains only new test files + root test.sh)
[ ] solution.patch exists (contains only implementation changes)
[ ] Dockerfile exists
[ ] repo_url.txt exists (contains URL + immutable commit hash)

---

### B. Docker Full Matrix (run one final time clean from the pinned commit)

```bash
# CLEAN checkout, no patches applied
# WARNING: NEVER run `git clean -fd` without `-e challenge/` or it will delete all your work!
git checkout <commit-hash>
git reset --hard HEAD
git clean -fd -e challenge/
git apply challenge/<slug>/test.patch
docker build -t final-test challenge/<slug>/
docker run --rm --network none final-test ./test.sh --output_path /tmp/base.xml base   # MUST PASS
docker run --rm --network none final-test ./test.sh --output_path /tmp/new.xml new    # MUST FAIL

git apply solution.patch
docker build -t final-test .
docker run --rm --network none final-test ./test.sh --output_path /tmp/base.xml base   # MUST PASS
docker run --rm --network none final-test ./test.sh --output_path /tmp/new.xml new    # MUST PASS
```

[ ] base PASSES with only test.patch applied
[ ] new FAILS with only test.patch applied
[ ] base PASSES with both patches applied
[ ] new PASSES with both patches applied
[ ] All Docker runs use --network none

---

### C. Problem Description (problem.md)
[ ] 80-160 ASCII words (word count: ___)
[ ] Opens with the request as the first sentence (no preamble)
[ ] No headings (##, ###, bold acting as section titles)
[ ] No bulleted requirement lists
[ ] No code blocks
[ ] No external URLs
[ ] No non-ASCII characters
[ ] No em-dashes (—)
[ ] No motivation preamble ("currently X is missing...")
[ ] No private API names, internal file names, or implementation hints
[ ] Every clause corresponds to a tested behavior
[ ] P1: Aligns with repo philosophy
[ ] P2: No PR (open/merged/closed) already implements this
[ ] P3: Self-contained, solvable from repo + description alone
[ ] P4: Clear, concise, unambiguous
[ ] P5: Objectively testable
[ ] P6: Not prescriptive, no solution leaks
[ ] P7: Distinct from existing submissions (checked similarity)

---

### D. Test Patch (test.patch)
[ ] Contains ONLY new test files + root test.sh
[ ] No edits to upstream test files or shared helpers
[ ] No platform branding: "challenge", "olympus", "shipd", "mars" in filenames or code
[ ] New test filenames use openssl rand -hex 3 hex suffix
[ ] All new tests: 100% fail on base commit
[ ] All new tests: 100% pass with solution
[ ] No network access in any test
[ ] No timing dependencies, no random seeds, no goroutine scheduling assumptions
[ ] No over-pinning exact error text (unless contractually required)
[ ] Failure diagnostics intact (no catch-all messages hiding real failures)
[ ] Negative/failure paths tested as first-class behaviors
[ ] Coverage matrix complete (all claimed behaviors, edge cases, cross-products)
[ ] git apply --check passes cleanly on base commit

---

### E. test.sh
[ ] Accepts exactly one mode: base or new
[ ] Accepts exactly one --output_path (or --output_path=) value
[ ] Rejects: missing mode, missing output path, duplicate args, unknown args, extra args
[ ] Emits non-empty JUnit XML at the requested path (gotestsum or go-junit-report)
[ ] base mode: runs upstream touched-surface suite excluding only new challenge tests
[ ] base mode: does NOT use fail-fast flags
[ ] new mode: runs ONLY new challenge tests
[ ] Uses go test ./... when CHALLENGE_DOCKER=1
[ ] Returns real exit status (does not mask failures)
[ ] Never installs, fetches, or mutates outside the repo
[ ] test.sh is executable (chmod +x)

---

### F. Solution Patch (solution.patch)
[ ] Contains ONLY implementation changes (no tests, no Dockerfile, no docs, no test.sh)
[ ] No unrelated formatting or import churn
[ ] No dead code or padding
[ ] No whitespace-only hunks
[ ] No collateral edits outside feature scope
[ ] go vet passes (no new errors)
[ ] gofmt -l . shows no files (code is formatted)
[ ] goimports -l . shows no files
[ ] No AI slop: weird comments, unexplained defensive code, inconsistent patterns
[ ] Every new public symbol maps to a problem.md clause AND a test
[ ] Meaningful LOC: ___  (target 300-500; flag if below 250 or above 600 for human review)
[ ] Files touched: ___ (target 4+)
[ ] Subsystems touched: ___ (target 2+)
[ ] git apply --check passes cleanly (test.patch + solution.patch in sequence)

---

### G. Dockerfile
[ ] FROM public.ecr.aws/d3j8x8q7/olympus-base-go:latest
[ ] WORKDIR /app
[ ] COPY . .
[ ] ENV CHALLENGE_DOCKER=1
[ ] RUN go mod download (at build time, not runtime)
[ ] CMD ["/bin/bash"]
[ ] No test execution during build
[ ] No runtime network fetches
[ ] Built for amd64 (not arm64)
[ ] .dockerignore validated (not excluding required files)

---

### H. Pipeline Gates Passed

**Required (must be checked before submitting):**
[ ] 01 Scouter: repo passed all hard filters
[ ] 02 Pick Seed: no PR overlap found, seed is behavioral
[ ] 03A Problem+Tests: base PASS, new FAIL confirmed in Docker
[ ] 03B Problem Quality: PASS
[ ] 04 Solution Writer: base PASS, new PASS confirmed in Docker

**Optional quality steps (run when time allows — not required to submit):**
[ ] 05 Scope Check: READY (250+ LOC, 2+ subsystems)
[ ] 06 Full Review: ACCEPTED (3/3 all axes)
[ ] 08 Coverage Audit: CLEAN
[ ] 09 Anti-Shortcut: PASS
[ ] 10 Blind Testing: 1–4/5 legitimate passes

> **Note:** The platform auto-reviewer score you receive after submitting is feedback, not a gate. If Docker `NEW=PASS`, your challenge is valid. Do not hold back a submission waiting for a perfect auto-review score.

---

### I. Final Verification Commands (run before submit)

```bash
# Confirm both patches apply cleanly from pinned commit
git stash
git checkout <commit-hash>
git apply --check test.patch
git apply test.patch
git apply --check solution.patch
git apply solution.patch

# Confirm no platform branding
grep -r "olympus\|shipd\|mars\|challenge" test.patch

# Confirm problem.md is ASCII only
file problem.md  # should say "ASCII text"

# Confirm test.sh is executable
ls -la test.sh   # should show -rwxr-xr-x
```

---

## IF ANYTHING IS RED:
Fix it, re-run the relevant Docker checks, and re-run this checklist.
Do NOT submit until every item is GREEN.
