package esabb

/*Implementar una primitiva para el árbol binario EsABB(func(T, T) int) bool que reciba una función de comparación y determine
si el árbol cumple con la propiedad de ABB para dicha función de comparación. Indicar y justificar la complejidad del algoritmo
implementado.
A fines del ejercicio, considerar que la estructura del árbol es la indicada en el dorso a este examen.

type arbol[T any] struct {
izq *arbol
der *arbol
clave T
}*/

type arbol[T any] struct {
izq *arbol[T]
der *arbol[T]
clave T
}

func (ar *arbol[T]) EsABB(cmp func(T, T) int) bool {
	arr := []*arbol[T]{}
	arr = ar.InOrder(arr)
	for i := 1; i < len(arr); i++ {
		if cmp(arr[i-1].clave, arr[i].clave) > 0 {
			return false
		} 
	}
	return true
}

func (nodo *arbol[T]) InOrder(arr []*arbol[T]) []*arbol[T]  {
	if nodo == nil {
		return arr
	}
	nodo.izq.InOrder(arr)
	arr = append(arr, nodo)
	nodo.der.InOrder(arr)
	return arr
}