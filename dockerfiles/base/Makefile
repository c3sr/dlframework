
default: docker_build_cpu

build: docker_build output ## Build Docker image
push: docker_push output ## Push Docker image

release: docker_build docker_push ## Build and push Docker image

# Image and binary can be overidden with env vars.
DOCKER_IMAGE ?= c3sr/base

# Get System architecture
# for now... ask carl how best to get this
ARCH ?= $(shell go env GOARCH)

# Get the latest commit.
GIT_COMMIT = $(strip $(shell git rev-parse --short HEAD))

# Get the version number from the code
CODE_VERSION = v$(strip $(shell cat ../../VERSION))

# Find out if the working directory is clean
GIT_NOT_CLEAN_CHECK = $(shell git status --porcelain)

# Version of gcc to install
GCC_VERSION = 6

# Version of go to install
GO_VERSION ?= 1.15

# Version of cuda to use
CUDA_VERSION ?= 10.0

# If we're releasing to Docker Hub, and we're going to mark it with the latest tag, it should exactly match a version release
ifeq ($(MAKECMDGOALS),release)
# Use the version number as the release tag.
DOCKER_TAG = $(CODE_VERSION)

ifndef CODE_VERSION
$(error You need to create a VERSION file to build a release)
endif

# See what commit is tagged to match the version
VERSION_COMMIT = $(strip $(shell git rev-list $(CODE_VERSION) -n 1 | cut -c1-7))
ifneq ($(VERSION_COMMIT), $(GIT_COMMIT))
$(error echo You are trying to push a build based on commit $(GIT_COMMIT) but the tagged release version is $(VERSION_COMMIT))
endif

else
# Add the commit ref for development builds. Mark as dirty if the working directory isn't clean
DOCKER_TAG = $(CODE_VERSION)-$(GIT_COMMIT)$(DOCKER_TAG_SUFFIX)
endif

BUILD_ARGS = --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
			 			 --build-arg VERSION=$(CODE_VERSION) \
			       --build-arg VCS_URL=`git config --get remote.origin.url` \
		         --build-arg VCS_REF=$(GIT_COMMIT) \
		         --build-arg ARCH=$(ARCH) \
						 --build-arg GCC_VERSION=$(GCC_VERSION) \
						 --build-arg GO_VERSION=$(GO_VERSION)

docker_build_gpu: ## Build GPU Docker image
	@docker build $(BUILD_ARGS) \
		-f Dockerfile.$(ARCH)_gpu_cuda$(CUDA_VERSION) \
		-t $(DOCKER_IMAGE):$(ARCH)-gpu-go$(GO_VERSION)-cuda$(CUDA_VERSION)-$(DOCKER_TAG) .

docker_build_cpu:  ## Build CPU Docker image
	@docker build $(BUILD_ARGS) \
		-f Dockerfile.$(ARCH)_cpu \
		-t "$(DOCKER_IMAGE):$(ARCH)-cpu-go$(GO_VERSION)-$(DOCKER_TAG)" .

docker_build: docker_build_cpu docker_build_gpu

docker_push_cpu: ## Tag image as latest and Push to DockerHub
	docker tag $(DOCKER_IMAGE):$(ARCH)-cpu-go$(GO_VERSION)-$(DOCKER_TAG) $(DOCKER_IMAGE):$(ARCH)-cpu-go$(GO_VERSION)-latest
	docker push $(DOCKER_IMAGE):$(ARCH)-cpu-go$(GO_VERSION)-$(DOCKER_TAG)
	docker push $(DOCKER_IMAGE):$(ARCH)-cpu-go$(GO_VERSION)-latest

docker_push_gpu: ## Tag image as latest and Push to DockerHub
	docker tag $(DOCKER_IMAGE):$(ARCH)-gpu-go$(GO_VERSION)-cuda$(CUDA_VERSION)-$(DOCKER_TAG) $(DOCKER_IMAGE):$(ARCH)-gpu-go$(GO_VERSION)-cuda$(CUDA_VERSION)-latest
	docker push $(DOCKER_IMAGE):$(ARCH)-gpu-go$(GO_VERSION)-cuda$(CUDA_VERSION)-$(DOCKER_TAG)
	docker push $(DOCKER_IMAGE):$(ARCH)-gpu-go$(GO_VERSION)-cuda$(CUDA_VERSION)-latest

docker_push: docker_push_cpu docker_push_gpu

output:
	@echo Docker Image: $(DOCKER_IMAGE):$(DOCKER_TAG)

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: help

## See https://github.com/microscaling/microscaling/blob/master/Makefile
