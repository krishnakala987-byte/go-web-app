# go-web-app

A small Go web app taken all the way to production on AWS EKS. The app itself is simple on purpose. The point of this repo is the delivery pipeline around it.

## What it does

- Runs on a 2-node EKS cluster provisioned with eksctl.
- Uses a multi-stage Docker build that keeps the final image roughly 60% smaller by leaving build dependencies behind.
- Ships as a Helm chart, so the same app deploys to dev and prod with different values.
- Builds and pushes a versioned image to DockerHub on every merge to main, through GitHub Actions.
- Uses ArgoCD to watch the repo and sync manifests to the cluster, with health checks and Git-revert rollback when a deploy goes wrong.
- Adds Grafana dashboards, an HPA, RBAC, and liveness/readiness probes so it behaves like a real workload rather than a demo.

## Things I broke and fixed

Most of what I learned here came from things going wrong:

- Pending pods that turned out to be node group scaling limits.
- CrashLoopBackOff from Helm ownership conflicts on resources that already existed.
- ArgoCD sync errors caused by a namespace mismatch between the chart and the app.
- Stale images from relying on the latest tag instead of versioned ones.

## Stack

Go, Docker, Kubernetes, Helm, eksctl, AWS EKS, GitHub Actions, ArgoCD, Prometheus, Grafana.
