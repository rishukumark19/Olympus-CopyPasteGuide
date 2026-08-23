# STEP 12 — Re-Audit Mode
#
# Use this folder when you ALREADY HAVE built artifacts and want to:
#   a) Re-check the quality of existing work (second opinion, fresh eyes)
#   b) Verify nothing broke after a change
#   c) Run on a cheaper/smaller model safely without risking hallucinated PASSes
#   d) Resume audit from mid-pipeline without rebuilding from scratch
#
# YOU NEED ALL 5 ARTIFACTS READY BEFORE USING ANY FILE HERE:
#   - problem.md
#   - test.patch
#   - solution.patch
#   - Dockerfile
#   - repo_url.txt (URL + commit hash)
#
# ─────────────────────────────────────────────────────────────────────────────
# PICK YOUR MODE:
# ─────────────────────────────────────────────────────────────────────────────
#
# MODE A — Full Re-Audit (strong model, one pass)
#   → Run: Full.md
#   This chains the same gates as Steps 05 → 07 → 08 but in re-audit framing.
#   Use this when you want a thorough second opinion on a finished challenge.
#
# MODE B — Atomic Checks (any model, cheap-model safe)
#   Run each file independently for one focused question:
#   → Check_Problem.md   — P1-P7 audit on problem.md only
#   → Check_Tests.md     — T1-T8 audit on test.patch only
#   → Check_Solution.md  — S1-S4 audit on solution.patch only
#   Each prompt has a single focused question with a binary PASS/FAIL output.
#   Cheap models handle these well because they can't hide vague answers.
#
# MODE C — Local Verification (no LLM needed)
#   → Local.md
#   A set of terminal commands that verify facts the LLM can only claim.
#   Docker doesn't lie. Run this after ANY model tells you something passed.
#   Use this to catch hallucinations from any model — cheap or expensive.
#
# MODE D — Fast Recursive P/T/S Loop (any point in the pipeline)
#   → Quick_PTS.md
#   Reads the repo docs first, then recursively hunts P/T/S violations and fixes them.
#   Use this for a fast iteration loop: cheaper than the full 05 Review, more targeted
#   than atomic checks. Good for mid-build sanity checks or after any edit.
#   "Read the docs and recursively try to find P/T/S violations and fix it."
#
# ─────────────────────────────────────────────────────────────────────────────
# RECOMMENDED CHEAP-MODEL WORKFLOW:
# ─────────────────────────────────────────────────────────────────────────────
#
#   1. Run Local.md first (terminal commands — no model needed)
#   2. Run Check_Problem.md → get PASS or FAIL list
#   3. Run Check_Tests.md   → get PASS or FAIL list per test
#   4. Run Check_Solution.md → get PASS or FAIL list per hunk
#   5. If any FAIL: go to the relevant 06_Fix/ file to repair
#   6. After repairs: run Local.md again to confirm
#
# Note: Steps 05, 07, 08 in the main pipeline are still the authoritative gates.
# This folder gives you a lightweight way to spot-check without running the full pipeline.
