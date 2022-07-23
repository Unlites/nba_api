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

generate_mock:
	mockgen -source=internal/game/usecase.go \
	-destination=internal/game/mocks/usecase_mock.go
	mockgen -source=internal/game/repository.go \
	-destination=internal/game/mocks/repository_mock.go
	mockgen -source=internal/player/usecase.go \
	-destination=internal/player/mocks/usecase_mock.go
	mockgen -source=internal/player/repository.go \
	-destination=internal/player/mocks/repository_mock.go
	mockgen -source=internal/stat/usecase.go \
	-destination=internal/stat/mocks/usecase_mock.go
	mockgen -source=internal/stat/repository.go \
	-destination=internal/stat/mocks/repository_mock.go
	mockgen -source=internal/team/usecase.go \
	-destination=internal/team/mocks/usecase_mock.go
	mockgen -source=internal/team/repository.go \
	-destination=internal/team/mocks/repository_mock.go