#!make
include .env

build:
	docker-compose build

run:
	docker-compose up -d

stop:
	docker-compose stop

migrate:
	migrate -path ./migrations -database 'postgres://postgres:${PG_PASSWORD}@0.0.0.0:5435/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./migrations -database 'postgres://postgres:${PG_PASSWORD}@0.0.0.0:5435/postgres?sslmode=disable' down