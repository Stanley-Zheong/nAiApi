---
name: naiapi-default-frontend
description: Maintain nAiApi web/default React 19 admin frontend, routes, stores, Base UI/Tailwind components, i18n use, and Rsbuild/Bun workflows.
---

# Default Frontend

Use this skill for `web/default` application work: dashboard views, admin console pages, auth screens, channel/model/pricing/task/log pages, API clients, stores, route definitions, Base UI/Tailwind components, and frontend build/typecheck behavior.

## Required Reading

- `web/default/AGENTS.md` for frontend conventions.
- `web/default/package.json` for Bun scripts and dependency context.
- `web/default/src/routes/` before changing navigation or page ownership.
- `web/default/src/features/<feature>/` before touching a feature.
- `web/default/src/i18n/` and `i18n-translate` if user-visible strings are added or changed.

## Implementation Rules

- Use React 19, TypeScript, TanStack Router, Base UI, Tailwind CSS, Zustand, and existing project helpers.
- Do not import Semi Design into `web/default`.
- Do not destructure props unless the local file pattern does so for a concrete reason.
- Avoid nested ternaries and keep component state transitions explicit.
- Use `useTranslation()` and `t('English key')` for visible text.
- Build data-heavy admin screens for scanning and repeated operation, with stable dimensions and no nested UI cards.
- Prefer existing shared UI components in `src/components/ui` and feature-local patterns before adding new abstractions.

## Validation

For TypeScript changes, run:

```bash
cd web/default && bun run typecheck
```

The repository-learning baseline recorded that this command currently fails at the fixed revision because local frontend types/tooling are incomplete; keep new changes from adding additional failures and record the observed baseline.

## Depth Controls

Repository-learning controls live in:

- `../naiapi-repo-router/references/skill-depth.md`
- `../naiapi-repo-router/references/skill-depth.json`
- `../naiapi-repo-router/evals/evals.json`
- `../naiapi-repo-router/scripts/audit_skill_depth.py`
