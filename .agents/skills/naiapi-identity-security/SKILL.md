---
name: naiapi-identity-security
description: Maintain nAiApi auth, token/session trust boundaries, OAuth, passkeys, secure verification, SSRF protections, payment callbacks, and webhook security.
---

# Identity and Security

Use this skill for login/session auth, API token auth, OAuth providers, passkeys, secure verification, admin auth boundaries, SSRF-sensitive fetch/proxy/download behavior, payment callback validation, webhooks, and user-notification URLs.

## Authority Map

- `middleware/auth.go`, `middleware/header_nav.go`, and rate-limit middleware for request trust boundaries.
- `controller/user.go`, `controller/passkey.go`, `controller/secure_verification.go`, and `service/passkey/` for account and step-up verification flows.
- `oauth/` for provider-specific OAuth behavior.
- `model/token.go`, `model/user.go`, `model/passkey.go`, `model/topup.go`, and `model/subscription.go` for persisted trust state.
- `common/ssrf_protection.go`, `service/http_client.go`, `service/download.go`, `service/webhook.go`, `service/user_notify.go`, and video proxy controllers for outbound URL validation.
- Payment callback controllers for Stripe, EPay, Creem, Waffo, and Waffo Pancake.

## Security Invariants

- Token auth must respect token status, user status, group/model limits, read-only routes, and cache invalidation after sensitive user changes.
- `New-Api-User` header handling is an internal trust boundary; never expose it as user-controlled authentication.
- Passkey verification relies on WebAuthn origin/RPID checks plus short-lived session markers for secure verification.
- Outbound URL fetches and proxying must pass `ValidateURLWithFetchSetting` whenever user-configured URLs are involved.
- Payment callbacks must validate both payment identity and expected `PaymentProvider`; `PaymentMethod` alone is not enough to bind an order to a gateway.
- Preserve localized but non-leaking error behavior; log diagnostic details server-side when existing code does so.

## Validation

Run focused security and model tests:

```bash
GOCACHE=/private/tmp/naiapi-gocache go test ./middleware ./model ./service ./controller
```

When controller package tests hit current realtime sandbox limits, record the baseline and keep narrower packages passing.

## Depth Controls

Repository-learning controls live in:

- `../naiapi-repo-router/references/skill-depth.md`
- `../naiapi-repo-router/references/skill-depth.json`
- `../naiapi-repo-router/evals/evals.json`
- `../naiapi-repo-router/scripts/audit_skill_depth.py`
