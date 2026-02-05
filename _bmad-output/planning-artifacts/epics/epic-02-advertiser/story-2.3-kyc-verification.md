---
id: "STORY-2.3"
epic_id: "EPIC-02"
title: "Identity Verification (KYC)"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["kyc", "verification", "compliance"]
story_points: 13
sprint: null
start_date: null
due_date: null
time_estimate: "5d"
clickup_task_id: null
---

# Identity Verification (KYC)

## User Story

**As an** advertiser,
**I want** to verify my identity,
**So that** I can unlock higher spending limits and premium features.

## Acceptance Criteria

### Verification Benefits
- [ ] Without verification: $100 daily, $1,000 monthly limit
- [ ] With verification: $10,000 daily, $100,000 monthly limit
- [ ] Verification required for Premium/Enterprise tiers

### Individual Verification
- [ ] Upload government-issued ID (passport, driver's license, national ID)
- [ ] Selfie photo for face matching
- [ ] Proof of address for >$10,000/month spending
- [ ] Approval within 24-48 hours

### Business Verification
- [ ] Business registration certificate
- [ ] Tax ID (EIN, VAT number)
- [ ] Articles of incorporation
- [ ] Business bank statement
- [ ] ID of authorized representative
- [ ] Proof of business address
- [ ] Approval within 2-5 business days

### Verification Statuses
- [ ] Unverified: No documents submitted
- [ ] Pending: Documents submitted, under review
- [ ] Verified: Approved, higher limits unlocked
- [ ] Rejected: Documents not accepted, can resubmit
- [ ] Expired: Verification expired, need to re-verify

### Verification Expiration
- [ ] Individual: Valid for 2 years
- [ ] Business: Valid for 1 year
- [ ] Reminder sent 30 days before expiration

## Business Flow

### KYC Verification Flow

```
    Start Verification
           │
           v
    ┌──────────────────┐
    │ Select Type:     │
    │ Individual or    │
    │ Business?        │
    └────────┬─────────┘
             │
      ┌──────┴──────┐
      v             v
  Individual    Business
      │             │
      v             v
  Upload:       Upload:
  • ID          • Business registration
  • Selfie      • Tax ID
  • Address*    • Articles of incorporation
      │         • Bank statement
      │         • Representative ID
      │         • Business address
      │             │
      └──────┬──────┘
             v
    Documents Submitted
    (Status: Pending)
             │
             v
    ┌──────────────────┐
    │ Automated Checks │
    │ + Manual Review  │
    └────────┬─────────┘
             │
      ┌──────┼──────┐
      v      v      v
  Approved Rejected Expired
  (Verified) (Retry)  (Re-verify)
```

*Address proof required only for spending > $10,000/month

### Verification Status Reference

| Status | What It Means | Next Action |
|--------|---------------|-------------|
| Unverified | No documents submitted | Start verification |
| Pending | Documents under review | Wait 24-48 hours (Individual) or 2-5 days (Business) |
| Verified | Approved - limits unlocked | Enjoy higher limits |
| Rejected | Documents not accepted | Resubmit correct documents |
| Expired | Verification expired | Re-verify to maintain limits |

### Spending Limits Comparison

| Verification Status | Daily Limit | Monthly Limit |
|---------------------|-------------|---------------|
| Without Verification | $100 | $1,000 |
| With Verification | $10,000 | $100,000 |

### Verification Validity

| Type | Valid For | Reminder Sent |
|------|-----------|---------------|
| Individual | 2 years | 30 days before expiration |
| Business | 1 year | 30 days before expiration |

## Checklist (Subtasks)

- [ ] Create document upload interface
- [ ] Implement document type validation
- [ ] Integrate with KYC verification service
- [ ] Create automated document checks
- [ ] Build manual review queue for admins
- [ ] Implement verification status tracking
- [ ] Add limit unlocking on verification
- [ ] Create expiration tracking
- [ ] Send 30-day reminder notifications
- [ ] Implement re-verification flow
- [ ] Add rejection reason feedback
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
