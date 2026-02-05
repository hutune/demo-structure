---
id: "STORY-3.6"
epic_id: "EPIC-03"
title: "Store Performance Metrics"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["performance", "metrics", "analytics"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Store Performance Metrics

## User Story

**As a** supplier,
**I want** to view performance metrics for each store,
**So that** I can optimize my network and maximize revenue.

## Acceptance Criteria

### Store-Level Metrics
- [ ] Total impressions (daily, weekly, monthly)
- [ ] Revenue generated
- [ ] Device uptime percentage
- [ ] Content fill rate (% of time showing paid content)
- [ ] Average dwell time (if measurable)

### Network-Level Metrics
- [ ] Total impressions across all stores
- [ ] Total revenue across all stores
- [ ] Average uptime across network
- [ ] Best and worst performing stores
- [ ] Trend analysis (week over week, month over month)

### Device Health by Store
- [ ] Number of active devices
- [ ] Number of offline devices
- [ ] Average device health score
- [ ] Maintenance alerts

### Reporting
- [ ] Daily summary email (optional)
- [ ] Weekly performance report
- [ ] Monthly revenue statement
- [ ] Export to CSV/PDF

### Performance Alerts
- [ ] Store falls below 80% uptime
- [ ] Revenue drops >20% week over week
- [ ] Device goes offline for >1 hour
- [ ] Unusual impression patterns (potential fraud)

## Business Flow

### Quality Score Components

| Component | Weight | Description |
|-----------|--------|-------------|
| Device Uptime | 35% | Average uptime across all devices |
| Revenue Performance | 25% | Actual vs. expected revenue |
| Compliance | 20% | Policy violations (fewer = better) |
| Customer Ratings | 10% | Advertiser ratings of supplier quality |
| Growth | 10% | Month-over-month revenue growth |

### Device Uptime Scoring

| Uptime | Score |
|--------|-------|
| 98%+ | 100 |
| 95-97% | 90 |
| 90-94% | 75 |
| 85-89% | 60 |
| Below 85% | 40 |

### Revenue Performance Scoring

| Performance | Score |
|-------------|-------|
| At or above expected | 100 |
| 80-99% of expected | 80 |
| 60-79% of expected | 60 |
| Below 60% of expected | 40 |

### Compliance Scoring

| Violations (last 90 days) | Score |
|---------------------------|-------|
| 0 | 100 |
| 1 | 90 |
| 2 | 80 |
| 3+ | 70 or less |

### Performance Tiers

| Tier | Score Range | Benefits |
|------|-------------|----------|
| Platinum | 90-100 | Priority support, +5% revenue share, early feature access |
| Gold | 80-89 | Priority support, standard revenue share |
| Silver | 70-79 | Standard support, standard revenue share |
| Bronze | 60-69 | Standard support, improvement plan required |
| Poor | Below 60 | Account under review, possible suspension |

### Poor Performance Escalation Flow

```
Score < 60 for 2 consecutive months
         |
         v
Send warning email
         |
         v
Require improvement plan (7 days)
         |
         v
Assign support to help improve
         |
No improvement after 30 days
         |
         v
Suspend new device registrations
Reduce priority in campaign matching
         |
No improvement after 60 days
         |
         v
Suspend supplier account
Hold revenue pending resolution
```

### Revenue Calculation

```
Per Impression:
  Impression cost = Campaign CPM rate ÷ 1,000
  Supplier revenue = Impression cost × 80%
  Platform revenue = Impression cost × 20%

Example:
  Campaign CPM: $5.00
  One impression cost: $5.00 ÷ 1,000 = $0.005
  Supplier earns: $0.005 × 80% = $0.004
  Platform earns: $0.005 × 20% = $0.001
```

### Revenue Estimation

```
Daily Revenue per Device:
  (Hourly impressions × Operating hours × Average CPM × 80%) ÷ 1,000

Example:
  50 impressions/hour × 12 hours × $4.00 × 80% ÷ 1,000 = $1.92/day

Monthly Revenue Projection:
  Devices × Daily revenue × 30 days

Example:
  10 devices × $1.92/day × 30 = $576/month
```

## Checklist (Subtasks)

- [ ] Create store-level metrics aggregation
- [ ] Build network-level metrics dashboard
- [ ] Implement impression tracking per store
- [ ] Add revenue calculation per store
- [ ] Create uptime calculation service
- [ ] Build fill rate calculation
- [ ] Create performance alerts
- [ ] Implement email reports (daily, weekly, monthly)
- [ ] Add CSV/PDF export
- [ ] Create trend analysis charts
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
