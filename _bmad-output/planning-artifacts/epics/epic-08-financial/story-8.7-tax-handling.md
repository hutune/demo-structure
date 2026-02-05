---
id: "STORY-8.7"
epic_id: "EPIC-08"
title: "Tax Handling"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["tax", "compliance", "invoice"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Tax Handling

## User Story

**As a** platform operator,
**I want** to handle taxes correctly,
**So that** the platform is compliant with tax regulations.

## Acceptance Criteria

### Tax Collection (Advertisers)
- [ ] Calculate applicable sales tax
- [ ] Apply based on advertiser location
- [ ] Handle tax-exempt accounts
- [ ] Generate tax invoices

### Tax Reporting (Suppliers)
- [ ] Track supplier earnings for 1099
- [ ] Collect W-9 information
- [ ] Generate 1099-K forms (US)
- [ ] VAT handling (EU)

### Tax Rates
- [ ] Use tax rate API for accuracy
- [ ] Support multiple jurisdictions
- [ ] Update rates automatically
- [ ] Handle special cases (digital goods)

### Invoicing
- [ ] Monthly invoices for advertisers
- [ ] Invoice includes:
  - Advertising spend
  - Applicable taxes
  - Platform fees
  - Credits/adjustments
- [ ] Invoice available for download
- [ ] Invoice sent via email

### Tax Exemption
- [ ] Verify tax-exempt status
- [ ] Store exemption certificates
- [ ] Apply exemptions automatically
- [ ] Annual re-verification

## Business Flow

### Sales Tax Calculation (Advertisers)

```
Campaign budget: $1,000
Tax rate (based on billing address): 8.25%
Tax amount: $1,000 x 8.25% = $82.50
Total charge: $1,000 + $82.50 = $1,082.50
```

### Sales Tax Applicability

| Region | Tax Type |
|--------|----------|
| US advertisers | State sales tax on ad spending |
| EU advertisers | VAT on services |
| Other regions | Per local tax law |

### Sales Tax Business Rules

| Rule | Description |
|------|-------------|
| Calculation timing | Tax calculated at campaign creation |
| Display | Tax displayed separately in invoice |
| Exemptions | Tax-exempt accounts (with certificate) exempted |
| Remittance | Tax remitted to authorities monthly |

### Withholding Tax (Suppliers) - US

| Status | Withholding Rate |
|--------|------------------|
| W-9 form on file | 0% |
| No W-9 form | 24% backup withholding |

### Withholding Tax (Suppliers) - Non-US

| Status | Withholding Rate |
|--------|------------------|
| Tax treaty country | Treaty rate |
| No tax treaty | 30% |

### Supplier Withdrawal with Withholding Example

```
Withdrawal amount: $1,000
Supplier: US, no W-9

Fee: $10
Tax withholding: $1,000 x 24% = $240
Net to bank: $1,000 - $10 - $240 = $750
```

### Year-End Tax Reporting

| Form | Recipient | Due Date |
|------|-----------|----------|
| 1099-K | US suppliers | January 31 |
| 1042-S | Non-US suppliers | March 15 |

### KYC Verification Tiers for Financial Limits

| Tier | Requirements | Daily Limit |
|------|--------------|-------------|
| Basic | Email verified | $500 |
| Verified | Government ID + Selfie | $10,000 |
| Enterprise | Business registration + Tax ID | Custom |

### AML Transaction Monitoring Thresholds

| Indicator | Threshold | Action |
|-----------|-----------|--------|
| Large single deposit | $10,000+ | Flag for review |
| Multiple small deposits (structuring) | Total $10,000+ in day | Flag for review |
| Rapid deposit and withdrawal | Deposit >$5,000, withdraw within 24h | Flag for review |
| Unusual activity pattern | 10x normal activity | Flag for review |
| High-risk country origin | Any amount | Flag for review |

## Checklist (Subtasks)

- [ ] Integrate tax rate API
- [ ] Implement tax calculation
- [ ] Add tax-exempt handling
- [ ] Store exemption certificates
- [ ] Create invoice generation
- [ ] Build invoice template
- [ ] Add PDF generation
- [ ] Implement invoice delivery
- [ ] Collect W-9 from suppliers
- [ ] Build 1099-K generation
- [ ] Add tax reporting dashboard
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
