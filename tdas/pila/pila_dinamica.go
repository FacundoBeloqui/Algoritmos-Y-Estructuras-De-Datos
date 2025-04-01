package pila

/* Definición del struct pila proporcionado por la cátedra. */
type pilaDinamica[T any] struct {
	datos     []T
	cantidad  int
	capacidad int
}

func redimensionar[T any](p *pilaDinamica[T]) {
	nuevosDatos := make([]T, p.capacidad)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return p.datos[p.cantidad-1]
	}
}

func (p *pilaDinamica[T]) Apilar(t T) {
	const DOBLE = 2
	if p.cantidad == p.capacidad {
		p.capacidad *= DOBLE
		redimensionar(p)
	}
	p.datos[p.cantidad] = t
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	const MITAD = 2
	const CUADRUPLE = 4
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad*CUADRUPLE <= p.capacidad {
		p.capacidad /= MITAD
		redimensionar(p)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:     make([]T, 1),
		cantidad:  0,
		capacidad: 1,
	}
}
