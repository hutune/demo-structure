---
id: "STORY-7.5"
epic_id: "EPIC-07"
title: "Duplicate Detection"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["duplicate", "dedup", "fraud"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Duplicate Detection

## User Story

**As a** platform operator,
**I want** to detect and handle duplicate impressions,
**So that** billing is accurate and fraud is prevented.

## Acceptance Criteria

### Duplicate Types
- [ ] **Exact Duplicate**: Same event sent twice
- [ ] **Logical Duplicate**: Same play reported with different IDs
- [ ] **Replay Attack**: Old impression replayed

### Detection Methods
- [ ] Event ID uniqueness check
- [ ] Device + Content + Timestamp proximity
- [ ] Sequence number validation
- [ ] Signature verification

### Detection Rules
| Rule | Threshold | Action |
|------|-----------|--------|
| Same event_id | Exact match | Reject second |
| Same device+content | Within 5 seconds | Flag as duplicate |
| Sequence gap | > 10 missing | Investigate |
| Timestamp in past | > 1 hour old | Reject or flag |

### Handling Duplicates
- [ ] First occurrence: Keep
- [ ] Subsequent: Mark as duplicate
- [ ] Duplicates not billable
- [ ] Track duplicate rates per device

### Duplicate Reporting
- [ ] Daily duplicate report
- [ ] Duplicate rate by device
- [ ] Duplicate rate by supplier
- [ ] Trend analysis

## Business Flow

### Duplicate Detection Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                   DUPLICATE DETECTION FLOW                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   [Impression Received]                                         │
│          │                                                      │
│          ↓                                                      │
│   ┌─────────────────────────────┐                               │
│   │  Check 1: Exact Duplicate   │                               │
│   │  Same event sent twice?     │                               │
│   └──────────┬──────────────────┘                               │
│              │                                                  │
│       ┌──────┴──────┐                                           │
│       │             │                                           │
│       ↓             ↓                                           │
│    Unique     Already Exists                                    │
│       │             │                                           │
│       │             ↓                                           │
│       │        REJECT as Duplicate                              │
│       │                                                         │
│       ↓                                                         │
│   ┌─────────────────────────────┐                               │
│   │  Check 2: Logical Duplicate │                               │
│   │  Same device + content      │                               │
│   │  within 5 minutes?          │                               │
│   └──────────┬──────────────────┘                               │
│              │                                                  │
│       ┌──────┴──────┐                                           │
│       │             │                                           │
│       ↓             ↓                                           │
│    Clear       Too Soon                                         │
│       │        (within 5 min)                                   │
│       │             │                                           │
│       │             ↓                                           │
│       │        FLAG as Duplicate                                │
│       │                                                         │
│       ↓                                                         │
│   ┌─────────────────────────────┐                               │
│   │  Check 3: Replay Attack     │                               │
│   │  Old impression replayed?   │                               │
│   │  (More than 1 hour old)     │                               │
│   └──────────┬──────────────────┘                               │
│              │                                                  │
│       ┌──────┴──────┐                                           │
│       │             │                                           │
│       ↓             ↓                                           │
│    Fresh       Stale/Replayed                                   │
│       │             │                                           │
│       ↓             ↓                                           │
│   ACCEPTED    REJECT or FLAG                                    │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Duplicate Window Rules

| Same Ad + Same Device | Time Gap | System Action |
|-----------------------|----------|---------------|
| First occurrence | N/A | Recorded |
| Second occurrence | Within 5 minutes | Rejected as duplicate |
| Third occurrence | After 5 minutes from first | Recorded (new window) |

**Example:**
```
Campaign A on Device #123:
- 2:30 PM → Recorded ✓
- 2:33 PM → Duplicate, rejected ✗ (only 3 min gap)
- 2:36 PM → Recorded ✓ (new 5-minute window)
```

### Hourly Frequency Check

| Frequency vs Expected | System Response |
|-----------------------|----------------|
| Normal (as configured) | Accepted |
| 50% higher than expected | Flagged but accepted |
| 100% higher than expected | Held for admin review |

### Duplicate Handling Outcomes

| Outcome | Billing Impact | Tracking |
|---------|----------------|----------|
| First occurrence kept | Billable | Normal impression |
| Subsequent marked duplicate | Not billable | Logged for analysis |
| High duplicate rate | Alert triggered | Device investigated |

## Checklist (Subtasks)

- [ ] Implement event_id caching
- [ ] Create logical duplicate detection
- [ ] Add sequence number validation
- [ ] Implement replay detection
- [ ] Build deduplication pipeline
- [ ] Create duplicate marking
- [ ] Build duplicate reporting
- [ ] Add rate tracking by device
- [ ] Create trend analysis
- [ ] Set up alerts for high rates
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
