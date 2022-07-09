build:
	docker-compose build

run:
	docker-compose up

stop:
	docker-compose stop

migrate:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@0.0.0.0:5435/postgres?sslmode=disable' up