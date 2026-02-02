# RMN Backend Monorepo

## Quick commands
```bash
# Build all services
make build

# Test all services
make test

# Run a specific service
make run SERVICE=api-gateway

# Start local infrastructure
make infra-up

# Generate Kafka event code
make proto
```

## Architecture
- External API: HTTP REST via API Gateway (port 8080)
- Service-to-Service (sync): HTTP REST via pkg/httpclient
- Service-to-Service (async): Kafka events (see proto/events/)
- Database: CockroachDB (port 26257)
- Cache: Redis (port 6379)
- Storage: MinIO (port 9000)

## Directory structure
- `proto/` — Kafka event schemas (Protobuf)
- `generated/` — Generated code from proto
- `pkg/` — Shared Go libraries
- `services/` — Microservices
- `infrastructure/` — K8s manifests
- `argocd/` — GitOps config

## Services
| Service | Port | Description |
|---------|------|-------------|
| api-gateway | 8080 | API Gateway |
| auth-service | 8081 | Authentication |
| user-service | 8082 | User management |

## Conventions
- Module naming: `rmn-backend/services/{service-name}`
- Event naming: PascalCase verb (CampaignCreated, UserUpdated)
- API versioning: /api/v1/, /api/v2/
- Layer architecture: Handler → Logic → Repository

## Communication
- NO gRPC in this project
- Sync calls: HTTP REST only
- Async events: Kafka with Protobuf schemas
