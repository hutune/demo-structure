---
id: "STORY-2.2"
epic_id: "EPIC-02"
title: "Account Tiers & Limits"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["tiers", "limits", "subscription"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Account Tiers & Limits

## User Story

**As an** advertiser,
**I want** to understand my account tier and its limits,
**So that** I can plan my campaigns accordingly.

## Acceptance Criteria

### Tier Definitions
- [ ] FREE tier: $0/month, 2 campaigns, $500 max budget, $100 daily limit
- [ ] BASIC tier: $99/month, 5 campaigns, $2,000 max budget, $500 daily limit
- [ ] PREMIUM tier: $499/month, 20 campaigns, $10,000 max budget, $2,000 daily limit
- [ ] ENTERPRISE tier: Custom pricing, unlimited features

### Hard Limits (Cannot Exceed)
- [ ] Number of active campaigns
- [ ] Budget per campaign
- [ ] Team member count

### Soft Limits (Warning + Block)
- [ ] Daily spending: warning at 80%, blocked at 100%
- [ ] Monthly spending: warning at 80%, blocked at 100%

### Limit Enforcement
- [ ] At 80% of limit: send warning notification
- [ ] At 100% of limit: block action with upgrade suggestion
- [ ] Enterprise accounts: allow up to 10% temporary overage

### Feature Access by Tier
| Feature | FREE | BASIC | PREMIUM | ENTERPRISE |
|---------|------|-------|---------|------------|
| Content Uploads | 10 | 50 | 200 | Unlimited |
| Team Members | 1 | 3 | 10 | Unlimited |
| Support | Community | Email (48h) | Priority (24h) | Dedicated |
| API Access | No | No | Yes | Yes |
| Advanced Reports | No | Yes | Yes | Yes |

## Business Flow

### Tier Progression Flow

```
    FREE Tier
        │
        │  Meet requirements?
        │  ├── Payment method added
        │  └── Want more features
        v
    BASIC Tier ($99/month)
        │
        │  Growing needs?
        │  ├── More campaigns needed
        │  ├── Higher budgets required
        │  └── API access needed
        v
    PREMIUM Tier ($499/month)
        │
        │  Enterprise needs?
        │  ├── Unlimited requirements
        │  ├── Custom limits
        │  └── Dedicated support
        v
    ENTERPRISE Tier (Custom)
```

### Limit Enforcement Flow

```
┌─────────────────────────────────────────────────────────────────────┐
│                    LIMIT ENFORCEMENT                                │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   At 80% of limit:                                                  │
│   └── You receive a warning notification                            │
│       "You've used 80% of your daily spending limit"                │
│                                                                     │
│   At 100% of limit:                                                 │
│   └── Action is blocked                                             │
│       "Daily spending limit reached. Upgrade for higher limits."    │
│                                                                     │
│   Enterprise accounts only:                                         │
│   └── Can exceed by up to 10% (temporary overage allowed)           │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Spending Limits Reference

| Tier | Daily Limit | Monthly Limit | Per Campaign |
|------|-------------|---------------|---------------|
| FREE | $100 | $1,000 | $500 |
| BASIC | $500 | $5,000 | $2,000 |
| PREMIUM | $2,000 | $50,000 | $10,000 |
| ENTERPRISE | Custom | Custom | Custom |

## Checklist (Subtasks)

- [ ] Create tier configuration table
- [ ] Implement limit checking service
- [ ] Create 80% warning notification
- [ ] Implement 100% blocking logic
- [ ] Add Enterprise 10% overage allowance
- [ ] Create tier comparison UI
- [ ] Add current usage display in dashboard
- [ ] Implement feature flags by tier
- [ ] Create upgrade prompts when limits reached
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
