package pila

const CAPACIDAD_INICIAL = 1
const DOBLE = 2
const MITAD = 2
const CUADRUPLE = 4

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{
		datos:    make([]T, CAPACIDAD_INICIAL),
		cantidad: 0,
	}
}

func redimensionar[T any](p *pilaDinamica[T], nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
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
	if p.cantidad == len(p.datos) {
		redimensionar(p, len(p.datos)*DOBLE)
	}
	p.datos[p.cantidad] = t
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad*CUADRUPLE <= len(p.datos) {
		redimensionar(p, len(p.datos)/MITAD)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}
