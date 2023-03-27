DOCKER_TAG ?= movier
docker-build:
	$(call describe_job,"Building docker image '$(DOCKER_TAG)'")
	docker build -f .infra/Dockerfile -t $(DOCKER_TAG) .