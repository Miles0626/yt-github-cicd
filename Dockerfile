# 第一阶段：编译阶段 builder，升级go1.26
FROM golang:1.26-alpine AS builder

# 设置容器内工作目录
WORKDIR /app

# 先拷贝依赖文件，缓存依赖层（优化构建速度）
COPY go.mod ./
RUN go mod tidy

# 拷贝全部源码
COPY . .

# 静态编译，CGO关闭，生成无系统依赖的二进制
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o hello-app main.go

# 第二阶段：运行阶段，仅用轻量alpine
FROM alpine:latest
# 安装基础证书，防止https拉取镜像仓库报错（可选）
RUN apk --no-cache add ca-certificates

WORKDIR /app
# 从builder阶段复制编译好的程序
COPY --from=builder /app/hello-app ./

# 容器启动命令
CMD ["./hello-app"]