# Gestor de Vehículos - TP5

Aplicación web para gestionar vehículos mediante renderizado del lado del servidor (SSR) con Go y Templ.

## Tecnologías

- Go
- Templ
- PostgreSQL
- SQLc
- Docker Compose
- Pico.css

## Funcionalidades

- Listar los vehículos registrados.
- Registrar vehículos mediante un formulario HTML.
- Guardar los datos en PostgreSQL.
- Actualizar la página después de registrar un vehículo mediante una redirección HTTP (patrón PRG).

## Requisitos

- Go instalado.
- Docker y Docker Compose.
- SQLc instalado.
- Templ instalado en `~/go/bin/templ`.

## Ejecución

Desde la raíz del proyecto, ejecutar:

```bash
make run
```

Abrir en el navegador:

http://localhost:8080

Para detener el servidor, presionar `Ctrl+C`.

## Pruebas

Para ejecutar las pruebas, ejecutar:

```bash
make test
```

## Estructura del proyecto

- `main.go`: servidor HTTP y controladores.
- `views/`: componentes Templ y código Go generado.
- `db/`: código generado por SQLc
