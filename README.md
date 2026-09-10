# servidorEntrega
# Servidor de Vehículos

Proyecto realizado en Go utilizando PostgreSQL y sqlc.

El proyecto incluye un servidor HTTP, una base de datos PostgreSQL y operaciones CRUD para vehículos.

## Tecnologías utilizadas

* Go
* PostgreSQL
* Docker
* Docker Compose
* sqlc
* Make

## Estructura del proyecto

```text
servidorEntrega/
├── db/
│   ├── db.go
│   ├── models.go
│   └── queries.sql.go
├── db_test.go
├── schema.sql
├── queries.sql
├── sqlc.yaml
├── docker-compose.yml
├── Makefile
├── main.go
├── Static/
    ├──index.html
├── go.mod
└── go.sum
```

## Base de datos

La base de datos utiliza PostgreSQL y contiene la tabla `vehiculos`.

Cada vehículo tiene:

* ID
* Marca
* Modelo
* Año
* Precio de venta
* Kilometraje

La base de datos se inicia automáticamente mediante Docker Compose.

## Requisitos

Para ejecutar el proyecto se necesita tener instalado:

* Go
* Docker
* Docker Compose
* sqlc
* Make

## Ejecutar las pruebas

Clonar el repositorio y entrar en la carpeta del proyecto:

```bash
git clone https://github.com/francisco2015solane-web/servidorEntrega.git
cd servidorEntrega
```

Luego ejecutar:

make test

El comando `make test` realiza automáticamente las siguientes tareas:

1. Elimina los contenedores y volúmenes anteriores.
2. Ejecuta `sqlc generate`.
3. Inicia PostgreSQL mediante Docker Compose.
4. Espera a que PostgreSQL esté disponible.
5. Ejecuta los tests con `go test ./...`.
6. Elimina los contenedores y volúmenes al finalizar.

## Tests CRUD

Los tests verifican las operaciones principales sobre la tabla `vehiculos`:

* Crear un vehículo.
* Obtener un vehículo por ID.
* Actualizar un vehículo.
* Listar los vehículos.
* Eliminar un vehículo.
* Comprobar que el vehículo eliminado ya no existe.

## Servidor

El servidor HTTP se ejecuta con:

```bash
go run main.go
```

Por defecto queda disponible en:

http://localhost:8080

