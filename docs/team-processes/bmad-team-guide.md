---
clickup_workspace_id: "90182277854"
clickup_doc_id: "2kzmgppy-1378"
title: "Hướng Dẫn Sử Dụng BMAD Cho Team"
---

# 📖 Từ Điển BMAD Cho Team RMN

> **Tài liệu hướng dẫn đầy đủ** cho team members mới và tham khảo khi cần.

---

## 🚀 Bước 1: Pull Code Mới Nhất

**⚠️ QUAN TRỌNG: Luôn pull trước khi bắt đầu làm việc!**

```bash
# Vào thư mục project
cd demo-structure

# Pull code mới nhất
git pull origin main
```

> 💡 **Chưa có repo trên máy?** (Chỉ dành cho member mới lần đầu)
> ```bash
> git clone https://github.com/mtsgn/demo-structure.git
> cd demo-structure
> ```

---

## 🛠️ Bước 2: Mở Claude Code

```bash
# Mở Claude Code tại thư mục demo-structure
claude
```

### Kiểm tra BMAD đã sẵn sàng
```bash
# Trong Claude Code, gõ:
/bmad-help
```

Nếu thấy menu workflows → BMAD đã hoạt động ✅

---

## 🎯 Bước 3: Chọn Agent Theo Role

Load agent trước khi bắt đầu để có context phù hợp:

| Role | Command | Agent Name |
|------|---------|------------|
| Business Analyst | `/bmad-agent-ba` | Mary 📊 |
| Product Manager | `/bmad-agent-pm` | John 📋 |
| Architect | `/bmad-agent-arch` | Winston 🏗️ |
| Developer (BE) | `/bmad-agent-dev` | Amelia 💻 |
| UX Designer | `/bmad-agent-ux` | Sally 🎨 |
| Scrum Master | `/bmad-agent-sm` | Bob 🏃 |
| QA/Test | `/bmad-agent-quinn` | Murat 🧪 |
| Solo Dev | `/bmad-agent-solo` | Barry 🚀 |
| Tech Writer | `/bmad-agent-writer` | Paige 📚 |

---

## 📋 Bước 4: Chạy Commands Theo Công Việc

### 📊 BA (Business Analyst)
| Command | Mô tả | Khi nào dùng |
|---------|-------|-------------|
| `/ba-create-brief` | Tạo product brief | Bắt đầu project mới |
| `/ba-research` | Research workflow | Cần thu thập thông tin |
| `/ba-brainstorm` | Brainstorming | Cần ý tưởng mới |
| `/ba-create-stories` | Tạo epics/stories | Sau khi có architecture |

### 📋 PM (Product Manager)
| Command | Mô tả | Khi nào dùng |
|---------|-------|-------------|
| `/pm-create-prd` | Tạo PRD | Sau product brief |
| `/pm-sprint-planning` | Lập kế hoạch sprint | Bắt đầu sprint |
| `/pm-sprint-status` | Xem trạng thái | Check progress |
| `/pm-retro` | Retrospective | Cuối epic |
| `/pm-correct-course` | Điều chỉnh hướng | Có thay đổi lớn |

### 🏗️ Architect
| Command | Mô tả | Khi nào dùng |
|---------|-------|-------------|
| `/arch-create` | Tạo architecture | Sau PRD |
| `/arch-diagram` | Tạo diagram | Cần visualize hệ thống |
| `/arch-dataflow` | Data flow diagram | Thiết kế data flow |
| `/arch-flowchart` | Flowchart | Vẽ luồng xử lý |
| `/arch-context` | Generate context | Onboarding project mới |
| `/arch-check-readiness` | Kiểm tra sẵn sàng | Trước khi code |

### 💻 Dev Backend
| Command | Mô tả | Khi nào dùng |
|---------|-------|-------------|
| `/dev-be-story` | Implement story | Làm story từ sprint |
| `/dev-be-review` | Code review | Sau khi code xong |
| `/dev-be-quick` | Quick dev | Task nhỏ, bug fix |
| `/dev-be-create-story` | Tạo story | Chuẩn bị story |
| `/dev-be-docs` | Document project | Viết docs |

### 🎨 Dev Frontend
| Command | Mô tả | Khi nào dùng |
|---------|-------|-------------|
| `/dev-fe-ux` | Tạo UX design | Thiết kế UX |
| `/dev-fe-wireframe` | Tạo wireframe | Mockup UI |
| `/dev-fe-spec` | Quick spec | Spec nhỏ |

### 🧪 QA
| Command | Mô tả | Khi nào dùng |
|---------|-------|-------------|
| `/qa-automate` | QA automation | Setup test automation |

---

## 🔄 Bước 5: Commit và Push

Sau khi hoàn thành công việc:

```bash
# 1. Kiểm tra thay đổi
git status

# 2. Add files
git add .

# 3. Commit với message rõ ràng
git commit -m "feat: Mô tả ngắn gọn thay đổi"

# 4. Push lên
git push origin main
```

### Commit Message Convention
| Prefix | Khi nào dùng |
|--------|-------------|
| `feat:` | Thêm feature mới |
| `fix:` | Sửa bug |
| `docs:` | Cập nhật documentation |
| `refactor:` | Refactor code |
| `chore:` | Các việc khác |

---

## 📁 Cấu Trúc Thư Mục Quan Trọng

```
demo-structure/
├── .claude/              ← 🎯 AI IDE integration
│   ├── commands/         ← Slash commands (41 files)
│   ├── agents/           ← Custom agents
│   ├── hooks/            ← Automation hooks
│   └── scripts/          ← Helper scripts
│
├── _bmad/                ← 🧠 BMAD Framework core
│   ├── core/             ← Core module
│   └── bmm/              ← BMM module (9 agents, 30 workflows)
│
├── _bmad-output/         ← 📦 Output artifacts
│   ├── planning_artifacts/       ← PRD, architecture, epics
│   └── implementation_artifacts/ ← Sprints, stories
│
└── docs/                 ← 📚 Documentation
    ├── bmad_onboarding/  ← Học BMAD
    └── team-processes/   ← File này
```

---

## 🚦 Luồng Làm Việc Chuẩn

### Project Mới (Full Flow)
```
1. /ba-create-brief    → Product Brief
2. /pm-create-prd      → PRD
3. /arch-create        → Architecture
4. /ba-create-stories  → Epics & Stories
5. /arch-check-readiness → Quality Check
6. /pm-sprint-planning → Sprint Plan
7. /dev-be-story       → Implement (lặp lại)
8. /dev-be-review      → Code Review
9. /pm-retro           → Retrospective
```

### Task Nhỏ/Bug Fix (Quick Flow)
```
/dev-be-quick          → Direct implementation
```

### Đang Stuck?
```
/bmad-help             → Được hướng dẫn next step
```

---

## ❓ FAQ

### Q: Mới vào team, bắt đầu từ đâu?
**A:** Clone repo → Đọc `docs/bmad_onboarding/01-getting-started/` → Thử `/bmad-help`

### Q: Không biết dùng command nào?
**A:** Gõ `/bmad-help` để được gợi ý

### Q: Muốn hiểu sâu hơn về BMAD?
**A:** Đọc `docs/bmad_onboarding/`

### Q: Có lỗi khi chạy command?
**A:** Kiểm tra: Đã `git pull` chưa? Đang ở đúng thư mục chưa? Thử mở lại Claude Code.

---

## 📞 Hỗ Trợ

- **Docs chi tiết:** `docs/bmad_onboarding/`
- **Commands reference:** `docs/bmad_onboarding/05-references/02-commands-reference.md`
- **Best practices:** `docs/bmad_onboarding/05-references/01-best-practices-by-role.md`
