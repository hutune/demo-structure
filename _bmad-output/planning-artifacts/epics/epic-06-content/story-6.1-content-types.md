---
id: "STORY-6.1"
epic_id: "EPIC-06"
title: "Content Types & Specifications"
status: "to-do"
priority: "high"
assigned_to: null
tags: ["content", "types", "specifications"]
story_points: 5
sprint: null
start_date: null
due_date: null
time_estimate: "2d"
clickup_task_id: null
---

# Content Types & Specifications

## User Story

**As an** advertiser,
**I want** to know the content specifications for different ad formats,
**So that** I can create assets that work well on the digital screens.

## Acceptance Criteria

### Content Types
- [ ] **Static Image**: JPG, PNG
- [ ] **Video**: MP4 with H.264 codec
- [ ] **Animated Image**: GIF (limited use)
- [ ] **HTML5**: Interactive content package

### Specifications by Type

#### Static Image
| Property | Requirement |
|----------|-------------|
| Formats | JPG, PNG |
| Max Size | 10 MB |
| Resolutions | 1920x1080, 1080x1920, 3840x2160 |
| Color Space | sRGB |
| Safe Zone | 90% of frame |

#### Video
| Property | Requirement |
|----------|-------------|
| Format | MP4 |
| Codec | H.264 |
| Max Size | 100 MB |
| Max Duration | 60 seconds (15s or 30s recommended) |
| Min Resolution | 1080p |
| Frame Rate | 24, 25, 30 fps |
| Audio | Optional, AAC codec |

#### HTML5
| Property | Requirement |
|----------|-------------|
| Format | ZIP package |
| Max Size | 5 MB |
| Entry Point | index.html |
| Supported APIs | CSS3, Canvas, WebGL (limited) |
| External Resources | Not allowed (self-contained) |

### Resolution Support
- [ ] Landscape: 1920x1080 (Full HD), 3840x2160 (4K)
- [ ] Portrait: 1080x1920 (Full HD), 2160x3840 (4K)
- [ ] Auto-scaling for different device sizes

## Business Flow

### Supported Content Types by Tier

```
+------------------+----------------------------------+
|     TIER         |       ALLOWED FORMATS            |
+------------------+----------------------------------+
| All Tiers        | JPG, PNG, GIF, WebP, MP4         |
+------------------+----------------------------------+
| Professional+    | + SVG, WebM, MOV, MP3, AAC,      |
|                  |   WAV, HTML5                     |
+------------------+----------------------------------+
```

### Content Requirements Summary

| Content Type | Maximum Size | Minimum Resolution | Use Case |
|--------------|--------------|-------------------|----------|
| Image | 10 MB | 1920 × 1080 | Static display ads, banners |
| Video | 500 MB | 1920 × 1080, 10-60 sec | Video ads, motion graphics |
| Audio | 50 MB | 128 kbps | Audio-enabled displays |
| Document | 20 MB | N/A | Menu boards, information |
| Rich Media | 50 MB | N/A | Interactive ads (HTML5) |

### Aspect Ratio Compatibility

```
Content Upload          Device Screen Match
     |                        |
     v                        v
+----------+            +----------+
| 16:9     |  =======>  | 16:9     |  ✓ Perfect match
|Landscape |            |Landscape |
+----------+            +----------+

+----------+            +----------+
| 9:16     |  =======>  | 9:16     |  ✓ Perfect match
| Portrait |            | Portrait |
+----------+            +----------+

+----------+            +----------+
| 16:9     |  --X-->    | 9:16     |  ⚠ Mismatch!
|Landscape |            | Portrait |  Options:
+----------+            +----------+  - Crop/fit (letterbox)
                                      - Use different asset
```

### Resolution Validation Rules

| Device Display | Recommended Content Resolution | Result if Lower |
|----------------|-------------------------------|----------------|
| 1080p Display | 1920 × 1080 or higher | Warning: may appear pixelated |
| 4K Display | 3840 × 2160 or higher | Warning: may appear pixelated |

## Checklist (Subtasks)

- [ ] Define content type specifications
- [ ] Create specification documentation
- [ ] Build specification display UI
- [ ] Implement format validation
- [ ] Add resolution validation
- [ ] Create file size validation
- [ ] Add video duration validation
- [ ] Build HTML5 package validator
- [ ] Create specification helper/tooltip
- [ ] Write unit tests
- [ ] Write integration tests

## Updates

<!-- 
Dev comments will be added here by AI when updating via chat.
Format: **YYYY-MM-DD HH:MM** - @author: Message
-->
