---
id: "STORY-5.7"
epic_id: "EPIC-05"
title: "Targeting Configuration"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["targeting", "audience", "location"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Targeting Configuration

## User Story

**As an** advertiser,
**I want** to target specific locations, devices, and audiences,
**So that** my ads reach the most relevant viewers.

## Acceptance Criteria

### Geographic Targeting
- [ ] By country
- [ ] By state/province
- [ ] By city
- [ ] By postal/zip code
- [ ] By radius around a point
- [ ] By specific stores

### Store Targeting
- [ ] All available stores (default)
- [ ] Specific supplier networks
- [ ] Store categories (grocery, convenience, etc.)
- [ ] Individual store selection
- [ ] Exclude specific stores

### Device Targeting
- [ ] All device types (default)
- [ ] Specific device types (large display, kiosk, etc.)
- [ ] Device placement (entrance, checkout, aisle)
- [ ] Exclude specific devices

### Audience Targeting (if available)
- [ ] Demographics (age, gender - if measurable)
- [ ] Shopping behavior segments
- [ ] Time-based (day vs night shoppers)
- [ ] First-time vs returning visitors

### Targeting Reach Estimation
- [ ] Show estimated devices reached
- [ ] Show estimated daily impressions
- [ ] Update in real-time as targeting changes
- [ ] Warn if reach is very low

### Competitor Blocking Awareness
- [ ] Show blocked brands for selected stores
- [ ] Indicate reduced reach due to blocking
- [ ] Suggest alternative stores

## Business Flow

### How Store Targeting Works

```
Campaign selects target stores
         |
         v
System applies supplier blocking rules
         |
         v
Result: Eligible stores (can show ads)
        Blocked stores (cannot show ads)
         |
         v
At activation:
    Content distributed to eligible store devices only
```

### Blocking Rules

Suppliers can block campaigns based on:
- Brand name match
- Industry category match
- Keyword match

**Campaign Impact**:
- Blocked stores never receive campaign content
- Advertiser sees blocked stores and reasons
- Budget allocated only for eligible stores

### Store Selection Flow

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

**Targeting Criteria** (if criteria-based):
- Regions/Areas
- Store categories (Supermarket, Mall, etc.)
- Minimum foot traffic
- Distance from location

**Minimum Requirements**:
- At least 1 eligible store after blocking rules applied
- Maximum 1,000 stores per campaign

### Dynamic Blocking Changes

If supplier adds blocking rule after campaign starts:

1. Campaign immediately removed from that store
2. Store moved from eligible to blocked
3. Advertiser notified: "Store X now blocks your campaign"
4. Content removed from store devices
5. No impressions counted from that point

### Estimation Accuracy

System estimates impressions based on:
- Store foot traffic
- Number of devices per store
- Average visitor dwell time
- Campaign duration
- Historical performance

Estimate includes 30% conservative buffer for accuracy.

## Checklist (Subtasks)

- [ ] Build geographic targeting UI
- [ ] Implement map-based location picker
- [ ] Add radius targeting
- [ ] Create store selection interface
- [ ] Implement supplier network targeting
- [ ] Add store category filtering
- [ ] Build device type selection
- [ ] Add device placement targeting
- [ ] Create audience segment picker
- [ ] Implement reach estimation
- [ ] Add competitor block warnings
- [ ] Create targeting summary display
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
