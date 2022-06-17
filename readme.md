# Simple Bank - Go using framework Gin

# Section 1: Working with database

# 4. Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc

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

# .\sqlc init ( .\sqlc version )

# Chạy này thì mới tạo được file trong sqlc:

    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc version
    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc init
    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate


    i think your volume mapping -v parameter isn't correct.
    it must map your current_folder/src to /src inside the container
    docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

    can you try with %cd%?
    i know a student who make it work with this command in the Makefile:
    sqlc:
        docker run --rm -v "$(CURDIR):/src" -w //src kjconroy/sqlc generate
    some other students make it work by using WSL:
    - Setup WSL
    https://docs.microsoft.com/en-us/windows/wsl/install-win10
    - And run the linux binary of sqlc like you would normally run an executable.
    wsl sqlc generate
    or
    root@LAPTOP-PFVT9C7K:/mnt/c/Git/sqlc_1.12.0_linux_amd64# ./sqlc generate -f "../study/simple-bank/sqlc.yaml"

    other students run this directly in the PowerShell:
    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc version
    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc init
    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

    can you try those solutions and tell me which one works for you?
    i will try to find a Windows laptop and see how it works later
