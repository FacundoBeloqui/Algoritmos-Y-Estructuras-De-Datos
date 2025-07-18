package ordenaralumnos

import (
	"tdas/cola_prioridad"
)

/*
Trabajamos para una escuela primaria muy estructurada. En dicha escuela hay k cursos, cada uno con m alumnos (es
decir, hay un total de n = k · m alumnos). Todas las mañanas hay que armar filas para cantar Aurora en el patio del
colegio. Primero los alumnos se ubican en una fila correspondiente a su curso, de menor a mayor altura para cantar.
Una vez terminado, proceden a entrar a la escuela de a un alumno por vez, pero deben hacerlo de menor a mayor altura.
Es decir, se debe ordenar a todos los alumnos de menor a mayor. Nosotros sabemos que esto es ineficiente (suelen usar
mergesort, así que es O(n log n)), y desaprovechamos que los alumnos ya estaban ordenados por cursos. Implementar
un algoritmo que reciba k filas (arreglos) de alumnos, cada una previamente ordenada de menor a mayor altura, y nos
devuelva un único arreglo con todos los alumnos ordeados por altura en tiempo menor a O(n log n). Indicar y justificar
la complejidad del algoritmo implementado.
*/
type compuesto struct {
	vector int
	posicion int
	valor int
}
func cmp(a, b compuesto) int {
	return b.valor - a.valor
}

func KMergeAlumnos(cursos [][]int) []int {
	aux := make([]compuesto, len(cursos))
	contador := 0
	for i := range cursos {
		aux = append(aux, compuesto{i, 0, cursos[i][0]})
		contador += len(cursos[i])
	}
	heap := cola_prioridad.CrearHeapArr(aux, cmp)
	res := make([]int, contador)
	for !heap.EstaVacia() {
		minimo := heap.Desencolar()
		altura := minimo.valor
		res = append(res, altura)
		if minimo.posicion + 1 < len(cursos[minimo.vector]) {
			heap.Encolar(compuesto{minimo.vector, minimo.posicion+1, cursos[minimo.vector][minimo.posicion+1]})
		}
	}
	return res
}