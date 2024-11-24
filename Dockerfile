# 使用官方的 Go 运行时作为父镜像
FROM golang:1.17-alpine AS builder

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到容器中的 /app 目录
COPY . .

# 构建可执行文件
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用一个更小的基础镜像来减小最终镜像的大小
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制可执行文件到最终镜像
COPY --from=builder /app/main .

# 暴露应用监听的端口
EXPOSE 8080

# 设置容器启动时运行的命令
CMD ["./main"]