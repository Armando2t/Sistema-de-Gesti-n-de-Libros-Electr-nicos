package controles

import (
	"fmt"
	"SistemaGestionLibrosElectronicos/modelos"
	"SistemaGestionLibrosElectronicos/repositorios"
)

// Función para registrar un libro
func RegistrarLibro() {
	var libro modelos.Libro
	fmt.Println("Registrar Libro:")
	fmt.Print("Título: ")
	fmt.Scan(&libro.Titulo)
	fmt.Print("Autor: ")
	fmt.Scan(&libro.Autor)
	fmt.Print("Género: ")
	fmt.Scan(&libro.Genero)
	libro.Disponible = true

	// Registra el libro
	repositorios.RegistrarLibro(libro)
	fmt.Println("Libro registrado exitosamente.")
}

// Función para listar todos los libros
func ListarLibros() {
	fmt.Println("Listado de Libros:")
	libros := repositorios.ListarLibros()
	for _, libro := range libros {
		disponibilidad := "Disponible"
		if !libro.Disponible {
			disponibilidad = "No disponible"
		}
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Género: %s, Estado: %s\n",
			libro.ID, libro.Titulo, libro.Autor, libro.Genero, disponibilidad)
	}
}
