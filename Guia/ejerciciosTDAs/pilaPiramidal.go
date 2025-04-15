package ejerciciosTDAs

import (
	"tdas/pila"
)

/*
Dada una pila de punteros a enteros, escribir una función que determine si es piramidal.
Una pila de enteros es piramidal si cada elemento es menor a su elemento inferior (en el sentido que va desde el tope de la pila hacia el otro extremo).
La pila no debe ser modificada.
*/

func EsPiramidal(pila pila.Pila[int]) bool {
	var arreglo []int
	for !pila.EstaVacia() {
		arreglo = append(arreglo, pila.Desapilar())
	}
	for i := len(arreglo) - 1; i >= 0; i-- {
		pila.Apilar(arreglo[i])
	}
	for i := 0; i < len(arreglo)-1; i++ {
		if arreglo[i] >= arreglo[i+1] {
			return false
		}
	}
	return true
}
