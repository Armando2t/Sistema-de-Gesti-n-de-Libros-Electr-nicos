package controles

import (
	"fmt"
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
	"time"
)

// Función para aplicar una multa a un préstamo
func AplicarMulta() {
	var prestamoID int
	fmt.Println("Aplicar Multa:")
	fmt.Print("ID del préstamo: ")
	fmt.Scan(&prestamoID)

	// Verificar si el préstamo existe
	if prestamoID <= 0 {
		fmt.Println("Préstamo no válido.")
		return
	}

	prestamo := repositorios.ListarPrestamos()[prestamoID-1]
	if time.Now().After(prestamo.FechaFin) {
		// Aplica una multa si el plazo ha expirado
		multa := modelos.Multa{
			UsuarioID: prestamo.UsuarioID,
			Monto:     50.0, // Monto fijo para la multa
			Estado:    "Pendiente",
		}
		repositorios.RegistrarMulta(multa)
		fmt.Println("Multa aplicada.")
	} else {
		fmt.Println("No hay multa, el préstamo está dentro del plazo.")
	}
}
