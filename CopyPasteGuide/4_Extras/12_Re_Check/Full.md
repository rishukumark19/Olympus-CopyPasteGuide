You are doing a full quality re-audit of a completed Go Olympus challenge. Execute immediately. Do not fix files. Report findings only. Return a verdict for each axis. Do not ask questions.

### Artifacts:

problem.md:
(read problem.md from your current working directory)

test.patch:
(read test.patch from your current working directory)

solution.patch:
(read solution.patch from your current working directory)

Dockerfile:
(read Dockerfile from your current working directory)

Repo + commit: (read repo URL and commit from repo_url.txt and commit.txt in your current working directory)

test.sh (extracted from test.patch):
(read test.sh from your current working directory)

---

### Audit Sequence

Run all checks below. Do not stop early. Report incomplete checks as findings.

**1. Problem Description Audit (P1-P7)**

Check:
- P1: Aligns with repo philosophy
- P2: Not solved by any open/merged/closed PR
- P3: Self-contained, solvable from repo + description alone
- P4: Clear, concise, unambiguous
- P5: Objectively testable
- P6: Non-prescriptive, no solution leaks
- P7: Materially distinct

Also check format:
- 80-160 ASCII words
- No headings, bullets, code blocks, URLs, em-dashes, non-ASCII
- Opens with the request as first sentence
- No hard line-wrapping at ~70-85 chars
- No platform branding

Also check common writing mistakes:
- Framed as external ("the repo lacks X", "X is currently not supported") instead of maintainer prose
- Written as a list of requests or snappy instructions instead of flowing prose
- Uses code snippets where plain English would work equally well
- Lists discoverable repo details (behavior or implementation) the solver finds in the codebase anyway
- AI-tone markers: em-dashes, "Additionally,", "Furthermore,", "It is worth noting"

**2. Test Coverage Audit (T1-T8)**

For each test in test.patch:
- T1: Would it fail on base commit for its own missing behavior? (Not just compilation)
- T2: Is it deterministic across machines and runs?
- T3: Could a wrong/shortcut implementation still pass it?
- T4: Does it cover the relevant behavior + edge cases?
- T5: Does it only assert behavior in problem.md or discoverable from pinned repo?
- T6: Is it fully offline?
- T7: Does it over-pin exact output/error text without contractual reason?
- T8: Are failure diagnostics visible (no catch-all hiding real failures)?

Also check test.sh:
- Accepts exactly one mode (base or new) and one --output_path
- base mode covers all touched upstream surfaces
- new mode runs only challenge tests
- No fail-fast flags
- Emits real JUnit XML

**3. Solution Quality Audit (S1-S4)**

For each hunk in solution.patch:
- S1: Does it implement a requirement from problem.md?
- S2: Does it follow existing repo patterns and avoid regressions?
- S3: Is it limited to the feature scope (no collateral edits)?
- S4: No AI slop (weird comments, dead code, unexplained defensive logic)?

Also check:
- Every hunk has at least one test in test.patch (if not: either add test or remove hunk)
- Meaningful LOC count: 250-600 range
- 4+ existing non-test files touched
- 2+ subsystems touched
- go vet / gofmt / goimports clean

**4. Cross-Alignment Check**

- Every problem.md clause has a test AND a solution hunk
- Every test maps to a problem.md clause
- Every solution hunk maps to a test
- No orphan clauses, orphan tests, or orphan hunks

---

### Output

Respond in exactly this format — no other text:

Problem Description: <PASS or ISSUES FOUND>
Issues:
1. [P#] <specific finding, exact quote, repair instruction>

Tests: <PASS or ISSUES FOUND>
Issues:
1. [T#] <test name, what's wrong, what shortcut survives if T3/T4>

Solution: <PASS or ISSUES FOUND>
Issues:
1. [S#] <hunk location, what's wrong, repair instruction>

Cross-Alignment: <PASS or ISSUES FOUND>
Issues:
1. <orphan clause/test/hunk>

Overall: <CLEAN (all four PASS) or NEEDS WORK (list which axes have issues)>

Findings ordered by: Docker-provable failures first, then coverage gaps, then wording concerns.
Do not say "looks good" — only report violations or explicitly confirm PASS per criterion.
