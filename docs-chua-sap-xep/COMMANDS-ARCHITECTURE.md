# 📖 Commands, Skills, Agents Architecture

Giải thích cấu trúc và cách hoạt động của Commands, Skills, Agents trong Claude Code.

---

## 🏗️ Component Hierarchy

```
┌──────────────────────────────────────────────┐
│              CLAUDE CODE IDE                 │
└──────────────────────────────────────────────┘
                    │
      ┌─────────────┼─────────────┐
      │             │             │
      ▼             ▼             ▼
 ┌─────────┐  ┌─────────┐  ┌─────────┐
 │COMMANDS │  │ SKILLS  │  │ AGENTS  │
 └─────────┘  └─────────┘  └─────────┘
      │             │             │
      └─────────────┼─────────────┘
                    │
      ┌─────────────┴─────────────┐
      │                           │
      ▼                           ▼
 ┌──────────┐              ┌──────────┐
 │ SCRIPTS  │              │  TOOLS   │
 └──────────┘              └──────────┘
```

---

## 🎯 Component Definitions

### **1. Commands** (Slash Commands)

**File:** `.claude/commands/*.md`

**Purpose:** User-invoked actions (như terminal commands)

**Invocation:** `/command-name`

**Structure:**
```markdown
---
name: command-name
description: What this command does
command: /command-name
implementation: path/to/script.sh or python script.py
---

# Command Documentation
Instructions for AI on how to execute this command...
```

**Can Call:**
- ✅ **Agents** (delegate to specific agent)
- ✅ **Skills** (activate skills)
- ✅ **Scripts** (run shell/python scripts)
- ✅ **Tools** (use file operations, run commands)
- ✅ **Other Commands** (chain commands)
- ✅ **Workflows** (trigger BMAD workflows)

**Examples:**
```yaml
/tdd
  → Calls: ec-tdd-guide agent
  → Uses: tdd-workflow skill
  → Runs: test commands

/instinct-status
  → Runs: python3 instinct-cli.py status
  → Reads: .claude/skills/ec-learning/instincts/
  → Displays: learned patterns

/evolve
  → Runs: python3 instinct-cli.py evolve
  → Reads: instincts
  → Creates: new commands/skills/agents

/build-fix
  → Calls: ec-build-fixer agent
  → Runs: npm run build
  → Uses: file operations to fix errors
  → Loops: until errors resolved
```

---

### **2. Skills** (Auto-Activated Intelligence)

**File:** `.claude/skills/skill-name/SKILL.md`

**Purpose:** Auto-triggered behaviors based on keywords

**Invocation:** **AUTOMATIC** (no manual call needed)

**Structure:**
```markdown
---
name: skill-name
description: "Keywords that trigger this skill:
  Actions: design, create, build...
  Projects: dashboard, website...
  Elements: button, modal, form..."
---

# Skill Instructions
When activated, follow these guidelines...
```

**Can Call:**
- ✅ **Agents** (delegate complex tasks)
- ✅ **Scripts** (helper utilities)
- ✅ **Tools** (file operations)
- ✅ **Other Skills** (compose behaviors)
- ✅ **Data Files** (load patterns, templates)

**Examples:**
```yaml
ec-learning (continuous-learning-v2):
  Triggers: Always observing (background)
  Actions: Captures patterns, saves instincts
  Calls: observer agent, instinct-cli.py script
  
ui-ux-pro-max:
  Triggers: "design", "dashboard", "glassmorphism"
  Actions: Provides design patterns
  Reads: data/ (67 styles, 97 palettes)
  
ec-tdd:
  Triggers: "test", "TDD", "implement feature"
  Actions: Enforces TDD workflow
  Calls: ec-tdd-guide agent
  
ec-verify:
  Triggers: "verify", "validation", "check"
  Actions: Runs verification loops
  Uses: eval harness
```

---

### **3. Agents** (Specialized AI Personas)

**File:** `.claude/agents/agent-name.md`

**Purpose:** Specialized AI with specific expertise and instructions

**Invocation:** 
- Automatic (delegated by commands/skills)
- Manual: Mention in chat

**Structure:**
```markdown
---
name: agent-name
description: Agent specialization
role: Specific role/expertise
---

# Agent Instructions
You are a [specialized role]...

## Capabilities
- Capability 1
- Capability 2

## Process
1. Step 1
2. Step 2
```

**Can Call:**
- ✅ **Skills** (use supporting skills)
- ✅ **Tools** (file operations, terminal)
- ✅ **Scripts** (helper utilities)
- ✅ **Other Agents** (delegate sub-tasks)
- ✅ **Commands** (trigger workflows)

**Examples:**
```yaml
ec-tdd-guide:
  Role: TDD expert
  Called by: /tdd command, BMAD dev workflow
  Process: Define interfaces → Write tests → Implement → Refactor
  Uses: test frameworks, coverage tools
  
ec-security:
  Role: Security reviewer
  Called by: BMAD code-review, explicit mention
  Process: Scan for vulnerabilities → Check auth → Review secrets
  Uses: security scanning tools
  
ec-build-fixer:
  Role: Build error resolver
  Called by: /build-fix command
  Process: Run build → Parse errors → Fix one by one → Verify
  Uses: build tools, file operations
  
ec-refactor:
  Role: Code cleanup specialist
  Called by: /refactor-clean command
  Process: Find dead code → Categorize → Safe deletion → Test
  Uses: knip, depcheck, ts-prune tools
```

---

## 🔄 Integration Patterns

### **Pattern 1: Command → Agent → Tools**

```
USER: /tdd
  │
  ▼
COMMAND: tdd.md
  │ "invoke tdd-guide agent"
  ▼
AGENT: ec-tdd-guide.md
  │ "follow TDD workflow"
  ├─→ Define interfaces
  ├─→ Write tests (use file tools)
  ├─→ Run tests (use terminal)
  ├─→ Implement code (use file tools)
  ├─→ Verify coverage (use test tools)
  └─→ Return: Result
```

### **Pattern 2: Command → Script → Output**

```
USER: /instinct-status
  │
  ▼
COMMAND: instinct-status.md
  │ "run instinct-cli.py status"
  ▼
SCRIPT: instinct-cli.py
  │ Read .claude/skills/ec-learning/instincts/
  │ Parse instinct files
  │ Format output
  └─→ Display: Status report
```

### **Pattern 3: Skill Auto-Activates → Agent**

```
USER: "Create a glassmorphic dashboard"
  │
  ▼
CLAUDE AI detects keywords
  │
  ├─→ SKILL: ui-ux-pro-max
  │     │ "glassmorphic" + "dashboard" detected
  │     ├─→ Load design patterns
  │     ├─→ Load color palettes
  │     └─→ Provide design guidance
  │
  └─→ Return: Design + Code
```

### **Pattern 4: BMAD Workflow → Delegates**

```
USER: @bmad-bmm-dev-story
  │
  ▼
BMAD WORKFLOW: dev-story
  │ Load workflow.xml + dev-story/workflow.yaml
  │
  ├─→ Need TDD?
  │   └─→ SKILL: ec-tdd activates
  │       └─→ AGENT: ec-tdd-guide helps
  │
  ├─→ Need UI design?
  │   └─→ SKILL: ui-ux-pro-max activates
  │       └─→ Provides design patterns
  │
  ├─→ Need security check?
  │   └─→ AGENT: ec-security reviews
  │
  └─→ Learning observes
      └─→ SKILL: ec-learning captures patterns
```

### **Pattern 5: Command Chain**

```
USER: /evolve
  │
  ▼
COMMAND: evolve.md
  │ "cluster instincts into new structures"
  │
  ├─→ Run: instinct-cli.py evolve
  │   └─→ Analyzes patterns
  │
  ├─→ Creates: New command files
  │   └─→ Writes: .claude/commands/new-command.md
  │
  ├─→ Creates: New skill files
  │   └─→ Writes: .claude/skills/new-skill/SKILL.md
  │
  └─→ Creates: New agent files
      └─→ Writes: .claude/agents/new-agent.md
```

---

## 📋 What Commands Can Call

```yaml
Commands can invoke:

1. Agents:
   ✅ "invoke the tdd-guide agent"
   ✅ "delegate to ec-security agent"
   ✅ Load agent instructions and execute

2. Skills:
   ✅ "activate ui-ux-pro-max skill"
   ✅ "use continuous-learning skill"
   ✅ Skills can also auto-activate

3. Scripts:
   ✅ Shell scripts: bash script.sh
   ✅ Python scripts: python3 script.py
   ✅ Node scripts: node script.js
   ✅ Any executable: /path/to/binary

4. Tools (Built-in):
   ✅ File operations (read, write, edit)
   ✅ Terminal commands (run_in_terminal)
   ✅ Search (grep, semantic_search)
   ✅ Git operations
   ✅ Code analysis

5. Workflows:
   ✅ BMAD workflows: @bmad-xxx
   ✅ Custom workflows
   ✅ Multi-step orchestrations

6. Other Commands:
   ✅ Chain commands: /tdd → /verify → /test-coverage
   ✅ Compose workflows

7. Data Files:
   ✅ JSON configs
   ✅ YAML configs
   ✅ Templates
   ✅ Knowledge bases
```

---

## 🎯 Real Examples from Your Setup

### **Example 1: `/tdd` Command**

```markdown
File: .claude/commands/tdd.md

Command invokes:
  → Agent: ec-tdd-guide
    → Uses: TDD methodology
    → Calls: test frameworks (jest, pytest, etc.)
    → Uses: file tools to write tests
    → Uses: terminal to run tests
    → Uses: coverage tools to verify
```

### **Example 2: `/evolve` Command**

```markdown
File: .claude/commands/evolve.md

Command runs:
  → Script: python3 instinct-cli.py evolve
    → Reads: .claude/skills/ec-learning/instincts/
    → Analyzes: pattern clusters
    → Creates: new commands/skills/agents
    → Writes: .claude/commands/*.md
    → Writes: .claude/skills/*/SKILL.md
    → Writes: .claude/agents/*.md
```

### **Example 3: `ui-ux-pro-max` Skill**

```markdown
File: .claude/skills/ui-ux-pro-max/SKILL.md

Skill activates on keywords:
  "glassmorphism", "dashboard", "design", "button"

When activated:
  → Reads: data/ (67 styles, 97 palettes, 57 fonts)
  → Applies: design rules (100 reasoning rules)
  → Provides: code templates
  → Can call: color generation scripts
  → Can call: font pairing utilities
```

### **Example 4: BMAD `@bmad-bmm-dev-story` Workflow**

```markdown
File: .agent/workflows/bmad-bmm-dev-story.md

Workflow orchestrates:
  → Loads: workflow.xml (core OS)
  → Loads: dev-story/workflow.yaml (config)
  → Delegates to:
    - ec-tdd-guide (if TDD needed)
    - ui-ux-pro-max (if UI needed)
    - ec-security (if security review needed)
    - ec-learning (always observing)
  → Uses: file tools, terminal, git
  → Saves: artifacts to _bmad-output/
```

---

## ✅ Summary Table

| Component | Location | Invocation | Can Call |
|-----------|----------|------------|----------|
| **Command** | `.claude/commands/*.md` | `/command-name` | Agents, Skills, Scripts, Tools, Workflows, Commands |
| **Skill** | `.claude/skills/*/SKILL.md` | Auto (keywords) | Agents, Scripts, Tools, Skills, Data |
| **Agent** | `.claude/agents/*.md` | Auto (delegated) or manual mention | Skills, Tools, Scripts, Agents, Commands |
| **Workflow** | `.agent/workflows/*.md` (BMAD) | `@bmad-xxx` | Agents, Skills, Commands, Tools, Scripts |
| **Script** | `.claude/skills/*/scripts/*.py` | Called by commands/skills/agents | Tools, Other scripts |

---

## 🎮 Usage Guidelines

### **When to Use Commands:**

```bash
# Explicit workflows you want to invoke
/tdd              # Start TDD workflow
/instinct-status  # Check learned patterns
/evolve           # Generate new structures
/build-fix        # Fix build errors
```

### **When Skills Auto-Activate:**

```bash
# Just mention keywords naturally
"Create a glassmorphic dashboard"  # → ui-ux-pro-max
"Implement feature X with tests"   # → ec-tdd
"Review security vulnerabilities"   # → ec-security (via agent)
```

### **When to Use BMAD:**

```bash
# Full project workflows
@bmad-bmm-dev-story        # Development workflow
@bmad-bmm-code-review      # Code review workflow
@bmad-bmm-create-prd       # Planning workflow
# BMAD orchestrates all sub-components automatically
```

---

## 🔗 Integration Best Practices

```yaml
Priority:
  1. BMAD Workflows (primary orchestrator)
     → @bmad-xxx commands
     
  2. ECC Commands (specialized tasks)
     → /tdd, /evolve, /instinct-status
     
  3. Skills (automatic support)
     → Auto-activate based on context
     
  4. Direct Agent Mention (rarely needed)
     → "Use ec-security agent to review..."

Rule:
  ✅ Let BMAD orchestrate
  ✅ Let skills auto-activate
  ✅ Use commands for explicit workflows
  ✅ Rarely call agents directly
```

---

**Last updated: January 28, 2026**
