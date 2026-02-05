---
id: "STORY-7.1"
epic_id: "EPIC-07"
title: "Impression Recording"
status: "to-do"
priority: "critical"
assigned_to: null
tags: ["impression", "recording", "tracking"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Impression Recording

## User Story

**As a** platform operator,
**I want** to accurately record every ad impression,
**So that** billing is accurate and analytics are reliable.

## Acceptance Criteria

### Impression Data Captured
- [ ] Device ID
- [ ] Campaign ID
- [ ] Ad group ID
- [ ] Content ID
- [ ] Start timestamp (UTC)
- [ ] End timestamp (UTC)
- [ ] Duration (seconds)
- [ ] Store ID
- [ ] Supplier ID

### Recording Trigger
- [ ] Record when content starts playing
- [ ] Record when content finishes
- [ ] Record if content is interrupted
- [ ] Handle offline mode (queue locally)

### Data Quality
- [ ] Timestamps synchronized with NTP
- [ ] Device clock drift detection
- [ ] Sequence number for ordering
- [ ] Checksum for integrity

### Delivery Methods
- [ ] Real-time streaming (preferred)
- [ ] Batch upload (fallback)
- [ ] Store-and-forward for offline

### Handling Edge Cases
- [ ] Device goes offline mid-play: Mark partial
- [ ] Content loops: Count each play
- [ ] Skip button pressed: Record actual duration
- [ ] Device restart: Resume sequence

## Business Flow

### Recording Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                    IMPRESSION RECORDING FLOW                    │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   [Ad Starts Playing on Screen]                                 │
│          │                                                      │
│          ↓                                                      │
│   Record Start Time                                             │
│   Capture Device + Store Info                                   │
│          │                                                      │
│          ↓                                                      │
│   ┌──────────────┐                                              │
│   │ During Play  │                                              │
│   │              │                                              │
│   │ • Monitor    │                                              │
│   │   duration   │                                              │
│   │ • Capture    │                                              │
│   │   screenshot │                                              │
│   │   (40-60%)   │                                              │
│   └──────┬───────┘                                              │
│          │                                                      │
│          ↓                                                      │
│   ┌──────┴──────┬─────────────┐                                 │
│   │             │             │                                 │
│   ↓             ↓             ↓                                 │
│ Ad Finishes  Interrupted   Device Goes                          │
│ Normally     (Skip/Error)  Offline                              │
│   │             │             │                                 │
│   ↓             ↓             ↓                                 │
│ Full Play    Partial       Queue Locally                        │
│ Recorded     Recorded      (up to 4 hours)                      │
│   │             │             │                                 │
│   └─────────────┴─────────────┘                                 │
│                 │                                               │
│                 ↓                                               │
│   ┌─────────────────────────────────┐                           │
│   │     Send to Platform            │                           │
│   │  • Real-time stream (preferred) │                           │
│   │  • Batch upload (fallback)      │                           │
│   │  • Store-and-forward (offline)  │                           │
│   └─────────────────────────────────┘                           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Minimum Playback Requirements

| Playback Completed | Status | Outcome |
|--------------------|--------|---------|
| 80% or more | Valid Impression | Sent for verification |
| Below 80% | Partial | May not be billable |
| Interrupted/Skip | Partial | Record actual duration |

### Offline Handling Rules

| Scenario | Rule | Time Limit |
|----------|------|------------|
| Network temporarily down | Queue impressions locally | Up to 4 hours |
| Network restored | Send all queued impressions | Immediately |
| Impressions older than 4 hours | Rejected as too stale | Not accepted |

### Clock Synchronization Requirements

| Time Difference from Server | System Response |
|-----------------------------|----------------|
| Within 10 minutes | Accepted normally |
| 10-30 minutes off | Flagged, reduced quality score |
| More than 30 minutes off | Rejected |
| Time in the future | Rejected (possible fraud) |

## Checklist (Subtasks)

- [ ] Design impression event schema
- [ ] Create device recording logic
- [ ] Implement NTP synchronization
- [ ] Build real-time streaming endpoint
- [ ] Create batch upload endpoint
- [ ] Implement local queue for offline
- [ ] Add checksum validation
- [ ] Build ingestion pipeline
- [ ] Create sequence validation
- [ ] Handle partial impressions
- [ ] Implement acknowledgment
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
