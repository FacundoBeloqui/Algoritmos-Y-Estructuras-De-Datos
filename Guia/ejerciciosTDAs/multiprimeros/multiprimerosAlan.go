package multiprimeros

/*
(★★) Implementar la primitiva func (cola *colaEnlazada[T]) Multiprimeros(k int) []T que dada una cola y un número kk,
devuelva los primeros kk elementos de la cola, en el mismo orden en el que habrían salido de la cola.
En caso que la cola tenga menos de kk elementos. Si hay menos elementos que kk en la cola, devolver un slice del tamaño de la cola.
Indicar y justificar el orden de ejecución del algoritmo.
*/

func (cola *ColaEnlazada[T]) Multiprimeros(k int) []T {
	var arr []T
	for k != 0 && cola.primero != nil {
		arr = append(arr, cola.primero.dato)
		cola.primero = cola.primero.siguiente
		k--
	}
	return arr
}
