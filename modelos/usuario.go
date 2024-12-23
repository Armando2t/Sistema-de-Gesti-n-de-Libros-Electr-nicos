// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Definición de la estructura de datos para los usuarios.

package modelos

// Usuario representa un usuario en el sistema.
// Campos:
// - ID: Identificador único del usuario.
// - Nombre: Nombre del usuario (requerido).
// - Apellido: Apellido del usuario (requerido).
// - Correo: Correo electrónico del usuario (requerido, formato válido).
type Usuario struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre" validate:"required"`
	Apellido string `json:"apellido" validate:"required"`
	Correo   string `json:"correo" validate:"required,email"`
}
