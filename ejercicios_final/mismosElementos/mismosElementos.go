package mismoselementos

import (
	"tdas/diccionario"
)

/*
Implementar un algoritmo que reciba dos arreglos desordenados y determine si ambos arreglos tienen los mismos
elementos (y en mismas cantidades). Indicar y justificar la complejidad del algoritmo implementado.
*/

func TienenMismosElementos(cadena1, cadena2 string) bool {
	if len(cadena1) == len(cadena2) {
		dicc1 := diccionario.CrearHash[rune, int]()
		dicc2 := diccionario.CrearHash[rune, int]()
		for _, e := range cadena1 {
			if !dicc1.Pertenece(e) {
				dicc1.Guardar(e, 0)
			}
			dicc1.Guardar(e, dicc1.Obtener(e)+1)
		}
		for _, e := range cadena2 {
			if !dicc2.Pertenece(e) {
				dicc2.Guardar(e, 0)
			}
			dicc2.Guardar(e, dicc2.Obtener(e)+1)
		}
		for iter := dicc1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			clave, cantidad := iter.VerActual()
			if dicc2.Pertenece(clave) && dicc2.Obtener(clave) == cantidad {
				continue
			}
			return false
		}
		return true
	}
	return false
} 
