# TP2 – Persistiendo el Dominio

Proyecto realizado en Go para definir la estructura de datos de una aplicación de vehículos y preparar la capa de acceso a una base de datos PostgreSQL.

El objetivo de este trabajo es definir el esquema de la base de datos, escribir las consultas SQL para las operaciones CRUD y generar el código Go mediante `sqlc`.

## Tecnologías utilizadas

- Go
- PostgreSQL
- SQL
- sqlc
- Docker y Docker Compose

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
├── go.mod
└── go.sum
```

## Diseño de la base de datos

Se utiliza PostgreSQL para almacenar los vehículos en una tabla llamada `vehiculos`.

Cada vehículo contiene los siguientes datos:

- ID
- Marca
- Modelo
- Año
- Precio de venta
- Kilometraje

La estructura de la tabla se define en `schema.sql`.

## Consultas SQL

El archivo `queries.sql` contiene las consultas necesarias para realizar las operaciones CRUD:

- Crear un vehículo.
- Obtener un vehículo por su ID.
- Listar los vehículos.
- Actualizar un vehículo.
- Eliminar un vehículo.

Las consultas utilizan las anotaciones `-- name:` requeridas por `sqlc`.

## Generación del código

La configuración de `sqlc` se encuentra en `sqlc.yaml`. Este archivo indica dónde están el esquema y las consultas, y dónde debe generarse el código Go.

Para generar el código de acceso a datos, ejecutar:

```bash
sqlc generate
```

El código generado se encuentra en la carpeta `db/`.

## Pruebas

El proyecto incluye pruebas para verificar las operaciones de acceso a datos.

Si Docker, PostgreSQL y las herramientas necesarias están instalados, se pueden ejecutar las pruebas mediante:

```bash
make test
```

## Objetivo alcanzado

Este trabajo prepara la estructura de la base de datos y la capa de acceso a datos para que puedan utilizarse en los siguientes prácticos.
