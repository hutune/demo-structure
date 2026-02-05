---
id: "STORY-8.3"
epic_id: "EPIC-08"
title: "Top-up & Payments"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["payment", "topup", "checkout"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Top-up & Payments

## User Story

**As an** advertiser,
**I want** to add funds to my wallet easily,
**So that** I can run advertising campaigns without interruption.

## Acceptance Criteria

### Payment Methods
- [ ] Credit/debit card (Visa, Mastercard, Amex)
- [ ] Bank transfer (ACH, SEPA, wire)
- [ ] PayPal (if available)
- [ ] Invoice/net terms (for qualified accounts)

### Top-up Flow
1. **Select Amount**
   - [ ] Preset amounts ($100, $500, $1000, $5000)
   - [ ] Custom amount input
   - [ ] Minimum: $50
   - [ ] Maximum: Based on tier limits

2. **Select Payment Method**
   - [ ] Choose saved method
   - [ ] Add new payment method
   - [ ] One-time payment option

3. **Review & Confirm**
   - [ ] Show amount and fees
   - [ ] Show estimated new balance
   - [ ] Terms confirmation

4. **Process Payment**
   - [ ] Payment processed
   - [ ] Funds added to wallet
   - [ ] Receipt sent via email

### Saved Payment Methods
- [ ] Save card for future use
- [ ] Nickname payment methods
- [ ] Set default payment method
- [ ] Delete saved methods
- [ ] PCI compliant storage (tokenization)

### Auto Top-up
- [ ] Enable auto top-up when balance falls below threshold
- [ ] Set threshold amount
- [ ] Set top-up amount
- [ ] Set maximum auto top-ups per month

### Failed Payments
- [ ] Retry logic for declined cards
- [ ] Notify advertiser of failure
- [ ] Suggest alternative payment methods

## Business Flow

### Top-up Processing Flow

```
User initiates top-up
       |
       v
System validates:
- Amount within limits ($50 - $10,000)
- Daily limit not exceeded
- Wallet status is Active
       |
       v
Create pending transaction
Add amount to Pending Balance
       |
       v
Process payment via gateway
       |
       +---> Success: Move to Available Balance
       |     Notify user: "Top-up successful"
       |
       +---> Failure: Remove from Pending Balance
       |     Notify user: "Top-up failed"
       |
       +---> Requires 3D Secure: User completes verification
             Then process result
```

### Transaction Limits

| Limit Type | Amount |
|------------|--------|
| Minimum top-up | $50 |
| Maximum top-up | $10,000 |
| Maximum transactions per day | 10 |
| Cooldown between transactions | 1 minute |

### Daily Limits by Verification Level

| Verification Level | Daily Limit |
|-------------------|-------------|
| Unverified | $500 |
| Verified | $10,000 |
| Enterprise | Custom (negotiated) |

### Supported Payment Methods

| Method | Processing Time | Fees |
|--------|-----------------|------|
| Credit Card | Instant | 2.9% + $0.30 |
| Debit Card | Instant | 2.9% + $0.30 |
| Bank Transfer | 1-3 business days | $0 |

### Top-up Business Rules

| Rule | Requirement |
|------|-------------|
| Minimum amount | $50 per transaction |
| Maximum amount | $10,000 per transaction |
| Daily limit check | Enforced before processing |
| Failed payments | Automatically reversed |
| 3D Secure | Required for card payments (EU regulation) |

## Checklist (Subtasks)

- [ ] Integrate payment gateway (Stripe)
- [ ] Create top-up amount selection
- [ ] Build payment method selector
- [ ] Implement saved payment methods
- [ ] Add new card flow
- [ ] Create payment processing
- [ ] Handle webhook confirmations
- [ ] Credit wallet on success
- [ ] Generate receipts
- [ ] Implement auto top-up
- [ ] Add failed payment handling
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
