.PHONY: proto run

proto: services/**/proto/*.pb.go

%.pb.go: %.proto
	protoc --go_out=plugins=grpc:. $<

data/bindata.go: data/*.json
	go-bindata -o data/bindata.go -pkg data data/*.json

GO_SRCS=$(shell find . -name "*.go" | grep -v "vendor")

build: $(GO_SRCS)
	go mod vendor
	docker-compose build

run:
	docker-compose up --remove-orphans
