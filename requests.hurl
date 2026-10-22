# 1. Listar vehículos
GET http://localhost:8080/vehiculos
HTTP 200


# 2. Crear vehículo
POST http://localhost:8080/vehiculos
Content-Type: application/json

{
  "marca": "Ford",
  "modelo": "Focus",
  "anio": 2019,
  "precio_venta": "12000000",
  "kilometraje": 70000
}

HTTP 201
[Captures]
id: jsonpath "$.ID"
[Asserts]
jsonpath "$.Marca" == "Ford"
jsonpath "$.Modelo" == "Focus"
jsonpath "$.Anio" == 2019


# 3. Obtener vehículo creado
GET http://localhost:8080/vehiculos/{{id}}
HTTP 200
[Asserts]
jsonpath "$.Marca" == "Ford"
jsonpath "$.Modelo" == "Focus"
jsonpath "$.Anio" == 2019


# 4. Actualizar vehículo
PUT http://localhost:8080/vehiculos/{{id}}
Content-Type: application/json

{
  "marca": "Ford",
  "modelo": "Focus Titanium",
  "anio": 2020,
  "precio_venta": "14000000",
  "kilometraje": 60000
}

HTTP 200
[Asserts]
jsonpath "$.Modelo" == "Focus Titanium"
jsonpath "$.Anio" == 2020
jsonpath "$.Kilometraje" == 60000


# 5. Verificar actualización
GET http://localhost:8080/vehiculos/{{id}}
HTTP 200
[Asserts]
jsonpath "$.Modelo" == "Focus Titanium"
jsonpath "$.Anio" == 2020
jsonpath "$.Kilometraje" == 60000


# 6. Eliminar vehículo
DELETE http://localhost:8080/vehiculos/{{id}}
HTTP 204


# 7. Verificar eliminación
GET http://localhost:8080/vehiculos/{{id}}
HTTP 404


# 8. Probar validación de marca vacía
POST http://localhost:8080/vehiculos
Content-Type: application/json

{
  "marca": "",
  "modelo": "Focus",
  "anio": 2019,
  "precio_venta": "12000000",
  "kilometraje": 70000
}

HTTP 400


# 9. Probar ID inexistente
GET http://localhost:8080/vehiculos/999999
HTTP 404
