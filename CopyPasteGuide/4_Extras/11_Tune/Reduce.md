You are trimming a Go Olympus challenge solution that is too large or contains padding. Execute immediately. Do not explain, do not ask questions. Trim solution.patch to the minimum correct implementation that passes all tests. Do NOT change problem.md or test.patch.

### Current problem.md:
(read problem.md from your current working directory)

### Current test.patch:
(read test.patch from your current working directory)

### Current solution.patch:
(read solution.patch from your current working directory)

### Repository:
URL: (read from repo_url.txt in your current working directory)
Commit: (read from commit.txt in your current working directory)

### Critic Findings (LOC/padding issue):
[PASTE THE RELEVANT CRITIC FINDINGS]

---

### Reduction Rules

Target after reduction:
- 300-500 meaningful LOC (never below 250)
- Every remaining hunk has at least one direct behavioral test in test.patch
- No hunk can be removed without causing a new test to fail

Meaningful LOC definition (what counts):
- Logic lines with actual branching, computation, state mutation, or behavioral effect
NOT counting: blank lines, comment-only lines, imports, declarations without logic, braces, generated files, mechanical propagation, whitespace-only lines, boilerplate

Hunk-by-hunk audit:
For every hunk in solution.patch, ask: "If I remove this hunk, do any new tests fail?"
- If YES → hunk is verified, keep it
- If NO → hunk is unverified. Either:
  a. Remove the hunk (it is unnecessary padding or dead code), OR
  b. If it IS genuinely required behavior: add the missing behavioral test (requires re-running Docker loop)

Remove from solution.patch:
- Extra constructor fields that are never read by tests
- Convenience aliases or wrapper functions with no test
- Duplicate helpers that can be replaced by existing repo-native paths
- Broad error suppression (catching errors and ignoring them without behavioral test)
- Over-decomposition: functions that do one thing that could be inline
- Mechanical propagation: wiring a value through 5 layers with no test at intermediate layers
- Unrelated formatting, import reordering, whitespace changes in untouched files

Style and conventions:
- Run go vet, gofmt, goimports after trimming
- Collapsed code must still follow existing repo patterns (S4)
- No AI slop introduced during simplification

If LOC after reduction would fall below 250:
- Do not continue reducing
- Report that the scope itself is too thin and recommend Expand.md instead

### Output

Respond in this order:

Hunk-by-hunk analysis:
- Hunk <file:line>: KEEP / REMOVE / NEEDS-TEST — <one line reason>

**solution.patch:**
Save the revised patch directly to the file. Do not print it in chat.

Final LOC estimate: <per-file breakdown, total meaningful LOC>

Hunks needing new tests (flag for human review — do NOT write the tests):
- <hunk> requires a test for <behavior> before it can be kept
