package division_y_conquista

/*
(★★★★★) Implementar una función (que utilice división y conquista) de orden O(n log n)
que dado un arreglo de n números enteros devuelva true o false según si existe
algún elemento que aparezca más de la mitad de las veces. Justificar el orden de la solución. Ejemplos:

[1, 2, 1, 2, 3] -> false
[1, 1, 2, 3] -> false
[1, 2, 3, 1, 1, 1] -> true
[1] -> true


*/

func MasDeLaMitad(arr []int) bool {
	candidato, contador := encontrarCandidato(arr)
	contador = 0
	for _, num := range arr {
		if num == candidato {
			contador++
		}
	}
	return contador > len(arr)/2
}

func encontrarCandidato(arr []int) (int, int) {
	candidato := 0
	contador := 0

	for _, num := range arr {
		if contador == 0 {
			candidato = num
			contador = 1
		} else if candidato == num {
			contador++
		} else {
			contador--
		}
	}

	return candidato, contador
}
