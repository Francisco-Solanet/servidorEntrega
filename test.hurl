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

GET http://localhost:8080/vehiculos/{{id}}
HTTP 200
