package controles

import (
	"fmt"
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
)

// Función para registrar un usuario
func RegistrarUsuario() {
	var usuario modelos.Usuario
	fmt.Println("Registrar Usuario:")
	fmt.Print("Nombre: ")
	fmt.Scan(&usuario.Nombre)
	fmt.Print("Apellido: ")
	fmt.Scan(&usuario.Apellido)
	fmt.Print("Cédula: ")
	fmt.Scan(&usuario.Cedula)
	fmt.Print("Dirección: ")
	fmt.Scan(&usuario.Direccion)
	fmt.Print("Correo: ")
	fmt.Scan(&usuario.Correo)
	fmt.Print("Rol (Admin/Usuario): ")
	fmt.Scan(&usuario.Rol)

	// Registra el usuario
	repositorios.RegistrarUsuario(usuario)
	fmt.Println("Usuario registrado exitosamente.")
}

// Función para listar todos los usuarios
func ListarUsuarios() {
	fmt.Println("Listado de Usuarios:")
	usuarios := repositorios.ListarUsuarios()
	for _, usuario := range usuarios {
		fmt.Printf("ID: %d, Nombre: %s %s, Cédula: %s, Correo: %s, Rol: %s\n",
			usuario.ID, usuario.Nombre, usuario.Apellido, usuario.Cedula, usuario.Correo, usuario.Rol)
	}
}
