#!/usr/bin/env bash
set -Eeuo pipefail

BIN="./.test-api-bin"
LOG="./.test-api.log"
APP_PID=""

# Limpiar recursos incluso si falla alguna prueba.
cleanup() {
    resultado=$?
    trap - EXIT

    if [[ -n "$APP_PID" ]]; then
        kill "$APP_PID" 2>/dev/null || true
        wait "$APP_PID" 2>/dev/null || true
    fi

    docker compose down -v --remove-orphans || true
    rm -f "$BIN" "$LOG"

    exit "$resultado"
}

trap cleanup EXIT

echo "==> Limpiando contenedores y volumenes anteriores"
docker compose down -v --remove-orphans

echo "==> Generando codigo con sqlc"
sqlc generate

echo "==> Compilando la aplicacion"
go build -o "$BIN" .

echo "==> Iniciando PostgreSQL"
docker compose up -d

echo "==> Esperando a que PostgreSQL este listo"
db_lista=false

for intento in {1..60}; do
    if docker compose exec -T postgres \
        pg_isready -U postgres -d vehiculos >/dev/null 2>&1; then
        db_lista=true
        break
    fi

    sleep 1
done

if [[ "$db_lista" != true ]]; then
    echo "ERROR: PostgreSQL no estuvo listo a tiempo"
    exit 1
fi

echo "==> Iniciando la aplicacion"
"$BIN" >"$LOG" 2>&1 &
APP_PID=$!

echo "==> Esperando a que la API este disponible"
api_lista=false

for intento in {1..60}; do
    if curl -fsS http://localhost:8080/vehiculos >/dev/null 2>&1; then
        api_lista=true
        break
    fi

    if ! kill -0 "$APP_PID" 2>/dev/null; then
        echo "ERROR: La aplicacion se cerro inesperadamente"
        cat "$LOG"
        exit 1
    fi

    sleep 1
done

if [[ "$api_lista" != true ]]; then
    echo "ERROR: La API no estuvo disponible a tiempo"
    cat "$LOG"
    exit 1
fi

echo "==> Ejecutando pruebas de la API: requests.hurl"
hurl --test requests.hurl

echo "==> Ejecutando pruebas adicionales: test.hurl"
hurl --test test.hurl

echo "==> Todas las pruebas de la API finalizaron correctamente"
