FROM golang:1.18.5-alpine3.16 as builder

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

COPY . .
RUN go mod download


RUN cd app/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../baize

FROM ubuntu:jammy-20220801


WORKDIR /usr/local/baize
COPY --from=builder /build/baize .

VOLUME ["/usr/local/baize/config","/usr/local/baize/file","/usr/local/baize/baizeLog"]
EXPOSE 8080
ENTRYPOINT ["./baize"]
