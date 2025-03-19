# Docker compose
DOCKER_COMPOSE_FILE := docker-compose.yml
CONTAINER_NAME := ndot-postgres

# Database configuration
DB_NAME := ndot
DB_USER := postgres
DB_PASSWORD := postgres
DB_PORT := 5432

# Migration
MIGRATE_CMD := migrate
MIGRATION_DIR := db/migration

.PHONY: composeup composedown startdb stopdb createdb dropdb migrateup migratedown

composeup:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

composedown:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

startdb:
	docker start $(CONTAINER_NAME)

stopdb:
	docker stop $(CONTAINER_NAME)

createdb:
	docker exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME);"

dropdb:
	docker exec -it $(CONTAINER_NAME) psql -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

migrateup:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable up

migratedown:
	$(MIGRATE_CMD) -path $(MIGRATION_DIR) -database postgresql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)?sslmode=disable down

sqlc:
	sqlc generate