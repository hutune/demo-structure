---
id: "STORY-1.9"
epic_id: "EPIC-01"
title: "Account Status Management"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["account", "status", "suspension", "deletion"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Account Status Management

## User Story

**As a** platform administrator,
**I want** to manage user account statuses,
**So that** I can suspend, reactivate, or close accounts as needed.

## Acceptance Criteria

### Account Statuses

| Status | Description | Can Login | Can Transact |
|--------|-------------|-----------|--------------|
| UNVERIFIED | Created but email not verified | No | No |
| ACTIVE | Normal operation | Yes | Yes |
| SUSPENDED | Temporarily disabled by admin | No | No |
| BANNED | Permanently disabled | No | No |
| CLOSED | User requested closure | No | No |

### Status Transitions
- [ ] UNVERIFIED → ACTIVE (email verified)
- [ ] ACTIVE → SUSPENDED (admin action)
- [ ] SUSPENDED → ACTIVE (admin reactivation)
- [ ] ACTIVE → BANNED (severe policy violation)
- [ ] ACTIVE → CLOSED (user request or admin action)
- [ ] CLOSED → ACTIVE (restore within 30 days)

### Suspension
- [ ] Admin can suspend accounts with reason
- [ ] Reason categories: Policy violation, Payment issues, Suspicious activity, User request
- [ ] Suspended users cannot log in
- [ ] All active campaigns pause immediately
- [ ] Cannot create new campaigns or top up wallet
- [ ] Can view reports and data (read-only via support)
- [ ] Team members lose access
- [ ] Email notification sent on suspension
- [ ] Admin can reactivate suspended accounts

### Account Closure (User-Initiated)
- [ ] User can request account closure from settings
- [ ] Must verify password to confirm
- [ ] 30-day grace period before permanent deletion
- [ ] User can cancel closure within grace period
- [ ] Remaining wallet balance must be withdrawn or forfeited
- [ ] Active subscriptions must be cancelled first
- [ ] All data permanently deleted after 30 days (GDPR)

### Account Ban (Admin-Initiated)
- [ ] Reserved for severe policy violations
- [ ] Immediate effect - no grace period
- [ ] User cannot create new accounts
- [ ] Email/phone blacklisted
- [ ] All data retained for legal purposes
- [ ] Appeals process available

### Soft Delete vs Hard Delete
- [ ] Soft delete: Account marked as CLOSED, data retained 30 days
- [ ] Hard delete: Permanent removal after 30 days
- [ ] Some data retained for legal/compliance (7 years for financial)
- [ ] User notified at each stage

## Business Flow

### User Status Flow
```
         +-------------+
         | UNVERIFIED  |
         +------+------+
     Email      |
     Verified   v
         +-------------+
         |   ACTIVE    |
         +------+------+
                |
     +----------+----------+----------+
     |          |          |          |
     v          v          v          v
+-----------+ +------+ +--------+ +--------+
| SUSPENDED | | CLOSED| | BANNED | | ACTIVE |
+-----------+ +------+ +--------+ +--------+
 (by admin)   (by user) (severe)   (normal)
     |            |
     v            v
  Can be      Restore within
reactivated    30 days
```

### Suspension Flow
```
Admin initiates suspension
          |
          v
Select suspension reason:
  - Policy violation
  - Payment issues
  - Suspicious activity
  - User request
          |
          v
Add notes (optional)
          |
          v
Confirm suspension
          |
          v
System immediately:
  - Sets status = SUSPENDED
  - Pauses all active campaigns
  - Revokes team member access
  - Sends notification to user
          |
          v
User receives email:
"Your account has been suspended"
with reason and appeal info
```

### Account Closure Flow (User-Initiated)
```
User requests account closure
          |
          v
System checks:
  - Active subscriptions?
  - Positive wallet balance?
  - Pending transactions?
          |
     +----+----+
     |         |
   Blockers   Clear
     |         |
     v         v
  Show:    Verify password
  "Please       |
   resolve      v
   first"   Schedule closure
             (30 days)
                 |
                 v
            Status = CLOSED
                 |
                 v
            Send confirmation
            email with
            cancel link
                 |
                 v
            Day 25: Reminder email
                 |
                 v
            Day 30: Permanent deletion
                 |
                 v
            Final confirmation email
```

### Reactivation Flow
```
Admin or User initiates reactivation
          |
          v
Check current status
          |
     +----+----+----+
     |    |         |
SUSPENDED CLOSED   BANNED
     |    |         |
     v    v         v
  Admin  Within   Appeal
  can    30 days? required
reactivate   |
     |   +---+---+
     |   |       |
     |  Yes      No
     |   |       |
     v   v       v
  Status =   Cannot
  ACTIVE   reactivate
     |
     v
  Notify user
  "Your account is active"
```

### Status Capabilities Matrix

| Capability | ACTIVE | SUSPENDED | CLOSED | BANNED |
|------------|--------|-----------|--------|--------|
| Login | ✅ | ❌ | ❌ | ❌ |
| View dashboard | ✅ | ❌ | ❌ | ❌ |
| Run campaigns | ✅ | ❌ | ❌ | ❌ |
| Create campaigns | ✅ | ❌ | ❌ | ❌ |
| Top up wallet | ✅ | ❌ | ❌ | ❌ |
| Withdraw balance | ✅ | Via support | Via support | ❌ |
| View reports | ✅ | Via support | Via support | ❌ |
| Export data | ✅ | Via support | Via support | ❌ |
| Appeal | N/A | ✅ | N/A | ✅ |
| Reactivate | N/A | ✅ (admin) | ✅ (30 days) | ❌ |

### Suspension Reasons

| Reason | Severity | Typical Duration | Auto-Resolve? |
|--------|----------|------------------|---------------|
| Policy violation - Minor | Low | 7 days | After review |
| Policy violation - Major | High | 30+ days | Manual only |
| Payment issues | Medium | Until resolved | Yes, on payment |
| Suspicious activity | High | Until verified | Manual only |
| User request | Low | As requested | User can cancel |

### Data Retention After Closure

| Data Type | Retention | Reason |
|-----------|-----------|--------|
| Account info | 30 days | Grace period |
| Transaction records | 7 years | Tax/legal compliance |
| Campaign history | 30 days | Then anonymized |
| Content uploads | 30 days | Then deleted |
| Personal data | 30 days | GDPR compliance |

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
