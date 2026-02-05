---
id: "STORY-1.8"
epic_id: "EPIC-01"
title: "KYC Verification (All User Types)"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["kyc", "verification", "identity", "compliance"]
story_points: 13
sprint: null
start_date: null
due_date: null
time_estimate: "5d"
clickup_task_id: null
---

# KYC Verification (All User Types)

## User Story

**As a** platform user (Advertiser, Supplier, or Admin),
**I want** to verify my identity through KYC process,
**So that** I can unlock higher spending limits and access premium features.

## Acceptance Criteria

### Verification Types
- [ ] Individual KYC (for personal accounts)
- [ ] Business KYC (for company accounts)
- [ ] Re-verification (when documents expire)

### Individual Verification
- [ ] User can upload government-issued ID (passport, driver's license, national ID)
- [ ] User takes selfie photo for face matching
- [ ] System performs automated document verification
- [ ] For spending > $10,000/month: proof of address required (utility bill, bank statement < 3 months old)
- [ ] Automated verification completes in < 5 minutes (if documents pass)
- [ ] Manual review escalation if automated checks fail
- [ ] Approval within 24-48 hours

### Business Verification
- [ ] User uploads business registration certificate
- [ ] Tax ID required (EIN, VAT number, etc.)
- [ ] Articles of incorporation for corporations
- [ ] Business bank statement (< 3 months old)
- [ ] ID of authorized representative
- [ ] Proof of business address
- [ ] Admin reviews all business documents manually
- [ ] Approval within 2-5 business days

### Verification Status Flow
- [ ] UNVERIFIED: No documents submitted
- [ ] PENDING: Documents submitted, under review
- [ ] VERIFIED: Approved, higher limits unlocked
- [ ] REJECTED: Documents not accepted, can resubmit with reason
- [ ] EXPIRED: Verification expired (Individual: 2 years, Business: 1 year)

### Spending Limits by Verification

| User Type | Without KYC | With KYC |
|-----------|-------------|----------|
| Advertiser | Daily: $100, Monthly: $1,000 | Daily: $10,000, Monthly: $100,000 |
| Supplier | Daily: $50, Monthly: $500 | Daily: $5,000, Monthly: $50,000 |

### Expiration & Renewal
- [ ] Individual verification valid for 2 years
- [ ] Business verification valid for 1 year
- [ ] Email reminder sent 30 days before expiration
- [ ] Warning notification 7 days before expiration
- [ ] Account limits revert to unverified if expired
- [ ] User can renew by re-uploading documents

### Document Security
- [ ] All documents encrypted at rest
- [ ] Documents accessible only by authorized admins
- [ ] Audit log for all document access
- [ ] Documents auto-deleted 90 days after verification expiry (GDPR compliance)

### Rejection Handling
- [ ] Clear rejection reason provided to user
- [ ] User notified via email and in-app notification
- [ ] User can resubmit corrected documents
- [ ] Max 3 attempts within 30 days (anti-fraud measure)

## Business Flow

### Individual KYC Flow
```
User initiates verification
          |
          v
Upload government ID
          |
          v
Take selfie photo
          |
          v
Automated checks:
  - Document authenticity
  - Face matching
  - Data extraction
          |
     +----+----+
     |         |
    Pass      Fail
     |         |
     v         v
Verify in  Manual review
< 5 min    by admin
     |         |
     +----+----+
          |
          v
  Spending > $10k/month?
          |
     +----+----+
     |         |
    Yes        No
     |         |
     v         v
Request    Status =
address    VERIFIED
proof
     |
     v
Upload proof
     |
     v
Status = VERIFIED
```

### Business KYC Flow
```
User initiates business verification
          |
          v
Upload business documents:
  - Registration certificate
  - Tax ID
  - Articles of incorporation
  - Bank statement
  - Representative ID
  - Business address proof
          |
          v
Documents queued for manual review
          |
          v
Admin checks:
  - Business registration validity
  - Tax ID verification
  - Bank statement authenticity
  - Representative authorization
          |
     +----+----+
     |         |
  Approved  Rejected
     |         |
     v         v
Status =   Notify user
VERIFIED   with reason
             |
             v
        User can resubmit
```

### Verification Expiration Flow
```
Verification approaching expiration
          |
          v
30 days before: Email reminder sent
          |
          v
7 days before: Warning notification
          |
          v
Expiration date reached
          |
          v
Status changes to EXPIRED
          |
          v
Spending limits revert to unverified
          |
          v
User prompted to renew
```

### Verification Status & Limits

| Status | Daily Limit (Advertiser) | Monthly Limit (Advertiser) | Can Access Premium? |
|--------|--------------------------|---------------------------|---------------------|
| UNVERIFIED | $100 | $1,000 | No |
| PENDING | $100 | $1,000 | No |
| VERIFIED | $10,000 | $100,000 | Yes |
| REJECTED | $100 | $1,000 | No |
| EXPIRED | $100 | $1,000 | No (was Yes before) |

### Document Requirements

#### Individual Documents
| Document | Required | Notes |
|----------|----------|-------|
| Government ID | Yes | Passport, Driver's License, or National ID |
| Selfie Photo | Yes | For face matching verification |
| Address Proof | If > $10k/month | Utility bill, Bank statement (< 3 months) |

#### Business Documents
| Document | Required | Notes |
|----------|----------|-------|
| Business Registration | Yes | Certificate of Incorporation, Business License |
| Tax ID | Yes | EIN (US), VAT (EU), or equivalent |
| Articles of Incorporation | Yes | For corporations |
| Bank Statement | Yes | Business account, < 3 months old |
| Representative ID | Yes | Authorized signer's ID |
| Business Address Proof | Yes | Official document with address |

### Validity Periods

| Verification Type | Valid For | Grace Period | Renewal Reminder |
|-------------------|-----------|--------------|------------------|
| Individual | 2 years | 14 days | 30 days before |
| Business | 1 year | 7 days | 30 days before |

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
