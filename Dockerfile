FROM golang:alpine as builder

WORKDIR /go/src/github.com/wangrui19970405/wu-shi-admin/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="wangrui19970405@gmail.com"

WORKDIR /go/src/github.com/wangrui19970405/wu-shi-admin/server

COPY --from=0 /go/src/github.com/wangrui19970405/wu-shi-admin/server/server ./
COPY --from=0 /go/src/github.com/wangrui19970405/wu-shi-admin/server/resource ./resource/
COPY --from=0 /go/src/github.com/wangrui19970405/wu-shi-admin/server/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
