# Claude code, Team how to.

\# Claude Code 팀 협업 가이드: Agent, Subagent, Skill 완전 정복

Claude Code는 \*\*3단계 계층 구조\*\*—Main Agent, Subagent, Skill—를 통해 팀 협업과 자동화를 지원합니다. 3인 소규모 팀(백엔드/프론트/PM)이 RMN SaaS 플랫폼 개발을 완전 자동화하려면 이 시스템의 전략적 활용이 핵심입니다. 팀 구독 환경에서 \`.claude/\` 디렉토리 구조와 \`CLAUDE.md\`를 Git으로 공유하면 즉시 동일한 개발 환경을 구축할 수 있습니다.

\---

\## Agent, Subagent, Skill의 핵심 차이점

Claude Code의 agent 시스템은 \*\*작업 분리\*\*와 \*\*컨텍스트 관리\*\*라는 두 가지 목표를 달성합니다. 각 구성요소의 역할이 명확히 다릅니다.

| 구성요소 | 목적 | 컨텍스트 | 호출 방식 | 저장 위치 |
|---------|------|---------|----------|----------|
| \*\*Main Agent\*\* | 핵심 대화 처리 | 공유 대화창 | 항상 활성화 | 내장 |
| \*\*Subagent\*\* | 전문 작업 위임 | \*\*독립된 컨텍스트\*\* | 자동/명시적 위임 | \`.claude/agents/\` |
| \*\*Skill\*\* | 지식/역량 주입 | 현재 컨텍스트에 로드 | 관련성 기반 자동 트리거 | \`.claude/skills/\` |

\*\*Subagent\*\*는 별도의 컨텍스트 창에서 작업하여 메인 대화를 오염시키지 않습니다. 코드 리뷰, 탐색, 복잡한 연구 등 독립 작업에 적합합니다. 반면 \*\*Skill\*\*은 현재 컨텍스트에 지식을 주입하는 방식으로, 점진적 공개(Progressive Disclosure) 메커니즘을 통해 필요한 정보만 로드합니다.

\### 내장 Subagent 3종

Claude Code는 세 가지 내장 subagent를 제공합니다:

\- \*\*Explore Subagent\*\*: Haiku 모델 사용, 읽기 전용으로 파일 탐색과 코드 검색에 특화
\- \*\*Plan Subagent\*\*: Sonnet 모델 사용, 코드베이스 분석 및 계획 수립용
\- \*\*General-purpose Subagent\*\*: Sonnet 모델 사용, 수정 포함 복잡한 연구 작업용

\---

\## 설정 파일 구조와 팀 공유 방법

\### 권장 디렉토리 구조

\`\`\`
your-rmn-project/
├── [CLAUDE.md](http://CLAUDE.md) # 프로젝트 메모리 (Git 커밋)
├── .mcp.json # MCP 서버 설정 (Git 커밋)
├── .claude/
│ ├── settings.json # 팀 공유 설정 (Git 커밋)
│ ├── settings.local.json # 개인 설정 (.gitignore)
│ │
│ ├── agents/ # 커스텀 Subagent
│ │ ├── [code-reviewer.md](http://code-reviewer.md)
│ │ ├── [golang-expert.md](http://golang-expert.md)
│ │ └── [flutter-specialist.md](http://flutter-specialist.md)
│ │
│ ├── commands/ # 슬래시 커맨드
│ │ ├── [test.md](http://test.md)
│ │ ├── [review.md](http://review.md)
│ │ └── deploy/
│ │ └── [staging.md](http://staging.md)
│ │
│ └── skills/ # Skills
│ ├── api-design/
│ │ └── [SKILL.md](http://SKILL.md)
│ └── testing-patterns/
│ └── [SKILL.md](http://SKILL.md)
│
└── agent\_docs/ # 점진적 공개용 문서
├── building\_the\_project.md
└── [architecture.md](http://architecture.md)
\`\`\`

\### Git을 통한 팀 공유

\*\*커밋 대상 (팀 공유):\*\*
\- \`CLAUDE.md\` - 프로젝트 컨텍스트
\- \`.claude/settings.json\` - 권한, 환경변수
\- \`.claude/commands/\` - 팀 슬래시 커맨드
\- \`.claude/agents/\` - 공유 subagent
\- \`.mcp.json\` - MCP 서버 설정

\*\*.gitignore 대상 (개인용):\*\*
\`\`\`gitignore
[CLAUDE.local.md](http://CLAUDE.local.md)
.claude/settings.local.json
.claude/local/
\`\`\`

팀원이 저장소를 clone하면 \`claude\` 명령어 실행 시 \*\*자동으로 동일한 환경\*\*이 적용됩니다.

\---

\## RMN SaaS 프로젝트용 [CLAUDE.md](http://CLAUDE.md) 템플릿

CLAUDE.md는 \*\*60~300줄 이내\*\*로 유지하는 것이 핵심입니다. LLM이 안정적으로 따를 수 있는 명령어 수가 제한되어 있기 때문입니다.

\`\`\`markdown
\# RMN SaaS 광고 관리 플랫폼

\## 프로젝트 개요
Retail Media Network 기반 SaaS 광고 관리 플랫폼.
Flutter 프론트엔드, Golang 백엔드, PostgreSQL, Kubernetes 운영.

\## 핵심 디렉토리
\- \`frontend/\` - Flutter 앱 (lib/features/, lib/shared/, lib/core/)
\- \`backend/\` - Go 서비스 (cmd/, internal/, pkg/, api/)
\- \`infra/\` - Kubernetes 매니페스트, Terraform
\- \`docs/\` - API 문서, PRD

\## 명령어
\`\`\`bash
_\# 백엔드_
cd backend && go build ./... && go test ./...
golangci-lint run

_\# 프론트엔드_
cd frontend && flutter analyze && flutter test
flutter build apk --release
\`\`\`

\## 코드 컨벤션
\### Golang
\- 인터페이스로 동작 정의, 생성자 함수로 DI
\- 에러는 항상 체크, %w로 래핑
\- Repository 패턴 + Service 레이어

\### Flutter
\- Feature-First Clean Architecture (FFCA)
\- Riverpod 상태관리, flutter\_hooks 사용
\- 다크 모드 전용, flex 기반 반응형 UI

\## 중요 사항
\- .env 파일 절대 커밋 금지
\- SQL 쿼리는 반드시 파라미터화
\- API 변경 시 OpenAPI 스펙 업데이트 필수

\## 추가 문서
복잡한 작업 전 \`agent\_docs/\` 참조:
\- \`architecture.md\` - 시스템 설계
\- \`api\_conventions.md\` - API 규칙
\`\`\`

\---

\## 역할별 맞춤 Agent/Skill 구성

\### 백엔드 개발자용 (Golang)

\*\*\`.claude/agents/golang-expert.md\`\*\*
\`\`\`markdown
\---
name: golang-expert
description: Go 코드 작성, 리뷰, 테스트 생성. Go 관련 작업 시 사용.
tools: Read, Write, Edit, Bash, Grep, Glob
model: claude-sonnet-4-20250514
\---

당신은 15년 경력의 Go 백엔드 전문가입니다.

\## 책임
\- Clean Architecture 원칙 준수
\- Table-driven 테스트 작성
\- 에러 핸들링 패턴 적용
\- PostgreSQL 쿼리 최적화

\## 코드 스타일
\- context.Context 필수 사용
\- 인터페이스 우선 설계
\- testify 어설션 사용
\`\`\`

\*\*\`.claude/commands/go-test.md\`\*\*
\`\`\`markdown
\---
name: go-test
description: Go 테이블 기반 테스트 생성
allowed-tools: Read, Write, Bash
\---

$ARGUMENTS 파일에 대한 테이블 기반 테스트를 생성하세요:
1\. 해피 패스, 에러 케이스, 엣지 케이스 포함
2\. testify/assert 사용
3\. 모킹은 외부 의존성만 대상
4\. \`go test -v\` 실행하여 검증
\`\`\`

\### 프론트엔드 개발자용 (Flutter)

\*\*\`.claude/agents/flutter-specialist.md\`\*\*
\`\`\`markdown
\---
name: flutter-specialist
description: Flutter 위젯 개발, 상태관리, UI 구현. Flutter 작업 시 사용.
tools: Read, Write, Edit, Bash, Grep, Glob
model: claude-sonnet-4-20250514
\---

당신은 Flutter/Dart 전문가입니다.

\## 아키텍처
\- Feature-First Clean Architecture
\- Riverpod + flutter\_hooks 조합
\- GoRouter 네비게이션

\## 위젯 원칙
\- 작고 조합 가능한 위젯 선호
\- 하드코딩 크기 대신 flex 값 사용
\- print() 대신 dart:developer의 log() 사용

\## 테스트
\- Widget Test로 UI 컴포넌트 검증
\- Unit Test로 비즈니스 로직 검증
\`\`\`

\*\*\`.claude/commands/flutter-feature.md\`\*\*
\`\`\`markdown
\---
name: flutter-feature
description: Flutter 피처 모듈 스캐폴딩 생성
\---

'$ARGUMENTS' 피처를 Feature-First 구조로 생성하세요:

\`\`\`
lib/features/$ARGUMENTS/
├── data/
│ ├── repositories/
│ └── models/
├── domain/
│ └── entities/
├── presentation/
│ ├── pages/
│ ├── widgets/
│ └── providers/
\`\`\`

각 레이어에 기본 파일과 Riverpod Provider를 생성하세요.
\`\`\`

\### PM용 문서화 도구

\*\*\`.claude/agents/prd-writer.md\`\*\*
\`\`\`markdown
\---
name: prd-writer
description: PRD, 기능 명세서, 릴리즈 노트 작성. 문서 작업 시 사용.
tools: Read, Write, Grep, Glob
model: claude-sonnet-4-20250514
\---

당신은 'Sarah'라는 이름의 시니어 PM입니다.

\## PRD 프레임워크 (15개 섹션)
1\. Executive Summary
2\. Problem Statement
3\. Goals & Success Metrics
4\. User Stories
5\. Functional Requirements
6\. Non-Functional Requirements
7\. Technical Constraints
8\. Dependencies
9\. Acceptance Criteria
10\. Out of Scope
11\. Risks & Mitigations
12\. Timeline
13\. Stakeholders
14\. Open Questions
15\. Appendix

PRD 작성 전 충분한 질문으로 요구사항을 명확히 하세요.
\`\`\`

\*\*\`.claude/commands/release-notes.md\`\*\*
\`\`\`markdown
\---
name: release-notes
description: 릴리즈 노트 자동 생성
\---

최근 $ARGUMENTS개 커밋을 분석하여 릴리즈 노트를 생성하세요:

1\. \`git log --oneline -n $ARGUMENTS\` 실행
2\. 커밋 유형별 분류 (feat, fix, docs, refactor)
3\. 사용자 관점의 변경사항 요약
4\. Breaking Changes 명시
5\. \`docs/CHANGELOG.md\` 업데이트
\`\`\`

\---

\## 팀 공유 설정 파일 예시

\*\*\`.claude/settings.json\`\*\*
\`\`\`json
{
"permissions": {
"allow": \[
"Bash(go build:\*)",
"Bash(go test:\*)",
"Bash(golangci-lint:\*)",
"Bash(flutter:\*)",
"Bash(git commit:\*)",
"Bash(git push:\*)",
"Bash(kubectl get:\*)"
\],
"deny": \[
"Bash(rm -rf:\*)",
"Read(.env\*)",
"Read(./secrets/\*\*)",
"Bash(kubectl delete:\*)"
\]
},
"env": {
"GO\_ENV": "development",
"FLUTTER\_ENV": "development"
},
"hooks": {
"PostToolUse": \[
{
"matcher": "Write(\*.go)",
"hooks": \[{
"type": "command",
"command": "gofmt -w \\"$file\\""
}\]
},
{
"matcher": "Write(\*.dart)",
"hooks": \[{
"type": "command",
"command": "dart format \\"$file\\""
}\]
}
\]
}
}
\`\`\`

이 설정은 Go/Dart 파일 저장 시 자동 포맷팅을 적용하고, 위험한 명령어 실행을 차단합니다.

\---

\## 코드 리뷰 및 테스트 자동화 실전 예시

\### GitHub Actions 통합

\*\*\`.github/workflows/claude-review.yml\`\*\*
\`\`\`yaml
name: Claude Code Review
on:
pull\_request:
types: \[opened, synchronize\]

jobs:
review:
runs-on: ubuntu-latest
steps:
\- uses: anthropics/claude-code-action@v1
with:
anthropic\_api\_key: ${{ secrets.ANTHROPIC\_API\_KEY }}
claude\_args: |
\--max-turns 5
settings: |
{
"permissions": {
"allow": \["Bash", "Read"\],
"deny": \["Write", "WebFetch"\]
}
}
\`\`\`

\### 코드 리뷰 자동화 커맨드

\*\*\`.claude/commands/review.md\`\*\*
\`\`\`markdown
\---
name: review
description: PR 코드 리뷰 수행
allowed-tools: Read, Grep, Glob, Bash(git diff:\*)
\---

현재 브랜치의 변경사항을 리뷰하세요:

1\. \`git diff main...HEAD\` 실행
2\. 보안 취약점 체크 (SQL Injection, XSS, CSRF)
3\. 성능 이슈 식별 (N+1 쿼리, 비효율적 알고리즘)
4\. 코딩 컨벤션 준수 확인
5\. 테스트 커버리지 검토
6\. 80점 이상 신뢰도의 이슈만 보고
\`\`\`

\---

\## 팀 구독 기능과 협업 활용법

\### Claude Code 팀 플랜 핵심 기능

Claude Code 팀 구독은 \*\*Standard Seat($25/월)\*\*과 \*\*Premium Seat($150/월)\*\* 두 가지 옵션을 제공합니다. 팀 플랜의 핵심 협업 기능은 다음과 같습니다:

\- \*\*중앙 집중식 관리\*\*: 시트 할당, 사용량 모니터링, 비용 제어
\- \*\*커넥터 지원\*\*: GitHub, Slack, Google Drive, Gmail 연동
\- \*\*사용량 한도\*\*: 5시간당 약 \*\*225 메시지\*\*, 주간 \*\*50~95시간\*\* Sonnet 4 사용
\- \*\*추가 사용량\*\*: 관리자 설정 한도 내에서 API 요금으로 추가 사용 가능

\### 3인 팀을 위한 워크플로우 패턴

\*\*온보딩 프로세스:\*\*
1\. 저장소 clone ([CLAUDE.md](http://CLAUDE.md), settings.json 자동 포함)
2\. 터미널에서 \`claude\` 실행
3\. \`/init\`으로 환경 확인
4\. 개인 설정 필요시 \`CLAUDE.local.md\` 생성

\*\*병렬 작업 패턴:\*\*
\`\`\`bash
_\# Git worktree로 병렬 Claude 세션 운영_
git worktree add ../rmn-feature-auth feature/auth
git worktree add ../rmn-feature-ads feature/ads
_\# 각 worktree에서 독립적 Claude 세션 실행_
\`\`\`

\---

\## 흔한 실수와 회피 방법

\### 피해야 할 5가지 실수

| 실수 | 문제점 | 해결책 |
|-----|-------|-------|
| [CLAUDE.md](http://CLAUDE.md) 과다 작성 | 명령어 준수율 전반적 저하 | 60~300줄 유지, 점진적 공개 활용 |
| \`/init\` 무검토 사용 | 부정확한 정보 포함 가능 | 항상 검토 후 수정 |
| Claude를 린터로 사용 | 비용 과다, 속도 저하 | Hooks로 gofmt, dart format 자동 실행 |
| 민감 정보 포함 | 보안 위험 | API 키, 비밀번호 절대 포함 금지 |
| 컨텍스트 관리 미흡 | 관련 없는 정보 누적 | 작업 간 \`/clear\` 사용 |

\### Best Practices 요약

\- \*\*CLAUDE.md는 고레버리지 포인트\*\*: 프롬프트처럼 신중하게 작성하고 반복 개선
\- \*\*CODEOWNERS로 보호\*\*: [CLAUDE.md](http://CLAUDE.md) 변경 시 시니어 리뷰 필수화
\- \*\*결정론적 도구 우선\*\*: 린터, 포맷터는 Hooks로, Claude는 창의적 작업에 집중
\- \*\*점진적 공개 활용\*\*: 상세 문서는 \`agent\_docs/\`에 분리하고 CLAUDE.md에서 참조

\---

\## 즉시 적용 가능한 시작 체크리스트

1\. ✅ 프로젝트 루트에 \`CLAUDE.md\` 생성 (60줄 이내)
2\. ✅ \`.claude/settings.json\`에 팀 권한 규칙 설정
3\. ✅ \`.claude/agents/\`에 역할별 subagent 3개 추가
4\. ✅ \`.claude/commands/\`에 핵심 워크플로우 커맨드 생성
5\. ✅ \`.gitignore\`에 \`CLAUDE.local.md\`, \`.claude/settings.local.json\` 추가
6\. ✅ 팀원에게 clone 후 \`claude\` 실행만으로 환경 준비 완료 확인

이 구조를 따르면 3인 팀이 \*\*동일한 Claude Code 환경\*\*에서 RMN SaaS 플랫폼의 \*\*개발 전체 자동화\*\*를 효과적으로 달성할 수 있습니다.