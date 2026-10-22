# TP4 - Interfaz de Usuario

## Descripción

Aplicación web para administrar un catálogo de vehículos. Permite consultar, agregar, modificar y eliminar vehículos mediante una interfaz web conectada a una API desarrollada en Go y una base de datos PostgreSQL.

## Tecnologías

- HTML
- CSS
- JavaScript
- Go
- PostgreSQL
- Docker

## Funcionalidades

- Listar los vehículos registrados.
- Agregar vehículos.
- Modificar los datos de un vehículo.
- Eliminar vehículos.

## Ejecución

1. Iniciar PostgreSQL:

   ```bash
   docker compose up -d
   ```

2. Iniciar el servidor:

   ```bash
   go run .
   ```

3. Abrir `frontend/index.html` en el navegador.

## Estructura

- `frontend/index.html`: estructura de la interfaz.
- `frontend/styles.css`: estilos visuales.
- `frontend/app.js`: lógica y comunicación con la API.
- `main.go`: servidor y endpoints de la API.
- `db/`: código generado por sqlc para acceder a la base de datos.
- `schema.sql`: estructura de la base de datos.
- `docker-compose.yml`: configuración de PostgreSQL.
