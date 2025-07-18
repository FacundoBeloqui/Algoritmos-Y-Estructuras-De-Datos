package pilaordenada

import (
	"tdas/pila"
)

/*
Implementar una función que dada una pila, determine si la misma se encuentra ordenada (es decir, se ingresaron los
elementos de menor a mayor). La pila debe quedar en el mismo estado al original al terminar la ejecución de la función.
Indicar y justificar la complejidad de la función.
*/
func PilaEstaOrdenada(p pila.Pila[int]) bool {
	pila_aux := pila.CrearPilaDinamica[int]()
	estaOrdenada := true
	for !p.EstaVacia() {
		desapilado := p.Desapilar()

		if !p.EstaVacia() {
			if desapilado < p.VerTope(){
				estaOrdenada = false
			}
		}
		
		pila_aux.Apilar(desapilado)
	}
	for !pila_aux.EstaVacia() {
		p.Apilar(pila_aux.Desapilar())
	}
	return estaOrdenada
}
