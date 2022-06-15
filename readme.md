# Simple Bank - Go using framework Gin

# Section 1: Working with database

# 3. How to write & run database migration in Golang

# postgres:
    docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
# createdb:
    docker exec -it postgres12 createdb --username=root --owner=root simple_bank

# dropdb:
    docker exec -it postgres12 dropdb simple_bank

# .PHONY: createdb

# Run migration-up : 
    migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/postgres?sslmode=disable" -verbose up

# Run migration-down : 
    migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/postgres?sslmode=disable" -verbose down
