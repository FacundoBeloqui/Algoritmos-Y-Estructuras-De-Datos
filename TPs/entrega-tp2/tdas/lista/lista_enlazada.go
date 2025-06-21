package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func crearNuevoNodo[T any](elemento T) *nodoLista[T] {
	return &nodoLista[T]{
		dato:      elemento,
		siguiente: nil,
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{
		primero: nil,
		ultimo:  nil,
		largo:   0,
	}
}

func (lista *listaEnlazada[T]) verificarListaVacia() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(t T) {
	nuevoNodo := crearNuevoNodo(t)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	nuevoNodo.siguiente = lista.primero
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(t T) {
	nuevoNodo := crearNuevoNodo(t)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.verificarListaVacia()
	valor := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return valor
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.verificarListaVacia()
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.verificarListaVacia()
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for nodo := lista.primero; nodo != nil; nodo = nodo.siguiente {
		if !visitar(nodo.dato) {
			break
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
}

func (iter *iterListaEnlazada[T]) verIteradorEsNil() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	iter.verIteradorEsNil()
	return iter.actual.dato //

}
func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}
func (iter *iterListaEnlazada[T]) Siguiente() {
	iter.verIteradorEsNil()
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente

}

func (iter *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := crearNuevoNodo(elemento)
	nuevoNodo.siguiente = iter.actual

	if iter.anterior == nil {
		iter.lista.primero = nuevoNodo
	} else {
		iter.anterior.siguiente = nuevoNodo
	}

	if iter.actual == nil {
		iter.lista.ultimo = nuevoNodo
	}

	iter.actual = nuevoNodo
	iter.lista.largo++
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	iter.verIteradorEsNil()
	dato := iter.actual.dato

	if iter.anterior != nil {
		iter.anterior.siguiente = iter.actual.siguiente
	} else {
		iter.lista.primero = iter.actual.siguiente
	}

	if iter.actual == iter.lista.ultimo {
		iter.lista.ultimo = iter.anterior
	}

	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return dato
}
