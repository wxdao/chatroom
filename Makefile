export PATH := $(shell go env GOPATH)/bin:dep/protoc/bin:$(PATH)

# --------------------------

dep/protoc/bin/protoc:
	mkdir -p dep/protoc
	curl -L -o dep/protoc/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip
	cd dep/protoc; unzip -o protoc.zip

dep-install: dep/protoc/bin/protoc
	go install github.com/golang/protobuf/protoc-gen-go
	go get ./...

# --------------------------

internal/pkg/api/pb/push.pb.go: internal/pkg/api/pb/push.proto
	protoc -I internal/pkg/api/pb $? --go_out=plugins=grpc:internal/pkg/api/pb

proto: internal/pkg/api/pb/push.pb.go

# --------------------------

test: proto
	go test -v ./...

.PHONY: dep-install proto test
