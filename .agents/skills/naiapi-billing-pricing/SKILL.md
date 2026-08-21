---
name: naiapi-billing-pricing
description: Maintain nAiApi model pricing, tiered billing expressions, quota conversion, pre-consume, settlement, and pricing display behavior.
---

# Billing and Pricing

Use this skill for model ratios, group ratios, tiered or dynamic billing expressions, token/quota accounting, pre-consume/refund/settlement behavior, pricing APIs, sync endpoints, and frontend pricing editors or displays.

## Required Reading

Read `pkg/billingexpr/expr.md` before changing any tiered or dynamic billing code.

Then read the authority that matches the change:

- `pkg/billingexpr/` for expression parsing, compilation, runtime variables, and tests.
- `relay/helper/price.go` for model price resolution.
- `service/billing_session.go`, `service/text_quota.go`, `service/quota.go`, and `service/tiered_settle.go` for quota lifecycle.
- `setting/ratio_setting/`, `setting/billing_setting/`, and `setting/model_setting/` for pricing configuration.
- `controller/pricing.go` and `controller/ratio_sync.go` for API/admin surfaces.
- `web/default/src/features/pricing/` and related stores when pricing behavior reaches the default frontend.

## Non-Negotiable Semantics

- `p` and `c` are automatically excluded from chargeable token totals in expression mode; do not double-count them in manual formulas.
- `len` is the normalized tier selector for tiered conditions when expressions depend on request or completion length.
- Pre-consume must snapshot enough pricing context for settlement to reproduce the intended formula even if settings drift later.
- Output-only or completion-only models are not free just because prompt price is zero; verify group ratio, completion ratio, and tier expression behavior together.
- Convert money or ratio outputs through the same quota helper path used by existing code; avoid ad hoc float rounding.
- Do not add database-specific storage for expression data; use cross-database compatible text/JSON storage with common JSON wrappers.

## Validation

Run focused billing tests first:

```bash
GOCACHE=/private/tmp/naiapi-gocache go test ./pkg/billingexpr ./service ./relay/helper ./setting/billing_setting ./setting/ratio_setting
```

For changes that affect task pricing, also use `naiapi-task-relay-lifecycle`.

## Depth Controls

Repository-learning controls live in:

- `../naiapi-repo-router/references/skill-depth.md`
- `../naiapi-repo-router/references/skill-depth.json`
- `../naiapi-repo-router/evals/evals.json`
- `../naiapi-repo-router/scripts/audit_skill_depth.py`
