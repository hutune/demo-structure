---
id: "STORY-1.6"
epic_id: "EPIC-01"
title: "Device & Session Management"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["device", "session", "security", "trusted-device"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Device & Session Management

## User Story

**As a** user,
**I want** to view and manage all devices that are logged into my account,
**So that** I can ensure my account security and remove unauthorized access.

## Acceptance Criteria

### View Active Sessions
- [ ] User can see list of all active sessions/devices
- [ ] Each session shows: device name, browser, OS, location (city/country), last active time
- [ ] Current session is clearly marked
- [ ] Sessions sorted by last active (most recent first)

### Session Details Display
| Information | Example |
|-------------|---------|
| Device Type | Mobile, Desktop, Tablet |
| Device Name | iPhone 15 Pro, MacBook Pro |
| Browser | Chrome 120, Safari 17 |
| Operating System | iOS 17, macOS Sonoma |
| Location | Ho Chi Minh City, Vietnam |
| IP Address | Partially masked (192.168.xxx.xxx) |
| Last Active | 5 minutes ago, Yesterday at 3:00 PM |
| Login Time | Feb 5, 2026 at 10:30 AM |

### Remote Logout
- [ ] User can logout from any specific session (except current)
- [ ] User can "Logout from all other devices" (one-click)
- [ ] Logged out session is terminated immediately
- [ ] User receives notification when remotely logged out
- [ ] Confirmation required before mass logout

### Trusted Devices
- [ ] User can mark a device as "Trusted"
- [ ] Trusted devices skip 2FA for 30 days
- [ ] Maximum 5 trusted devices per account
- [ ] User can remove trusted status anytime
- [ ] Trusted device automatically removed after 90 days of inactivity

### Device Recognition
- [ ] System creates device fingerprint on first login
- [ ] Same device recognized on subsequent logins
- [ ] "New device" alert for unrecognized devices
- [ ] Device history retained for 1 year

### Session Limits
| Account Type | Max Concurrent Sessions |
|--------------|------------------------|
| Free | 3 |
| Basic | 5 |
| Premium | 10 |
| Enterprise | Unlimited |

- [ ] Oldest session auto-terminated when limit reached
- [ ] User notified when session terminated due to limit

## Business Flow

### View & Manage Sessions Flow
```
User opens "Security Settings"
          |
          v
Navigate to "Active Sessions"
          |
          v
System displays all active sessions:
+------------------------------------------+
| Current Session (This device)            |
| Chrome on MacBook Pro                    |
| Ho Chi Minh City • Active now            |
+------------------------------------------+
| iPhone 15 Pro                            |
| Safari on iOS 17                         |
| Hanoi • 2 hours ago         [Logout] [⭐]|
+------------------------------------------+
| Windows PC                               |
| Firefox on Windows 11                    |
| Da Nang • Yesterday         [Logout] [⭐]|
+------------------------------------------+
          |
          v
[Logout from all other devices]
```

### Remote Logout Flow
```
User clicks "Logout" on specific session
          |
          v
Confirm: "Logout from iPhone 15 Pro?"
          |
     +----+----+
     |         |
  Cancel    Confirm
     |         |
     v         v
  Close    Terminate session
  dialog   immediately
               |
               v
        Refresh session list
               |
               v
        Notify other device:
        "You have been logged out"
```

### Trusted Device Flow
```
User logs in with 2FA
          |
          v
2FA verified successfully
          |
          v
Prompt: "Trust this device?"
    "Skip 2FA for 30 days on this device"
          |
     +----+----+
     |         |
    No        Yes
     |         |
     v         v
  Normal    Add to trusted
  login     devices list
               |
               v
        Check: < 5 trusted devices?
               |
          +----+----+
          |         |
         Yes        No
          |         |
          v         v
        Save     Prompt: "Remove
        device   oldest trusted
                 device?"
```

### New Device Detection Flow
```
User attempts login
          |
          v
System checks device fingerprint
          |
          v
Device recognized?
          |
     +----+----+
     |         |
    Yes        No (New Device)
     |         |
     v         v
  Normal    Send alert email:
  login     "New login detected"
               |
               v
        Show in-app notification
               |
               v
        Prompt for 2FA (even if trusted
        devices exist)
               |
               v
        Offer to trust new device
```

### Session Limit Enforcement
```
User logs in (new session)
          |
          v
Count current active sessions
          |
          v
Sessions >= limit?
          |
     +----+----+
     |         |
    No        Yes
     |         |
     v         v
  Create   Find oldest session
  session       |
               v
        Terminate oldest session
               |
               v
        Notify user on that device:
        "Session ended - logged in elsewhere"
               |
               v
        Create new session
```

### Device Information Collected

| Data Point | Purpose | Retention |
|------------|---------|-----------|
| Device Type | Display & recognition | 1 year |
| Browser/OS | Display & recognition | 1 year |
| IP Address | Location & security | 90 days |
| Login Time | Audit trail | 1 year |
| Last Active | Session management | Real-time |
| Device Fingerprint | Recognition | 1 year |

### Trusted Device Rules

| Rule | Value |
|------|-------|
| Max trusted devices | 5 per account |
| Trust duration | 30 days (renewable on use) |
| Auto-remove after inactivity | 90 days |
| Removed on password change | Yes |
| Removed on 2FA disable | Yes |

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
