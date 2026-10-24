# TP6 - Interfaces Reactivas con HTMX

## Descripción

Aplicación web para gestionar vehículos, desarrollada en Go y conectada a una base de datos PostgreSQL. Utiliza Templ para generar el HTML y HTMX para actualizar partes de la interfaz sin recargar la página completa.

## Funcionalidades

- Consultar la lista de vehículos registrados.
- Registrar vehículos mediante un formulario.
- Actualizar la lista automáticamente después de crear un vehículo.
- Eliminar vehículos directamente desde la tabla.
- Solicitar confirmación antes de eliminar un vehículo.
- Guardar los datos en PostgreSQL.

## Tecnologías utilizadas

- **Go:** servidor HTTP y lógica de la aplicación.
- **Templ:** generación de componentes HTML.
- **HTMX:** solicitudes HTTP y actualización parcial de la interfaz.
- **PostgreSQL:** almacenamiento de los vehículos.
- **SQLc:** generación del código Go para acceder a la base de datos.
- **Docker Compose:** ejecución de PostgreSQL.
- **Pico.css:** estilos de la interfaz.

## Funcionamiento de HTMX

El formulario de registro utiliza `hx-post` para enviar los datos a `/vehiculos`. La respuesta del servidor contiene la lista actualizada, que reemplaza el elemento `#entity-list` mediante `hx-target` y `hx-swap="outerHTML"`.

Cada vehículo tiene un botón de eliminación que utiliza `hx-delete` para enviar una solicitud a `/vehiculos/{id}`. Con `hx-target="closest tr"` y `hx-swap="outerHTML"`, se elimina de la interfaz la fila correspondiente cuando la operación finaliza correctamente.

Las interacciones se realizan sin escribir archivos JavaScript propios. HTMX se incorpora mediante un script externo en el layout.

## Requisitos

- Go.
- Docker y Docker Compose.
- Conexión a Internet para descargar las dependencias y herramientas necesarias.

El Makefile utiliza Templ `v0.2.793` y SQLc `v1.31.1`.

## Ejecución

Desde la raíz del proyecto, ejecutar:

```bash
make run-app
```

El comando:

1. Genera los componentes Templ.
2. Genera el código de acceso a datos mediante SQLc.
3. Detiene los contenedores anteriores y elimina sus volúmenes.
4. Inicia PostgreSQL.
5. Espera a que la base de datos esté disponible.
6. Ejecuta la aplicación Go.
7. Al finalizar la aplicación, detiene los contenedores y elimina sus volúmenes.

Abrir en el navegador:

http://localhost:8080

Para detener la aplicación, presionar `Ctrl+C`.

**Importante:** el comando elimina los volúmenes de Docker al iniciar y al finalizar. Los datos almacenados en esos volúmenes se pierden.

## Pruebas

Para generar el código, iniciar PostgreSQL y ejecutar las pruebas automatizadas:

```bash
make test
```

Para detener los contenedores y eliminar sus volúmenes manualmente:

```bash
make clean
```

## Estructura del proyecto

- `main.go`: servidor HTTP y handlers.
- `views/layout.templ`: estructura general de la página e inclusión de HTMX.
- `views/entity_form.templ`: formulario para registrar vehículos.
- `views/entity_list.templ`: tabla de vehículos y botones de eliminación.
- `views/index.templ`: composición de la página principal.
- `views/*_templ.go`: código generado por Templ.
- `db/`: código generado por SQLc.
- `schema.sql`: estructura de la base de datos.
- `queries.sql`: consultas SQL.
- `sqlc.yaml`: configuración de SQLc.
- `docker-compose.yml`: configuración de PostgreSQL.
- `makefile`: comandos para generar código, ejecutar la aplicación, probar y limpiar.

## Base de datos

La aplicación utiliza la base de datos `vehiculos`, configurada en `docker-compose.yml`. La estructura inicial se define en `schema.sql` y las consultas se encuentran en `queries.sql`.

## Entrega

La implementación del TP6 se encuentra en la rama `tp6`.
