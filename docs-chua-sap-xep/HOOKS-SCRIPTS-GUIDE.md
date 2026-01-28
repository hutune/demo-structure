# 🪝 Hooks & Scripts Architecture

Giải thích chi tiết về Hooks (event-driven automation) và Scripts (helper utilities) trong Claude Code.

---

## 🎯 Overview

```
┌────────────────────────────────────────────┐
│         CLAUDE CODE IDE                    │
└────────────────────────────────────────────┘
                │
    ┌───────────┴───────────┐
    │                       │
    ▼                       ▼
┌─────────┐           ┌─────────┐
│  HOOKS  │           │ SCRIPTS │
│(Events) │           │(Helpers)│
└─────────┘           └─────────┘
    │                       │
    │ Auto-trigger          │ Explicitly called
    │ on events             │ by commands/skills
    │                       │
    ▼                       ▼
Execute                 Execute
automatically          when invoked
```

---

## 🪝 HOOKS: Event-Driven Automation

### **Hooks Là Gì?**

**Hooks = Tự động chạy khi có events nhất định**

- Không cần gọi manual
- Trigger bởi events trong IDE
- Chạy trong background
- Có thể block/modify actions

**Location:** `.claude/hooks/hooks.json`

---

### **Hook Events (Khi Nào Chạy?)**

```yaml
PreToolUse:
  Trigger: TRƯỚC KHI tool sẽ chạy
  Purpose: Validate, block, modify
  Examples:
    - Block dev server outside tmux
    - Check console.log before commit
    - Validate file paths
    - Security checks

PostToolUse:
  Trigger: SAU KHI tool đã chạy xong
  Purpose: Post-process, log, analyze
  Examples:
    - Log PR URL after creation
    - Analyze build output
    - Save session state
    - Update instincts

SessionStart:
  Trigger: KHI BẮT ĐẦU session mới
  Purpose: Setup, restore state
  Examples:
    - Load previous context
    - Detect package manager
    - Initialize settings
    - Restore workspace

SessionEnd:
  Trigger: KHI KẾT THÚC session
  Purpose: Cleanup, save state
  Examples:
    - Save learned patterns
    - Backup context
    - Generate summary
    - Export instincts

PreCompact:
  Trigger: TRƯỚC KHI compact context (khi context đầy)
  Purpose: Save important state
  Examples:
    - Save current task state
    - Backup conversation
    - Mark important messages
    - Snapshot progress

Stop:
  Trigger: SAU MỖI AI RESPONSE
  Purpose: Quality checks, reminders
  Examples:
    - Check console.log in code
    - Verify test coverage
    - Check for TODOs
    - Evaluate session quality
```

---

### **Hook Structure (hooks.json)**

```json
{
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "tool == \"Bash\" && tool_input.command matches \"git push\"",
        "hooks": [
          {
            "type": "command",
            "command": "node script.js"
          }
        ],
        "description": "What this hook does"
      }
    ],
    
    "PostToolUse": [...],
    "SessionStart": [...],
    "Stop": [...]
  }
}
```

**Components:**
- `matcher`: Điều kiện trigger (tool name, command pattern)
- `hooks`: Array of actions to execute
- `type`: "command" (run script)
- `command`: Script to execute
- `description`: Human-readable explanation

---

### **Real Examples from Your Setup**

#### **Example 1: Block Dev Server Outside Tmux (PreToolUse)**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"(npm run dev|pnpm dev)\"",
  "hooks": [
    {
      "type": "command",
      "command": "node -e \"console.error('[Hook] BLOCKED: Dev server must run in tmux');process.exit(1)\""
    }
  ],
  "description": "Block dev servers outside tmux - ensures you can access logs"
}
```

**Trigger:** Khi AI sắp chạy `npm run dev`  
**Action:** Block command và show error  
**Why:** Đảm bảo dev server chạy trong tmux để access logs

#### **Example 2: Git Push Reminder (PreToolUse)**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"git push\"",
  "hooks": [
    {
      "type": "command",
      "command": "node -e \"console.error('[Hook] Review changes before push...')\""
    }
  ],
  "description": "Reminder before git push to review changes"
}
```

**Trigger:** Trước khi `git push`  
**Action:** Show reminder  
**Why:** Nhắc nhở review code trước khi push

#### **Example 3: Block Random .md Files (PreToolUse)**

```json
{
  "matcher": "tool == \"Write\" && tool_input.file_path matches \"\\\\.(md|txt)$\" && !(tool_input.file_path matches \"README\\\\.md\")",
  "hooks": [
    {
      "type": "command",
      "command": "node -e \"console.error('[Hook] BLOCKED: Use README.md for docs');process.exit(1)\""
    }
  ],
  "description": "Block creation of random .md files - keeps docs consolidated"
}
```

**Trigger:** Khi AI sắp tạo .md file (không phải README.md)  
**Action:** Block và suggest dùng README.md  
**Why:** Tránh tạo nhiều .md files rải rác

#### **Example 4: Log PR URL (PostToolUse)**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"gh pr create\"",
  "hooks": [
    {
      "type": "command",
      "command": "node -e \"/* extract PR URL from output and log it */\""
    }
  ],
  "description": "Log PR URL and provide review command after PR creation"
}
```

**Trigger:** Sau khi `gh pr create` chạy xong  
**Action:** Extract PR URL và show review command  
**Why:** Dễ dàng review PR vừa tạo

#### **Example 5: Session Start (SessionStart)**

```json
{
  "matcher": "*",
  "hooks": [
    {
      "type": "command",
      "command": "node \"${CLAUDE_PLUGIN_ROOT}/scripts/hooks/session-start.js\""
    }
  ],
  "description": "Load previous context and detect package manager on new session"
}
```

**Trigger:** Khi bắt đầu session mới  
**Action:** Load previous context, detect package manager  
**Why:** Restore state từ session trước

#### **Example 6: Check Console.log (Stop)**

```javascript
// check-console-log.js
// Runs after EACH AI response
// Checks git diff for console.log statements

const files = execSync('git diff --name-only HEAD')
  .split('\n')
  .filter(f => /\.(ts|tsx|js|jsx)$/.test(f));

for (const file of files) {
  const content = fs.readFileSync(file, 'utf8');
  if (content.includes('console.log')) {
    console.error(`[Hook] WARNING: console.log found in ${file}`);
  }
}
```

**Trigger:** Sau mỗi AI response  
**Action:** Scan modified files cho console.log  
**Why:** Nhắc nhở xóa debug statements

---

## 📜 SCRIPTS: Helper Utilities

### **Scripts Là Gì?**

**Scripts = Helper programs được gọi EXPLICITLY**

- Không tự động chạy
- Được gọi bởi: Commands, Skills, Agents, Hooks
- Thực hiện tasks cụ thể
- Có thể dùng bất kỳ language (JS, Python, Bash, etc.)

**Locations:**
- `.claude/scripts/hooks/*.js` → Scripts for hooks
- `.claude/skills/*/scripts/*.py` → Scripts for skills
- `.claude/agents/*/scripts/*` → Scripts for agents

---

### **Script Types**

#### **1. Hook Scripts**

**Purpose:** Execute logic for hooks

**Location:** `.claude/scripts/hooks/`

**Examples:**
```
check-console-log.js      → Check for console.log
evaluate-session.js       → Evaluate session quality
pre-compact.js            → Save state before compact
session-end.js            → Cleanup at session end
session-start.js          → Setup at session start
suggest-compact.js        → Suggest when to compact
```

**Called by:** Hooks automatically

#### **2. Skill Scripts**

**Purpose:** Helper logic for skills

**Location:** `.claude/skills/skill-name/scripts/`

**Example: instinct-cli.py**
```python
#!/usr/bin/env python3
"""
Instinct CLI - Manage instincts for Continuous Learning v2

Commands:
  status   - Show all instincts and their status
  import   - Import instincts from file or URL
  export   - Export instincts to file
  evolve   - Cluster instincts into skills/commands/agents
"""

# Can be called by:
# - /instinct-status command
# - /evolve command
# - ec-learning skill
# - Hooks (SessionEnd)
```

**Called by:** Commands (`/instinct-status`, `/evolve`)

#### **3. Agent Scripts**

**Purpose:** Helper logic for agents

**Location:** `.claude/agents/agent-name/scripts/`

**Examples:**
- Security scanning scripts
- Test generation scripts
- Code analysis scripts
- Refactoring utilities

**Called by:** Agents during execution

#### **4. Utility Scripts**

**Purpose:** General utilities

**Location:** `.claude/scripts/utils/`

**Examples:**
- Data processing
- File manipulation
- API calls
- Database queries

**Called by:** Any component

---

### **Script Invocation**

```yaml
From Commands:
  command: /instinct-status
  runs: python3 instinct-cli.py status

From Skills:
  skill: ec-learning
  runs: python3 instinct-cli.py evolve

From Agents:
  agent: ec-security
  runs: bash security-scan.sh

From Hooks:
  hook: SessionEnd
  runs: node session-end.js

From Other Scripts:
  script: evolve.py
  imports: from utils import cluster_patterns
```

---

## 🔄 Hooks vs Scripts

```
┌──────────────┬─────────────────┬─────────────────┐
│   Aspect     │     HOOKS       │    SCRIPTS      │
├──────────────┼─────────────────┼─────────────────┤
│ Trigger      │ Automatic       │ Explicit call   │
│              │ (event-based)   │ (manual invoke) │
├──────────────┼─────────────────┼─────────────────┤
│ When         │ PreToolUse,     │ When called by  │
│              │ PostToolUse,    │ commands/skills │
│              │ SessionStart,   │ agents/hooks    │
│              │ Stop, etc.      │                 │
├──────────────┼─────────────────┼─────────────────┤
│ Purpose      │ Automation,     │ Helper logic,   │
│              │ validation,     │ data processing │
│              │ guards          │ utilities       │
├──────────────┼─────────────────┼─────────────────┤
│ Can Block    │ ✅ Yes          │ ❌ No           │
│ Actions      │ (PreToolUse)    │ (just returns)  │
├──────────────┼─────────────────┼─────────────────┤
│ Examples     │ Block git push, │ instinct-cli.py │
│              │ Check logs,     │ analyze code    │
│              │ Save state      │ process data    │
└──────────────┴─────────────────┴─────────────────┘
```

---

## 🎮 Real Use Cases

### **Use Case 1: Continuous Learning Flow**

```
AI Response Completed
  │
  ▼
HOOK: Stop (PostToolUse)
  │ "After each response"
  │
  ├─→ SCRIPT: evaluate-session.js
  │     └─→ Analyze what happened
  │
  └─→ SKILL: ec-learning
        │ "Observe patterns"
        │
        └─→ SCRIPT: instinct-cli.py
              └─→ Save learned instincts
                  └─→ .claude/skills/ec-learning/instincts/
```

**Flow:**
1. User interacts with AI
2. AI completes response
3. **Hook triggers** (Stop event)
4. **Script runs** (evaluate-session.js)
5. **Skill activates** (ec-learning)
6. **Script called** (instinct-cli.py)
7. Instincts saved to disk

### **Use Case 2: Git Push Safety**

```
AI about to run: git push
  │
  ▼
HOOK: PreToolUse
  │ "Before tool execution"
  │
  ├─→ Matcher: "git push" detected
  │
  └─→ ACTION:
        ├─→ Show warning
        ├─→ Can block (exit 1)
        └─→ Or allow with reminder
```

**Flow:**
1. AI wants to `git push`
2. **Hook intercepts** (PreToolUse)
3. **Matcher checks** command
4. **Script runs** inline or external
5. Can **block** or **allow** with warning

### **Use Case 3: Session Restore**

```
New Session Starts
  │
  ▼
HOOK: SessionStart
  │
  ├─→ SCRIPT: session-start.js
  │     ├─→ Load .claude/homunculus/state.json
  │     ├─→ Restore previous context
  │     ├─→ Detect package manager
  │     └─→ Initialize settings
  │
  └─→ SESSION READY with restored state
```

**Flow:**
1. New session begins
2. **Hook triggers** automatically
3. **Script loads** previous state
4. Context restored seamlessly

### **Use Case 4: Command Calls Script**

```
USER: /instinct-status
  │
  ▼
COMMAND: instinct-status.md
  │ implementation: python3 instinct-cli.py status
  │
  └─→ SCRIPT: instinct-cli.py
        ├─→ Read: .claude/skills/ec-learning/instincts/
        ├─→ Parse instinct files
        ├─→ Calculate confidence
        └─→ Display formatted output

OUTPUT: Instinct status report
```

**Flow:**
1. User calls `/instinct-status`
2. **Command definition** specifies script
3. **Script runs** with arguments
4. Output returned to user

---

## 📋 Your Current Setup

### **Hooks Installed:**

```yaml
PreToolUse:
  1. Block dev server outside tmux
  2. Remind to use tmux for long commands
  3. Warn before git push
  4. Block random .md file creation
  5. Suggest manual compaction

PostToolUse:
  1. Log PR URL after creation
  2. Analyze build output

SessionStart:
  1. Load previous context
  2. Detect package manager

PreCompact:
  1. Save state before compaction

Stop:
  1. Check console.log in modified files
  2. Evaluate session quality
```

### **Scripts Installed:**

```yaml
Hook Scripts (.claude/scripts/hooks/):
  - check-console-log.js      → Check debug statements
  - evaluate-session.js       → Session quality
  - pre-compact.js            → Pre-compact save
  - session-end.js            → Cleanup
  - session-start.js          → Restore state
  - suggest-compact.js        → Compaction suggestions

Skill Scripts (.claude/skills/ec-learning/scripts/):
  - instinct-cli.py           → Manage instincts
    ├─ status command
    ├─ import command
    ├─ export command
    └─ evolve command
```

---

## ✅ Best Practices

### **When to Create Hooks:**

```yaml
✅ Safety checks:
   - Block dangerous commands
   - Validate inputs
   - Security checks

✅ Automation:
   - Save state automatically
   - Log important events
   - Update instincts

✅ Reminders:
   - Before git operations
   - Before deployments
   - After builds

✅ Quality gates:
   - Check console.log
   - Verify test coverage
   - Lint code
```

### **When to Create Scripts:**

```yaml
✅ Complex logic:
   - Data processing
   - File analysis
   - Pattern clustering

✅ Reusable utilities:
   - Shared between commands
   - Shared between skills
   - Helper functions

✅ External integrations:
   - API calls
   - Database queries
   - Service interactions

✅ Heavy computation:
   - Code analysis
   - ML/AI processing
   - Large file processing
```

---

## 🚀 Quick Reference

```bash
# Hooks = Automatic
PreToolUse     → Before tool runs (can block)
PostToolUse    → After tool completes
SessionStart   → New session begins
SessionEnd     → Session ends
PreCompact     → Before context compact
Stop           → After each AI response

# Scripts = Explicit
python3 script.py      → Called by commands
node script.js         → Called by hooks
bash script.sh         → Called by agents
```

**Hooks trigger automatically, Scripts are called explicitly!**

---

**Last updated: January 28, 2026**
