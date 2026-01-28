# 🚀 Quick Start Guide - BMAD + Everything CC

**Setup Date:** January 28, 2026  
**Target:** Sẵn sàng trước Tết Nguyên Đán

---

## 📋 Bước 1: Installation (30 phút)

```bash
# Chạy script cài đặt
cd "/Users/mazhnguyen/Library/CloudStorage/GoogleDrive-work.huutrung@gmail.com/My Drive/KwayVina/RMN"
chmod +x install-bmad-enhanced.sh
./install-bmad-enhanced.sh
```

**Kết quả:**
- ✅ BMAD Method (đã có)
- ✅ Everything CC tools (thêm vào)
- ✅ Continuous learning active
- ✅ No conflicts

---

## 🎯 Bước 2: Verify Installation (5 phút)

```bash
# Check structure
ls -la .claude/
ls -la .claude/agents/
ls -la .claude/skills/

# Should see:
# .claude/agents/ec-security.md
# .claude/agents/ec-tdd-guide.md
# .claude/skills/ec-learning/
```

---

## 📚 Bước 3: Hiểu Command Priority (10 phút)

### **BMAD Commands (Primary - Luôn dùng):**

```yaml
Project Management:
  @bmad-help                      # Help command
  @bmad-brainstorming            # Brainstorm ideas
  @bmad-bmm-analyst              # Business analysis
  @bmad-bmm-pm                   # Project management

Planning:
  @bmad-bmm-create-product-brief # Product brief
  @bmad-bmm-create-prd          # PRD
  @bmad-bmm-create-architecture  # Architecture
  @bmad-bmm-create-epics-and-stories # User stories

Implementation:
  @bmad-bmm-sprint-planning      # Sprint planning
  @bmad-bmm-dev-story           # Development
  @bmad-bmm-code-review         # Code review

Testing:
  @bmad-bmm-testarch-framework   # Test framework
  @bmad-bmm-testarch-automate    # Test automation
```

### **Everything CC (Specialists - Khi cần):**

```yaml
ECC có 20 commands (đã bỏ conflicts):

Learning & Evolution:
  /instinct-status        # Check learned patterns
  /instinct-export        # Export patterns
  /instinct-import        # Import patterns
  /evolve                 # Generate workflows from patterns
  /learn                  # Manual learning trigger

Development:
  /tdd                    # TDD workflow
  /e2e                    # E2E testing
  /build-fix              # Fix build errors
  /refactor-clean         # Refactor code cleanup

Quality:
  /verify                 # Verification loop
  /eval                   # Eval harness
  /test-coverage          # Test coverage

Go-specific:
  /go-build               # Go build
  /go-test                # Go test
  /go-review              # Go review

Documentation:
  /update-docs            # Update docs
  /update-codemaps        # Update code maps

Project:
  /checkpoint             # Save checkpoint
  /setup-pm               # Setup PM
  /skill-create           # Create custom skill

❌ Removed (conflicts with BMAD):
  /plan → use @bmad-bmm-pm
  /code-review → use @bmad-bmm-code-review
  /orchestrate → use @bmad-party-mode
```

---

## 🎮 Bước 4: First Workflow (30 phút)

### **Test 1: Brainstorming**

```bash
# Trong AI IDE (Claude Code/Windsurf):
@bmad-brainstorming

# AI sẽ hỏi về project idea
# Follow prompts
# Check output: _bmad-output/planning-artifacts/brainstorming/
```

### **Test 2: Product Brief**

```bash
@bmad-bmm-create-product-brief

# Create a simple product brief
# Check: _bmad-output/planning-artifacts/product-brief/
```

### **Test 3: Check Learning**

```bash
# After some sessions, check learned patterns
# In chat: "Show me learned patterns" or "check instinct status"
# Learning system observes in background
```

---

## 📖 Bước 5: Typical Development Flow

### **Workflow A: New Project từ đầu**

```yaml
Week 1: Discovery
  Day 1: @bmad-brainstorming
  Day 2: @bmad-bmm-research (nếu cần)
  Day 3: @bmad-bmm-create-product-brief
  Output: Product brief

Week 1-2: Planning
  Day 4: @bmad-bmm-create-prd
  Day 5: @bmad-bmm-create-architecture
  Day 6: @bmad-bmm-create-epics-and-stories
  Output: PRD, Architecture, Stories

Week 2-3: Implementation
  Day 7: @bmad-bmm-sprint-planning
  Day 8-10: @bmad-bmm-dev-story (multiple times)
           → ec-tdd-guide auto helps with TDD
           → ec-learning observes patterns
  Day 11: @bmad-bmm-code-review
          → ec-security reviews if needed
  Output: Working code + tests
```

### **Workflow B: Quick Feature**

```yaml
Single Feature (1-2 days):
  1. @bmad-bmm-quick-spec
     → Create quick spec
  
  2. @bmad-bmm-quick-dev
     → Implement with TDD (ec-tdd-guide helps)
  
  3. @bmad-bmm-code-review
     → Quality check (ec-security if needed)
  
  4. Done!
     → ec-learning captured pattern
```

---

## 🔧 Bước 6: Common Commands

### **Getting Help:**

```bash
@bmad-help                    # General help
"What workflows available?"   # List workflows
"How to start project?"       # Guidance
```

### **Check Status:**

```bash
# Project status
@bmad-bmm-sprint-status

# Learned patterns (after using a while)
"Show learned patterns"
"Check instinct status"
```

### **Quality Checks:**

```bash
# Code review
@bmad-bmm-code-review

# Security (auto-invoked when needed)
# Just mention: "Check security"

# Refactoring
"Clean up dead code"  # ec-refactor helps
```

---

## 🎯 Workflow Cheat Sheet

```
┌─────────────────────────────────────────────────┐
│           COMMAND QUICK REFERENCE               │
├─────────────────────────────────────────────────┤
│                                                 │
│ START PROJECT:                                  │
│   @bmad-brainstorming                          │
│   @bmad-bmm-create-product-brief               │
│                                                 │
│ PLAN PROJECT:                                   │
│   @bmad-bmm-create-prd                         │
│   @bmad-bmm-create-architecture                │
│   @bmad-bmm-create-epics-and-stories           │
│                                                 │
│ IMPLEMENT:                                      │
│   @bmad-bmm-sprint-planning                    │
│   @bmad-bmm-dev-story (multiple)               │
│   @bmad-bmm-code-review                        │
│                                                 │
│ QUICK WORK:                                     │
│   @bmad-bmm-quick-spec                         │
│   @bmad-bmm-quick-dev                          │
│                                                 │
│ HELP:                                           │
│   @bmad-help                                   │
│                                                 │
└─────────────────────────────────────────────────┘
```

---

## 📊 Integration Architecture

```
USER REQUEST
      │
      ▼
┌─────────────────────┐
│   BMAD Workflows    │ ← You invoke these
│   (Primary)         │
└─────────────────────┘
      │
      ├─→ Needs security? ──→ ec-security (auto)
      │
      ├─→ Needs TDD? ────────→ ec-tdd-guide (auto)
      │
      ├─→ Needs refactor? ───→ ec-refactor (auto)
      │
      └─→ Learning ──────────→ ec-learning (always on)
                                   ↓
                              Captures patterns
                                   ↓
                              Can evolve later
```

---

## ⚠️ Important Notes

### **Do's:**

```yaml
✅ Always use BMAD commands (@bmad-xxx)
✅ Let ECC tools activate automatically
✅ Check learned patterns weekly
✅ Read _bmad-output/ for all artifacts
✅ Follow BMAD phases (Discovery → Planning → Implementation)
```

### **Don'ts:**

```yaml
❌ Don't use ECC commands directly (no /tdd, /code-review)
❌ Don't skip BMAD workflows
❌ Don't manually invoke ECC agents (let BMAD delegate)
❌ Don't mix methodologies (stick to BMAD flow)
```

---

## 🎊 Timeline Trước Tết

```
TODAY (Jan 28):
  ✅ Run install-bmad-enhanced.sh (30 min)
  ✅ Read this QUICK-START.md (30 min)
  ✅ Test @bmad-brainstorming (30 min)
  Total: 1.5 hours

TOMORROW (Jan 29):
  ✅ Try @bmad-bmm-create-product-brief (1 hour)
  ✅ Try @bmad-bmm-create-prd (1 hour)
  ✅ Practice workflows (1 hour)
  Total: 3 hours

Jan 30-31:
  ✅ Complete 1 small project end-to-end
  ✅ Test all main workflows
  ✅ Get comfortable
  Total: 4-6 hours

BEFORE TẾT:
  ✅ Comfortable với BMAD + ECC
  ✅ Ready for real projects
  ✅ Learning system active
```

---

## 🚀 Quick Troubleshooting

### **Problem: Commands không work**

```bash
# Check if in correct directory
pwd
# Should be: /Users/mazhnguyen/.../RMN

# Check .agent folder exists
ls .agent/workflows/

# Try with @ prefix
@bmad-help
```

### **Problem: Không thấy output**

```bash
# Check output folder
ls _bmad-output/planning-artifacts/
ls _bmad-output/implementation-artifacts/

# Artifacts are created after workflow completes
```

### **Problem: Confusion về command nào dùng**

```bash
# Rule đơn giản:
# - All commands start with @bmad-xxx
# - ECC is invisible (auto-activated)
# - When in doubt: @bmad-help
```

---

## 📖 Additional Resources

```
Read These (in order):
  1. .claude/PROJECT-CONFIG.md (integration details)
  2. notes/01-getting-started/01-bmad-quickstart.md
  3. notes/02-core-concepts/01-folder-structure.md
  4. notes/03-workflows/01-all-workflows.md

Check These Regularly:
  - _bmad-output/ (your artifacts)
  - .claude/skills/ec-learning/instincts/ (learned patterns)
```

---

## ✅ Success Checklist

```yaml
Installation:
  □ install-bmad-enhanced.sh ran successfully
  □ .claude/ folder exists with agents/ and skills/
  □ No error messages during install

First Test:
  □ @bmad-help works
  □ @bmad-brainstorming works
  □ Output appears in _bmad-output/

Understanding:
  □ Know to use @bmad-xxx commands
  □ Know ECC is background/auto
  □ Know where to find outputs
  □ Read PROJECT-CONFIG.md

Ready:
  □ Comfortable with basic workflows
  □ Can start real project
  □ Learning system active
  □ Sẵn sàng sau Tết! 🎊
```

---

## 🎯 Final Words

**You have:**
- ✅ BMAD Method (complete framework)
- ✅ Everything CC (quality tools)
- ✅ Continuous learning (auto-improve)
- ✅ No conflicts (properly configured)
- ✅ Ready to use!

**Remember:**
1. Use BMAD commands (@bmad-xxx)
2. Let ECC work in background
3. Check learned patterns weekly
4. Trust the system!

**Chúc mừng năm mới! 🎊🎉**

---

*Last Updated: January 28, 2026*
