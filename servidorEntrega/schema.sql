CREATE TABLE vehiculos (
    id SERIAL PRIMARY KEY,
    marca VARCHAR(50) NOT NULL,
    modelo VARCHAR(50) NOT NULL,
    anio INT NOT NULL,
    precio_venta DECIMAL(12,2) NOT NULL,
    kilometraje INT NOT NULL
);
