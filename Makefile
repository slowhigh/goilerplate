.PHONY:

# ==============================================================================
# Docker Control

HUB_NAME = someday94
REPO_PREFIX = goilerplate
DOCKERFILE_DIR_PATH = ./docker

push_api:
	@read -p "Enter Version Name:" version; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-api:$$version -f $(DOCKERFILE_DIR_PATH)/api/Dockerfile .; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-api:$$version; \

push_postgres:
	@read -p "Enter Version Name:" version; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-postgres:$$version -f $(DOCKERFILE_DIR_PATH)/postgres/Dockerfile .; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-postgres:$$version; \

push_redis:
	@read -p "Enter Version Name:" version; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-redis:$$version -f $(DOCKERFILE_DIR_PATH)/redis/Dockerfile .; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-redis:$$version; \

push_pgadmin4:
	@read -p "Enter Version Name:" version; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-pgadmin4:$$version -f $(DOCKERFILE_DIR_PATH)/pgadmin4/Dockerfile .; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-pgadmin4:$$version; \

push_all:
	@read -p "Enter Version Name:" version; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-api:$$version -f $(DOCKERFILE_DIR_PATH)/api/Dockerfile .; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-postgres:$$version -f $(DOCKERFILE_DIR_PATH)/postgres/Dockerfile .; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-redis:$$version -f $(DOCKERFILE_DIR_PATH)/redis/Dockerfile .; \
	docker build -t $(HUB_NAME)/$(REPO_PREFIX)-pgadmin4:$$version -f $(DOCKERFILE_DIR_PATH)/pgadmin4/Dockerfile .; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-api:$$version; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-postgres:$$version; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-redis:$$version; \
	docker push $(HUB_NAME)/$(REPO_PREFIX)-pgadmin4:$$version; \



# ==============================================================================
# Kubectl Control

deploy-to-docker-desktop:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
	kubectl apply -f k8s/base/metrics-server.yaml
	sleep 30
	kubectl apply -f k8s/postgres-volume.yaml
	kubectl apply -f k8s/postgres-deployment.yaml
	kubectl apply -f k8s/postgres-service.yaml
	kubectl apply -f k8s/redis-volume.yaml
	kubectl apply -f k8s/redis-deployment.yaml
	kubectl apply -f k8s/redis-service.yaml
	kubectl apply -f k8s/pgAdmin4-deployment.yaml
	kubectl apply -f k8s/pgAdmin4-service.yaml
	kubectl apply -f k8s/api-deployment.yaml
	kubectl apply -f k8s/api-service.yaml
	kubectl apply -f k8s/api-hpa.yaml
	kubectl apply -f k8s/ingress-internal.yaml

delete-to-docker-desktop:
	kubectl delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
	kubectl delete -f k8s/base/metrics-server.yaml
	kubectl delete -f k8s/postgres-volume.yaml
	kubectl delete -f k8s/postgres-deployment.yaml
	kubectl delete -f k8s/postgres-service.yaml
	kubectl delete -f k8s/redis-volume.yaml
	kubectl delete -f k8s/redis-deployment.yaml
	kubectl delete -f k8s/redis-service.yaml
	kubectl delete -f k8s/pgAdmin4-deployment.yaml
	kubectl delete -f k8s/pgAdmin4-service.yaml
	kubectl delete -f k8s/api-deployment.yaml
	kubectl delete -f k8s/api-service.yaml
	kubectl delete -f k8s/api-hpa.yaml
	kubectl delete -f k8s/ingress-internal.yaml

deploy-to-cloud:
	kubectl apply -f k8s/base/kube-flannel.yaml
	sleep 30
	kubectl apply -f k8s/base/metrics-server.yaml
	kubectl apply -f k8s/base/metallb-native.yaml
	kubectl apply -f k8s/ipaddress-pools.yaml
	kubectl apply -f k8s/base/ingress-nginx.yaml
	sleep 30
	kubectl apply -f k8s/postgres-volume.yaml
	kubectl apply -f k8s/postgres-deployment.yaml
	kubectl apply -f k8s/postgres-service.yaml
	kubectl apply -f k8s/redis-volume.yaml
	kubectl apply -f k8s/redis-deployment.yaml
	kubectl apply -f k8s/redis-service.yaml
	kubectl apply -f k8s/pgAdmin4-deployment.yaml
	kubectl apply -f k8s/pgAdmin4-service.yaml
	kubectl apply -f k8s/api-deployment.yaml
	kubectl apply -f k8s/api-service.yaml
	kubectl apply -f k8s/api-hpa.yaml
	kubectl apply -f k8s/secret-tls.yaml
	kubectl apply -f k8s/ingress-external.yaml

delete-to-cloud:
	kubectl delete -f k8s/base/kube-flannel.yaml
	kubectl delete -f k8s/base/metrics-server.yaml
	kubectl delete -f k8s/base/metallb-native.yaml
	kubectl delete -f k8s/ipaddress-pools.yaml
	kubectl delete -f k8s/base/ingress-nginx.yaml
	kubectl delete -f k8s/postgres-volume.yaml
	kubectl delete -f k8s/postgres-deployment.yaml
	kubectl delete -f k8s/postgres-service.yaml
	kubectl delete -f k8s/redis-volume.yaml
	kubectl delete -f k8s/redis-deployment.yaml
	kubectl delete -f k8s/redis-service.yaml
	kubectl delete -f k8s/pgAdmin4-deployment.yaml
	kubectl delete -f k8s/pgAdmin4-service.yaml
	kubectl delete -f k8s/api-deployment.yaml
	kubectl delete -f k8s/api-service.yaml
	kubectl delete -f k8s/api-hpa.yaml
	kubectl delete -f k8s/secret-tls.yaml
	kubectl delete -f k8s/ingress-external.yaml













