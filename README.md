<h1 align="center">
  <b>Goilerplate</b>
</h1>

<p align="center">
  For the perfect template. 🔥🔥🔥
</p>

![image](https://github.com/user-attachments/assets/bcecce5a-e09c-4e76-96c9-6562acc4fe65)


## Contents
- [Features](#features)
  - [Structure](#structure)
  - [Infra](#infra)
- [Kubernetes](#kubernetes)
  - [Network](#network)
  - [Load Balancer](#load-balancer)
  - [Ingress](#ingress)
  - [Auto Scaling](#auto-scaling)
  - [API](#api)
  - [PostgreSQL](#postgresql)
  - [pgAdmin4](#pgadmin4)
  - [Redis](#redis)
- [Quick start](#quick-start)
  - [Kubernetes in Docker Desktop (localhost)](#kubernetes-in-docker-desktop-localhost)
  - [Kubernetes in Public Cloud (goilerplate.com)](#kubernetes-in-public-cloud-goilerplatecom)
- [Makefile Macros](#makefile-macros)
  - [Docker Control](#docker-control)
  - [Kubectl Control](#kubectl-control)

## Features

#### Structure
- [x] **Routing** - Gin Web Framework ---------------------------- [📚](https://gin-gonic.com/docs) [:octocat:](https://github.com/gin-gonic/gin)
- [x] **CLI** - Cobra ------------------------------------------------- [📚](https://cobra.dev) [:octocat:](https://github.com/spf13/cobra)
- [x] **DI pattern** - Fx --------------------------------------------- [📚](https://uber-go.github.io/fx/get-started) [:octocat:](https://github.com/uber-go/fx)
- [x] **Environment** - Viper --------------------------------------- [:octocat:](https://github.com/spf13/viper)
- [x] **Logging** - Zap ---------------------------------------------- [:octocat:](https://github.com/uber-go/zap)
- [x] **PostgreSQL ORM** - GORM --------------------------------- [📚](https://gorm.io/docs) [:octocat:](https://github.com/go-gorm/gorm)
- [x] **Redis ORM** - Go-Redis ------------------------------------- [📚](https://redis.uptrace.dev/guide) [:octocat:](https://github.com/go-redis/redis)
- [x] **DB Viewer** - pgAdmin4 (Web) ----------------------------- [📚](https://www.pgadmin.org/docs/pgadmin4/latest/index.html) [🐳](https://hub.docker.com/r/dpage/pgadmin4) [:octocat:](https://github.com/pgadmin-org/pgadmin4)
- [x] **Authentication** - JWT (Access + Refresh) ----------------- [📚](https://golang-jwt.github.io/jwt/) [:octocat:](https://github.com/golang-jwt/jwt)
- [x] **Makefile** - make -------------------------------------------- [📚](https://www.gnu.org/savannah-checkouts/gnu/make/manual/make.html)
- [x] **CI/CD** - GitHub-Actions ------------------------------------ [📚](https://docs.github.com/en/actions)

#### Infra
- [x] **CNI(Container Network Interface)**- flannel --------------- [:octocat:](https://github.com/flannel-io/flannel)
- [x] **Load Balancer** - MetalLB ----------------------------------- [📚](https://metallb.universe.tf) [:octocat:](https://github.com/metallb/metallb)
- [x] **Ingress** - NGINX -------------------------------------------- [📚](https://kubernetes.github.io/ingress-nginx/deploy)
- [x] **AutoScaling** - k8s-HPA ------------------------------------- [📚](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)

## Kubernetes

#### Network
- [flannel](k8s/base/kube-flannel.yaml)

#### Load Balancer
- [MetalLB](k8s/base/metallb-native.yaml)
- [IPAddressPool](k8s/ipaddress-pools.yaml)

#### Ingress
- [NGINX Ingress Controller](k8s/base/ingress-nginx.yaml)
- [Ingress (Internal)](k8s/ingress-internal.yaml)
- [Ingress (External)](k8s/ingress-external.yaml)
- [Secret (TLS)](k8s/secret-tls.yaml)

#### Auto Scaling
- [Metrics Server](k8s/base/metrics-server.yaml)
- [HPA (HorizontalPodAutoscaler)](k8s/api-hpa.yaml)

#### API
- [Deployment](k8s/api-deployment.yaml)
- [Service](k8s/api-service.yaml)

#### PostgreSQL
- [Deployment](k8s/postgres-deployment.yaml)
- [Service](k8s/postgres-service.yaml)
- [Volume](k8s/postgres-volume.yaml)

#### pgAdmin4
- [Deployment](k8s/pgAdmin4-deployment.yaml)
- [Service](k8s/pgAdmin4-service.yaml)

#### Redis
- [Deployment](k8s/redis-deployment.yaml)
- [Service](k8s/redis-service.yaml)
- [Volume](k8s/redis-volume.yaml)


## Quick start
Make sure you have `make` installed.
```bash
$ sudo apt install make
```

### Kubernetes in **Docker Desktop** (localhost)
```bash
$ make deploy-to-docker-desktop
```

##### API Server URL (api.localhost)
[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.postman.com/goilerplate/workspace/goilerplate/documentation/9185268-4a71a5ff-e919-4fd4-b88c-2b7b972c7aef?entity=&branch=&version=)

##### pgAdmin4 URL
[https://pgadimin4.localhost](https://pgadimin4.localhost)

### Kubernetes in **Public Cloud** (goilerplate.com)
```bash
# Check the IP of the Kubernetes control plane.
$ kubectl cluster-info
Kubernetes control plane is running at https://10.0.0.8:6443
CoreDNS is running at https://10.0.0.8:6443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

# Edit the addresses of ipaddress-pools.yaml.
$ vi k8s/ipaddress-pools.yaml
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: first-pool
  namespace: metallb-system
spec:
  addresses:
  - 10.0.0.8-10.0.0.8             <---- edit IP
---
apiVersion: metallb.io/v1beta1
kind: L2Advertisement
metadata:
  name: l2-advert
  namespace: metallb-system

$ make deploy-to-cloud
```

##### API Server (api.goilerplate.com)
[![Run in Postman](https://run.pstmn.io/button.svg)](https://www.postman.com/goilerplate/workspace/goilerplate/documentation/9185268-4a71a5ff-e919-4fd4-b88c-2b7b972c7aef?entity=&branch=&version=)

##### pgAdmin4
[https://pgadimin4.localhost](https://pgadimin4.localhost)

## Makefile Macros
#### Docker Control
```bash
# Push the API Docker Image to the DockerHub.
$ make push_api

# Push the PostgreSQL Docker Image to the DockerHub.
$ make push_postgres

# Push the Redi Docker Image to the DockerHub.
$ make push_redis

# Push the pgAdmin4 Docker Image to the DockerHub.
$ make push_pgadmin4

# Push the All Docker Image to the DockerHub.
$ make push_all
```

#### Kubectl Control
```bash
# Deploy to the Docker Desktop. (WSL + Docker Desktop)
$ make deploy-to-docker-desktop

# Undeploy to the Docker Desktop. (WSL + Docker Desktop)
$ make delete-to-docker-desktop

# Deploy to the Public Cloud.
$ make deploy-to-cloud

# Undeploy to the Public Cloud.
$ make delete-to-cloud
```
