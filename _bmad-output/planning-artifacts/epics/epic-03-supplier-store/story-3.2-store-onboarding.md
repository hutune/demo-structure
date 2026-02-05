---
id: "STORY-3.2"
epic_id: "EPIC-03"
title: "Store Onboarding & Configuration"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["store", "onboarding", "configuration"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Store Onboarding & Configuration

## User Story

**As a** supplier,
**I want** to add and configure my store locations,
**So that** advertisers can target my stores for their campaigns.

## Acceptance Criteria

### Store Information
- [ ] Store name
- [ ] Full address with geocoding
- [ ] Store category (grocery, convenience, pharmacy, etc.)
- [ ] Store size (small, medium, large)
- [ ] Operating hours for each day
- [ ] Timezone (auto-detected from location)
- [ ] Photos of store interior/exterior

### Location Configuration
- [ ] Address validation and geocoding
- [ ] Map pin placement for precise location
- [ ] Service area definition (radius or polygon)
- [ ] Nearby points of interest

### Operating Hours
- [ ] Set hours for each day of week
- [ ] Support for 24-hour operation
- [ ] Holiday schedule overrides
- [ ] Special event hours

### Store Categories
- [ ] Grocery / Supermarket
- [ ] Convenience Store
- [ ] Pharmacy / Drugstore
- [ ] Department Store
- [ ] Shopping Mall
- [ ] Gas Station
- [ ] Restaurant / Café
- [ ] Other Retail

### Store Status
- [ ] Draft: Configuration in progress
- [ ] Pending: Awaiting verification
- [ ] Active: Operational and available for campaigns
- [ ] Paused: Temporarily unavailable
- [ ] Closed: Permanently closed

## Business Flow

### Store Verification Flow

```
Automated verification:
      |
      +---> Address geocoding (get coordinates)
      |
      +---> Commercial address check (not residential)
      |
      +---> Business presence verification (map imagery)
      |
      v
If all pass --> Store verified

If any fail --> Manual verification required:
      |
      +---> Supplier uploads storefront photo
      |
      +---> Supplier uploads business license for location
      |
      +---> Review within 3 business days
```

### Store Status Definitions

| Status | Description | Devices Serve Ads |
|--------|-------------|-------------------|
| Active | Operational | Yes |
| Inactive | Temporarily paused | No |
| Suspended | Platform-initiated block | No |
| Closed | Permanently closed | No |

### Device Limits per Store

| Square Footage | Maximum Devices |
|----------------|-----------------|
| Under 1,000 sq ft | 1 device |
| 1,000 - 2,999 sq ft | 2 devices |
| 3,000 - 4,999 sq ft | 3 devices |
| 5,000 - 9,999 sq ft | 5 devices |
| 10,000+ sq ft | 10 devices |

### Store Suspension Flow

```
Store suspension triggers:
      |
      +---> Quality score below 50 for 7+ consecutive days
      +---> Multiple fraud reports
      +---> Content policy violations
      +---> Device tampering or unauthorized relocation
      |
      v
During suspension:
      +---> No ads served
      +---> Revenue held in pending state
      +---> Supplier notified with reason
      +---> 14-day window to resolve issues
      |
      v
Resolution:
      +---> If resolved: Reactivate, release held revenue
      +---> If unresolved: Store permanently deactivated
```

### Operating Hours Rules

| Rule | Description |
|------|-------------|
| Ads during hours | Ads only served during operating hours |
| Standby mode | Devices enter standby outside hours |
| Impressions | No impressions counted when store is closed |
| 24/7 stores | Set all days to full 24 hours |
| Holiday hours | Can set special hours for specific dates |

## Checklist (Subtasks)

- [ ] Create store listing page for supplier
- [ ] Build store creation wizard
- [ ] Implement address validation with geocoding
- [ ] Add map-based location picker
- [ ] Create operating hours editor
- [ ] Implement timezone detection
- [ ] Add store category selection
- [ ] Create photo upload feature
- [ ] Implement store status management
- [ ] Build store verification queue
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
