---
id: "STORY-1.7"
epic_id: "EPIC-01"
title: "Security Notifications"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["security", "notifications", "alerts", "email"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Security Notifications

## User Story

**As a** user,
**I want** to receive notifications about security-related activities on my account,
**So that** I can detect and respond to unauthorized access quickly.

## Acceptance Criteria

### Notification Triggers
| Event | Email | In-App | Push (Mobile) |
|-------|-------|--------|---------------|
| New device login | ✅ Always | ✅ | ✅ |
| Login from new location | ✅ Always | ✅ | ✅ |
| Password changed | ✅ Always | ✅ | ✅ |
| Password reset requested | ✅ Always | ✅ | ❌ |
| 2FA enabled | ✅ Always | ✅ | ✅ |
| 2FA disabled | ✅ Always | ✅ | ✅ |
| Recovery code used | ✅ Always | ✅ | ✅ |
| Account locked (failed attempts) | ✅ Always | ❌ | ❌ |
| Successful login after lockout | ✅ Always | ✅ | ✅ |
| Session terminated remotely | ❌ | ✅ | ✅ |
| Trusted device added | ✅ | ✅ | ✅ |
| OAuth provider linked | ✅ | ✅ | ✅ |
| OAuth provider unlinked | ✅ | ✅ | ✅ |
| Email address changed | ✅ (both old & new) | ✅ | ✅ |
| Account suspended | ✅ Always | ❌ | ❌ |
| KYC status changed | ✅ | ✅ | ✅ |

### Notification Content
- [ ] Clear description of what happened
- [ ] When it happened (date, time, timezone)
- [ ] Where it happened (device, location, IP)
- [ ] Action buttons: "This was me" / "Secure my account"
- [ ] Link to security settings

### "Secure My Account" Quick Actions
When user clicks "This wasn't me":
- [ ] Force logout from all devices
- [ ] Lock account temporarily (15 minutes)
- [ ] Prompt to change password
- [ ] Review recent activity

### Email Notification Format
```
+--------------------------------------------------+
|  🔔 Security Alert from RMN Platform             |
+--------------------------------------------------+
|                                                  |
|  New login to your account                       |
|                                                  |
|  Device: iPhone 15 Pro                           |
|  Location: Hanoi, Vietnam                        |
|  Time: Feb 5, 2026 at 3:45 PM (GMT+7)           |
|  IP: 103.xxx.xxx.xxx                            |
|                                                  |
|  +----------------+  +---------------------+     |
|  | This was me ✓  |  | Secure my account ⚠ |     |
|  +----------------+  +---------------------+     |
|                                                  |
|  If you don't recognize this activity,          |
|  click "Secure my account" immediately.         |
|                                                  |
+--------------------------------------------------+
```

### Notification Preferences
- [ ] User can customize notification channels (email on/off, push on/off)
- [ ] Critical security events CANNOT be disabled (password change, 2FA change)
- [ ] User can set "quiet hours" for non-critical notifications
- [ ] Digest option: group non-critical alerts into daily summary

### Notification Preference Levels

| Preference Level | Events Included |
|-----------------|-----------------|
| **Critical** (cannot disable) | Password change, 2FA change, Account lock, Email change |
| **High** (default on) | New device, New location, OAuth changes |
| **Medium** (default on) | Trusted device, Remote logout, KYC updates |
| **Low** (default off) | Session expired, Login success |

### Suspicious Activity Detection
- [ ] Multiple failed login attempts (5+ in 5 minutes)
- [ ] Login from unusual location (different country)
- [ ] Login at unusual time (based on user's patterns)
- [ ] Rapid device switching
- [ ] Multiple password reset requests

## Business Flow

### New Device Login Notification Flow
```
User logs in from new device
          |
          v
System detects: Device not recognized
          |
          v
Create security event record
          |
          v
Send notifications (parallel):
     +----+----+----+
     |    |    |    |
     v    v    v    v
  Email  Push  In-App  SMS
  (always) (if enabled) (always) (if 2FA phone)
          |
          v
User receives alert
          |
     +----+----+
     |         |
"This was me" "Secure account"
     |         |
     v         v
  Mark as    Force logout all
  verified   + Lock 15 min
             + Prompt password change
```

### Suspicious Activity Response Flow
```
System detects suspicious pattern
          |
          v
Evaluate risk level:
  - Failed attempts: Low
  - New country: Medium
  - Multiple patterns: High
          |
          v
Risk level determines response:
          |
     +----+----+----+
     |    |         |
    Low  Medium    High
     |    |         |
     v    v         v
  Alert  Alert +   Alert +
  only   extra     require
         verify    re-auth
```

### "Secure My Account" Flow
```
User clicks "Secure my account"
          |
          v
+------------------------------------------+
|  🔒 Secure Your Account                  |
|                                          |
|  We'll help you secure your account:     |
|                                          |
|  ☑ Log out from all devices             |
|  ☑ Lock account for 15 minutes          |
|  ☐ Change password (recommended)         |
|  ☐ Review trusted devices               |
|  ☐ Enable 2FA (if not enabled)          |
|                                          |
|  [Secure Now]                            |
+------------------------------------------+
          |
          v
Execute selected actions
          |
          v
Show confirmation:
"Your account has been secured"
          |
          v
Guide to review security settings
```

### Notification Delivery Priority

| Channel | Delivery Time | Retry Policy |
|---------|---------------|--------------|
| In-App | Immediate | No retry needed |
| Push | < 5 seconds | 3 retries over 1 minute |
| Email | < 1 minute | 5 retries over 1 hour |
| SMS | < 30 seconds | 3 retries over 5 minutes |

### Security Event Log

All security events stored for:
- User self-review: 90 days
- Admin audit: 1 year
- Legal compliance: 7 years (if required)

| Event Data Stored |
|-------------------|
| Event type |
| Timestamp (UTC) |
| Device information |
| IP address |
| Location (city, country) |
| User action taken |
| Resolution status |

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
