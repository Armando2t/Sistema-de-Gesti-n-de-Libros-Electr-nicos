// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Repositorio para manejar las operaciones relacionadas con los libros.

package repositorios

import (
	"database/sql"
	"SistemaGestionLibrosElectronicos/modelos"
)

// RegistrarLibro agrega un nuevo libro a la base de datos.
func RegistrarLibro(db *sql.DB, libro modelos.Libro) error {
	query := "INSERT INTO Libro (Titulo, Autor) VALUES (?, ?)"
	_, err := db.Exec(query, libro.Titulo, libro.Autor)
	return err
}

// BuscarLibros busca libros por título o autor.
func BuscarLibros(db *sql.DB, criterio string) ([]modelos.Libro, error) {
	query := "SELECT IDLibro, Titulo, Autor FROM Libro WHERE Titulo LIKE ? OR Autor LIKE ?"
	rows, err := db.Query(query, "%"+criterio+"%", "%"+criterio+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []modelos.Libro
	for rows.Next() {
		var libro modelos.Libro
		if err := rows.Scan(&libro.ID, &libro.Titulo, &libro.Autor); err != nil {
			return nil, err
		}
		libros = append(libros, libro)
	}
	return libros, nil
}

// ActualizarLibro actualiza la información de un libro existente.
func ActualizarLibro(db *sql.DB, libro modelos.Libro) error {
	query := "UPDATE Libro SET Titulo = ?, Autor = ? WHERE IDLibro = ?"
	_, err := db.Exec(query, libro.Titulo, libro.Autor, libro.ID)
	return err
}
