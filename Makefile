export PATH := $(shell go env GOPATH)/bin:dep/protoc/bin:$(PATH)

PROTO_TARGETS = pkg/message/message.pb.go \
				pkg/push/push.pb.go

# --------------------------

dep/protoc/bin/protoc:
	mkdir -p dep/protoc
	curl -L -o dep/protoc/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip
	cd dep/protoc; unzip -o protoc.zip

dep-install: dep/protoc/bin/protoc
	go install github.com/golang/protobuf/protoc-gen-go
	go get ./...

# --------------------------

$(PROTO_TARGETS): %.pb.go: %.proto
	protoc -I $(dir $@) $< --go_out=plugins=grpc:$(dir $@)

proto: $(PROTO_TARGETS)

# --------------------------

test: proto
	go test -v ./...

# --------------------------

oneserver: proto
	CGO_ENABLED=0 go build -o output/oneserver ./cmd/oneserver

client: proto
	CGO_ENABLED=0 go build -o output/client ./cmd/client


.PHONY: dep-install proto test oneserver client
