UID=$(shell id -u)
GID=$(shell id -g)

DOCKER_CMD = docker-compose
DOCKER_CMD_RUN = $(DOCKER_CMD) run -u $(UID):$(GID)
DOCKER_CMD_START = $(DOCKER_CMD) up -d server
RUN_IN_DEV = $(DOCKER_CMD_RUN) --rm server

build: docker
	$(RUN_IN_DEV) go build -mod=readonly -o main.go

.PHONY: docker-start
docker-start: docker-build
	$(DOCKER_CMD_START)

.PHONY: docker-build
docker-build:
	$(DOCKER_CMD) build

.PHONY: docker-stop
docker-stop:
	$(DOCKER_CMD) stop

docker-clean:
	$(DOCKER_CMD) down -v --remove-orphans || true

.PHONY: test test-unit
test: test-unit

test-unit: docker-build #gogen-test
	$(RUN_IN_DEV) go test ./...

test-coverage:
	$(RUN_IN_DEV) scripts/coverage-report.sh

exec:
	docker exec -it carpedia_server_1 bash

GOGEN = go generate
include mock.mk