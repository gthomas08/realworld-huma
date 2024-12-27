# Variables
DB_HOST ?= localhost  # Defaults to localhost if not set

## run: run the application
.PHONY: run
run:
	go run ./cmd/api

## jet: run Jet generator 
.PHONY: jet
jet:
	jet -dsn="postgresql://admin:root@$$DB_HOST:5432/postgres?sslmode=disable" -path=gen --ignore-tables=goose_db_version

## migrate: run database migrations
.PHONY: migrate
migrate:
	goose -dir ./migrations postgres "host=$$DB_HOST user=admin password=root dbname=postgres sslmode=disable" up