// Autor: Armando Topón
// Fecha: 24/11/2024
// Descripción: Simulación de repositorio en memoria para gestionar usuarios.

package repositorios

import (
	"SistemaGestionLibrosElectronicos/modelos"
	"errors"
	"fmt"
)

// Slice para almacenar los usuarios simulando la base de datos
var usuarios = []modelos.Usuario{}
var idCounter = 1

// RegistrarUsuario simula agregar un usuario al "almacenamiento"
func RegistrarUsuario(usuario modelos.Usuario) error {
	usuario.ID = idCounter
	idCounter++
	usuarios = append(usuarios, usuario)
	fmt.Println("Usuario registrado correctamente")
	return nil
}

// ListarUsuarios simula la recuperación de todos los usuarios
func ListarUsuarios() ([]modelos.Usuario, error) {
	if len(usuarios) == 0 {
		return nil, errors.New("no hay usuarios registrados")
	}
	return usuarios, nil
}

// ActualizarUsuario simula la actualización de un usuario por ID
func ActualizarUsuario(id int, nuevosDatos modelos.Usuario) error {
	for i, usuario := range usuarios {
		if usuario.ID == id {
			usuarios[i].Nombre = nuevosDatos.Nombre
			usuarios[i].Apellido = nuevosDatos.Apellido
			usuarios[i].Correo = nuevosDatos.Correo
			fmt.Println("Usuario actualizado correctamente")
			return nil
		}
	}
	return errors.New("usuario no encontrado")
}
