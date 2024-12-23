// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Definición de la estructura de datos para los préstamos.

package modelos

// Prestamo representa un préstamo en el sistema.
// Campos:
// - ID: Identificador único del préstamo.
// - IDUsuario: Identificador del usuario asociado (requerido).
// - IDLibro: Identificador del libro prestado (requerido).
// - FechaInicio: Fecha de inicio del préstamo (requerido).
// - FechaFin: Fecha de finalización del préstamo (requerido).
type Prestamo struct {
	ID         int    `json:"id"`
	IDUsuario  int    `json:"id_usuario" validate:"required"`
	IDLibro    int    `json:"id_libro" validate:"required"`
	FechaInicio string `json:"fecha_inicio" validate:"required"`
	FechaFin    string `json:"fecha_fin" validate:"required"`
}
