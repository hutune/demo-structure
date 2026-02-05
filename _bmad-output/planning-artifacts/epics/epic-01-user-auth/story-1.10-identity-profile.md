---
id: "STORY-1.10"
epic_id: "EPIC-01"
title: "Profile Management"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["profile", "settings", "gdpr", "personal-data"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Profile Management

## User Story

**As a** user,
**I want** to view and update my personal profile information,
**So that** I can keep my account details accurate and manage my data.

## Acceptance Criteria

### Profile Information

| Field | Editable | Validation |
|-------|----------|------------|
| Email | Yes (with re-verification) | Valid email format |
| Display Name | Yes | 2-50 characters |
| Profile Picture | Yes | Image file, max 5MB |
| Phone Number | Yes | Valid phone format |
| Timezone | Yes | Standard timezone list |
| Language | Yes | Supported languages |
| Date Format | Yes | DD/MM/YYYY or MM/DD/YYYY |

### Display Name
- [ ] User can update display name anytime
- [ ] 2-50 characters allowed
- [ ] Shown throughout the platform
- [ ] Changes take effect immediately

### Profile Picture
- [ ] User can upload profile picture
- [ ] Supported formats: JPG, PNG, GIF
- [ ] Maximum file size: 5MB
- [ ] Auto-cropped to square
- [ ] Can remove picture (revert to default avatar)

### Email Change
- [ ] User can request email change
- [ ] Must verify current password first
- [ ] Verification email sent to NEW email address
- [ ] Change not effective until new email verified
- [ ] Notification sent to OLD email about the change
- [ ] 24-hour window to cancel change (from old email)

### Phone Number
- [ ] User can add/update phone number
- [ ] Phone verification via SMS OTP
- [ ] Verified phone can be used for 2FA backup
- [ ] User can remove phone (if not used for 2FA)

### Preferences
- [ ] Timezone selection (for displaying dates/times)
- [ ] Language preference (UI language)
- [ ] Date format preference
- [ ] Notification preferences (link to Story 1.7)

### Activity History
- [ ] User can view recent account activity
- [ ] Shows: logins, password changes, security events
- [ ] Last 90 days of activity
- [ ] Exportable as CSV

### Data Export (GDPR)
- [ ] User can request full data export
- [ ] Export includes all personal data stored
- [ ] Generated within 24 hours
- [ ] Download link sent via email
- [ ] Link valid for 7 days
- [ ] Format: JSON or CSV (user choice)

### Account Deletion Request
- [ ] User can request account deletion
- [ ] Must verify password to confirm
- [ ] 30-day grace period before deletion
- [ ] Can cancel deletion within grace period
- [ ] After 30 days: permanent deletion
- [ ] Confirmation email sent at each stage

## Business Flow

### Email Change Flow
```
User requests email change
          |
          v
Enter new email address
          |
          v
Verify current password
          |
     +----+----+
     |         |
  Invalid    Valid
     |         |
     v         v
  Error    Send verification to
           NEW email
               |
               v
        Send notification to
        OLD email (with cancel link)
               |
               v
        User clicks verify link
        in NEW email
               |
               v
        Email changed successfully
               |
               v
        Notify OLD email:
        "Email change completed"
```

### Email Change Cancellation
```
Notification sent to old email:
"Someone requested to change your email"
          |
          v
User clicks "Cancel this change"
          |
          v
Verify it's within 24 hours
          |
     +----+----+
     |         |
  Expired   Valid
     |         |
     v         v
  Error:   Cancel email change
  "Too     request
   late"        |
               v
        Notify user:
        "Email change cancelled"
               |
               v
        Log security event
```

### Profile Picture Upload Flow
```
User selects image file
          |
          v
Validate file:
  - Format: JPG/PNG/GIF?
  - Size: <= 5MB?
          |
     +----+----+
     |         |
   Fail      Pass
     |         |
     v         v
  Error:    Show crop preview
  "Invalid     |
   file"       v
        User adjusts crop area
               |
               v
        [Save] / [Cancel]
               |
        +------+------+
        |             |
     Cancel        Save
        |             |
        v             v
     Discard     Upload & process
                      |
                      v
                 Update profile
                      |
                      v
                 Show new picture
```

### Data Export Flow (GDPR)
```
User requests data export
          |
          v
Select format: JSON or CSV
          |
          v
Verify password
          |
          v
Request queued for processing
          |
          v
Show message: "Your data export
is being prepared. You'll receive
an email within 24 hours."
          |
          v
Background job collects all data:
  - Profile information
  - Activity history
  - Campaign data
  - Content uploads
  - Financial transactions
  - Preferences
          |
          v
Generate export file
          |
          v
Send email with download link
(valid for 7 days)
          |
          v
User downloads data
```

### Account Deletion Flow
```
User requests account deletion
          |
          v
Show warning:
"This will permanently delete:
 - All your campaigns
 - All your content
 - All your data"
          |
          v
Confirm: Enter password
          |
          v
Check: Active subscriptions?
       Positive wallet balance?
          |
     +----+----+
     |         |
    Yes        No
     |         |
     v         v
  Block:    Schedule deletion
  "Please   (30 days from now)
   cancel        |
   subs first"   v
           Send confirmation email
                 |
                 v
           Account status = PENDING_DELETION
                 |
                 v
           Send reminder at Day 25
                 |
                 v
           Day 30: Permanent deletion
                 |
                 v
           Final confirmation email
```

### Profile Data Categories

| Category | Includes |
|----------|----------|
| **Personal Info** | Name, email, phone, picture |
| **Preferences** | Timezone, language, notifications |
| **Security** | 2FA status, trusted devices, login history |
| **Activity** | Account events, security events |
| **Business Data** | Campaigns, content, transactions (if applicable) |

### Data Retention Rules

| Data Type | Retention Period | After Deletion |
|-----------|------------------|----------------|
| Profile info | Until account deleted | 30 days grace |
| Activity log | 90 days visible | 1 year archived |
| Security events | 1 year | 7 years (compliance) |
| Export files | 7 days download | Auto-deleted |
| Deleted account data | 30 days recoverable | Permanent delete |

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
