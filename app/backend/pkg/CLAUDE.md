# RMN Backend - Shared Libraries (pkg/)

## Module
`rmn-backend/pkg`

## Packages

| Package | Purpose |
|---------|---------|
| `config/` | Viper-based config loading |
| `database/` | CockroachDB connection, Goose migrations |
| `redis/` | Redis client wrapper |
| `kafka/` | Kafka publisher/subscriber |
| `httpclient/` | Internal HTTP client for service-to-service |
| `middleware/` | Common HTTP middlewares |
| `token/` | PASETO token management |
| `errorx/` | Error types and codes |
| `logger/` | Zerolog wrapper |
| `httpserver/` | HTTP server setup |
| `utils/` | Utility functions |

## Usage in services
```go
import (
    "rmn-backend/pkg/config"
    "rmn-backend/pkg/database"
    "rmn-backend/pkg/logger"
)
```

## Adding a new package
1. Create directory under pkg/
2. Add Go files
3. Run `go work sync` from backend root
4. Import in services
