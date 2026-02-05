---
id: "STORY-1.1"
epic_id: "EPIC-01"
title: "User Registration & Email Verification"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["registration", "verification", "onboarding"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# User Registration & Email Verification

## User Story

**As a** new user,
**I want** to register an account with my email and password,
**So that** I can access the RMN platform.

## Acceptance Criteria

- [ ] User can enter email and password to create account
- [ ] System validates email uniqueness (one account per email)
- [ ] System validates password requirements (12+ chars, uppercase + lowercase + number + symbol)
- [ ] System sends verification email with link
- [ ] Verification link expires in 24 hours
- [ ] Maximum 3 verification emails per hour (resend limit)
- [ ] Account status starts as UNVERIFIED
- [ ] Account becomes ACTIVE after email verification
- [ ] Account expires if not verified within 24 hours
- [ ] All failed registration attempts are logged for security audit

## Business Flow

### Self-Registration Flow
```
User enters email + password
          |
          v
  System validates email uniqueness
          |
          v
  System validates password requirements
          |
          v
  System sends verification email
          |
          v
  User clicks verification link
          |
          v
  Email verified --> Status = ACTIVE
```

### Account Status After Registration
```
+-------------+     Email      +--------+
| UNVERIFIED  |--------------->| ACTIVE |
+-------------+    Verified    +--------+
      |
      v
  (24h timeout)
      |
      v
 Account Expires
```

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
