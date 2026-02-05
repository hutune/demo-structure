---
id: "STORY-5.4"
epic_id: "EPIC-05"
title: "Campaign Editing & Cloning"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["campaign", "editing", "cloning"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Campaign Editing & Cloning

## User Story

**As an** advertiser,
**I want** to edit existing campaigns and clone successful ones,
**So that** I can optimize and replicate my advertising efforts.

## Acceptance Criteria

### Edit by Status
| Status | Editable Fields |
|--------|----------------|
| Draft | All fields |
| Pending Review | None (must withdraw first) |
| Scheduled | Budget, dates, targeting |
| Active | Budget (increase), end date, targeting |
| Paused | Same as Active |
| Completed | None (clone instead) |
| Cancelled | None (clone instead) |

### Edit Restrictions for Active Campaigns
- [ ] Budget can be increased, not decreased
- [ ] End date can be extended, not shortened
- [ ] Content changes require re-submission
- [ ] Some targeting changes apply next day

### Edit Workflow
- [ ] Click Edit on campaign
- [ ] Make changes
- [ ] Preview changes
- [ ] Confirm changes
- [ ] Changes logged in history

### Campaign Cloning
- [ ] Clone any campaign (including completed)
- [ ] Choose what to copy:
  - [ ] Budget settings
  - [ ] Schedule (adjust dates)
  - [ ] Targeting
  - [ ] Content
- [ ] New campaign created as Draft
- [ ] Original campaign linked for reference

### Version History
- [ ] Track all changes
- [ ] Show who made changes
- [ ] Show when changes were made
- [ ] Compare versions

## Business Flow

### Edit Permissions by Status

| Field | DRAFT | PENDING | SCHEDULED | ACTIVE | PAUSED |
|-------|-------|---------|-----------|--------|--------|
| Name | Edit | Edit | Edit | Edit | Edit |
| Description | Edit | Edit | Edit | Edit | Edit |
| Brand Name | Edit | Edit | Read-only | Read-only | Read-only |
| Category | Edit | Edit | Read-only | Read-only | Read-only |
| Budget | Edit | Edit | Increase only | Increase only | Increase only |
| Start Date | Edit | Edit | Delay only | Read-only | Read-only |
| End Date | Edit | Edit | Edit | Extend only | Extend only |
| Daily Cap | Edit | Edit | Edit | Edit | Edit |
| Target Stores | Edit | Edit | Add only | Add only | Add only |
| Content | Edit | Read-only* | Read-only | Add only** | Read-only |
| Priority | Edit | Edit | Edit | Edit | Edit |

*Content changes in PENDING_APPROVAL require re-submission
**Can add new approved assets to ACTIVE campaign

### Budget Change Rules

**Increase Budget**:
- Allowed in: DRAFT, PENDING_APPROVAL, SCHEDULED, ACTIVE, PAUSED
- Requires: Sufficient wallet balance
- Process: Additional hold created, budget updated immediately

**Decrease Budget**:
- Allowed in: DRAFT only
- Not allowed once submitted (budget is held)

### Date Change Rules

**Start Date**:
- Can delay in SCHEDULED status (must maintain 24-hour lead time)
- Cannot advance (make earlier) after SCHEDULED

**End Date**:
- Can extend if budget remaining
- Can shorten in SCHEDULED (triggers recalculation)

### Store Targeting Change Rules

**Add Stores**:
- Allowed in: SCHEDULED, ACTIVE, PAUSED
- New stores checked against blocking rules
- Content distribution begins immediately

**Remove Stores**:
- Allowed in: DRAFT, PENDING_APPROVAL only
- Cannot remove once campaign is scheduled/active

### Edit Audit Trail

All campaign edits are logged:
- What field changed
- Old value and new value
- Who made the change
- When the change was made

Audit logs retained for 7 years.

## Checklist (Subtasks)

- [ ] Implement edit permissions by status
- [ ] Create edit form with field restrictions
- [ ] Add change preview
- [ ] Implement version history tracking
- [ ] Create clone dialog
- [ ] Add clone options selection
- [ ] Implement date shifting for clones
- [ ] Link cloned campaign to original
- [ ] Create version comparison view
- [ ] Add edit confirmation workflow
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
