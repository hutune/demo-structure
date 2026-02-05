---
id: "STORY-7.3"
epic_id: "EPIC-07"
title: "Proof-of-Play"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["proof", "evidence", "reporting"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Proof-of-Play

## User Story

**As an** advertiser,
**I want** evidence that my ads were actually displayed,
**So that** I can trust the billing and verify campaign delivery.

## Acceptance Criteria

### Proof Types
- [ ] **Log Records**: Timestamped play logs
- [ ] **Screenshots**: Periodic screen captures
- [ ] **Video Clips**: Short recordings (optional)
- [ ] **Attestation**: Cryptographic proof

### Log Records
- [ ] Play start and end times
- [ ] Device and store information
- [ ] Content displayed
- [ ] Duration
- [ ] Sequence number

### Screenshot Capture
- [ ] Capture during content play
- [ ] Low resolution to save bandwidth
- [ ] Store for 90 days
- [ ] Watermarked with timestamp
- [ ] Privacy-compliant (no people identification)

### Proof-of-Play Report
- [ ] Downloadable per campaign
- [ ] Filter by date range
- [ ] Filter by store/device
- [ ] Include summary statistics
- [ ] Include sample screenshots
- [ ] PDF and CSV formats

### Access Control
- [ ] Advertiser: Own campaigns only
- [ ] Supplier: Own devices only
- [ ] Admin: All data

## Business Flow

### Proof Collection Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                    PROOF-OF-PLAY COLLECTION                     │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   [Ad Starts Playing]                                           │
│          │                                                      │
│          ↓                                                      │
│   Record Play Start Time                                        │
│   + Device Information                                          │
│   + Store Location                                              │
│          │                                                      │
│          ↓                                                      │
│   ┌──────────────────────────────┐                              │
│   │ At 40-60% of Ad Duration:   │                               │
│   │                              │                               │
│   │ • Capture Screenshot         │                               │
│   │ • Create Image Fingerprint   │                               │
│   │ • Generate Digital Signature │                               │
│   │ • Record GPS (if available)  │                               │
│   └──────────────┬───────────────┘                              │
│                  │                                              │
│                  ↓                                              │
│   ┌──────────────┴───────────────┐                              │
│   │    Screenshot Selection      │                              │
│   │                              │                              │
│   │  • 5% randomly selected      │                              │
│   │    for full upload           │                              │
│   │  • All flagged impressions   │                              │
│   │    uploaded                  │                              │
│   │  • Fingerprints always sent  │                              │
│   └──────────────┬───────────────┘                              │
│                  │                                              │
│                  ↓                                              │
│   [Ad Finishes Playing]                                         │
│          │                                                      │
│          ↓                                                      │
│   Send Complete Proof Package:                                  │
│   • Timestamp (start + end)                                     │
│   • Screenshot fingerprint                                      │
│   • Digital signature                                           │
│   • Location data                                               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Proof Elements Required

| Proof Type | What It Is | Why It Matters |
|------------|------------|----------------|
| Log Records | Timestamped play start/end | Proves timing is accurate |
| Screenshot | Picture captured during playback | Proves the right content was shown |
| Digital Signature | Unique code from the device | Proves the report came from actual device |
| Location (optional) | GPS coordinates | Proves the device is at store location |

### Screenshot Rules

| Rule | Requirement |
|------|-------------|
| Capture timing | Between 40-60% of ad duration |
| Full upload rate | 5% of impressions (randomly selected) |
| Flagged impressions | Always fully uploaded |
| Fingerprint | Always sent with every impression |
| Storage duration | 30 days, then deleted |
| Privacy | No people identification allowed |

### Proof Retention Schedule

| Data Type | Retention Period | After Expiry |
|-----------|------------------|-------------|
| Full screenshots | 30 days | Deleted for privacy/storage |
| Screenshot fingerprints | Longer term | Kept for audit purposes |
| Play logs | 90 days | Archived |
| Digital signatures | 90 days | Archived |

## Checklist (Subtasks)

- [ ] Design screenshot capture system
- [ ] Implement random capture trigger
- [ ] Add watermarking service
- [ ] Create screenshot upload flow
- [ ] Build attestation signing
- [ ] Create proof linking to impressions
- [ ] Build proof-of-play report
- [ ] Add date/location filters
- [ ] Implement PDF generation
- [ ] Implement CSV export
- [ ] Set up 90-day retention
- [ ] Add access control
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
