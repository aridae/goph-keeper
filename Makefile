LOCALBIN = $(PWD)/bin

.PHONY: smartimports
smartimports: export SMARTIMPORTS := ${LOCALBIN}/smartimports
smartimports:
	test -f ${SMARTIMPORTS} || GOBIN=${LOCALBIN} go install github.com/pav5000/smartimports/cmd/smartimports@latest
	PATH=${PATH}:${LOCALBIN} ${SMARTIMPORTS} -path . -exclude ./static ./../_mock -local github.com/aridae/go-metrics-store

.PHONY: generate-mocks
generate-mocks: export MOCKGEN := ${LOCALBIN}/mockgen
generate-mocks:
	test -f ${MOCKGEN} || GOBIN=${LOCALBIN} go install go.uber.org/mock/mockgen@latest
	PATH=${PATH}:${LOCALBIN} go generate -run mockgen $(shell find . -d -name '_mock')

.PHONY: lint
lint: export GOLANGCILINT := ${LOCALBIN}/golangci-lint
lint:
	test -f ${GOLANGCILINT} || GOBIN=${LOCALBIN} go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	PATH=${PATH}:${LOCALBIN} ${GOLANGCILINT} run

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: test-coverage
test-coverage: export COVERAGE_OUT_FILE := ./coverage.out
test-coverage:
	go test ./... -coverpkg=./... -coverprofile=${COVERAGE_OUT_FILE} -vet=all
	go tool cover -func=${COVERAGE_OUT_FILE}

# usage: make bench out=filename.txt
.PHONY: bench
bench:
	go test ./... -bench=./... -benchmem > ${out}

# usage: make benchstat old=old.txt new=new.txt
.PHONY: benchstat
benchstat: export BENCHSTATBIN := ${LOCALBIN}/benchstat
benchstat:
	test -f ${BENCHSTATBIN} || GOBIN=${LOCALBIN} go install golang.org/x/perf/cmd/benchstat@latest
	PATH=${PATH}:${LOCALBIN} ${BENCHSTATBIN} ${old} ${new}

.PHONY: install-protoc
install-protoc: export PROTOC_VERSION := 30.0
install-protoc: export PROTOC_ZIP := protoc-$(PROTOC_VERSION)-osx-x86_64.zip # версия только для osx
install-protoc:
	$(info Installing protoc...)
	curl -o ${LOCALBIN}/${PROTOC_ZIP} -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
	unzip -o ${LOCALBIN}/${PROTOC_ZIP} -d ${LOCALBIN}/protoc
	rm -f ${LOCALBIN}/${PROTOC_ZIP}

.PHONY: install-protoc-gen-go
install-protoc-gen-go: export PROTOC_GEN_GO_BIN := ${LOCALBIN}/protoc-gen-go
install-protoc-gen-go: export PROTOC_GEN_GO_GRPC_BIN := ${LOCALBIN}/protoc-gen-go-grpc
install-protoc-gen-go: install-protoc
	$(info Installing binary dependencies...)
	test -f ${PROTOC_GEN_GO_BIN} || GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	test -f ${PROTOC_GEN_GO_GRPC_BIN} || GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: generate-pb
generate-pb: install-protoc-gen-go
	$(info Generating proto and grpc implementation...)
	PATH=${PATH}:${LOCALBIN} ${LOCALBIN}/protoc/bin/protoc --proto_path=api/server/grpc/goph-keeper/user \
		--go_out=pkg/pb/goph-keeper/user --go_opt=paths=source_relative \
		--go-grpc_out=pkg/pb/goph-keeper/user --go-grpc_opt=paths=source_relative \
		users_service.proto
	PATH=${PATH}:${LOCALBIN} ${LOCALBIN}/protoc/bin/protoc --proto_path=api/server/grpc/goph-keeper/secret \
		--go_out=pkg/pb/goph-keeper/secret --go_opt=paths=source_relative \
		--go-grpc_out=pkg/pb/goph-keeper/secret --go-grpc_opt=paths=source_relative \
		secrets_service.proto

GOOS=darwin
GOARCH=amd64

VERSION = v0.0.1
COMMIT = $(shell git rev-parse HEAD)
DATE = $(shell date "+%Y-%m-%d")
LDFLAGS = -X main.buildVersion=$(VERSION) -X main.buildDate=$(DATE) -X main.buildCommit=$(COMMIT)

PHONY: build-client
build-client: export CLIENTBIN := ${LOCALBIN}/client
build-client:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ${CLIENTBIN} -ldflags "$(LDFLAGS)" cmd/client/main.go
