PROTO_DIR := proto
DOCKER_COMPOSE := docker/docker-compose.yaml

.PHONY: proto up down test tidy

proto:
	buf generate

up:
	docker compose -f $(DOCKER_COMPOSE) up -d

down:
	docker compose -f $(DOCKER_COMPOSE) down

test:
	go test ./...

tidy:
	go mod tidy
