# RMN Platform - Demo Structure

## Project overview
RMN (Retail Media Network) advertising management platform.

## Directory structure
```
demo-structure/
├── app/
│   ├── backend/     # Golang microservices monorepo
│   └── fe/          # Flutter Web frontend
├── _bmad/           # BMAD framework
├── _bmad-output/    # Planning artifacts (PRD, Architecture)
├── docs/            # Documentation
└── backend/         # [LEGACY] Original backend code
```

## Quick commands

### Backend
```bash
cd app/backend
make build          # Build all services
make test           # Test all services
make run SERVICE=x  # Run specific service
make infra-up       # Start local infrastructure
```

### Frontend
```bash
cd app/fe
flutter pub get     # Install dependencies
flutter run -d chrome   # Run dev server
flutter test        # Run tests
```

## Architecture
- Backend: Go microservices with Kafka events
- Frontend: Flutter Web
- Database: CockroachDB
- Cache: Redis
- Events: Kafka
- Storage: MinIO (S3 compatible)

## Communication Patterns
- Client → Backend: HTTP REST via API Gateway
- Service ↔ Service (sync): HTTP REST
- Service ↔ Service (async): Kafka Events
- NO gRPC in this project

## Documentation
- PRD: `_bmad-output/planning-artifacts/prd.md`
- Architecture: `_bmad-output/planning-artifacts/architecture-monorepo.md`
- Epics: `_bmad-output/planning-artifacts/epics/`
