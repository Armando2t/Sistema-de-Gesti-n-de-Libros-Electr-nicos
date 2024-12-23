package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// ConectarBD establece la conexión a la base de datos.
func ConectarBD() (*sql.DB, error) {
	// Cargar el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env, utilizando variables del entorno")
	}

	// Obtener credenciales desde variables de entorno o el archivo .env.
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Verificar si las variables están configuradas.
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbName == "" {
		return nil, fmt.Errorf("faltan variables de entorno para la conexión a la base de datos")
	}

	// Construir el DSN (Data Source Name).
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	// Abrir la conexión.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la conexión a la base de datos: %w", err)
	}

	// Verificar la conexión.
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al verificar la conexión con la base de datos: %w", err)
	}

	log.Println("Conexión exitosa a la base de datos")
	return db, nil
}

// CerrarBD cierra la conexión a la base de datos.
// Parámetro:
// - db: La instancia de conexión a la base de datos que se desea cerrar.
func CerrarBD(db *sql.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error al cerrar la conexión a la base de datos: %v\n", err)
		} else {
			log.Println("Conexión cerrada correctamente")
		}
	}
}
