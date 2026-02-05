---
id: "STORY-6.2"
epic_id: "EPIC-06"
title: "Content Lifecycle"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["content", "lifecycle", "status"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Content Lifecycle

## User Story

**As an** advertiser,
**I want** my content to progress through defined lifecycle stages,
**So that** I can track content status from upload to delivery.

## Acceptance Criteria

### Content Statuses
| Status | Description |
|--------|-------------|
| Uploading | File transfer in progress |
| Processing | Transcoding, validation |
| Pending Review | Awaiting moderation |
| Approved | Ready for use in campaigns |
| Rejected | Failed moderation |
| In Use | Assigned to active campaign |
| Archived | No longer in active use |
| Deleted | Marked for deletion |

### Status Transitions
```
Uploading → Processing → Pending Review → Approved → In Use
                              ↓                        ↓
                          Rejected                 Archived
                              ↓                        ↓
                          (Edit)               Deleted/Restored
```

### Automatic Transitions
- [ ] Uploading → Processing: Upload complete
- [ ] Processing → Pending Review: Processing complete
- [ ] Approved → In Use: Assigned to active campaign
- [ ] In Use → Approved: Campaign ends

### Manual Transitions
- [ ] Pending Review → Approved: Moderator approves
- [ ] Pending Review → Rejected: Moderator rejects
- [ ] Any → Archived: User archives
- [ ] Archived → Approved: User restores

### Delete Behavior
- [ ] Soft delete (mark as deleted)
- [ ] Cannot delete if in active campaign
- [ ] Permanent delete after 30 days
- [ ] Admin can force delete

## Business Flow

### Complete Content Status Flow

```
+----------+      File received       +--------------------+
| UPLOADED |------------------------->| PROCESSING         |
+----------+                          +--------------------+
                                             |
                                    Validation complete
                                             |
                                             v
                                      +--------------------+
                                      | PENDING_APPROVAL   |
                                      +--------------------+
                                             |
                         +-------------------+-------------------+
                         |                                       |
                    Approved                                 Rejected
                         |                                       |
                         v                                       v
                   +----------+                            +----------+
                   | APPROVED |                            | REJECTED |
                   +----------+                            +----------+
                         |                                       |
               Assigned to campaign                        Can appeal or
                         |                                  upload new
                         v
                   +----------+
                   | ACTIVE   | <-- Currently in campaigns
                   +----------+
                         |
               Campaign ends or user archives
                         |
                         v
                   +----------+
                   | ARCHIVED |
                   +----------+
```

### Status Usage Rules

| Status | Can Use in Campaign? | Notes |
|--------|---------------------|-------|
| UPLOADED | No | File received, awaiting processing |
| PROCESSING | No | System validating and preparing |
| PENDING_APPROVAL | No | Awaiting moderation review |
| APPROVED | Yes | Passed moderation, ready for use |
| REJECTED | No | Failed moderation or validation |
| ACTIVE | Yes | Currently being used in campaigns |
| ARCHIVED | No* | *Existing campaigns continue, no new assignments |

### Deletion Flow

```
User clicks Delete
        |
        v
+------------------+
| In active        |----YES---> Show warning:
| campaign?        |            "Cannot delete while in use"
+------------------+
        | NO
        v
+------------------+
| SOFT DELETE      |  (Hidden from library, 30-day recovery)
+------------------+
        |
   After 30 days
        |
        v
+------------------+
| PERMANENT DELETE |  (Cannot be recovered)
+------------------+
```

### Archival Rules

| Action | Behavior |
|--------|----------|
| Manual Archive | Move to Archived folder, not in main library |
| Auto-Archive | After 365 days unused, notification sent 30 days before |
| Restore | Move back to active library, immediately available |
| Existing Campaigns | Continue without interruption |

## Checklist (Subtasks)

- [ ] Implement content status enum
- [ ] Create state machine for transitions
- [ ] Add transition validation
- [ ] Implement auto-transitions
- [ ] Build status change notifications
- [ ] Create soft delete functionality
- [ ] Add permanent delete scheduler
- [ ] Implement restore from archive
- [ ] Create status history log
- [ ] Add in-use detection
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
