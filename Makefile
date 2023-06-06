# Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
# This is licensed software from AccelByte Inc, for limitations
# and restrictions contact your company contract manager.

SHELL := /bin/bash

GOLANG_DOCKER_IMAGE := golang:1.19

# TODO: do we need to adjust this too? e.g. the BUILDER value
#IMAGE_NAME := $(shell basename "$$(pwd)")-app
#BUILDER := grpc-plugin-server-builder

proto:
	rm -rfv pkg/pb/*
	mkdir -p pkg/pb
	docker run -t --rm -u $$(id -u):$$(id -g) -v $$(pwd):/data/ -w /data/ rvolosatovs/protoc:4.0.0 \
			--proto_path=pkg/proto \
			--go_out=pkg/pb \
			--go_opt=paths=source_relative \
			--go-grpc_out=pkg/pb \
			--go-grpc_opt=paths=source_relative \
			pkg/proto/accelbyte-async-api/iam/oauth/v1/*.proto

lint: proto
	rm -f lint.err
	find -type f -iname go.mod -exec dirname {} \; | while read DIRECTORY; do \
		echo "# $$DIRECTORY"; \
		docker run -t --rm -u $$(id -u):$$(id -g) -v $$(pwd):/data/ -w /data/ -e GOCACHE=/data/.cache/go-build -e GOLANGCI_LINT_CACHE=/data/.cache/go-lint golangci/golangci-lint:v1.42.1\
				sh -c "cd $$DIRECTORY && golangci-lint -v --timeout 5m --max-same-issues 0 --max-issues-per-linter 0 --color never run || touch /data/lint.err"; \
	done
	[ ! -f lint.err ] || (rm lint.err && exit 1)

build: proto
	docker run -t --rm -u $$(id -u):$$(id -g) -v $$(pwd):/data/ -w /data/ -e GOCACHE=/data/.cache/go-build $(GOLANG_DOCKER_IMAGE) \
		sh -c "go build"

#image:
#	docker buildx build -t ${IMAGE_NAME} --load .
#
#imagex:
#	docker buildx inspect $(BUILDER) || docker buildx create --name $(BUILDER) --use
#	docker buildx build -t ${IMAGE_NAME} --platform linux/arm64/v8,linux/amd64 .
#	docker buildx build -t ${IMAGE_NAME} --load .
#	docker buildx rm --keep-state $(BUILDER)
#
#imagex_push:
#	@test -n "$(IMAGE_VERSION)" || (echo "IMAGE_VERSION is not set"; exit 1)
#	docker buildx inspect $(BUILDER) || docker buildx create --name $(BUILDER) --use
#	docker buildx build -t ${IMAGE_PREFIX}${IMAGE_NAME}:${IMAGE_VERSION} --platform linux/arm64/v8,linux/amd64 --push .
#	docker buildx rm --keep-state $(BUILDER)
