package division_y_conquista

//(★) Explicar por qué el siguiente código no es de división y conquista.

// Algoritmo ¿por D&C? Para obtener el máximo de un arreglo

func Maximo(arreglo []int) int {
	if len(arreglo) == 1 {
		return arreglo[0]
	}
	maxRestante := Maximo(arreglo[0 : len(arreglo)-1])
	if arreglo[len(arreglo)-1] > maxRestante {
		return arreglo[len(arreglo)-1]
	} else {
		return maxRestante
	}
}

//no es de division y conquista, ya que en cada llamada recursiva el arreglo se está reduciendo en un
//elemento en vez de dividirse en 2 por lo que la reduccion es lineal en lugar de logaritmica.
