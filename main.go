package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	"ejercicioEntrega.com/servidorEntrega/db"
)

type Handler struct {
	queries *db.Queries
}

type VehiculoRequest struct {
	Marca       string `json:"marca"`
	Modelo      string `json:"modelo"`
	Anio        int32  `json:"anio"`
	PrecioVenta string `json:"precio_venta"`
	Kilometraje int32  `json:"kilometraje"`
}

func main() {
	// Conexión con PostgreSQL
	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/vehiculos?sslmode=disable",
	)

	if err != nil {
		log.Fatal("Error al abrir la conexión:", err)
	}

	defer conn.Close()

	// Comprobar conexión
	if err := conn.Ping(); err != nil {
		log.Fatal("No se pudo conectar con PostgreSQL:", err)
	}

	fmt.Println("Conectado a PostgreSQL")

	// Crear repositorio generado por sqlc
	queries := db.New(conn)

	fmt.Println("Repositorio sqlc creado")

	// Crear handler
	handler := &Handler{
		queries: queries,
	}

	// Rutas
	http.HandleFunc("/vehiculos", handler.vehiculosHandler)
	http.HandleFunc("/vehiculos/", handler.vehiculosHandler)

	fmt.Println("Servidor iniciado en http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux)))
}

// Permite las peticiones del frontend y responde las solicitudes CORS.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Decide qué operación realizar según el método HTTP.
func (h *Handler) vehiculosHandler(w http.ResponseWriter, r *http.Request) {
	// Si la ruta es /vehiculos
	if r.URL.Path == "/vehiculos" {
		switch r.Method {
		case http.MethodGet:
			h.listarVehiculos(w, r)
		case http.MethodPost:
			h.crearVehiculo(w, r)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	// Si la ruta es /vehiculos/{id}
	if strings.HasPrefix(r.URL.Path, "/vehiculos/") {
		idTexto := strings.TrimPrefix(r.URL.Path, "/vehiculos/")

		id, err := strconv.ParseInt(idTexto, 10, 32)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.obtenerVehiculo(w, r, int32(id))
		case http.MethodPut:
			h.actualizarVehiculo(w, r, int32(id))
		case http.MethodDelete:
			h.eliminarVehiculo(w, r, int32(id))
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
		return
	}

	http.NotFound(w, r)
}

// POST /vehiculos
func (h *Handler) crearVehiculo(w http.ResponseWriter, r *http.Request) {
	var request VehiculoRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if request.Marca == "" || request.Modelo == "" {
		http.Error(w, "La marca y el modelo son obligatorios", http.StatusBadRequest)
		return
	}

	vehiculo, err := h.queries.CreateVehiculo(
		context.Background(),
		db.CreateVehiculoParams{
			Marca:       request.Marca,
			Modelo:      request.Modelo,
			Anio:        request.Anio,
			PrecioVenta: request.PrecioVenta,
			Kilometraje: request.Kilometraje,
		},
	)

	if err != nil {
		http.Error(w, "Error al crear el vehículo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehiculo)
}

// GET /vehiculos
func (h *Handler) listarVehiculos(w http.ResponseWriter, r *http.Request) {
	vehiculos, err := h.queries.ListVehiculos(context.Background())
	if err != nil {
		http.Error(w, "Error al obtener los vehículos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehiculos)
}

// GET /vehiculos/{id}
func (h *Handler) obtenerVehiculo(w http.ResponseWriter, r *http.Request, id int32) {
	vehiculo, err := h.queries.GetVehiculo(context.Background(), id)

	if err == sql.ErrNoRows {
		http.Error(w, "Vehículo no encontrado", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Error al obtener el vehículo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehiculo)
}

// PUT /vehiculos/{id}
func (h *Handler) actualizarVehiculo(w http.ResponseWriter, r *http.Request, id int32) {
	var request VehiculoRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if request.Marca == "" || request.Modelo == "" {
		http.Error(w, "La marca y el modelo son obligatorios", http.StatusBadRequest)
		return
	}

	vehiculo, err := h.queries.UpdateVehiculo(
		context.Background(),
		db.UpdateVehiculoParams{
			ID:          id,
			Marca:       request.Marca,
			Modelo:      request.Modelo,
			Anio:        request.Anio,
			PrecioVenta: request.PrecioVenta,
			Kilometraje: request.Kilometraje,
		},
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Vehículo no encontrado", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Error al actualizar el vehículo", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehiculo)
}

// DELETE /vehiculos/{id}
func (h *Handler) eliminarVehiculo(w http.ResponseWriter, r *http.Request, id int32) {
	err := h.queries.DeleteVehiculo(context.Background(), id)

	if err != nil {
		http.Error(w, "Error al eliminar el vehículo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
