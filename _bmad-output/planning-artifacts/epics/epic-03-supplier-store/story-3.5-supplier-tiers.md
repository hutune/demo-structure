---
id: "STORY-3.5"
epic_id: "EPIC-03"
title: "Supplier Tiers & Benefits"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["tiers", "benefits", "revenue-share"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Supplier Tiers & Benefits

## User Story

**As a** supplier,
**I want** to understand my tier level and benefits,
**So that** I can maximize my revenue and grow my network.

## Acceptance Criteria

### Tier Definitions
| Tier | Active Devices | Revenue Share | Min Monthly Impressions |
|------|----------------|---------------|-------------------------|
| BRONZE | 1-10 | 40% | 0 |
| SILVER | 11-50 | 50% | 100,000 |
| GOLD | 51-200 | 60% | 500,000 |
| PLATINUM | 201+ | 70% | 2,000,000 |

### Tier Evaluation
- [ ] Tiers evaluated monthly on 1st
- [ ] Based on: device count + monthly impressions
- [ ] Must meet BOTH criteria to qualify
- [ ] Downgrade if criteria not met for 2 consecutive months

### Tier Benefits
- [ ] BRONZE: Basic reporting, email support
- [ ] SILVER: Enhanced reporting, priority email (24h)
- [ ] GOLD: Advanced analytics, phone support, dedicated account manager
- [ ] PLATINUM: Custom reporting, 24/7 support, strategic partnership benefits

### Tier Notifications
- [ ] Monthly tier status report
- [ ] Alert when close to next tier
- [ ] Warning when at risk of downgrade
- [ ] Celebration notification on upgrade

### Dashboard Display
- [ ] Current tier with badge
- [ ] Progress toward next tier
- [ ] Current month metrics
- [ ] Revenue share percentage

## Business Flow

### Tier Structure

| Feature | Starter | Professional | Enterprise |
|---------|---------|--------------|------------|
| Monthly Fee | Free | $99 | Custom (from $499) |
| Max Devices | 5 | 50 | Unlimited |
| Max Stores | 3 | 20 | Unlimited |
| Revenue Share | 80% | 80% (85% if Platinum) | 85% base (90% if Platinum) |
| Payment Schedule | Monthly | Weekly/Biweekly/Monthly | Weekly/Daily (custom) |
| Minimum Payout | $100 | $50 | $25 |
| Blocking Rules | 5 | 20 | Unlimited |
| Support SLA | Email (48h) | Email + Chat (24h) | Dedicated manager (4h) |
| Analytics History | 30 days | 1 year | Unlimited |

### Tier Limit Enforcement

```
Device Limit Reached:
      |
      v
Registration rejected
      |
      v
Prompt: "You've reached your device limit. Upgrade to add more devices."
      |
      v
Show upgrade options

---

Approaching Limit (at 80%):
      |
      v
Notification sent
      |
      v
ROI calculation shown
(potential revenue with more devices vs. upgrade cost)
```

### Tier Upgrade Flow

```
Step 1: Charge prorated amount for current billing period
      |
      v
Step 2: Update tier immediately
      |
      v
Step 3: New limits take effect
      |
      v
Step 4: Confirmation email sent
```

### Tier Downgrade Flow

```
Step 1: Scheduled for end of current billing period
      |
      v
Step 2: Check - Exceeds new tier limits?
      |
     YES --> Must deactivate devices/stores first
      |        Cannot downgrade while over limit
      |
     NO
      |
      v
Step 3: Prorated refund if applicable
```

### Enterprise Onboarding Flow

```
Step 1: Submit inquiry form
      |
      v
Step 2: Sales call within 2 business days
      |
      v
Step 3: Custom pricing based on scale and requirements
      |
      v
Step 4: Contract signed
      |
      v
Step 5: Manual account upgrade by admin
```

### Platinum Performance Bonus

| Condition | Benefit |
|-----------|----------|
| Score ≥ 90 for 3 consecutive months | Revenue share increases to 85% (or 90% for Enterprise) |
| Score drops below 90 | 30-day grace period |
| After grace period | Reverts to standard share |

## Checklist (Subtasks)

- [ ] Create tier configuration
- [ ] Implement monthly tier evaluation job
- [ ] Add tier upgrade logic
- [ ] Implement 2-month downgrade rule
- [ ] Build tier dashboard widget
- [ ] Create progress indicator toward next tier
- [ ] Add tier change notifications
- [ ] Implement revenue share calculation
- [ ] Create tier benefits display
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
