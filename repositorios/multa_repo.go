// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Repositorio para manejar las operaciones relacionadas con las multas.

package repositorios

import (
	"database/sql"
	"SistemaGestionLibrosElectronicos/modelos"
)

// RegistrarMulta agrega una multa por retraso.
func RegistrarMulta(db *sql.DB, multa modelos.Multa) error {
	query := "INSERT INTO Multa (IDUsuario, Monto, Estado) VALUES (?, ?, ?)"
	_, err := db.Exec(query, multa.IDUsuario, multa.Monto, multa.Estado)
	return err
}

// ConsultarMultas obtiene todas las multas de un usuario.
func ConsultarMultas(db *sql.DB, idUsuario int) ([]modelos.Multa, error) {
	query := "SELECT IDMulta, Monto, Estado FROM Multa WHERE IDUsuario = ?"
	rows, err := db.Query(query, idUsuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var multas []modelos.Multa
	for rows.Next() {
		var multa modelos.Multa
		if err := rows.Scan(&multa.ID, &multa.Monto, &multa.Estado); err != nil {
			return nil, err
		}
		multas = append(multas, multa)
	}
	return multas, nil
}

// RegistrarPago actualiza el estado de una multa como "Pagada".
func RegistrarPago(db *sql.DB, idMulta int) error {
	query := "UPDATE Multa SET Estado = 'Pagada' WHERE IDMulta = ?"
	_, err := db.Exec(query, idMulta)
	return err
}
