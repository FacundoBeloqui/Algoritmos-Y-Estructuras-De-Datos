package ejerciciosTDAs

import (
	"tdas/cola"
	"tdas/pila"
)

/*
Implementar una función que reciba un arreglo genérico e invierta su orden, utilizando los TDAs vistos. Indicar y justificar el orden de ejecución.
*/

func InvertirArreglo[T any](arr *[]T) {
	p := pila.CrearPilaDinamica[T]()
	c := cola.CrearColaEnlazada[T]()

	for _, i := range *arr {
		p.Apilar(i)
	}

	for range *arr {
		c.Encolar(p.Desapilar())
	}

	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = c.Desencolar()
	}
}
