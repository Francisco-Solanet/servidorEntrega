.PHONY: run-app run templ sqlc up test clean

TEMPL_VERSION := v0.2.793
SQLC_VERSION := v1.31.1

# Genera el código, prepara la base de datos y ejecuta la aplicación.
run-app:
	@set -e; \
	trap 'docker compose down -v --remove-orphans' EXIT; \
	echo "Generando componentes Templ..."; \
	go run github.com/a-h/templ/cmd/templ@$(TEMPL_VERSION) generate; \
	echo "Generando consultas SQLc..."; \
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@$(SQLC_VERSION) generate; \
	echo "Limpiando contenedores y volumenes anteriores..."; \
	docker compose down -v --remove-orphans; \
	echo "Iniciando PostgreSQL..."; \
	docker compose up -d; \
	echo "Esperando a PostgreSQL..."; \
	until docker compose exec -T postgres pg_isready -U postgres -d vehiculos > /dev/null 2>&1; do sleep 1; done; \
	echo "Iniciando aplicacion..."; \
	go run .

run: run-app

templ:
	go run github.com/a-h/templ/cmd/templ@$(TEMPL_VERSION) generate

sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@$(SQLC_VERSION) generate

up:
	docker compose up -d

test:
	$(MAKE) templ
	$(MAKE) sqlc
	$(MAKE) up
	@echo "Esperando a PostgreSQL..."
	@until docker compose exec -T postgres pg_isready -U postgres -d vehiculos > /dev/null 2>&1; do sleep 1; done
	go test ./...

clean:
	docker compose down -v --remove-orphans
