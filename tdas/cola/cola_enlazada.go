package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNuevoNodo[T any](elemento T) *nodoCola[T] {
	return &nodoCola[T]{
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
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(t T) {
	nuevoNodo := crearNuevoNodo[T](t)
	if c.EstaVacia() {
		c.primero = nuevoNodo
	} else {
		c.ultimo.prox = nuevoNodo
	}
	c.ultimo = nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	valor := c.primero.dato
	c.primero = c.primero.prox
	if c.primero == nil {
		c.ultimo = nil
	}
	return valor

}

/*Implementar para la cola enlazada la primitiva Consumir(accion func (T)) que aplique la función accion a todos
los elementos de la cola. Al terminar la ejecución, la cola debe quedar vacía. Se espera que se implemente sin utilizar
otras primitivas, para demostrar el conocimiento sobre estructuras enlazadas. Indicar y justificar la complejidad de la
primitiva.*/

func (c *colaEnlazada[T]) Consumir(accion func(T)) {
	actual := c.primero
	for actual != nil {
		accion(actual.dato)
		actual = actual.prox
	}
	c.primero = nil
	c.ultimo = nil
}
