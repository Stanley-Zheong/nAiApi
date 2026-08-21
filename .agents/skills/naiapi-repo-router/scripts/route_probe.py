#!/usr/bin/env python3
"""Deterministic route checks for the nAiApi repository router eval suite."""

from __future__ import annotations

import argparse
import json
from pathlib import Path


def route(prompt: str) -> str:
    text = prompt.lower()
    if "which repo-local skill" in text or "decide which" in text or "spans relay and billing" in text:
        return "naiapi-repo-router"
    if "shadcn" in text:
        return "shadcn-ui"
    if "classic" in text and "default" in text and "commit" in text:
        return "classic-to-default-sync"
    if "i18n" in text or "locale" in text or "missing keys" in text or "untranslated" in text:
        return "i18n-translate"
    if (
        "docker" in text
        or "release" in text
        or "frontend dist" in text
        or "embedded dist" in text
        or "startup diagnostics" in text
        or "performance metrics" in text
    ):
        return "naiapi-ops-release-diagnostics"
    if "passkey" in text or "ssrf" in text or "webhook" in text or "payment" in text or "token auth" in text:
        return "naiapi-identity-security"
    if "sqlite" in text or "mysql" in text or "postgres" in text or "migration" in text or "gorm" in text:
        return "naiapi-persistence-migrations"
    if "task" in text or "sora" in text or "otherratios" in text or "forcepreconsume" in text or "polling" in text:
        return "naiapi-task-relay-lifecycle"
    if "billing" in text or "pricing" in text or "quota" in text or "pre-consume" in text or "settlement" in text:
        return "naiapi-billing-pricing"
    if "provider" in text or "adapter" in text or "streamoptions" in text or "streaming" in text or "model mapping" in text:
        return "naiapi-relay-channel-adapter"
    if "web/default" in text or "react" in text or "tanstack" in text or "tailwind" in text:
        return "naiapi-default-frontend"
    return "naiapi-repo-router"


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--evals", required=True, type=Path)
    parser.add_argument("--out", type=Path)
    args = parser.parse_args()

    payload = json.loads(args.evals.read_text(encoding="utf-8"))
    results = []
    ok = True
    for row in payload.get("evals", []):
        actual = route(str(row.get("prompt", "")))
        expected = row.get("expected_route", [])
        forbidden = row.get("forbidden_routes", [])
        passed = actual in expected and actual not in forbidden
        ok = ok and passed
        results.append(
            {
                "id": row.get("id"),
                "expected_route": expected,
                "actual_route": actual,
                "forbidden_routes": forbidden,
                "status": "passed" if passed else "failed",
            }
        )

    report = {"status": "ok" if ok else "failed", "results": results}
    text = json.dumps(report, indent=2) + "\n"
    if args.out:
        args.out.write_text(text, encoding="utf-8")
    print(text, end="")
    return 0 if ok else 1


if __name__ == "__main__":
    raise SystemExit(main())
