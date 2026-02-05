---
id: "STORY-6.4"
epic_id: "EPIC-06"
title: "Content Moderation"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["moderation", "review", "compliance"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Content Moderation

## User Story

**As a** platform operator,
**I want** to moderate all advertising content,
**So that** inappropriate content is not displayed on screens.

## Acceptance Criteria

### Moderation Levels
- [ ] **Level 1**: Automated AI moderation
- [ ] **Level 2**: Human review (flagged or all)

### Automated Checks
- [ ] Adult/explicit content detection
- [ ] Violence detection
- [ ] Hate speech/symbols detection
- [ ] Political content detection
- [ ] Brand safety compliance
- [ ] Copyright infringement indicators

### Content Guidelines
- [ ] No nudity or explicit content
- [ ] No violence or graphic imagery
- [ ] No hate speech or discrimination
- [ ] No political advertising (unless allowed)
- [ ] No misleading claims
- [ ] No competitor disparagement
- [ ] No prohibited products (tobacco, etc.)

### Moderation Queue
- [ ] View pending content
- [ ] Preview content (image/video player)
- [ ] One-click approve
- [ ] Reject with reason
- [ ] Escalate to senior moderator
- [ ] Bulk actions

### Rejection Reasons
- [ ] Explicit content
- [ ] Violent content
- [ ] Hate speech
- [ ] Political content
- [ ] Misleading claims
- [ ] Copyright violation
- [ ] Technical issues
- [ ] Other (with notes)

### Appeals Process
- [ ] Advertiser can appeal rejection
- [ ] Appeal reviewed by senior moderator
- [ ] Appeal result is final

## Business Flow

### Two-Tier Moderation Process

```
        CONTENT UPLOADED
              |
              v
    +------------------+
    | TIER 1: AI SCAN  |
    | (All uploads)    |
    +------------------+
              |
              v
    +------------------+
    | Calculate AI     |
    | Confidence Score |
    +------------------+
              |
    +---------+---------+---------+
    |         |                   |
  90-100    70-89              Below 70
    |         |                   |
    v         v                   v
+--------+ +--------+        +----------+
| AUTO-  | | FLAG   |        | AUTO-    |
| APPROVE| | FOR    |        | REJECT   |
+--------+ | REVIEW |        +----------+
    |      +--------+             |
    |         |                   v
    |         v            Notify advertiser
    |   +-----------+      with reason
    |   | TIER 2:   |
    |   | MANUAL    |
    |   | REVIEW    |
    |   +-----------+
    |         |
    |   +-----+-----+
    |   |           |
    v   v           v
+----------+   +----------+
| APPROVED |   | REJECTED |
+----------+   +----------+
```

### AI Scoring Actions

| Score Range | Action | SLA |
|-------------|--------|-----|
| 90-100 | Auto-approve, ready immediately | Instant |
| 70-89 | Flag for manual review | 24 hours (Standard), 4 hours (Enterprise) |
| Below 70 | Auto-reject with reason | Instant |

### Content Policy Categories

| Category | Action |
|----------|--------|
| **Prohibited** (Always Rejected) | |
| Adult/sexual content | Auto-reject |
| Violence, gore, graphic imagery | Auto-reject |
| Hate speech, discrimination | Auto-reject |
| Misleading/deceptive claims | Auto-reject |
| Illegal products or services | Auto-reject |
| Copyright/trademark infringement | Auto-reject |
| **Restricted** (Special Approval) | |
| Alcohol advertising | Requires age-gating approval |
| Gambling | Licensed advertisers only |
| Political campaigns | Special approval required |
| Healthcare/pharmaceutical | Regulatory compliance check |
| Financial services | Disclosure requirements |

### Manual Review Workflow

```
Review Queue (Priority Order)
           |
           v
+------------------------+
| 1. Enterprise customers|
| 2. Standard customers  |
| 3. Free tier          |
+------------------------+
           |
           v
+------------------------+
| Reviewer sees:         |
| - Thumbnail           |
| - AI flags            |
| - Metadata            |
+------------------------+
           |
     +-----+-----+
     |     |     |
     v     v     v
APPROVE REJECT REQUEST
                CHANGES
     |     |     |
     v     v     v
Notify  Reason  Specific
user    required edits
                 needed
```

### Appeal Process Flow

```
     CONTENT REJECTED
           |
           v
+---------------------+
| Advertiser submits  |
| appeal with         |
| explanation         |
+---------------------+
           |
           v
+---------------------+
| (Optional) Upload   |
| revised version     |
+---------------------+
           |
           v
+---------------------+
| ESCALATED TO        |
| SENIOR REVIEWER     |
+---------------------+
           |
    Review SLA:
    - Standard: 48 hours
    - Enterprise: 8 hours
           |
     +-----+-----+
     |     |     |
     v     v     v
APPROVE REJECT REQUEST
(overturn)(uphold) CHANGES
           |
           v
    Decision is FINAL

    Limit: 1 appeal per asset
```

## Checklist (Subtasks)

- [ ] Integrate AI moderation service
- [ ] Build moderation queue interface
- [ ] Create content preview player
- [ ] Implement approve/reject workflow
- [ ] Add rejection reason selection
- [ ] Create escalation feature
- [ ] Implement bulk actions
- [ ] Build appeals workflow
- [ ] Add moderator assignment
- [ ] Create moderation metrics dashboard
- [ ] Set up SLA tracking
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
