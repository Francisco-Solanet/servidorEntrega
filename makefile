.PHONY: run templ sqlc up test test-api clean

run: templ up
	@echo "Esperando a PostgreSQL..."
	@until docker compose exec -T postgres pg_isready -U postgres -d vehiculos > /dev/null 2>&1; do sleep 1; done
	go run .

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

test-api:
	bash ./test-api.sh

clean:
	docker compose down -v --remove-orphans
