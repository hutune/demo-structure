---
id: "EPIC-02"
title: "Advertiser Account Management"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["epic", "advertiser", "account", "onboarding", "kyc"]
start_date: null
due_date: null
clickup_task_id: null
business_rule_ref: "02-advertiser.md"
---

# Epic 02: Advertiser Account Management

## Overview

This epic covers all advertiser-specific account functionality including onboarding, account tiers, spending limits, identity verification (KYC), and team management. Advertisers are businesses or individuals who want to display ads on digital screens in retail locations.

## Business Value

- Enable seamless advertiser onboarding experience
- Provide tiered service levels with appropriate limits
- Ensure compliance through identity verification
- Enable team collaboration for larger advertisers

## Goals

- Implement complete advertiser signup flow
- Support four account tiers (FREE, BASIC, PREMIUM, ENTERPRISE)
- Implement spending limits and enforcement
- Enable KYC verification for higher limits
- Support team member management

## User Stories

| ID | Title | Status | Assignee |
|----|-------|--------|----------|
| STORY-2.1 | Advertiser Onboarding | to-do | - |
| STORY-2.2 | Account Tiers & Limits | to-do | - |
| STORY-2.3 | Identity Verification (KYC) | to-do | - |
| STORY-2.4 | Team Management | to-do | - |
| STORY-2.5 | Account Status & Suspension | to-do | - |
| STORY-2.6 | Tier Upgrade & Downgrade | to-do | - |

## Success Metrics

- Onboarding completion rate > 80%
- KYC verification turnaround < 48 hours
- Account tier upgrade rate > 20%
- Team collaboration adoption in eligible tiers > 50%

## Dependencies

- Epic-01: User & Authentication
- Payment gateway integration
- KYC verification service

## Technical Notes

### Account Tiers
| Tier | Monthly Fee | Max Campaigns | Max Budget/Campaign | Max Daily Spend |
|------|-------------|---------------|---------------------|-----------------|
| FREE | $0 | 2 | $500 | $100 |
| BASIC | $99 | 5 | $2,000 | $500 |
| PREMIUM | $499 | 20 | $10,000 | $2,000 |
| ENTERPRISE | Custom | Unlimited | Custom | Custom |

## Reference Documents

- [02-advertiser.md](../../../../docs/_rmn-arms-docs/business-rules%20(en%20ver)/02-advertiser.md)

## Updates

<!-- 
Progress updates will be added here.
Format: **YYYY-MM-DD** - Status update
-->
