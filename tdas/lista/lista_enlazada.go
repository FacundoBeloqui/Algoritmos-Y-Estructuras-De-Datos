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
func (l listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(t T) {
	nuevoNodo := crearNuevoNodo(t)
	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = l.primero
	}
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
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	valor := l.primero.dato
	l.primero = l.primero.siguiente
	if l.primero == nil {
		l.ultimo = nil
	}
	return valor
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l listaEnlazada[T]) Iterar(visitar func(T) bool) {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) Iterador() IteradorLista[T] {
	//TODO implement me
	panic("implement me")
}

func (i *iterListaEnlazada[T]) VerActual() T {
	if i.actual == nil {
		panic("El iterador ya termino de recorrer")
	} else {
		return i.actual.dato
	}
}
func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}
func (i *iterListaEnlazada[T]) Siguiente() {
	if i.actual == nil {
		panic("El iterador ya termino de recorrer")
	} else {
		i.anterior = i.actual
		i.actual = i.actual.siguiente
	}
}

func (i *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := crearNuevoNodo(elemento)
	if i.actual == i.lista.primero {
		i.lista.primero = nuevoNodo

	} else if i.actual == nil {
		i.lista.ultimo = nuevoNodo

	} else {
		i.anterior = nuevoNodo
	}
	nuevoNodo.siguiente = i.actual
	i.actual = nuevoNodo
	i.lista.largo++
}

func (i *iterListaEnlazada[T]) Borrar() T {
	if i.actual == nil {
		panic("No hay elemento para borrar")
	}
	dato := i.actual.dato
	if i.anterior != nil {
		i.actual = i.actual.siguiente
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
