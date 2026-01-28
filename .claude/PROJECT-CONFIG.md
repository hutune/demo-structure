# BMAD + Everything CC Integration Config

## Setup Date
January 28, 2026

## Primary: BMAD Method
All project workflows use BMAD as primary orchestrator.

## Added: Everything CC Tools
Selected components for quality & learning.

## Command Priority

### Use BMAD (Primary):
- @bmad-bmm-analyst or /analyst
- @bmad-bmm-pm or /pm  
- @bmad-bmm-architect or /architect
- @bmad-bmm-dev or /dev
- @bmad-bmm-create-prd
- @bmad-bmm-sprint-planning
- @bmad-bmm-dev-story
- @bmad-bmm-code-review

### Use Everything CC (Specialists):
Files in .claude/agents/ and .claude/skills/

**Agents:**
- ec-security: Security deep-dive
- ec-tdd-guide: TDD workflow
- ec-build-fixer: Build errors
- ec-refactor: Code cleanup
- ec-learning: Continuous learning (background)

**Commands (20):**
- /instinct-status - Check learned patterns
- /evolve - Generate workflows from patterns
- /tdd - TDD workflow
- /e2e - E2E testing
- /build-fix - Fix build errors
- /refactor-clean - Refactor code
- /verify - Verification loop
- /eval - Eval harness
- /learn - Manual learning trigger
- ... (11 more specialist commands)

### Conflicts Removed:
- ❌ REMOVED /plan (use @bmad-bmm-pm)
- ❌ REMOVED /code-review (use @bmad-bmm-code-review)
- ❌ REMOVED /orchestrate (use @bmad-party-mode)
- ❌ SKIP ECC planner agent (use BMAD /pm)
- ❌ SKIP ECC architect agent (use BMAD /architect)

### Integration Flow:
1. BMAD orchestrates
2. ECC provides specialists
3. Learning observes (background)
4. No conflicts!

## Learning System Active
- Location: .claude/skills/ec-learning/
- Instincts: .claude/skills/ec-learning/instincts/
- Commands: /instinct-status, /evolve (when available)

## Adding More Plugins

All Claude Code plugins install to `.claude/`:

### Structure:
```
.claude/
├── agents/          ← All agents from all plugins
├── skills/          ← All skills from all plugins
├── commands/        ← All commands from all plugins
└── hooks/           ← All hooks from all plugins
```

### Example: UI/UX Pro Max
```bash
# Clone
git clone https://github.com/nextlevelbuilder/ui-ux-pro-max-skill.git

# Copy to .claude/skills/
cp -r ui-ux-pro-max-skill/.claude/skills/ui-ux-pro-max .claude/skills/

# Done! Auto-activates when you mention:
# - "design", "UI", "dashboard", "glassmorphism"
# - "color palette", "typography", "button"
# - etc. (see skill description)
```

### How Skills Work:
- **No commands needed!** Skills auto-activate based on keywords
- BMAD can delegate to skills automatically
- ECC can use skills automatically  
- You can mention skill explicitly: "Use ui-ux-pro-max skill"

### Integration:
```
@bmad-bmm-ux-designer
  → May use ui-ux-pro-max skill (if installed)
  → May use other design skills (if installed)
  → All automatic!
```

## Planned Integrations

### Tomorrow (January 29, 2026):

**ClickUp Automation:**
```json
// .claude/hooks/hooks.json - Custom hook to add
{
  "matcher": "tool == \"Bash\" && tool_input.command matches \"git commit\"",
  "hooks": [
    {
      "type": "command",
      "command": "node .claude/scripts/hooks/sync-clickup.js"
    }
  ],
  "description": "Sync BMAD artifacts with ClickUp tasks"
}
```

**Team Documentation:**
- Create: `docs/team-best-practices.md`
- Content: BMAD workflows, command reference, conventions
- Audience: Team members

### Future Plugins (Install as needed):

**UI/UX Pro Max:**
```bash
# When needed for design work
git clone https://github.com/nextlevelbuilder/ui-ux-pro-max-skill.git
cp -r ui-ux-pro-max-skill/.claude/skills/ui-ux-pro-max .claude/skills/
# See: PLUGIN-INSTALL-GUIDE.md
```

**Other Potential Additions:**
- Security scanning tools
- Language-specific analyzers
- Custom team workflows
- Project-specific automations
