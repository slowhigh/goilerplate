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

