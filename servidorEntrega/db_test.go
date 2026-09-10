package main

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"

	"ejercicioEntrega.com/servidorEntrega/db"
)

func TestQueries_CRUD(t *testing.T) {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=vehiculos sslmode=disable")
	if err != nil {
		t.Fatalf("error al conectar: %v", err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		t.Fatalf("error al conectar con la base: %v", err)
	}

	queries := db.New(conn)
	ctx := context.Background()

	// Crear vehículo
	vehiculo, err := queries.CreateVehiculo(ctx, db.CreateVehiculoParams{
		Marca:       "Toyota",
		Modelo:      "Corolla",
		Anio:        2020,
		PrecioVenta: "15000.00",
		Kilometraje: 50000,
	})
	if err != nil {
		t.Fatalf("error al crear vehículo: %v", err)
	}

	if vehiculo.Marca != "Toyota" {
		t.Errorf("marca incorrecta: %s", vehiculo.Marca)
	}

	// Obtener vehículo
	vehiculo, err = queries.GetVehiculo(ctx, vehiculo.ID)
	if err != nil {
		t.Fatalf("error al obtener vehículo: %v", err)
	}

	if vehiculo.Modelo != "Corolla" {
		t.Errorf("modelo incorrecto: %s", vehiculo.Modelo)
	}

	// Actualizar vehículo
	vehiculo, err = queries.UpdateVehiculo(ctx, db.UpdateVehiculoParams{
		ID:          vehiculo.ID,
		Marca:       "Toyota",
		Modelo:      "Corolla X",
		Anio:        2021,
		PrecioVenta: "16000.00",
		Kilometraje: 55000,
	})
	if err != nil {
		t.Fatalf("error al actualizar vehículo: %v", err)
	}

	if vehiculo.Modelo != "Corolla X" {
		t.Errorf("modelo no actualizado: %s", vehiculo.Modelo)
	}

	// Listar vehículos
	vehiculos, err := queries.ListVehiculos(ctx)
	if err != nil {
		t.Fatalf("error al listar vehículos: %v", err)
	}

	encontrado := false

	for _, v := range vehiculos {
		if v.ID == vehiculo.ID {
			encontrado = true
			break
		}
	}

	if !encontrado {
		t.Error("el vehículo actualizado no aparece en la lista")
	}

	// Eliminar vehículo
	err = queries.DeleteVehiculo(ctx, vehiculo.ID)
	if err != nil {
		t.Fatalf("error al eliminar vehículo: %v", err)
	}

	// Comprobar que fue eliminado
	_, err = queries.GetVehiculo(ctx, vehiculo.ID)
	if err == nil {
		t.Error("el vehículo sigue existiendo después de eliminarlo")
	}
}
