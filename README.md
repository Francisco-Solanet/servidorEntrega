# Gestor de Vehículos - TP5

## Descripción

Aplicación web para gestionar vehículos mediante renderizado del lado del servidor (SSR), desarrollada con Go y Templ.

La aplicación permite consultar los vehículos registrados y agregar nuevos vehículos mediante formularios HTML. Los datos se almacenan en una base de datos PostgreSQL.

Se utiliza el patrón Post/Redirect/Get (PRG) para actualizar la página después de registrar un vehículo.

## Tecnologías utilizadas

- **Go:** lenguaje utilizado para implementar el servidor.
- **Templ:** generación de componentes HTML desde Go.
- **PostgreSQL:** almacenamiento de los vehículos.
- **sqlc:** generación de código Go a partir de consultas SQL.
- **Docker Compose:** administración del contenedor de PostgreSQL.
- **Pico.css:** estilos de la interfaz web.
- **Make:** automatización de tareas del proyecto.

## Funcionalidades

- Listar los vehículos registrados.
- Registrar vehículos mediante un formulario HTML.
- Guardar los datos en PostgreSQL.
- Actualizar la página después de registrar un vehículo mediante una redirección HTTP (patrón PRG).
- Renderizar las páginas en el servidor mediante Templ.

## Requisitos previos

Para ejecutar el proyecto se necesita tener instalado:

- Go.
- Docker y Docker Compose.
- sqlc.
- Templ.
- Make.

Templ debe estar instalado en `~/go/bin/templ`, ya que el Makefile utiliza esa ruta para generar los componentes.

## Ejecución de la aplicación

Desde la raíz del proyecto, ejecutar:

```bash
make run-app
```

Este comando realiza automáticamente las siguientes tareas:

1. Genera los componentes HTML mediante Templ.
2. Genera el código de acceso a datos mediante sqlc.
3. Elimina los contenedores y volúmenes anteriores del proyecto.
4. Inicia PostgreSQL mediante Docker Compose.
5. Espera a que PostgreSQL esté listo para aceptar conexiones.
6. Ejecuta la aplicación Go.
7. Al finalizar la aplicación, elimina los contenedores y volúmenes mediante Docker Compose.

Una vez iniciada la aplicación, abrir en el navegador:

http://localhost:8080

Para detener el servidor, presionar `Ctrl+C`.

**Importante:** la ejecución de `make run-app` elimina los volúmenes anteriores de PostgreSQL. Por lo tanto, los datos almacenados se pierden al iniciar una nueva ejecución y al finalizarla.

## Pruebas

Para ejecutar las pruebas del proyecto, utilizar:

```bash
make test
```

Este comando genera los componentes Templ, genera las consultas SQL mediante sqlc, inicia PostgreSQL, espera a que esté disponible y ejecuta las pruebas de Go.

## Estructura del proyecto

- `main.go`: servidor HTTP y controladores de la aplicación.
- `views/`: componentes Templ y código Go generado.
- `db/`: código generado por sqlc para acceder a la base de datos.
- `schema.sql`: estructura de la base de datos.
- `queries.sql`: consultas SQL utilizadas por sqlc.
- `sqlc.yaml`: configuración de sqlc.
- `docker-compose.yml`: configuración del servicio PostgreSQL.
- `makefile`: comandos para automatizar la ejecución, generación y pruebas.
- `go.mod` y `go.sum`: dependencias del proyecto.
