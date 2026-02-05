---
id: "STORY-8.6"
epic_id: "EPIC-08"
title: "Refunds & Credits"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["refund", "credit", "adjustment"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Refunds & Credits

## User Story

**As an** advertiser,
**I want** to receive refunds or credits when appropriate,
**So that** I'm not charged for issues beyond my control.

## Acceptance Criteria

### Refund Scenarios
- [ ] Campaign cancelled before start
- [ ] Technical issues prevented delivery
- [ ] Content moderation delays
- [ ] Disputed impressions found invalid

### Credit Scenarios
- [ ] Promotional credits
- [ ] Goodwill credits
- [ ] Dispute resolution credits
- [ ] Service level agreement violations

### Refund Types
| Type | Returns To | Timeline |
|------|------------|----------|
| Wallet Refund | Wallet balance | Immediate |
| Payment Refund | Original payment method | 5-10 days |

### Refund Process
1. **Request Submitted**
   - [ ] Advertiser requests or admin initiates
   - [ ] Reason provided
   - [ ] Amount calculated

2. **Review**
   - [ ] Validate refund eligibility
   - [ ] Approve or deny
   - [ ] Calculate final amount

3. **Processing**
   - [ ] Credit wallet or refund payment
   - [ ] Update transaction records
   - [ ] Send confirmation

### Credit Application
- [ ] Credits applied to wallet balance
- [ ] Credits expire after 1 year
- [ ] Credits used before wallet balance
- [ ] Credits cannot be refunded

## Business Flow

### Automatic Refund Scenarios

| Scenario | Refund Amount | Timing |
|----------|---------------|--------|
| Campaign cancelled before start | 100% of budget | Immediate |
| Unused budget at completion | Remaining budget | Automatic |
| Failed payment reversal | Full amount | Immediate |

### Manual Refund Scenarios

| Scenario | Process |
|----------|----------|
| Customer service request | Admin reviews, approves/rejects |
| Platform error | Admin initiates correction |
| Dispute settlement | Admin determines amount |

### Refund Approval Thresholds

| Refund Amount | Approval Required |
|---------------|-------------------|
| Under $1,000 | Auto-approve |
| $1,000 or more | Admin approval |

### Refund Processing Flow

```
Refund requested or triggered
       |
       v
Validate refund eligibility
       |
       v
Determine refund method:
       |
       +---> Within 90 days: Refund to original payment method
       |
       +---> After 90 days: Wallet credit only
       |
       +---> Payment method unavailable: Wallet credit
       |
       v
Process refund (1-3 business days)
       |
       v
Update transaction records
Send confirmation to advertiser
```

### Refund Method Priority

1. **Original payment method** (if within 90 days)
2. **Wallet credit** (if payment method unavailable)
3. **Bank transfer** (last resort)

### Refund Business Rules

| Rule | Requirement |
|------|-------------|
| Refund window | 90 days for payment gateway refund |
| After 90 days | Wallet credit only |
| Refund fees | None (platform absorbs) |
| Processing time | 1-3 business days |

### Playtime Guarantee Shortfall Refund

```
Package: 10,000 impressions within 1 month
Price: $500 fixed

If shortfall occurs:
- Option A: Extend campaign until target reached
- Option B: Refund proportional amount

Example:
- Played: 8,500 impressions
- Shortfall: 1,500 impressions
- Refund calculation: (1,500 / 10,000) x $500 = $75
```

## Checklist (Subtasks)

- [ ] Create refund request flow
- [ ] Build refund review interface
- [ ] Implement wallet refund
- [ ] Integrate payment gateway refunds
- [ ] Create credit application
- [ ] Implement credit expiration
- [ ] Add credit usage priority
- [ ] Build refund history
- [ ] Create credit balance display
- [ ] Add refund notifications
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
