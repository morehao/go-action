# 外部传入参数（只用于 build 和 run）
DeployMode ?=
DeployVersion ?=

VERSION_PKG = github.com/morehao/go-action/bizcase/gobuild/version

# build 前参数校验
check-vars:
ifndef DeployMode
	$(error ❌ deployMode is required. Use: make build deployMode=dev deployVersion=1.0.0)
endif
ifndef DeployVersion
	$(error ❌ deployVersion is required. Use: make build deployMode=dev deployVersion=1.0.0)
endif

# 编译带变量的可执行文件
build: check-vars
	@echo "🔧 Building with:"
	@echo "  - Deploy Mode:     $(DeployMode)"
	@echo "  - Deploy Version:  $(DeployVersion)"
	go build -ldflags "\
		-X '$(VERSION_PKG).deployMode=$(DeployMode)' \
		-X '$(VERSION_PKG).deployVersion=$(DeployVersion)'" \
		-o myapp main.go

# 编译并运行
run: build
	@echo "🚀 Running..."
	./myapp
