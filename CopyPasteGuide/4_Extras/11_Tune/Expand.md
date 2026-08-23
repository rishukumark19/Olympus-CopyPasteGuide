You are expanding a Go Olympus challenge that is too thin in scope. Execute immediately. Do not explain, do not ask questions. Expand the behavioral scope through natural same-workflow behaviors only. Do not invent unrelated requirements.

### Original seed and expansion plan:
[PASTE SEED SUMMARY FROM STEP 02]

### Current problem.md:
(read problem.md from your current working directory)

### Current test.patch:
(read test.patch from your current working directory)

### Current solution.patch:
(read solution.patch from your current working directory)

### Repository:
URL: (read from repo_url.txt in your current working directory)
Commit: (read from commit.txt in your current working directory)

### Critic Findings (LOC/scope issue):
[PASTE THE RELEVANT CRITIC FINDINGS]

---

### Expansion Rules

Target after expansion:
- 300-500 meaningful implementation LOC (never exceed 600)
- 4+ existing non-test files touched
- 2+ subsystems touched
- Problem.md stays 80-160 words

Meaningful LOC definition (what counts):
- Logic lines with actual branching, computation, state mutation, or behavioral effect
NOT counting: blank lines, comment-only lines, imports, declarations without logic, braces, generated files, mechanical propagation, whitespace-only lines, boilerplate

How to expand (in order of preference):

1. Same-workflow lifecycle behaviors: persistence/restart paths, error/rejection paths,
   idempotency, rollback, concurrent access, ordering constraints — all within the SAME user workflow.

2. Named variants or backends: if the original seed involves one backend/mode/dialect,
   extend to the next adjacent one in the SAME workflow. Do NOT add a completely different workflow.

3. Observability / readback behaviors: if the feature creates state, add that the state
   is correctly reflected in list/get/status APIs — same workflow, just more complete.

Do NOT expand by:
- Adding unrelated features or a second disconnected workflow
- Adding padding: extra constructors, aliases, convenience wrappers with no behavioral test
- Adding migration/compatibility code that is not actually required for the problem
- Inventing behaviors not discoverable from the repo

For each new behavior added:
- Add the observable public API test for it in test.patch (it must fail on base, pass with solution)
- Add the implementation in solution.patch
- Add the corresponding clause in problem.md (if the solver needs to know it)
- Confirm the new clause is testable and non-prescriptive

Reference-to-test closure: After expanding, every new solution hunk must have a corresponding test. If a hunk has no observable test, remove it from solution.patch.

### Output

Respond in this order:

Expansion analysis: <which natural same-workflow behaviors can be added, with evidence from repo>

**problem.md:**
Save the revised text directly to the file. Do not print it in chat.

**test.patch:**
Save the revised patch directly to the file. Do not print it in chat.

**solution.patch:**
Save the revised patch directly to the file. Do not print it in chat.

LOC estimate: <which files, how many new lines per file, total meaningful LOC>
