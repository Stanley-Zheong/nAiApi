<div align="center">

![new-api](/web/default/public/logo.png)

# nAiApi

🍥 **新一代大模型网关与 AI 资产管理系统**（基于 [QuantumNous/new-api](https://github.com/QuantumNous/new-api) 的 fork）

</div>

## 📌 项目速览

| 维度 | 内容 |
|---|---|
| **项目类型** | 商业化项目 |
| **关注等级** | 高 |
| **阻塞点** | 无硬阻塞。`go build ./...` 会因 `web/classic/dist` 前端 embed 未构建而失败（稳定门禁用 `go build ./relay/...` + 单包 `go test`）；cohere/moonshot/baidu 等部分上游适配器仍为 `TODO implement me`；`relay/channel/claude` 此前的 3 个失败测试已修复并合回 `main`（提交 `278cfbf1`，HES-97），`go build ./relay/...` + 单包 `go test ./relay/channel/claude/` 现为绿色 |
| **恢复条件** | 随时可继续开发；需全量构建前先构建前端产物（`web/default` 用 Bun） |
| **整理日期** | 2026-06-08 |

## Multica 项目信息

| 字段 | 内容 |
|---|---|
| **项目名称** | nAiApi |
| **Multica 项目 ID** | `42e5fc42-30cd-4505-a884-9f9237fb78d8` |
| **负责人** | `hgpt`（agent，`a58a7982-e2e0-4cb9-8710-d2e6cdab0351`） |
| **GitHub** | `https://github.com/Stanley-Zheong/nAiApi.git` |
| **本地目录** | `/Users/laosanzheong/Documents/codebases/nAiApi` |
| **运行环境** | Go 1.25.1 + Gin/GORM；前端为 React/TypeScript，包管理器用 Bun；数据库需兼容 SQLite、MySQL、PostgreSQL，可选 Redis，支持 Docker 部署。 |

> 本仓库 `nAiApi` 是 **New API**（QuantumNous/new-api）的下游 fork。其内核、品牌、Go 模块路径（`github.com/QuantumNous/new-api`）与 New API 完全一致，相关署名与品牌信息受项目策略保护、不可移除。
> 多语言版上游 README：[简体中文](./README.zh_CN.md) · [繁體中文](./README.zh_TW.md) · [English](./README.en.md) · [Français](./README.fr.md) · [日本語](./README.ja.md)。
> 官方文档：<https://docs.newapi.pro/>。

---

## 项目目标

nAiApi 是一个用 Go 编写的 **AI API 网关 / 代理**。它把 40+ 家上游 AI 服务商（OpenAI、Claude、Gemini、Azure、AWS Bedrock、DeepSeek、通义千问、智谱、文心、Mistral、Cohere 等）聚合到**统一的 API 之后**，并在其上提供用户管理、令牌（API Key）分组、计费与额度核算、限流、渠道路由与可视化管理后台。

它要解决的核心问题：

- **接口碎片化**：不同厂商的请求/响应格式、鉴权方式、流式协议各不相同。网关用一套兼容 OpenAI 的协议（同时支持 Claude Messages、Google Gemini 等原生格式）屏蔽这些差异。
- **统一的资产与成本治理**：把分散的上游 Key、账号、模型权限集中起来，做组织级的按次 / 按量 / 缓存命中计费、额度分配与用量分析。
- **可私有化、可分发**：支持 Docker、裸机、宝塔面板部署，并提供 Electron 桌面端封装。

服务对象：需要做大模型聚合、组织级鉴权、多模型管理、用量分析、成本核算与私有化部署的团队与个人（仅限**合法、已获授权**的场景，详见下文「约束」）。

---

## 核心功能

- **统一多厂商网关**：`relay/channel/` 下有 40+ 个服务商适配器，统一对外暴露兼容 OpenAI 的接口。
- **多种 API 格式**：Chat Completions、OpenAI Responses、OpenAI Realtime（含 Azure）、Claude Messages、Google Gemini、Embeddings、Image、Audio、Video、Rerank（Cohere / Jina）、Midjourney-Proxy、Suno、Dify ChatFlow。
- **格式互转**：OpenAI ⇄ Claude Messages、OpenAI → Gemini、Gemini → OpenAI（仅文本，暂不支持 function calling）、思维链转正文（thinking-to-content）。OpenAI ⇄ OpenAI Responses 互转**开发中**。
- **推理力度（Reasoning Effort）**：对 OpenAI（`o3-mini-*`、`gpt-5-*`）、Claude thinking、Gemini thinking 系列模型支持 high/medium/low 等档位与 thinking budget。
- **智能路由**：渠道加权随机、失败自动重试、用户级模型限流。
- **计费与额度**：内部充值与额度分配（EPay、Stripe、Creem、Waffo 等支付渠道），组织级按次/按量/缓存命中成本核算；表达式驱动的分层动态计费（见 `pkg/billingexpr/`）。
- **鉴权与登录**：JWT、WebAuthn/Passkeys，以及 GitHub / Discord / LinuxDO / Telegram / OIDC 等 OAuth 登录。
- **管理后台**：现代化 React 前端、可视化数据看板、令牌分组与模型限制、用户与权限管理。
- **多语言**：后端 en/zh；前端 zh/en/fr/ru/ja/vi。
- **可观测性**：Pyroscope 持续性能剖析、pprof、错误日志开关。
- **桌面端**：`electron/` 提供 Windows / macOS / Linux 的系统托盘桌面封装。

---

## 技术架构

### 技术栈

- **后端**：Go 1.25（`go.mod` 声明 `go 1.25.1`）、Gin Web 框架、GORM v2 ORM。
- **前端**：默认主题 React 19 + TypeScript + Rsbuild + Base UI + Tailwind CSS；经典主题 React 18 + Vite + Semi Design。前端包管理器统一用 **Bun**。
- **数据库**：SQLite / MySQL(≥5.7.8) / PostgreSQL(≥9.6)，三者必须同时兼容。
- **缓存**：Redis（go-redis）+ 内存缓存。
- **鉴权**：JWT、WebAuthn/Passkeys、多家 OAuth。
- **打包**：Docker / Docker Compose；前端构建产物通过 Go `embed` 内嵌进二进制（`main.go` 内嵌 `web/default/dist` 与 `web/classic/dist`）。

### 目录结构

```
main.go        — 入口，初始化资源、embed 前端产物、启动 Gin
router/        — HTTP 路由（API / relay / dashboard / web）
controller/    — 请求处理器（约 71 个）
service/       — 业务逻辑
model/         — 数据模型与 DB 访问（GORM）
relay/         — AI API 中继/代理
  relay/channel/ — 各厂商适配器（openai/ claude/ gemini/ aws/ … 共 40+）
middleware/    — 鉴权、限流、CORS、日志、分发
setting/       — 配置管理（ratio / model / operation / system / performance / 计费 / 支付）
common/        — 共享工具（JSON 封装、加密、Redis、env、限流等）
dto/           — 请求/响应传输结构
constant/      — 常量（API 类型、渠道类型、context key）
types/         — 类型定义（中继格式、文件来源、错误）
oauth/         — OAuth 各 provider 实现
pkg/           — 内部包（cachex、ionet、billingexpr、perf_metrics 等）
i18n/          — 后端国际化（go-i18n，en/zh）
electron/      — Electron 桌面端封装
web/default/   — 默认前端（React 19 / Rsbuild / Base UI / Tailwind）
web/classic/   — 经典前端（React 18 / Vite / Semi Design）
docs/          — 部署/渠道/翻译术语表/OpenAPI 文档
bin/           — DB 迁移 SQL 与脚本
```

### 模块划分与数据流

分层架构：**Router → Controller → Service → Model**；中继请求走 **Router → middleware（鉴权/分发） → relay → relay/channel/<厂商>适配器 → 上游**。

```
客户端 ──HTTP──▶ Gin Router ──▶ middleware（鉴权 / 限流 / 渠道分发）
                                  │
                                  ├─ 管理类请求 ─▶ controller ─▶ service ─▶ model ─▶ DB(SQLite/MySQL/PG)
                                  │
                                  └─ 模型推理请求 ─▶ relay ─▶ relay/channel/<provider> 适配器
                                          │（格式转换、计费预扣/结算、流式透传）
                                          └────────────────────────────▶ 上游 AI 服务商
```

> 工程约定详见 `CLAUDE.md` / `AGENTS.md`，关键规则包括：JSON 必须走 `common/json.go` 封装、DB 代码三库兼容、前端用 Bun、新增渠道需确认 `StreamOptions`、中继请求 DTO 用指针保留显式零值、分层计费需先读 `pkg/billingexpr/expr.md`。

---

## 当前进展

- **基线**：fork 跟随上游 New API 主线，已合并大量上游 fix/feat（如渠道分组过滤修复、公共模块导航访问控制、API Key 额度校验、用量日志过滤、上游请求 ID 追踪、付费功能合规确认等，见 `git log`）。
- **安全加固（fork 自有，2026-05-19）**：见 `SECURITY_FIX_SUMMARY.md` 与 `SECURITY.md`。已完成前端依赖升级（axios 等，修复多项漏洞）、CORS 收敛（新增 `ALLOWED_ORIGINS`、限定允许头部）等，自评安全分从 7.2/10 提升至约 8.8/10。
- **前端**：默认主题（`web/default`）持续迭代，经典主题（`web/classic`）作为兼容保留。
- **桌面端**：`electron/` 已可在 macOS/Windows/Linux 以系统托盘形式运行（依赖预编译 Go 二进制）。
- **Realtime WebSocket（2026-06-12）**：握手允许 `openai-beta.realtime-v1` 子协议，兼容 OpenAI Realtime 客户端的 `Sec-WebSocket-Protocol` 协商。
- **自动化任务**：`cross-loop.tasks.yaml`（本地、未纳入版本控制的辅助文件）记录了基于 cross-loop 的批量改造计划，主要面向 `relay/channel/*` 适配器的同构改造。

---

## TODO / 待办

来自代码 TODO、文档与构建说明：

- ~~**Claude 文件内容处理测试失败（已知）**：`relay/channel/claude` 有 3 个失败测试（`...IgnoresUnsupportedFileContent` / `...SupportsPDFFileContent` / `...ConvertsTextFileContentToText`），`cross-loop.tasks.yaml` 将其列为待修复首要项。~~ ✅ 已修复并合回 `main`（提交 `278cfbf1`，HES-97）：在 OpenAI→Claude 转换中新增 `dto.ContentTypeFile` 分支，PDF 转 `document`、纯文本解码为 `text`、其余不支持类型跳过；三个测试现已通过。
- **格式互转补全**：OpenAI ⇄ OpenAI Responses 互转**开发中**；Gemini → OpenAI 暂不支持 function calling（README 标注）。
- **渠道适配器 TODO**：`relay/channel/cohere`、`moonshot`、`baidu` 等适配器中存在多处 `TODO implement me` / `TODO: fix this`（如 cohere 流式 handler）；`relay/claude_handler.go:95` 有「临时处理」标记。
- ~~**Electron 从源码构建文档缺失**：`electron/README.md` 中「Option B: Build from source」一节仅为 `TODO`。~~ ✅ 已补齐：`electron/README.md` 现含完整的「先构建前端（`make build-all-frontends`）→ `CGO_ENABLED=1 go build` → `electron/build.sh` 一键打包」说明，并修复了 `build.sh` 中失效的前端构建路径。
- **依赖漏洞收尾（推断）**：`SECURITY_FIX_SUMMARY.md` 指出仍有 minimist、dompurify 等间接依赖漏洞待上游升级后清理。
- **测试与构建门禁修复（推断）**：见下「注意事项」中关于 `go build ./...` 与若干 `go test` 失败的说明，待补齐前端产物/修复预存在用例后才能形成稳定绿色门禁。

> 上游功能路线图（如更多模型与接口支持）以官方文档为准：<https://docs.newapi.pro/>。

---

## 注意事项

- **`go build ./...` 会失败**：`main.go` 通过 `//go:embed` 内嵌 `web/default/dist` 与 `web/classic/dist`。**未先构建前端**时整仓库编译会失败。可用的稳定门禁是 `go build ./relay/...`。完整构建请用 `make all`（先 `bun install && bun run build` 构建两套前端，再 `go run main.go`）。
- **预存在的测试/检查噪音**：`go vet ./relay/...` 有预存在的 "unreachable code" 告警；`go test ./relay/...`、`relay/helper` 存在预存在失败用例（`relay/channel/claude` 已于 HES-97 修复并合回 `main`，现为绿色）。可靠绿色的范围是部分适配器（如 `relay/channel/claude`、`relay/channel/aws/...`、`gemini`、`minimax`）。
- **前端必须用 Bun**：`bun install` / `bun run dev` / `bun run build` / `bun run i18n:*`，工作目录为 `web/default/`（经典主题为 `web/classic/`）。
- **三库兼容硬约束**：所有 DB 代码必须同时兼容 SQLite / MySQL / PostgreSQL。优先用 GORM 抽象；不可避免的原生 SQL 需处理列引用符、布尔值、保留字（`group`/`key`）差异，SQLite 不支持 `ALTER COLUMN`。详见 `CLAUDE.md` Rule 2 与 `model/main.go`。
- **JSON 统一封装**：业务代码禁止直接调用 `encoding/json` 的 marshal/unmarshal，必须走 `common.Marshal`/`common.Unmarshal` 等封装（`CLAUDE.md` Rule 1）。
- **中继请求 DTO 零值**：从客户端解析再转发上游的可选标量字段必须用指针 + `omitempty`，以区分「未传」与「显式传 0/false」（`CLAUDE.md` Rule 6）。
- **运行环境**：Go 1.25+；远程库 MySQL ≥ 5.7.8 或 PostgreSQL ≥ 9.6；本地默认 SQLite（Docker 需挂载 `/data`）。
- **多机部署**：必须设置 `SESSION_SECRET`（否则登录态不一致）；共享 Redis 必须设置 `CRYPTO_SECRET`（否则数据无法解密）。
- **保护性信息**：`new-api` / `QuantumNous` 相关品牌、署名、模块路径、文档与配置受 `CLAUDE.md` Rule 5 保护，**禁止删除或重命名**。

---

## 约束

- **技术约束**
  - 三数据库（SQLite/MySQL/PostgreSQL）必须同时兼容；禁止使用无跨库兜底的 MySQL/PostgreSQL 专属函数与类型（如裸 `JSONB`、`GROUP_CONCAT`、PG JSON 操作符）。
  - JSON 操作、计费表达式、渠道 `StreamOptions`、请求 DTO 零值等须遵循 `CLAUDE.md` 规则。
  - 前端工具链固定为 Bun + Rsbuild/Vite；后端入口依赖前端构建产物（embed）。
- **业务/合规约束**
  - 本项目仅用于**合法且已获授权**的 AI API 网关、组织级鉴权、多模型管理、用量分析、成本核算与私有化部署场景。
  - 使用者须合法获取上游 API Key、账号、模型服务与接口权限，并遵守上游服务条款及所在司法辖区的法律法规。
  - 面向公众提供生成式 AI 服务时，须自行完成备案、许可、内容安全、实名、日志留存、税务、支付与上游授权等义务。
- **已知限制**
  - 整仓库 `go build ./...` 在未构建前端时不可用；部分 `go test` 与适配器存在预存在失败/未实现项（见 TODO）。
  - Gemini → OpenAI 转换仅支持文本；OpenAI ⇄ Responses 互转尚在开发。
  - 部分间接前端依赖仍有未修复漏洞，待上游升级。

---

## 快速开始

```bash
# Docker Compose（推荐）
git clone https://github.com/QuantumNous/new-api.git   # 或本 fork 仓库
cd new-api
nano docker-compose.yml      # 按需修改配置
docker-compose up -d         # 默认端口 3000

# 本地开发（需 Go 1.25+ 与 Bun）
make dev        # docker 起依赖(docker-compose.dev.yml) + 前端 dev server
# 或完整构建后端嵌入前端运行：
make all        # 构建两套前端并 go run main.go
```

部署完成后访问 `http://localhost:3000`。更多部署方式（Docker 命令、宝塔面板、环境变量）见 `docs/installation/` 与官方文档 <https://docs.newapi.pro/>。

---

## 许可与署名

- 许可证见 `LICENSE`，第三方依赖许可见 `THIRD-PARTY-LICENSES.md`、`NOTICE`。
- 本项目基于 **New API**（QuantumNous/new-api）开发，相关品牌与署名信息受保护、保留所有上游归属。
