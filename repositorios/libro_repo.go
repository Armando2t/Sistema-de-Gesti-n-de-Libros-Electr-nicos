package repositorios

import (
	"SistemaGestionLibrosElectronicos/modelos"
)

// Lista temporal de libros
var libros []modelos.Libro

// Función para registrar un libro nuevo en la lista
func RegistrarLibro(libro modelos.Libro) {
	libro.ID = len(libros) + 1  // Asigna un ID único basado en el tamaño de la lista
	libros = append(libros, libro)
}

// Función para listar todos los libros
func ListarLibros() []modelos.Libro {
	return libros
}

// Función para actualizar la información de un libro
func ActualizarLibro(id int, libro modelos.Libro) {
	for i, l := range libros {
		if l.ID == id {
			libros[i] = libro
			break
		}
	}
}
