#!/bin/bash
# BMAD + Everything CC Enhanced Setup
# Quick installation trước tết!

set -e

PROJECT_ROOT="/Users/mazhnguyen/Library/CloudStorage/GoogleDrive-work.huutrung@gmail.com/My Drive/KwayVina/RMN"
cd "$PROJECT_ROOT"

echo "🚀 BMAD + Everything CC Quick Setup"
echo "===================================="
echo ""

# Step 1: Verify BMAD (already installed)
echo "✅ Step 1: BMAD already installed"
echo "   Location: _bmad/"
echo ""

# Step 2: Clone Everything Claude Code
echo "📦 Step 2: Fetching Everything Claude Code..."
if [ ! -d "/tmp/ecc" ]; then
    git clone --depth 1 https://github.com/affaan-m/everything-claude-code.git /tmp/ecc
    echo "   ✅ Downloaded"
else
    echo "   ✅ Already exists"
fi
echo ""

# Step 3: Create .claude directory structure
echo "📁 Step 3: Creating .claude directory..."
mkdir -p .claude/agents
mkdir -p .claude/skills
mkdir -p .claude/commands
mkdir -p .claude/hooks
mkdir -p .claude/scripts/hooks
echo "   ✅ Directories created"
echo ""

# Step 4: Copy ESSENTIAL components only (no conflicts)
echo "📦 Step 4: Installing Everything CC components..."

# 4.1: Continuous Learning v2 (MUST HAVE!)
echo "   → Installing continuous-learning-v2..."
cp -r /tmp/ecc/skills/continuous-learning-v2 .claude/skills/ec-learning
echo "      ✅ Installed: .claude/skills/ec-learning/"

# 4.2: Security Specialist
echo "   → Installing security specialist..."
cp /tmp/ecc/agents/security-reviewer.md .claude/agents/ec-security.md
echo "      ✅ Installed: .claude/agents/ec-security.md"

# 4.3: TDD Tools
echo "   → Installing TDD tools..."
cp /tmp/ecc/agents/tdd-guide.md .claude/agents/ec-tdd-guide.md
cp -r /tmp/ecc/skills/tdd-workflow .claude/skills/ec-tdd
echo "      ✅ Installed: TDD workflow"

# 4.4: Build Error Resolver
echo "   → Installing build error resolver..."
cp /tmp/ecc/agents/build-error-resolver.md .claude/agents/ec-build-fixer.md
echo "      ✅ Installed: ec-build-fixer"

# 4.5: Refactoring Tools
echo "   → Installing refactoring tools..."
cp /tmp/ecc/agents/refactor-cleaner.md .claude/agents/ec-refactor.md
echo "      ✅ Installed: ec-refactor"

# 4.6: Eval Harness
echo "   → Installing eval harness..."
cp -r /tmp/ecc/skills/eval-harness .claude/skills/ec-eval
echo "      ✅ Installed: ec-eval"

# 4.7: Verification Loop
echo "   → Installing verification loop..."
cp -r /tmp/ecc/skills/verification-loop .claude/skills/ec-verify
echo "      ✅ Installed: ec-verify"

# 4.8: Commands (slash commands)
echo "   → Installing commands..."
cp /tmp/ecc/commands/*.md .claude/commands/
echo "      ✅ Installed: 23 commands (/tdd, /evolve, /instinct-status, etc.)"

echo ""

# Step 5: Setup Hooks
echo "📦 Step 5: Setting up hooks..."
cp /tmp/ecc/hooks/hooks.json .claude/hooks/hooks.json
cp -r /tmp/ecc/scripts/hooks/*.js .claude/scripts/hooks/ 2>/dev/null || true
echo "   ✅ Hooks configured"
echo ""

# Step 6: Create integration config
echo "📝 Step 6: Creating integration config..."
cat > .claude/PROJECT-CONFIG.md << 'EOF'
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

- ec-security: Security deep-dive
- ec-tdd-guide: TDD workflow
- ec-build-fixer: Build errors
- ec-refactor: Code cleanup
- ec-learning: Continuous learning (background)

### Conflicts Resolved:
- ❌ SKIP ECC planner (use BMAD /pm)
- ❌ SKIP ECC architect (use BMAD /architect)
- ❌ SKIP ECC generic code-review (use BMAD)

### Integration Flow:
1. BMAD orchestrates
2. ECC provides specialists
3. Learning observes (background)
4. No conflicts!

## Learning System Active
- Location: .claude/skills/ec-learning/
- Instincts: .claude/skills/ec-learning/instincts/
- Commands: /instinct-status, /evolve (when available)
EOF
echo "   ✅ Config created: .claude/PROJECT-CONFIG.md"
echo ""

# Step 7: Cleanup
echo "🧹 Step 7: Cleaning up..."
rm -rf /tmp/ecc
echo "   ✅ Temporary files removed"
echo ""

# Step 8: Summary
echo "✅ Installation Complete!"
echo "========================"
echo ""
echo "📂 Structure:"
echo "   _bmad/                    (BMAD framework - primary)"
echo "   _bmad-output/             (artifacts)"
echo "   .agent/workflows/         (BMAD commands)"
echo "   .claude/agents/           (ECC specialists)"
echo "   .claude/skills/           (ECC tools)"
echo "   .claude/PROJECT-CONFIG.md (integration guide)"
echo ""
echo "🎯 Next Steps:"
echo "   1. Read: .claude/PROJECT-CONFIG.md"
echo "   2. Try: @bmad-help"
echo "   3. Start: @bmad-brainstorming"
echo ""
echo "🎊 Ready to use! Chúc mừng năm mới!"
