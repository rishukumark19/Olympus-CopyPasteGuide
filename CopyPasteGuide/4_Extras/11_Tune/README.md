# STEP 11 — Difficulty Tuning
#
# Use this folder AFTER Step 09 (Blind Testing) reveals the challenge is too easy,
# too hard, or the LOC is out of range.
#
# HOW TO PICK THE RIGHT FILE:
#
#   Step 09 → 5/5 pass (too easy)           → Harder.md
#   Step 09 → 0/5 pass (too hard/unfair)    → Easier.md
#   Step 09 → 0/5, agents understood task   → After_Runs.md (add surgical hints)
#   Step 09 → some pass, missing behaviors  → After_Runs.md (close the gaps)
#   Critic flags LOC < 250 or scope thin    → Expand.md
#   Critic flags LOC > 600 or padding found → Reduce.md
#   Step 08 found surviving mutations       → Harder.md (also works mid-pipeline)
#
# WHAT HAPPENS AFTER TUNING:
#
#   After Harder    → re-run Docker matrix → re-run Step 08 → re-run Step 09
#   After Easier    → re-run Docker matrix → re-run Step 05 → re-run Step 09
#   After After_Runs → re-run Step 03B (problem check) → re-run Step 09
#   After Expand → restart from Step 03A (new problem.md + test.patch needed)
#   After Reduce → re-run Step 04 (trim solution.patch) → re-run Step 05
#
# STALENESS NOTE:
#   Any change to test.patch or solution.patch from this folder
#   makes ALL Step 09 rollouts stale. Re-run Step 09 before submitting.
