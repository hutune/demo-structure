---
id: "STORY-8.4"
epic_id: "EPIC-08"
title: "Supplier Payouts"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["payout", "supplier", "revenue-share"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Supplier Payouts

## User Story

**As a** supplier,
**I want** to receive my share of advertising revenue regularly,
**So that** I can monetize my retail locations effectively.

## Acceptance Criteria

### Revenue Share
- [ ] Based on supplier tier (40-70%)
- [ ] Platform takes remainder
- [ ] Calculated per impression

### Payout Schedule
| Threshold | Schedule |
|-----------|----------|
| < $100 | Accumulate |
| $100+ | Monthly payout |
| $1,000+ | Bi-weekly available |
| $5,000+ | Weekly available |

### Payout Process
1. **Calculation Period**
   - [ ] Calculate earnings from impressions
   - [ ] Apply revenue share percentage
   - [ ] Deduct any fees or adjustments

2. **Review Period**
   - [ ] 7-day hold for disputes/verification
   - [ ] Admin can flag for review
   - [ ] Supplier can view pending amount

3. **Payout Execution**
   - [ ] Transfer to registered bank account
   - [ ] 3-5 business days for arrival
   - [ ] Confirmation notification

### Payout Dashboard
- [ ] Current balance (unpaid)
- [ ] Pending payout
- [ ] Recent payouts
- [ ] Earnings breakdown by store
- [ ] Revenue share tier

### Payout Methods
- [ ] Bank transfer (ACH, SEPA, wire)
- [ ] PayPal (if available)
- [ ] Hold for accumulation (optional)

## Business Flow

### Revenue Distribution Model

```
Impression cost paid by advertiser: $1.00
       |
       v
Supplier receives: $1.00 x 80% = $0.80
Platform keeps: $1.00 x 20% = $0.20
```

### Revenue Accrual Flow

```
Impression verified
       |
       v
Calculate supplier share (80%)
       |
       v
Credit to supplier Held Balance
       |
       v (7-day anti-fraud hold)
Release to Available Balance
       |
       v
Ready for withdrawal
```

### 7-Day Hold Period Rules

| Rule | Description |
|------|-------------|
| Hold duration | All revenue held for 7 days before release |
| Purpose | Protects against fraudulent impressions |
| Dispute handling | If disputed within 7 days, held until resolved |
| Auto-release | Released after 7 days if no disputes |

### Withdrawal Processing Flow

```
Supplier requests withdrawal
       |
       v
Validate:
- Amount >= $50
- Available Balance >= Amount + Fee
- Bank account verified
- No pending disputes
       |
       v
Calculate:
- Withdrawal amount
- Fee
- Tax withholding (if applicable)
- Net amount to bank
       |
       v
Available Balance -= Amount
Pending Balance += Amount
       |
       v
Approval:
- Auto-approve if < $5,000
- Manual review if >= $5,000
       |
       v
Process bank transfer (3-5 business days)
       |
       +---> Success: Remove from Pending, notify supplier
       |
       +---> Failure: Refund to Available Balance
```

### Withdrawal Limits & Fees

| Limit Type | Amount |
|------------|--------|
| Minimum withdrawal | $50 |
| Maximum withdrawal | No limit |

| Withdrawal Amount | Fee |
|-------------------|-----|
| Under $500 | $5 |
| $500 - $5,000 | $10 |
| Over $5,000 | $25 |

### Withdrawal Business Rules

| Rule | Requirement |
|------|-------------|
| Minimum amount | $50 |
| Processing time | 3-5 business days |
| Auto-approval threshold | Under $5,000 |
| Failed withdrawal | Automatically refunded |
| Retry limit | 3 attempts maximum |
| After 3 failures | Manual intervention required |

## Checklist (Subtasks)

- [ ] Implement revenue share calculation
- [ ] Create earnings tracking per store
- [ ] Build payout period aggregation
- [ ] Implement hold period
- [ ] Create payout batch processing
- [ ] Integrate bank transfer API
- [ ] Build payout dashboard
- [ ] Add earnings breakdown view
- [ ] Create payout history
- [ ] Implement payout notifications
- [ ] Add manual payout adjustment
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
