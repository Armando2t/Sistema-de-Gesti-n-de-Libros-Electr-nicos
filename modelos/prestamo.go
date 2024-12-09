package modelos

import "time"

// Estructura de un Préstamo de libro en el sistema.
type Prestamo struct {
	ID         int
	UsuarioID  int
	LibroID    int
	FechaInicio time.Time
	FechaFin   time.Time
}
