# 📦 Plugin Installation Guide

Cách cài thêm plugins vào BMAD + ECC setup.

---

## 🎯 Convention: `.claude/` Folder

**TẤT CẢ plugins/skills/agents đều vào `.claude/`:**

```
.claude/
├── agents/          ← Tất cả agents (từ mọi plugins)
├── skills/          ← Tất cả skills (từ mọi plugins)
├── commands/        ← Tất cả commands (từ mọi plugins)
├── hooks/           ← Event hooks
└── PROJECT-CONFIG.md
```

**Giống như:**
- `.vscode/` → VS Code settings
- `node_modules/` → NPM packages
- **`.claude/` → Claude Code plugins**

---

## 📥 Cài Plugin Mới

### **Step 1: Clone Repository**

```bash
cd /tmp
git clone --depth 1 <plugin-repo-url>
```

### **Step 2: Check Structure**

```bash
ls -la <plugin-folder>/.claude/

# Should see:
# - agents/ (nếu có agents)
# - skills/ (nếu có skills)
# - commands/ (nếu có commands)
# - .claude-plugin/ (plugin metadata)
```

### **Step 3: Copy to Project**

```bash
# Copy everything to your project
cp -r <plugin-folder>/.claude/* .claude/

# Hoặc selective copy:
cp -r <plugin-folder>/.claude/skills/skill-name .claude/skills/
```

### **Step 4: Check Conflicts**

```bash
# Check if any commands conflict with BMAD
ls .claude/commands/

# If conflicts exist, remove them:
rm .claude/commands/conflicting-command.md
```

### **Step 5: Verify**

```bash
ls .claude/agents/
ls .claude/skills/
ls .claude/commands/

# All installed plugins should appear
```

---

## 🎨 Example: UI/UX Pro Max

### **Install:**

```bash
# 1. Clone
cd /tmp
git clone --depth 1 https://github.com/nextlevelbuilder/ui-ux-pro-max-skill.git

# 2. Copy skill
cp -r ui-ux-pro-max-skill/.claude/skills/ui-ux-pro-max \
  "/path/to/RMN/.claude/skills/"

# 3. Verify
ls "/path/to/RMN/.claude/skills/ui-ux-pro-max/"
# Should see: SKILL.md, data/, scripts/

# 4. Done!
```

### **Usage:**

**Tự động (recommended):**
```bash
# Just mention keywords:
"Create a glassmorphic dashboard with bento grid layout"

# ui-ux-pro-max skill AUTO-ACTIVATES
# Provides: design patterns, colors, fonts
```

**Với BMAD:**
```bash
@bmad-bmm-ux-designer

# BMAD UX Designer workflow
# → Automatically uses ui-ux-pro-max skill if installed
# → No manual invocation needed!
```

**Explicit:**
```bash
"Use ui-ux-pro-max skill to design a modern landing page"
```

---

## 🔄 Cách Skills Hoạt Động

### **Auto-Activation:**

Skills tự động activate dựa trên **description** trong `SKILL.md`:

```yaml
# Example: ui-ux-pro-max/SKILL.md
description: "UI/UX design intelligence. 
  Actions: plan, build, create, design, implement...
  Projects: website, landing page, dashboard...
  Elements: button, modal, navbar, sidebar...
  Styles: glassmorphism, minimalism, brutalism..."
```

**Khi bạn mention keywords → skill activates!**

### **Integration với BMAD:**

```
YOU: "Design a dashboard"
  │
  ▼
@bmad-bmm-ux-designer
  │
  ├─→ BMAD orchestrates
  │
  ├─→ Need design patterns?
  │   → ui-ux-pro-max skill helps (if installed)
  │
  ├─→ Need TDD?
  │   → ec-tdd-guide helps
  │
  ├─→ Learning observes
  │   → ec-learning captures
  │
  └─→ Result: Complete UX design + code + tests
```

**Tất cả tự động!**

---

## ⚠️ Handling Conflicts

### **Check for Command Conflicts:**

```bash
# List all commands
ls .claude/commands/

# Example conflicts with BMAD:
# - plan.md → use @bmad-bmm-pm
# - code-review.md → use @bmad-bmm-code-review
# - orchestrate.md → use @bmad-party-mode
```

### **Remove Conflicts:**

```bash
cd .claude/commands/
rm plan.md code-review.md orchestrate.md

# Or keep and document in PROJECT-CONFIG.md
```

### **Priority Rules:**

1. **BMAD = Primary** (always use first)
2. **Plugin Specialists** (use when BMAD delegates)
3. **Learning** (always observing in background)

---

## 📋 Installed Plugins Checklist

```yaml
Current Setup:
  ✅ BMAD Method (v6.0.0-alpha.23)
     - 47 workflows
     - 10 agents
     - Location: _bmad/, .agent/workflows/
  
  ✅ Everything Claude Code (Selected)
     - 4 agents: ec-security, ec-tdd-guide, ec-build-fixer, ec-refactor
     - 4 skills: ec-learning, ec-tdd, ec-eval, ec-verify
     - 20 commands (conflicts removed)
     - Location: .claude/
  
  □ UI/UX Pro Max
     - 1 skill: ui-ux-pro-max
     - 67 styles, 97 palettes, 57 fonts
     - Auto-activates on design keywords
     - Location: .claude/skills/ui-ux-pro-max/ (if installed)
```

---

## 🚀 Quick Commands

```bash
# List installed plugins
ls .claude/agents/
ls .claude/skills/
ls .claude/commands/

# Test a skill
# (trong AI chat):
"Use <skill-name> skill to ..."

# Test BMAD delegation
@bmad-bmm-ux-designer  # Should use ui-ux-pro-max if installed

# Check learning
/instinct-status  # See captured patterns
```

---

## 📖 More Plugins

**Where to find:**
- GitHub search: "claude code skill"
- GitHub search: "claude code plugin"
- Claude Code marketplace (future)

**Popular plugins:**
- Everything Claude Code (installed ✅)
- UI/UX Pro Max (design intelligence)
- Code review plugins
- Language-specific tools (Go, Python, etc.)
- Testing frameworks
- Documentation generators

---

## ✅ Best Practices

```yaml
DO:
  ✅ Install plugins to .claude/
  ✅ Check for command conflicts
  ✅ Remove conflicting commands
  ✅ Let BMAD orchestrate (primary)
  ✅ Let skills auto-activate
  ✅ Document installed plugins in PROJECT-CONFIG.md

DON'T:
  ❌ Install plugins to custom folders
  ❌ Skip conflict checks
  ❌ Manually invoke skills (let them auto-activate)
  ❌ Mix multiple primary orchestrators
  ❌ Override BMAD commands
```

---

**Last updated: January 28, 2026**
