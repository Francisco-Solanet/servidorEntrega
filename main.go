package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	"ejercicioEntrega.com/servidorEntrega/db"
	"ejercicioEntrega.com/servidorEntrega/views"
)

type Handler struct {
	queries *db.Queries
}

func main() {
	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/vehiculos?sslmode=disable",
	)
	if err != nil {
		log.Fatal("Error al abrir la conexión:", err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatal("No se pudo conectar con PostgreSQL:", err)
	}

	handler := &Handler{
		queries: db.New(conn),
	}

	http.HandleFunc("/", handler.indexHandler)
	http.HandleFunc("/vehiculos", handler.vehiculosHandler)
	http.HandleFunc("/vehiculos/", handler.vehiculoIDHandler)

	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// GET / muestra la página con los vehículos.
func (h *Handler) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	vehiculos, err := h.queries.ListVehiculos(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener los vehículos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := views.IndexPage(vehiculos).Render(r.Context(), w); err != nil {
		log.Println("Error al renderizar la página:", err)
	}
}

// POST /vehiculos crea un vehículo y devuelve la lista actualizada.
func (h *Handler) vehiculosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Formulario inválido", http.StatusBadRequest)
		return
	}

	marca := strings.TrimSpace(r.FormValue("marca"))
	modelo := strings.TrimSpace(r.FormValue("modelo"))
	anio, errAnio := strconv.ParseInt(r.FormValue("anio"), 10, 32)
	precio := strings.TrimSpace(r.FormValue("precio_venta"))
	kilometraje, errKilometraje := strconv.ParseInt(r.FormValue("kilometraje"), 10, 32)

	if marca == "" || modelo == "" || errAnio != nil || errKilometraje != nil {
		http.Error(w, "Completá correctamente todos los campos", http.StatusBadRequest)
		return
	}

	if anio < 1900 || anio > 2100 || kilometraje < 0 || precio == "" {
		http.Error(w, "Los valores ingresados no son válidos", http.StatusBadRequest)
		return
	}

	_, err := h.queries.CreateVehiculo(
		r.Context(),
		db.CreateVehiculoParams{
			Marca:       marca,
			Modelo:      modelo,
			Anio:        int32(anio),
			PrecioVenta: precio,
			Kilometraje: int32(kilometraje),
		},
	)
	if err != nil {
		log.Println("Error al crear el vehículo:", err)
		http.Error(w, "Error al guardar el vehículo", http.StatusInternalServerError)
		return
	}

	vehiculos, err := h.queries.ListVehiculos(r.Context())
	if err != nil {
		http.Error(w, "Error al obtener los vehículos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := views.EntityList(vehiculos).Render(r.Context(), w); err != nil {
		log.Println("Error al renderizar la lista:", err)
	}
}

// DELETE /vehiculos/{id} elimina un vehículo.
func (h *Handler) vehiculoIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	idTexto := strings.TrimPrefix(r.URL.Path, "/vehiculos/")
	id, err := strconv.ParseInt(idTexto, 10, 32)
	if err != nil || id <= 0 {
		http.Error(w, "ID de vehículo inválido", http.StatusBadRequest)
		return
	}

	err = h.queries.DeleteVehiculo(r.Context(), int32(id))
	if err != nil {
		log.Println("Error al eliminar el vehículo:", err)
		http.Error(w, "Error al eliminar el vehículo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
