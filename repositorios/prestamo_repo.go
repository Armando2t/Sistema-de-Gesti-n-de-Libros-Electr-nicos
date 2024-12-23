// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Repositorio para manejar las operaciones relacionadas con los préstamos.

package repositorios

import (
	"database/sql"
	"SistemaGestionLibrosElectronicos/modelos"
)

// RegistrarPrestamo agrega un nuevo préstamo a la base de datos.
func RegistrarPrestamo(db *sql.DB, prestamo modelos.Prestamo) error {
	query := "INSERT INTO Prestamo (IDUsuario, IDLibro, FechaInicio, FechaFin) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, prestamo.IDUsuario, prestamo.IDLibro, prestamo.FechaInicio, prestamo.FechaFin)
	return err
}

// ExtenderPrestamo actualiza la fecha de devolución de un préstamo.
func ExtenderPrestamo(db *sql.DB, idPrestamo int, nuevaFechaFin string) error {
	query := "UPDATE Prestamo SET FechaFin = ? WHERE IDPrestamo = ?"
	_, err := db.Exec(query, nuevaFechaFin, idPrestamo)
	return err
}

// ConsultarHistorial obtiene el historial de préstamos de un usuario.
func ConsultarHistorial(db *sql.DB, idUsuario int) ([]modelos.Prestamo, error) {
	query := "SELECT IDPrestamo, IDLibro, FechaInicio, FechaFin FROM Prestamo WHERE IDUsuario = ?"
	rows, err := db.Query(query, idUsuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prestamos []modelos.Prestamo
	for rows.Next() {
		var prestamo modelos.Prestamo
		if err := rows.Scan(&prestamo.ID, &prestamo.IDLibro, &prestamo.FechaInicio, &prestamo.FechaFin); err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}
	return prestamos, nil
}
