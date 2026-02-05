---
id: "STORY-2.6"
epic_id: "EPIC-02"
title: "Tier Upgrade & Downgrade"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["tier", "upgrade", "subscription"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Tier Upgrade & Downgrade

## User Story

**As an** advertiser,
**I want** to upgrade or downgrade my account tier,
**So that** I can adjust my plan based on my needs.

## Acceptance Criteria

### Upgrade Process
- [ ] Go to Account Settings → Upgrade
- [ ] Compare current tier with new tier
- [ ] Choose monthly or annual billing (annual saves 10%)
- [ ] Confirm payment
- [ ] Upgrade takes effect immediately
- [ ] Mid-cycle upgrades are prorated

### Downgrade Process
- [ ] Submit request through support ticket
- [ ] Effective at end of current billing period
- [ ] Must be within new tier's limits before downgrade
- [ ] System checks: campaigns, content, team members

### Billing Rules
- [ ] Annual billing gets 10% discount
- [ ] Upgrade: prorated charge for remaining days
- [ ] Downgrade: no refund, effective next cycle
- [ ] Enterprise tier requires sales contact

### Pre-Downgrade Checks
- [ ] Active campaigns within new limit
- [ ] Content uploads within new limit
- [ ] Team members within new limit
- [ ] Block downgrade if limits exceeded

## Business Flow

### Tier Upgrade Flow

```
    Go to Account Settings → Upgrade
           │
           v
    ┌──────────────────────────┐
    │ Compare Tiers            │
    │ (Current vs New)         │
    └────────────┬─────────────┘
                 │
                 v
    ┌──────────────────────────┐
    │ Select New Tier          │
    │ (Higher than current)    │
    └────────────┬─────────────┘
                 │
         ┌───────┴───────┐
         v               v
    Enterprise?      Other Tiers
         │               │
         v               v
    Contact Sales   Choose Billing:
    Team            ├── Monthly
                    └── Annual (10% off)
                         │
                         v
                    Confirm Payment
                    (Prorated charge)
                         │
                         v
                 ┌───────────────────┐
                 │ UPGRADE IMMEDIATE │
                 │ New limits active │
                 └───────────────────┘
```

### Tier Downgrade Flow

```
    Submit Downgrade Request
    (via Support Ticket)
           │
           v
    ┌──────────────────────────┐
    │ System Checks Usage      │
    ├──────────────────────────┤
    │ ✓ Active campaigns       │
    │ ✓ Content uploads        │
    │ ✓ Team members           │
    └────────────┬─────────────┘
                 │
          ┌──────┴──────┐
          v             v
    Within New      Exceeds New
    Tier Limits     Tier Limits
          │             │
          v             v
    Request          BLOCKED
    Approved         (Reduce usage
          │          first)
          v
    Scheduled for End
    of Billing Cycle
          │
          v
    ┌───────────────────────┐
    │ DOWNGRADE EFFECTIVE   │
    │ at next billing date  │
    └───────────────────────┘
```

### Billing Rules Reference

| Action | When It Happens | Billing Impact |
|--------|-----------------|----------------|
| Upgrade | Immediate | Prorated charge for remaining days |
| Downgrade | End of cycle | No refund, effective next period |
| Annual billing | At purchase | 10% discount applied |
| Enterprise | After sales contact | Custom pricing |

### Pre-Downgrade Limit Checks

| Resource | FREE | BASIC | PREMIUM |
|----------|------|-------|----------|
| Active Campaigns | ≤ 2 | ≤ 5 | ≤ 20 |
| Content Uploads | ≤ 10 | ≤ 50 | ≤ 200 |
| Team Members | ≤ 1 | ≤ 3 | ≤ 10 |

## Checklist (Subtasks)

- [ ] Create tier comparison page
- [ ] Implement upgrade selection flow
- [ ] Add proration calculation
- [ ] Integrate annual billing discount
- [ ] Create payment processing for upgrade
- [ ] Implement immediate tier change
- [ ] Build downgrade request flow
- [ ] Add usage validation for downgrade
- [ ] Create scheduled downgrade job
- [ ] Add Enterprise contact form
- [ ] Send confirmation emails
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
