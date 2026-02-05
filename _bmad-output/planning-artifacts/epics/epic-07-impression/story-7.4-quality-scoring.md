---
id: "STORY-7.4"
epic_id: "EPIC-07"
title: "Quality Scoring"
status: "to-do"
priority: "medium"
assigned_to: null
tags: ["quality", "scoring", "pricing"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Quality Scoring

## User Story

**As a** platform operator,
**I want** to score impression quality,
**So that** pricing can reflect actual value delivered.

## Acceptance Criteria

### Quality Factors
| Factor | Weight | Description |
|--------|--------|-------------|
| Duration | 30% | Full play vs partial |
| Device Health | 25% | Device operating condition |
| Time Match | 20% | Within daypart schedule |
| Visibility | 15% | No obstructions detected |
| Uniqueness | 10% | Not duplicate or loop |

### Quality Score Range
- [ ] 100: Perfect impression
- [ ] 80-99: High quality
- [ ] 60-79: Standard quality
- [ ] 40-59: Below average
- [ ] 0-39: Low quality (may not be billable)

### Scoring Rules

#### Duration Score
| Completion | Score |
|------------|-------|
| 100% | 100 |
| 75-99% | 80 |
| 50-74% | 60 |
| 25-49% | 40 |
| <25% | 20 |

#### Device Health Score
| Health | Score |
|--------|-------|
| Excellent (90-100) | 100 |
| Good (75-89) | 80 |
| Fair (50-74) | 60 |
| Poor (<50) | 40 |

### Quality-Based Pricing
- [ ] Premium quality (90+): Full CPM
- [ ] Standard quality (70-89): 90% CPM
- [ ] Below average (50-69): 70% CPM
- [ ] Low quality (<50): 50% CPM or no charge

## Business Flow

### Quality Score Calculation

```
┌─────────────────────────────────────────────────────────────────┐
│                    QUALITY SCORE COMPONENTS                     │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   Technical Quality (30% of score)                              │
│   ├── Did the ad play fully?                                    │
│   ├── Is the device clock accurate?                             │
│   ├── Is the network connection good?                           │
│   └── Is the device healthy?                                    │
│                                                                 │
│   Proof Quality (25% of score)                                  │
│   ├── Was a screenshot captured?                                │
│   ├── Is the digital signature valid?                           │
│   └── Is GPS data available?                                    │
│                                                                 │
│   Viewability (20% of score)                                    │
│   ├── Is the screen brightness sufficient?                      │
│   ├── Is the surrounding area visible (not too dark)?           │
│   └── Was audio enabled (if applicable)?                        │
│                                                                 │
│   Location (15% of score)                                       │
│   ├── Is GPS available?                                         │
│   └── How close is the device to the registered store?          │
│                                                                 │
│   Timing (10% of score)                                         │
│   ├── Was it during store operating hours?                      │
│   └── Was it during peak traffic hours?                         │
│                                                                 │
│   ═══════════════════════════════════════════════════           │
│                  FINAL QUALITY SCORE (0-100)                    │
│   ═══════════════════════════════════════════════════           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

### Quality Tier Classification

| Score Range | Tier | What It Means | Processing |
|-------------|------|---------------|------------|
| 90-100 | Premium | Excellent quality | Fast-tracked verification |
| 70-89 | Standard | Good quality | Normal processing |
| 50-69 | Basic | Acceptable, minor issues | May require review |
| Below 50 | Poor | Questionable | Held for review |
| Below 30 | Very Poor | Unreliable | Manual review required |

### Impact on Stakeholders

**For Advertisers:**
| Quality Tier | Benefit |
|--------------|--------|
| Premium | Full confidence in ad delivery |
| Standard | Good value, reliable impressions |
| Basic | Acceptable but may have minor issues |
| Poor | Can choose not to pay for these |

**For Suppliers:**
| Quality Trend | Meaning |
|---------------|--------|
| Consistently high | Device performing well |
| Consistently low | Device may need maintenance |
| Declining over time | Investigation recommended |

### Duration Score Breakdown

| Completion Percentage | Score Contribution |
|-----------------------|-------------------|
| 100% (full play) | 100 points |
| 75-99% | 80 points |
| 50-74% | 60 points |
| 25-49% | 40 points |
| Less than 25% | 20 points |

## Checklist (Subtasks)

- [ ] Define quality scoring algorithm
- [ ] Implement duration scoring
- [ ] Implement device health scoring
- [ ] Implement time match scoring
- [ ] Implement visibility scoring
- [ ] Implement uniqueness scoring
- [ ] Create quality score calculation job
- [ ] Add price multiplier calculation
- [ ] Build quality dashboard
- [ ] Create quality trend reports
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
