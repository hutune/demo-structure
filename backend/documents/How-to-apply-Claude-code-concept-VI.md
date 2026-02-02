# How to Apply Claude Code (Concept)

# Hướng dẫn cộng tác đội Claude Code – Dự án RMN
## Cấu trúc tách Frontend / Backend+Infra

---

## 1. Tổng quan cấu trúc repository

```
GitHub Organization: rmn-platform/
│
├── 📦 rmn-frontend ← Đội frontend
│   └── Ứng dụng Flutter Web
│
├── 📦 rmn-backend ← Đội backend + hạ tầng
│   ├── Vi mô dịch vụ Golang
│   ├── Manifest Kubernetes
│   └── Đặc tả API (OpenAPI) ⭐ Nguồn chân lý
│
└── (tùy chọn) 📦 rmn-api-spec ← Chỉ đặc tả API (khi quản lý riêng)
```

### Chiến lược chia sẻ đặc tả API

**Khuyến nghị: Repo backend sở hữu đặc tả API → Frontend tiêu thụ**

```
[rmn-backend]                    [rmn-frontend]
api-spec/openapi.yaml ──────▶ scripts/sync-api.sh
     (git submodule hoặc tải về)
     ▼
Phát hành qua GitHub Release / gói npm
```

---

## 2. Cấu trúc repo Frontend (rmn-frontend)

```
rmn-frontend/
├── CLAUDE.md              # Ngữ cảnh frontend
├── .claude/
│   ├── settings.json      # Cài đặt chia sẻ đội
│   ├── settings.local.json # Cá nhân (gitignore)
│   └── commands/
│       ├── feature-create.md
│       ├── widget-create.md
│       ├── api-sync.md     # Đồng bộ đặc tả API
│       └── pr-create.md
│
├── lib/
│   ├── api/
│   │   ├── generated/      # Từ OpenAPI (không sửa)
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
├── api-spec/               # Đồng bộ từ backend
│   └── openapi.yaml        # (submodule hoặc tải về)
│
├── scripts/
│   ├── sync-api.sh        # Script đồng bộ đặc tả API
│   └── gen-api-client.sh  # Sinh client Dart
│
├── pubspec.yaml
└── Makefile
```

### 2.1 CLAUDE.md Frontend

```markdown
# RMN Frontend (Flutter)

## Tổng quan dự án
Web frontend nền tảng quản lý quảng cáo RMN

## Công nghệ
- Flutter 3.x / Dart
- State: flutter_bloc
- API client: dio + openapi-generator
- Điều hướng: go_router
- Test: flutter_test, mockito

## Tích hợp backend
- Đặc tả API: api-spec/openapi.yaml (đồng bộ từ repo backend)
- API client: lib/api/generated/ (tự sinh, không sửa)
- Endpoint theo môi trường:
  - dev: https://api-dev.rmn-platform.com
  - stg: https://api-stg.rmn-platform.com
  - prd: https://api.rmn-platform.com

## Cấu trúc thư mục
- lib/features/[tính-năng]/ → module tính năng (BLoC, pages, widgets)
- lib/shared/ → widget dùng chung, utils
- lib/api/generated/ → code sinh (không sửa!)

## Lệnh
```bash
# Chạy dev
flutter run -d chrome

# Đồng bộ đặc tả API (từ repo backend)
make sync-api

# Sinh lại API client
make gen-api

# Test
flutter test

# Build
flutter build web --dart-define=ENV=dev
```

## Quy ước code
- Bắt buộc BLoC: Event → BLoC → State
- Phân tầng widget: page > section > component
- Gọi API: dùng Repository pattern
- Xử lý lỗi: pattern Either<Failure, Success>

## Quy tắc nhánh
- feature/[tên-tính-năng] (vd: feature/campaign-list)
- Tiêu đề PR: rõ ràng, tiếng Anh hoặc tiếng địa phương

## Lưu ý
- Không sửa trực tiếp lib/api/generated/
- Cần thay đổi API thì yêu cầu đội backend → cập nhật đặc tả → sync
```

### 2.2 settings.json Frontend

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

### 2.3 Lệnh slash Frontend

**.claude/commands/feature-create.md**
```markdown
# Tạo module tính năng

Tạo module tính năng mới theo cấu trúc Flutter chuẩn.

## Cấu trúc cần tạo
```
lib/features/[tên-tính-năng]/
├── bloc/
│   ├── [tính-năng]_bloc.dart
│   ├── [tính-năng]_event.dart
│   └── [tính-năng]_state.dart
├── pages/
│   └── [tính-năng]_page.dart
├── widgets/
│   └── .gitkeep
└── repository/
    └── [tính-năng]_repository.dart
```

## Các bước
1. Tạo file theo cấu trúc trên
2. Viết boilerplate BLoC
3. Thêm route vào go_router
4. Tạo file test cơ bản trong test/features/[tính-năng]/

## Tên tính năng
$ARGUMENTS
```

**.claude/commands/api-sync.md**
```markdown
# Đồng bộ đặc tả API

Lấy đặc tả API mới nhất từ repo backend và sinh lại client.

## Các bước
1. Chạy `make sync-api` (tải openapi.yaml từ repo backend)
2. Kiểm tra và tóm tắt thay đổi
3. Chạy `make gen-api` (sinh lại Dart client)
4. Mô tả thay đổi chính trong client được sinh
5. Liệt kê Repository/BLoC bị ảnh hưởng

## Chạy
$ARGUMENTS
```

---

## 3. Cấu trúc repo Backend+Hạ tầng (rmn-backend)

```
rmn-backend/
├── CLAUDE.md              # Ngữ cảnh toàn bộ backend
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
├── api-spec/               # ⭐ Đặc tả API (nguồn chân lý)
│   ├── CLAUDE.md
│   └── openapi.yaml
│
├── proto/                  # Schema sự kiện Kafka
│   ├── CLAUDE.md
│   ├── buf.yaml
│   └── events/
│       ├── campaign/v1/
│       ├── user/v1/
│       └── billing/v1/
│
├── services/               # Vi mô dịch vụ
│   ├── CLAUDE.md          # Ngữ cảnh chung backend
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
├── pkg/                    # Thư viện dùng chung
│   ├── kafka/
│   ├── middleware/
│   └── config/
│
├── infrastructure/        # Code hạ tầng
│   ├── CLAUDE.md
│   ├── terraform/         # (tùy chọn) Tài nguyên cloud
│   ├── helm/              # Helm chart dùng chung
│   │   ├── base/
│   │   └── charts/
│   └── argocd/
│       └── appset.yaml
│
├── .github/
│   └── workflows/
│       ├── api-gateway.yaml
│       ├── user-service.yaml
│       └── publish-api-spec.yaml  # Phát hành đặc tả API
│
├── go.work
├── Makefile
└── docker-compose.yaml    # Môi trường dev local
```

### 3.1 CLAUDE.md gốc Backend

```markdown
# RMN Backend + Hạ tầng

## Tổng quan dự án
Backend và hạ tầng nền tảng quản lý quảng cáo RMN

## Kiến trúc
- Vi mô dịch vụ: Golang
- Hàng đợi tin: Kafka (hướng sự kiện)
- Cơ sở dữ liệu: PostgreSQL
- Cache: Redis
- Hạ tầng: Kubernetes (ArgoCD GitOps)

## Môi trường
- dev: phát triển
- stg: staging
- prd: production

## Cấu trúc thư mục
- api-spec/ → Đặc tả OpenAPI (cho frontend)
- proto/ → Schema sự kiện Kafka (Protobuf)
- services/ → các vi mô dịch vụ
- pkg/ → thư viện Go dùng chung
- infrastructure/ → K8s, Terraform, ArgoCD

## Danh sách dịch vụ
| Dịch vụ          | Port | Vai trò               |
|------------------|------|------------------------|
| api-gateway      | 8080 | Điểm vào API bên ngoài |
| user-service     | 8081 | Người dùng / xác thực  |
| campaign-service | 8082 | Quản lý campaign       |
| billing-service  | 8083 | Thanh toán / quyết toán |
| device-service   | 8084 | Thiết bị signage       |

## Lệnh
```bash
# Chạy stack local (Kafka, PostgreSQL, Redis)
docker-compose up -d

# Chạy một dịch vụ
cd services/[tên-dịch-vụ] && go run ./cmd/...

# Test toàn bộ
make test

# Sinh Proto
buf generate

# Sinh từ đặc tả API
make gen-api

# Lint
make lint
```

## Quy tắc nhánh
- feature/[dịch-vụ]-[tính-năng] (vd: feature/campaign-create-api)
- infra/[thay-đổi] (vd: infra/add-redis-cluster)

## Phát hành đặc tả API
Khi api-spec/openapi.yaml thay đổi → tự tạo GitHub Release.
Đội frontend có thể đồng bộ phiên bản mới nhất.
```

### 3.2 CLAUDE.md chung cho Services (services/CLAUDE.md)

```markdown
# Backend Services

## Công nghệ
- Go 1.23+
- HTTP: chi router
- gRPC: grpc-go
- Kafka: segmentio/kafka-go
- DB: pgx v5
- Cấu hình: viper
- Log: zerolog

## Cấu trúc dịch vụ chuẩn
```
services/[tên-dịch-vụ]/
├── cmd/[tên-dịch-vụ]/
│   └── main.go
├── internal/
│   ├── handler/   # Handler HTTP/gRPC
│   ├── service/   # Logic nghiệp vụ
│   ├── repository/# Truy cập dữ liệu
│   ├── event/     # Kafka publish/subscribe
│   └── domain/    # Mô hình domain
├── go.mod
├── Dockerfile
└── chart/         # Helm chart
    ├── values.yaml
    ├── values-dev.yaml
    ├── values-stg.yaml
    └── values-prd.yaml
```

## Quy ước code
- Tuân thủ Go project layout chuẩn
- Lỗi: fmt.Errorf("context: %w", err)
- Log: zerolog, log có cấu trúc
- Test: table-driven

## Quy ước Kafka
- Đặt tên topic: [domain].[event] (vd: campaign.created)
- Schema sự kiện: proto/events/[domain]/v1/
- Consumer group: [tên-dịch-vụ]-[env]

## Giao tiếp giữa dịch vụ
- Đồng bộ: gRPC (trực tiếp)
- Bất đồng bộ: Kafka (publish/subscribe)
- Bên ngoài: REST qua api-gateway
```

### 3.3 CLAUDE.md Hạ tầng (infrastructure/CLAUDE.md)

```markdown
# Infrastructure

## Cấu trúc Kubernetes
- ArgoCD triển khai GitOps
- Helm chart triển khai dịch vụ
- Tách file values theo môi trường

## Môi trường
| Môi trường | Cluster | Namespace |
|------------|---------|-----------|
| dev | rmn-dev | rmn-dev   |
| stg | rmn-stg | rmn-stg   |
| prd | rmn-prd | rmn-prd   |

## ArgoCD ApplicationSet
Tự phát hiện thư mục services/*/chart/

## Lệnh
```bash
# Thiết lập K8s local (kind)
make setup-local-k8s

# Lint Helm chart
helm lint services/*/chart/

# Xem trước manifest
helm template services/api-gateway/chart -f services/api-gateway/chart/values-dev.yaml
```

## Lưu ý
- Không dùng kubectl trực tiếp lên prd → chỉ qua ArgoCD
- Secret dùng External Secrets Operator
- Không đưa thông tin nhạy cảm vào values-prd.yaml
```

### 3.4 settings.json Backend

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

### 3.5 Lệnh slash Backend

**.claude/commands/api-add.md**
```markdown
# Thêm endpoint API

Thêm endpoint mới vào đặc tả OpenAPI và sinh code liên quan.

## Các bước
1. Đọc api-spec/openapi.yaml
2. Thêm endpoint theo yêu cầu (paths, schemas)
3. Chạy `make gen-api` (sinh code Go)
4. Thêm handler stub trong api-gateway
5. Tóm tắt thay đổi và soạn tin nhắn thông báo đội frontend

## Mô tả API
$ARGUMENTS
```

**.claude/commands/service-create.md**
```markdown
# Tạo vi mô dịch vụ

Tạo vi mô dịch vụ Golang mới theo cấu trúc chuẩn.

## Cấu trúc cần tạo
```
services/[tên-dịch-vụ]/
├── cmd/[tên-dịch-vụ]/main.go
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

## Các bước
1. Tạo cấu trúc như trên
2. Thêm dịch vụ vào go.work
3. Thêm target dịch vụ vào Makefile
4. Tạo workflow GitHub Actions
5. Viết CLAUDE.md cho dịch vụ

## Tên dịch vụ và mô tả
$ARGUMENTS
```

**.claude/commands/event-create.md**
```markdown
# Tạo sự kiện Kafka

Định nghĩa sự kiện Kafka mới bằng schema Protobuf.

## Các bước
1. Tạo/sửa file .proto trong proto/events/[domain]/v1/
2. Định nghĩa message sự kiện
3. Chạy `buf lint`
4. Chạy `buf generate`
5. Thêm code publish trong dịch vụ phát hành
6. (Tùy chọn) Thêm handler trong dịch vụ subscribe

## Mô tả sự kiện
$ARGUMENTS
```

---

## 4. Quy trình cộng tác giữa đội

### 4.1 Quy trình thay đổi API

```
┌─────────────────────────────────────────────────────────────┐
│ Quy trình thay đổi API                                       │
└─────────────────────────────────────────────────────────────┘

[1] Yêu cầu từ frontend
    └─▶ Slack/Issue: "Cần thêm tham số lọc cho API danh sách campaign"

[2] Công việc backend (rmn-backend)
    ├─▶ /api-add: thêm query status, date_from cho GET /campaigns
    ├─▶ Triển khai handler api-gateway
    ├─▶ Triển khai logic campaign-service
    ├─▶ Tạo PR và review
    └─▶ Merge vào main → Tự phát hành đặc tả API

[3] Công việc frontend (rmn-frontend)
    ├─▶ /api-sync (đồng bộ đặc tả mới nhất)
    ├─▶ Dùng tham số mới trong Repository
    ├─▶ Cập nhật BLoC/UI
    └─▶ Tạo PR và merge
```

### 4.2 Cách đồng bộ đặc tả API

**Tùy chọn A: Tải từ GitHub Release (khuyến nghị)**

```bash
# rmn-frontend/scripts/sync-api.sh
#!/bin/bash
LATEST=$(gh release view --repo rmn-platform/rmn-backend --json tagName -q .tagName)
curl -L "https://github.com/rmn-platform/rmn-backend/releases/download/${LATEST}/openapi.yaml" \
  -o api-spec/openapi.yaml
echo "Synced to version: ${LATEST}"
```

**Tùy chọn B: Git Submodule**

```bash
# Trong rmn-frontend
git submodule add https://github.com/rmn-platform/rmn-backend.git external/backend
# Symlink api-spec → external/backend/api-spec
```

### 4.3 Chia sẻ ngữ cảnh Claude Code giữa nhiều repo

**Cách: dùng tùy chọn --add-dir**

Khi dev frontend cần tham chiếu đặc tả API backend:

```bash
# Đã clone cả hai repo
cd ~/projects/rmn-frontend
claude --add-dir ~/projects/rmn-backend/api-spec

# Giờ có thể làm frontend với đặc tả API backend trong ngữ cảnh
> Dựa vào api-spec/openapi.yaml, triển khai Repository campaign giúp tôi
```

---

## 5. Onboarding theo đội

### 5.1 Onboarding đội Frontend

```bash
# 1. Clone repo
git clone https://github.com/rmn-platform/rmn-frontend.git
cd rmn-frontend

# 2. Cài dependency
flutter pub get

# 3. Đồng bộ đặc tả API
make sync-api
make gen-api

# 4. Chạy dev server
flutter run -d chrome

# 5. Khởi động Claude Code
claude
> /help   # Xem lệnh có sẵn
```

### 5.2 Onboarding đội Backend

```bash
# 1. Clone repo
git clone https://github.com/rmn-platform/rmn-backend.git
cd rmn-backend

# 2. Chạy hạ tầng local
docker-compose up -d   # Kafka, PostgreSQL, Redis

# 3. Dependency và sinh code
go work sync
buf generate
make gen-api

# 4. Chạy một dịch vụ
cd services/api-gateway && go run ./cmd/api-gateway

# 5. Khởi động Claude Code
claude
> /help
```

---

## 6. Tóm tắt cấu hình theo môi trường

### Frontend (rmn-frontend)

| File | Git | Mục đích |
|------|-----|----------|
| CLAUDE.md | ✅ | Ngữ cảnh dự án |
| .claude/settings.json | ✅ | Cài đặt chia sẻ đội |
| .claude/settings.local.json | ❌ | Cá nhân |
| .claude/commands/*.md | ✅ | Lệnh chia sẻ đội |
| lib/api/generated/ | ✅ | Code sinh (không sửa) |

### Backend (rmn-backend)

| File | Git | Mục đích |
|------|-----|----------|
| CLAUDE.md | ✅ | Ngữ cảnh toàn bộ |
| services/CLAUDE.md | ✅ | Chung backend |
| services/*/CLAUDE.md | ✅ | Chi tiết từng dịch vụ |
| infrastructure/CLAUDE.md | ✅ | Ngữ cảnh hạ tầng |
| .claude/settings.json | ✅ | Cài đặt chia sẻ đội |
| .claude/settings.local.json | ❌ | Cá nhân |

---

## 7. Kênh giao tiếp

```
┌─────────────────────────────────────────────────────────┐
│ Yêu cầu thay đổi API                                    │
│ ───────────────────                                    │
│ Slack: #rmn-api-changes                                 │
│ GitHub Issue: repo rmn-backend, nhãn "api-request"      │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│ Thông báo hoàn thành thay đổi API                       │
│ ───────────────────────────────                        │
│ GitHub Actions → Slack:                                 │
│ "🚀 API v1.2.3 đã phát hành - frontend cần sync"        │
└─────────────────────────────────────────────────────────┘
```

---

## 8. So sánh tóm tắt

| Hạng mục | Frontend (rmn-frontend) | Backend (rmn-backend) |
|----------|-------------------------|------------------------|
| Công nghệ | Flutter/Dart | Golang + K8s |
| Đặc tả API | Đồng bộ rồi dùng | Nguồn chân lý |
| Ngữ cảnh Claude | Tối ưu cho FE | BE + hạ tầng tích hợp |
| Triển khai | CDN/S3 | ArgoCD GitOps |
| Lệnh chính | /feature-create, /api-sync | /service-create, /api-add |

Với cấu trúc này, mỗi đội có thể **làm việc độc lập** đồng thời **cộng tác rõ ràng qua đặc tả API**.
