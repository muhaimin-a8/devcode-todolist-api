include .env
MYSQL_URI = "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_HOST):$(MYSQL_PORT))/$(MYSQL_DBNAME)"

# -------------------------------------------------------------------------------------------
# Install Required Tools
# -------------------------------------------------------------------------------------------
#
# install-deps: install all required tools to run this app
install-deps:
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# -------------------------------------------------------------------------------------------
# Database Migrations
# -------------------------------------------------------------------------------------------
#
# migrate-create: create new migration with `name` as parameter
migrate-create:
	migrate create -ext sql -dir migrations "$(name)"

# migrate-up: migrate up all databases migrations, required `POSTGRES_URI`
migrate-up:
	migrate -database $(MYSQL_URI) -path migrations up

# migrate-up: migrate down all databases migrations, required `POSTGRES_URI`
migrate-down:
	migrate -database $(MYSQL_URI) -path migrations down


# -------------------------------------------------------------------------------------------
# Run and Build
# -------------------------------------------------------------------------------------------
#
# run: run app immediately without build
run:
	cd cmd; go run main.go

# build: building app to ./bin/devcode_todolist_api
build:
	@ echo "building app to ./bin/devcode_todolist_api"
	cd cmd; go build -o ../bin/devcode_todolist_api

# -------------------------------------------------------------------------------------------
# Testing
# -------------------------------------------------------------------------------------------
#
# test-unit: run all unit tests
test-unit:
