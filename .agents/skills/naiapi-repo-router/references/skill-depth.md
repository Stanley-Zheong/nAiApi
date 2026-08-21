# nAiApi Repository Skill Depth

This repository-learning delivery is scoped to fixed revision `23385008d692b6e6f255156cd820f231d3623d72` and the previous-12-month history window ending 2026-08-21.

Machine-readable files:

- [capability coverage](capability-coverage.json)
- [skill depth contract](skill-depth.json)
- [route evals](../evals/evals.json)
- [depth audit](../scripts/audit_skill_depth.py)
- [router skill](../SKILL.md)
- [project rules](../../../../AGENTS.md)

## Coverage Model

The router splits current repository work into these owned task families:

- [relay adapters](../../naiapi-relay-channel-adapter/SKILL.md): provider adapters, relay formats, DTO conversion, streaming, and model mapping.
- [billing and pricing](../../naiapi-billing-pricing/SKILL.md): model ratios, expression pricing, quota conversion, pre-consume, and settlement.
- [task lifecycle](../../naiapi-task-relay-lifecycle/SKILL.md): async task submit, remix, polling, conversion, billing adjustment, and refunds.
- [persistence](../../naiapi-persistence-migrations/SKILL.md): GORM models, migrations, cross-database SQL, and cache-backed state.
- [identity and security](../../naiapi-identity-security/SKILL.md): auth, tokens, OAuth, Passkey, SSRF, webhooks, and payment callbacks.
- [default frontend](../../naiapi-default-frontend/SKILL.md): `web/default` React routes, feature modules, state, styling, and frontend verification.
- [ops and diagnostics](../../naiapi-ops-release-diagnostics/SKILL.md): Docker, CI, release packaging, embedded dist, Electron, performance metrics, and logs.
- [classic sync](../../classic-to-default-sync/SKILL.md): commit-driven `web/classic` to `web/default` parity work.
- [i18n](../../i18n-translate/SKILL.md): `web/default` locale sync and translations.
- [shadcn/ui](../../shadcn-ui/SKILL.md): project-aware shadcn/ui component composition.

## History Evidence

The full code-read ledger is stored in the repository card evidence directory as `evidence/commit-code-read-coverage.jsonl`; it contains 1876 commit rows. The summary file records these major historical task counts: relay provider adapters 733, default frontend 601, admin dashboard API 582, billing/pricing 303, ops/release 224, identity/security 200, persistence/migrations 194, async task relay 178, diagnostics 172, frontend i18n 90, payment/subscription 48, docs/local skills 22, and frontend classic 10.

This depth contract does not treat deleted legacy paths as active capabilities. Historical rows marked `retired-path-supporting-evidence` remain evidence, not owners.

## Validation Level

All delivered skills are declared at L1: current-state implementation anchors, deterministic route probes, skill validation, and targeted test/baseline evidence. No L2 or L3 hidden replay claim is made in this delivery.

Known fixed-revision baselines are recorded in the repository card evidence directory:

- `evidence/go-test-targeted.txt`: targeted Go package tests passed.
- `evidence/go-test-all-baseline.md`: `go test ./...` fails at the fixed revision because embedded frontend dist files are absent, one controller realtime test hits sandbox bind limits, and two relay helper/provider expectations fail.
- `evidence/web-default-typecheck.txt`: `bun run typecheck` fails at the fixed revision due local missing Rsbuild/TanStack/Node types and a `baseUrl` deprecation warning.

## Drift Rules

When a future change adds a new provider family, task platform, billing expression primitive, security boundary, frontend framework path, or release target, update both `capability-coverage.json` and `skill-depth.json`, add a representative eval in `evals/evals.json`, and rerun `audit_skill_depth.py`.
