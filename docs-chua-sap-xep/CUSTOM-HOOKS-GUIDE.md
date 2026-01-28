# 🛠️ Customizing Hooks & Scripts Guide

Hướng dẫn tạo custom hooks và scripts cho project của bạn.

---

## 📦 Nguồn Gốc: Hooks & Scripts Có Sẵn Ở Đâu?

### **1. Everything Claude Code (ECC)**

```yaml
Source: https://github.com/affaan-m/everything-claude-code

Provides:
  hooks/hooks.json           → Hook definitions template
  scripts/hooks/*.js         → Hook implementation scripts
  skills/*/scripts/*.py      → Skill utility scripts
  
Installed to:
  .claude/hooks/hooks.json
  .claude/scripts/hooks/*.js
  .claude/skills/ec-learning/scripts/instinct-cli.py
```

**Hooks hiện có từ ECC:**
- ✅ Block dev server outside tmux
- ✅ Remind to use tmux for long commands
- ✅ Warn before git push
- ✅ Block random .md file creation
- ✅ Suggest manual compaction
- ✅ Log PR URL after creation
- ✅ Session start/end automation
- ✅ Check console.log

**Scripts hiện có từ ECC:**
- ✅ `check-console-log.js` - Check debug statements
- ✅ `session-start.js` - Restore session
- ✅ `session-end.js` - Save session
- ✅ `pre-compact.js` - Save before compact
- ✅ `suggest-compact.js` - Compaction hints
- ✅ `instinct-cli.py` - Manage learned patterns

---

### **2. BMAD Method**

```yaml
Source: BMAD framework (installed in _bmad/)

Provides:
  workflows/              → Workflow definitions
  agents/                 → Agent definitions
  tasks/                  → Task definitions
  
Does NOT provide:
  ❌ Hooks (không có)
  ❌ Scripts (không có)
```

**BMAD focuses on:**
- Workflows & orchestration
- Agent personas
- Artifact management
- Development methodology

**BMAD KHÔNG có hooks/scripts!**

---

### **3. UI/UX Pro Max**

```yaml
Source: https://github.com/nextlevelbuilder/ui-ux-pro-max-skill

Provides:
  .claude/skills/ui-ux-pro-max/SKILL.md
  src/ui-ux-pro-max/data/       → Design data
  src/ui-ux-pro-max/scripts/    → Helper scripts
  
Does NOT provide:
  ❌ Hooks (không có)
```

---

## ✅ Custom Hooks & Scripts: CÓ THỂ!

### **Bạn HOÀN TOÀN có thể custom:**

```yaml
Can Add:
  ✅ New hooks to .claude/hooks/hooks.json
  ✅ New scripts to .claude/scripts/
  ✅ Modify existing hooks
  ✅ Remove unwanted hooks
  ✅ Create project-specific automation

Should Not:
  ❌ Modify ECC source files directly
  ❌ Break existing hook structure
```

---

## 🎯 Khi Nào Nên Custom Hooks?

### **Scenario 1: Project-Specific Rules**

```yaml
Example: Enforce naming convention

Hook:
  Event: PreToolUse (file creation)
  Check: File name matches convention
  Action: Block if invalid
  
Why custom?
  ✅ Specific to YOUR project
  ✅ Not universal rule
  ✅ Team convention
```

**Implementation:**

```json
{
  "matcher": "tool == \"Write\" && tool_input.file_path matches \"\\.component\\.(tsx|jsx)$\"",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/check-component-name.js"
    }
  ],
  "description": "Enforce PascalCase for React components"
}
```

---

### **Scenario 2: Integration with External Tools**

```yaml
Example: Auto-update Jira on git commit

Hook:
  Event: PostToolUse (git commit)
  Action: Extract ticket ID from commit message
  Call: Jira API to update ticket
  
Why custom?
  ✅ Your team uses Jira
  ✅ Workflow automation
  ✅ Integration specific to you
```

**Implementation:**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"git commit\"",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/update-jira.js"
    }
  ],
  "description": "Update Jira ticket status on commit"
}
```

---

### **Scenario 3: Security/Compliance Rules**

```yaml
Example: Block commits with secrets

Hook:
  Event: PreToolUse (git commit)
  Check: Scan for API keys, passwords, tokens
  Action: Block if found
  
Why custom?
  ✅ Company security policy
  ✅ Prevent data leaks
  ✅ Compliance requirement
```

**Implementation:**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"git commit\"",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/scan-secrets.js"
    }
  ],
  "description": "Block commits containing secrets"
}
```

---

### **Scenario 4: Performance Monitoring**

```yaml
Example: Track AI response time

Hook:
  Event: Stop (after each response)
  Action: Log response time
  Save: Performance metrics
  
Why custom?
  ✅ Monitor AI efficiency
  ✅ Identify slow operations
  ✅ Optimize workflow
```

**Implementation:**

```json
{
  "matcher": "*",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/track-performance.js"
    }
  ],
  "description": "Track and log AI response time"
}
```

---

### **Scenario 5: Team Notifications**

```yaml
Example: Notify team on deployment

Hook:
  Event: PostToolUse (deploy command)
  Action: Send Slack notification
  Include: Who, what, when, status
  
Why custom?
  ✅ Team collaboration
  ✅ Transparency
  ✅ Your Slack workspace
```

**Implementation:**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"(npm run deploy|pnpm deploy)\"",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/notify-slack.js"
    }
  ],
  "description": "Notify team on deployment"
}
```

---

## 🎯 Khi Nào Nên Custom Scripts?

### **Scenario 1: Complex Validation Logic**

```yaml
Example: Validate API schema

Where: Hook calls script
Logic:
  - Load OpenAPI spec
  - Parse modified files
  - Validate against spec
  - Report mismatches

Why script?
  ✅ Too complex for inline hook
  ✅ Reusable across hooks/commands
  ✅ Testable independently
```

**Implementation:**

```javascript
// .claude/scripts/hooks/validate-api-schema.js
const fs = require('fs');
const yaml = require('yaml');

// Read OpenAPI spec
const spec = yaml.parse(fs.readFileSync('./api-spec.yaml', 'utf8'));

// Validate modified files
const files = getModifiedFiles();
for (const file of files) {
  if (file.includes('/api/')) {
    validateAgainstSpec(file, spec);
  }
}
```

---

### **Scenario 2: External Service Integration**

```yaml
Example: Update Notion database

Where: Command calls script
Logic:
  - Read BMAD artifacts
  - Parse PRD, stories, tasks
  - Update Notion pages
  - Sync status

Why script?
  ✅ API integration
  ✅ Complex data transformation
  ✅ Error handling
  ✅ Rate limiting
```

**Implementation:**

```python
# .claude/scripts/sync-notion.py
import notion_client
import yaml

def sync_prd_to_notion():
    # Read BMAD PRD
    with open('_bmad-output/planning-artifacts/prd.yaml') as f:
        prd = yaml.safe_load(f)
    
    # Update Notion
    notion = notion_client.Client(auth=os.environ["NOTION_TOKEN"])
    notion.pages.update(page_id=..., properties=...)
```

---

### **Scenario 3: Code Generation**

```yaml
Example: Generate CRUD endpoints

Where: Custom command calls script
Input: Model definition
Output: API routes + tests + docs

Why script?
  ✅ Template processing
  ✅ File generation
  ✅ Boilerplate reduction
```

**Implementation:**

```javascript
// .claude/scripts/generate-crud.js
const fs = require('fs');
const Handlebars = require('handlebars');

function generateCRUD(modelName) {
  const template = fs.readFileSync('./templates/crud.hbs', 'utf8');
  const compiled = Handlebars.compile(template);
  
  // Generate routes
  fs.writeFileSync(
    `./api/${modelName}.route.ts`,
    compiled({ model: modelName })
  );
  
  // Generate tests
  // Generate docs
}
```

---

## 📝 Custom Hook Template

```json
{
  "$schema": "https://json.schemastore.org/claude-code-settings.json",
  "hooks": {
    "PreToolUse": [
      {
        "matcher": "YOUR_MATCHER_HERE",
        "hooks": [
          {
            "type": "command",
            "command": "YOUR_COMMAND_HERE"
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

### **Matcher Examples:**

```javascript
// Match specific tool
"tool == \"Bash\""
"tool == \"Write\""
"tool == \"Edit\""

// Match command pattern
"tool_input.command matches \"npm run\""
"tool_input.command matches \"git (commit|push)\""
"tool_input.command matches \"deploy\""

// Match file pattern
"tool_input.file_path matches \"\\.ts$\""
"tool_input.file_path matches \"src/api/\""
"tool_input.file_path matches \"\\.test\\.(ts|js)$\""

// Combine conditions
"tool == \"Bash\" && tool_input.command matches \"test\""
"tool == \"Write\" && tool_input.file_path matches \"\\.tsx$\""

// Match everything
"*"
```

---

## 📝 Custom Script Template

```javascript
#!/usr/bin/env node
/**
 * Custom Hook Script: [Name]
 * 
 * Purpose: [What this script does]
 * Trigger: [Which hook calls this]
 * Input: [stdin JSON if any]
 * Output: [what it returns]
 */

const fs = require('fs');

// Read stdin if hook passes data
let data = '';
process.stdin.on('data', chunk => {
  data += chunk;
});

process.stdin.on('end', () => {
  try {
    // Parse hook data if provided
    const hookData = data ? JSON.parse(data) : null;
    
    // YOUR LOGIC HERE
    const result = yourCustomLogic(hookData);
    
    // For blocking hooks: exit 1 to block
    if (result.shouldBlock) {
      console.error('[Hook] BLOCKED: ' + result.reason);
      process.exit(1);
    }
    
    // For info hooks: just log
    if (result.info) {
      console.error('[Hook] ' + result.info);
    }
    
    // Pass through original data (for PostToolUse)
    if (data) {
      console.log(data);
    }
    
    process.exit(0);
  } catch (error) {
    console.error('[Hook] Error:', error.message);
    process.exit(1);
  }
});

function yourCustomLogic(hookData) {
  // Implement your logic
  return {
    shouldBlock: false,
    info: 'Processing complete',
    reason: null
  };
}
```

---

## 🚀 Real Custom Examples

### **Example 1: Enforce Test Coverage**

```json
// .claude/hooks/hooks.json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"git commit\"",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/check-coverage.js"
    }
  ],
  "description": "Block commits if test coverage < 80%"
}
```

```javascript
// .claude/scripts/hooks/check-coverage.js
const { execSync } = require('child_process');

const coverage = execSync('npm run test:coverage -- --silent', {
  encoding: 'utf8'
});

const match = coverage.match(/All files\s+\|\s+(\d+)/);
const percent = match ? parseInt(match[1]) : 0;

if (percent < 80) {
  console.error(`[Hook] BLOCKED: Test coverage is ${percent}%, need 80%+`);
  console.error('[Hook] Run: npm run test:coverage');
  process.exit(1);
}

console.error(`[Hook] ✅ Test coverage: ${percent}%`);
```

---

### **Example 2: Auto-Format Before Commit**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"git commit\"",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/auto-format.js"
    }
  ],
  "description": "Auto-format modified files before commit"
}
```

```javascript
// .claude/scripts/hooks/auto-format.js
const { execSync } = require('child_process');

const files = execSync('git diff --cached --name-only --diff-filter=ACMR', {
  encoding: 'utf8'
}).split('\n').filter(Boolean);

for (const file of files) {
  if (/\.(ts|tsx|js|jsx)$/.test(file)) {
    console.error(`[Hook] Formatting: ${file}`);
    execSync(`prettier --write ${file}`);
    execSync(`git add ${file}`);
  }
}

console.error('[Hook] ✅ All files formatted');
```

---

### **Example 3: Sync with External System**

```json
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"git push\"",
  "hooks": [
    {
      "type": "command",
      "command": "python3 .claude/scripts/hooks/sync-external.py",
      "async": true
    }
  ],
  "description": "Sync changes to external system (async)"
}
```

```python
# .claude/scripts/hooks/sync-external.py
import requests
import os

# Get current branch and commit
branch = os.popen('git branch --show-current').read().strip()
commit = os.popen('git rev-parse HEAD').read().strip()

# Notify external system
response = requests.post(
    'https://api.example.com/webhook',
    json={'branch': branch, 'commit': commit}
)

print(f'[Hook] ✅ Synced to external system: {response.status_code}')
```

---

## ⚠️ Best Practices

### **DO:**

```yaml
✅ Test hooks thoroughly before enabling
✅ Provide clear error messages
✅ Make hooks fast (< 1 second if possible)
✅ Use async for long-running tasks
✅ Document what each hook does
✅ Version control your custom hooks
✅ Handle errors gracefully
✅ Provide way to bypass if needed
```

### **DON'T:**

```yaml
❌ Block critical operations without escape hatch
❌ Make hooks that take too long
❌ Hardcode sensitive data in hooks
❌ Modify ECC source files directly
❌ Create hooks for everything (keep it minimal)
❌ Forget to test edge cases
❌ Leave debugging console.log in production hooks
```

---

## 📊 Decision Matrix

```
┌─────────────────────┬──────────────┬─────────────┐
│      NEED           │  USE THIS    │   WHY       │
├─────────────────────┼──────────────┼─────────────┤
│ Universal rules     │ Keep ECC     │ Tested,     │
│ (tmux, console.log) │ defaults     │ maintained  │
├─────────────────────┼──────────────┼─────────────┤
│ Project-specific    │ Custom hook  │ Your rules  │
│ conventions         │              │             │
├─────────────────────┼──────────────┼─────────────┤
│ External            │ Custom       │ Integration │
│ integrations        │ hook+script  │ needs       │
├─────────────────────┼──────────────┼─────────────┤
│ Complex validation  │ Custom       │ Reusable    │
│                     │ script       │ testable    │
├─────────────────────┼──────────────┼─────────────┤
│ Simple checks       │ Inline hook  │ No overhead │
│ (one-liner)         │ command      │             │
└─────────────────────┴──────────────┴─────────────┘
```

---

## ✅ Summary

### **Nguồn Gốc:**

```yaml
ECC:
  ✅ Provides hooks.json template
  ✅ Provides common hook scripts
  ✅ Provides instinct-cli.py
  
BMAD:
  ❌ No hooks
  ❌ No scripts
  
UI/UX Pro Max:
  ❌ No hooks
  ⚠️ Has helper scripts (for skill only)
```

### **Custom:**

```yaml
Can Do:
  ✅ Add new hooks
  ✅ Modify existing hooks
  ✅ Create custom scripts
  ✅ Remove unwanted hooks
  
Should Do:
  ✅ When you have project-specific rules
  ✅ When you need external integrations
  ✅ When you have compliance requirements
  
Should NOT Do:
  ❌ For everything (keep minimal)
  ❌ Modify ECC source directly
  ❌ Create slow hooks (< 1s ideal)
```

---

**Last updated: January 28, 2026**
