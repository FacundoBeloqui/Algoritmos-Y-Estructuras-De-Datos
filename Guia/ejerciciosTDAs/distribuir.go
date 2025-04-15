package ejerciciosTDAs

import (
	"tdas/cola"
	"tdas/pila"
)

/*
Se desea extraer los elementos de una pila de manera tal que los que corresponden a “posiciones” impares (es decir: el
primero, el tercero, el quinto, etc. al ser desapilados) queden encolados en una cola, y los de las posiciones pares queden
en la pila. Los elementos de dicha cola deben poder ser desencolados manteniendo el mismo orden que tenían en la pila
original.
Implementar en Go una función Distribuir[T any](Pila[T]) Cola[T] que reciba una pila y devuelva una cola
cumpliendo el comportamiento pedido.
Por ejemplo, si se aplica la función Distribuir a la pila p = [1, 2, 4, 3, 15, 6, 0, 8, 9, 20] (tope = 20), esta
devuelve una cola c = [20, 8, 6, 3, 2] (primero = 20) y la pila queda p = [1, 4, 15, 0, 9] (tope = 9).
*/
func Distribuir[T any](p pila.Pila[T]) cola.Cola[T] {
	c := cola.CrearColaEnlazada[T]()
	pilaAux := pila.CrearPilaDinamica[T]()
	for i := 1; !p.EstaVacia(); i++ {
		if i%2 != 0 {
			c.Encolar(p.Desapilar())
		} else {
			pilaAux.Apilar(p.Desapilar())
		}
	}
	for !pilaAux.EstaVacia() {
		p.Apilar(pilaAux.Desapilar())
	}
	return c
}
