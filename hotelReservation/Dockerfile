FROM golang:1.13
ARG MODE=""
COPY . /go/src/github.com/harlow/go-micro-services
WORKDIR /go/src/github.com/harlow/go-micro-services
RUN go install -ldflags="-s -w" -mod=vendor --tags=$MODE ./cmd/...
