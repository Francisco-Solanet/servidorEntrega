# TP3 – Exponiendo la Lógica como API

## Descripción

En esta entrega se desarrolla una API REST para gestionar vehículos. El servidor web se conecta a una base de datos PostgreSQL y utiliza `sqlc` para ejecutar las consultas SQL.

La API permite crear, listar, consultar, actualizar y eliminar vehículos.

## Tecnologías utilizadas

- Go
- PostgreSQL
- sqlc
- Docker Compose
- Hurl para las pruebas HTTP

## Estructura del proyecto

- `main.go`: servidor HTTP y handlers de la API.
- `db/`: código generado por sqlc para acceder a la base de datos.
- `schema.sql`: definición de la estructura de la base de datos.
- `queries.sql`: consultas SQL utilizadas por sqlc.
- `sqlc.yaml`: configuración de sqlc.
- `db_test.go`: pruebas de la capa de acceso a datos.
- `requests.hurl`: solicitudes HTTP de ejemplo.
- `test.hurl`: pruebas de la API.
- `docker-compose.yml`: configuración de los servicios.
- `Makefile`: comandos para ejecutar tareas del proyecto.
- `go.mod` y `go.sum`: dependencias del proyecto.
- `Static/index.html`: página web estática, si se utiliza en esta entrega.

## Entidad: Vehículo

Los vehículos contienen los siguientes atributos:

- `id`: identificador del vehículo.
- `marca`: marca del vehículo.
- `modelo`: modelo del vehículo.
- `anio`: año de fabricación.
- `precio_venta`: precio de venta.
- `kilometraje`: kilometraje registrado.

## Endpoints de la API

| Método HTTP | Ruta | Descripción |
|---|---|---|
| POST | `/vehiculos` | Crea un vehículo. |
| GET | `/vehiculos` | Lista todos los vehículos. |
| GET | `/vehiculos/{id}` | Obtiene un vehículo por su ID. |
| PUT | `/vehiculos/{id}` | Actualiza un vehículo existente. |
| DELETE | `/vehiculos/{id}` | Elimina un vehículo. |

## Validaciones y respuestas

- Se devuelve `400 Bad Request` cuando el JSON es inválido, el ID no es válido o faltan la marca o el modelo.
- Se devuelve `404 Not Found` cuando el vehículo solicitado no existe en las operaciones de consulta o actualización.
- Se devuelve `405 Method Not Allowed` cuando se utiliza un método HTTP no permitido para la ruta.
- Se devuelve `500 Internal Server Error` ante errores internos de la base de datos.
- Se devuelve `201 Created` al crear un vehículo correctamente.
- Se devuelve `204 No Content` al eliminar un vehículo correctamente.

## Ejecución

1. Iniciar los servicios configurados en Docker Compose:

   ```bash
   docker compose up -d
   ```

2. Ejecutar el servidor:

   ```bash
   go run main.go
   ```

3. Acceder a la API en `http://localhost:8080`.

## Pruebas

Las pruebas HTTP se pueden ejecutar con Hurl:

```bash
hurl --test requests.hurl
hurl --test test.hurl
```

La ejecución requiere que la base de datos esté disponible y que los servicios necesarios estén iniciados.

