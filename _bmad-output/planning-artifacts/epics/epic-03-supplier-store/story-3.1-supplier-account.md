---
id: "STORY-3.1"
epic_id: "EPIC-03"
title: "Supplier Account Management"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["supplier", "account", "onboarding"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Supplier Account Management

## User Story

**As a** retail business owner,
**I want** to create a supplier account and manage my business details,
**So that** I can monetize my store locations with digital advertising.

## Acceptance Criteria

### Supplier Registration
- [ ] Register with email and password
- [ ] Verify email address
- [ ] Provide company name and legal entity details
- [ ] Provide tax ID for revenue reporting
- [ ] Add banking details for payouts
- [ ] Upload business documents for verification

### Required Company Information
- [ ] Legal company name
- [ ] Trading name (if different)
- [ ] Business registration number
- [ ] Tax identification number
- [ ] Registered address
- [ ] Primary contact details

### Banking Details for Payouts
- [ ] Bank name
- [ ] Account holder name
- [ ] Account number
- [ ] Routing number / SWIFT code
- [ ] Currency preference

### Account Verification
- [ ] Business documents reviewed by admin
- [ ] Banking details verified with micro-deposit
- [ ] Status: Pending → Verified → Active

### Account Statuses
- [ ] Pending: Registration submitted, under review
- [ ] Active: Verified and operational
- [ ] Suspended: Temporarily disabled
- [ ] Terminated: Permanently closed

## Business Flow

### Supplier Registration Flow

```
Step 1: Create user account
              |
              v
Step 2: Enter business information
              |
              v
Step 3: Upload verification documents
              |
              v
Step 4: Register first store
              |
              v
Step 5: Register first device
              |
              v
Step 6: Platform reviews application
              |
              v
Step 7: Approval --> Status = ACTIVE
```

### Supplier Status Transitions

```
+----------------------+     Verified     +--------+
| PENDING_REGISTRATION |----------------->| ACTIVE |
+----------------------+                  +--------+
                                              |
                              +---------------+---------------+
                              |               |               |
                              v               v               v
                        +----------+   +-----------+   +------------------+
                        | INACTIVE |   | SUSPENDED |   | PERMANENTLY_BANNED |
                        +----------+   +-----------+   +------------------+
                        (voluntary)    (platform)      (severe violation)
```

### Status Definitions

| Status | Description | Can Earn Revenue |
|--------|-------------|------------------|
| Pending Registration | Registration submitted, awaiting verification | No |
| Active | Fully operational | Yes |
| Inactive | Voluntarily paused by supplier | No |
| Suspended | Temporarily blocked by platform | No (held) |
| Permanently Banned | Permanent termination | No (forfeited) |

### Required Documents by Business Type

| Business Type | Required Documents |
|--------------|-------------------|
| Individual | Government ID, Proof of address, Bank statement |
| Sole Proprietorship | Business license, Tax ID letter, Bank statement |
| LLC | Articles of Organization, Tax ID letter, Bank statement, Operating Agreement |
| Corporation | Certificate of Incorporation, Tax ID letter, Bank statement, Bylaws |
| Partnership | Partnership Agreement, Tax ID letter, Bank statement |

### Revenue Share Model

| Party | Share |
|-------|-------|
| Supplier | 80% of impression revenue |
| Platform | 20% of impression revenue |

## Checklist (Subtasks)

- [ ] Create supplier registration flow
- [ ] Build company information form
- [ ] Implement tax ID validation
- [ ] Add banking details form
- [ ] Implement document upload
- [ ] Create admin verification queue
- [ ] Add micro-deposit verification
- [ ] Implement account status management
- [ ] Create supplier dashboard
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
