# RMN Frontend (Flutter Web)

## Quick commands
```bash
# Install dependencies
flutter pub get

# Run dev server
flutter run -d chrome

# Run tests
flutter test

# Build for production
flutter build web --release
```

## Architecture
- Feature-First Clean Architecture
- State management: Riverpod / flutter_bloc
- Routing: go_router
- API client: dio

## Directory structure
- `lib/core/` — Config, constants, theme, utils
- `lib/shared/` — Shared widgets, models, services
- `lib/api/` — API client, models, repositories
- `lib/features/` — Feature modules

## Feature module structure
```
features/{feature}/
├── data/
│   ├── models/
│   └── repositories/
├── domain/
│   └── entities/
└── presentation/
    ├── pages/
    ├── widgets/
    └── providers/
```

## Conventions
- Widget naming: PascalCase
- File naming: snake_case
- One widget per file for complex components
