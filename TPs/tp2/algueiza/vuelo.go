package algueiza

type Vuelo interface {
	NumeroVuelo() int
	Aerolinea() string
	Origen() string
	Destino() string
	Matricula() string
	Prioridad() int
	Fecha() string
	Atraso() int
	TiempoDeVuelo() int
	Cancelado() int
}
