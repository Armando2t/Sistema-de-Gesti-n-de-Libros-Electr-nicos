// Autor: Armando Topón
// Fecha: 24/11/2024
// Descripción: Controlador para la lógica de negocio relacionada con la gestión de usuarios.

package controles

import (
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
	"fmt"
)

func RegistrarUsuarioControl(usuario *modelos.Usuario) {
	err := repositorios.RegistrarUsuario(usuario)
	if err != nil {
		fmt.Println("Error al registrar usuario:", err)
	}
}

func ListarUsuariosControl() {
	usuarios, err := repositorios.ListarUsuarios()
	if err != nil {
		fmt.Println("Error al listar usuarios:", err)
		return
	}
	for _, usuario := range usuarios {
		fmt.Printf("ID: %d, Nombre: %s %s, Correo: %s\n", usuario.ID, usuario.Nombre, usuario.Apellido, usuario.Correo)
	}
}

func ActualizarUsuarioControl(usuario *modelos.Usuario) {
	err := repositorios.ActualizarUsuario(usuario)
	if err != nil {
		fmt.Println("Error al actualizar usuario:", err)
	}
}
