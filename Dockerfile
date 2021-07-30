FROM golang:1.16-alpine

RUN apk --update upgrade \
    && apk --no-cache --no-progress add git mercurial bash gcc musl-dev curl tar ca-certificates tzdata \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/bitmyth/hooks

COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on GOPROXY=https://goproxy.cn go mod download

COPY . .

EXPOSE 8000

CMD exec go run src/main.go
