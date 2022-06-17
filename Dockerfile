FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o bubble .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

#设置工作目录
WORKDIR $GOPATH /work/gotest

COPY ./wait-for.sh /work/gotest/
COPY ./config /work/gotest/config
COPY ./static /work/gotest/static
COPY ./views /work/gotest/views

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bubble /work/gotest/

RUN set -eux; \
    apt-get update; \
    apt-get install -y \
        --no-install-recommends \
        netcat; \
        chmod 755 /work/gotest/wait-for.sh

# 需要运行的命令
# ENTRYPOINT ["/bubble", "conf/config.ini"]