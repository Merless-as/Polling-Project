build:
	docker-compose build Polling-Project

run:
	docker-compose up Polling-Project

migrate:
	migrate -path ./schema -database  'postgres://postgres:postgres@0.0.0.0:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go