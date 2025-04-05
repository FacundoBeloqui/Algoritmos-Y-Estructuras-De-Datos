package cola

type NodoCola[T any] struct {
	dato T
	prox *NodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *NodoCola[T]
	ultimo  *NodoCola[T]
}

func CrearNuevoNodo[T any](elemento T) *NodoCola[T] {
	return &NodoCola[T]{
		dato: elemento,
		prox: nil,
	}
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
	}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	} else {
		return c.primero.dato
	}
}

func (c *colaEnlazada[T]) Encolar(t T) {
	nuevoNodo := CrearNuevoNodo[T](t)
	if c.EstaVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else {
		c.ultimo.prox = nuevoNodo
		c.ultimo = nuevoNodo
	}

}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	} else {
		valor := c.primero.dato
		c.primero = c.primero.prox
		if c.primero == nil {
			c.ultimo = nil
		}
		return valor
	}
}
