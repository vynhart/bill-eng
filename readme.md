# System Requirement

- MySql

## For development

- Go version ^1.21
- [sqlc](https://github.com/sqlc-dev/sqlc) to generate golang from SQL statement.

# Setup

## Set `.env`

This service use `.env` file for configuration. Copy `.env.example` to `.env`
and adjust the value inside the file accordingly.

## Create Database

To be able to run database migration, you should create the database manually. So get into mysql console and run

```
CREATE DATABASE bill_engine_db;
CREATE DATABASE bill_engine_test_db;
```

adjust the db_name based on configuration on `.env` file accordingly and give
privilege to the user specified in `.env` file.

## Run Migration

Run database migration by using `migrate` command:

```
go run . migrate
```

# Testing

This service requires database to run the test.

First, make sure you've setup database and run migration on testing env:

```
ENV=testing go run . migrate
```

Then run test:

```
go test ./...
```