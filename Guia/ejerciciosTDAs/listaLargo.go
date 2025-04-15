package ejerciciosTDAs

/*
Implementar una primitiva Largo() int para la lista enlazada que permita obtener la cantidad de elementos que
tiene. Indicar (y justificar) el orden de la primitiva. Considerar que la estructura interna de la lista es:
*/
type lista[T any] struct {
	dato      T
	siguiente *lista[T]
}

func (l *lista[T]) Largo() int {
	var largo int
	actual := l
	for actual != nil {
		largo++
		actual = actual.siguiente
	}
	return largo
}
