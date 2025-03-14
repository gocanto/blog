include .env

PADDING="    "
GREEN = \033[0;32m
YELLOW=\033[1;33m
NC = \033[0m
BLUE=\033[0;34m

# --- App Configuration
ROOT_NETWORK=gocanto
ROOT_PATH=$(shell pwd)
ROOT_ENV_FILE=$(ROOT_PATH)/.env
ROOT_EXAMPLE_ENV_FILE=$(ROOT_PATH)/.env.example

# --- Database Configuration
# --- Docker
DB_DOCKER_SERVICE_NAME=postgres
DB_DOCKER_CONTAINER_NAME=gocanto-db
# --- Paths
DB_ROOT_PATH=$(ROOT_PATH)/database
DB_SSL_PATH=$(DB_ROOT_PATH)/ssl
DB_DATA_PATH=$(DB_ROOT_PATH)/data
# --- SSL
DB_SERVER_CRT=$(DB_SSL_PATH)/server.crt
DB_SERVER_CSR=$(DB_SSL_PATH)/server.csr
DB_SERVER_KEY=$(DB_SSL_PATH)/server.key
# --- Migrations
DB_MIGRATE_PATH=$(ROOT_PATH)/database/migrations
DB_MIGRATE_VOL_MAP=$(DB_MIGRATE_PATH):$(DB_MIGRATE_PATH)

.PHONY: flush env\:generate db\:sql db\:up db\:ping db\:bash db\:fresh db\:logs db\:delete db\:dev\:crt\:fresh
.PHONY: db\:dev\:crt\:list migrate\:up migrate\:down migrate\:create migrate\:up\:force

flush:
	rm -rf $(DB_DATA_PATH) && \
	docker compose down --remove-orphans && \
	docker container prune -f && \
	docker image prune -f && \
	docker volume prune -f && \
	docker network prune -f && \
	docker ps

env\:generate:
	rm -f $(ROOT_ENV_FILE) && cp $(ROOT_EXAMPLE_ENV_FILE) $(ROOT_ENV_FILE)

db\:sql:
	# --- Works with your local PG installation.
	cd  $(EN_DB_BIN_DIR) && \
	./psql -h $(ENV_DB_HOST) -U $(ENV_DB_USER_NAME) -d $(ENV_DB_DATABASE_NAME) -p $(ENV_DB_PORT)

db\:up:
	docker compose up $(DB_DOCKER_SERVICE_NAME) -d

db\:ping:
	docker port $(DB_DOCKER_CONTAINER_NAME)

db\:bash:
	docker exec -it $(DB_DOCKER_CONTAINER_NAME) bash

db\:fresh:
	make db:delete && make db:up

db\:logs:
	docker logs -f $(DB_DOCKER_CONTAINER_NAME)

db\:delete:
	docker compose down $(DB_DOCKER_SERVICE_NAME) --remove-orphans && \
	rm -rf $(DB_DATA_PATH) && \
	docker ps

db\:dev\:crt\:fresh:
	make prune && \
	rm -rf $(DB_SERVER_CRT) && rm -rf $(DB_SERVER_CSR) && rm -rf $(DB_SERVER_KEY) && make prune && \
	openssl genpkey -algorithm RSA -out $(DB_SSL_PATH)/server.key && \
    openssl req -new -key $(DB_SERVER_KEY) -out $(DB_SERVER_CSR) && \
    openssl x509 -req -days 365 -in $(DB_SERVER_CSR) -signkey $(DB_SERVER_KEY) -out $(DB_SERVER_CRT) && \
    chmod 600 $(DB_SERVER_KEY) && chmod 600 $(DB_SERVER_CRT)

db\:dev\:crt\:list:
	docker exec -it $(DB_DOCKER_CONTAINER_NAME) ls -l /etc/ssl/private/server.key && \
	docker exec -it $(DB_DOCKER_CONTAINER_NAME) ls -l /etc/ssl/certs/server.crt

migrate\:up:
	@echo "\n${BLUE}${PADDING}--- Running DB Migrations ---\n${NC}"
	@docker run -v $(DB_MIGRATE_VOL_MAP) --network ${ROOT_NETWORK} migrate/migrate -verbose -path=$(DB_MIGRATE_PATH) -database $(ENV_DB_URL) up
	@echo "\n${GREEN}${PADDING}--- Done Running DB Migrations ---\n${NC}"

migrate\:down:
	@echo "\n${BLUE}${PADDING}--- Running DB Migrations ---\n${NC}"
	@docker run -v $(DB_MIGRATE_VOL_MAP) --network ${ROOT_NETWORK} migrate/migrate -verbose -path=$(DB_MIGRATE_PATH) -database $(ENV_DB_URL) down 1
	@echo "\n${GREEN}${PADDING}--- Done Running DB Migrations ---\n${NC}"

migrate\:create:
	docker run -v $(DB_MIGRATE_VOL_MAP) --network ${ROOT_NETWORK} migrate/migrate create -ext sql -dir $(DB_MIGRATE_PATH) -seq $(name)

migrate\:up\:force:
	migrate -path $(DB_MIGRATE_PATH) -database $(ENV_DB_URL) force $(version)
