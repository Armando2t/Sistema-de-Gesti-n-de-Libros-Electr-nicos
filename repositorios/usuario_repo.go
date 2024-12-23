// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Repositorio para manejar las operaciones relacionadas con los usuarios.

package repositorios

import (
	"database/sql"
	"SistemaGestionLibrosElectronicos/modelos"
)

// RegistrarUsuario agrega un nuevo usuario a la base de datos.
func RegistrarUsuario(db *sql.DB, usuario modelos.Usuario) error {
	query := "INSERT INTO Usuario (Nombre, Apellido, Correo) VALUES (?, ?, ?)"
	_, err := db.Exec(query, usuario.Nombre, usuario.Apellido, usuario.Correo)
	return err
}

// ListarUsuarios obtiene todos los usuarios de la base de datos.
func ListarUsuarios(db *sql.DB) ([]modelos.Usuario, error) {
	query := "SELECT IDUsuario, Nombre, Apellido, Correo FROM Usuario"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []modelos.Usuario
	for rows.Next() {
		var usuario modelos.Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Apellido, &usuario.Correo); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

// ActualizarUsuario actualiza la información de un usuario existente.
func ActualizarUsuario(db *sql.DB, usuario modelos.Usuario) error {
	query := "UPDATE Usuario SET Nombre = ?, Apellido = ?, Correo = ? WHERE IDUsuario = ?"
	_, err := db.Exec(query, usuario.Nombre, usuario.Apellido, usuario.Correo, usuario.ID)
	return err
}
