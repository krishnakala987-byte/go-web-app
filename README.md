# go-web-app

A Go web application that I took from a single binary running on my laptop all the way to a containerized deployment on AWS EKS, with the CI/CD pipeline, GitOps, and monitoring wired up around it. I built this to practice the full path a real service travels: build, containerize, push to a registry, deploy to Kubernetes, and then actually watch it run in Grafana.

## What's in here

The app itself is a small Go server that renders a few static HTML pages (home, about, projects, contact). The interesting part is everything built around it:

- A `Dockerfile` to build the image and a `docker-compose.yml` to run it locally
- A Helm chart (`go-web-app-chart`) that templates the Kubernetes manifests
- Plain Kubernetes manifests under `k8s/` (deployment, service, ingress)
- A GitHub Actions workflow that builds the image and pushes it to Docker Hub
- ArgoCD, so the cluster syncs itself from Git
- Prometheus and Grafana for metrics

## Architecture

```
Developer (git push)
        |
        v
GitHub repository
        |
        v
GitHub Actions  (build -> Docker image -> push to Docker Hub)
        |
        v
Docker Hub registry
        |
        v
ArgoCD  (detects the change and syncs the cluster)
        |
        v
AWS EKS cluster  (Deployment + Service + Ingress)
        |
        v
Prometheus + Grafana
```

The point of setting it up this way is that, once it's running, a push to `main` is all it takes. Actions builds and pushes the image, ArgoCD notices and rolls it out. No manual `kubectl apply` on every change.

## Tools used

| Tool | What it does here |
|---|---|
| Go | the web app itself |
| HTML / CSS | the static pages it serves |
| Docker | packaging the app into an image |
| Docker Hub | where the image is stored |
| Kubernetes | running the app |
| Helm | templating the Kubernetes manifests |
| AWS EKS | the managed Kubernetes cluster |
| eksctl | creating the cluster |
| kubectl | talking to the cluster |
| ArgoCD | GitOps, keeping the cluster in sync with Git |
| GitHub Actions | the CI/CD pipeline |
| Prometheus | collecting metrics |
| Grafana | dashboards |

## Project structure

```
go-web-app/
|-- .github/
|   `-- workflows/
|       `-- ci-cd.yml           # GitHub Actions pipeline
|-- go-web-app-chart/           # Helm chart
|   |-- Chart.yaml
|   |-- values.yaml
|   `-- templates/
|-- k8s/
|   |-- deployment.yaml
|   |-- service.yaml
|   `-- ingress.yaml
|-- static/
|   |-- home.html
|   |-- about.html
|   |-- projects.html
|   |-- contact.html
|   `-- style.css
|-- Dockerfile
|-- docker-compose.yml
|-- go.mod
|-- main.go
|-- main_test.go
`-- README.md
```

## Running it

### Prerequisites

You'll need Go 1.21+, Docker and a Docker Hub account, kubectl, Helm v3+, eksctl, and the AWS CLI configured with your credentials. The ArgoCD CLI is optional.

### Locally

```bash
git clone https://github.com/krishnakala987-byte/go-web-app.git
cd go-web-app
go run main.go
```

It comes up at http://localhost:8080.

### With Docker

```bash
docker build -t krishna2915/go-web-app:latest .
docker run -p 8080:8080 krishna2915/go-web-app:latest
```

Push it to Docker Hub:

```bash
docker login
docker push krishna2915/go-web-app:latest
```

### On AWS EKS

Create the cluster:

```bash
eksctl create cluster \
  --name go-web-app-cluster \
  --region us-east-1 \
  --nodegroup-name workers \
  --node-type t3.small \
  --nodes 2
```

Check the nodes came up:

```bash
kubectl get nodes
```

Deploy with the Helm chart:

```bash
helm install go-web-app ./go-web-app-chart
# after changes:
helm upgrade go-web-app ./go-web-app-chart
```

### Monitoring

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install monitoring prometheus-community/kube-prometheus-stack \
  -n monitoring --create-namespace
```

Check it's running:

```bash
kubectl get pods -n monitoring
kubectl get svc -n monitoring
```

### ArgoCD

```bash
helm repo add argo https://argoproj.github.io/argo-helm
helm repo update
kubectl create namespace argocd
helm install argocd argo/argo-cd -n argocd
```

Reach the dashboard and get the admin password:

```bash
kubectl port-forward service/argocd-server -n argocd 8080:443
kubectl -n argocd get secret argocd-initial-admin-secret \
  -o jsonpath="{.data.password}" | base64 -d
```

## The CI/CD pipeline

The workflow at `.github/workflows/ci-cd.yml` runs on every push to `main`. It checks out the code, sets up Docker Buildx, logs in to Docker Hub using repository secrets, builds the image, and pushes it. ArgoCD then picks up the new image and updates the cluster, so there's no manual deploy step.

It needs two GitHub secrets:

| Secret | What it is |
|---|---|
| `DOCKER_USERNAME` | your Docker Hub username |
| `DOCKER_PASSWORD` | your Docker Hub password or access token |

## Checking the GitOps loop actually works

To confirm the whole thing end to end: edit some text in `static/home.html`, commit and push to `main`, watch the Actions run finish, open the ArgoCD dashboard and watch it auto-sync, then get the LoadBalancer URL with `kubectl get svc` and refresh the app. The change should be live.

## Troubleshooting

These are the things that actually tripped me up while building it:

| Problem | Cause | Fix |
|---|---|---|
| Pods stuck in `Pending` | not enough node capacity | scale the nodegroup (below) |
| `CrashLoopBackOff` | bad config or a startup error | check `kubectl logs <pod>` and `kubectl describe pod <pod>` |
| Helm ownership errors | resources created by hand before Helm | delete the conflicting resources and reinstall cleanly |
| ArgoCD sync failing | namespace mismatch in the manifests | fix the namespace and re-sync |

Scaling the nodegroup (the fix for pending pods):

```bash
eksctl scale nodegroup \
  --cluster go-web-app-cluster \
  --name workers \
  --nodes 3 --nodes-min 2 --nodes-max 3 \
  --region us-east-1
```

Commands I kept coming back to:

```bash
kubectl describe pod <pod-name>            # events and errors
kubectl logs <pod-name>                    # app logs
kubectl get all                            # everything at a glance
kubectl rollout restart deployment <name>  # force a restart
```

## Things I'd still like to add

Terraform for the infrastructure instead of standing the cluster up with eksctl by hand, HTTPS with cert-manager and Let's Encrypt, a Horizontal Pod Autoscaler, the AWS Load Balancer Controller, Loki for log aggregation, SonarQube for code quality, and separate dev / staging / prod environments.

## Author

Krishna Kala. The application is under the MIT License (see `LICENSE`).
