BINARY_NAME=blog_api
ENVIRONMENT=development
HTTP_SERVER_ADDRESS=0.0.0.0:9093
GRPC_SERVER_ADDRESS=0.0.0.0:9090
REDIS_ADDRESS=localhost:6379
DB_HOST=localhost
DB_PORT=5432
DB_NAME=blog
DB_USER=postgres
DB_PASSWORD=secret
DB_URL=postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
URL_FRONTEND=http://localhost:5173
TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
MINIO_ENDPOINT=http://minio:9000
MINIO_ACCESS_KEY_ID=secret
MINIO_SECRET_ACCESS_KEY=IENOZ7MJzi4Tpwn7ntiZd3zAOqUrOvRjI4qpXHxE
MINIO_USE_SSL=false
MINIO_BUCKET_NAME=blog
MINIO_URL_RESULT=http://localhost:9000/blog/
EMAIL_SENDER_NAME=sender_name
EMAIL_SENDER_ADDRESS=yourself@gmail.com
EMAIL_SENDER_PASSWORD=yourselfgmailpassword

## postgres: Start PostgreSQL container
postgres:
	docker compose up postgres -d

## postgresdown: Stop PostgreSQL container
postgresdown:
	docker compose down postgres

## redis: Start Redis container
redis:
	docker compose up redis -d

## redisdown: Stop Redis container
redisdown:
	docker compose down redis

## migrate: Create a new migration file
migrate:
	migrate create -ext sql -dir db/migration -seq $(name)

## migrateup: Apply all migrations
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

## migrateup1: Apply the next migration
migrateup_no:
	migrate -path db/migration -database "$(DB_URL)" -verbose up $(no)

## migratedown: Rollback all migrations
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

## migratedown1: Rollback the previous migration
migratedown_no:
	migrate -path db/migration -database "$(DB_URL)" -verbose down $(no)

migrate_version:
	migrate -path db/migration -database "$(DB_URL)" -verbose version

migrate_goto:
	migrate -path db/migration -database "$(DB_URL)" -verbose goto $(no)

migrate_force:
	migrate -path db/migration -database "$(DB_URL)" -verbose force $(no)

## sqlc: Generate code from SQL queries
sqlc:
	sqlc generate

## proto: Generate code from
proto:
	rm -rf pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

## evans: Run evans for gRPC
evans:
	evans --host localhost --port 9090 -r repl

## build: Build binary
build:
	@echo "Building back end..."
	go build -o ${BINARY_NAME} .
	@echo "Binary built!"

## run: builds and runs the application
run: build
	@echo "Starting back end..."
	@env HTTP_SERVER_ADDRESS=${HTTP_SERVER_ADDRESS} \
	GRPC_SERVER_ADDRESS=${GRPC_SERVER_ADDRESS} \
	ENVIRONMENT=${ENVIRONMENT} \
	REDIS_ADDRESS=${REDIS_ADDRESS} \
	DB_SOURCE=${DB_URL} \
	GIN_MODE=debug \
	URL_FRONTEND=${URL_FRONTEND} \
	TOKEN_SYMMETRIC_KEY=${TOKEN_SYMMETRIC_KEY} \
	MIGRATION_URL=${MIGRATION_URL} \
	DB_DRIVER=${DB_DRIVER} \
	MINIO_ENDPOINT=${MINIO_ENDPOINT} \
	MINIO_ACCESS_KEY_ID=${MINIO_ACCESS_KEY_ID} \
	MINIO_SECRET_ACCESS_KEY=${MINIO_SECRET_ACCESS_KEY} \
	MINIO_USE_SSL=${MINIO_USE_SSL} \
	MINIO_BUCKET_NAME=${MINIO_BUCKET_NAME} \
	MINIO_URL_RESULT=${MINIO_URL_RESULT} \
	EMAIL_SENDER_NAME=${EMAIL_SENDER_NAME} \
	EMAIL_SENDER_ADDRESS=${EMAIL_SENDER_ADDRESS} \
	EMAIL_SENDER_PASSWORD=${EMAIL_SENDER_PASSWORD} \
	./${BINARY_NAME} &
	@echo "Back end started!"

## stop: stops the running application
stop:
	@echo "Stopping back end..."
	@if [ -f "${BINARY_NAME}" ]; then \
		pkill -SIGTERM -f "./${BINARY_NAME}"; \
		rm ${BINARY_NAME}; \
	fi
	@echo "Stopped back end!"

## restart: stops and starts the running application
restart: stop run

.PHONY: postgres postgresdown redis redisdown migrate migrateup migrateup1 migratedown migratedown1 sqlc proto evans build run stop restart
