package modelos

// Estructura de una Multa en el sistema.
type Multa struct {
	ID        int
	UsuarioID int
	Monto     float64
	Estado    string
}
