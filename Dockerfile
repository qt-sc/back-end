FROM golang:1.15.3 as build

# 容器环境变量添加
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

MAINTAINER Qt-sc

WORKDIR $GOPATH/src/server
ADD . $GOPATH/src/server
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./server"]