---
id: "STORY-5.6"
epic_id: "EPIC-05"
title: "Scheduling & Dayparting"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["scheduling", "dayparting", "time"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Scheduling & Dayparting

## User Story

**As an** advertiser,
**I want** to control when my ads are shown,
**So that** I can reach my audience at the most effective times.

## Acceptance Criteria

### Campaign Schedule
- [ ] Start date and time
- [ ] End date and time (or "run indefinitely")
- [ ] Timezone selection (default: advertiser's timezone)
- [ ] Immediate start option

### Dayparting (Time-of-Day Targeting)
- [ ] Select specific hours of day
- [ ] Select specific days of week
- [ ] Create multiple time blocks
- [ ] Visual schedule builder (grid)

### Daypart Templates
- [ ] **Morning Rush**: 6am - 9am weekdays
- [ ] **Lunch**: 11am - 2pm daily
- [ ] **After Work**: 5pm - 8pm weekdays
- [ ] **Weekend**: Saturday & Sunday all day
- [ ] **Prime Time**: 6pm - 10pm daily
- [ ] **Custom**: User-defined

### Timezone Handling
- [ ] Campaign timezone for scheduling
- [ ] Device timezone for delivery
- [ ] Convert campaign times to device local time
- [ ] Handle devices in multiple timezones

### Schedule Conflicts
- [ ] Check store operating hours
- [ ] Warn if targeting closed stores
- [ ] Auto-adjust to store hours (optional)

## Business Flow

### Scheduling Rules

| Setting | Rule |
|---------|------|
| Minimum Lead Time | 24 hours between submission and start date |
| Maximum Duration | 365 days per campaign |
| Start Date | At least 24 hours from now |
| End Date | Must be after start date |

### Campaign Activation Process

```
At campaign start_date:
         |
         v
Pre-flight checks:
    ✓ Campaign status = SCHEDULED
    ✓ Remaining budget > 0
    ✓ At least 1 device online
    ✓ Content accessible
         |
         v
Distribute content to devices
         |
         v
Update status → ACTIVE
         |
         v
Enable impression tracking
         |
         v
Notify advertiser: "Campaign is now live"
```

### Activation Edge Cases

**All Devices Offline**:
- Campaign still marked ACTIVE
- Content push retried every 5 minutes
- If no device online after 1 hour: Alert advertiser
- No billing until first impression

**Content Distribution Failed**:
- Retry up to 3 times
- If still failed: Status → ERROR
- Urgent notification to admin and advertiser

**Start Date Missed** (system downtime):
- Activate immediately if:
  - Start date passed but end date not reached
  - Budget remaining
- Log late activation for reporting

### Date Change Rules

**Start Date**:
- Can delay in SCHEDULED status (must maintain 24-hour lead time)
- Cannot advance (make earlier) after SCHEDULED

**End Date**:
- Can extend if budget remaining
- Can shorten in SCHEDULED (triggers recalculation)

## Checklist (Subtasks)

- [ ] Create schedule configuration form
- [ ] Build daypart grid component
- [ ] Implement daypart templates
- [ ] Add custom time block creation
- [ ] Implement timezone selection
- [ ] Create timezone conversion logic
- [ ] Add store hours validation
- [ ] Build schedule preview
- [ ] Create schedule summary display
- [ ] Add "run indefinitely" option
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
