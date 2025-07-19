package ejerciciotopk

import (
	"tdas/cola_prioridad"
)

/*Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en el que
para cada posición i de dicho arreglo, contenga el resultado de la multiplicación de los primeros k máximos del arreglo A
entre las posición [0;i] (incluyendo a i). Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor -1.
Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 15, 60, 60, 160]. Indicar
y justificar la complejidad del algoritmo implementado.*/

func cmp(a, b int) int {
	return b - a
}

func Multiplicar(arr []int, k int) []int {
	nuevoArr := make([]int, len(arr))
	for i := 0; i < k - 1; i++ {
		nuevoArr[i] = -1
	}
	heap := cola_prioridad.CrearHeapArr(arr[:k], cmp)
	for i := k-1; i < len(arr); i++ {
		if i != k-1 && arr[i] > heap.VerMax() {
			heap.Desencolar()
			heap.Encolar(arr[i])
		}
		contador := 1
		arrTemporal := make([]int, 0, k)
		for !heap.EstaVacia() {
			elemento := heap.Desencolar()
			arrTemporal = append(arrTemporal, elemento)
			contador *= elemento
		}
		nuevoArr[i] = contador
		for j := 0; j < k; j++{
			heap.Encolar(arrTemporal[j])
		}
	}
	return nuevoArr
}