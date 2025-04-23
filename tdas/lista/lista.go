package lista

type Lista[T any] interface {
	//EstaVacia devuelve true si la lista esta vacia y false en caso contrario
	EstaVacia() bool

	//InsertarPrimero inserta un elemento al principio de la lista.
	InsertarPrimero(T)

	//InsertarUltimo insertar un elemento al final de la lista.
	InsertarUltimo(T)

	//BorrarPrimero borra el primer elemento de la lista. SI esta vacia
	// entra en panico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	//VerPrimero devuelve el primer elemento de la lista. Si esta vacia
	// entra en panico con un mensaje "La lista esta vacia".
	VerPrimero() T

	//VerUltimo devuelve el ultimo elemento de la lista. Si esta vacia
	// entra en panico con un mensaje "La lista esta vacia".
	VerUltimo() T

	//Largo devuelve el largo de la lista.
	Largo() int

	//Iterar llama al iterador interno de la lista.
	Iterar(visitar func(T) bool)

	//Iterador crea el iterador externo de la lista.
	Iterador() IteradorLista[T]
}
type IteradorLista[T any] interface {
	//VerActual devuelve el elemento de la posicion actual del iterador.
	VerActual() T

	//HaySiguiente devuelve true si el Iterador no termino de recorrer todavia y False en caso contrario.
	HaySiguiente() bool

	//Siguiente cambia a la proxima posicion del Iterador.
	Siguiente()

	//Insertar inserta un elemento a la izquierda del elemento de la posicion actual del Iterador.
	Insertar(T)

	//Borrar borra el elemento de la posicion actual del Iterador.
	Borrar() T
}
