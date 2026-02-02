# RMN Backend Infrastructure

Kubernetes manifests and GitOps overlays for backend services.

## Structure
- `base/` — shared namespace, configmap
- `overlays/dev` — development
- `overlays/stg` — staging
- `overlays/prd` — production

## Usage
Applied via ArgoCD (see argocd/appset.yaml).
