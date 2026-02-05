---
id: "STORY-1.13"
epic_id: "EPIC-01"
title: "Employee Invitation Flow"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["invitation", "team", "onboarding"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Employee Invitation Flow

## User Story

**As a** company admin,
**I want** to invite team members to join my company account,
**So that** they can help manage campaigns and operations.

## Acceptance Criteria

### Invitation Process
- [ ] Admin can send invitation to email address
- [ ] Invitation email contains unique link
- [ ] Invitation expires after 7 days
- [ ] Only one active invitation per email per company
- [ ] Admin can cancel invitation before acceptance
- [ ] Expired invitations can be resent by admin

### Invitation Acceptance
- [ ] Invitee clicks link to accept invitation
- [ ] Invitee enters personal information
- [ ] If invitee has no account, create one (skip email verification)
- [ ] Account created with status = ACTIVE
- [ ] Account linked to company with assigned role
- [ ] One user can be on multiple company teams

### Invitation Status
- [ ] PENDING: Sent, awaiting response
- [ ] ACCEPTED: Invitee joined
- [ ] EXPIRED: 7 days passed without action
- [ ] REVOKED: Admin cancelled

## Business Flow

### Employee Invitation Flow
```
Admin sends invitation to email
          |
          v
  Invitee receives email with link
          |
          v
  Invitee clicks link, verifies identity
          |
          v
  Invitee enters personal information
          |
          v
  Account created --> Status = ACTIVE
          |
          v
  Account linked to company with assigned role
```

### Invitation Status Flow
```
+----------+     Accepted     +----------+
| PENDING  |----------------->| ACCEPTED |
+----------+                  +----------+
     |
     +--------+--------+
     |                 |
     v                 v
+---------+      +---------+
| EXPIRED |      | REVOKED |
+---------+      +---------+
 (7 days)       (by admin)
```

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
