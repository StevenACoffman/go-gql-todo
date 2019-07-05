FROM golang:1.12.5-alpine

RUN apk add --no-cache git

ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOPATH=/go \
    GOBIN=$GOPATH/bin \
    CODE_PATH=/app \
    GO111MODULE=on \
    GOPROXY=https://proxy.golang.org \
    LOG_LEVEL=debug

RUN go get github.com/cespare/reflex

RUN mkdir -p $CODE_PATH

WORKDIR $CODE_PATH

EXPOSE 8080
