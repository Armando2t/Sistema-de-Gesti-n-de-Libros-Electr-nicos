package controles

import (
	"fmt"
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
)

// Función para registrar un préstamo
func RegistrarPrestamo() {
	var prestamo modelos.Prestamo
	var usuarioID, libroID int
	fmt.Println("Registrar Préstamo:")

	// Solicita el ID del usuario y el libro
	fmt.Print("ID del Usuario: ")
	fmt.Scan(&usuarioID)
	fmt.Print("ID del Libro: ")
	fmt.Scan(&libroID)

	// Validar si el usuario existe
	if usuarioID <= 0 {
		fmt.Println("Usuario no válido.")
		return
	}

	// Validar si el libro existe y está disponible
	if libroID <= 0 {
		fmt.Println("Libro no válido.")
		return
	}

	// Crea un préstamo con los datos proporcionados
	prestamo.UsuarioID = usuarioID
	prestamo.LibroID = libroID

	// Registra el préstamo
	repositorios.RegistrarPrestamo(prestamo)
	fmt.Println("Préstamo registrado exitosamente.")
}
