package repositorios

import (
	"SistemaGestionLibrosElectronicos/modelos"
)

// Lista temporal de usuarios
var usuarios []modelos.Usuario

// Función para registrar un usuario nuevo en la lista
func RegistrarUsuario(usuario modelos.Usuario) {
	usuario.ID = len(usuarios) + 1  // Asigna un ID único basado en el tamaño de la lista
	usuarios = append(usuarios, usuario)
}

// Función para listar todos los usuarios
func ListarUsuarios() []modelos.Usuario {
	return usuarios
}

// Función para actualizar la información de un usuario
func ActualizarUsuario(id int, usuario modelos.Usuario) {
	for i, u := range usuarios {
		if u.ID == id {
			usuarios[i] = usuario
			break
		}
	}
}
