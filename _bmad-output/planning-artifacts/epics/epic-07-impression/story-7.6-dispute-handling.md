---
id: "STORY-7.6"
epic_id: "EPIC-07"
title: "Dispute Handling"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["dispute", "resolution", "support"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Dispute Handling

## User Story

**As an** advertiser or supplier,
**I want** to dispute impression counts if I believe they are inaccurate,
**So that** billing issues can be investigated and resolved fairly.

## Acceptance Criteria

### Dispute Submission
- [ ] Submit dispute through portal
- [ ] Select time period in question
- [ ] Describe discrepancy
- [ ] Upload supporting evidence (optional)
- [ ] Provide expected count vs actual

### Dispute Types
- [ ] **Undercounting**: Fewer impressions than expected
- [ ] **Overcounting**: More charges than expected
- [ ] **Quality Issue**: Impressions not as described
- [ ] **Technical Issue**: Device/network problems

### Dispute Workflow
```
Submitted → Under Review → Investigating → Resolved/Rejected
               ↓                ↓
         Need Info → Customer provides
```

### Investigation Process
- [ ] Review impression logs
- [ ] Check proof-of-play evidence
- [ ] Analyze device health data
- [ ] Compare with expected delivery
- [ ] Consult technical team if needed

### Resolution Options
- [ ] **Approved**: Credit issued to advertiser
- [ ] **Partially Approved**: Partial credit
- [ ] **Rejected**: No change, with explanation
- [ ] **Technical Error**: System fix + credit

### SLA
| Priority | Response Time | Resolution Time |
|----------|---------------|-----------------|
| Standard | 24 hours | 5 business days |
| Urgent | 4 hours | 2 business days |
| Critical | 1 hour | 24 hours |

## Business Flow

### Dispute Resolution Flow

```
┌─────────────────────────────────────────────────────────────────┐
│                    DISPUTE RESOLUTION FLOW                      │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   Advertiser files dispute                                      │
│   (within 30 days of impression)                                │
│          │                                                      │
│          ↓                                                      │
│   ┌─────────────────────────────────┐                           │
│   │ Billing temporarily reversed    │                           │
│   │ • Advertiser credited           │                           │
│   │ • Supplier revenue held         │                           │
│   └─────────────┬───────────────────┘                           │
│                 │                                               │
│                 ↓                                               │
│   Admin assigned to investigate                                 │
│   (Must resolve within 72 hours)                                │
│                 │                                               │
│                 ↓                                               │
│   ┌─────────────────────────────────┐                           │
│   │ Admin reviews evidence:         │                           │
│   │ • Screenshot and proof data     │                           │
│   │ • Device logs and history       │                           │
│   │ • Location data                 │                           │
│   │ • Pattern analysis              │                           │
│   └─────────────┬───────────────────┘                           │
│                 │                                               │
│          ┌──────┴──────┐                                        │
│          │             │                                        │
│          ↓             ↓                                        │
│       Dispute       Dispute                                     │
│       Upheld        Denied                                      │
│          │             │                                        │
│          ↓             ↓                                        │
│   ┌────────────┐ ┌─────────────────┐                            │
│   │ Refund to  │ │ No refund       │                            │
│   │ Advertiser │ │ Supplier paid   │                            │
│   │            │ │                 │                            │
│   │ Supplier   │ │ Advertiser can  │                            │
│   │ loses      │ │ appeal within   │                            │
│   │ revenue    │ │ 7 days          │                            │
│   └────────────┘ └─────────────────┘                            │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Valid Dispute Reasons

| Reason | Example |
|--------|--------|
| Invalid Proof | Screenshot shows different content than campaign |
| Device Offline | Device was marked offline at the reported time |
| Wrong Location | Device was not at the correct store |
| Duplicate Charge | Same impression charged twice |
| Wrong Content | Different ad was shown than assigned |
| Poor Quality | Playback was unacceptably poor |

### Dispute Timeline Rules

| Action | Time Limit |
|--------|------------|
| File a dispute | Within 30 days of impression |
| Admin assignment | Immediately upon filing |
| Investigation completion | Within 72 hours |
| Appeal a denied dispute | Within 7 days of denial |

### Resolution Outcomes

**If Dispute Upheld (Advertiser Wins):**
| Impact | Action |
|--------|--------|
| Advertiser | Receives full refund |
| Supplier | Loses revenue for that impression |
| Device | Reputation score drops |
| Device with 5+ chargebacks | May be put in maintenance |

**If Dispute Denied (Supplier Wins):**
| Impact | Action |
|--------|--------|
| Billing | Original charge stands |
| Supplier | Revenue released from hold |
| Advertiser | Can appeal within 7 days |

### Resolution Timeline by Customer Tier

| Customer Tier | Standard Resolution | Priority Resolution |
|---------------|---------------------|--------------------|
| Free/Basic | 72 hours | Not available |
| Premium | 48 hours | 24 hours |
| Enterprise | 24 hours | 8 hours |

## Checklist (Subtasks)

- [ ] Create dispute submission form
- [ ] Build dispute type selection
- [ ] Add evidence upload
- [ ] Implement dispute queue
- [ ] Create assignment workflow
- [ ] Build investigation dashboard
- [ ] Add data comparison tools
- [ ] Implement resolution workflow
- [ ] Create credit issuance
- [ ] Build notification system
- [ ] Track SLA compliance
- [ ] Create dispute reporting
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
