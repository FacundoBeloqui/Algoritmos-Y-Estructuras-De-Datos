package representacionarreglo

import "tdas/cola"

/*Implementar en Go una primitiva que reciba un árbol binario que representa un heap (árbol binario izquierdista, que
cumple la propiedad de heap), y devuelva la representación en arreglo del heap. La firma de la primitiva debe ser
RepresentacionArreglo() []T. Indicar y justificar la complejidad de la primitiva. La estructura del árbol binario es:
type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}*/

type ab[T any] struct {
izquierda *ab[T]
derecha *ab[T]
dato T
}

func (arbol *ab[T]) RepresentacionArreglo() []T {
	arreglo := []T{}
	cola := cola.CrearColaEnlazada[*ab[T]]()
	cola.Encolar(arbol)
	for !cola.EstaVacia() {
		nodo := cola.Desencolar()
		arreglo = append(arreglo, nodo.dato)
		if nodo.izquierda != nil {
			cola.Encolar(nodo.izquierda)
		}
		if nodo.derecha != nil {
			cola.Encolar(nodo.derecha)
		}
	}
	return arreglo
}