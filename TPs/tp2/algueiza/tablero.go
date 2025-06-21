package algueiza

type Tablero interface {
	AgregarVuelo(vuelo)
	VerTablero(int, string, string, string) ([][]string, error)
	InfoVuelo(int) ([]string, error)
	PrioridadVuelos(int) []string
	SiguienteVuelo(string, string, string) ([]string, bool)
	Borrar(string, string) ([]string, error)
}
