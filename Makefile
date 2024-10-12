## jet: run Jet generator 
.PHONY: jet
jet:
	jet -dsn="postgresql://admin:root@localhost:5432/postgres?sslmode=disable" -path=gen --ignore-tables=goose_db_version
