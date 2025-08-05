# 第一阶段：构建 Go 应用
FROM golang:1.23 as builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

# 静态编译（Linux 下推荐）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

# 第二阶段：创建极小镜像
FROM scratch
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
