# RMN Backend - Kafka Event Schemas

## Structure
```
proto/
├── buf.yaml           # Buf config
├── buf.gen.yaml       # Code generation
├── common/
│   └── types.proto    # Shared types
└── events/
    ├── auth/v1/
    ├── campaign/v1/
    ├── billing/v1/
    └── ...
```

## Commands
```bash
# Lint protos
buf lint

# Generate Go code
buf generate
```

## Event naming
- Topic: `{domain}.{event}` (e.g., `campaign.created`)
- Message: PascalCase verb (e.g., `CampaignCreated`)

## Adding a new event
1. Create/update .proto file in events/{domain}/v1/
2. Run `buf lint`
3. Run `buf generate`
4. Use generated code in services
