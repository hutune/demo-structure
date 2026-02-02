# Mono repo Golang Microservices Team

For Kafka-based Golang microservices developed with Kubernetes and Claude Code, **monorepo is the clear answer**. Reduced coordination overhead, simplified Kafka schema sharing, AI-driven development efficiency, and the success of companies like Uber and Monzo running thousands of Go microservices in a monorepo all support this. In contrast, Segment’s multi-repo microservices failure is a cautionary tale.

Independent deployment and loose coupling are fully achievable in a monorepo with the right structure. Key insight: **repository layout does not determine service coupling; code design does.**

---

## Why Monorepo Favors Small Go Teams

Expert opinion for 1–2 developer teams is clear. Ricardo Soares (LinkedIn) says: "When the company is very small, the monorepo approach is the most dynamic. Multi-repo is generally much more costly for small teams." Earthly’s Brandon Schurman agrees: "Monorepo is especially attractive when working with small teams."

**Go workspaces (go.work) change the game.** Workspaces introduced in Go 1.18 remove the friction of managing multiple modules in a monorepo. Each service keeps its own `go.mod`, but with `go.work` you can develop all services at once. Shared library changes apply immediately, no `replace` directives during development, and gopls supports cross-module navigation and refactoring.

For small teams, the real issue with multi-repo is **coordination overhead**. In multi-repo, a change spanning several services means: update library repo → PR → merge → tag release → bump dependency in service-1 → PR → merge → repeat for each service. In a monorepo? A single atomic commit. Segment felt this painfully—with 120+ repos, "three full-time engineers spent most of their time just keeping the system alive."

| Factor | Monorepo (1–2 people) | Multi-repo (1–2 people) |
|--------|------------------------|--------------------------|
| Cross-service refactoring | Single commit | Multiple coordinated PRs |
| Shared library update | Immediate | Tag → release → bump in each consumer |
| CI/CD maintenance | One pipeline | N pipelines to manage |
| Code review context | Full visibility | Split across repos |
| Cognitive load | One mental model | Track which repo has what |

---

## Kafka Event Schema Sharing: Monorepo Simplifies Everything

In Kafka-based event-driven architecture, schema management is central—and much simpler in a monorepo. If services share event contracts, they should share schema definitions too.

**Recommended approach: Protobuf + Buf CLI in the monorepo**

```
monorepo/
├── proto/
│   ├── buf.yaml       # Buf config
│   ├── buf.gen.yaml   # Code generation config
│   ├── common/
│   │   └── types.proto    # Shared types (timestamp, ID)
│   └── events/
│       ├── order/v1/events.proto
│       └── user/v1/events.proto
├── generated/
│   └── go/            # Generated Go code
├── services/
│   └── [all services import generated/go]
```

Running `buf generate` produces Go structs. All services import `github.com/yourorg/monorepo/generated/go/events`. Schema and consumer code update atomically in one commit—no version coordination.

For **runtime validation**, use Confluent Schema Registry in **BACKWARD** compatibility mode. Main production settings:

- `auto.register.schemas=false` (disable auto-registration)
- `normalize.schemas=true` (enable schema normalization)
- Use TopicNameStrategy for topic naming: `{topic}-value`

The multi-repo alternative—a dedicated schema repo published as a Go module—adds significant overhead: after a schema change you publish the library, then bump the version in every consumer. That’s the pattern that contributed to Segment’s microservices collapse.

---

## Achieving Independent Deployment in a Monorepo

Per-service independent deployment is fully possible in a monorepo using **path-based triggers** in GitHub Actions. This is not a compromise—it’s the same pattern used by Uber and Monzo.

**Per-service GitHub Actions workflow:**

```yaml
name: user-service
on:
  push:
    paths:
      - 'services/user-service/**'
      - 'pkg/**'              # Shared libraries
      - 'proto/events/user/**' # Service-specific schema
    branches: [main]

jobs:
  build-deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.23'
          cache-dependency-path: 'services/user-service/go.sum'
      - name: Test
        run: cd services/user-service && go test ./...
      - name: Build and push
        run: |
          docker build -t ghcr.io/${{ github.repository }}/user-service:${{ github.sha }} \
            -f services/user-service/Dockerfile .
          docker push ghcr.io/${{ github.repository }}/user-service:${{ github.sha }}
```

For small teams with fewer than 10 services, **a custom script using `git diff`** to detect affected services is enough—no Bazel or Nx. As services grow, generate workflow files from a template.

---

## Organizing Kubernetes Manifests with GitOps

**Recommended structure: per-service Helm charts + ArgoCD ApplicationSets**

```
monorepo/
├── services/
│   ├── user-service/
│   │   ├── src/
│   │   └── chart/           # Per-service Helm chart
│   │       ├── Chart.yaml
│   │       ├── values.yaml  # Defaults
│   │       ├── values-dev.yaml  # Dev overrides
│   │       ├── values-stg.yaml  # Staging overrides
│   │       ├── values-prd.yaml  # Production overrides
│   │       └── templates/
├── infrastructure/
│   ├── base/                # Shared K8s resources
│   └── overlays/
│       ├── dev/
│       ├── stg/
│       └── prd/
└── argocd/
    └── appset.yaml          # ApplicationSet for all services
```

**ArgoCD ApplicationSet** discovers services dynamically:

```yaml
apiVersion: argoproj.io/v1alpha1
kind: ApplicationSet
metadata:
  name: microservices
spec:
  goTemplate: true
  generators:
    - git:
        repoURL: https://github.com/org/monorepo.git
        revision: HEAD
        directories:
          - path: services/*/chart
  template:
    metadata:
      name: '{{index .path.segments 1}}-{{.values.env}}'
    spec:
      project: default
      source:
        repoURL: https://github.com/org/monorepo.git
        path: '{{.path.path}}'
        helm:
          valueFiles:
            - values-{{.values.env}}.yaml
      destination:
        server: https://kubernetes.default.svc
        namespace: '{{index .path.segments 1}}'
```

With this pattern, adding a new service is just creating a directory that includes a chart—ArgoCD picks it up automatically.

---

## Real-World Validation from Go-Centric Companies

Companies running thousands of Go microservices successfully all use a monorepo:

**Uber**: 3,000+ microservices, ~50 million lines of Go, 1,000+ commits per day in a Go monorepo. They migrated from multi-repo to monorepo in 2018 due to dependency management issues.

**Monzo**: 2,800+ microservices, all in Go, all in one monorepo. Platform engineer Will Sewell: "Having a monorepo is incredibly helpful for managing some of the challenges generally associated with microservices." He highlights consistent dependency versions, compile-time feedback across services, and large-scale refactoring capability.

**Segment’s lesson**: They started as a monolith, moved to microservices across 120+ separate repos, then **returned to a monolith in 2017**. Engineers described "shared library hell"—different versions per service, testing nightmares, operational overhead scaling with the number of services. "Instead of moving faster, a small team was swallowed by exploding complexity."

The pattern is clear: companies that succeed at microservices at scale use a monorepo. Those that struggled with multi-repo microservices either switched back or invested heavily in infra—not feasible for a 1–2 person team.

---

## Claude Code Works Much Better in a Monorepo

In AI-assisted development **context is decisive**—and a monorepo gives context advantages that a multi-repo setup can’t match without substantial engineering.

**Monorepo advantages for Claude Code:**
- Full visibility of how services interact with shared libraries
- A single hierarchical CLAUDE.md structure provides full context
- Cross-service refactoring understood with full context
- One developer reduced CLAUDE.md from **47k to 9k words** using the monorepo’s hierarchy

**Multi-repo issues**: Teams report **40–60% of token budget** spent on cross-repo context duplication. When Claude only sees one repo, it "has no idea what the API of another service looks like."

**Recommended CLAUDE.md structure:**

```
monorepo/
├── CLAUDE.md              # Root: architecture overview, common commands
├── services/
│   ├── user-service/
│   │   └── CLAUDE.md      # Per-service: API contract, dependencies
│   └── order-service/
│       └── CLAUDE.md
├── proto/
│   └── CLAUDE.md          # Schema conventions, generation commands
└── pkg/
    └── CLAUDE.md          # Shared library patterns
```

Example root CLAUDE.md content:

```markdown
# Platform Monorepo

## Quick commands
- Full build: `make build`
- Service test: `cd services/<name> && go test ./...`
- Proto generation: `buf generate`
- Staging deploy: `make deploy-stg SERVICE=<name>`

## Architecture
- Services communicate via Kafka events (see proto/events/)
- Shared libraries live in pkg/ — import as internal packages
- Each service has its own Helm chart under chart/

## Conventions
- Event naming: PascalCase verb (OrderCreated, UserUpdated)
- Service directory name matches Kubernetes namespace
```

---

## Recommended Monorepo Structure for RMN Project

Based on the research above, a structure tuned for Golang microservices + Flutter + Kafka + Kubernetes:

```
platform/
├── CLAUDE.md              # AI context: architecture, commands
├── go.work                # Go workspace (gitignore)
├── buf.yaml               # Protobuf config
├── buf.gen.yaml
│
├── proto/                 # Event schemas
│   └── events/
│       ├── order/v1/
│       └── user/v1/
├── generated/             # Generated code (committed)
│   └── go/
│
├── services/              # Go microservices
│   ├── api-gateway/
│   │   ├── go.mod
│   │   ├── main.go
│   │   ├── Dockerfile
│   │   └── chart/         # Helm chart
│   │       ├── Chart.yaml
│   │       ├── values.yaml
│   │       ├── values-dev.yaml
│   │       ├── values-stg.yaml
│   │       └── values-prd.yaml
│   ├── user-service/
│   └── order-service/
│
├── pkg/                   # Shared Go libraries
│   ├── kafka/             # Kafka client wrapper
│   ├── middleware/        # Common HTTP/gRPC middleware
│   └── config/            # Config loading
│
├── web/                   # Flutter web
│   └── flutter_app/
│
├── infrastructure/        # Shared K8s resources
│   ├── kafka/
│   └── monitoring/
│
├── argocd/                # GitOps config
│   └── appset.yaml
│
├── .github/
│   └── workflows/
│       ├── user-service.yaml   # Per-service CI
│       ├── order-service.yaml
│       └── flutter-web.yaml
│
└── Makefile               # Common commands
```

---

## Summary: Decision Matrix by Requirement

| Requirement | Monorepo solution | What multi-repo needs |
|-------------|-------------------|------------------------|
| **1–2 developers** | Natural fit — minimal coordination | Significant overhead managing repos |
| **Independent deployment** | Path-based GitHub Actions triggers | Already independent (but more pipelines) |
| **Loose coupling** | Determined by code design, not repo layout | Same discipline required |
| **Kafka schema sharing** | Single proto/ dir, atomic updates | Dedicated schema repo + versioning |
| **K8s (dev/stg/prd)** | Env-specific Helm values in repo | Same complexity either way |
| **Claude Code** | Full context, hierarchical CLAUDE.md | 40–60% token overhead for context |
| **GitHub workflows** | Single repo search, unified PRs | Cross-repo coordination |

**Final recommendation**: Adopt the monorepo with the structure above. Use Go workspaces for local development, Buf for Protobuf, path-based GitHub Actions for independent deployment, and ArgoCD ApplicationSets for GitOps. Small team size is actually one of the strongest use cases for a monorepo—you get the benefits (atomic commits, shared code, unified tooling) while avoiding the downsides that require tools like Bazel at scale.

If services are tightly related, you can start with a single root `go.mod`; if you want clearer boundaries, use per-service `go.mod` and `go.work`. Both work; the latter is more flexible if you later need to extract a service. Key insight: **loose coupling comes from interface design and event contracts, not from splitting repositories.**
