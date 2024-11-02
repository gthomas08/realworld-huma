## jet: run Jet generator 
.PHONY: jet
jet:
	jet -dsn="postgresql://admin:root@localhost:5432/postgres?sslmode=disable" -path=gen --ignore-tables=goose_db_version

## migrate: run database migrations
.PHONY: migrate
migrate:
	goose -dir ./migrations postgres "user=admin dbname=postgres password=root sslmode=disable" up
