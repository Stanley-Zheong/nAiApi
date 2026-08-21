---
name: naiapi-task-relay-lifecycle
description: Maintain nAiApi async task relay submission, remix, polling, provider task adapters, task billing adjustments, refunds, and OpenAI video compatibility.
---

# Task Relay Lifecycle

Use this skill for async task platforms such as Midjourney, Suno, Sora, Kling, Vidu, Jimeng, Doubao, Gemini video, Vertex video, Ali video, Hailuo, video proxying, task polling, remix/continue, and task billing settlement.

## Authority Map

- `relay/channel/adapter.go` for the `TaskAdaptor` contract.
- `relay/relay_task.go` for submit flow, origin task resolution, pre-consume, submit-time adjustment, and task log creation.
- `relay/channel/task/` for provider-specific task adapters.
- `service/task_billing.go` and `service/task_polling.go` for task quota recalculation and terminal-state settlement.
- `model/task.go` for task persistence, private data, billing context, and retry/poll state.
- `controller/task_video.go`, `controller/video_proxy.go`, `controller/video_proxy_gemini.go`, and `router/video-router.go` for public task/video surfaces.

## Lifecycle Contract

1. `ValidateRequestAndSetAction` determines platform action and rejects unsupported shapes before billing.
2. `ResolveOriginTask` owns remix/continuation lookup, locked-channel behavior, and inherited `OtherRatios`.
3. `EstimateBilling` must run before pre-consume when request metadata changes duration, size, resolution, input mode, or discount.
4. `ForcePreConsume` disables trust bypass for async tasks that must reserve full quota before upstream submission.
5. `AdjustBillingOnSubmit` may refine `OtherRatios` from upstream response data before task persistence.
6. `AdjustBillingOnComplete` may replace the final quota at terminal state; failure paths must refund or recalculate through existing services.
7. Provider adapters must parse terminal task state and expose OpenAI-compatible video output only through the established converter.

## Validation

Run focused task tests and relevant provider tests:

```bash
GOCACHE=/private/tmp/naiapi-gocache go test ./service ./relay/channel/task/... ./relay/channel
```

For SSRF or fetch/proxy changes, also use `naiapi-identity-security`.

## Depth Controls

Repository-learning controls live in:

- `../naiapi-repo-router/references/skill-depth.md`
- `../naiapi-repo-router/references/skill-depth.json`
- `../naiapi-repo-router/evals/evals.json`
- `../naiapi-repo-router/scripts/audit_skill_depth.py`
