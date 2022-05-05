GOPATH = $(shell go env GOPATH)
WIRE = "$(shell ls ${GOPATH}/bin | grep wire)"
GOMOCK = "$(shell ls ${GOPATH}/bin | grep gomock)"

.PHONY: up
up:
	docker-compose up

.PHONY: donw
down:
	docker-compose down

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
	go run github.com/swaggo/swag/cmd/swag init -g cmd/main.go --output  ./docs

# デバッグ用途で、手動でimageをpushしたい時に利用.
# 下記コマンドを実行するのにperman認証が必要.
.PHONY: image-build-push
image-build-push:
	aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 176285604616.dkr.ecr.ap-northeast-1.amazonaws.com
	docker build -t api .
	docker tag api:latest 176285604616.dkr.ecr.ap-northeast-1.amazonaws.com/api:latest
	docker push 176285604616.dkr.ecr.ap-northeast-1.amazonaws.com/api:latest

# wireによるコード生成
.PHONY: wire
wire:
ifneq (wire, $(shell echo $(WIRE)))
	go install github.com/google/wire/cmd/wire@latest
endif
	$(GOPATH)/bin/wire
