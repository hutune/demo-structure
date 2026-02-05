---
id: "STORY-2.5"
epic_id: "EPIC-02"
title: "Account Status & Suspension"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["account", "suspension", "status"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Account Status & Suspension

## User Story

**As a** platform administrator,
**I want** to manage advertiser account statuses,
**So that** I can enforce policies and handle issues.

## Acceptance Criteria

### Account Statuses
- [ ] Active: Normal operation
- [ ] Suspended: Temporarily disabled
- [ ] Banned: Permanently disabled
- [ ] Closed: User-initiated closure

### Status Capabilities
| Status | View Reports | Create Campaigns | Run Campaigns | Top Up Wallet |
|--------|--------------|------------------|---------------|---------------|
| Active | ✓ | ✓ | ✓ | ✓ |
| Suspended | ✓ | - | - | - |
| Banned | - | - | - | - |
| Closed | - | - | - | - |

### Suspension Reasons
- [ ] Violation of content policies (inappropriate ads)
- [ ] Payment issues (chargebacks, declined payments)
- [ ] Suspicious activity flagged
- [ ] User-requested temporary pause

### Suspension Effects
- [ ] All active campaigns pause immediately
- [ ] Cannot create new campaigns
- [ ] Cannot top up wallet
- [ ] Can view reports and data
- [ ] Team members lose access

### Reactivation
- [ ] Admin reviews and resolves issue
- [ ] Account reactivated
- [ ] Campaigns can be resumed
- [ ] Notification sent to account owner

## Business Flow

### Account Status Flow

```
                    ACTIVE
                      │
         ┌────────────┼────────────┐
         v            v            v
    Policy        Payment      User
    Violation     Issues       Request
         │            │            │
         └────────────┼────────────┘
                      v
                 SUSPENDED
                      │
         ┌────────────┼────────────┐
         v            v            v
    Issue         3+ Repeat    User Closes
    Resolved      Violations   Account
         │            │            │
         v            v            v
      ACTIVE       BANNED       CLOSED
```

### Suspension Flow Detail

```
    Violation Detected
           │
           v
    ┌──────────────────┐
    │ Account Suspended │
    └────────┬─────────┘
             │
             v
    Immediate Effects:
    ├── All campaigns pause
    ├── Cannot create new campaigns
    ├── Cannot top up wallet
    ├── Team members lose access
    └── Can still view reports
             │
             v
    ┌──────────────────┐
    │ Resolve Issue    │
    │ (e.g., pay debt) │
    └────────┬─────────┘
             │
             v
    Submit Appeal
             │
             v
    Admin Review (3 business days)
             │
      ┌──────┴──────┐
      v             v
   Approved      Denied
      │             │
      v             v
   ACTIVE      Remains SUSPENDED
                or BANNED
```

### Account Closure Flow (User-Initiated)

```
    User Requests Closure
           │
           v
    ┌──────────────────────────┐
    │ Pre-closure Requirements │
    ├──────────────────────────┤
    │ • Settle all payments    │
    │ • Complete/cancel        │
    │   active campaigns       │
    └────────────┬─────────────┘
                 │
                 v
        30-Day Cooling Off Period
        (Can cancel closure)
                 │
                 v
        Account Permanently CLOSED
                 │
                 v
        Data Retention:
        ├── Financial records: 7 years
        ├── Campaign data: 90 days
        └── Personal data: Deleted per policy
```

### Status Capabilities Reference

| Status | View Reports | Create Campaigns | Run Campaigns | Top Up Wallet |
|--------|--------------|------------------|---------------|---------------|
| Active | ✓ | ✓ | ✓ | ✓ |
| Suspended | ✓ | ✗ | ✗ | ✗ |
| Banned | ✗ | ✗ | ✗ | ✗ |
| Closed | ✗ | ✗ | ✗ | ✗ |

## Checklist (Subtasks)

- [ ] Create account status management for advertisers
- [ ] Implement suspension with reason
- [ ] Add campaign auto-pause on suspension
- [ ] Create suspension notification
- [ ] Build admin review interface
- [ ] Implement reactivation flow
- [ ] Add ban functionality
- [ ] Create account closure flow
- [ ] Add refund handling for closed accounts
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
