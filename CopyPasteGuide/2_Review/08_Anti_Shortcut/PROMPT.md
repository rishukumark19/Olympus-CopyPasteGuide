You are a false positive hunter and mutation tester for a Go Olympus challenge. Execute immediately. Do not explain, do not ask questions. Determine whether an incorrect implementation can pass the complete test suite while violating the actual task contract.

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.

**Execution order (follow this to save tokens — do not run everything at once):**
1. Establish contract (align problem.md, tests, solution)
2. Use Test Fairness suggestions early — review every 💡 before writing mutations (they surface real gaps cheaply)
3. Atomize the prompt and build the requirement-to-test map
4. Decompose reference solution into states
5. Run VCA / stub the feature — catch major structural gaps first
6. Actively search for FPs (mutation testing)
7. Test partially correct and hardcoded implementations
8. Audit edge/boundary/error/state/interaction coverage
9. Audit assertion strength
10. Fairness-check every new assertion before adding it
11. Final adversarial pass — run the simplest wrong implementation that survives everything

### Artifacts:

problem.md:
(read problem.md from your current working directory)

test.patch:
(read test.patch from your current working directory)

solution.patch:
(read solution.patch from your current working directory)

test.sh:
(read test.sh from your current working directory)

Repo: (read repo URL and commit from repo_url.txt and commit.txt in your current working directory)

Available agent patches (if any): [PASTE or write NONE]

---

### 1. Establish and Align the Contract

Treat ONLY these as requirements:
1. Behavior explicitly required by the problem statement.
2. Behavior clearly discoverable from existing public repository behavior.

Do NOT treat additional behavior implemented only by the reference solution as required.

Verify that problem.md, tests, and reference solution describe the same observable behavior.
Look for: contradictions, requirements in prompt absent from tests, tested requirements absent from prompt, reference behavior not actually required, requirement leaks, implementation details exposed as requirements.

### 2. Atomize the Prompt

Break every prompt clause into independently testable behavioral atoms.
Extract atoms for: positive behavior, negative behavior, defaults, every enumerated value, documented errors/failures, boundaries, state transitions, repeated operations, ordering, interactions, distinct valid input classes.
Treat symmetric operations independently (set/reset, apply/rollback, encode/decode, enable/disable, add/remove, start/stop, success/failure). Testing one side does not prove the other.

### 3. Decompose the Reference Solution into States

Break the reference solution into logical states. For each state, identify: public behavior, validation logic, error handling, edge cases, defaults, state updates, branches, input distinctions, transitions in/out.
For each state: "What functionality must work correctly for this state to be valid?"

### 4. Build a Requirement-to-Test Map

For every behavioral atom: identify the exact test/assertion verifying it, explain what incorrect behavior it rejects, classify as: explicitly required, repo-discoverable, flexible/underspecified, or outside scope.
Flag atoms without meaningful verification.

### 5. Actively Search for False Positives

Take the reference implementation. Inspect hunk by hunk and branch by branch.
For every behavioral atom, construct the smallest implementation that violates it.

Mutation examples:
- Flip a boolean
- Change a constant
- Delete a branch
- Stub a branch to a no-op
- Remove a match arm
- Ignore an argument/input
- Return a fixed value or always return default
- Skip validation
- Ignore an error
- Skip a state update
- Preserve stale state
- Break one side of a symmetric operation
- Collapse distinct inputs to same result
- Remove boundary handling
- Handle only examples appearing in tests

ACTUALLY RUN the relevant test suite after each mutation. Never conclude a mutation "should fail" without executing it.

Record: behavioral atom, mutation, violated behavior, test result, whether mutation survived.
If a mutation stays green → that behavior is not adequately verified.

### 6. Test Partially Correct and Hardcoded Implementations

Attempt candidates that:
- Implement A and B but ignore C
- Support only documented examples
- Handle success but not failure
- Handle one side of a symmetric pair
- Work only on first invocation
- Ignore certain arguments
- Return expected constants
- Recognize only tested inputs
- Implement only happy path

Vary valid inputs beyond existing fixtures. If different valid inputs should produce different observable outcomes, verify tests enforce those distinctions.

### 7. Edge, Boundary, Error, State, and Interaction Coverage

Where required by contract, check:
- Empty inputs, single and multiple values, min/max boundaries, just below/at/above boundaries
- Duplicates, missing optional values, invalid values, different ordering
- Initialized/uninitialized state, repeated operations
- Success followed by failure, failure followed by success
- Forward and reverse state transitions, interactions between components

For documented failures: failure occurs when required, incorrect success rejected, required public error behavior preserved, state correct after failure, subsequent behavior correct where specified.
Do NOT require exact error text unless established by the contract.

### 8. Audit Assertion Strength

Identify assertions that execute code without proving required behavior:
- is_ok(), is_some(), length > 0, contains(...), object exists, truthy checks

If contract requires exact observable result, verify it directly.
For suspicious assertions: deliberately change the expected behavior and rerun — if it stays green, the assertion is decorative.

### 9. Stub the Feature

Create extremely incomplete implementation while preserving compilation.
Replace functionality with no-ops. Return defaults. Remove most new logic. Stub major branches.
Run complete suite. Feature-specific tests should be heavily red. Investigate anything green.

### 10. Prove a False Positive

A valid FP exists ONLY when ALL of these hold:
1. Candidate passes every relevant new test
2. It violates an explicit or repository-discoverable requirement
3. A fair public API probe exposes the violation
4. The probe fails on the incorrect candidate
5. The exact probe passes on the reference implementation

For every potential FP:
1. Identify violated contract atom
2. Construct smallest fair public API probe
3. Run against incorrect candidate → confirm fails
4. Run against reference → confirm passes
5. Add ONLY after proving behavior is required

### 11. Fairness Gate

Every assertion must answer: "Could an agent know this requirement from problem.md or clearly discoverable public repository behavior?"
If no: either state it clearly in problem.md, or remove/reject the assertion.
Never keep a hidden requirement merely because it kills a mutation.

Flag tests requiring: unstated representations, unstated algorithms, implementation-specific strategies, internal helpers/private APIs, specific internal file layouts, exact error text not in contract, external protocol semantics not named in prompt.

### 12. Final Adversarial Pass

Ask: "What is the simplest wrong implementation that can still pass everything?"
Create it and RUN it.
Try different behavioral atoms until you cannot construct a meaningfully incorrect implementation that simultaneously:
1. Passes the complete suite
2. Violates an explicit or repository-discoverable requirement
3. Can be exposed by a fair public API probe

PASS requires executable adversarial testing. Never declare PASS from: tests looking comprehensive, coverage appearing high, static analysis finding nothing, existing examples passing.

### Output

Respond in exactly this format:

STATUS: <PASS or NEEDS CHANGES>

Contract Atoms: <every independently testable requirement>
State Model: <reference states/milestones and required behavior within each>
Coverage Map: <requirement/state → exact test/assertion → classification>
Proven False Positives: <violated requirement, mutation, existing suite result, probe, candidate result, reference result>
Surviving Mutations: <mutations that remained green>
Missing Coverage: <behavioral, edge, boundary, error, state, interaction gaps>
Weak Assertions: <assertions that fail to prove required behavior>
Hardcoding Risks: <shortcut candidates capable of passing>
Unfair Tests: <assertions enforcing behavior outside contract>
Stub Results: <results from deliberately incomplete implementation>
Alignment Issues: <mismatches between prompt, tests, reference>
Changes Made: <each test/prompt/solution change and its contractual justification>
Remaining FP Risks: <anything not conclusively validated>
Final Adversarial Result: <last incorrect candidates attempted and whether suite rejected them>

Do not declare PASS until executable adversarial testing completed and no known fair contract violation can pass.
