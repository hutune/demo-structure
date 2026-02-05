---
id: "STORY-5.2"
epic_id: "EPIC-05"
title: "Campaign Lifecycle"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["campaign", "lifecycle", "status"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Campaign Lifecycle

## User Story

**As an** advertiser,
**I want** my campaigns to progress through defined lifecycle stages,
**So that** I can track and manage campaign status.

## Acceptance Criteria

### Campaign Statuses
| Status | Description | Can Edit? |
|--------|-------------|-----------|
| Draft | Being configured | Full edit |
| Pending Review | Awaiting moderation | No edit |
| Rejected | Failed review | Edit to resubmit |
| Scheduled | Approved, waiting to start | Limited edit |
| Active | Currently running | Limited edit |
| Paused | Temporarily stopped | Resume only |
| Completed | Finished successfully | Read only |
| Cancelled | Terminated early | Read only |

### Status Transitions
```
Draft → Pending Review → Scheduled → Active → Completed
                ↓              ↘        ↓
             Rejected          Paused → Active
                ↓                       ↓
         (Edit) Draft              Cancelled
```

### Automated Transitions
- [ ] Scheduled → Active: When start time reached
- [ ] Active → Completed: When end time reached
- [ ] Active → Paused: When daily/total budget exhausted
- [ ] Paused → Active: At midnight if budget reset (daily)

### Manual Transitions
- [ ] Draft → Pending Review: Advertiser submits
- [ ] Pending Review → Scheduled: Moderator approves
- [ ] Pending Review → Rejected: Moderator rejects
- [ ] Active → Paused: Advertiser pauses
- [ ] Paused → Active: Advertiser resumes
- [ ] Any → Cancelled: Advertiser cancels

### Edit Restrictions by Status
- [ ] Scheduled: Can edit dates, budget, targeting
- [ ] Active: Can edit budget, end date, targeting
- [ ] Active: Cannot change content (requires new review)

## Business Flow

### Campaign Status Flow

```
+-------+      Submit      +------------------+
| DRAFT |----------------->| PENDING_APPROVAL |
+-------+                  +------------------+
                                   |
                      +------------+------------+
                      |                         |
                      v                         v
              +-----------+              +-----------+
              | SCHEDULED |              | REJECTED  |
              +-----------+              +-----------+
                    |
                    | Start date reached
                    v
              +--------+
              | ACTIVE |<---------+
              +--------+          |
                    |             |
         +----------+------+      |
         |                 |      |
         v                 v      |
    +--------+        +-----------+
    | PAUSED |------->| (Resume)  |
    +--------+        +-----------+
         |
         v
   +-----------+       +-----------+
   | COMPLETED |       | CANCELLED |
   +-----------+       +-----------+
```

### Status Descriptions

| Status | Description | User Can Edit |
|--------|-------------|---------------|
| DRAFT | Being created, not submitted | Full editing |
| PENDING_APPROVAL | Awaiting admin review | Limited editing |
| SCHEDULED | Approved, waiting for start date | Limited editing |
| ACTIVE | Currently serving ads | Very limited editing |
| PAUSED | Temporarily stopped | Limited editing |
| COMPLETED | Ended normally | Read-only |
| CANCELLED | Terminated by user | Read-only |
| REJECTED | Admin rejected | Can resubmit after fixes |

### Status Transition Rules

| From | To | Triggered By |
|------|-----|--------------|
| DRAFT | PENDING_APPROVAL | Advertiser submits |
| DRAFT | SCHEDULED | Auto-approve (if criteria met) |
| PENDING_APPROVAL | SCHEDULED | Admin approves |
| PENDING_APPROVAL | REJECTED | Admin rejects |
| SCHEDULED | ACTIVE | Start date reached |
| ACTIVE | PAUSED | User pauses or system auto-pause |
| ACTIVE | COMPLETED | End date reached or budget exhausted |
| PAUSED | ACTIVE | User resumes |
| PAUSED | CANCELLED | User cancels |
| PAUSED | COMPLETED | End date reached while paused |
| Any non-terminal | CANCELLED | User cancels |

### Campaign Completion Flow

```
End condition met (date reached or budget spent)
         |
         v
Allow 5-minute grace period (in-flight impressions)
         |
         v
Final reconciliation:
    - Process remaining impressions
    - Calculate final totals
         |
         v
Release unused budget to wallet
         |
         v
Generate final report
         |
         v
Notify advertiser:
    "Campaign completed"
    - Total impressions
    - Total spent
    - Effective CPM
    - Refund amount (if any)
```

## Checklist (Subtasks)

- [ ] Implement campaign status enum
- [ ] Create state machine for transitions
- [ ] Add transition validation rules
- [ ] Implement automated start scheduler
- [ ] Implement automated completion scheduler
- [ ] Add budget exhaustion detection
- [ ] Create pause/resume functionality
- [ ] Implement cancel flow
- [ ] Add status change notifications
- [ ] Create status history log
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
