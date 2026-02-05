---
id: "STORY-3.4"
epic_id: "EPIC-03"
title: "Competitor Blocking"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["competitor", "blocking", "exclusivity"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Competitor Blocking

## User Story

**As a** supplier,
**I want** to block competitor brands from advertising in my stores,
**So that** I can protect my relationships with key brand partners.

## Acceptance Criteria

### Block Management
- [ ] Add brands to blocked list
- [ ] Search existing advertisers by brand name
- [ ] Add custom brand names not yet on platform
- [ ] Specify blocked product categories
- [ ] Set blocking scope (all stores or specific stores)

### Blocking Levels
- [ ] Brand-level: Block specific advertiser
- [ ] Category-level: Block entire category (e.g., all soft drinks)
- [ ] Product-level: Block specific products

### Blocking Scope
- [ ] All stores (supplier-wide)
- [ ] Specific stores only
- [ ] Store groups / regions

### System Enforcement
- [ ] Blocked brands cannot target supplier's stores
- [ ] Existing campaigns with blocked brands stop at those stores
- [ ] Advertisers see "Store unavailable" with no reason exposed

### Block Management
- [ ] View all current blocks
- [ ] Remove blocks at any time
- [ ] View block history
- [ ] Receive notification when block affects campaigns

## Business Flow

### Blocking Rule Matching Flow

```
Campaign seeks matching devices
         |
         v
For each potential device:
         |
         v
Check supplier's blocking rules
         |
         +---> Advertiser blocked? --> Exclude device
         |
         +---> Industry blocked? --> Exclude device
         |
         +---> Keyword found in content? --> Exclude device
         |
         v
If no blocks --> Include device in campaign
```

### Blocking Rule Types

| Type | Description | Example |
|------|-------------|----------|
| Specific Advertiser | Block by advertiser name | Block "Adidas Inc." |
| Industry Category | Block by industry | Block "Athletic Footwear" |
| Keyword | Block content containing keywords | Block "Nike", "Adidas" |

### Blocking Rule Scope

| Scope | Description |
|-------|-------------|
| All Stores | Rule applies to all supplier's stores |
| Specific Stores | Rule applies only to selected stores |

### Keyword Matching Rules

| Rule | Description |
|------|-------------|
| Case handling | Case-insensitive matching |
| Partial match | Finds "Adidas" in "New Adidas Shoes" |
| What is checked | Campaign name, content file name, advertiser name |

### Rule Priority

If multiple rules conflict, the most restrictive rule wins.

```
Example:
  Rule 1: Block "Athletic Footwear" industry
  Rule 2: Allow advertiser "Nike" (explicit)
  Result: Nike is BLOCKED (industry rule is more restrictive)

Override: Explicit allowlist rules can override blocking rules.
```

### Rule Limits by Tier

| Tier | Max Rules | Max Blocked Advertisers | Max Keywords |
|------|-----------|------------------------|---------------|
| Starter | 5 | 10 | 20 |
| Professional | 20 | 50 | 100 |
| Enterprise | Unlimited | Unlimited | Unlimited |

### Rule Change Effects

| Effect | Description |
|--------|-------------|
| Timing | Changes take effect immediately |
| Current campaigns | NOT interrupted |
| Future matching | Uses new rules |
| Advertiser notification | Affected advertisers notified if campaign becomes ineligible |

## Checklist (Subtasks)

- [ ] Create block management interface
- [ ] Implement brand search functionality
- [ ] Add custom brand entry
- [ ] Build category blocking feature
- [ ] Implement store-level scope selection
- [ ] Create campaign impact scanner
- [ ] Add real-time campaign adjustment
- [ ] Build block history view
- [ ] Create notification for affected advertisers
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
