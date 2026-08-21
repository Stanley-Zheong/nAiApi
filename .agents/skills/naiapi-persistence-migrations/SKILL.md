---
name: naiapi-persistence-migrations
description: Maintain nAiApi GORM models, migrations, cache-backed persistence, and SQLite/MySQL/PostgreSQL compatibility.
---

# Persistence and Migrations

Use this skill for model structs, GORM queries, migrations, raw SQL, cache-backed model state, token/channel/task persistence, and any behavior that must work across SQLite, MySQL, and PostgreSQL.

## Authority Map

- `model/main.go` for database selection, compatibility flags, migration order, reserved column quoting, and boolean constants.
- `model/*.go` for GORM models, query helpers, cache invalidation, and transaction boundaries.
- `common/database.go`, Redis/cache helpers, and `pkg/cachex/` when state is cached.
- Migration tests in `model/*_test.go` before modifying schema-sensitive code.

## Compatibility Rules

- Prefer GORM abstractions over raw SQL.
- If raw SQL is unavoidable, branch for PostgreSQL versus MySQL/SQLite using `common.UsingPostgreSQL`, `common.UsingSQLite`, and `common.UsingMySQL`.
- Use `commonGroupCol`, `commonKeyCol`, `commonTrueVal`, and `commonFalseVal` for reserved column names and boolean SQL values.
- Do not use `AUTO_INCREMENT`, `SERIAL`, JSONB-only operators, MySQL-only functions, or `ALTER COLUMN` without cross-database fallback.
- Store JSON-like configuration in portable text fields unless an existing model already defines a compatible abstraction.
- Use common JSON wrappers in GORM scanner/valuer helpers and model serialization code.

## Validation

Run SQLite-backed model tests locally and keep external database tests guarded by DSN env vars:

```bash
GOCACHE=/private/tmp/naiapi-gocache go test ./model ./common ./pkg/cachex
```

For migrations, verify tests still skip cleanly without `TEST_MYSQL_DSN` and `TEST_POSTGRES_DSN`.

## Depth Controls

Repository-learning controls live in:

- `../naiapi-repo-router/references/skill-depth.md`
- `../naiapi-repo-router/references/skill-depth.json`
- `../naiapi-repo-router/evals/evals.json`
- `../naiapi-repo-router/scripts/audit_skill_depth.py`
