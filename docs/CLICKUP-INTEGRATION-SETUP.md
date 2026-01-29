# ClickUp-GitHub Integration Setup

## For Team Members (Developers)

**Không cần setup gì thêm!** Chỉ cần:

```bash
# Clone repo
git clone https://github.com/hutune/demo-structure.git
cd demo-structure

# Daily workflow
git pull origin main
# ... edit stories in _bmad-output/stories/ ...
git add .
git commit -m "Update story-1.1"
git push origin main
# → ClickUp auto-updated!
```

---

## For Admin (One-Time Setup)

### 1. GitHub Repository Secret

**Required**: Add ClickUp API key

1. Go to: Repository → Settings → Secrets → Actions
2. New repository secret:
   - **Name**: `CLICKUP_API_KEY`
   - **Value**: `pk_xxx` (your ClickUp API token)

### 2. ClickUp List IDs

Current configuration (in `.github/workflows/sync-clickup.yml`):

```yaml
CLICKUP_EPICS_LIST_ID: "901815396322"    # Epics / Initiatives
CLICKUP_STORIES_LIST_ID: "901815396340"   # Product Backlog
```

To change, update these values in the workflow file.

---

## Troubleshooting

### Error: "refusing to allow OAuth App to create workflow"

**Cause**: Your git credentials don't have `workflow` scope.

**Fix**: Use Personal Access Token (PAT) with workflow scope:
1. GitHub → Settings → Developer settings → Personal access tokens
2. Generate new token with `repo` + `workflow` scopes
3. Use token as password when pushing

### Error: "exit code 128" in GitHub Action

**Cause**: Workflow lacks permission to commit.

**Fix**: Already fixed in workflow with:
```yaml
permissions:
  contents: write
```

---

## File Structure

```
_bmad-output/
├── epics/
│   └── epic-001-authentication.md    # → ClickUp Epic
└── stories/
    ├── story-1.1-login-page.md       # → ClickUp Task
    └── story-1.2-user-registration.md
```

## Status Mapping

| Markdown Status | Epics List | Product Backlog |
|-----------------|------------|-----------------|
| `to-do` | "to do" | "Open" |
| `in-progress` | "in progress" | "scoping" |
| `done` | "complete" | - |
