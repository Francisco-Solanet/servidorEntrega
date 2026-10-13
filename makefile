.PHONY: test test-api sqlc clean up

test:
	$(MAKE) clean
	$(MAKE) sqlc
	$(MAKE) up
	@echo "Esperando a PostgreSQL..."
	@until docker compose exec -T postgres pg_isready -U postgres -d vehiculos > /dev/null 2>&1; do sleep 1; done
	@go test ./...; resultado=$$?; $(MAKE) clean; exit $$resultado

test-api:
	bash ./test-api.sh

sqlc:
	sqlc generate

up:
	docker compose up -d

clean:
	docker compose down -v --remove-orphans
