# Claude Code – Hướng dẫn làm việc theo đội

# Hướng dẫn cộng tác đội với Claude Code: Nắm vững Agent, Subagent và Skill

Claude Code hỗ trợ cộng tác đội và tự động hóa qua **cấu trúc 3 tầng**—Main Agent, Subagent và Skill. Để đội 3 người (backend/frontend/PM) tự động hóa toàn bộ phát triển nền tảng RMN SaaS, việc dùng chiến lược hệ thống này là then chốt. Với gói đăng ký đội, chia sẻ cấu trúc thư mục `.claude/` và `CLAUDE.md` qua Git giúp thiết lập cùng một môi trường phát triển ngay lập tức.

---

## Khác biệt cốt lõi: Agent, Subagent, Skill

Hệ thống agent của Claude Code phục vụ hai mục tiêu: **tách nhiệm vụ** và **quản lý ngữ cảnh**. Mỗi thành phần có vai trò rõ ràng.

| Thành phần | Mục đích | Ngữ cảnh | Cách gọi | Vị trí |
|------------|----------|----------|----------|--------|
| **Main Agent** | Xử lý hội thoại chính | Cửa sổ chat dùng chung | Luôn bật | Tích hợp sẵn |
| **Subagent** | Công việc chuyên môn được ủy thác | **Ngữ cảnh tách biệt** | Ủy thác tự động hoặc rõ ràng | `.claude/agents/` |
| **Skill** | Bơm kiến thức/năng lực | Nạp vào ngữ cảnh hiện tại | Tự kích hoạt theo mức độ liên quan | `.claude/skills/` |

**Subagent** chạy trong cửa sổ ngữ cảnh riêng nên không làm rối hội thoại chính. Phù hợp cho công việc độc lập như code review, khám phá codebase, nghiên cứu sâu. **Skill** bơm kiến thức vào ngữ cảnh hiện tại và chỉ nạp phần cần thiết qua cơ chế progressive disclosure.

### Ba Subagent tích hợp sẵn

Claude Code cung cấp ba subagent có sẵn:

- **Explore Subagent**: Dùng Haiku, chỉ đọc; tập trung duyệt file và tìm kiếm code
- **Plan Subagent**: Dùng Sonnet; phân tích codebase và lập kế hoạch
- **Subagent đa năng**: Dùng Sonnet; nghiên cứu phức tạp kèm chỉnh sửa

---

## Cấu trúc cấu hình và chia sẻ đội

### Cấu trúc thư mục khuyến nghị

```
your-rmn-project/
├── CLAUDE.md              # Bộ nhớ dự án (commit vào Git)
├── .mcp.json              # Cấu hình MCP server (commit vào Git)
├── .claude/
│   ├── settings.json      # Cài đặt chia sẻ đội (commit vào Git)
│   ├── settings.local.json # Cài đặt cá nhân (.gitignore)
│   │
│   ├── agents/            # Subagent tùy chỉnh
│   │   ├── code-reviewer.md
│   │   ├── golang-expert.md
│   │   └── flutter-specialist.md
│   │
│   ├── commands/          # Lệnh slash
│   │   ├── test.md
│   │   ├── review.md
│   │   └── deploy/
│   │       └── staging.md
│   │
│   └── skills/            # Skills
│       ├── api-design/
│       │   └── SKILL.md
│       └── testing-patterns/
│           └── SKILL.md
│
└── agent_docs/            # Tài liệu progressive disclosure
    ├── building_the_project.md
    └── architecture.md
```

### Chia sẻ đội qua Git

**Commit (chia sẻ đội):**
- `CLAUDE.md` — ngữ cảnh dự án
- `.claude/settings.json` — quyền, biến môi trường
- `.claude/commands/` — lệnh slash của đội
- `.claude/agents/` — subagent dùng chung
- `.mcp.json` — cấu hình MCP server

**Thêm vào .gitignore (cá nhân):**
```gitignore
CLAUDE.local.md
.claude/settings.local.json
.claude/local/
```

Khi thành viên clone repo và chạy lệnh `claude`, **cùng môi trường được áp dụng tự động**.

---

## Mẫu CLAUDE.md cho dự án RMN SaaS

Nên giữ **CLAUDE.md trong khoảng 60–300 dòng** để LLM tuân thủ hướng dẫn ổn định.

```markdown
# Nền tảng quản lý quảng cáo RMN SaaS

## Tổng quan dự án
Nền tảng quản lý quảng cáo SaaS dựa trên Retail Media Network.
Frontend Flutter, backend Golang, PostgreSQL, vận hành trên Kubernetes.

## Thư mục chính
- `frontend/` — Ứng dụng Flutter (lib/features/, lib/shared/, lib/core/)
- `backend/` — Dịch vụ Go (cmd/, internal/, pkg/, api/)
- `infra/` — Manifest Kubernetes, Terraform
- `docs/` — Tài liệu API, PRD

## Lệnh
```bash
# Backend
cd backend && go build ./... && go test ./...
golangci-lint run

# Frontend
cd frontend && flutter analyze && flutter test
flutter build apk --release
```

## Quy ước code
### Golang
- Định nghĩa hành vi bằng interface, DI qua hàm khởi tạo
- Luôn kiểm tra lỗi, bọc bằng %w
- Repository pattern + tầng Service

### Flutter
- Feature-First Clean Architecture (FFCA)
- Riverpod quản lý state, flutter_hooks
- Ưu tiên dark mode, UI responsive dựa flex

## Lưu ý quan trọng
- Không bao giờ commit file .env
- Câu truy vấn SQL phải được tham số hóa
- Cập nhật spec OpenAPI khi thay đổi API

## Tài liệu bổ sung
Với tác vụ phức tạp, xem `agent_docs/`:
- `architecture.md` — thiết kế hệ thống
- `api_conventions.md` — quy tắc API
```

---

## Thiết lập Agent/Skill theo vai trò

### Cho lập trình viên Backend (Golang)

**.claude/agents/golang-expert.md**
```markdown
---
name: golang-expert
description: Viết code Go, review, sinh test. Dùng cho công việc liên quan Go.
tools: Read, Write, Edit, Bash, Grep, Glob
model: claude-sonnet-4-20250514
---

Bạn là chuyên gia backend Go 15 năm kinh nghiệm.

## Trách nhiệm
- Tuân thủ Clean Architecture
- Viết test theo bảng (table-driven)
- Áp dụng pattern xử lý lỗi
- Tối ưu truy vấn PostgreSQL

## Phong cách code
- Luôn dùng context.Context
- Thiết kế interface trước
- Dùng assertion testify
```

**.claude/commands/go-test.md**
```markdown
---
name: go-test
description: Sinh test Go theo bảng (table-driven)
allowed-tools: Read, Write, Bash
---

Sinh test theo bảng cho $ARGUMENTS:
1. Bao gồm happy path, trường hợp lỗi, edge case
2. Dùng testify/assert
3. Mock chỉ dependency bên ngoài
4. Chạy `go test -v` để kiểm chứng
```

### Cho lập trình viên Frontend (Flutter)

**.claude/agents/flutter-specialist.md**
```markdown
---
name: flutter-specialist
description: Phát triển widget Flutter, quản lý state, UI. Dùng cho công việc Flutter.
tools: Read, Write, Edit, Bash, Grep, Glob
model: claude-sonnet-4-20250514
---

Bạn là chuyên gia Flutter/Dart.

## Kiến trúc
- Feature-First Clean Architecture
- Riverpod + flutter_hooks
- GoRouter cho điều hướng

## Nguyên tắc widget
- Ưu tiên widget nhỏ, có thể kết hợp
- Dùng giá trị flex thay cho kích thước hardcode
- Dùng dart:developer log() thay cho print()

## Test
- Widget test cho component UI
- Unit test cho logic nghiệp vụ
```

**.claude/commands/flutter-feature.md**
```markdown
---
name: flutter-feature
description: Tạo khung (scaffold) module tính năng Flutter
---

Tạo tính năng '$ARGUMENTS' theo cấu trúc Feature-First:

```
lib/features/$ARGUMENTS/
├── data/
│   ├── repositories/
│   └── models/
├── domain/
│   └── entities/
├── presentation/
│   ├── pages/
│   ├── widgets/
│   └── providers/
```

Tạo file cơ bản và Riverpod Provider ở mỗi tầng.
```

### Cho PM: Công cụ tài liệu

**.claude/agents/prd-writer.md**
```markdown
---
name: prd-writer
description: PRD, đặc tả tính năng, release notes. Dùng cho công việc tài liệu.
tools: Read, Write, Grep, Glob
model: claude-sonnet-4-20250514
---

Bạn là Sarah, PM cấp cao.

## Khung PRD (15 phần)
1. Executive Summary
2. Problem Statement
3. Goals & Success Metrics
4. User Stories
5. Functional Requirements
6. Non-Functional Requirements
7. Technical Constraints
8. Dependencies
9. Acceptance Criteria
10. Out of Scope
11. Risks & Mitigations
12. Timeline
13. Stakeholders
14. Open Questions
15. Appendix

Làm rõ yêu cầu bằng đủ câu hỏi trước khi viết PRD.
```

**.claude/commands/release-notes.md**
```markdown
---
name: release-notes
description: Tự sinh release notes
---

Sinh release notes từ $ARGUMENTS commit gần nhất:

1. Chạy `git log --oneline -n $ARGUMENTS`
2. Nhóm theo loại (feat, fix, docs, refactor)
3. Tóm tắt thay đổi theo góc nhìn người dùng
4. Nêu rõ Breaking Changes
5. Cập nhật `docs/CHANGELOG.md`
```

---

## Ví dụ cài đặt chia sẻ đội

**.claude/settings.json**
```json
{
  "permissions": {
    "allow": [
      "Bash(go build:*)",
      "Bash(go test:*)",
      "Bash(golangci-lint:*)",
      "Bash(flutter:*)",
      "Bash(git commit:*)",
      "Bash(git push:*)",
      "Bash(kubectl get:*)"
    ],
    "deny": [
      "Bash(rm -rf:*)",
      "Read(.env*)",
      "Read(./secrets/**)",
      "Bash(kubectl delete:*)"
    ]
  },
  "env": {
    "GO_ENV": "development",
    "FLUTTER_ENV": "development"
  },
  "hooks": {
    "PostToolUse": [
      {
        "matcher": "Write(*.go)",
        "hooks": [{
          "type": "command",
          "command": "gofmt -w \"$file\""
        }]
      },
      {
        "matcher": "Write(*.dart)",
        "hooks": [{
          "type": "command",
          "command": "dart format \"$file\""
        }]
      }
    ]
  }
}
```

Cấu hình này tự format Go/Dart khi lưu và chặn lệnh nguy hiểm.

---

## Ví dụ tự động hóa code review và test

### Tích hợp GitHub Actions

**.github/workflows/claude-review.yml**
```yaml
name: Claude Code Review
on:
  pull_request:
    types: [opened, synchronize]

jobs:
  review:
    runs-on: ubuntu-latest
    steps:
      - uses: anthropics/claude-code-action@v1
        with:
          anthropic_api_key: ${{ secrets.ANTHROPIC_API_KEY }}
          claude_args: |
            --max-turns 5
          settings: |
            {
              "permissions": {
                "allow": ["Bash", "Read"],
                "deny": ["Write", "WebFetch"]
              }
            }
```

### Lệnh tự động code review

**.claude/commands/review.md**
```markdown
---
name: review
description: Thực hiện code review PR
allowed-tools: Read, Grep, Glob, Bash(git diff:*)
---

Review thay đổi trên nhánh hiện tại:

1. Chạy `git diff main...HEAD`
2. Kiểm tra lỗ hổng bảo mật (SQL Injection, XSS, CSRF)
3. Gắn cờ vấn đề hiệu năng (N+1 query, thuật toán kém)
4. Kiểm tra tuân thủ quy ước code
5. Xem xét độ phủ test
6. Chỉ báo cáo vấn đề có độ tin cậy ≥ 80%
```

---

## Gói đăng ký đội và cách dùng cộng tác

### Tính năng gói Team Claude Code

Gói đăng ký đội Claude Code có **Standard Seat ($25/tháng)** và **Premium Seat ($150/tháng)**. Tính năng cộng tác chính:

- **Quản lý tập trung**: Gán seat, theo dõi sử dụng, kiểm soát chi phí
- **Connector**: GitHub, Slack, Google Drive, Gmail
- **Giới hạn sử dụng**: ~**225 tin nhắn** mỗi 5 giờ, **50–95 giờ** Sonnet 4 mỗi tuần
- **Vượt hạn**: Dùng thêm qua API trong giới hạn admin cấu hình

### Mẫu quy trình cho đội 3 người

**Onboarding:**
1. Clone repo (CLAUDE.md, settings.json đã kèm)
2. Chạy `claude` trong terminal
3. Dùng `/init` để xác nhận môi trường
4. Tạo `CLAUDE.local.md` nếu cần ghi đè cá nhân

**Làm việc song song:**
```bash
# Chạy nhiều phiên Claude song song bằng Git worktree
git worktree add ../rmn-feature-auth feature/auth
git worktree add ../rmn-feature-ads feature/ads
# Chạy phiên Claude độc lập trong mỗi worktree
```

---

## Sai lầm thường gặp và cách tránh

### Năm sai lầm nên tránh

| Sai lầm | Hệ quả | Cách xử lý |
|---------|--------|------------|
| Viết CLAUDE.md quá dài | Tuân thủ hướng dẫn giảm | Giữ 60–300 dòng, dùng progressive disclosure |
| Dùng /init không review | Có thể chứa thông tin sai | Luôn review và sửa |
| Dùng Claude làm linter | Tốn chi phí, chậm | Dùng hook cho gofmt, dart format |
| Đưa thông tin nhạy cảm | Rủi ro bảo mật | Không bao giờ đưa API key hay mật khẩu vào ngữ cảnh |
| Quản lý ngữ cảnh kém | Thông tin không liên quan tích tụ | Dùng `/clear` giữa các tác vụ |

### Tóm tắt best practices

- **CLAUDE.md là điểm đòn bẩy cao**: Viết và lặp như prompt
- **Bảo vệ bằng CODEOWNERS**: Yêu cầu review cấp cao cho thay đổi CLAUDE.md
- **Ưu tiên công cụ xác định**: Dùng hook cho linter/formatter; dùng Claude cho công việc sáng tạo
- **Dùng progressive disclosure**: Đặt tài liệu chi tiết trong `agent_docs/` và tham chiếu từ CLAUDE.md

---

## Checklist bắt đầu nhanh

1. ✅ Thêm `CLAUDE.md` ở gốc dự án (≤ 60 dòng để bắt đầu)
2. ✅ Cấu hình quyền đội trong `.claude/settings.json`
3. ✅ Thêm 3 subagent theo vai trò trong `.claude/agents/`
4. ✅ Thêm lệnh quy trình chính trong `.claude/commands/`
5. ✅ Thêm `CLAUDE.local.md`, `.claude/settings.local.json` vào `.gitignore`
6. ✅ Xác nhận thành viên sẵn sàng sau khi clone và chạy `claude`

Với cấu trúc này, đội 3 người có thể đạt **tự động hóa phát triển toàn bộ** cho nền tảng RMN SaaS trong **cùng môi trường Claude Code**.
