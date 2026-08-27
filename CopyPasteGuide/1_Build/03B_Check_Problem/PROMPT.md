You are auditing a problem.md for a Go Olympus challenge. Execute immediately. Do not fix anything — only report findings. Do not ask questions.

The solver will see ONLY this description and the pinned repo. They will NOT see the tests or solution.

### Working Context
<!-- Fill in BOTH fields before pasting -->
REPO_LOCAL_PATH: [ABSOLUTE PATH TO CLONED REPO — e.g. C:\Users\you\repos\myrepo]
CHALLENGE_SLUG:  [SLUG — the folder name inside challenge/ — e.g. least-latency-selection-policy]

**Navigation:** Use your shell/terminal to run:
```
cd <REPO_LOCAL_PATH>
cd challenge/<CHALLENGE_SLUG>
```
All subsequent file reads happen from this directory. Do NOT proceed if `problem.md` is not present.

### problem.md to audit:
(read problem.md from your current working directory — the challenge folder)

Count the words yourself after reading. Report the count in your output.

---

### Check 1: Format and Structure
Report any violation:
- Contains headings (##, ###, bold section titles acting as headings)
- Contains bulleted requirement lists (- or * used for requirements)
- Contains code blocks (``` or `)
- Contains motivation preamble ("currently X is missing", "the repo does not support", "this PR adds")
- Does NOT open with the request itself as the first sentence
- Contains non-ASCII characters
- Contains em-dashes (—) — this is an AI-generation tell
- Hard-wrapped paragraphs where lines break at ~70-85 chars mid-sentence — AI tell
- Word count outside 80-160 words (fail if below 80 or above 200, warn if 161-200) — use your own word count from reading the file

### Check 2: Content Fairness
Report any violation:
- Names private helpers, internal files, storage keys, fixture names, or implementation types
- Names specific internal method signatures when not required for fairness (the solver can discover these)
- States something the solver cannot determine from the pinned repo + this description alone (P3 violation)
- Leaks implementation approach, algorithm, or file structure (P6 violation)
- Contains external URLs (solver agents cannot fetch them)
- Contradicts repo design or proposes something maintainers have declined

### Check 3: Behavioral Completeness
Report any violation:
- A key behavior that the tests verify but the description does not mention
- A description clause with no corresponding test (would be removed to stay honest)
- Ordering or precedence dependencies that tests rely on but description leaves ambiguous
- Error/rejection behavior that tests assert but description does not specify

### Check 4: P1-P7 Audit
For each criterion, report PASS or specific issue:
- P1: Does it align with the repo's philosophy?
- P2: Does it avoid describing behavior already in an open/merged/closed PR?
- P3: Is it self-contained — solvable from repo + description alone?
- P4: Is it clear, concise, and unambiguous?
- P5: Is success objectively testable?
- P6: Is it non-prescriptive (no solution leaks)?
- P7: Is it materially distinct (not a rewording of an existing submission)?

### Check 5: Common Writing Mistakes
Report any violation:
- Framed as if the repo is external ("Library X currently lacks Y" / "the repo does not support") — must read as a maintainer issue, not an outsider report
- Written as a list of requests or snappy instructions instead of flowing natural prose
- Uses code snippets where plain English would work equally well (exception: exact shape, signature, or format required for fairness)
- Lists discoverable repo details (behavioral or implementation) that the solver will find in the codebase anyway
- AI-generated tone markers: em-dashes (—), hard line-wrapping at ~70-85 chars, "Additionally", "Furthermore", "It is worth noting"

### Output

Respond in exactly this format — no other text:

Word count: <N>

PASS
(if zero violations across all 5 checks)

-- or --

ISSUES FOUND:
1. [Check N] <specific issue and suggested fix>
2. [Check N] <specific issue and suggested fix>
