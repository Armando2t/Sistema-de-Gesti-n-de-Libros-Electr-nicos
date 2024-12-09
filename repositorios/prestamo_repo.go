package repositorios

import (
	"SistemaGestionLibrosElectronicos/modelos"
	"time"
)

// Lista temporal de préstamos
var prestamos []modelos.Prestamo

// Función para registrar un préstamo de libro
func RegistrarPrestamo(prestamo modelos.Prestamo) {
	prestamo.ID = len(prestamos) + 1  // Asigna un ID único basado en el tamaño de la lista
	prestamo.FechaInicio = time.Now()
	prestamo.FechaFin = prestamo.FechaInicio.Add(5 * 24 * time.Hour) // 5 días de préstamo
	prestamos = append(prestamos, prestamo)
}

// Función para listar todos los préstamos
func ListarPrestamos() []modelos.Prestamo {
	return prestamos
}
