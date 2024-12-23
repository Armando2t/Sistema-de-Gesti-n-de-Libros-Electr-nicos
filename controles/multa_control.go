// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Controlador para manejar las operaciones relacionadas con las multas.

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

// ManejarMultas gestiona las solicitudes relacionadas con las multas.
// Métodos soportados:
// - POST: Registrar una nueva multa.
// - PUT: Registrar el pago de una multa existente.
// - GET: Consultar las multas asociadas a un usuario.
func ManejarMultas(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConectarBD()
	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	defer config.CerrarBD(db)

	switch r.Method {
	case "POST":
		var multa modelos.Multa
		if err := json.NewDecoder(r.Body).Decode(&multa); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		// Validar los datos de la multa.
		if err := ValidarEstructura(multa); err != nil {
			http.Error(w, "Error de validación: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := repositorios.RegistrarMulta(db, multa); err != nil {
			http.Error(w, "Error al registrar multa", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Multa registrada correctamente")

	case "PUT":
		idMulta, err := strconv.Atoi(r.URL.Query().Get("idMulta"))
		if err != nil {
			http.Error(w, "ID de multa inválido", http.StatusBadRequest)
			return
		}

		if err := repositorios.RegistrarPago(db, idMulta); err != nil {
			http.Error(w, "Error al registrar pago de multa", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Pago de multa registrado correctamente")

	case "GET":
		idUsuario, err := strconv.Atoi(r.URL.Query().Get("idUsuario"))
		if err != nil {
			http.Error(w, "ID de usuario inválido", http.StatusBadRequest)
			return
		}

		multas, err := repositorios.ConsultarMultas(db, idUsuario)
		if err != nil {
			http.Error(w, "Error al consultar multas", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(multas)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// VerHistorialMultas maneja la consulta del historial de multas desde un formulario web.
func VerHistorialMultas(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/multas_form.tmpl")
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

	multas, err := repositorios.ConsultarMultas(db, idUsuario)
	if err != nil {
		http.Error(w, "Error al consultar multas", http.StatusInternalServerError)
		return
	}

	var pendientes, pagadas []modelos.Multa
	for _, multa := range multas {
		if multa.Estado == "Pendiente" {
			pendientes = append(pendientes, multa)
		} else if multa.Estado == "Pagada" {
			pagadas = append(pagadas, multa)
		}
	}

	// Datos dinámicos para la plantilla.
	data := struct {
		Usuario    int
		Pendientes []modelos.Multa
		Pagadas    []modelos.Multa
	}{
		Usuario:    idUsuario,
		Pendientes: pendientes,
		Pagadas:    pagadas,
	}

	tmpl.Execute(w, data)
}
