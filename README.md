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
- `makefile`: comandos para automatizar las tareas del proyecto.
- `.env.example`: ejemplo de configuración de las variables de entorno.
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

## Configuración del entorno

Antes de ejecutar el proyecto, crear el archivo `.env` a partir del archivo de ejemplo:

```bash
cp .env.example .env
```

El archivo `.env` contiene las variables de configuración de PostgreSQL y la URL de conexión utilizada por la aplicación.

Este archivo es local y está excluido del repositorio mediante `.gitignore`. Se debe utilizar `.env.example` como referencia para configurar el entorno.

## Ejecución automática de las pruebas

Desde la raíz del proyecto, ejecutar:

```bash
cp .env.example .env
make test-api
```

Este comando ejecuta el script `test-api.sh`, que realiza las siguientes tareas:

1. Carga las variables de entorno desde `.env`.
2. Elimina los contenedores y volúmenes anteriores del proyecto.
3. Genera el código Go a partir de las consultas SQL mediante `sqlc generate`.
4. Compila la aplicación.
5. Inicia PostgreSQL mediante Docker Compose.
6. Espera a que PostgreSQL esté listo para aceptar conexiones.
7. Inicia la aplicación Go.
8. Espera a que la API esté disponible.
9. Ejecuta las pruebas HTTP de `requests.hurl`.
10. Ejecuta las pruebas HTTP de `test.hurl`.
11. Detiene la aplicación y elimina los contenedores y volúmenes al finalizar.

No es necesario iniciar PostgreSQL ni el servidor manualmente.

**Importante:** la limpieza elimina los datos almacenados en el volumen de PostgreSQL del proyecto. Cada ejecución comienza con una base de datos nueva.

## Ejecución manual

Si se desea iniciar el entorno manualmente, ejecutar desde la raíz del proyecto:

1. Crear el archivo de configuración:

   ```bash
   cp .env.example .env
   ```

2. Cargar las variables de entorno para que estén disponibles para Go:

   ```bash
   set -a
   source .env
   set +a
   ```

3. Generar el código de acceso a datos e iniciar PostgreSQL:

   ```bash
   sqlc generate
   docker compose up -d
   ```

4. Esperar a que PostgreSQL esté listo:

   ```bash
   until docker compose exec -T postgres pg_isready -U "$POSTGRES_USER" -d "$POSTGRES_DB" > /dev/null 2>&1; do
       sleep 1
   done
   ```

5. Iniciar el servidor:

   ```bash
   go run .
   ```

6. En otra terminal, ejecutar las pruebas HTTP:

   ```bash
   hurl --test requests.hurl
   hurl --test test.hurl
   ```

7. Al finalizar, detener PostgreSQL:

   ```bash
   docker compose down
   ```

Este último comando detiene y elimina el contenedor, pero conserva los datos del volumen. Para eliminar también los datos almacenados, utilizar `docker compose down -v`.

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

El proyecto incluye un `makefile` para centralizar las tareas y un script Bash para ejecutar las pruebas de la API de principio a fin.

El comando recomendado para verificar la entrega es:

```bash
make test-api
```

Este comando permite ejecutar las pruebas sin tener que iniciar manualmente cada componente del sistema.
