---
id: "STORY-4.3"
epic_id: "EPIC-04"
title: "Device Provisioning"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["device", "provisioning", "setup"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Device Provisioning

## User Story

**As a** supplier or technician,
**I want** to provision and configure devices easily,
**So that** new devices can be deployed quickly.

## Acceptance Criteria

### Provisioning Methods
- [ ] QR code scan (preferred for field deployment)
- [ ] Manual entry of device code
- [ ] Bulk import via CSV

### Provisioning Steps
1. **Physical Setup**
   - [ ] Power on device
   - [ ] Connect to network (Ethernet or WiFi)
   - [ ] Device displays setup screen with QR code

2. **Registration**
   - [ ] Scan QR code with mobile app
   - [ ] Or enter device code manually
   - [ ] Select supplier account
   - [ ] Select target store

3. **Configuration**
   - [ ] Apply device type settings
   - [ ] Set orientation
   - [ ] Configure content slot duration
   - [ ] Set operating hours (inherit from store or custom)

4. **Activation**
   - [ ] Download initial content
   - [ ] Verify display working
   - [ ] Mark as Active

### Network Configuration
- [ ] Support Ethernet connection
- [ ] Support WiFi with WPA2
- [ ] Proxy server configuration
- [ ] Firewall-friendly ports

### Provisioning Status
- [ ] Waiting for connection
- [ ] Connected, awaiting registration
- [ ] Registered, downloading content
- [ ] Ready, pending activation
- [ ] Active

## Business Flow

### Step-by-Step Provisioning Process

```
┌─────────────────────────────────────────────────────────────────────┐
│                    DEVICE PROVISIONING FLOW                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   STEP 1: REGISTRATION (Done by Platform or Supplier)               │
│   ┌─────────────────────────────────────────────────────────┐       │
│   │ • Unique device code generated (DVC-XXXX-XXXX-XXXX)     │       │
│   │ • QR code created for easy setup                        │       │
│   │ • Security keys generated (one-time only)               │       │
│   │ • Registration valid for 30 days                        │       │
│   └────────────────────────┬────────────────────────────────┘       │
│                            │                                        │
│                            ↓                                        │
│   STEP 2: STORE ASSIGNMENT (Done by Supplier)                       │
│   ┌─────────────────────────────────────────────────────────┐       │
│   │ • Scan QR code on device (or enter code manually)       │       │
│   │ • Enter activation key                                  │       │
│   │ • Select which store to assign device to                │       │
│   └────────────────────────┬────────────────────────────────┘       │
│                            │                                        │
│                            ↓                                        │
│   STEP 3: FIRST CONNECTION (Device Powers On)                       │
│   ┌─────────────────────────────────────────────────────────┐       │
│   │ • Device connects to internet                           │       │
│   │ • Reports hardware information                          │       │
│   │ • Platform verifies and activates                       │       │
│   │ • Device receives configuration settings                │       │
│   │ • Downloads initial content playlist                    │       │
│   └─────────────────────────────────────────────────────────┘       │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Assignment Requirements

| Requirement | Description |
|-------------|-------------|
| Device Status | Must be in "Registered" status |
| Store Ownership | Store must belong to the supplier |
| Store Status | Store must be active |
| Activation Key | Must be valid and unused |

### What Happens After Assignment

| Setting | Source |
|---------|--------|
| Operating Hours | Inherited from store |
| Timezone | Inherited from store |
| Ad Slots Per Hour | Default: 12 slots |

### Important Security Notes

```
┌─────────────────────────────────────────────────────────────────────┐
│                    SECURITY REMINDER                                │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   ⚠️  The security key (private key) is shown ONLY ONCE             │
│       during registration.                                          │
│                                                                     │
│   ⚠️  If lost, the device must be RE-REGISTERED.                    │
│                                                                     │
│   ⚠️  QR code/activation key expires after 30 days.                 │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Supplier Notification

When provisioning completes successfully:
- Supplier receives notification: "Device [code] at [store name] is now active"
- Device status changes: Registered → Active

## Checklist (Subtasks)

- [ ] Create QR code generation for devices
- [ ] Build mobile app scan feature
- [ ] Implement manual code entry
- [ ] Create store selection interface
- [ ] Build configuration wizard
- [ ] Implement content download trigger
- [ ] Create activation confirmation
- [ ] Add bulk import via CSV
- [ ] Build network configuration options
- [ ] Create provisioning status tracking
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
