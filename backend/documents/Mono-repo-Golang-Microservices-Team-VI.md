# Mono repo Golang – Đội ngũ vi mô dịch vụ (Microservices)

Với vi mô dịch vụ Golang dựa trên Kafka phát triển cùng Kubernetes và Claude Code, **monorepo là lựa chọn rõ ràng**. Giảm chi phí điều phối, chia sẻ schema Kafka đơn giản hơn, hiệu quả phát triển dựa trên AI, và kinh nghiệm của các công ty như Uber, Monzo vận hành hàng nghìn vi mô dịch vụ Go trong một monorepo đều ủng hộ điều này. Ngược lại, thất bại vi mô dịch vụ đa repo của Segment là bài học cảnh báo.

Triển khai độc lập và ghép nối lỏng lẻo hoàn toàn có thể đạt được trong monorepo nếu có cấu trúc phù hợp. Ý chính: **cấu trúc repo không quyết định độ ghép nối giữa các dịch vụ; thiết kế mã nguồn mới quyết định.**

---

## Vì sao monorepo phù hợp với đội Go nhỏ

Ý kiến chuyên gia cho đội 1–2 lập trình viên rất rõ. Ricardo Soares (LinkedIn): "Khi công ty rất nhỏ, cách tiếp cận monorepo là năng động nhất. Đa repo thường tốn kém hơn nhiều cho đội nhỏ." Brandon Schurman (Earthly) đồng ý: "Monorepo đặc biệt hấp dẫn khi làm việc với đội nhỏ."

**Go workspace (go.work) thay đổi cuộc chơi.** Workspace từ Go 1.18 giúp quản lý nhiều module trong monorepo bớt ma sát. Mỗi dịch vụ giữ `go.mod` riêng, nhưng với `go.work` bạn có thể phát triển tất cả dịch vụ cùng lúc. Thay đổi thư viện dùng chung áp dụng ngay, không cần chỉ thị `replace` khi phát triển, và gopls hỗ trợ điều hướng và refactor xuyên module.

Với đội nhỏ, vấn đề thực sự của đa repo là **chi phí điều phối**. Trong đa repo, một thay đổi trải nhiều dịch vụ nghĩa là: cập nhật repo thư viện → PR → merge → tag release → nâng phiên bản dependency trong service-1 → PR → merge → lặp lại cho từng dịch vụ. Trong monorepo? Một commit nguyên tử. Segment đã trả giá—với hơn 120 repo, "ba kỹ sư toàn thời gian dành phần lớn thời gian chỉ để giữ hệ thống sống."

| Yếu tố | Monorepo (1–2 người) | Đa repo (1–2 người) |
|--------|----------------------|----------------------|
| Refactor xuyên dịch vụ | Một commit | Nhiều PR cần phối hợp |
| Cập nhật thư viện dùng chung | Áp dụng ngay | Tag → release → nâng version ở từng consumer |
| Bảo trì CI/CD | Một pipeline | Quản lý N pipeline |
| Ngữ cảnh code review | Thấy đủ | Phân tán theo repo |
| Gánh nhận thức | Một mô hình tư duy | Phải nhớ repo nào chứa gì |

---

## Chia sẻ schema sự kiện Kafka: Monorepo đơn giản hóa mọi thứ

Trong kiến trúc hướng sự kiện dựa trên Kafka, quản lý schema là trung tâm—và đơn giản hơn nhiều trong monorepo. Nếu các dịch vụ dùng chung hợp đồng sự kiện, chúng nên dùng chung định nghĩa schema.

**Cách làm khuyến nghị: Protobuf + Buf CLI trong monorepo**

```
monorepo/
├── proto/
│   ├── buf.yaml       # Cấu hình Buf
│   ├── buf.gen.yaml   # Cấu hình sinh code
│   ├── common/
│   │   └── types.proto    # Kiểu dùng chung (timestamp, ID)
│   └── events/
│       ├── order/v1/events.proto
│       └── user/v1/events.proto
├── generated/
│   └── go/            # Go code được sinh
├── services/
│   └── [mọi dịch vụ import generated/go]
```

Chạy `buf generate` tạo ra struct Go. Mọi dịch vụ import `github.com/yourorg/monorepo/generated/go/events`. Schema và mã consumer cập nhật nguyên tử trong một commit—không cần phối hợp phiên bản.

Để **kiểm tra lúc chạy**, dùng Confluent Schema Registry ở chế độ tương thích **BACKWARD**. Thiết lập chính cho production:

- `auto.register.schemas=false` (tắt đăng ký tự động)
- `normalize.schemas=true` (bật chuẩn hóa schema)
- Dùng TopicNameStrategy cho tên topic: `{topic}-value`

Cách đa repo—một repo schema riêng phát hành dưới dạng Go module—thêm chi phí đáng kể: sau mỗi thay đổi schema bạn phát hành thư viện rồi nâng version ở mọi consumer. Đó chính là kiểu mẫu góp phần vào sự sụp đổ vi mô dịch vụ của Segment.

---

## Đạt triển khai độc lập trong monorepo

Triển khai độc lập theo từng dịch vụ hoàn toàn khả thi trong monorepo nhờ **kích hoạt theo đường dẫn** trong GitHub Actions. Đây không phải thỏa hiệp—đó là cùng mẫu Uber và Monzo dùng.

**Workflow GitHub Actions theo dịch vụ:**

```yaml
name: user-service
on:
  push:
    paths:
      - 'services/user-service/**'
      - 'pkg/**'              # Thư viện dùng chung
      - 'proto/events/user/**' # Schema riêng dịch vụ
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

Với đội nhỏ dưới 10 dịch vụ, **script tùy chỉnh dùng `git diff`** để phát hiện dịch vụ bị ảnh hưởng là đủ—không cần Bazel hay Nx. Khi dịch vụ tăng, sinh file workflow từ template.

---

## Tổ chức manifest Kubernetes với GitOps

**Cấu trúc khuyến nghị: Helm chart theo dịch vụ + ArgoCD ApplicationSets**

```
monorepo/
├── services/
│   ├── user-service/
│   │   ├── src/
│   │   └── chart/           # Helm chart theo dịch vụ
│   │       ├── Chart.yaml
│   │       ├── values.yaml  # Mặc định
│   │       ├── values-dev.yaml  # Ghi đè Dev
│   │       ├── values-stg.yaml  # Ghi đè Staging
│   │       ├── values-prd.yaml  # Ghi đè Production
│   │       └── templates/
├── infrastructure/
│   ├── base/                # Tài nguyên K8s dùng chung
│   └── overlays/
│       ├── dev/
│       ├── stg/
│       └── prd/
└── argocd/
    └── appset.yaml          # ApplicationSet cho mọi dịch vụ
```

**ArgoCD ApplicationSet** tự phát hiện dịch vụ:

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

Với mẫu này, thêm dịch vụ mới chỉ cần tạo thư mục có chứa chart—ArgoCD tự nhận.

---

## Kiểm chứng thực tế từ công ty tập trung Go

Các công ty vận hành hàng nghìn vi mô dịch vụ Go thành công đều dùng monorepo:

**Uber**: Hơn 3.000 vi mô dịch vụ, ~50 triệu dòng Go, hơn 1.000 commit/ngày trong monorepo Go. Họ chuyển từ đa repo sang monorepo năm 2018 do vấn đề quản lý dependency.

**Monzo**: Hơn 2.800 vi mô dịch vụ, toàn bộ bằng Go, toàn bộ trong một monorepo. Kỹ sư nền tảng Will Sewell: "Có monorepo cực kỳ hữu ích để quản lý một số thách thức thường gắn với vi mô dịch vụ." Ông nhấn mạnh phiên bản dependency thống nhất, phản hồi lúc biên dịch xuyên dịch vụ, và khả năng refactor quy mô lớn.

**Bài học Segment**: Họ bắt đầu monolith, chuyển sang vi mô dịch vụ với hơn 120 repo riêng, rồi **quay lại monolith năm 2017**. Kỹ sư mô tả "địa ngục thư viện dùng chung"—phiên bản khác nhau mỗi dịch vụ, cơn ác mộng test, chi phí vận hành tăng theo số dịch vụ. "Thay vì đi nhanh hơn, đội nhỏ bị nuốt bởi độ phức tạp bùng nổ."

Mẫu rõ ràng: công ty thành công với vi mô dịch vụ ở quy mô lớn dùng monorepo. Những nơi vất vả với vi mô dịch vụ đa repo hoặc quay lại hoặc đầu tư nặng vào hạ tầng—không khả thi cho đội 1–2 người.

---

## Claude Code hoạt động tốt hơn nhiều trong monorepo

Trong phát triển hỗ trợ AI **ngữ cảnh là yếu tố quyết định**—và monorepo mang lại lợi thế ngữ cảnh mà thiết lập đa repo khó đạt được nếu không có kỹ thuật đáng kể.

**Lợi thế monorepo cho Claude Code:**
- Thấy rõ cách dịch vụ tương tác với thư viện dùng chung
- Một cấu trúc CLAUDE.md phân cấp cung cấp đủ ngữ cảnh
- Refactor xuyên dịch vụ được hiểu với đủ ngữ cảnh
- Một lập trình viên đã rút CLAUDE.md từ **47k xuống 9k từ** nhờ cấu trúc phân cấp monorepo

**Vấn đề đa repo**: Đội báo **40–60% ngân sách token** dùng cho trùng lặp ngữ cảnh xuyên repo. Khi Claude chỉ thấy một repo, nó "không biết API dịch vụ kia trông thế nào."

**Cấu trúc CLAUDE.md khuyến nghị:**

```
monorepo/
├── CLAUDE.md              # Gốc: tổng quan kiến trúc, lệnh chung
├── services/
│   ├── user-service/
│   │   └── CLAUDE.md      # Theo dịch vụ: hợp đồng API, dependency
│   └── order-service/
│       └── CLAUDE.md
├── proto/
│   └── CLAUDE.md          # Quy ước schema, lệnh sinh code
└── pkg/
    └── CLAUDE.md          # Mẫu thư viện dùng chung
```

Nội dung mẫu CLAUDE.md gốc:

```markdown
# Platform Monorepo

## Lệnh nhanh
- Build toàn bộ: `make build`
- Test dịch vụ: `cd services/<tên> && go test ./...`
- Sinh Proto: `buf generate`
- Deploy staging: `make deploy-stg SERVICE=<tên>`

## Kiến trúc
- Dịch vụ giao tiếp qua sự kiện Kafka (xem proto/events/)
- Thư viện dùng chung nằm trong pkg/ — import như package nội bộ
- Mỗi dịch vụ có Helm chart riêng trong chart/

## Quy ước
- Đặt tên sự kiện: PascalCase động từ (OrderCreated, UserUpdated)
- Tên thư mục dịch vụ trùng với namespace Kubernetes
```

---

## Cấu trúc monorepo khuyến nghị cho dự án RMN

Dựa trên nghiên cứu trên, cấu trúc tối ưu cho vi mô dịch vụ Golang + Flutter + Kafka + Kubernetes:

```
platform/
├── CLAUDE.md              # Ngữ cảnh AI: kiến trúc, lệnh
├── go.work                # Go workspace (gitignore)
├── buf.yaml               # Cấu hình Protobuf
├── buf.gen.yaml
│
├── proto/                 # Schema sự kiện
│   └── events/
│       ├── order/v1/
│       └── user/v1/
├── generated/             # Code sinh (được commit)
│   └── go/
│
├── services/              # Vi mô dịch vụ Go
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
├── pkg/                   # Thư viện Go dùng chung
│   ├── kafka/             # Wrapper client Kafka
│   ├── middleware/        # Middleware HTTP/gRPC chung
│   └── config/            # Load cấu hình
│
├── web/                   # Flutter web
│   └── flutter_app/
│
├── infrastructure/        # Tài nguyên K8s dùng chung
│   ├── kafka/
│   └── monitoring/
│
├── argocd/                # Cấu hình GitOps
│   └── appset.yaml
│
├── .github/
│   └── workflows/
│       ├── user-service.yaml   # CI theo dịch vụ
│       ├── order-service.yaml
│       └── flutter-web.yaml
│
└── Makefile               # Lệnh chung
```

---

## Tóm tắt: Ma trận quyết định theo yêu cầu

| Yêu cầu | Giải pháp monorepo | Đa repo cần gì |
|---------|--------------------|----------------|
| **1–2 lập trình viên** | Phù hợp tự nhiên — điều phối tối thiểu | Chi phí quản lý repo đáng kể |
| **Triển khai độc lập** | Kích hoạt GitHub Actions theo đường dẫn | Vốn đã độc lập (nhưng thêm pipeline) |
| **Ghép nối lỏng** | Do thiết kế mã quyết định, không phải cấu trúc repo | Cùng kỷ luật thiết kế |
| **Chia sẻ schema Kafka** | Một thư mục proto/, cập nhật nguyên tử | Repo schema riêng + quản lý phiên bản |
| **K8s (dev/stg/prd)** | Helm values theo môi trường trong repo | Độ phức tạp tương đương |
| **Claude Code** | Đủ ngữ cảnh, CLAUDE.md phân cấp | 40–60% token cho ngữ cảnh |
| **GitHub workflow** | Tìm trong một repo, PR thống nhất | Cần điều phối xuyên repo |

**Khuyến nghị cuối**: Áp dụng monorepo với cấu trúc trên. Dùng Go workspace cho phát triển local, Buf cho Protobuf, GitHub Actions theo đường dẫn cho triển khai độc lập, ArgoCD ApplicationSets cho GitOps. Quy mô đội nhỏ thực ra là một trong các trường hợp monorepo mạnh nhất—bạn được lợi (commit nguyên tử, mã dùng chung, công cụ thống nhất) mà tránh nhược điểm cần công cụ như Bazel ở quy mô lớn.

Nếu các dịch vụ gắn chặt, có thể bắt đầu với một `go.mod` gốc; nếu muốn ranh giới rõ hơn thì dùng `go.mod` và `go.work` theo dịch vụ. Cả hai đều ổn; cách sau linh hoạt hơn nếu sau này cần tách dịch vụ. Ý chính: **ghép nối lỏng đạt được nhờ thiết kế giao diện và hợp đồng sự kiện, không phải nhờ tách repo.**
