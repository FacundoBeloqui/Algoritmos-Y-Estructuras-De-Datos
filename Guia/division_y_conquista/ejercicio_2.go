package division_y_conquista

// Explicar por qué el siguiente código no es de división y conquista.

func maximo(arreglo []int) int {
	medio := len(arreglo) / 2
	maxIzquierda := _maximo(arreglo, 0, medio)
	maxDerecha := _maximo(arreglo, medio+1, len(arreglo)-1)
	if maxIzquierda > maxDerecha {
		return maxIzquierda
	} else {
		return maxDerecha
	}
}

func _maximo(arreglo []int, inicio int, fin int) int {
	maximo := arreglo[inicio]
	for i := inicio + 1; i <= fin; i++ {
		if maximo < arreglo[i] {
			maximo = arreglo[i]
		}
	}
	return maximo
}

//No es de division y conquista ya que el algoritmo si bien es recursivo a la hora de ir
// llamandose a si misma en los lados izquierdo y derecho, al haber un for y la ausencia de un caso base, no
// se puede considerar un algoritmo recursivo.
