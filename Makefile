.PHONY: update
update: tidy wire env

.PHONY: fmt
fmt: bin/gofumpt
	bin/gofumpt -w .

.PHONY: imports
imports: PROJECT_NAME=$(shell grep "^module " go.mod | awk '{print $$2}')
imports: bin/goimports
	bin/goimports -local ${PROJECT_NAME} -w .

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: env
env:
	cp .env.dist .env

.PHONY: wire
wire: bin/wire
	bin/wire config/wire/wire.go
	$(MAKE) fmt
	$(MAKE) imports

generate: docs
	go generate -tags wireinject ./...
	$(MAKE) fmt
	$(MAKE) imports

.PHONY: lint
lint: bin/golangci-lint
	bin/golangci-lint run

.PHONY: test
test: bin/gotestsum
	bin/gotestsum -- -race -p=1 ./...

abigen: bin/abigen
	bin/abigen --abi=./pkg/contract/Contract.abi --pkg=contract --out=./pkg/contract/contract.go

.PHONY: docs
docs: bin/swag
	bin/swag init
