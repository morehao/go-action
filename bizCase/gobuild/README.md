# gobuild

`golang` 构建工具case。
- `make run DeployMode=dev DeployVersion=1.0.0`: 构建并运行当前项目
  ``` bash
    ⚡  make run DeployMode=dev DeployVersion=1.0.0
    🔧 Building with:
    - Deploy Mode:     dev
    - Deploy Version:  1.0.0
    go build -ldflags "\
            -X 'github.com/morehao/go-action/bizCase/gobuild/version.DeployMode=dev' \
            -X 'github.com/morehao/go-action/bizCase/gobuild/version.DeployVersion=1.0.0'" \
            -o myapp main.go
    🚀 Running...
    ./myapp
    Version: 1.0.0
    Mode: dev
  ```