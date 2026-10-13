# TP3 – Exponiendo la Lógica como API

## Descripción

En esta entrega se desarrolla una API REST para gestionar vehículos, conectando el servidor web con la capa de acceso a datos generada mediante sqlc y una base de datos PostgreSQL.

La API permite realizar operaciones CRUD: crear, consultar, actualizar y eliminar vehículos.

## Tecnologías utilizadas

- Go
- PostgreSQL
- sqlc
- Docker y Docker Compose
- Hurl para las pruebas HTTP

## Estructura del proyecto

- `main.go`: servidor HTTP y handlers de la API REST.
- `db/`: código generado por sqlc para acceder a la base de datos.
- `schema.sql`: definición de la tabla de vehículos.
- `queries.sql`: consultas SQL utilizadas por sqlc.
- `sqlc.yaml`: configuración de generación de código.
- `db_test.go`: pruebas de la capa de acceso a datos.
- `requests.hurl`: ejemplos de solicitudes HTTP.
- `test.hurl`: pruebas de la API.
- `docker-compose.yml`: configuración de los servicios.
- `Makefile`: comandos para facilitar tareas de ejecución y pruebas.
- `go.mod` y `go.sum`: dependencias del proyecto.
- `Static/index.html`: página web estática, si se mantiene en esta entrega.

## Entidad: Vehículo

Cada vehículo se representa mediante los siguientes atributos:

- `id`: identificador.
- `marca`: marca del vehículo.
- `modelo`: modelo del vehículo.
- `anio`: año de fabricación.
- `precio_venta`: precio de venta.
- `kilometraje`: kilometraje registrado.

## Endpoints

La API permite realizar las siguientes operaciones:

| Método | Ruta | Operación |
|---|---|---|
| POST | `/vehiculos` | Crear un vehículo |
| GET | `/vehiculos` | Listar vehículos |
| GET | `/vehiculos/{id}` | Obtener un vehículo por ID |
| PUT | `/vehiculos/{id}` | Actualizar un vehículo |
| DELETE | `/vehiculos/{id}` | Eliminar un vehículo |

Las respuestas y los códigos HTTP dependen de cada operación y de sus resultados.

## Ejecución

Se utiliza Docker Compose para facilitar la ejecución de los servicios necesarios.

Desde la raíz del proyecto, iniciar los servicios con:

```bash
docker compose up -d
```

Luego, ejecutar el servidor Go:

```bash
go run main.go
```

El servidor atiende solicitudes en el puerto `8080`, según la configuración de la aplicación.

## Pruebas

Los archivos `requests.hurl` y `test.hurl` contienen ejemplos de solicitudes y pruebas HTTP para verificar el funcionamiento de la API.

Para ejecutar un archivo Hurl:

```bash
hurl --test requests.hurl
```

También se puede ejecutar:

```bash
hurl --test test.hurl
```

Las pruebas deben ejecutarse con la base de datos disponible y el servidor iniciado, cuando los archivos estén configurados para interactuar con el servidor en ejecución.

