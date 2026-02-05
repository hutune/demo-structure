---
id: "STORY-3.3"
epic_id: "EPIC-03"
title: "Device-Store Association"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["device", "store", "association"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Device-Store Association

## User Story

**As a** supplier,
**I want** to assign devices to specific stores and locations within stores,
**So that** content is displayed in the right physical locations.

## Acceptance Criteria

### Device Assignment
- [ ] View list of available (unassigned) devices
- [ ] Assign device to a specific store
- [ ] Specify device location within store (entrance, checkout, aisle, etc.)
- [ ] Device inherits store's timezone and operating hours
- [ ] One device can only be assigned to one store at a time

### Device Placement Locations
- [ ] Entrance / Exit
- [ ] Checkout / Point of Sale
- [ ] Aisle / Shelf
- [ ] Customer Service
- [ ] Waiting Area
- [ ] Window Display
- [ ] Other (custom)

### Device Details at Store
- [ ] Installation date
- [ ] Physical location description
- [ ] Mounting type (wall, stand, ceiling)
- [ ] Orientation (portrait, landscape)
- [ ] Photo of installed device

### Device Reassignment
- [ ] Unassign from current store
- [ ] Can reassign to different store
- [ ] History of store assignments tracked

### Rules
- [ ] Device must be provisioned before assignment
- [ ] Store must be Active status
- [ ] Respect store size device limits

## Business Flow

### Device Assignment Flow

```
Device Registration
      |
      v
Device Provisioned (verified)
      |
      v
Check: Store Active? --NO--> Cannot Assign
      |
     YES
      |
      v
Check: Store at device limit? --YES--> Cannot Assign
      |
     NO
      |
      v
Assign Device to Store
      |
      v
Device inherits:
      +---> Store timezone
      +---> Store operating hours
      |
      v
Device starts serving ads
```

### Device Limits by Store Size

| Square Footage | Maximum Devices |
|----------------|-----------------|
| Under 1,000 sq ft | 1 device |
| 1,000 - 2,999 sq ft | 2 devices |
| 3,000 - 4,999 sq ft | 3 devices |
| 5,000 - 9,999 sq ft | 5 devices |
| 10,000+ sq ft | 10 devices |

### Device Reassignment Flow

```
Current Assignment:
Device --> Store A
      |
      v
Unassign from Store A
      |
      v
Device becomes "unassigned"
      |
      v
Assign to Store B (if eligible)
      |
      v
History of assignments tracked
```

### Device Ownership Transfer

```
Step 1: Current supplier initiates transfer
      |
      v
Step 2: New supplier accepts (must have active account)
      |
      v
Step 3: Device re-paired with new account
      |
      v
Step 4: Historical data not accessible to new owner
      |
      v
Step 5: Final payout to original supplier
      |
      v
Step 6: New supplier earns from transfer date
```

### Transfer Restrictions

| Restriction | Description |
|-------------|-------------|
| Active campaigns | Cannot transfer during active campaigns |
| Verification | Both parties must be verified |
| Transfer fee | Platform charges $25 transfer fee |

## Checklist (Subtasks)

- [ ] Create device assignment interface
- [ ] Build store device list view
- [ ] Implement device placement form
- [ ] Add location type selection
- [ ] Create installation photo upload
- [ ] Implement device reassignment flow
- [ ] Add assignment history tracking
- [ ] Enforce store device limits
- [ ] Implement timezone inheritance
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
