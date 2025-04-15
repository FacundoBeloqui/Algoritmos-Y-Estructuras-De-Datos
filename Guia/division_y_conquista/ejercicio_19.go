package division_y_conquista

import "tdas/pila"

/*
Implementar una función insertarEnPos[T any](pila Pila[T], elemento T, n int) que inserte el elemento en
la posición n de la pila. Es decir, la pila debe quedar con sus elementos originales, más el nuevo elemento, el cuál debe
salir de la pila como n-ésimo elemento. En caso que n sea mayor a la cantidad de elementos de la pila, finalizar con
panic. Considerar que se espera una solución simple a un problema simple. Indicar y justificar la complejidad de la
función.
Ejemplo: pila = [1, 2, 3, 4] (tope 4), elemento = 70, n = 2, el estado final de la pila debería ser pila = [1, 2,
70, 3, 4]. Si n = 0, el resultado sería equivalente a simplemente apilar.
*/

func InsertarEnPos[T any](p pila.Pila[T], elemento T, n int) {
	pilaAux := pila.CrearPilaDinamica[T]()
	if n == 0 {
		p.Apilar(elemento)
	} else {
		for i := 0; i < n; i++ {
			if p.EstaVacia() {
				panic("n es mayor a la cantidad de elementos.")
			} else {
				pilaAux.Apilar(p.Desapilar())
			}
		}
		p.Apilar(elemento)
		for !pilaAux.EstaVacia() {
			p.Apilar(pilaAux.Desapilar())
		}
	}
}
