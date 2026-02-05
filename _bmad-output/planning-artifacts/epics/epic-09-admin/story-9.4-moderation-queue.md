---
id: "STORY-9.4"
epic_id: "EPIC-09"
title: "Moderation Queue"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["moderation", "queue", "review"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Moderation Queue

## User Story

**As a** content moderator,
**I want** an efficient queue to review content,
**So that** I can process submissions quickly and consistently.

## Acceptance Criteria

### Queue Types
- [ ] Content moderation (ad creatives)
- [ ] KYC document review
- [ ] Account verification
- [ ] Reported content
- [ ] Dispute review

### Queue Features
- [ ] Prioritized by age (oldest first)
- [ ] Filter by type, status, priority
- [ ] Assignment to moderators
- [ ] Claim next available
- [ ] Reassign items

### Content Review Interface
- [ ] Full content preview
- [ ] Side-by-side with guidelines
- [ ] One-click approve
- [ ] Reject with reason selection
- [ ] Escalate to senior moderator
- [ ] Request more info

### KYC Review Interface
- [ ] Document viewer (zoom, rotate)
- [ ] ID verification checklist
- [ ] Cross-reference information
- [ ] Approve/Reject with notes

### Quality Assurance
- [ ] Random sample review
- [ ] Moderator accuracy tracking
- [ ] Consistency checks
- [ ] Regular calibration sessions

### Metrics Dashboard
- [ ] Queue depth
- [ ] Average wait time
- [ ] Moderator throughput
- [ ] Approval/rejection rates

## Business Flow

### Content Moderation Flow

```
CONTENT UPLOADED
     │
     ▼
+------------------+
│ AI Moderation    │
│ Score (0-100)    │
+------------------+
     │
     ├─ Score ≥ 90 ────► Auto-APPROVE (<5 min)
     │
     ├─ Score 70-89 ───► MANUAL REVIEW
     │                    │
     │                    ├─ Enterprise → 8 hr queue
     │                    ├─ Premium → 24 hr queue
     │                    └─ Others → 48 hr queue
     │
     └─ Score < 70 ────► Auto-REJECT (<5 min)
     │
     ▼
+------------------+
│ User Notified    │
│ of Decision      │
+------------------+
```

### Queue Priority Order

1. Enterprise customer content
2. Campaign starting within 24 hours
3. Appeal reviews
4. Standard queue (FIFO)

### Moderation SLA by Customer Tier

| Customer Tier | Standard Queue | Priority Queue |
|---------------|----------------|----------------|
| FREE | 48 hours | N/A |
| BASIC | 24 hours | N/A |
| PREMIUM | 24 hours | 8 hours |
| ENTERPRISE | 8 hours | 4 hours |

### Moderator Decisions

| Decision | Effect | Notification |
|----------|--------|-------------|
| Approve | Content ready for use | "Content approved" |
| Reject | Cannot be used | Reason + resubmit option |
| Request Changes | Minor edits needed | Specific feedback |
| Escalate | Needs senior review | Processing notice |

### Notification Categories

```
NOTIFICATION CATEGORIES
├── ACCOUNT
│   ├── Security (login, password, 2FA)
│   ├── Verification (KYC, documents)
│   ├── Tier changes
│   └── Team management
├── FINANCIAL
│   ├── Balance alerts
│   ├── Top-up confirmations
│   ├── Withdrawal status
│   └── Invoice generated
├── CAMPAIGN
│   ├── Status changes
│   ├── Budget alerts
│   ├── Performance updates
│   └── Approval required
├── DEVICE
│   ├── Health alerts
│   ├── Offline warnings
│   └── Sync failures
├── CONTENT
│   ├── Moderation status
│   ├── Approval/rejection
│   └── Expiration warnings
└── SYSTEM
    ├── Scheduled maintenance
    ├── Feature announcements
    └── API deprecation
```

### Device Notification Rules

| Event | Priority | Escalation | Action |
|-------|----------|------------|--------|
| Device offline (10 min) | HIGH | → CRITICAL at 1 hour | Troubleshooting steps |
| Device back online | MEDIUM | N/A | Downtime summary |
| Health score < 70 | HIGH | Daily max | Improvement recommendations |
| Content sync failed | HIGH | After 3 retries | Manual intervention steps |

## Checklist (Subtasks)

- [ ] Design queue data model
- [ ] Create queue listing interface
- [ ] Build priority calculation
- [ ] Implement claim functionality
- [ ] Create content review UI
- [ ] Build KYC review UI
- [ ] Add approval workflow
- [ ] Create rejection workflow
- [ ] Implement escalation
- [ ] Build assignment management
- [ ] Create moderator dashboard
- [ ] Add QA sampling
- [ ] Build metrics reporting
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
