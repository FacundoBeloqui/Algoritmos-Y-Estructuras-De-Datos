package pila

/* Definición del struct pila proporcionado por la cátedra. */
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func redimensionar[T any](capacidad int, p *pilaDinamica[T]) {
	nuevosDatos := make([]T, capacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

func aumentarCapacidad[T any](p *pilaDinamica[T]) {
	capacidad := len(p.datos)
	if capacidad == 0 {
		capacidad++
	} else {
		capacidad *= 2
	}
	redimensionar(capacidad, p)
}

func disminuirCapacidad[T any](p *pilaDinamica[T]) {
	capacidad := len(p.datos)
	capacidad /= 2
	redimensionar(capacidad, p)
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia.")
	} else {
		return p.datos[p.cantidad-1]
	}
}

func (p *pilaDinamica[T]) Apilar(t T) {
	if p.cantidad == len(p.datos) {
		aumentarCapacidad(p)
	}
	p.datos[p.cantidad] = t
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia.")
	}
	if p.cantidad*4 <= len(p.datos) {
		disminuirCapacidad(p)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, 0),
		cantidad: 0,
	}
}
