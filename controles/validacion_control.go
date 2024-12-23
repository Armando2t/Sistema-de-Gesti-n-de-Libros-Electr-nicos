// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Archivo para manejar la instancia global del validador y validar estructuras.

package controles

import "github.com/go-playground/validator/v10"

// Crear una instancia global del validador.
var validate = validator.New()

// ValidarEstructura valida una estructura usando las reglas definidas en las etiquetas de validación.
// Parámetro:
// - i: La estructura que se va a validar.
// Retorna:
// - error: Un error si la validación falla; nil si es válida.
func ValidarEstructura(i interface{}) error {
	return validate.Struct(i)
}
