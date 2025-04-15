package division_y_conquista

// Implementar, por división y conquista, una función que determine el mínimo de un arreglo. Indicar y justificar el orden.

func Minimo(arreglo []int, inicio int, fin int) int {
	if inicio == fin {
		return arreglo[inicio]
	}
	medio := (inicio + fin) / 2

	minimoIzq := Minimo(arreglo, inicio, medio)
	minimoDer := Minimo(arreglo, medio+1, fin)
	if minimoIzq < minimoDer {
		return minimoIzq
	}
	return minimoDer
}

// El orden es O(n) usando el teorema maestro, lo cual tiene sentido, ya que no se puede encontrar el minimo sin mirar todos los elementos al menos una vez.
