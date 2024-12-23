// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Controlador para manejar las operaciones relacionadas con los libros.

package controles

import (
	"encoding/json"
	"net/http"
	"SistemaGestionLibrosElectronicos/config"
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
)

// ManejarLibros gestiona las solicitudes relacionadas con los libros.
// Métodos soportados:
// - POST: Registrar un nuevo libro.
// - PUT: Actualizar información de un libro existente.
// - GET: Buscar libros por título o autor.
func ManejarLibros(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConectarBD()
	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	defer config.CerrarBD(db)

	switch r.Method {
	case "POST":
		var libro modelos.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		// Validar los datos del libro.
		if err := ValidarEstructura(libro); err != nil {
			http.Error(w, "Error de validación: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := repositorios.RegistrarLibro(db, libro); err != nil {
			http.Error(w, "Error al registrar libro", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Libro registrado correctamente")

	case "PUT":
		var libro modelos.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		// Validar los datos antes de actualizar.
		if err := ValidarEstructura(libro); err != nil {
			http.Error(w, "Error de validación: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := repositorios.ActualizarLibro(db, libro); err != nil {
			http.Error(w, "Error al actualizar libro", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Libro actualizado correctamente")

	case "GET":
		criterio := r.URL.Query().Get("criterio")
		libros, err := repositorios.BuscarLibros(db, criterio)
		if err != nil {
			http.Error(w, "Error al buscar libros", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(libros)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
