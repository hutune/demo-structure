---
id: "STORY-7.2"
epic_id: "EPIC-07"
title: "Impression Verification"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["verification", "validation", "quality"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Impression Verification

## User Story

**As a** platform operator,
**I want** to verify that impressions are valid and billable,
**So that** only legitimate impressions are charged.

## Acceptance Criteria

### Verification Checks
- [ ] **Device Validity**: Device is registered and active
- [ ] **Campaign Validity**: Campaign is active at impression time
- [ ] **Content Match**: Content is assigned to campaign
- [ ] **Schedule Match**: Impression within scheduled time
- [ ] **Geographic Match**: Device in targeted location
- [ ] **Duration Check**: Minimum duration met (e.g., 2+ seconds)

### Verification Results
| Result | Description | Billable |
|--------|-------------|----------|
| Valid | All checks pass | Yes |
| Partial | Minimum duration met but not full | Partial |
| Invalid | Failed critical check | No |
| Suspicious | Passed but flagged | Under review |

### Minimum Duration Rules
- [ ] Static image: 2+ seconds
- [ ] Video: 50%+ of content length
- [ ] HTML5: 3+ seconds

### Fraud Detection
- [ ] Impossible play rate (too many impressions)
- [ ] Unusual patterns (spikes, regularity)
- [ ] Geographic anomalies
- [ ] Device fingerprint mismatch

### Verification Timing
- [ ] Real-time preliminary check
- [ ] Daily batch verification
- [ ] Weekly fraud analysis

## Business Flow

### Impression Lifecycle

```
┌─────────────────────────────────────────────────────────────────┐
│                     IMPRESSION LIFECYCLE                        │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   [Impression Received from Device]                             │
│          │                                                      │
│          ↓                                                      │
│   ┌─────────────┐                                               │
│   │  PENDING    │  ← System receives playback report            │
│   └──────┬──────┘                                               │
│          │                                                      │
│          │ Automatic verification checks                        │
│          ↓                                                      │
│   ┌──────┴──────┬─────────────┬─────────────┐                   │
│   │             │             │             │                   │
│   ↓             ↓             ↓             ↓                   │
│ All checks   Suspicious   Needs review   Clear fraud            │
│   pass        pattern      detected       detected              │
│   │             │             │             │                   │
│   ↓             ↓             ↓             ↓                   │
│ ┌─────────┐ ┌───────────┐ ┌───────────┐ ┌──────────┐            │
│ │VERIFIED │ │  FLAGGED  │ │UNDER_REVIEW│ │ REJECTED │            │
│ │         │ │ but OK    │ │           │ │          │            │
│ └────┬────┘ └─────┬─────┘ └─────┬─────┘ └──────────┘            │
│      │            │             │                               │
│      │            │      Admin reviews                          │
│      │            │             │                               │
│      ├────────────┴──────┬─────┴───────┐                        │
│      │                   │             │                        │
│      ↓                   ↓             ↓                        │
│  Billing              Approved      Rejected                    │
│  processed             by admin      by admin                   │
│      │                   │             │                        │
│      │                   ↓             ↓                        │
│      │              ┌─────────┐   ┌──────────┐                  │
│      │              │VERIFIED │   │ REJECTED │                  │
│      │              └─────────┘   └──────────┘                  │
│      ↓                                                          │
│  Can be disputed                                                │
│  within 30 days                                                 │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Verification Status Outcomes

| Status | What It Means | Billing Impact |
|--------|---------------|----------------|
| Pending | Just received, being checked | Not charged yet |
| Verified | Passed all checks | Advertiser charged, supplier earns revenue |
| Rejected | Failed checks | Not charged, logged for troubleshooting |
| Under Review | Suspicious, needs human review | Billing paused until resolved |
| Disputed | Advertiser challenged the charge | Under investigation |

### Verification Check Requirements

| Check | Requirement | If Failed |
|-------|-------------|----------|
| Device Validity | Device must be registered and Active | Rejected |
| Device Health | Health check within last 10 minutes | Flagged |
| Campaign Status | Campaign must be Active | Rejected |
| Budget Available | Campaign has remaining budget | Rejected |
| Playback Duration | At least 80% of ad duration | May not be billable |
| Time Accuracy | Within 10 minutes of server time | Flagged or Rejected |
| Location Match | Device at correct store location | Flagged or Rejected |

### Distance Verification Rules

| Distance from Store | System Action |
|---------------------|---------------|
| Less than 1 km | Normal (no concern) |
| 1-5 km | Flagged (minor concern) |
| 5-20 km | Held for admin review |
| More than 50 km | Rejected (device not at store) |

## Checklist (Subtasks)

- [ ] Implement device validity check
- [ ] Create campaign validity check
- [ ] Add content match validation
- [ ] Implement schedule matching
- [ ] Add geographic validation
- [ ] Create duration validation
- [ ] Build real-time verification
- [ ] Create batch verification job
- [ ] Implement fraud detection rules
- [ ] Build suspicious flag workflow
- [ ] Create verification dashboard
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
