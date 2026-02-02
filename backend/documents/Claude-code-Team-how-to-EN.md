# Claude Code, Team How To

# Claude Code Team Collaboration Guide: Mastering Agent, Subagent, and Skill

Claude Code supports team collaboration and automation through a **3-tier hierarchy**—Main Agent, Subagent, and Skill. For a small 3-person team (backend/frontend/PM) to fully automate development of an RMN SaaS platform, strategic use of this system is essential. With a team subscription, sharing the `.claude/` directory structure and `CLAUDE.md` via Git lets you set up the same development environment immediately.

---

## Core Differences: Agent, Subagent, Skill

Claude Code’s agent system serves two goals: **task separation** and **context management**. Each component has a distinct role.

| Component | Purpose | Context | Invocation | Location |
|-----------|---------|---------|------------|----------|
| **Main Agent** | Core conversation handling | Shared chat | Always active | Built-in |
| **Subagent** | Delegated specialized work | **Isolated context** | Auto or explicit delegation | `.claude/agents/` |
| **Skill** | Knowledge/capability injection | Loaded into current context | Relevance-based auto-trigger | `.claude/skills/` |

**Subagents** run in a separate context window so they don’t clutter the main conversation. They’re suited for independent work like code review, exploration, and deep research. **Skills** inject knowledge into the current context and load only what’s needed via progressive disclosure.

### Three Built-in Subagents

Claude Code provides three built-in subagents:

- **Explore Subagent**: Uses Haiku, read-only; focused on file browsing and code search
- **Plan Subagent**: Uses Sonnet; for codebase analysis and planning
- **General-purpose Subagent**: Uses Sonnet; for complex research including edits

---

## Config Structure and Team Sharing

### Recommended Directory Structure

```
your-rmn-project/
├── CLAUDE.md              # Project memory (commit to Git)
├── .mcp.json              # MCP server config (commit to Git)
├── .claude/
│   ├── settings.json      # Team-shared settings (commit to Git)
│   ├── settings.local.json # Personal settings (.gitignore)
│   │
│   ├── agents/            # Custom Subagents
│   │   ├── code-reviewer.md
│   │   ├── golang-expert.md
│   │   └── flutter-specialist.md
│   │
│   ├── commands/          # Slash commands
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
└── agent_docs/            # Progressive disclosure docs
    ├── building_the_project.md
    └── architecture.md
```

### Team Sharing via Git

**Commit (team-shared):**
- `CLAUDE.md` — project context
- `.claude/settings.json` — permissions, env vars
- `.claude/commands/` — team slash commands
- `.claude/agents/` — shared subagents
- `.mcp.json` — MCP server config

**Add to .gitignore (personal):**
```gitignore
CLAUDE.local.md
.claude/settings.local.json
.claude/local/
```

When a teammate clones the repo and runs the `claude` command, **the same environment is applied automatically**.

---

## CLAUDE.md Template for RMN SaaS Project

Keeping **CLAUDE.md between 60–300 lines** is important so the LLM can reliably follow the instructions.

```markdown
# RMN SaaS Ad Management Platform

## Project overview
Retail Media Network–based SaaS ad management platform.
Flutter frontend, Golang backend, PostgreSQL, Kubernetes.

## Key directories
- `frontend/` — Flutter app (lib/features/, lib/shared/, lib/core/)
- `backend/` — Go services (cmd/, internal/, pkg/, api/)
- `infra/` — Kubernetes manifests, Terraform
- `docs/` — API docs, PRD

## Commands
```bash
# Backend
cd backend && go build ./... && go test ./...
golangci-lint run

# Frontend
cd frontend && flutter analyze && flutter test
flutter build apk --release
```

## Code conventions
### Golang
- Define behavior with interfaces, DI via constructor functions
- Always check errors, wrap with %w
- Repository pattern + Service layer

### Flutter
- Feature-First Clean Architecture (FFCA)
- Riverpod for state, flutter_hooks
- Dark-mode first, flex-based responsive UI

## Important
- Never commit .env files
- SQL queries must be parameterized
- Update OpenAPI spec when API changes

## Additional docs
For complex tasks, see `agent_docs/`:
- `architecture.md` — system design
- `api_conventions.md` — API rules
```

---

## Role-Based Agent/Skill Setup

### For Backend Developer (Golang)

**.claude/agents/golang-expert.md**
```markdown
---
name: golang-expert
description: Go code writing, review, test generation. Use for Go-related work.
tools: Read, Write, Edit, Bash, Grep, Glob
model: claude-sonnet-4-20250514
---

You are a 15-year Go backend expert.

## Responsibilities
- Follow Clean Architecture
- Write table-driven tests
- Apply error-handling patterns
- Optimize PostgreSQL queries

## Code style
- Always use context.Context
- Interface-first design
- Use testify assertions
```

**.claude/commands/go-test.md**
```markdown
---
name: go-test
description: Generate Go table-driven tests
allowed-tools: Read, Write, Bash
---

Generate table-driven tests for $ARGUMENTS:
1. Include happy path, error cases, edge cases
2. Use testify/assert
3. Mock only external dependencies
4. Run `go test -v` to verify
```

### For Frontend Developer (Flutter)

**.claude/agents/flutter-specialist.md**
```markdown
---
name: flutter-specialist
description: Flutter widget development, state management, UI. Use for Flutter work.
tools: Read, Write, Edit, Bash, Grep, Glob
model: claude-sonnet-4-20250514
---

You are a Flutter/Dart specialist.

## Architecture
- Feature-First Clean Architecture
- Riverpod + flutter_hooks
- GoRouter for navigation

## Widget principles
- Prefer small, composable widgets
- Use flex values instead of hardcoded sizes
- Use dart:developer log() instead of print()

## Testing
- Widget tests for UI components
- Unit tests for business logic
```

**.claude/commands/flutter-feature.md**
```markdown
---
name: flutter-feature
description: Scaffold a Flutter feature module
---

Create the '$ARGUMENTS' feature in Feature-First structure:

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

Add baseline files and Riverpod providers in each layer.
```

### For PM: Documentation Tools

**.claude/agents/prd-writer.md**
```markdown
---
name: prd-writer
description: PRD, feature specs, release notes. Use for documentation work.
tools: Read, Write, Grep, Glob
model: claude-sonnet-4-20250514
---

You are Sarah, a senior PM.

## PRD framework (15 sections)
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

Clarify requirements with enough questions before writing the PRD.
```

**.claude/commands/release-notes.md**
```markdown
---
name: release-notes
description: Auto-generate release notes
---

Generate release notes from the last $ARGUMENTS commits:

1. Run `git log --oneline -n $ARGUMENTS`
2. Group by type (feat, fix, docs, refactor)
3. Summarize changes from user perspective
4. Call out Breaking Changes
5. Update `docs/CHANGELOG.md`
```

---

## Example Team-Shared Settings

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

This setup auto-formats Go/Dart on save and blocks dangerous commands.

---

## Code Review and Test Automation Examples

### GitHub Actions Integration

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

### Code Review Automation Command

**.claude/commands/review.md**
```markdown
---
name: review
description: Perform PR code review
allowed-tools: Read, Grep, Glob, Bash(git diff:*)
---

Review changes on the current branch:

1. Run `git diff main...HEAD`
2. Check for security issues (SQL injection, XSS, CSRF)
3. Flag performance issues (N+1 queries, inefficient algorithms)
4. Verify coding conventions
5. Review test coverage
6. Report only issues with confidence ≥ 80%
```

---

## Team Subscription and Collaboration

### Claude Code Team Plan Features

Claude Code team subscription offers **Standard Seat ($25/month)** and **Premium Seat ($150/month)**. Main collaboration features:

- **Centralized management**: Seat assignment, usage monitoring, cost control
- **Connectors**: GitHub, Slack, Google Drive, Gmail
- **Usage limits**: ~**225 messages** per 5 hours, **50–95 hours** Sonnet 4 per week
- **Overage**: Additional usage via API within admin-set limits

### Workflow Patterns for a 3-Person Team

**Onboarding:**
1. Clone repo (CLAUDE.md, settings.json included)
2. Run `claude` in terminal
3. Use `/init` to confirm environment
4. Create `CLAUDE.local.md` if you need personal overrides

**Parallel work:**
```bash
# Run parallel Claude sessions with Git worktrees
git worktree add ../rmn-feature-auth feature/auth
git worktree add ../rmn-feature-ads feature/ads
# Run independent Claude sessions in each worktree
```

---

## Common Mistakes and How to Avoid Them

### Five Mistakes to Avoid

| Mistake | Problem | Fix |
|---------|---------|-----|
| Overwriting CLAUDE.md | Lower instruction compliance | Keep 60–300 lines, use progressive disclosure |
| Using /init without review | Can include wrong info | Always review and edit |
| Using Claude as linter | Higher cost, slower | Use hooks for gofmt, dart format |
| Including secrets | Security risk | Never put API keys or passwords in context |
| Poor context management | Unrelated info piles up | Use `/clear` between tasks |

### Best Practices Summary

- **CLAUDE.md is high leverage**: Write and iterate like a prompt
- **Protect with CODEOWNERS**: Require senior review for CLAUDE.md changes
- **Prefer deterministic tools**: Use hooks for linters/formatters; use Claude for creative work
- **Use progressive disclosure**: Put detailed docs in `agent_docs/` and reference from CLAUDE.md

---

## Quick Start Checklist

1. ✅ Add `CLAUDE.md` at project root (≤ 60 lines to start)
2. ✅ Set team permissions in `.claude/settings.json`
3. ✅ Add 3 role-based subagents under `.claude/agents/`
4. ✅ Add key workflow commands under `.claude/commands/`
5. ✅ Add `CLAUDE.local.md`, `.claude/settings.local.json` to `.gitignore`
6. ✅ Confirm teammates are ready after clone + `claude` run

With this structure, a 3-person team can effectively get **full development automation** for the RMN SaaS platform in a **shared Claude Code environment**.
