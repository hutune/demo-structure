---
id: "STORY-2.4"
epic_id: "EPIC-02"
title: "Team Management"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["team", "collaboration", "roles"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Team Management

## User Story

**As an** advertiser account owner,
**I want** to invite team members with specific roles,
**So that** my team can collaborate on campaigns.

## Acceptance Criteria

### Team Member Limits by Tier
- [ ] FREE: 1 member (owner only)
- [ ] BASIC: 3 members
- [ ] PREMIUM: 10 members
- [ ] ENTERPRISE: Unlimited

### Team Roles
- [ ] Owner: Everything (only account creator)
- [ ] Admin: Everything except billing changes
- [ ] Campaign Manager: Create, edit, manage campaigns
- [ ] Content Manager: Upload and manage ad content
- [ ] Analyst: View reports only
- [ ] Viewer: View everything, change nothing

### Invitation Process
- [ ] Enter email address
- [ ] Select role
- [ ] Invitation email sent
- [ ] Invitation expires after 7 days
- [ ] Must have available slot in tier limit
- [ ] User creates account if needed
- [ ] One user can be on multiple teams

### Member Removal
- [ ] Access revoked immediately
- [ ] Campaigns they created stay with account
- [ ] They cannot see any account data after removal

## Business Flow

### Team Member Invitation Flow

```
    Owner/Admin initiates invite
           │
           v
    ┌──────────────────┐
    │ Check team slot  │
    │ availability     │
    └────────┬─────────┘
             │
      ┌──────┴──────┐
      v             v
   No Slots      Slots Available
   Available          │
      │               v
      v         Enter email address
   Show upgrade       │
   prompt             v
                 Select role
                      │
                      v
               Send invitation email
                      │
                      v
               ┌─────────────────┐
               │ Invitation sent │
               │ (Valid 7 days)  │
               └────────┬────────┘
                        │
                 ┌──────┴──────┐
                 v             v
            Accepted       Expired
                 │         (No action)
                 v
         ┌───────────────┐
         │ Has account?  │
         └───────┬───────┘
                 │
          ┌──────┴──────┐
          v             v
        Yes            No
          │             │
          │             v
          │      Create account
          │             │
          └──────┬──────┘
                 v
          Join team with
          assigned role
```

### Team Roles & Permissions

| Role | Campaigns | Content | Billing | Reports | Team |
|------|-----------|---------|---------|---------|------|
| Owner | Full | Full | Full | Full | Full |
| Admin | Full | Full | View only | Full | Manage |
| Campaign Manager | Full | View | - | View | - |
| Content Manager | View | Full | - | View | - |
| Analyst | View | View | - | Full | - |
| Viewer | View | View | - | View | - |

### Team Member Limits by Tier

| Tier | Team Members Allowed |
|------|----------------------|
| FREE | 1 (owner only) |
| BASIC | 3 |
| PREMIUM | 10 |
| ENTERPRISE | Unlimited |

## Checklist (Subtasks)

- [ ] Create team member listing page
- [ ] Implement invitation flow
- [ ] Create role selection component
- [ ] Add team slot checking
- [ ] Implement invitation acceptance
- [ ] Create member removal feature
- [ ] Add role change feature
- [ ] Build permission checking for each role
- [ ] Add invitation management (resend, cancel)
- [ ] Create upgrade prompt when slots full
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
