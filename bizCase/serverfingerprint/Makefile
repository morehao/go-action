# 变量定义
APP_NAME := server-fingerprint
VERSION := 1.0.0
DOCKER_IMAGE := $(APP_NAME):$(VERSION)
DOCKER_IMAGE_LATEST := $(APP_NAME):latest

# Go 相关变量
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod

# 构建输出
BINARY_NAME := server-fingerprint
BINARY_LINUX := $(BINARY_NAME)_linux

.PHONY: help build clean test run docker-build docker-run docker-stop k8s-deploy k8s-delete

# 默认目标
help: ## 显示帮助信息
	@echo "可用的 make 命令："
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# 本地开发
deps: ## 下载依赖
	$(GOMOD) download
	$(GOMOD) tidy

build: ## 构建应用
	$(GOBUILD) -o $(BINARY_NAME) -v ./

build-linux: ## 构建 Linux 版本
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_LINUX) -v ./

clean: ## 清理构建文件
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)

test: ## 运行测试
	$(GOTEST) -v ./...

run: ## 本地运行
	$(GOBUILD) -o $(BINARY_NAME) -v ./
	./$(BINARY_NAME)

dev: ## 开发模式运行（带热重载，需要安装 air）
	@which air > /dev/null || (echo "请先安装 air: go install github.com/cosmtrek/air@latest" && exit 1)
	air

# Docker 相关
docker-build: ## 构建 Docker 镜像
	docker build -t $(DOCKER_IMAGE_LATEST) -f ./Dockerfile .
	echo "✅ Docker 镜像 $(DOCKER_IMAGE_LATEST) 已构建完成"

docker-run: ## 运行 Docker 容器
	docker run -d \
		--name $(APP_NAME) \
		-p 8080:8080 \
		--rm \
		$(DOCKER_IMAGE_LATEST)

docker-stop: ## 停止 Docker 容器
	docker stop $(APP_NAME) || true

docker-logs: ## 查看 Docker 容器日志
	docker logs -f $(APP_NAME)

docker-compose-up: ## 使用 docker-compose 启动
	docker-compose up -d

docker-compose-down: ## 使用 docker-compose 停止
	docker-compose down

docker-compose-logs: ## 查看 docker-compose 日志
	docker-compose logs -f

# Kubernetes 相关
k8s-build-image: ## 为 K8s 构建镜像
	docker build -t $(DOCKER_IMAGE_LATEST) .

k8s-deploy: ## 部署到 Kubernetes
	kubectl apply -f k8s-manifests.yaml

k8s-delete: ## 从 Kubernetes 删除
	kubectl delete -f k8s-manifests.yaml

k8s-status: ## 查看 K8s 部署状态
	kubectl get all -n server-fingerprint

k8s-logs: ## 查看 K8s Pod 日志
	kubectl logs -f -l app=server-fingerprint -n server-fingerprint

k8s-port-forward: ## 端口转发到本地
	kubectl port-forward service/server-fingerprint-service 8080:80 -n server-fingerprint

# 测试相关
test-api: ## 测试 API 接口
	@echo "测试健康检查接口..."
	curl -s http://localhost:8080/health | jq .
	@echo "\n测试指纹接口..."
	curl -s http://localhost:8080/fingerprint | jq .

test-docker-api: ## 测试 Docker 容器 API
	@echo "等待容器启动..."
	sleep 5
	@echo "测试健康检查接口..."
	curl -s http://localhost:8080/health | jq .
	@echo "\n测试指纹接口..."
	curl -s http://localhost:8080/fingerprint | jq .

# 完整部署流程
deploy-local: clean build run ## 本地完整部署

deploy-docker: docker-build docker-stop docker-run ## Docker 完整部署
	@echo "等待容器启动..."
	sleep 5
	make test-docker-api

deploy-k8s: k8s-build-image k8s-deploy ## K8s 完整部署
	@echo "等待 Pod 启动..."
	kubectl wait --for=condition=ready pod -l app=server-fingerprint -n server-fingerprint --timeout=60s
	@echo "部署完成，运行 'make k8s-port-forward' 进行端口转发测试"

# 清理所有环境
clean-all: clean docker-stop k8s-delete ## 清理所有环境
	docker rmi $(DOCKER_IMAGE_LATEST) || true
	docker rmi $(DOCKER_IMAGE) || true

# 安装开发工具
install-tools: ## 安装开发工具
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 代码质量检查
lint: ## 代码质量检查
	@which golangci-lint > /dev/null || (echo "请先运行 make install-tools 安装 golangci-lint" && exit 1)
	golangci-lint run

# 格式化代码
fmt: ## 格式化代码
	go fmt ./...
	go vet ./...