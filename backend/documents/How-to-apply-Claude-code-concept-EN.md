# How to Apply Claude Code (Concept)

# RMN Project Claude Code Team Collaboration Guide
## Frontend / Backend+Infra Separated Structure

---

## 1. Repository Structure Overview

```
GitHub Organization: rmn-platform/
│
├── 📦 rmn-frontend ← Frontend team
│   └── Flutter Web app
│
├── 📦 rmn-backend ← Backend + Infra team
│   ├── Golang microservices
│   ├── Kubernetes manifests
│   └── API spec (OpenAPI) ⭐ Source of truth
│
└── (optional) 📦 rmn-api-spec ← API spec only (when managed separately)
```

### API Spec Sharing Strategy

**Recommended: Backend repo owns API spec → Frontend consumes**

```
[rmn-backend]                    [rmn-frontend]
api-spec/openapi.yaml ──────▶ scripts/sync-api.sh
     (git submodule or download)
     ▼
Deploy via GitHub Release / npm package
```

---

## 2. Frontend Repo Structure (rmn-frontend)

```
rmn-frontend/
├── CLAUDE.md              # Frontend context
├── .claude/
│   ├── settings.json      # Team-shared settings
│   ├── settings.local.json # Personal (gitignore)
│   └── commands/
│       ├── feature-create.md
│       ├── widget-create.md
│       ├── api-sync.md     # API spec sync
│       └── pr-create.md
│
├── lib/
│   ├── api/
│   │   ├── generated/      # From OpenAPI (do not edit)
│   │   └── client.dart
│   ├── features/
│   │   ├── campaign/
│   │   ├── dashboard/
│   │   └── auth/
│   ├── shared/
│   │   ├── widgets/
│   │   └── utils/
│   └── main.dart
│
├── api-spec/               # Synced from backend
│   └── openapi.yaml       # (submodule or download)
│
├── scripts/
│   ├── sync-api.sh        # API spec sync script
│   └── gen-api-client.sh  # Dart client generation
│
├── pubspec.yaml
└── Makefile
```

### 2.1 Frontend CLAUDE.md

```markdown
# RMN Frontend (Flutter)

## Project overview
RMN ad management platform web frontend

## Tech stack
- Flutter 3.x / Dart
- State: flutter_bloc
- API client: dio + openapi-generator
- Routing: go_router
- Testing: flutter_test, mockito

## Backend integration
- API spec: api-spec/openapi.yaml (synced from backend repo)
- API client: lib/api/generated/ (auto-generated, do not edit)
- Per-environment endpoints:
  - dev: https://api-dev.rmn-platform.com
  - stg: https://api-stg.rmn-platform.com
  - prd: https://api.rmn-platform.com

## Directory structure
- lib/features/[feature]/ → feature modules (BLoC, pages, widgets)
- lib/shared/ → shared widgets, utils
- lib/api/generated/ → generated code (do not edit!)

## Commands
```bash
# Dev server
flutter run -d chrome

# Sync API spec (from backend repo)
make sync-api

# Regenerate API client
make gen-api

# Tests
flutter test

# Build
flutter build web --dart-define=ENV=dev
```

## Coding conventions
- BLoC pattern required: Event → BLoC → State
- Widget layering: page > section > component
- API calls: Repository pattern
- Error handling: Either<Failure, Success> pattern

## Branch rules
- feature/[feature-name] (e.g. feature/campaign-list)
- PR title: clear, in English or local language

## Important
- Do not edit lib/api/generated/ directly
- For API changes, request from backend team → update spec → sync
```

### 2.2 Frontend settings.json

```json
{
  "model": "claude-sonnet-4-20250514",
  "permissions": {
    "allow": [
      "Read",
      "Write(lib/**)",
      "Write(test/**)",
      "Bash(flutter:*)",
      "Bash(dart:*)",
      "Bash(make:*)",
      "Bash(git add:*)",
      "Bash(git commit:*)",
      "Bash(git push:*)",
      "Bash(git checkout:*)",
      "Bash(git branch:*)",
      "Bash(gh pr:*)"
    ],
    "deny": [
      "Read(.env*)",
      "Write(.env*)",
      "Write(lib/api/generated/**)",
      "Bash(rm -rf:*)",
      "Bash(sudo:*)"
    ]
  },
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write(lib/**/*.dart)",
        "hooks": [{
          "type": "command",
          "command": "dart format $CLAUDE_FILE_PATH"
        }]
      }
    ]
  }
}
```

### 2.3 Frontend Slash Commands

**.claude/commands/feature-create.md**
```markdown
# Create feature module

Creates a new feature module in standard Flutter structure.

## Structure to create
```
lib/features/[feature-name]/
├── bloc/
│   ├── [feature]_bloc.dart
│   ├── [feature]_event.dart
│   └── [feature]_state.dart
├── pages/
│   └── [feature]_page.dart
├── widgets/
│   └── .gitkeep
└── repository/
    └── [feature]_repository.dart
```

## Steps
1. Create files as above
2. Add BLoC boilerplate
3. Add route in go_router
4. Add basic tests under test/features/[feature]/

## Feature name
$ARGUMENTS
```

**.claude/commands/api-sync.md**
```markdown
# API spec sync

Fetches latest API spec from backend repo and regenerates client.

## Steps
1. Run `make sync-api` (download openapi.yaml from backend repo)
2. Summarize changes
3. Run `make gen-api` (regenerate Dart client)
4. Summarize main changes in generated client
5. List affected Repository/BLoC

## Run
$ARGUMENTS
```

---

## 3. Backend+Infra Repo Structure (rmn-backend)

```
rmn-backend/
├── CLAUDE.md              # Full backend context
├── .claude/
│   ├── settings.json
│   ├── settings.local.json
│   └── commands/
│       ├── service-create.md
│       ├── event-create.md
│       ├── api-add.md
│       ├── deploy.md
│       └── pr-create.md
│
├── api-spec/              # ⭐ API spec (source of truth)
│   ├── CLAUDE.md
│   └── openapi.yaml
│
├── proto/                 # Kafka event schemas
│   ├── CLAUDE.md
│   ├── buf.yaml
│   └── events/
│       ├── campaign/v1/
│       ├── user/v1/
│       └── billing/v1/
│
├── services/              # Microservices
│   ├── CLAUDE.md         # Backend common context
│   ├── api-gateway/
│   │   ├── CLAUDE.md
│   │   ├── cmd/
│   │   ├── internal/
│   │   ├── go.mod
│   │   └── Dockerfile
│   ├── user-service/
│   ├── campaign-service/
│   ├── billing-service/
│   └── device-service/
│
├── pkg/                   # Shared libraries
│   ├── kafka/
│   ├── middleware/
│   └── config/
│
├── infrastructure/        # Infra code
│   ├── CLAUDE.md
│   ├── terraform/        # (optional) Cloud resources
│   ├── helm/             # Shared Helm charts
│   │   ├── base/
│   │   └── charts/
│   └── argocd/
│       └── appset.yaml
│
├── .github/
│   └── workflows/
│       ├── api-gateway.yaml
│       ├── user-service.yaml
│       └── publish-api-spec.yaml  # API spec release
│
├── go.work
├── Makefile
└── docker-compose.yaml   # Local dev
```

### 3.1 Backend Root CLAUDE.md

```markdown
# RMN Backend + Infrastructure

## Project overview
RMN ad management platform backend and infrastructure

## Architecture
- Microservices: Golang
- Message queue: Kafka (event-driven)
- Database: PostgreSQL
- Cache: Redis
- Infra: Kubernetes (ArgoCD GitOps)

## Environments
- dev: development
- stg: staging
- prd: production

## Directory structure
- api-spec/ → OpenAPI spec (for frontend)
- proto/ → Kafka event schemas (Protobuf)
- services/ → microservices
- pkg/ → shared Go libraries
- infrastructure/ → K8s, Terraform, ArgoCD

## Services
| Service         | Port | Role              |
|-----------------|------|--------------------|
| api-gateway     | 8080 | External API entry |
| user-service    | 8081 | Users / auth       |
| campaign-service| 8082 | Campaigns          |
| billing-service | 8083 | Billing / settlement |
| device-service  | 8084 | Signage devices    |

## Commands
```bash
# Local stack (Kafka, PostgreSQL, Redis)
docker-compose up -d

# Run a service
cd services/[service-name] && go run ./cmd/...

# All tests
make test

# Proto generation
buf generate

# Generate from API spec
make gen-api

# Lint
make lint
```

## Branch rules
- feature/[service]-[feature] (e.g. feature/campaign-create-api)
- infra/[change] (e.g. infra/add-redis-cluster)

## API spec release
On api-spec/openapi.yaml change, GitHub Release is created automatically.
Frontend can sync to latest version.
```

### 3.2 Services Common CLAUDE.md (services/CLAUDE.md)

```markdown
# Backend Services

## Tech stack
- Go 1.23+
- HTTP: chi router
- gRPC: grpc-go
- Kafka: segmentio/kafka-go
- DB: pgx v5
- Config: viper
- Logging: zerolog

## Standard service layout
```
services/[service-name]/
├── cmd/[service-name]/
│   └── main.go
├── internal/
│   ├── handler/   # HTTP/gRPC handlers
│   ├── service/   # Business logic
│   ├── repository/# Data access
│   ├── event/     # Kafka publish/subscribe
│   └── domain/    # Domain models
├── go.mod
├── Dockerfile
└── chart/         # Helm chart
    ├── values.yaml
    ├── values-dev.yaml
    ├── values-stg.yaml
    └── values-prd.yaml
```

## Coding conventions
- Follow standard Go project layout
- Errors: fmt.Errorf("context: %w", err)
- Logging: zerolog, structured logs
- Tests: table-driven tests

## Kafka conventions
- Topic naming: [domain].[event] (e.g. campaign.created)
- Event schema: proto/events/[domain]/v1/
- Consumer group: [service-name]-[env]

## Inter-service communication
- Sync: gRPC (direct)
- Async: Kafka (publish/subscribe)
- External: REST via api-gateway
```

### 3.3 Infrastructure CLAUDE.md (infrastructure/CLAUDE.md)

```markdown
# Infrastructure

## Kubernetes layout
- ArgoCD for GitOps deployment
- Helm charts for services
- Per-environment values files

## Environments
| Env | Cluster | Namespace |
|-----|---------|-----------|
| dev | rmn-dev | rmn-dev   |
| stg | rmn-stg | rmn-stg   |
| prd | rmn-prd | rmn-prd   |

## ArgoCD ApplicationSet
Auto-discovers services/*/chart/

## Commands
```bash
# Local K8s (kind)
make setup-local-k8s

# Helm lint
helm lint services/*/chart/

# Manifest preview
helm template services/api-gateway/chart -f services/api-gateway/chart/values-dev.yaml
```

## Important
- No direct kubectl to prd → ArgoCD only
- Secrets via External Secrets Operator
- Never put secrets in values-prd.yaml
```

### 3.4 Backend settings.json

```json
{
  "model": "claude-sonnet-4-20250514",
  "permissions": {
    "allow": [
      "Read",
      "Write(services/**)",
      "Write(pkg/**)",
      "Write(proto/**)",
      "Write(api-spec/**)",
      "Write(infrastructure/**)",
      "Bash(go:*)",
      "Bash(buf:*)",
      "Bash(make:*)",
      "Bash(docker:*)",
      "Bash(docker-compose:*)",
      "Bash(helm lint:*)",
      "Bash(helm template:*)",
      "Bash(git:*)",
      "Bash(gh pr:*)"
    ],
    "deny": [
      "Read(.env*)",
      "Read(**/secrets/**)",
      "Read(**/*secret*)",
      "Write(.env*)",
      "Write(**/secrets/**)",
      "Bash(rm -rf:*)",
      "Bash(sudo:*)",
      "Bash(kubectl delete:*)",
      "Bash(kubectl apply:*)",
      "Bash(helm install:*)",
      "Bash(helm upgrade:*)"
    ]
  },
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write(services/**/*.go)",
        "hooks": [{
          "type": "command",
          "command": "gofmt -w $CLAUDE_FILE_PATH"
        }]
      }
    ]
  }
}
```

### 3.5 Backend Slash Commands

**.claude/commands/api-add.md**
```markdown
# Add API endpoint

Adds a new endpoint to OpenAPI spec and generates related code.

## Steps
1. Read api-spec/openapi.yaml
2. Add requested endpoint (paths, schemas)
3. Run `make gen-api` (generate Go code)
4. Add handler stub in api-gateway
5. Summarize changes and draft message for frontend team

## API description
$ARGUMENTS
```

**.claude/commands/service-create.md**
```markdown
# Create microservice

Creates a new Golang microservice in standard layout.

## Structure to create
```
services/[service-name]/
├── cmd/[service-name]/main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── event/
│   └── domain/
├── go.mod
├── Dockerfile
├── CLAUDE.md
└── chart/
    ├── Chart.yaml
    ├── values.yaml
    ├── values-dev.yaml
    ├── values-stg.yaml
    ├── values-prd.yaml
    └── templates/
```

## Steps
1. Create structure as above
2. Add service to go.work
3. Add service target to Makefile
4. Add GitHub Actions workflow
5. Write service CLAUDE.md

## Service name and description
$ARGUMENTS
```

**.claude/commands/event-create.md**
```markdown
# Create Kafka event

Defines a new Kafka event with Protobuf schema.

## Steps
1. Create/update .proto under proto/events/[domain]/v1/
2. Define event message
3. Run `buf lint`
4. Run `buf generate`
5. Add publish code in publishing service
6. (Optional) Add handler in subscribing service

## Event description
$ARGUMENTS
```

---

## 4. Cross-Team Collaboration Workflow

### 4.1 API Change Process

```
┌─────────────────────────────────────────────────────────────┐
│ API change workflow                                          │
└─────────────────────────────────────────────────────────────┘

[1] Frontend request
    └─▶ Slack/Issue: "Add filter params to campaign list API"

[2] Backend work (rmn-backend)
    ├─▶ /api-add: add status, date_from query params to GET /campaigns
    ├─▶ Implement api-gateway handler
    ├─▶ Implement campaign-service logic
    ├─▶ Open PR and review
    └─▶ Merge to main → API spec auto-released

[3] Frontend work (rmn-frontend)
    ├─▶ /api-sync (sync latest spec)
    ├─▶ Use new params in Repository
    ├─▶ Update BLoC/UI
    └─▶ Open PR and merge
```

### 4.2 API Spec Sync Options

**Option A: GitHub Release download (recommended)**

```bash
# rmn-frontend/scripts/sync-api.sh
#!/bin/bash
LATEST=$(gh release view --repo rmn-platform/rmn-backend --json tagName -q .tagName)
curl -L "https://github.com/rmn-platform/rmn-backend/releases/download/${LATEST}/openapi.yaml" \
  -o api-spec/openapi.yaml
echo "Synced to version: ${LATEST}"
```

**Option B: Git Submodule**

```bash
# From rmn-frontend
git submodule add https://github.com/rmn-platform/rmn-backend.git external/backend
# Symlink api-spec → external/backend/api-spec
```

### 4.3 Sharing Claude Code Context Across Repos

**Method: use --add-dir**

When a frontend dev needs to reference backend API spec:

```bash
# With both repos cloned
cd ~/projects/rmn-frontend
claude --add-dir ~/projects/rmn-backend/api-spec

# Now you can work on frontend with backend API spec in context
> Implement campaign Repository based on api-spec/openapi.yaml
```

---

## 5. Team Onboarding

### 5.1 Frontend Team Onboarding

```bash
# 1. Clone repo
git clone https://github.com/rmn-platform/rmn-frontend.git
cd rmn-frontend

# 2. Install dependencies
flutter pub get

# 3. Sync API spec
make sync-api
make gen-api

# 4. Run dev server
flutter run -d chrome

# 5. Start Claude Code
claude
> /help   # List available commands
```

### 5.2 Backend Team Onboarding

```bash
# 1. Clone repo
git clone https://github.com/rmn-platform/rmn-backend.git
cd rmn-backend

# 2. Start local infra
docker-compose up -d   # Kafka, PostgreSQL, Redis

# 3. Dependencies and codegen
go work sync
buf generate
make gen-api

# 4. Run a service
cd services/api-gateway && go run ./cmd/api-gateway

# 5. Start Claude Code
claude
> /help
```

---

## 6. Per-Environment Config Summary

### Frontend (rmn-frontend)

| File | Git | Purpose |
|------|-----|---------|
| CLAUDE.md | ✅ | Project context |
| .claude/settings.json | ✅ | Team settings |
| .claude/settings.local.json | ❌ | Personal |
| .claude/commands/*.md | ✅ | Team commands |
| lib/api/generated/ | ✅ | Generated (do not edit) |

### Backend (rmn-backend)

| File | Git | Purpose |
|------|-----|---------|
| CLAUDE.md | ✅ | Full context |
| services/CLAUDE.md | ✅ | Backend common |
| services/*/CLAUDE.md | ✅ | Per-service |
| infrastructure/CLAUDE.md | ✅ | Infra context |
| .claude/settings.json | ✅ | Team settings |
| .claude/settings.local.json | ❌ | Personal |

---

## 7. Communication Channels

```
┌─────────────────────────────────────────────────────────┐
│ API change requests                                     │
│ ─────────────────                                      │
│ Slack: #rmn-api-changes                                 │
│ GitHub Issue: rmn-backend repo, label "api-request"      │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│ API change completion                                   │
│ ───────────────────                                    │
│ GitHub Actions → Slack:                                 │
│ "🚀 API v1.2.3 released - frontend sync needed"         │
└─────────────────────────────────────────────────────────┘
```

---

## 8. Summary Comparison

| Item | Frontend (rmn-frontend) | Backend (rmn-backend) |
|------|-------------------------|------------------------|
| Tech | Flutter/Dart | Golang + K8s |
| API spec | Synced consumer | Source of truth |
| Claude context | FE-optimized | BE + infra integrated |
| Deploy | CDN/S3 | ArgoCD GitOps |
| Main commands | /feature-create, /api-sync | /service-create, /api-add |

With this structure, each team can **work independently** while **collaborating clearly via the API spec**.
