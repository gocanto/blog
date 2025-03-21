ifneq (,$(wildcard .env))
  include .env
endif

PADDING="    "
GREEN = \033[0;32m
YELLOW=\033[1;33m
NC = \033[0m
BLUE=\033[0;34m

SOURCE ?= go_bindata
DATABASE ?= postgres
VERSION ?= $(shell git describe --tags 2>/dev/null | cut -c 2-)
REPO_OWNER ?= $(shell cd .. && basename "$$(pwd)")

# ------ App Configuration
ROOT_NETWORK ?= gocanto
ROOT_PATH ?= $(shell pwd)
ROOT_ENV_FILE ?= $(ROOT_PATH)/.env
ROOT_EXAMPLE_ENV_FILE? = $(ROOT_PATH)/.env.example
STORAGE_PATH ?= $(ROOT_PATH)/storage
BIN_PATH ?= $(ROOT_PATH)/bin
BIN_LOGS_PATH ?= $(ROOT_PATH)/bin/storage/logs

# ------ Database Configuration
# --- Docker
DB_DOCKER_SERVICE_NAME ?= postgres
DB_DOCKER_CONTAINER_NAME ?= gocanto-db
# --- Paths
DB_ROOT_PATH ?= $(ROOT_PATH)/database
DB_SSL_PATH ?= $(DB_ROOT_PATH)/ssl
DB_DATA_PATH ?= $(DB_ROOT_PATH)/data
# --- SSL
DB_SERVER_CRT ?= $(DB_SSL_PATH)/server.crt
DB_SERVER_CSR ?= $(DB_SSL_PATH)/server.csr
DB_SERVER_KEY ?= $(DB_SSL_PATH)/server.key
# --- Migrations
DB_MIGRATE_PATH ?= $(ROOT_PATH)/database/migrations
DB_MIGRATE_VOL_MAP ?= $(DB_MIGRATE_PATH):$(DB_MIGRATE_PATH)

flush:
	rm -rf $(DB_DATA_PATH) && \
	docker compose down --remove-orphans && \
	docker container prune -f && \
	docker image prune -f && \
	docker volume prune -f && \
	docker network prune -f && \
	docker ps

dependencies:
	$(call external_deps,'.')
	$(call external_deps,'./bin/...')
	$(call external_deps,'./cmd/...')
	$(call external_deps,'./packages/...')

build\:api:
	make logs:bin:fresh && \
	CGO_ENABLED=0 go build -a -ldflags='-X main.Version=$(VERSION)' -o "$(ROOT_PATH)/bin/api" -tags '$(DATABASE) $(SOURCE)' $(ROOT_PATH)/cmd/api

build\:api\:linux:
	make logs:bin:fresh && \
	cd $(ROOT_PATH)/cmd/api && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o "$(ROOT_PATH)/bin/api" -ldflags='-X main.Version=$(VERSION) -extldflags "-static"' -tags '$(DATABASE) $(SOURCE)' $(ROOT_PATH)/cmd/api

api\:run:
	cd $(BIN_PATH) && ./api

api\:air:
	air $(ROOT_PATH)/cmd/api/main.go

api\:release:
	git tag v$(V)
	@read -p "Press enter to confirm and push to origin ..." && git push origin v$(V)

env\:init:
	rm -f $(ROOT_ENV_FILE) && cp $(ROOT_EXAMPLE_ENV_FILE) $(ROOT_ENV_FILE)

db\:local:
	# --- Works with your local PG installation.
	cd  $(EN_DB_BIN_DIR) && \
	./psql -h $(ENV_DB_HOST) -U $(ENV_DB_USER_NAME) -d $(ENV_DB_DATABASE_NAME) -p $(ENV_DB_PORT)

db\:up:
	docker compose up $(DB_DOCKER_SERVICE_NAME) -d && \
	make db:logs

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

db\:secure:
	make flush && \
	rm -rf $(DB_SERVER_CRT) && rm -rf $(DB_SERVER_CSR) && rm -rf $(DB_SERVER_KEY) && \
	openssl genpkey -algorithm RSA -out $(DB_SERVER_KEY) && \
    openssl req -new -key $(DB_SERVER_KEY) -out $(DB_SERVER_CSR) && \
    openssl x509 -req -days 365 -in $(DB_SERVER_CSR) -signkey $(DB_SERVER_KEY) -out $(DB_SERVER_CRT) && \
    chmod 600 $(DB_SERVER_KEY) && chmod 600 $(DB_SERVER_CRT)

db\:secure\:show::
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
	#migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force VERSION
	docker run -v $(DB_MIGRATE_VOL_MAP) --network ${ROOT_NETWORK} migrate/migrate migrate -path $(DB_MIGRATE_PATH) -database $(ENV_DB_URL) force $(version)

logs\:clear:
	find $(STORAGE_PATH)/logs -maxdepth 1 -type f -not -name ".gitkeep" -delete

logs\:bin\:fresh:
	@rm -rf "$(BIN_LOGS_PATH)"
	@mkdir -m 777 $(BIN_LOGS_PATH)
	@touch $(BIN_LOGS_PATH)/.gitkeep

define external_deps
	@echo '-- $(1)';  go list -f '{{join .Deps "\n"}}' $(1) | grep -v github.com/$(REPO_OWNER)/blog | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'
endef

.PHONY: flush dependencies
.PHONY: build\:api
.PHONY: api\:air api\:build api\:release api\:run
.PHONY: env\:init
.PHONY: db\:local db\:up db\:ping db\:bash db\:fresh db\:logs db\:delete db\:secure db\:secure\:show
.PHONY: migrate\:up migrate\:down migrate\:create db\:migrate\:force
.PHONY: logs\:clear logs\:bin\:fresh

RAND = $(shell echo $$RANDOM)
