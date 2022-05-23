GOPATH = $(shell go env GOPATH)
WIRE = "$(shell ls ${GOPATH}/bin | grep wire)"
GOMOCK = "$(shell ls ${GOPATH}/bin | grep gomock)"

.PHONY: up
up:
	docker-compose up

.PHONY: donw
down:
	docker-compose down -v

.PHONY: integration-test
integration-test:
	cd integration-test && make test

.PHONY: test
test:
	go test $(FLAGS) ./...

.PHONY: docker-build
docker-build:
	docker-compose build

.PHONY: docker-build-nocache
docker-build-nocache:
	docker-compose build --no-cache

.PHONY: generate
generate:
	go generate -x ./...

# wireによるコード生成
.PHONY: wire
wire:
ifneq (wire, $(shell echo $(WIRE)))
	go install github.com/google/wire/cmd/wire@latest
endif
	$(GOPATH)/bin/wire

.PHTONY: gql
gql:
	go run github.com/99designs/gqlgen