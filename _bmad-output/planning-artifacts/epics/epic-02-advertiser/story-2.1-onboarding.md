---
id: "STORY-2.1"
epic_id: "EPIC-02"
title: "Advertiser Onboarding"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["onboarding", "registration", "advertiser"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Advertiser Onboarding

## User Story

**As a** new advertiser,
**I want** to complete the signup process and set up my profile,
**So that** I can start creating advertising campaigns.

## Acceptance Criteria

### Step 1: Register with Email
- [ ] Enter email and password
- [ ] Receive verification email
- [ ] Click link to verify email

### Step 2: Set Up Profile
- [ ] Enter brand name (required, for competitor blocking)
- [ ] Select business type (individual, small business, agency, etc.)
- [ ] Choose industry (retail, food & beverage, electronics, etc.)
- [ ] Provide billing address (for tax purposes)
- [ ] Provide billing contact (name and email)

### Step 3: Account Created
- [ ] Account starts on FREE tier
- [ ] Wallet balance initialized to $0
- [ ] Welcome email sent with onboarding guide

### Step 4: Add Payment Method (Optional)
- [ ] Add credit card or bank account
- [ ] Payment method verified
- [ ] Can now top up wallet

### Step 5: First Campaign Guidance
- [ ] Show guided setup with tutorials
- [ ] Provide sample campaign templates

## Business Flow

### Advertiser Registration Flow

```
┌─────────────────────────────────────────────────────────────────────┐
│                    NEW ADVERTISER SIGNUP FLOW                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   Step 1: Register with Email                                       │
│   ├── Enter email and password                                      │
│   ├── Receive verification email                                    │
│   └── Click link to verify email                                    │
│                                                                     │
│   Step 2: Set Up Profile                                            │
│   ├── Enter company or brand name                                   │
│   ├── Select business type (individual, small business, etc.)       │
│   ├── Choose your industry                                          │
│   └── Provide billing address                                       │
│                                                                     │
│   Step 3: Account Created!                                          │
│   ├── You start on the FREE tier                                    │
│   ├── Wallet balance: $0                                            │
│   └── Ready to add payment method                                   │
│                                                                     │
│   Step 4: Add Payment Method (Optional but Recommended)             │
│   ├── Add credit card or bank account                               │
│   └── Can now top up wallet                                         │
│                                                                     │
│   Step 5: Verify Identity (Optional - for Higher Limits)            │
│   ├── Submit ID documents                                           │
│   ├── Admin reviews                                                 │
│   └── Unlock higher spending limits                                 │
│                                                                     │
│   Step 6: Create First Campaign                                     │
│   └── Follow guided setup with tutorials                            │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Required Information

| Type | Information | Purpose |
|------|-------------|----------|
| **Required** | Email address | Account login |
| **Required** | Brand name | Competitor blocking |
| **Required** | Business type | Service customization |
| **Required** | Industry | Content guidelines |
| **Required** | Billing address | Tax purposes |
| **Required** | Billing contact | Invoice communication |
| **Optional** | Website URL | Profile enhancement |
| **Optional** | Tax ID | Required for verification |

## Checklist (Subtasks)

- [ ] Create advertiser registration flow
- [ ] Build profile setup wizard (multi-step form)
- [ ] Implement brand name validation
- [ ] Add business type selection
- [ ] Add industry selection
- [ ] Implement billing address form
- [ ] Create wallet initialization
- [ ] Add payment method integration
- [ ] Design welcome email template
- [ ] Create first campaign tutorial
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
