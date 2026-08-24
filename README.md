# Olympus Challenge — Copy-Paste Pipeline & Workflow

A streamlined, step-by-step copy-paste pipeline and guide for building, reviewing, validating, and tuning benchmark coding challenges for Olympus / ShipdAI.

## Quick Links

- 📖 **[Copy-Paste Pipeline Guide](CopyPasteGuide/README.md)** — Complete step-by-step walkthrough and prompt files.
- 📁 **[Pipeline Prompts](CopyPasteGuide/)**:
  - [`1_Build/`](CopyPasteGuide/1_Build/) — Steps 01 to 04B: Repo discovery, seed selection, problem/tests generation, Docker validation, reference solution.
  - [`2_Review/`](CopyPasteGuide/2_Review/) — Steps 05 to 08: Full review, targeted fixes, coverage check, and anti-shortcut mutations.
  - [`3_Validate/`](CopyPasteGuide/3_Validate/) — Steps 09 to 10: Blind agent runs and final pre-submission checklist.
  - [`4_Extras/`](CopyPasteGuide/4_Extras/) — Steps 11 to 12: Difficulty tuning and rapid re-check tools.
- 📝 **[Original Notes & Reference Docs](Original_Notes/)** — Underlying reference documentation and interface notes.

## Structure

```
.
├── CopyPasteGuide/
│   ├── 1_Build/          # 01_Find_Repo, 02_Pick_Seed, 03A_Build, 03B_Check_Problem, 03C_Docker_Matrix, 04_Build_Solution, 04B_Docker_Matrix
│   ├── 2_Review/         # 05_Review, 06_Fix, 07_Coverage, 08_Anti_Shortcut
│   ├── 3_Validate/       # 09_Blind_Test, 10_Submit
│   ├── 4_Extras/         # 11_Tune, 12_Re_Check
│   └── README.md         # Full Pipeline Walkthrough & Reference
├── Original_Notes/       # Platform documentation and raw reference notes
└── .gitignore
```
