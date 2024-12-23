// Autor: Armando Topón
// Fecha: 22/12/2024
// Descripción: Punto de entrada principal del sistema.

package main

import (
	"log"
	"net/http"
	"SistemaGestionLibrosElectronicos/config"
	"SistemaGestionLibrosElectronicos/controles"
)

func main() {
	// Configuración inicial de la base de datos.
	db, err := config.ConectarBD()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Configuración de rutas.
	http.HandleFunc("/usuarios", controles.ManejarUsuarios)
	http.HandleFunc("/libros", controles.ManejarLibros)
	http.HandleFunc("/prestamos", controles.ManejarPrestamos)
	http.HandleFunc("/multas", controles.ManejarMultas)
	
	// Rutas para funcionalidades con templates.
	http.HandleFunc("/multas/historial", controles.VerHistorialMultas)
	http.HandleFunc("/prestamos/activos", controles.VerPrestamosActivos)

	// Iniciar el servidor.
	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}




