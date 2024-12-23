// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Controlador para manejar las operaciones relacionadas con los préstamos.

package controles

import (
	"encoding/json"
	"html/template"
	"net/http"
	"SistemaGestionLibrosElectronicos/config"
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
	"strconv"
)

func ManejarPrestamos(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConectarBD()
	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	defer config.CerrarBD(db)

	switch r.Method {
	case "POST":
		var prestamo modelos.Prestamo
		if err := json.NewDecoder(r.Body).Decode(&prestamo); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		if err := ValidarEstructura(prestamo); err != nil {
			http.Error(w, "Error de validación: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := repositorios.RegistrarPrestamo(db, prestamo); err != nil {
			http.Error(w, "Error al registrar préstamo", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Préstamo registrado correctamente")

	case "PUT":
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		nuevaFecha := r.URL.Query().Get("nuevaFecha")
		if err := repositorios.ExtenderPrestamo(db, id, nuevaFecha); err != nil {
			http.Error(w, "Error al extender préstamo", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Préstamo extendido correctamente")

	case "GET":
		idUsuarioStr := r.URL.Query().Get("idUsuario")
		idUsuario, err := strconv.Atoi(idUsuarioStr)
		if err != nil {
			http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
			return
		}

		prestamos, err := repositorios.ConsultarHistorial(db, idUsuario)
		if err != nil {
			http.Error(w, "Error al consultar historial de préstamos", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(prestamos)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// VerPrestamosActivos maneja la consulta de préstamos activos desde un formulario web.
func VerPrestamosActivos(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/prestamos_form.tmpl")
	if err != nil {
		http.Error(w, "Error al cargar la plantilla", http.StatusInternalServerError)
		return
	}

	idUsuarioStr := r.URL.Query().Get("idUsuario")
	if idUsuarioStr == "" {
		// Si no se ingresó un ID, renderizar el formulario vacío.
		tmpl.Execute(w, nil)
		return
	}

	idUsuario, err := strconv.Atoi(idUsuarioStr)
	if err != nil {
		http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
		return
	}

	db, err := config.ConectarBD()
	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	defer config.CerrarBD(db)

	prestamos, err := repositorios.ConsultarHistorial(db, idUsuario)
	if err != nil {
		http.Error(w, "Error al consultar préstamos", http.StatusInternalServerError)
		return
	}

	// Datos dinámicos para la plantilla.
	data := struct {
		Usuario   int
		Prestamos []modelos.Prestamo
	}{
		Usuario:   idUsuario,
		Prestamos: prestamos,
	}

	tmpl.Execute(w, data)
}
