package calculadoraPolaca

type Operacion interface {
	Simbolo() string
	Prioridad() int
	Asociatividad() Asociatividad
}
