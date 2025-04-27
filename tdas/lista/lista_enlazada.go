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

func verificarListaVacia[T any](l *listaEnlazada[T]) {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(t T) {
	nuevoNodo := crearNuevoNodo(t)
	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	}
	nuevoNodo.siguiente = l.primero
	l.primero = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(t T) {
	nuevoNodo := crearNuevoNodo(t)
	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.siguiente = nuevoNodo
	}
	l.ultimo = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	verificarListaVacia(l)
	valor := l.primero.dato
	l.primero = l.primero.siguiente
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--
	return valor
}

func (l *listaEnlazada[T]) VerPrimero() T {
	verificarListaVacia(l)
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	verificarListaVacia(l)
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for nodo := l.primero; nodo != nil; nodo = nodo.siguiente {
		if !visitar(nodo.dato) {
			break
		}
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{
		actual:   l.primero,
		anterior: nil,
		lista:    l,
	}
}

func verIteradorEsNil[T any](i *iterListaEnlazada[T]) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (i *iterListaEnlazada[T]) VerActual() T {
	verIteradorEsNil(i)
	return i.actual.dato //

}
func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}
func (i *iterListaEnlazada[T]) Siguiente() {
	verIteradorEsNil(i)
	i.anterior = i.actual
	i.actual = i.actual.siguiente

}

func (i *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := crearNuevoNodo(elemento)
	nuevoNodo.siguiente = i.actual

	if i.anterior == nil {
		i.lista.primero = nuevoNodo
	} else {
		i.anterior.siguiente = nuevoNodo
	}

	if i.actual == nil {
		i.lista.ultimo = nuevoNodo
	}

	i.actual = nuevoNodo
	i.lista.largo++
}

func (i *iterListaEnlazada[T]) Borrar() T {
	verIteradorEsNil(i)
	dato := i.actual.dato

	if i.anterior != nil {
		i.anterior.siguiente = i.actual.siguiente
	} else {
		i.lista.primero = i.actual.siguiente
	}

	if i.actual == i.lista.ultimo {
		i.lista.ultimo = i.anterior
	}

	i.actual = i.actual.siguiente
	i.lista.largo--
	return dato
}
