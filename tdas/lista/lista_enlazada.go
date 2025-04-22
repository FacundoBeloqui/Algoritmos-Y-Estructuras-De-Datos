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
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) InsertarPrimero(t T) {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) InsertarUltimo(t T) {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) BorrarPrimero() T {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) VerPrimero() T {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) VerUltimo() T {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) Largo() int {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) Iterar(visitar func(T) bool) {
	//TODO implement me
	panic("implement me")
}

func (l listaEnlazada[T]) Iterador() IteradorLista[T] {
	//TODO implement me
	panic("implement me")
}
