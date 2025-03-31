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
   			service.proto
	PATH=${PATH}:${LOCALBIN} ${LOCALBIN}/protoc/bin/protoc --proto_path=api/server/grpc/goph-keeper/secret \
		--go_out=pkg/pb/goph-keeper/secret --go_opt=paths=source_relative \
		--go-grpc_out=pkg/pb/goph-keeper/secret --go-grpc_opt=paths=source_relative \
		service.proto