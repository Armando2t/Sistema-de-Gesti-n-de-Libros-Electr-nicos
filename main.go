// Autor: Armando Topón
// Fecha: 08/12/2024
// Descripción: Menú principal para la gestión de usuarios en el sistema.
package main

import (
	"fmt"
	"SistemaGestionLibrosElectronicos/controles"
)

// Función principal que muestra el menú
func main() {
	for {
		menu()
	}
}

// Función para mostrar el menú
func menu() {
	var opcion int
	fmt.Println("\nSistema de Gestión de Libros Electrónicos")
	fmt.Println("1. Registrar Usuario")
	fmt.Println("2. Listar Usuarios")
	fmt.Println("3. Registrar Libro")
	fmt.Println("4. Listar Libros")
	fmt.Println("5. Registrar Préstamo")
	fmt.Println("6. Aplicar Multa")
	fmt.Println("7. Salir")
	fmt.Print("Seleccione una opción: ")
	fmt.Scan(&opcion)

	switch opcion {
	case 1:
		controles.RegistrarUsuario()
	case 2:
		controles.ListarUsuarios()
	case 3:
		controles.RegistrarLibro()
	case 4:
		controles.ListarLibros()
	case 5:
		controles.RegistrarPrestamo()
	case 6:
		controles.AplicarMulta()
	case 7:
		fmt.Println("Saliendo...")
		return
	default:
		fmt.Println("Opción no válida.")
	}
}
