## Expansion — Challenge is Below Scope

The threshold check returned **🔴 NEEDS EXPANSION**.

This is handled by the existing workflow step:

**→ Use `4_Extras/11_Tune/Expand.md`**

Paste the findings from `1_Check.md` into that prompt. It will grow the challenge scope through natural same-workflow behaviors without breaking existing tests.

After expanding, re-run the Docker matrix (`2_Review/06_Fix/4_Docker_Matrix.md`) and then re-run `1_Check.md`.

---

### What counts as natural expansion

✅ Valid: persistence, execution, validation, serialization, CLI/API variants, recovery, restart paths, lifecycle/ordering behavior, error paths, concurrency behavior.

🚫 Invalid: API inventory, schema plumbing, wrappers, convenience helpers, unrelated parity, test-only scope.

---

### The loop

```
1_Check.md → 🔴 NEEDS EXPANSION
    → 4_Extras/11_Tune/Expand.md  (expand scope)
    → 2_Review/06_Fix/4_Docker_Matrix.md  (verify)
    → 1_Check.md  (re-audit)
```
