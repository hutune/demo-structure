---
id: "STORY-5.1"
epic_id: "EPIC-05"
title: "Campaign Types & Structure"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["campaign", "types", "structure"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Campaign Types & Structure

## User Story

**As an** advertiser,
**I want** to understand the different campaign types available,
**So that** I can choose the right format for my advertising goals.

## Acceptance Criteria

### Campaign Types
- [ ] **Brand Awareness**: Maximize impressions across locations
- [ ] **Engagement**: Optimize for interaction (touch, QR scans)
- [ ] **Traffic**: Drive foot traffic to specific locations
- [ ] **Sales Lift**: Measure impact on sales at retail locations

### Campaign Structure
```
Campaign
├── Objectives (goal and KPIs)
├── Budget (total and daily limits)
├── Schedule (start, end, dayparting)
├── Targeting (locations, devices, audience)
└── Ad Groups
    ├── Ad Group 1
    │   ├── Targeting refinements
    │   └── Creative assets
    └── Ad Group 2
        ├── Targeting refinements
        └── Creative assets
```

### Campaign Properties
- [ ] Name and description
- [ ] Campaign type / objective
- [ ] Start and end dates
- [ ] Total budget
- [ ] Daily budget cap
- [ ] Billing model (CPM, CPC, CPV, flat rate)

### Ad Group Properties
- [ ] Name
- [ ] Subset targeting (optional)
- [ ] Creative rotation (sequential, random, weighted)
- [ ] Assigned content assets

### Content per Ad Group
- [ ] Multiple creatives per ad group
- [ ] A/B testing capability
- [ ] Performance tracking per creative

## Business Flow

### Campaign Information Requirements

| Information | Required | Description |
|-------------|----------|-------------|
| Name | Yes | Campaign identifier, unique per advertiser |
| Description | No | Optional notes about the campaign |
| Brand Name | Yes | For competitor blocking matching |
| Category | Yes | Industry category |
| Budget | Yes | Total budget for the campaign |
| Start Date | Yes | When campaign begins |
| End Date | Yes | When campaign ends |
| Target Stores | Yes | Where ads will show |
| Content | Yes | What content will display |
| Priority | Yes | Ad slot preference (1-10) |
| Daily Cap | No | Optional daily budget limit |

### Campaign Categories

- Food & Beverage
- Electronics
- Fashion & Apparel
- Health & Beauty
- Home & Garden
- Automotive
- Entertainment
- Financial Services
- Telecom
- Other

### Campaign Priority Scale

| Priority | Level | Default Budget Range |
|----------|-------|---------------------|
| 1-3 | Low | Under $500 |
| 4-6 | Normal | $500 - $2,000 |
| 7-8 | High | $2,000 - $10,000 |
| 9-10 | Premium | Over $10,000 |

> Advertisers can manually adjust priority ±2 levels from default.

### How Priority Affects Ad Serving

```
Higher priority = More frequent ad slot allocation

When device requests next ad:
    1. Get all eligible active campaigns
    2. Calculate weight: Priority × Remaining Budget Ratio
    3. Select campaign using weighted random
```

**Example**:
- Campaign A: Priority 10, 90% budget left → Weight 9.0
- Campaign B: Priority 7, 50% budget left → Weight 3.5
- Campaign C: Priority 5, 100% budget left → Weight 5.0

Campaign A has highest chance of being selected.

## Checklist (Subtasks)

- [ ] Define campaign type configurations
- [ ] Create campaign data model
- [ ] Create ad group data model
- [ ] Implement campaign-ad group relationships
- [ ] Build campaign type selection UI
- [ ] Create structure visualization
- [ ] Add campaign type descriptions
- [ ] Implement billing model options
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
