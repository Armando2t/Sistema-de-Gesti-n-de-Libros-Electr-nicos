package repositorios

import (
	"SistemaGestionLibrosElectronicos/modelos"
)

// Lista temporal de multas
var multas []modelos.Multa

// Función para registrar una multa
func RegistrarMulta(multa modelos.Multa) {
	multa.ID = len(multas) + 1  // Asigna un ID único basado en el tamaño de la lista
	multas = append(multas, multa)
}

// Función para listar todas las multas
func ListarMultas() []modelos.Multa {
	return multas
}
