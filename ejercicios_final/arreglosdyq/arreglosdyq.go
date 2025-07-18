package arreglosdyq

/*
Dados dos arreglos ordenados A y B, donde B tiene “un elemento menos que A”, implementar un algoritmo de división y
conquista que permita obtener el valor faltante de A en B. Ejemplo, si A = {2, 4, 6, 8, 9, 10, 12} y B = {2, 4,
6, 8, 10, 12}, entonces la salida del algoritmo debe ser o bien la posición 4, o el valor 9 (lo que decidan que devuelva).
Indicar y justificar adecuadamente la complejidad del algoritmo implementado.
*/

func Valor_faltante(arr1, arr2 []int) int {
	return valor_faltante_rec(arr1, arr2, 0, len(arr1)-1)
}
func valor_faltante_rec(arr1, arr2 []int, inicio, fin int) int {
	if inicio == fin {
		return arr1[inicio]
	}
	medio := (inicio + fin) / 2
	if arr1[medio] == arr2[medio] {
		return valor_faltante_rec(arr1, arr2, medio + 1, fin)
	}
	return valor_faltante_rec(arr1, arr2, inicio, medio)
}