
ifeq ($(POSTGRES_SETUP),)
	POSTGRES_SETUP = postgresql://user:5432@localhost:5430/OzonIntro
endif

LOCAL_BIN:=$(CURDIR)/bin
MIGRATION_FOLDER = $(CURDIR)/migrations

.all-deps: .bin-deps .add-deps

.add-deps:
	$(info Installing dependencies...)
	go get -u github.com/jackc/pgx/v5/pgxpool
	go get -u github.com/99designs/gqlgen
	go get -u github.com/99designs/gqlgen/graphql/handler

.bin-deps:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	GOBIN=$(LOCAL_BIN) go install github.com/99designs/gqlgen@latest

build-compose:
	docker-compose build

compose-up-postgres:
	docker-compose up -d postgres

compose-down:
	docker-compose down

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create postgres sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP)" down
generate-schema:
	go run github.com/99designs/gqlgen generate
#Используй sudo если нет доступа к pgdata
build-ozonapi: 
	docker build -t ozonapi . 
up-ozonapi: 
	docker run -p 8080:8080  ozonapi