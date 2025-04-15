package division_y_conquista

// Implementar, por división y conquista, una función que dado un arreglo y su largo, determine si el mismo se encuentra ordenado. Indicar y justificar el orden.

func EstaOrdenadoRecursivo(arr []int, inicio, fin int) bool {
	if inicio >= fin {
		return true
	}
	medio := (inicio + fin) / 2
	if medio-1 >= 0 && arr[medio-1] > arr[medio] {
		return false
	}
	if medio+1 <= len(arr)-1 && arr[medio+1] < arr[medio] {
		return false
	}
	if arr[medio-1] < arr[medio] {
		return EstaOrdenadoRecursivo(arr, inicio, medio)
	}
	return EstaOrdenadoRecursivo(arr, medio+1, fin)
}

func EstaOrdenado(arr []int) bool {
	return EstaOrdenadoRecursivo(arr, 0, len(arr)-1)
}
