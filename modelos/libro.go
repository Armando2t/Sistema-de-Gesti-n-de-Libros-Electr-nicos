// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Definición de la estructura de datos para los libros.

package modelos

// Libro representa un libro en el sistema.
// Campos:
// - ID: Identificador único del libro.
// - Titulo: Título del libro (requerido).
// - Autor: Autor del libro (requerido).
type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo" validate:"required"`
	Autor  string `json:"autor" validate:"required"`
}

