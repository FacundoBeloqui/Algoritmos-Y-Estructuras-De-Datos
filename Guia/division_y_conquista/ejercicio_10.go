package division_y_conquista

//(★★★) Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado
//(todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar. Indicar y justificar el orden.

func ElementoDesordenadoRecursivo(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if arr[medio-1] > arr[medio] || medio > arr[medio+1] {
		return medio
	}
	if arr[medio-1] < arr[medio] {
		return ElementoDesordenadoRecursivo(arr, inicio, medio)
	}
	return ElementoDesordenadoRecursivo(arr, medio+1, fin)

}

func ElementoDesordenado(arr []int) int {
	return ElementoDesordenadoRecursivo(arr, 0, len(arr)-1)
}
