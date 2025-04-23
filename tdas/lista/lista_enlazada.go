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

/*
func (l *listaEnlazada[T]) verificarListaVacia(){
	if l.EstaVacia(){
		panic("La lista esta vacia")
	}
}
*/

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
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	//l.verificarListaVacia()
	valor := l.primero.dato
	l.primero = l.primero.siguiente
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo-- 
	return valor
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	//l.verificarListaVacia()
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	//l.verificarListaVacia()
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

/*
func (i *iterListaEnlazada[T]) VerIteradorEsNil() { perdon por los nombres de las funciones soy un desastre jajjaja
	if !i.HaySiguiente() {
		panic("El iterador ya termino de recorrer")
	}
}
*/

func (i *iterListaEnlazada[T]) VerActual() T {
	if i.actual == nil {
		panic("El iterador ya termino de recorrer")
	}
	//i.VerIteradorEsNil()
	return i.actual.dato //

}
func (i *iterListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}
func (i *iterListaEnlazada[T]) Siguiente() {
	if i.actual == nil {
		panic("El iterador ya termino de recorrer")
	}
	//i.VerIteradorEsNil()
	i.anterior = i.actual         
	i.actual = i.actual.siguiente 

}

func (i *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := crearNuevoNodo(elemento)
	/*
		nuevoNodo.siguiente = i.actual

		if i.anterior == nil o i.actual == i.lista.primero{
			i.lista.primero = nuevoNodo
		} else {
			i.anterior.siguiente = nuevoNodo
		}

		if i.actual == nil{
			i.lista.ultimo.siguiente = nuevoNodo //esta no estoy segura de si va
			i.lista.ultimo = nuevoNodo
		}

		i.actual = nuevoNodo
		i.lista.largo++
	*/

	if i.actual == i.lista.primero {
		i.lista.primero = nuevoNodo

	} else if i.actual == nil {
		i.lista.ultimo.siguiente = nuevoNodo 
		i.lista.ultimo = nuevoNodo           

	} else {
		i.anterior.siguiente = nuevoNodo 
	}

	i.actual = nuevoNodo
	i.lista.largo++
}

func (i *iterListaEnlazada[T]) Borrar() T {
	if !i.HaySiguiente() {
		panic("No hay elemento para borrar")
	}
	//i.VerIteradorEsNil()
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
