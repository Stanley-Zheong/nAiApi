---
name: naiapi-repo-router
description: Route nAiApi maintenance requests to the right repo-local skill for relay, billing, tasks, persistence, identity, frontend, release, diagnostics, i18n, and UI work.
---

# nAiApi Repository Router

Use this skill first for work in this repository when the request does not name a more specific repo-local skill. Route by owned business task family, not by whichever file name appears first.

## Required Preflight

1. Read `AGENTS.md` and `CLAUDE.md` before changing repository code.
2. If touching `web/default`, read `web/default/AGENTS.md` and use Bun commands.
3. If touching tiered or dynamic billing, read `pkg/billingexpr/expr.md` before editing.
4. Preserve protected project identity references and metadata exactly.
5. Use `common.Marshal`, `common.Unmarshal`, `common.UnmarshalJsonStr`, `common.DecodeJson`, and `common.GetJsonType` for JSON operations in business code.
6. Keep database code compatible with SQLite, MySQL, and PostgreSQL.
7. For upstream relay request DTOs, preserve explicit zero or false values with pointer fields and `omitempty`.
8. For new channels, confirm provider support for `StreamOptions`; add supported channels to `streamSupportedChannels`.

## Route Table

| Trigger | Primary Skill |
| --- | --- |
| Provider channel, relay format, `/v1/chat/completions`, `/v1/responses`, realtime, image/audio/embedding/rerank conversion, model mapping, streaming usage | `naiapi-relay-channel-adapter` |
| Model ratio, group ratio, tiered expression pricing, token/quota accounting, pre-consume, settlement, pricing UI/API | `naiapi-billing-pricing` |
| Async video/music/image task submit, remix, poll, fetch, task refund/recalculation, `OtherRatios`, task webhooks | `naiapi-task-relay-lifecycle` |
| GORM model, migration, cross-database raw SQL, cache/schema behavior, SQLite/MySQL/PostgreSQL compatibility | `naiapi-persistence-migrations` |
| User/session/token auth, OAuth, passkeys, secure verification, SSRF, payment callback trust boundary, webhook/download fetch security | `naiapi-identity-security` |
| `web/default` React app, TanStack Router routes, Base UI, Tailwind, frontend stores/API clients, admin console views | `naiapi-default-frontend` |
| User supplies a commit and asks to compare or port `web/classic` changes into `web/default` | `classic-to-default-sync` |
| Frontend translations, missing locale keys, i18next files, `bun run i18n:*` | `i18n-translate` |
| shadcn/ui components, registries, theming, or component composition inside `web/default` | `shadcn-ui` |
| Docker, compose, CI, release packaging, embedded frontend dists, Electron, perf metrics, logs, operational diagnostics | `naiapi-ops-release-diagnostics` |

## Boundary Rules

- If the task asks for Company Jarvis changes, write nothing there and return to the coordinator.
- If the task asks to mark repository learning complete, only the coordinator may set `completed`; the worker delivery status is `delivered-awaiting-coordinator-verification`.
- If a change crosses multiple task families, start with the skill owning the highest-risk state transition, then explicitly note secondary skills.

## Depth Controls

Repository-learning coverage and proof are tracked in:

- `references/capability-coverage.json`
- `references/skill-depth.md`
- `references/skill-depth.json`
- `evals/evals.json`
- `scripts/audit_skill_depth.py`

Run the audit after editing repo-local skills:

```bash
python3 .agents/skills/naiapi-repo-router/scripts/audit_skill_depth.py --repo . --router naiapi-repo-router --skills-root .agents/skills
```
