# Go Web App DevOps Project

A complete end-to-end DevOps implementation of a Go web application using Docker, Kubernetes, Helm, AWS EKS, ArgoCD, GitHub Actions, Prometheus and Grafana.

---

# Project Overview

This project started as a simple Go portfolio web application and was gradually upgraded step-by-step into a complete production-style DevOps deployment.

Instead of only deploying a container manually, the goal was to learn and implement:

* Containerization
* Kubernetes orchestration
* Helm packaging
* AWS EKS cluster management
* Monitoring
* GitOps
* CI/CD automation
* Real production-style deployment workflow

This project helped me understand how modern DevOps teams build, deploy, monitor and manage applications in real cloud environments.

---

# Technologies Used

| Tool / Service | Purpose                       |
| -------------- | ----------------------------- |
| Go             | Backend web application       |
| HTML/CSS       | Frontend portfolio pages      |
| Docker         | Containerization              |
| DockerHub      | Image registry                |
| Kubernetes     | Container orchestration       |
| Helm           | Kubernetes package management |
| AWS EKS        | Managed Kubernetes cluster    |
| kubectl        | Kubernetes CLI                |
| eksctl         | EKS cluster management        |
| ArgoCD         | GitOps continuous deployment  |
| GitHub Actions | CI/CD automation              |
| Prometheus     | Monitoring                    |
| Grafana        | Visualization and dashboards  |
| VS Code        | Development environment       |
| Ubuntu WSL     | Linux terminal environment    |
| Git & GitHub   | Version control               |

---

# Complete Project Journey

# Phase 1 — Building the Application

The project started with creating a simple portfolio web application using Go.

Files created:

```text
main.go
static/
├── home.html
├── about.html
├── projects.html
├── contact.html
└── style.css
```

Goal:

* Learn Go basics
* Understand serving static files
* Build a simple real project instead of only tutorials

Command used:

```bash
go run main.go
```

Purpose:

Runs the Go application locally.

---

# Phase 2 — Docker Containerization

After the application worked locally, the next step was containerization.

A Dockerfile was created.

Purpose:

* Package application with dependencies
* Run consistently on any machine
* Prepare for Kubernetes deployment

Important commands:

```bash
docker build -t krishna2915/go-web-app:latest .
```

Purpose:

Builds Docker image from Dockerfile.

---

```bash
docker images
```

Purpose:

Lists available Docker images.

---

```bash
docker run -p 8080:8080 krishna2915/go-web-app:latest
```

Purpose:

Runs container locally.

---

```bash
docker ps
```

Purpose:

Shows running containers.

---

# Phase 3 — DockerHub Integration

The image was pushed to DockerHub.

Purpose:

* Store images remotely
* Allow Kubernetes to pull images
* Enable CI/CD automation later

Commands:

```bash
docker login
```

Purpose:

Authenticate DockerHub.

---

```bash
docker push krishna2915/go-web-app:latest
```

Purpose:

Push Docker image to DockerHub.

---

# Phase 4 — Kubernetes Deployment

Kubernetes manifests were created manually.

Files created:

```text
k8s/
├── deployment.yaml
├── service.yaml
└── ingress.yaml
```

What I learned:

* Deployments manage pods
* Services expose applications
* Ingress manages routing
* Kubernetes desired state model

Commands used:

```bash
kubectl apply -f deployment.yaml
```

Purpose:

Creates Kubernetes deployment.

---

```bash
kubectl apply -f service.yaml
```

Purpose:

Creates Kubernetes service.

---

```bash
kubectl get pods
```

Purpose:

Checks pod status.

---

```bash
kubectl get svc
```

Purpose:

Checks services and external access.

---

```bash
kubectl describe pod <pod-name>
```

Purpose:

Troubleshoot pod issues.

---

```bash
kubectl logs <pod-name>
```

Purpose:

View application logs.

---

# Phase 5 — Helm Implementation

The project was upgraded from raw YAML files to Helm charts.

Helm structure:

```text
go-web-app-chart/
├── Chart.yaml
├── values.yaml
└── templates/
```

Why Helm was used:

* Better Kubernetes management
* Reusable templates
* Easier upgrades
* Production standard deployment method

Commands:

```bash
helm create go-web-app-chart
```

Purpose:

Creates Helm chart structure.

---

```bash
helm install go-web-app ./go-web-app-chart
```

Purpose:

Installs Helm chart into Kubernetes.

---

```bash
helm upgrade go-web-app ./go-web-app-chart
```

Purpose:

Updates existing Helm deployment.

---

# Phase 6 — AWS EKS Cluster Setup

The project was moved from local Kubernetes to AWS EKS.

Why EKS:

* Managed Kubernetes service
* Real cloud deployment
* Production-level infrastructure
* Industry standard

Cluster creation:

```bash
eksctl create cluster \
--name go-web-app-cluster \
--region us-east-1 \
--nodegroup-name workers \
--node-type t3.small \
--nodes 2
```

Purpose:

Creates EKS cluster with worker nodes.

---

Useful commands:

```bash
kubectl get nodes
```

Purpose:

Checks connected Kubernetes worker nodes.

---

```bash
eksctl get nodegroup --cluster go-web-app-cluster --region us-east-1
```

Purpose:

Checks EKS nodegroups.

---

# Scaling Issue Faced

Problem:

Pods stayed in Pending state.

Reason:

Worker nodes were insufficient.

Error observed:

```text
0/2 nodes are available: Too many pods
```

Troubleshooting:

* Used kubectl describe pod
* Checked scheduler events
* Identified resource limitation

Fix:

Scaled EKS nodegroup.

Command:

```bash
eksctl scale nodegroup \
--cluster go-web-app-cluster \
--name workers \
--nodes 3 \
--nodes-min 2 \
--nodes-max 3 \
--region us-east-1
```

Lesson learned:

Always verify worker node capacity when pods remain pending.

---

# Phase 7 — Monitoring Setup

Prometheus and Grafana were installed.

Purpose:

* Monitor cluster health
* Monitor workloads
* Visualize metrics
* Learn observability

Commands:

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
```

Purpose:

Adds Prometheus Helm repository.

---

```bash
helm install monitoring prometheus-community/kube-prometheus-stack -n monitoring --create-namespace
```

Purpose:

Installs Prometheus and Grafana stack.

---

Useful commands:

```bash
kubectl get pods -n monitoring
```

```bash
kubectl get svc -n monitoring
```

---

# Phase 8 — ArgoCD GitOps Setup

ArgoCD was implemented for GitOps deployment.

What is GitOps?

Git becomes the single source of truth.

Instead of manually applying Kubernetes changes:

```bash
kubectl apply -f deployment.yaml
```

ArgoCD automatically:

* Detects GitHub changes
* Syncs cluster state
* Deploys updates
* Maintains desired state

This removes manual deployment steps.

---

# ArgoCD Installation

Initial installation using manifest created multiple issues.

Problem faced:

* CRD conflicts
* Metadata ownership issues
* CrashLoopBackOff
* Pending pods

Troubleshooting steps:

* Checked logs using:

```bash
kubectl logs <pod-name> -n argocd
```

* Described pods:

```bash
kubectl describe pod <pod-name> -n argocd
```

* Checked scheduler events
* Removed conflicting resources
* Reinstalled ArgoCD properly using Helm

Final successful installation:

```bash
helm repo add argo https://argoproj.github.io/argo-helm
```

```bash
helm repo update
```

```bash
kubectl create namespace argocd
```

```bash
helm install argocd argo/argo-cd -n argocd
```

---

# Accessing ArgoCD UI

Command:

```bash
kubectl port-forward service/argocd-server -n argocd 8080:443
```

Purpose:

Access ArgoCD dashboard locally.

URL:

```text
http://localhost:8080
```

Get admin password:

```bash
kubectl -n argocd get secret argocd-initial-admin-secret \
-o jsonpath="{.data.password}" | base64 -d
```

---

# GitOps Validation

Test performed:

* Modified portfolio text
* Pushed code to GitHub
* ArgoCD automatically detected change
* Application synced automatically
* Deployment updated without kubectl apply

This confirmed complete GitOps workflow.

---

# Phase 9 — GitHub Actions CI/CD

GitHub Actions workflow created:

```text
.github/workflows/ci-cd.yml
```

Purpose:

Automate:

* Build
* Docker image creation
* DockerHub push

Workflow process:

```text
Git Push
↓
GitHub Actions
↓
Docker Build
↓
DockerHub Push
↓
ArgoCD Detects Change
↓
Kubernetes Updates Automatically
```

---

# CI/CD Pipeline Steps

GitHub Actions performed:

* Checkout code
* Setup Docker Buildx
* Login to DockerHub
* Build Docker image
* Push Docker image

DockerHub credentials stored securely using GitHub Secrets.

Secrets used:

```text
DOCKER_USERNAME
DOCKER_PASSWORD
```

---

# GitHub Actions Commands Used

Workflow trigger:

```yaml
on:
  push:
    branches:
      - main
```

Purpose:

Runs pipeline automatically on push to main branch.

---

# Project Architecture

```text
Developer
   ↓
GitHub Repository
   ↓
GitHub Actions CI/CD
   ↓
DockerHub Registry
   ↓
ArgoCD GitOps
   ↓
AWS EKS Cluster
   ↓
Kubernetes Deployment
   ↓
Prometheus + Grafana Monitoring
```

---

# Folder Structure

```text
GO-WEB-APP/
├── .github/
│   └── workflows/
│       └── ci-cd.yml
├── go-web-app-chart/
├── k8s/
├── static/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── main.go
├── main_test.go
├── README.md
└── LICENSE
```

---

# Important Troubleshooting Learned

# 1. Pending Pods

Cause:

Insufficient node resources.

Fix:

Scale nodegroup.

---

# 2. CrashLoopBackOff

Cause:

Incorrect configuration or startup failures.

Commands used:

```bash
kubectl logs <pod-name>
```

```bash
kubectl describe pod <pod-name>
```

---

# 3. Helm Ownership Errors

Cause:

Resources created manually before Helm installation.

Fix:

Delete conflicting resources and reinstall cleanly.

---

# 4. ArgoCD Sync Failed

Cause:

Namespace metadata issue.

Fix:

Corrected namespace naming and reapplied sync.

---

# Important Kubernetes Commands

```bash
kubectl get pods
```

View pods.

---

```bash
kubectl get svc
```

View services.

---

```bash
kubectl get ingress
```

View ingress resources.

---

```bash
kubectl get all
```

View all resources.

---

```bash
kubectl delete pod <pod-name>
```

Delete pod.

---

```bash
kubectl rollout restart deployment <deployment-name>
```

Restart deployment.

---

```bash
kubectl describe deployment <deployment-name>
```

Inspect deployment configuration.

---

# Important Docker Commands

```bash
docker build -t image-name .
```

Build image.

---

```bash
docker push image-name
```

Push image.

---

```bash
docker ps
```

Running containers.

---

# Important Helm Commands

```bash
helm list
```

View Helm releases.

---

```bash
helm uninstall <release-name>
```

Remove Helm release.

---

# Important ArgoCD Validation

How to verify GitOps is working:

1. Change application text.
2. Commit and push code.
3. Watch GitHub Actions succeed.
4. Open ArgoCD dashboard.
5. Observe automatic sync.
6. Refresh LoadBalancer URL.
7. Verify updated application.

---

# How to Find LoadBalancer URL

Command:

```bash
kubectl get svc
```

Look for:

```text
EXTERNAL-IP
```

Example:

```text
ad125ef6ad1514d368643f1e88847b3f-1253962898.us-east-1.elb.amazonaws.com
```

---

# Screenshot Guide

# 1. EKS Cluster Proof

Command:

```bash
kubectl get nodes
```

Take screenshot of worker nodes.

---

# 2. Kubernetes Pods

Command:

```bash
kubectl get pods -A
```

Take screenshot of running pods.

---

# 3. ArgoCD Sync

Open:

```text
http://localhost:8080
```

Take screenshot showing:

* Synced
* Healthy
* Deployment tree

---

# 4. GitHub Actions Success

Open:

```text
GitHub → Actions
```

Take screenshot showing green successful workflow.

---

# 5. Monitoring Dashboard

Open Grafana dashboard and capture:

* CPU metrics
* Node metrics
* Cluster metrics

---

# 6. Live Website

Open LoadBalancer URL and take screenshot.

---

# Key Lessons Learned

* Kubernetes troubleshooting is very important.
* Monitoring is essential in production.
* GitOps removes manual deployment steps.
* CI/CD automation saves time and reduces errors.
* Helm simplifies Kubernetes management.
* Cloud infrastructure requires proper scaling.
* Logs and describe commands are critical for debugging.

---

# Future Improvements

Possible future upgrades:

* Terraform infrastructure provisioning
* HTTPS with cert-manager
* Horizontal Pod Autoscaler
* AWS Load Balancer Controller
* Loki logging stack
* SonarQube code analysis
* Multi-environment deployments

---

# Final Outcome

This project evolved from a simple Go application into a complete production-style DevOps implementation.

It includes:

* Cloud deployment
* Kubernetes orchestration
* Monitoring
* GitOps
* CI/CD automation
* Automated deployments
* Production-style architecture

This project significantly improved my understanding of real-world DevOps workflows and modern cloud-native deployment practices.

---

# Author

Krishna Kala

Cloud & DevOps Engineer
