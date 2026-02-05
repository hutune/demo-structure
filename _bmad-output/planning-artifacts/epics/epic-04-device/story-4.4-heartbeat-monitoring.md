---
id: "STORY-4.4"
epic_id: "EPIC-04"
title: "Heartbeat & Monitoring"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["heartbeat", "monitoring", "alerts"]
story_points: 8
sprint: null
start_date: null
due_date: null
time_estimate: "3d"
clickup_task_id: null
---

# Heartbeat & Monitoring

## User Story

**As a** platform operator,
**I want** devices to send regular heartbeats with status information,
**So that** I can monitor device health and detect issues quickly.

## Acceptance Criteria

### Heartbeat Frequency
- [ ] Normal operation: every 60 seconds
- [ ] Issue detected: every 15 seconds
- [ ] Recovery mode: every 30 seconds

### Heartbeat Payload
- [ ] Device ID
- [ ] Timestamp
- [ ] Current status (playing, idle, error)
- [ ] Network connectivity (signal strength, latency)
- [ ] Storage usage (% used)
- [ ] Temperature (if sensor available)
- [ ] Currently playing content
- [ ] Content playlist version
- [ ] Firmware version

### Offline Detection
- [ ] Device marked offline after 3 missed heartbeats (3 minutes)
- [ ] Alert sent to supplier after 5 minutes offline
- [ ] Critical alert after 30 minutes offline
- [ ] Affected campaigns notified after 1 hour offline

### Real-time Dashboard
- [ ] Map view showing device locations
- [ ] Color-coded status (green, yellow, red)
- [ ] Click device for details
- [ ] Filter by status, supplier, store

### Historical Data
- [ ] Store heartbeat data for 30 days
- [ ] Uptime calculations based on heartbeats
- [ ] Trend analysis for recurring issues

## Business Flow

### What Is a Heartbeat?

Every device sends a "heartbeat" signal to the platform every 5 minutes. This tells the platform the device is still online and reports its current performance.

### Heartbeat Cycle

```
┌─────────────────────────────────────────────────────────────────────┐
│                    HEARTBEAT CYCLE                                  │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   Every 5 minutes, the device reports:                              │
│                                                                     │
│   ┌─────────────────────┐                                           │
│   │ System Health       │                                           │
│   │ • Processor usage   │                                           │
│   │ • Memory usage      │                                           │
│   │ • Storage space     │                                           │
│   │ • Network speed     │                                           │
│   └─────────────────────┘                                           │
│                                                                     │
│   ┌─────────────────────┐                                           │
│   │ Playback Status     │                                           │
│   │ • Screen on/off     │                                           │
│   │ • Currently playing │                                           │
│   │ • Last ad shown     │                                           │
│   └─────────────────────┘                                           │
│                                                                     │
│   ┌─────────────────────┐                                           │
│   │ Errors (if any)     │                                           │
│   │ • Error messages    │                                           │
│   │ • Error count       │                                           │
│   └─────────────────────┘                                           │
│                                                                     │
│   Platform responds with:                                           │
│   • Confirmation ("All OK")                                         │
│   • New content to download (if available)                          │
│   • Configuration updates (if any)                                  │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

### Offline Detection & Alert Escalation

| Time Without Heartbeat | What Happens |
|------------------------|---------------|
| 5 minutes | One missed beat - still considered online |
| 10 minutes (2 missed) | Status changes to **Offline** |
| 15 minutes | Supplier receives first alert |
| 60 minutes (1 hour) | Supplier receives urgent alert |
| 4 hours | Critical alert - may need on-site check |

### When Device Goes Offline

| Impact | Description |
|--------|-------------|
| Ad Rotation | Device removed from ad rotation |
| Impressions | No impressions counted |
| Revenue | Supplier doesn't earn revenue for this device |
| Campaigns | Campaigns targeting only this device are paused |

### When Device Returns Online

| Action | Description |
|--------|-------------|
| Status Change | Automatically changes back to Active |
| Ad Service | Resumes serving ads immediately |
| Content Sync | Syncs to catch up on any content changes |

### Alert Channels

| Channel | Use Case |
|---------|----------|
| In-app notification | All alerts |
| Email | All alerts |
| SMS | Urgent alerts only (if configured) |

## Checklist (Subtasks)

- [ ] Create heartbeat receiving endpoint
- [ ] Implement heartbeat data storage
- [ ] Build offline detection job
- [ ] Create alert generation system
- [ ] Build real-time map dashboard
- [ ] Add device status color coding
- [ ] Implement device detail popup
- [ ] Create filter controls
- [ ] Build historical data queries
- [ ] Implement uptime calculation
- [ ] Create alert escalation workflow
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
