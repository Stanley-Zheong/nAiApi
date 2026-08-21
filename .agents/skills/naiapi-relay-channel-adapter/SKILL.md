---
name: naiapi-relay-channel-adapter
description: Maintain nAiApi relay provider adapters, request/response conversions, streaming behavior, model mapping, and channel registration across upstream AI providers.
---

# Relay Channel Adapter

Use this skill for provider-specific relay work: adding or fixing a channel, request conversion, response parsing, streaming usage, model mapping, upstream headers, multipart/file conversion, realtime, Responses API, embeddings, rerank, audio, images, or compatibility between provider formats.

## Authority Map

Read the relevant authority before editing:

- `relay/channel/adapter.go` for the `Adaptor` interface contract.
- `relay/relay_adaptor.go` for API type and provider selection.
- `relay/common/relay_info.go` for relay metadata, model mapping, retry, and `streamSupportedChannels`.
- `controller/relay.go` and `middleware/distributor.go` for preflight, quota, retry, and channel selection.
- `relay/compatible_handler.go` for OpenAI-compatible stream options and request normalization.
- `dto/openai_request.go` plus provider-specific DTO files before changing relay request structs.
- The concrete `relay/channel/<provider>/` package and nearby tests for provider behavior.

## Workflow

1. Classify the relay format first: OpenAI-compatible chat, Claude messages, Gemini, Responses, realtime, image, audio, embedding, rerank, or task delegation.
2. Preserve explicit client JSON intent. Optional upstream scalar request fields must be pointer types with `omitempty`; `0`, `0.0`, and `false` sent by a client must survive re-marshal.
3. Use `common.*` JSON wrappers for marshal, unmarshal, decode, and JSON-string parsing.
4. Keep provider conversion as a narrow adapter concern. Do not push provider quirks into generic controllers unless the controller already owns that invariant.
5. For streaming, verify `info.SupportStreamOptions`, `request.Stream`, and usage chunk behavior together. Add the channel to `streamSupportedChannels` only after confirming upstream support.
6. When mapping models, verify both downstream model name and upstream model name behavior through `RelayInfo`.
7. Add or update focused tests in the provider package, DTO package, or relay helper package before broad test runs.

## Validation

Prefer targeted tests before full-suite commands:

```bash
GOCACHE=/private/tmp/naiapi-gocache go test ./dto ./relay/channel ./relay/channel/<provider> ./relay/common ./relay/helper
```

If the full suite fails due known current baselines, record the failing package and preserve the targeted pass evidence.

## Depth Controls

Repository-learning controls live in:

- `../naiapi-repo-router/references/skill-depth.md`
- `../naiapi-repo-router/references/skill-depth.json`
- `../naiapi-repo-router/evals/evals.json`
- `../naiapi-repo-router/scripts/audit_skill_depth.py`
