const API_URL = "http://localhost:8080/vehiculos";

const formulario = document.getElementById("form-vehiculo");
const listaVehiculos = document.getElementById("lista-vehiculos");
const mensaje = document.getElementById("mensaje");
const btnGuardar = document.getElementById("btn-guardar");
const btnCancelar = document.getElementById("btn-cancelar");

// Guarda el ID del vehículo que se está modificando.
let vehiculoEnEdicion = null;

// Obtiene los vehículos de la API y actualiza la lista.
async function cargarVehiculos() {
    try {
        const respuesta = await fetch(API_URL);

        if (!respuesta.ok) {
            throw new Error("No se pudieron obtener los vehículos.");
        }

        const vehiculos = await respuesta.json();
        listaVehiculos.innerHTML = "";

        if (!vehiculos || vehiculos.length === 0) {
            listaVehiculos.textContent = "No hay vehículos registrados.";
            return;
        }

        vehiculos.forEach((vehiculo) => {
            const tarjeta = document.createElement("article");
            tarjeta.classList.add("vehiculo");

            const titulo = document.createElement("h3");
            titulo.textContent = `${vehiculo.Marca} ${vehiculo.Modelo}`;

            const datos = document.createElement("p");
            datos.textContent =
                `ID: ${vehiculo.ID} | Año: ${vehiculo.Anio} | ` +
                `Precio: $${vehiculo.PrecioVenta} | ` +
                `Kilometraje: ${vehiculo.Kilometraje} km`;

            const botonModificar = document.createElement("button");
            botonModificar.type = "button";
            botonModificar.classList.add("modificar");
            botonModificar.dataset.id = vehiculo.ID;
            botonModificar.textContent = "Modificar";

            const botonEliminar = document.createElement("button");
            botonEliminar.type = "button";
            botonEliminar.classList.add("eliminar");
            botonEliminar.dataset.id = vehiculo.ID;
            botonEliminar.textContent = "Eliminar";

            tarjeta.append(titulo, datos, botonModificar, botonEliminar);
            listaVehiculos.appendChild(tarjeta);
        });
    } catch (error) {
        listaVehiculos.textContent =
            "Error al cargar los vehículos. Verificá que el servidor esté funcionando.";
        console.error(error);
    }
}

// Carga los datos de un vehículo en el formulario.
function modificarVehiculo(vehiculo) {
    vehiculoEnEdicion = vehiculo.ID;

    formulario.elements["marca"].value = vehiculo.Marca;
    formulario.elements["modelo"].value = vehiculo.Modelo;
    formulario.elements["anio"].value = vehiculo.Anio;
    formulario.elements["precio_venta"].value = vehiculo.PrecioVenta;
    formulario.elements["kilometraje"].value = vehiculo.Kilometraje;

    btnGuardar.textContent = "Guardar cambios";
    btnCancelar.hidden = false;
    mensaje.textContent = `Modificando el vehículo ID ${vehiculo.ID}.`;

    formulario.scrollIntoView({ behavior: "smooth" });
}

// Vuelve al modo de agregar vehículos.
function cancelarModificacion() {
    vehiculoEnEdicion = null;
    formulario.reset();

    btnGuardar.textContent = "Agregar vehículo";
    btnCancelar.hidden = true;
    mensaje.textContent = "";
}

// Agrega un vehículo o guarda sus modificaciones.
formulario.addEventListener("submit", async (evento) => {
    evento.preventDefault();

    const vehiculo = {
        marca: formulario.elements["marca"].value.trim(),
        modelo: formulario.elements["modelo"].value.trim(),
        anio: Number(formulario.elements["anio"].value),
        precio_venta: formulario.elements["precio_venta"].value,
        kilometraje: Number(formulario.elements["kilometraje"].value)
    };

    const modificando = vehiculoEnEdicion !== null;
    const url = modificando
        ? `${API_URL}/${vehiculoEnEdicion}`
        : API_URL;

    try {
        const respuesta = await fetch(url, {
            method: modificando ? "PUT" : "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(vehiculo)
        });

        if (!respuesta.ok) {
            const detalle = await respuesta.text();
            throw new Error(detalle || "No se pudo guardar el vehículo.");
        }

        cancelarModificacion();
        mensaje.textContent = modificando
            ? "Vehículo modificado correctamente."
            : "Vehículo agregado correctamente.";

        await cargarVehiculos();
    } catch (error) {
        mensaje.textContent = modificando
            ? "Error al modificar el vehículo."
            : "Error al agregar el vehículo.";

        console.error(error);
    }
});

// Gestiona los botones de las tarjetas.
listaVehiculos.addEventListener("click", async (evento) => {
    const botonModificar = evento.target.closest(".modificar");
    const botonEliminar = evento.target.closest(".eliminar");

    if (botonModificar) {
        const id = Number(botonModificar.dataset.id);

        try {
            const respuesta = await fetch(`${API_URL}/${id}`);

            if (!respuesta.ok) {
                throw new Error("No se pudo obtener el vehículo.");
            }

            const vehiculo = await respuesta.json();
            modificarVehiculo(vehiculo);
        } catch (error) {
            mensaje.textContent = "Error al cargar el vehículo para modificar.";
            console.error(error);
        }

        return;
    }

    if (botonEliminar) {
        const id = botonEliminar.dataset.id;

        try {
            const respuesta = await fetch(`${API_URL}/${id}`, {
                method: "DELETE"
            });

            if (!respuesta.ok) {
                throw new Error("No se pudo eliminar el vehículo.");
            }

            if (String(vehiculoEnEdicion) === id) {
                cancelarModificacion();
            }

            mensaje.textContent = "Vehículo eliminado correctamente.";
            await cargarVehiculos();
        } catch (error) {
            mensaje.textContent = "Error al eliminar el vehículo.";
            console.error(error);
        }
    }
});

// Cancela la edición.
btnCancelar.addEventListener("click", cancelarModificacion);

// Carga los vehículos al abrir la página.
cargarVehiculos();
