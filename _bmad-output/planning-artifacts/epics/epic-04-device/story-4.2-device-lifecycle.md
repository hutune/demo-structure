---
id: "STORY-4.2"
epic_id: "EPIC-04"
title: "Device Lifecycle Management"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["device", "lifecycle", "status"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Device Lifecycle Management

## User Story

**As a** supplier or administrator,
**I want** to manage devices through their entire lifecycle,
**So that** I can track devices from acquisition to retirement.

## Acceptance Criteria

### Device Lifecycle Stages
- [ ] Ordered: Device ordered, not yet received
- [ ] Received: Device received, not yet configured
- [ ] Provisioning: Being set up and configured
- [ ] Active: Deployed and operational
- [ ] Maintenance: Temporarily offline for service
- [ ] Decommissioned: Retired from service

### Status Transitions
| From | To | Trigger |
|------|-----|---------|
| Ordered | Received | Mark as received |
| Received | Provisioning | Start setup |
| Provisioning | Active | Successful activation |
| Active | Maintenance | Manual or automated |
| Maintenance | Active | Service completed |
| Any | Decommissioned | Manual retirement |

### Device Operations
- [ ] Register new device (serial number, device type)
- [ ] Update device details
- [ ] Change device status
- [ ] Transfer between stores
- [ ] Decommission device

### Device Inventory
- [ ] View all devices with status
- [ ] Filter by status, store, supplier
- [ ] Search by serial number or ID
- [ ] Bulk operations (status change)

### Audit Trail
- [ ] Log all status changes
- [ ] Track who made changes
- [ ] Record timestamps
- [ ] View device history

## Business Flow

### Device Lifecycle Stages

```
┌─────────────────────────────────────────────────────────────────────┐
│                     DEVICE LIFECYCLE                                │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   [Device Manufactured/Purchased]                                   │
│              │                                                      │
│              ↓                                                      │
│   ┌─────────────────┐                                               │
│   │   REGISTERED    │ ← Device code generated,                      │
│   │                 │   waiting to be assigned to store             │
│   └────────┬────────┘                                               │
│            │                                                        │
│     Supplier assigns to store                                       │
│            │                                                        │
│            ↓                                                        │
│   ┌─────────────────┐                                               │
│   │     ACTIVE      │ ← Device is online, serving ads               │
│   │                 │                                               │
│   └────────┬────────┘                                               │
│            │                                                        │
│     ┌──────┴──────┐                                                 │
│     │             │                                                 │
│  Network       Scheduled                                            │
│  lost          maintenance                                          │
│     │             │                                                 │
│     ↓             ↓                                                 │
│ ┌─────────┐  ┌─────────────┐                                        │
│ │ OFFLINE │  │ MAINTENANCE │                                        │
│ └────┬────┘  └──────┬──────┘                                        │
│      │              │                                               │
│      │   Returns    │                                               │
│      └──────┬───────┘                                               │
│             ↓                                                       │
│      ┌─────────────────┐                                            │
│      │     ACTIVE      │                                            │
│      └────────┬────────┘                                            │
│               │                                                     │
│       Permanently removed                                           │
│               │                                                     │
│               ↓                                                     │
│      ┌─────────────────┐                                            │
│      │ DECOMMISSIONED  │ ← Cannot be reactivated                    │
│      └─────────────────┘                                            │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Status Descriptions

| Status | What It Means | Can Serve Ads? |
|--------|---------------|----------------|
| **Registered** | Device added to system, not yet installed | No |
| **Active** | Device is online and working normally | Yes |
| **Offline** | Lost connection (no heartbeat for 10+ minutes) | No |
| **Maintenance** | Temporarily down for scheduled work | No |
| **Decommissioned** | Permanently removed from service | No |

### Device Transfer Flow

Devices can be moved between stores (same supplier only):

```
┌─────────────────────────────────────────────────────────────────────┐
│                    DEVICE TRANSFER PROCESS                          │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   1. Supplier initiates transfer                                    │
│      │                                                              │
│      ↓                                                              │
│   2. Check for active campaigns exclusive to old store              │
│      │                                                              │
│      ├── If any found → Supplier warned, campaigns will pause       │
│      │                                                              │
│      ↓                                                              │
│   3. Device reassigned to new store                                 │
│      • Gets new store's operating hours                             │
│      • Gets new store's timezone                                    │
│      • Location updated                                             │
│      │                                                              │
│      ↓                                                              │
│   4. Campaigns recalculated                                         │
│      • Old store-specific campaigns paused                          │
│      • New store's eligible campaigns activated                     │
│      │                                                              │
│      ↓                                                              │
│   5. Content syncs to match new campaigns                           │
│      │                                                              │
│      ↓                                                              │
│   6. Transfer complete, audit log created                           │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Decommissioning Rules

| Rule | Description |
|------|-------------|
| **Permanent** | Cannot reactivate a decommissioned device |
| **Data Preserved** | Historical impression data kept for reports |
| **Code Reuse** | Device code can be reused after 1 year |
| **Campaign Handling** | Active campaigns automatically remove this device |
| **Grace Period** | 5 minutes to complete in-progress impressions |

## Checklist (Subtasks)

- [ ] Create device registration form
- [ ] Implement serial number validation
- [ ] Build device inventory listing
- [ ] Add status filter and search
- [ ] Create status change workflow
- [ ] Implement status transition rules
- [ ] Add store transfer feature
- [ ] Build decommission flow
- [ ] Create audit log for changes
- [ ] Implement device history view
- [ ] Add bulk operations
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
