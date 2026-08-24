# Check_Problem.md — 4 separate atomic audits for problem.md

**Automatic Navigation:** Before reading any files, use your shell/terminal to list the contents of `challenge/` and `cd` into the correct challenge folder (the one containing `problem.md`, `test.patch`, `solution.patch`, and `Dockerfile`). Do NOT proceed if you are in the wrong directory.
# Run each pass independently. Paste one block at a time to a fresh LLM session.
# All 4 must return PASS before problem.md is considered clean.

---

## PASS 1 of 4 — Format Check

You are doing a single-focus format audit of a problem.md file. Execute immediately. Answer PASS or list every violation. Do not explain what is fine — only report violations. Do not ask questions.

### problem.md:
(read problem.md from your current working directory)

### Word count: [PASTE WORD COUNT]

Check ONLY these format rules (report each violation separately):
1. Word count below 80 or above 200 → FAIL (warn if 161-200)
2. Contains any heading: ## or ### or bold text acting as a section title → FAIL
3. Contains any bulleted requirement list (lines starting with - or * as requirements) → FAIL
4. Contains any code block (``` or `) → FAIL
5. Does NOT open with the actual request as the first sentence (starts with preamble instead) → FAIL
6. Contains non-ASCII characters → FAIL
7. Contains em-dashes (—) → FAIL
8. Lines are hard-wrapped at ~70-85 characters mid-sentence (AI generation tell) → FAIL
9. Contains external URLs (http:// or https://) → FAIL
10. Contains the words: olympus, shipd, mars, challenge (any case) → FAIL

Output format:
PASS
-- or --
FAIL:
1. [rule number] exact quote of the violation

---

## PASS 2 of 4 — Content Fairness Check

You are doing a single-focus fairness audit of a problem.md file. Execute immediately. The solver has access ONLY to this description and the pinned repository. Nothing else. Answer PASS or list every violation. Do not explain what is fine — only report violations. Do not ask questions.

### problem.md:
(read problem.md from your current working directory)

Check ONLY these fairness rules (report each violation separately):
1. Names a private helper function, internal file path, storage key, fixture name, or implementation type that the solver cannot find from public API alone → FAIL
2. Names a specific internal method signature when not required (solver can discover it from the repo) → FAIL
3. States something the solver cannot determine from the pinned repo + this description alone → FAIL (P3)
4. Leaks the implementation approach, algorithm choice, or specific file to edit → FAIL (P6)
5. Contains external URLs (solver agents cannot fetch them) → FAIL
6. Proposes behavior that contradicts the repo's design philosophy or that maintainers have explicitly declined → FAIL (P1)
7. Describes behavior already implemented in an open, closed, or merged PR → FAIL (P2)

Output format:
PASS
-- or --
FAIL:
1. [rule number] exact quote of the violation and why it violates the rule

---

## PASS 3 of 4 — Behavioral Completeness Check

You are doing a single-focus completeness audit of a problem.md file alongside its test.patch. Execute immediately. Answer PASS or list every violation. Do not explain what is fine — only report violations. Do not ask questions.

### problem.md:
(read problem.md from your current working directory)

### test.patch (just the test assertions, not the full diff if large):
[PASTE test.patch OR PASTE ASSERTION SUMMARY]

Check ONLY these completeness rules:
1. A key behavior that the tests assert but the description does not mention at all → FAIL (solver cannot implement it — unfair T5/P3)
2. A description clause that has NO corresponding test assertion → FAIL (either add the test or remove the clause)
3. Ordering or precedence that tests rely on but the description leaves ambiguous → FAIL
4. Error or rejection behavior that tests assert but the description does not specify → FAIL

Output format:
PASS
-- or --
FAIL:
1. [rule number] exact clause from problem.md or exact test name, and what is missing

---

## PASS 4 of 4 — P1-P7 Final Confirmation

You are doing a final P1-P7 confirmation audit of a problem.md. Execute immediately. For each criterion below, answer only PASS or give the specific issue. One line per criterion. Do not ask questions.

### problem.md:
(read problem.md from your current working directory)

Criteria:
P1: Aligns with the repo's design philosophy (not a feature the repo would reject)
P2: Not already implemented in any open, merged, or closed PR
P3: Self-contained — solvable from the pinned repo + this description alone
P4: Clear, concise, unambiguous — no guessing required about what to build
P5: Objectively testable — success can be determined by automated assertions
P6: Non-prescriptive — does not leak the solution approach, algorithm, or internal structure
P7: Materially distinct from existing submissions — not a rewording of the same behavior

Output format (one line per criterion):
P1: PASS
P2: PASS
P3: FAIL — <specific issue>
P4: PASS
P5: PASS
P6: PASS
P7: PASS
