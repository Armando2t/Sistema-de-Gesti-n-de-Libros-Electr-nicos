// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Controlador para manejar las operaciones relacionadas con los usuarios.

package controles

import (
	"encoding/json"
	"net/http"
	"SistemaGestionLibrosElectronicos/config"
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
)

// ManejarUsuarios gestiona las solicitudes relacionadas con los usuarios.

func ManejarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := config.ConectarBD()
	if err != nil {
		http.Error(w, "Error al conectar con la base de datos", http.StatusInternalServerError)
		return
	}
	defer config.CerrarBD(db)

	switch r.Method {
	case "POST":
		var usuario modelos.Usuario
		if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		// Validar los datos del usuario.
		if err := ValidarEstructura(usuario); err != nil {
			http.Error(w, "Error de validación: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := repositorios.RegistrarUsuario(db, usuario); err != nil {
			http.Error(w, "Error al registrar usuario", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode("Usuario registrado correctamente")

	case "PUT":
		var usuario modelos.Usuario
		if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}

		// Validar los datos antes de actualizar.
		if err := ValidarEstructura(usuario); err != nil {
			http.Error(w, "Error de validación: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := repositorios.ActualizarUsuario(db, usuario); err != nil {
			http.Error(w, "Error al actualizar usuario", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Usuario actualizado correctamente")

	case "GET":
		usuarios, err := repositorios.ListarUsuarios(db)
		if err != nil {
			http.Error(w, "Error al listar usuarios", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(usuarios)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

