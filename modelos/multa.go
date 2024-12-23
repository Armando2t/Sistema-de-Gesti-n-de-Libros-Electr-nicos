// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Definición de la estructura de datos para las multas.

package modelos

// Multa representa una multa en el sistema.
// Campos:
// - ID: Identificador único de la multa.
// - IDUsuario: Identificador del usuario asociado (requerido).
// - Monto: Monto de la multa (requerido).
// - Estado: Estado de la multa (requerido, valores posibles: "Pendiente", "Pagada").
type Multa struct {
	ID        int     `json:"id"`
	IDUsuario int     `json:"id_usuario" validate:"required"`
	Monto     float64 `json:"monto" validate:"required"`
	Estado    string  `json:"estado" validate:"required,oneof=Pendiente Pagada"`
}
