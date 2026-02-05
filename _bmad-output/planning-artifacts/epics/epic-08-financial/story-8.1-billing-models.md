---
id: "STORY-8.1"
epic_id: "EPIC-08"
title: "Billing Models"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["billing", "pricing", "cpm"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Billing Models

## User Story

**As an** advertiser,
**I want** to choose how I'm billed for advertising,
**So that** I can align costs with my campaign objectives.

## Acceptance Criteria

### Billing Model Options
| Model | Full Name | Description |
|-------|-----------|-------------|
| CPM | Cost Per Mille | Pay per 1,000 impressions |
| CPC | Cost Per Click | Pay per interaction/tap |
| CPV | Cost Per View | Pay per completed video view |
| Flat | Flat Rate | Fixed daily/weekly fee |

### CPM (Cost Per Mille)
- [ ] Charged per 1,000 verified impressions
- [ ] Industry-standard formula: (Impressions / 1000) × CPM rate
- [ ] Rate varies by location, time, device type
- [ ] Minimum bid applies

### CPC (Cost Per Click)
- [ ] Charged per verified interaction
- [ ] Requires touch-enabled devices
- [ ] Click tracking with anti-fraud
- [ ] Minimum bid applies

### CPV (Cost Per View)
- [ ] Charged per completed video view
- [ ] "Completed" = 75%+ of video duration
- [ ] Partial views not charged
- [ ] Rate based on video length

### Flat Rate
- [ ] Fixed price for guaranteed placement
- [ ] Daily, weekly, or campaign-long
- [ ] Specific devices/locations reserved
- [ ] Premium pricing

### Rate Cards
- [ ] Dynamic rates by supply/demand
- [ ] Location-based pricing tiers
- [ ] Time-of-day multipliers
- [ ] Volume discounts

## Business Flow

### CPM Calculation Flow

```
Advertiser selects campaign parameters
       |
       v
System determines applicable rates:
- Base CPM rate
- Time-slot multiplier
- Venue-type multiplier
- Location premium
       |
       v
Calculate effective CPM
       |
       v
Estimate total impressions from budget
       |
       v
Display cost preview to advertiser
```

### Cost Formula

```
Cost = (Number of Impressions / 1000) x CPM Rate

Example:
- CPM Rate: $5.00
- Total Impressions: 10,000
- Total Cost: (10,000 / 1000) x $5.00 = $50.00
```

### Time-Slot Pricing Multipliers

| Time Slot            | Multiplier | Effective CPM (if base = $5) |
|----------------------|------------|------------------------------|
| Evening (18:00-21:00)| +70%       | $8.50                        |
| Lunch (11:00-13:00)  | +30%       | $6.50                        |
| Morning (09:00-11:00)| +0%        | $5.00                        |
| Early (06:00-09:00)  | -50%       | $2.50                        |
| Night (21:00-06:00)  | -70%       | $1.50                        |

### Venue-Type Pricing Multipliers

| Store Type             | Multiplier | Effective CPM (if base = $5) |
|------------------------|------------|------------------------------|
| Flagship Store         | +200%      | $15.00                       |
| Large Shopping Mall    | +100%      | $10.00                       |
| Large Supermarket      | +50%       | $7.50                        |
| Convenience Store      | +0%        | $5.00 (base)                 |
| Small Grocery          | -30%       | $3.50                        |

### Location Premium Within Store

```
[Entrance]
    |
    v
+----------+
| Entry    | <-- CPM +50% (high visibility, all customers see)
| Screen   |
+----------+
    |
    v
+------------------+
|  Shopping Area   |
|  +------+        |
|  |Screen|        | <-- Base CPM (normal foot traffic)
|  +------+        |
+------------------+
    |
    v
+----------+
| Checkout | <-- CPM +100% (longest dwell time, highest attention)
| Screen   |
+----------+
```

### Playtime Guarantee Rules

| Rule | Description |
|------|-------------|
| Package Definition | Guaranteed number of impressions within time period |
| Shortfall Handling | Extend campaign OR refund proportional amount |
| Verification | Guarantee applies to verified impressions only |
| Exclusions | Device downtime excluded from guarantee calculation |

## Checklist (Subtasks)

- [ ] Design billing model configuration
- [ ] Implement CPM calculation
- [ ] Implement CPC tracking and billing
- [ ] Implement CPV tracking and billing
- [ ] Create flat rate reservations
- [ ] Build rate card management
- [ ] Add location-based pricing
- [ ] Implement time multipliers
- [ ] Add volume discounts
- [ ] Create billing preview
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
