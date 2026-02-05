---
id: "STORY-4.5"
epic_id: "EPIC-04"
title: "Device Health Scoring"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["health", "scoring", "quality"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Device Health Scoring

## User Story

**As a** platform operator,
**I want** a health score for each device,
**So that** I can prioritize maintenance and identify problematic devices.

## Acceptance Criteria

### Health Score Components
| Factor | Weight | Criteria |
|--------|--------|----------|
| Connectivity | 30% | Uptime, missed heartbeats |
| Temperature | 20% | Within normal range |
| Storage | 15% | Available space |
| Content Freshness | 25% | Playlist up to date |
| Display Quality | 10% | No reported issues |

### Score Calculation
- [ ] Score 0-100 (higher = healthier)
- [ ] Calculated every hour
- [ ] Historical scores stored for trend

### Health Categories
- [ ] Excellent (90-100): No action needed
- [ ] Good (75-89): Monitor
- [ ] Fair (50-74): Schedule maintenance
- [ ] Poor (25-49): Urgent attention
- [ ] Critical (0-24): Immediate action

### Health Alerts
- [ ] Drop to Fair: Email notification
- [ ] Drop to Poor: Priority ticket created
- [ ] Drop to Critical: Immediate alert + phone call

### Health Dashboard
- [ ] Fleet health overview
- [ ] Distribution of health scores
- [ ] Trend over time
- [ ] Devices needing attention list

## Business Flow

### How Health Score Is Calculated

Every device receives a health score from 0 to 100. This helps identify problem devices that need attention.

```
┌─────────────────────────────────────────────────────────────────────┐
│                    HEALTH SCORE COMPONENTS                          │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   Uptime (40% of score)                                             │
│   ├── How often is the device online?                               │
│   ├── Target: 95% or higher                                         │
│   └── Excellent: 99% or higher                                      │
│                                                                     │
│   Performance (30% of score)                                        │
│   ├── Processor usage (should be below 80%)                         │
│   ├── Memory usage (should be below 80%)                            │
│   ├── Network speed (latency below 100ms)                           │
│   └── Content load time (below 5 seconds)                           │
│                                                                     │
│   Reliability (20% of score)                                        │
│   ├── Successful ad plays vs attempts                               │
│   ├── Successful content downloads                                  │
│   └── Few errors (less than 5 per day)                              │
│                                                                     │
│   Compliance (10% of score)                                         │
│   ├── Only plays approved content                                   │
│   ├── Follows operating hours                                       │
│   └── Security checks pass                                          │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Health Score Thresholds

| Score | Rating | What Happens |
|-------|--------|---------------|
| 90-100 | Excellent | Premium campaign eligible |
| 70-89 | Good | Normal operation |
| 50-69 | Fair | Supplier notified of issues |
| Below 50 | Poor | Admin review required |

### What Lowers Health Score

| Issue | Impact |
|-------|--------|
| Frequent offline periods | Major drop |
| High processor/memory usage | Moderate drop |
| Slow network connection | Moderate drop |
| Failed content downloads | Moderate drop |
| Many errors | Moderate drop |
| Playing unapproved content | Major drop |

### Health Score Flow

```
┌─────────────────────────────────────────────────────────────────────┐
│                    HEALTH SCORE MONITORING                          │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   [Hourly Health Check]                                             │
│          │                                                          │
│          ↓                                                          │
│   ┌─────────────────┐                                               │
│   │ Calculate Score │                                               │
│   │ (0-100 points)  │                                               │
│   └────────┬────────┘                                               │
│            │                                                        │
│            ↓                                                        │
│   ┌─────────────────────────────────────────┐                       │
│   │ Score 90-100: Excellent                 │→ No action needed     │
│   ├─────────────────────────────────────────┤                       │
│   │ Score 70-89: Good                       │→ Monitor              │
│   ├─────────────────────────────────────────┤                       │
│   │ Score 50-69: Fair                       │→ Email notification   │
│   ├─────────────────────────────────────────┤                       │
│   │ Score 25-49: Poor                       │→ Priority ticket      │
│   ├─────────────────────────────────────────┤                       │
│   │ Score 0-24: Critical                    │→ Immediate alert      │
│   └─────────────────────────────────────────┘                       │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Device Reputation System

Each device builds a reputation over time:

| Factor | Effect on Reputation |
|--------|----------------------|
| Starting Score | 80 (new devices) |
| Clean operation | Improves reputation |
| High uptime | Improves reputation |
| Rejected impressions | Drops reputation |
| Chargebacks | Drops reputation |
| Fraud flags | Major reputation drop |

**Low Reputation Consequences**:
- Score below 30: All impressions reviewed manually
- Score below 10: Device may be decommissioned

## Checklist (Subtasks)

- [ ] Implement health score calculation
- [ ] Create hourly calculation job
- [ ] Build score storage
- [ ] Create health category classification
- [ ] Implement alert triggers
- [ ] Build fleet health dashboard
- [ ] Create score distribution chart
- [ ] Add trend analysis
- [ ] Build attention-needed list
- [ ] Implement maintenance ticket creation
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
