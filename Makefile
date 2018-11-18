# Variables
SERVICE=api-gateway
IMG_HUB?=registry.test.io/api-gateway
TAG?=latest
# Version information
VERSION=1.0.0
REVISION=${shell git rev-parse --short HEAD}
RELEASE=production
BUILD_HASH=${shell git rev-parse HEAD}
BUILD_TIME=${shell date "+%Y-%m-%d@%H:%M:%SZ%z"}
LD_FLAGS:=-X main.Version=$(VERSION) -X main.Revision=$(REVISION) -X main.Release=$(RELEASE) -X main.BuildHash=$(BUILD_HASH) -X main.BuildTime=$(BUILD_TIME)

run:prepare build
	@-docker service rm $(SERVICE)	
	@docker service create --name $(SERVICE) --network devel -p 8080:8080 -e GRPC_GO_LOG_SEVERITY_LEVEL=INFO $(IMG_HUB)/$(SERVICE):$(TAG)

prepare:
	@go get google.golang.org/grpc
	@go get github.com/gogo/protobuf/protoc-gen-gogofast
	@-docker swarm init
	@-docker network create --driver=overlay devel	

build:
	@go build -ldflags="$(LD_FLAGS)" -o bundles/$(SERVICE) cmd/main.go	
	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .

test:
	@go test -cover ./...

# PHONY
.PHONY : test test-integration generate fmt
