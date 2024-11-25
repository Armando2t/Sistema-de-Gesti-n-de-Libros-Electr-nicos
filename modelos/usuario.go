// Autor: Armando Topón
// Fecha: 24/11/2024
// Descripción: Definición de la estructura 'Usuario' con sus campos necesarios para la gestión de usuarios.

package modelos

type Usuario struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Cedula    string `json:"cedula"`
	Direccion string `json:"direccion"`
	Correo    string `json:"correo"`
	Rol       string `json:"rol"`
}
