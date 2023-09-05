#!make
include .env
export $(shell sed 's/=.*//' .env)

run:
	go run cmd/*.go

compile:
	go build -o app cmd/main.go

compose:
	docker compose -f .ci/docker/docker-compose.yaml up -d 

build:
	buf generate

lint:
	golangci-lint -c .golangci-lint.yaml -v run ./...

runner:
	gitlab-runner exec docker --docker-privileged  `cat .envrc |cut -f 2 -d" "|sed 's/^/--env /g'|xargs)` $(service)
