.PHONY: run-app run templ sqlc up test clean

# Prepara el proyecto, ejecuta la aplicación y limpia Docker al finalizar.
run-app:
	@set -e; \
	trap 'docker compose down -v --remove-orphans' EXIT; \
	echo "Generando componentes Templ..."; \
	~/go/bin/templ generate; \
	echo "Generando consultas SQLc..."; \
	sqlc generate; \
	echo "Limpiando contenedores y volumenes anteriores..."; \
	docker compose down -v --remove-orphans; \
	echo "Iniciando PostgreSQL..."; \
	docker compose up -d; \
	echo "Esperando a que PostgreSQL este listo..."; \
	until docker compose exec -T postgres pg_isready -U postgres -d vehiculos > /dev/null 2>&1; do sleep 1; done; \
	echo "Iniciando aplicacion..."; \
	go run .

# Alias para mantener compatibilidad.
run: run-app

templ:
	~/go/bin/templ generate

sqlc:
	sqlc generate

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
