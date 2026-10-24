-- name: CreateVehiculo :one
INSERT INTO vehiculos (
    marca,
    modelo,
    anio,
    precio_venta,
    kilometraje
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetVehiculo :one
SELECT *
FROM vehiculos
WHERE id = $1
LIMIT 1;


-- name: ListVehiculos :many
SELECT *
FROM vehiculos
ORDER BY id;


-- name: UpdateVehiculo :one
UPDATE vehiculos
SET
    marca = $2,
    modelo = $3,
    anio = $4,
    precio_venta = $5,
    kilometraje = $6
WHERE id = $1
RETURNING *;


-- name: DeleteVehiculo :exec
DELETE FROM vehiculos
WHERE id = $1;
