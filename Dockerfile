#使用golang镜像编译项目
FROM golang:1.20-alpine AS builder

#作者
MAINTAINER Leon

#指定docker中的工作目录
WORKDIR /go/src

#把本地项目拷贝到wordir中
ADD . /go/src

#执行编译
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo main.go

#拉取alpine镜像（空间小）
FROM alpine AS prod

# 拷贝执行文件
COPY --from=builder /go/src/main /main
# 拷贝配置文件
COPY --from=builder /go/src/config.yml ./config.yml

CMD ["./main"]
