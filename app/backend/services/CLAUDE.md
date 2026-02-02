# RMN Backend Services

## Standard service layout
```
services/{service-name}/
├── main.go                    # Entry point
├── go.mod                     # Dependencies
├── Dockerfile
├── config/
│   └── app.development.yaml
├── internal/
│   ├── app/
│   │   ├── server.go          # HTTP server
│   │   └── routes/
│   ├── common/
│   │   ├── app.go             # App context
│   │   └── errors/
│   ├── dto/                   # Data Transfer Objects
│   ├── models/                # Database models
│   ├── repositories/          # Data access
│   ├── logic/                 # Business logic
│   ├── services/              # DI container
│   ├── presentation/
│   │   ├── handlers/          # HTTP handlers
│   │   └── middlewares/
│   ├── events/                # Kafka handlers
│   └── migrations/
└── chart/                     # Helm chart
```

## Coding conventions
- Always use `context.Context` as first parameter
- Errors: `fmt.Errorf("context: %w", err)`
- Logging: zerolog, structured logs
- Tests: table-driven tests

## Layer responsibilities
- **Handler**: HTTP request/response, validation
- **Logic**: Business rules, orchestration
- **Repository**: Database operations only

## Creating a new service
1. Copy from existing service
2. Update module name in go.mod
3. Add to go.work
4. Create Dockerfile
5. Create Helm chart
6. Add GitHub Actions workflow
