package iterador_interno

import (
	"tdas/lista"
)

/*Implementar una función map[T any, V any](Lista[K], func(K) V) Lista[V] que dada una lista original, cree una nueva lista
con el resultado de aplicarle a cada elemento la función pasada por parámetro. Para que el ejercicio esté completamente bien, se
espera que se implemente utilizando el iterador interno de la lista. Indicar y justificar la complejidad de la función.*/

func Map[T any, V any, K any](l lista.Lista[K], funcion func(K) V) lista.Lista[V] {
	nuevaLista := lista.CrearListaEnlazada[V]()
	l.Iterar(func(k K) bool {
		nuevaLista.InsertarUltimo(funcion(k))
		return true
	})
	return nuevaLista
}