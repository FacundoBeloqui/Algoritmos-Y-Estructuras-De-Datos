package algueiza

type Interfaz interface {
	AgregarArchivo(string)
	VerTablero(int, string, string, string)
	InfoVuelo(int)
	PrioridadVuelos(int)
	SiguienteVuelo(string, string, string)
	Borrar(string)
}
