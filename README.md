<div align="center">

#  Go Web App — End-to-End DevOps Project

**A production-style Go web application deployed on AWS EKS with full CI/CD, GitOps, and observability**

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-Web_App-00ADD8?logo=go&logoColor=white)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Orchestrated-326CE5?logo=kubernetes&logoColor=white)](https://kubernetes.io/)
[![AWS EKS](https://img.shields.io/badge/AWS-EKS-FF9900?logo=amazonaws&logoColor=white)](https://aws.amazon.com/eks/)
[![ArgoCD](https://img.shields.io/badge/ArgoCD-GitOps-EF7B4D?logo=argo&logoColor=white)](https://argoproj.github.io/cd/)
[![GitHub Actions](https://img.shields.io/badge/CI%2FCD-GitHub_Actions-2088FF?logo=github-actions&logoColor=white)](https://github.com/features/actions)

*From a simple Go app to a fully automated cloud-native deployment — step by step.*

</div>

---

##  Overview

This project takes a Go portfolio web application and evolves it into a **complete production-grade DevOps pipeline** — covering every layer a modern cloud team uses: containerization, Kubernetes orchestration, Helm packaging, managed cloud infrastructure on AWS EKS, GitOps with ArgoCD, automated CI/CD via GitHub Actions, and full observability with Prometheus and Grafana.

---

##  Architecture

```
Developer (git push)
        │
        ▼
GitHub Repository
        │
        ▼
GitHub Actions CI/CD
(Build → Docker Image → DockerHub Push)
        │
        ▼
DockerHub Registry
        │
        ▼
ArgoCD (GitOps — detects change, auto-syncs)
        │
        ▼
AWS EKS Cluster
(Kubernetes Deployment + Service + Ingress)
        │
        ▼
Prometheus + Grafana Monitoring
```

---

## 🛠️ Tech Stack

| Tool / Service | Purpose |
|---|---|
| **Go** | Backend web application |
| **HTML / CSS** | Frontend portfolio pages |
| **Docker** | Containerization |
| **DockerHub** | Container image registry |
| **Kubernetes** | Container orchestration |
| **Helm** | Kubernetes package management |
| **AWS EKS** | Managed Kubernetes cluster |
| **eksctl** | EKS cluster provisioning |
| **kubectl** | Kubernetes CLI |
| **ArgoCD** | GitOps continuous deployment |
| **GitHub Actions** | CI/CD automation |
| **Prometheus** | Metrics & monitoring |
| **Grafana** | Visualization & dashboards |

---

##  Project Structure

```
go-web-app/
├── .github/
│   └── workflows/
│       └── ci-cd.yml           # GitHub Actions pipeline
├── go-web-app-chart/            # Helm chart
│   ├── Chart.yaml
│   ├── values.yaml
│   └── templates/
├── k8s/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
├── static/
│   ├── home.html
│   ├── about.html
│   ├── projects.html
│   ├── contact.html
│   └── style.css
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── main.go
├── main_test.go
├── LICENSE
└── README.md
```

---

##  Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.21+
- [Docker](https://docs.docker.com/get-docker/) & DockerHub account
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/docs/intro/install/) v3+
- [eksctl](https://eksctl.io/)
- [AWS CLI](https://aws.amazon.com/cli/) — configured with your credentials
- [ArgoCD CLI](https://argo-cd.readthedocs.io/en/stable/cli_installation/) *(optional)*

---

### 1. Run Locally

```bash
git clone https://github.com/krishnakala987-byte/go-web-app.git
cd go-web-app
go run main.go
```

App runs at `http://localhost:8080`

---

### 2. Build & Run with Docker

```bash
# Build image
docker build -t krishna2915/go-web-app:latest .

# Run container
docker run -p 8080:8080 krishna2915/go-web-app:latest
```

---

### 3. Push to DockerHub

```bash
docker login
docker push krishna2915/go-web-app:latest
```

---

### 4. Provision AWS EKS Cluster

```bash
eksctl create cluster \
  --name go-web-app-cluster \
  --region us-east-1 \
  --nodegroup-name workers \
  --node-type t3.small \
  --nodes 2
```

Verify nodes are ready:

```bash
kubectl get nodes
```

---

### 5. Deploy with Helm

```bash
# Install
helm install go-web-app ./go-web-app-chart

# Upgrade (after changes)
helm upgrade go-web-app ./go-web-app-chart
```

---

### 6. Install Monitoring (Prometheus + Grafana)

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm install monitoring prometheus-community/kube-prometheus-stack \
  -n monitoring --create-namespace
```

Check stack is running:

```bash
kubectl get pods -n monitoring
kubectl get svc -n monitoring
```

---

### 7. Install ArgoCD (GitOps)

```bash
helm repo add argo https://argoproj.github.io/argo-helm
helm repo update

kubectl create namespace argocd
helm install argocd argo/argo-cd -n argocd
```

Access the ArgoCD dashboard:

```bash
kubectl port-forward service/argocd-server -n argocd 8080:443
```

Open `http://localhost:8080` — get the admin password:

```bash
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d
```

---

##  CI/CD Pipeline

The GitHub Actions workflow at `.github/workflows/ci-cd.yml` triggers on every push to `main`:

```
git push → main
       │
       ▼
GitHub Actions
  ├── Checkout code
  ├── Setup Docker Buildx
  ├── Login to DockerHub (via GitHub Secrets)
  ├── Build Docker image
  └── Push to DockerHub
       │
       ▼
ArgoCD detects new image tag
       │
       ▼
EKS cluster auto-updated (no kubectl apply needed)
```

**Required GitHub Secrets:**

| Secret | Description |
|---|---|
| `DOCKER_USERNAME` | Your DockerHub username |
| `DOCKER_PASSWORD` | Your DockerHub password or token |

---

##  GitOps Validation

To verify the full GitOps loop is working:

1. Edit any text in `static/home.html`
2. Commit and push to `main`
3. Watch GitHub Actions complete successfully
4. Open ArgoCD dashboard → observe auto-sync
5. Get the LoadBalancer URL: `kubectl get svc`
6. Refresh the app URL — changes are live

---

##  Access Live App

```bash
kubectl get svc
```

Look for the `EXTERNAL-IP` field — this is your LoadBalancer URL:

```
ad125ef6...us-east-1.elb.amazonaws.com
```

---

##  Troubleshooting

| Problem | Cause | Fix |
|---|---|---|
| Pods stuck in `Pending` | Insufficient node capacity | Scale the nodegroup (see below) |
| `CrashLoopBackOff` | Bad config or startup error | Check `kubectl logs <pod>` and `kubectl describe pod <pod>` |
| Helm ownership errors | Resources created manually before Helm | Delete conflicting resources, reinstall cleanly |
| ArgoCD sync failed | Namespace metadata mismatch | Correct namespace name and re-sync |

**Scale EKS nodegroup (fix for Pending pods):**

```bash
eksctl scale nodegroup \
  --cluster go-web-app-cluster \
  --name workers \
  --nodes 3 \
  --nodes-min 2 \
  --nodes-max 3 \
  --region us-east-1
```

**Key debug commands:**

```bash
kubectl describe pod <pod-name>    # Inspect events & errors
kubectl logs <pod-name>            # View app logs
kubectl get all                    # Full resource overview
kubectl rollout restart deployment <name>   # Force restart
```

---

##  Roadmap

- [ ] Terraform for full infrastructure provisioning
- [ ] HTTPS with cert-manager + Let's Encrypt
- [ ] Horizontal Pod Autoscaler (HPA)
- [ ] AWS Load Balancer Controller
- [ ] Loki log aggregation stack
- [ ] SonarQube code quality analysis
- [ ] Multi-environment deployments (dev / staging / prod)

---

##  Contributing

1. Fork the repository
2. Create a feature branch — `git checkout -b feature/your-feature`
3. Commit your changes — `git commit -m 'feat: add your feature'`
4. Push — `git push origin feature/your-feature`
5. Open a Pull Request

---

##  License

This project is licensed under the [MIT License](LICENSE).

---

##  Author

**Krishna Kala**  
*Cloud & DevOps Engineer*

[![GitHub](https://img.shields.io/badge/GitHub-krishnakala987--byte-181717?logo=github)](https://github.com/krishnakala987-byte)

---

<div align="center">
<sub> If this project helped you, consider giving it a star!</sub>
</div>
