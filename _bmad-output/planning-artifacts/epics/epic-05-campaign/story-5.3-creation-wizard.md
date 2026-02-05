---
id: "STORY-5.3"
epic_id: "EPIC-05"
title: "Campaign Creation Wizard"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["campaign", "creation", "wizard"]
story_points: 13
sprint: null
start_date: null
due_date: null
time_estimate: "5d"
clickup_task_id: null
---

# Campaign Creation Wizard

## User Story

**As an** advertiser,
**I want** a step-by-step wizard to create campaigns,
**So that** I can easily set up effective advertising campaigns.

## Acceptance Criteria

### Step 1: Objective & Basic Info
- [ ] Campaign name (required)
- [ ] Campaign type / objective
- [ ] Description (optional)
- [ ] Save as draft available at any step

### Step 2: Budget & Billing
- [ ] Total campaign budget
- [ ] Daily budget cap (optional)
- [ ] Billing model selection (CPM, CPC, etc.)
- [ ] Bid amount (if applicable)
- [ ] Currency selection

### Step 3: Schedule
- [ ] Start date and time
- [ ] End date and time (or "run until budget depleted")
- [ ] Timezone selection
- [ ] Dayparting schedule (optional)

### Step 4: Targeting
- [ ] Geographic targeting (regions, cities, stores)
- [ ] Store categories (grocery, convenience, etc.)
- [ ] Device types
- [ ] Audience segments (if available)
- [ ] Competitor blocking awareness

### Step 5: Content
- [ ] Select from content library
- [ ] Upload new content
- [ ] Assign content to ad groups
- [ ] Set creative rotation

### Step 6: Review & Submit
- [ ] Review all settings
- [ ] Estimated reach calculator
- [ ] Budget summary
- [ ] Submit for review

### Wizard Features
- [ ] Progress indicator
- [ ] Back/Next navigation
- [ ] Save draft at any step
- [ ] Validation before next step
- [ ] Auto-save every 30 seconds

## Business Flow

### Campaign Creation Step-by-Step Process

```
Step 1: Basic Information
         |
         v
Step 2: Content Selection
         |
         v
Step 3: Store Targeting
         |
         v
Step 4: Schedule & Budget
         |
         v
Step 5: Review & Submit
         |
         +---> Auto-approve path (if eligible)
         |           |
         v           v
  Pending Approval   Scheduled
```

### Step 1: Basic Information Validation

| Field | Rule |
|-------|------|
| Name | 3-100 characters, unique per advertiser |
| Brand name | 2-50 characters, required for blocking rules |
| Category | Must select from predefined list |
| Description | Maximum 500 characters |

### Step 2: Content Selection Rules

- Minimum 1 content asset
- Maximum 10 content assets
- All assets must be owned by advertiser
- All assets must have APPROVED status
- Video duration: 10-60 seconds
- Image display duration: 10 seconds (default)
- Assets must meet minimum resolution (1920×1080)

### Step 3: Store Targeting Flow

```
Selected stores: 100
         |
         v
Apply blocking rules
         |
         v
Eligible stores: 85
Blocked stores: 15 (with reasons shown)
```

**Selection Options**:

| Method | Description |
|--------|-------------|
| Manual Selection | Pick individual stores from list |
| Criteria-Based | Select by region, category, foot traffic |

**Minimum Requirements**:
- At least 1 eligible store after blocking rules applied
- Maximum 1,000 stores per campaign

### Step 4: Schedule & Budget Settings

| Setting | Rule |
|---------|------|
| Start date | At least 24 hours from now |
| End date | After start date |
| Maximum duration | 365 days |
| Daily cap | Optional, minimum $10 if set |
| Minimum budget | $100 |
| Maximum budget | $1,000,000 |

### Step 5: Submit and Approval Determination

| Condition | Result |
|-----------|--------|
| Budget > $10,000 | Requires manual approval |
| Content has restricted categories | Requires manual approval |
| Standard campaign | Auto-approved → SCHEDULED |

### Budget Hold Process (on Submit)

```
Advertiser Wallet                    Campaign
+------------------+                +------------------+
| Available: $5000 |                | Budget: $1000    |
| Held: $0         |   Submit       | Remaining: $1000 |
+------------------+ ------------->  +------------------+
         |                                    |
         v                                    v
+------------------+                +------------------+
| Available: $4000 |                | Budget: $1000    |
| Held: $1000      |                | Remaining: $1000 |
+------------------+                +------------------+
```

## Checklist (Subtasks)

- [ ] Create wizard component framework
- [ ] Build Step 1: Objective form
- [ ] Build Step 2: Budget form
- [ ] Build Step 3: Schedule form
- [ ] Build Step 4: Targeting form
- [ ] Build Step 5: Content selection
- [ ] Build Step 6: Review summary
- [ ] Implement step validation
- [ ] Add save as draft feature
- [ ] Implement auto-save
- [ ] Create reach estimation service
- [ ] Build progress indicator
- [ ] Add back/next navigation
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
