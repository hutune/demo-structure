# How to apoly claude code (concept)

\# RMN 프로젝트 Claude Code 팀 협업 가이드
\## 프론트엔드 / 백엔드+인프라 분리 구조

\---

\## 1. 레포지토리 구조 개요

\`\`\`
GitHub Organization: rmn-platform/
│
├── 📦 rmn-frontend ← 프론트엔드 팀
│ └── Flutter Web 앱
│
├── 📦 rmn-backend ← 백엔드 + 인프라 팀
│ ├── Golang 마이크로서비스
│ ├── Kubernetes 매니페스트
│ └── API 명세 (OpenAPI) ⭐ 소스 오브 트루스
│
└── (선택) 📦 rmn-api-spec ← API 명세 전용 (별도 관리 시)
\`\`\`

\### API 명세 공유 전략

\*\*권장: 백엔드 레포에서 API 명세 관리 → 프론트엔드가 참조\*\*

\`\`\`
\[rmn-backend\] \[rmn-frontend\]
api-spec/openapi.yaml ──────▶ scripts/sync-api.sh
│ (git submodule 또는 다운로드)
▼
GitHub Release / npm package로 배포
\`\`\`

\---

\## 2. 프론트엔드 레포 구조 (rmn-frontend)

\`\`\`
rmn-frontend/
├── [CLAUDE.md](http://CLAUDE.md) # 프론트엔드 컨텍스트
├── .claude/
│ ├── settings.json # 팀 공유 설정
│ ├── settings.local.json # 개인 설정 (gitignore)
│ └── commands/
│ ├── [feature-create.md](http://feature-create.md)
│ ├── [widget-create.md](http://widget-create.md)
│ ├── [api-sync.md](http://api-sync.md) # API 명세 동기화
│ └── [pr-create.md](http://pr-create.md)
│
├── lib/
│ ├── api/
│ │ ├── generated/ # OpenAPI에서 생성 (수정 금지)
│ │ └── client.dart
│ ├── features/
│ │ ├── campaign/
│ │ ├── dashboard/
│ │ └── auth/
│ ├── shared/
│ │ ├── widgets/
│ │ └── utils/
│ └── main.dart
│
├── api-spec/ # 백엔드에서 동기화
│ └── openapi.yaml # (submodule 또는 다운로드)
│
├── scripts/
│ ├── [sync-api.sh](http://sync-api.sh) # API 명세 동기화 스크립트
│ └── [gen-api-client.sh](http://gen-api-client.sh) # Dart 클라이언트 생성
│
├── pubspec.yaml
└── Makefile
\`\`\`

\### 2.1 프론트엔드 [CLAUDE.md](http://CLAUDE.md)

\`\`\`markdown
\# RMN Frontend (Flutter)

\## 프로젝트 개요
RMN 광고 관리 플랫폼 웹 프론트엔드

\## 기술 스택
\- Flutter 3.x / Dart
\- 상태관리: flutter\_bloc
\- API 클라이언트: dio + openapi-generator
\- 라우팅: go\_router
\- 테스트: flutter\_test, mockito

\## 백엔드 연동
\- API 명세: api-spec/openapi.yaml (백엔드 레포에서 동기화)
\- API 클라이언트: lib/api/generated/ (자동 생성, 수정 금지)
\- 환경별 엔드포인트:
\- dev: [https://api-dev.rmn-platform.com](https://api-dev.rmn-platform.com)
\- stg: [https://api-stg.rmn-platform.com](https://api-stg.rmn-platform.com)
\- prd: [https://api.rmn-platform.com](https://api.rmn-platform.com)

\## 디렉토리 구조
\- lib/features/\[기능\]/ → 기능별 모듈 (BLoC, pages, widgets)
\- lib/shared/ → 공통 위젯, 유틸리티
\- lib/api/generated/ → 자동 생성 코드 (수정 금지!)

\## 주요 명령어
\`\`\`bash
_\# 개발 서버_
flutter run -d chrome

_\# API 명세 동기화 (백엔드 레포에서)_
make sync-api

_\# API 클라이언트 재생성_
make gen-api

_\# 테스트_
flutter test

_\# 빌드_
flutter build web --dart-define=ENV=dev
\`\`\`

\## 코딩 컨벤션
\- BLoC 패턴 필수: Event → BLoC → State
\- 위젯 분리: 페이지 > 섹션 > 컴포넌트
\- API 호출: Repository 패턴 사용
\- 에러 처리: Either<Failure, Success> 패턴

\## 브랜치 규칙
\- feature/\[기능명\] (예: feature/campaign-list)
\- PR 제목: 한글 또는 영문으로 명확하게

\## 주의사항
\- lib/api/generated/ 파일 직접 수정 금지
\- API 변경 필요시 백엔드 팀에 요청 → 명세 업데이트 후 sync
\`\`\`

\### 2.2 프론트엔드 settings.json

\`\`\`json
{
"model": "claude-sonnet-4-20250514",
"permissions": {
"allow": \[
"Read",
"Write(lib/\*\*)",
"Write(test/\*\*)",
"Bash(flutter:\*)",
"Bash(dart:\*)",
"Bash(make:\*)",
"Bash(git add:\*)",
"Bash(git commit:\*)",
"Bash(git push:\*)",
"Bash(git checkout:\*)",
"Bash(git branch:\*)",
"Bash(gh pr:\*)"
\],
"deny": \[
"Read(.env\*)",
"Write(.env\*)",
"Write(lib/api/generated/\*\*)",
"Bash(rm -rf:\*)",
"Bash(sudo:\*)"
\]
},
"hooks": {
"PostToolUse": \[
{
"matcher": "Write(lib/\*\*/\*.dart)",
"hooks": \[
{
"type": "command",
"command": "dart format $CLAUDE\_FILE\_PATH"
}
\]
}
\]
}
}
\`\`\`

\### 2.3 프론트엔드 슬래시 명령어

\*\*.claude/commands/feature-create.md\*\*
\`\`\`markdown
\# Feature 모듈 생성

새로운 기능 모듈을 Flutter 표준 구조로 생성합니다.

\## 생성할 구조
\`\`\`
lib/features/\[기능명\]/
├── bloc/
│ ├── \[기능\]\_bloc.dart
│ ├── \[기능\]\_event.dart
│ └── \[기능\]\_state.dart
├── pages/
│ └── \[기능\]\_page.dart
├── widgets/
│ └── .gitkeep
└── repository/
└── \[기능\]\_repository.dart
\`\`\`

\## 절차
1\. 위 구조대로 파일 생성
2\. BLoC 보일러플레이트 코드 작성
3\. go\_router에 라우트 추가
4\. 기본 테스트 파일 생성 (test/features/\[기능\]/)

\## 기능명
$ARGUMENTS
\`\`\`

\*\*.claude/commands/api-sync.md\*\*
\`\`\`markdown
\# API 명세 동기화

백엔드 레포에서 최신 API 명세를 가져와 클라이언트를 재생성합니다.

\## 절차
1\. \`make sync-api\` 실행 (백엔드 레포에서 openapi.yaml 다운로드)
2\. 변경사항 확인 및 요약
3\. \`make gen-api\` 실행 (Dart 클라이언트 재생성)
4\. 생성된 클라이언트의 주요 변경점 설명
5\. 영향받는 Repository/BLoC 목록 제시

\## 실행
$ARGUMENTS
\`\`\`

\---

\## 3. 백엔드+인프라 레포 구조 (rmn-backend)

\`\`\`
rmn-backend/
├── [CLAUDE.md](http://CLAUDE.md) # 백엔드 전체 컨텍스트
├── .claude/
│ ├── settings.json
│ ├── settings.local.json
│ └── commands/
│ ├── [service-create.md](http://service-create.md)
│ ├── [event-create.md](http://event-create.md)
│ ├── [api-add.md](http://api-add.md)
│ ├── [deploy.md](http://deploy.md)
│ └── [pr-create.md](http://pr-create.md)
│
├── api-spec/ # ⭐ API 명세 (소스 오브 트루스)
│ ├── [CLAUDE.md](http://CLAUDE.md)
│ └── openapi.yaml
│
├── proto/ # Kafka 이벤트 스키마
│ ├── [CLAUDE.md](http://CLAUDE.md)
│ ├── buf.yaml
│ └── events/
│ ├── campaign/v1/
│ ├── user/v1/
│ └── billing/v1/
│
├── services/ # 마이크로서비스
│ ├── [CLAUDE.md](http://CLAUDE.md) # 백엔드 공통 컨텍스트
│ ├── api-gateway/
│ │ ├── [CLAUDE.md](http://CLAUDE.md)
│ │ ├── cmd/
│ │ ├── internal/
│ │ ├── go.mod
│ │ └── Dockerfile
│ ├── user-service/
│ ├── campaign-service/
│ ├── billing-service/
│ └── device-service/
│
├── pkg/ # 공유 라이브러리
│ ├── kafka/
│ ├── middleware/
│ └── config/
│
├── infrastructure/ # 인프라 코드
│ ├── [CLAUDE.md](http://CLAUDE.md)
│ ├── terraform/ # (선택) 클라우드 리소스
│ ├── helm/ # 공통 Helm 차트
│ │ ├── base/
│ │ └── charts/
│ └── argocd/
│ └── appset.yaml
│
├── .github/
│ └── workflows/
│ ├── api-gateway.yaml
│ ├── user-service.yaml
│ └── publish-api-spec.yaml # API 명세 배포
│
├── [go.work](http://go.work)
├── Makefile
└── docker-compose.yaml # 로컬 개발 환경
\`\`\`

\### 3.1 백엔드 루트 [CLAUDE.md](http://CLAUDE.md)

\`\`\`markdown
\# RMN Backend + Infrastructure

\## 프로젝트 개요
RMN 광고 관리 플랫폼 백엔드 및 인프라

\## 아키텍처
\- 마이크로서비스: Golang
\- 메시지 큐: Kafka (이벤트 드리븐)
\- 데이터베이스: PostgreSQL
\- 캐시: Redis
\- 인프라: Kubernetes (ArgoCD GitOps)

\## 환경
\- dev: 개발 환경
\- stg: 스테이징 환경
\- prd: 프로덕션 환경

\## 디렉토리 구조
\- api-spec/ → OpenAPI 명세 (프론트엔드 공유용)
\- proto/ → Kafka 이벤트 스키마 (Protobuf)
\- services/ → 마이크로서비스들
\- pkg/ → 공유 Go 라이브러리
\- infrastructure/ → K8s, Terraform, ArgoCD

\## 서비스 목록
| 서비스 | 포트 | 역할 |
|-------|------|-----|
| api-gateway | 8080 | 외부 API 진입점 |
| user-service | 8081 | 사용자/인증 |
| campaign-service | 8082 | 캠페인 관리 |
| billing-service | 8083 | 과금/정산 |
| device-service | 8084 | 사이니지 장치 |

\## 주요 명령어
\`\`\`bash
_\# 로컬 환경 실행 (Kafka, PostgreSQL, Redis)_
docker-compose up -d

_\# 서비스 실행_
cd services/\[서비스명\] && go run ./cmd/...

_\# 전체 테스트_
make test

_\# Proto 생성_
buf generate

_\# API 명세에서 Go 코드 생성_
make gen-api

_\# 린트_
make lint
\`\`\`

\## 브랜치 규칙
\- feature/\[서비스\]-\[기능\] (예: feature/campaign-create-api)
\- infra/\[변경내용\] (예: infra/add-redis-cluster)

\## API 명세 배포
api-spec/openapi.yaml 변경 시 자동으로 GitHub Release 생성
프론트엔드 팀이 최신 버전 동기화 가능
\`\`\`

\### 3.2 서비스 공통 [CLAUDE.md](http://CLAUDE.md) (services/CLAUDE.md)

\`\`\`markdown
\# Backend Services

\## 기술 스택
\- Go 1.23+
\- HTTP: chi router
\- gRPC: grpc-go
\- Kafka: segmentio/kafka-go
\- DB: pgx v5
\- 설정: viper
\- 로깅: zerolog

\## 서비스 표준 구조
\`\`\`
services/\[서비스명\]/
├── cmd/\[서비스명\]/
│ └── main.go
├── internal/
│ ├── handler/ # HTTP/gRPC 핸들러
│ ├── service/ # 비즈니스 로직
│ ├── repository/ # 데이터 접근
│ ├── event/ # Kafka 발행/구독
│ └── domain/ # 도메인 모델
├── go.mod
├── Dockerfile
└── chart/ # Helm 차트
├── values.yaml
├── values-dev.yaml
├── values-stg.yaml
└── values-prd.yaml
\`\`\`

\## 코딩 컨벤션
\- 표준 Go 프로젝트 레이아웃 준수
\- 에러: fmt.Errorf("context: %w", err)
\- 로깅: zerolog 사용, 구조화된 로그
\- 테스트: 테이블 드리븐 테스트

\## Kafka 컨벤션
\- 토픽 네이밍: \[도메인\].\[이벤트\] (예: campaign.created)
\- 이벤트 스키마: proto/events/\[도메인\]/v1/
\- Consumer Group: \[서비스명\]-\[환경\]

\## 서비스 간 통신
\- 동기: gRPC (서비스 간 직접 호출)
\- 비동기: Kafka (이벤트 발행/구독)
\- 외부: REST API (api-gateway 경유)
\`\`\`

\### 3.3 인프라 [CLAUDE.md](http://CLAUDE.md) (infrastructure/CLAUDE.md)

\`\`\`markdown
\# Infrastructure

\## Kubernetes 구조
\- ArgoCD로 GitOps 배포
\- Helm 차트로 서비스 배포
\- 환경별 values 파일 분리

\## 환경
| 환경 | 클러스터 | 네임스페이스 |
|-----|---------|------------|
| dev | rmn-dev | rmn-dev |
| stg | rmn-stg | rmn-stg |
| prd | rmn-prd | rmn-prd |

\## ArgoCD ApplicationSet
services/\*/chart/ 디렉토리 자동 감지

\## 명령어
\`\`\`bash
_\# 로컬 K8s (kind) 설정_
make setup-local-k8s

_\# Helm 차트 린트_
helm lint services/\*/chart/

_\# 매니페스트 미리보기_
helm template services/api-gateway/chart -f services/api-gateway/chart/values-dev.yaml
\`\`\`

\## 주의사항
\- prd 환경 직접 kubectl 금지 → ArgoCD 통해서만
\- 시크릿은 External Secrets Operator 사용
\- values-prd.yaml에 민감정보 절대 금지
\`\`\`

\### 3.4 백엔드 settings.json

\`\`\`json
{
"model": "claude-sonnet-4-20250514",
"permissions": {
"allow": \[
"Read",
"Write(services/\*\*)",
"Write(pkg/\*\*)",
"Write(proto/\*\*)",
"Write(api-spec/\*\*)",
"Write(infrastructure/\*\*)",
"Bash(go:\*)",
"Bash(buf:\*)",
"Bash(make:\*)",
"Bash(docker:\*)",
"Bash(docker-compose:\*)",
"Bash(helm lint:\*)",
"Bash(helm template:\*)",
"Bash(git:\*)",
"Bash(gh pr:\*)"
\],
"deny": \[
"Read(.env\*)",
"Read(\*\*/secrets/\*\*)",
"Read(\*\*/\*secret\*)",
"Write(.env\*)",
"Write(\*\*/secrets/\*\*)",
"Bash(rm -rf:\*)",
"Bash(sudo:\*)",
"Bash(kubectl delete:\*)",
"Bash(kubectl apply:\*)",
"Bash(helm install:\*)",
"Bash(helm upgrade:\*)"
\]
},
"hooks": {
"PostToolUse": \[
{
"matcher": "Write(services/\*\*/\*.go)",
"hooks": \[
{
"type": "command",
"command": "gofmt -w $CLAUDE\_FILE\_PATH"
}
\]
}
\]
}
}
\`\`\`

\### 3.5 백엔드 슬래시 명령어

\*\*.claude/commands/api-add.md\*\*
\`\`\`markdown
\# API 엔드포인트 추가

OpenAPI 명세에 새 엔드포인트를 추가하고 관련 코드를 생성합니다.

\## 절차
1\. api-spec/openapi.yaml 읽기
2\. 요청된 엔드포인트 추가 (paths, schemas)
3\. \`make gen-api\` 실행 (Go 코드 생성)
4\. api-gateway에 핸들러 스텁 생성
5\. 변경사항 요약 및 프론트엔드 팀 알림 메시지 생성

\## API 설명
$ARGUMENTS
\`\`\`

\*\*.claude/commands/service-create.md\*\*
\`\`\`markdown
\# 마이크로서비스 생성

새로운 Golang 마이크로서비스를 표준 구조로 생성합니다.

\## 생성할 구조
\`\`\`
services/\[서비스명\]/
├── cmd/\[서비스명\]/main.go
├── internal/
│ ├── handler/
│ ├── service/
│ ├── repository/
│ ├── event/
│ └── domain/
├── go.mod
├── Dockerfile
├── [CLAUDE.md](http://CLAUDE.md)
└── chart/
├── Chart.yaml
├── values.yaml
├── values-dev.yaml
├── values-stg.yaml
├── values-prd.yaml
└── templates/
\`\`\`

\## 절차
1\. 위 구조대로 파일 생성
2\. go.work에 서비스 추가
3\. Makefile에 서비스 타겟 추가
4\. GitHub Actions 워크플로우 생성
5\. 서비스 [CLAUDE.md](http://CLAUDE.md) 작성

\## 서비스명과 설명
$ARGUMENTS
\`\`\`

\*\*.claude/commands/event-create.md\*\*
\`\`\`markdown
\# Kafka 이벤트 생성

Protobuf 스키마로 새 Kafka 이벤트를 정의합니다.

\## 절차
1\. proto/events/\[도메인\]/v1/ 에 .proto 파일 생성/수정
2\. 이벤트 메시지 정의
3\. \`buf lint\` 실행
4\. \`buf generate\` 실행
5\. 발행 서비스에 이벤트 발행 코드 추가
6\. (선택) 구독 서비스에 핸들러 추가

\## 이벤트 설명
$ARGUMENTS
\`\`\`

\---

\## 4. 팀 간 협업 워크플로우

\### 4.1 API 변경 프로세스

\`\`\`
┌─────────────────────────────────────────────────────────────┐
│ API 변경 워크플로우 │
└─────────────────────────────────────────────────────────────┘

\[1\] 프론트엔드 요청
└─▶ Slack/Issue: "캠페인 목록 API에 필터 파라미터 추가 필요"

\[2\] 백엔드 작업 (rmn-backend)
├─▶ /api-add GET /campaigns에 status, date\_from 쿼리 파라미터 추가
├─▶ api-gateway 핸들러 구현
├─▶ campaign-service 로직 구현
├─▶ PR 생성 및 리뷰
└─▶ main 머지 → API 명세 자동 배포

\[3\] 프론트엔드 작업 (rmn-frontend)
├─▶ /api-sync (최신 명세 동기화)
├─▶ Repository에서 새 파라미터 사용
├─▶ BLoC/UI 업데이트
└─▶ PR 생성 및 머지
\`\`\`

\### 4.2 API 명세 동기화 방법

\*\*옵션 A: GitHub Release 다운로드 (권장)\*\*

\`\`\`bash
_\# rmn-frontend/scripts/sync-api.sh_
_#!/bin/bash_
LATEST=$(gh release view --repo rmn-platform/rmn-backend --json tagName -q .tagName)
curl -L "[https://github.com/rmn-platform/rmn-backend/releases/download/$](https://github.com/rmn-platform/rmn-backend/releases/download/$){LATEST}/openapi.yaml" \\
\-o api-spec/openapi.yaml
echo "Synced to version: ${LATEST}"
\`\`\`

\*\*옵션 B: Git Submodule\*\*

\`\`\`bash
_\# rmn-frontend에서_
git submodule add [https://github.com/rmn-platform/rmn-backend.git](https://github.com/rmn-platform/rmn-backend.git) external/backend
_\# api-spec → external/backend/api-spec 심볼릭 링크_
\`\`\`

\### 4.3 멀티레포에서 Claude Code 컨텍스트 공유

\*\*방법: --add-dir 옵션 사용\*\*

프론트엔드 개발자가 백엔드 API 명세를 참조해야 할 때:

\`\`\`bash
_\# 두 레포를 모두 클론한 상태에서_
cd ~/projects/rmn-frontend
claude --add-dir ~/projects/rmn-backend/api-spec

_\# 이제 백엔드 API 명세를 참조하며 프론트 작업 가능_
\> api-spec/openapi.yaml을 보고 캠페인 Repository 구현해줘
\`\`\`

\---

\## 5. 팀별 온보딩

\### 5.1 프론트엔드 팀 온보딩

\`\`\`bash
_\# 1. 레포 클론_
git clone [https://github.com/rmn-platform/rmn-frontend.git](https://github.com/rmn-platform/rmn-frontend.git)
cd rmn-frontend

_\# 2. 의존성 설치_
flutter pub get

_\# 3. API 명세 동기화_
make sync-api
make gen-api

_\# 4. 개발 서버 실행_
flutter run -d chrome

_\# 5. Claude Code 시작_
claude
\> /help _\# 사용 가능한 명령어 확인_
\`\`\`

\### 5.2 백엔드 팀 온보딩

\`\`\`bash
_\# 1. 레포 클론_
git clone [https://github.com/rmn-platform/rmn-backend.git](https://github.com/rmn-platform/rmn-backend.git)
cd rmn-backend

_\# 2. 로컬 인프라 실행_
docker-compose up -d _\# Kafka, PostgreSQL, Redis_

_\# 3. 의존성 및 코드 생성_
go work sync
buf generate
make gen-api

_\# 4. 서비스 실행_
cd services/api-gateway && go run ./cmd/api-gateway

_\# 5. Claude Code 시작_
claude
\> /help
\`\`\`

\---

\## 6. 환경별 설정 요약

\### 프론트엔드 (rmn-frontend)

| 파일 | Git | 용도 |
|-----|-----|------|
| [CLAUDE.md](http://CLAUDE.md) | ✅ | 프로젝트 컨텍스트 |
| .claude/settings.json | ✅ | 팀 공유 설정 |
| .claude/settings.local.json | ❌ | 개인 설정 |
| .claude/commands/\*.md | ✅ | 팀 공유 명령어 |
| lib/api/generated/ | ✅ | 생성 코드 (수정 금지) |

\### 백엔드 (rmn-backend)

| 파일 | Git | 용도 |
|-----|-----|------|
| [CLAUDE.md](http://CLAUDE.md) | ✅ | 전체 컨텍스트 |
| services/CLAUDE.md | ✅ | 백엔드 공통 |
| services/\*/CLAUDE.md | ✅ | 서비스별 상세 |
| infrastructure/CLAUDE.md | ✅ | 인프라 컨텍스트 |
| .claude/settings.json | ✅ | 팀 공유 설정 |
| .claude/settings.local.json | ❌ | 개인 설정 |

\---

\## 7. 커뮤니케이션 채널

\`\`\`
┌─────────────────────────────────────────────────────────┐
│ API 변경 요청 │
│ ──────────── │
│ Slack: #rmn-api-changes │
│ GitHub Issue: rmn-backend 레포에 "api-request" 라벨 │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│ API 변경 완료 알림 │
│ ──────────────── │
│ GitHub Actions → Slack 자동 알림 │
│ "🚀 API v1.2.3 배포 완료 - 프론트 sync 필요" │
└─────────────────────────────────────────────────────────┘
\`\`\`

\---

\## 8. 요약 비교

| 항목 | 프론트엔드 (rmn-frontend) | 백엔드 (rmn-backend) |
|-----|-------------------------|---------------------|
| 기술 | Flutter/Dart | Golang + K8s |
| API 명세 | 동기화해서 사용 | 소스 오브 트루스 |
| Claude 컨텍스트 | FE 전용 최적화 | BE + 인프라 통합 |
| 배포 | CDN/S3 | ArgoCD GitOps |
| 주요 명령어 | /feature-create, /api-sync | /service-create, /api-add |

이 구조로 각 팀이 \*\*독립적으로 작업\*\*하면서도 \*\*API 명세를 통해 명확히 협업\*\*할 수 있습니다!