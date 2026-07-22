include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d chatgo-postgres

env-down:
	@docker compose down chatgo-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down chatgo-postgres port-forwarder && \
		sudo rm -rf $(PROJECT_ROOT)/out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq, Пример: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm chatgo-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсутствует необходимый параметр action, Пример: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm chatgo-postgres-migrate \
	-path /migrations \
	-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@chatgo-postgres:5432/$(POSTGRES_DB)?sslmode=disable \
	"$(action)"

logs-cleanup:
		@read -p "Очистить все log файлы? Опасность утери логов. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf $(PROJECT_ROOT)/out/logs && \
		echo "Файлы логов очищены"; \
	else \
		echo "Очистка логов отменена"; \
	fi

chatgoapp-run:
	@export LOGGER_FOLDER=$(PROJECT_ROOT)/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run $(PROJECT_ROOT)/cmd/chatgoapp/main.go

chatgoapp-deploy:
	@docker compose up -d --build chatgoapp

chatgoapp-undeploy:
	@docker compose down chatgoapp

swagger-gen:
	@docker compose run --rm swagger \
		init \
		-g cmd/chatgoapp/main.go \
		-o docs \
		--parseInternal \
		--parseDependency

ps:
	@docker compose ps

perms:
	@sudo chmod -R 755 $(PROJECT_ROOT)/out