# TP3 – Exponiendo la Lógica como API

## Descripción

En esta entrega se desarrolla una API REST para gestionar vehículos. El servidor está implementado en Go y se conecta a una base de datos PostgreSQL mediante consultas generadas con sqlc.

La API permite crear, listar, consultar, actualizar y eliminar vehículos. Además, se incorporan pruebas HTTP con Hurl y un script para automatizar la preparación del entorno, la ejecución de las pruebas y la limpieza posterior.

## Tecnologías utilizadas

- **Go:** implementación del servidor HTTP y la lógica de la API.
- **PostgreSQL:** base de datos para almacenar los vehículos.
- **sqlc:** generación de código Go a partir de consultas SQL.
- **Docker Compose:** creación y administración del contenedor de PostgreSQL.
- **Hurl:** pruebas de los endpoints HTTP.
- **Make:** automatización de tareas del proyecto.
- **Bash:** ejecución automatizada de las pruebas de la API.

## Estructura del proyecto

- `main.go`: servidor HTTP y handlers de la API REST.
- `db/`: código generado por sqlc para acceder a la base de datos.
- `schema.sql`: definición de la tabla de vehículos.
- `queries.sql`: consultas SQL utilizadas por sqlc.
- `sqlc.yaml`: configuración de sqlc.
- `db_test.go`: pruebas de la capa de acceso a datos.
- `requests.hurl`: pruebas de las operaciones de la API.
- `test.hurl`: pruebas HTTP adicionales.
- `test-api.sh`: script de automatización de las pruebas de la API.
- `docker-compose.yml`: configuración del servicio PostgreSQL.
- `Makefile`: comandos para automatizar las tareas del proyecto.
- `go.mod` y `go.sum`: dependencias del proyecto.
- `Static/index.html`: página web estática, si se utiliza en esta entrega.

## Modelo de datos

La entidad `Vehiculo` contiene los siguientes atributos:

| Atributo | Tipo | Descripción |
|---|---|---|
| `id` | SERIAL | Identificador único del vehículo. |
| `marca` | VARCHAR(50) | Marca del vehículo. |
| `modelo` | VARCHAR(50) | Modelo del vehículo. |
| `anio` | INT | Año de fabricación. |
| `precio_venta` | DECIMAL(12,2) | Precio de venta. |
| `kilometraje` | INT | Kilometraje registrado. |

Los campos son obligatorios y el identificador se genera automáticamente en PostgreSQL.

## Endpoints de la API

La API se ejecuta en el puerto `8080`.

| Método HTTP | Ruta | Descripción |
|---|---|---|
| `POST` | `/vehiculos` | Crea un vehículo. |
| `GET` | `/vehiculos` | Lista todos los vehículos. |
| `GET` | `/vehiculos/{id}` | Obtiene un vehículo por su identificador. |
| `PUT` | `/vehiculos/{id}` | Actualiza los datos de un vehículo. |
| `DELETE` | `/vehiculos/{id}` | Elimina un vehículo. |

### Códigos de respuesta

- `200 OK`: consulta o actualización realizada correctamente.
- `201 Created`: vehículo creado correctamente.
- `204 No Content`: eliminación realizada correctamente.
- `400 Bad Request`: JSON inválido, identificador inválido o campos obligatorios faltantes.
- `404 Not Found`: vehículo solicitado no encontrado.
- `405 Method Not Allowed`: método HTTP no permitido para la ruta.
- `500 Internal Server Error`: error interno durante una operación.

## Requisitos previos

Para ejecutar el proyecto se necesita tener instalados:

- Go.
- Docker y Docker Compose.
- sqlc.
- Hurl.
- curl.
- Make y Bash.

## Ejecución automática de las pruebas

Desde la raíz del proyecto, ejecutar:

```bash
make test-api
```

Este comando ejecuta el script `test-api.sh`, que realiza las siguientes tareas:

1. Elimina los contenedores y volúmenes anteriores del proyecto.
2. Genera el código Go a partir de las consultas SQL mediante `sqlc generate`.
3. Compila la aplicación.
4. Inicia PostgreSQL mediante Docker Compose.
5. Espera a que PostgreSQL esté listo para aceptar conexiones.
6. Inicia la aplicación Go.
7. Espera a que la API esté disponible.
8. Ejecuta las pruebas HTTP de `requests.hurl`.
9. Ejecuta las pruebas HTTP de `test.hurl`.
10. Detiene la aplicación y elimina los contenedores y volúmenes al finalizar.

No es necesario iniciar PostgreSQL ni el servidor manualmente.

**Importante:** la limpieza elimina los datos almacenados en el volumen de PostgreSQL del proyecto. Cada ejecución comienza con una base de datos nueva.

## Ejecución manual

Si se desea iniciar el entorno manualmente, ejecutar desde la raíz del proyecto:

1. Generar el código de acceso a datos:

   ```bash
   sqlc generate
   ```

2. Iniciar PostgreSQL:

   ```bash
   docker compose up -d
   ```

3. Esperar a que PostgreSQL esté listo y ejecutar el servidor:

   ```bash
   go run main.go
   ```

4. En otra terminal, ejecutar las pruebas HTTP:

   ```bash
   hurl --test requests.hurl
   hurl --test test.hurl
   ```

5. Al finalizar, detener los contenedores y eliminar los volúmenes:

   ```bash
   docker compose down -v --remove-orphans
   ```

## Pruebas de la API

Las pruebas automatizadas verifican las principales operaciones de la API:

- Listado de vehículos.
- Creación de un vehículo y comprobación de los datos devueltos.
- Consulta de un vehículo por su identificador.
- Actualización y verificación de los datos modificados.
- Eliminación y comprobación de que el vehículo ya no existe.
- Validación de campos obligatorios.
- Consulta de identificadores inexistentes.

Las pruebas se ejecutan mediante Hurl y se pueden lanzar todas juntas con `make test-api`.

## Base de datos

La base de datos se llama `vehiculos` y utiliza PostgreSQL. La configuración de conexión de la aplicación se encuentra en `main.go`, mientras que `docker-compose.yml` define el contenedor y `schema.sql` establece la estructura de la tabla.

## Automatización

El proyecto incluye un `Makefile` para centralizar las tareas y un script Bash para ejecutar las pruebas de la API de principio a fin.

El comando recomendado para verificar la entrega es:

```bash
make test-api
```

Este comando permite ejecutar las pruebas sin tener que iniciar manualmente cada componente del sistema.
