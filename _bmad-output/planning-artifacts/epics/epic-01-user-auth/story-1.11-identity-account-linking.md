---
id: "STORY-1.11"
epic_id: "EPIC-01"
title: "Account Linking & Context Switching"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["account", "multi-account", "context", "switching"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Account Linking & Context Switching

## User Story

**As a** user with multiple account types,
**I want** to switch between my Advertiser and Supplier accounts seamlessly,
**So that** I can manage all my business activities from one login.

## Acceptance Criteria

### Multiple Account Types
- [ ] One user (email) can have both Advertiser and Supplier accounts
- [ ] Each account type has separate data and permissions
- [ ] User can add a new account type from settings

### Context Selection
- [ ] When user logs in with multiple accounts, show account selection screen
- [ ] Display account type and company name for each option
- [ ] User must explicitly select context to proceed
- [ ] Last selected context is remembered for next login

### Context Switching
- [ ] User can switch context without re-login
- [ ] Context switch available from header/navigation
- [ ] Switching context loads appropriate dashboard
- [ ] Permissions are recalculated based on new context

### Data Isolation
- [ ] Each context only shows relevant data
- [ ] Actions logged with context identifier
- [ ] Cross-context data access not allowed

## Business Flow

### Login with Multiple Accounts
```
User logs in
      |
      v
System detects multiple accounts
      |
      v
Present account selection screen:
  +-----------------------------+
  | Select Account              |
  +-----------------------------+
  | [Advertiser] TechCorp Inc.  |
  | [Supplier] Store Network    |
  +-----------------------------+
      |
      v
User selects context
      |
      v
Load appropriate dashboard and permissions
```

### Context Switching Flow
```
User clicks account switcher
      |
      v
Show available contexts:
  - Current: Advertiser - TechCorp
  - Switch to: Supplier - Store Network
      |
      v
User selects new context
      |
      v
System recalculates permissions
      |
      v
Redirect to new context dashboard
```

### Adding New Account Type
```
User goes to Settings > Account
      |
      v
Click "Add Account Type"
      |
      v
Select: Advertiser or Supplier
      |
      v
Complete onboarding for new account type
      |
      v
New context available in switcher
```

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
