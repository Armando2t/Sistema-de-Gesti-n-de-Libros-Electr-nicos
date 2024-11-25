// Autor: Armando Topón
// Fecha: 24/11/2024
// Descripción: Probamos las funciones sin base de datos.

package main

import (
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
	"fmt"
)

func main() {
	//Registrar un usuario
	usuario := modelos.Usuario{
		Nombre:    "Armando",
		Apellido:  "Topon",
		Cedula:    "1720671591",
		Direccion: "Sangolqui",
		Correo:    "armando@local.com",
		Rol:       "Lector",
	}
	repositorios.RegistrarUsuario(usuario)

	//Listar todos los usuarios
	usuarios, err := repositorios.ListarUsuarios()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Lista de usuarios:")
		for _, u := range usuarios {
			fmt.Printf("ID: %d, Nombre: %s %s, Correo: %s\n", u.ID, u.Nombre, u.Apellido, u.Correo)
		}
	}

	//Actualizar un usuario
	nuevosDatos := modelos.Usuario{
		Nombre:   "Pedro",
		Apellido: "Jacome",
		Correo:   "pedro@local.com",
	}
	err = repositorios.ActualizarUsuario(1, nuevosDatos)
	if err != nil {
		fmt.Println("Error al actualizar:", err)
	}

	//Listar usuarios nuevamente para ver los cambios
	usuarios, _ = repositorios.ListarUsuarios()
	fmt.Println("Lista de usuarios después de la actualización:")
	for _, u := range usuarios {
		fmt.Printf("ID: %d, Nombre: %s %s, Correo: %s\n", u.ID, u.Nombre, u.Apellido, u.Correo)
	}
}
