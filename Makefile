include .env

install-binaries:
	go install github.com/mitranim/gow@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

run-watch:
	gow -e=go,mod,html -c run main.go

# MIGRATION

goose-status:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose status

.PHONY: goose-create
goose-create:
# make goose-create name="teste_db" 
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose create $(name) sql

goose-up:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose up

goose-down:
	GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) goose down

# PROD

prod-env:
	$(eval include prod.env)
	@echo $(ENV)

run-prod:
	$(eval include prod.env)
	ENV=prod go run main.go