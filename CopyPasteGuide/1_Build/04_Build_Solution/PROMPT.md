You are writing a solution patch for a Go Olympus challenge. Execute immediately. Do not explain. Do not ask questions. Produce ONLY solution.patch.

**ISOLATION RULE — strictly enforced:** You must NOT read, open, or reference test.patch in this session. Solve this from problem.md and the repository ONLY. This mirrors exactly what a competing agent sees. If you have seen test.patch in an earlier step in this session, start a new session.

**IMPLEMENTATION PLAN RULE — strictly enforced:** Before writing a single line of code, produce a written implementation plan covering: which files will be touched, what new types/fields/functions are needed, and how each problem.md clause maps to a code change. Only begin coding after this plan is clear. This prevents mid-implementation pivots that corrupt the patch.

### Working Context
<!-- Fill in BOTH fields before pasting -->
REPO_LOCAL_PATH: [ABSOLUTE PATH TO CLONED REPO — e.g. C:\Users\you\repos\myrepo]
CHALLENGE_SLUG:  [SLUG — the folder name inside challenge/ — e.g. least-latency-selection-policy]

**Navigation:** Use your shell/terminal to run:
```
cd <REPO_LOCAL_PATH>
```
The challenge folder is at `challenge/<CHALLENGE_SLUG>/` inside this repo root. Read `problem.md`, `repo_url.txt`, and `commit.txt` from that subfolder. All patch and linting commands run from the repo root (`<REPO_LOCAL_PATH>`).

### Problem Description
(read problem.md from the challenge folder)

### Repository
URL: (read from repo_url.txt in the challenge folder)
Commit: (read from commit.txt in the challenge folder)
Local path: (parent directory of the challenge folder — the cloned repo root)

### Solution Patch Rules

solution.patch contains implementation changes ONLY:
- No test files
- No Dockerfile changes
- No test.sh changes
- No docs
- No generated files
- No unrelated formatting/import churn
- No dead code or padding
- No whitespace-only hunks
- No collateral edits outside the feature's scope

Keep the implementation as short as possible while meeting requirements naturally.
Follow repo conventions, ownership boundaries, types, errors, serialization, and command wiring.

Every new public symbol, field, output, or persisted shape must map to at least one sentence in problem.md. If you cannot trace a change back to a clause in problem.md, remove it.

Remove from the patch: untested stats, fingerprints, extra constructor fields, aliases, convenience APIs, duplicate helpers, unused fields, broad error suppression, visible-test workarounds.

Solution quality rules (S1-S4):
- S1: Solution meets ALL requirements (every problem.md clause)
- S2: No regressions — does not break existing passing tests; follows existing code patterns
- S3: No irrelevant changes — if unrelated to the task, leave it as-is
- S4: No AI slop — no weird comments, unexplained defensive code, new coding patterns inconsistent with repo

After writing all implementation code:
1. Run: `go build ./...` — must succeed with zero errors
2. Run: `go vet ./...`
3. Run: `gofmt -l .` (no files should be listed)
4. Run: `goimports -l .` (no files should be listed)
5. Run: `git apply --check challenge/<slug>/solution.patch` on a **clean** checkout to verify patch integrity before Docker.

**Patch generation — do it RIGHT:**
```bash
git add -p   # stage only solution files, NOT test files or Dockerfiles
git diff --cached > challenge/<slug>/solution.patch
# Then verify:
git apply --check challenge/<slug>/solution.patch
```
Never generate a patch by hand-editing. Never include test files in solution.patch.

Meaningful LOC target: 500+ meaningful LOC across 5+ existing non-test files touching 2+ subsystems. Meaningful LOC excludes: blank lines, comment-only lines, imports, declarations without logic, braces, generated files, mechanical propagation. Do NOT pad to hit numbers — let the scope emerge from the requirements.

Generate solution.patch with: git diff --cached
Confirm: git apply --check passes, no platform branding, no test code included.

### Output

Write solution.patch into the challenge folder in your current working directory.
Then confirm:

solution.patch written to challenge folder ✓
Meaningful LOC: <count>
Files touched: <list>
Subsystems touched: <list>
