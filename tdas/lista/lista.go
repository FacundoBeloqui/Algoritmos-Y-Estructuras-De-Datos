package lista

type Lista[T any] interface {
	//EstaVacia devuelve true si la lista no tiene elementos y false en caso contrario
	EstaVacia() bool

	//InsertarPrimero inserta un elemento al principio de la lista.
	InsertarPrimero(T)

	//InsertarUltimo insertar un elemento al final de la lista.
	InsertarUltimo(T)

	//BorrarPrimero elimina el primer elemento de la lista y lo devuelve . Si esta vacia
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

	//Iterar llama al iterador interno de la lista. El cual permite recorrer la estructura sin que el ususario escriba explicitamente el bucle
	//que recorre los elementos, sino que le pasa una función que se aplicará a cada uno.
	Iterar(visitar func(T) bool)

	//Iterador crea el iterador externo de la lista. El cual permite recorrer la estructura paso a paso dandole el control al usuario.
	Iterador() IteradorLista[T]

}
type IteradorLista[T any] interface {
	//VerActual devuelve el elemento de la posicion actual del iterador.
	//Si ya termino de iterar entra en panico con un mensaje "El iterador termino de iterar".
	VerActual() T

	//HaySiguiente devuelve true si el Iterador no termino de recorrer todavia y False en caso contrario.
	HaySiguiente() bool

	//Siguiente cambia a la proxima posicion del Iterador.
	//Si ya termino de iterar entra en panico con un mensaje "El iterador termino de iterar".
	Siguiente()

	//Insertar inserta un elemento a la izquierda del elemento de la posicion actual del Iterador.
	Insertar(T)

	//Borrar elimina el elemento de la posicion actual del Iterador.
	//Si ya termino de iterar entra en panico con un mensaje "El iterador termino de iterar".
	Borrar() T

}
