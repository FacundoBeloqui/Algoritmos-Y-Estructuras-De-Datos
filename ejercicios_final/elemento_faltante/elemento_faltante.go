package elemento_faltante

func ElementoFaltante(arr []int) int {
	return elementoFaltanteR(arr, 0, len(arr) -1)
}
func elementoFaltanteR(arr []int, inicio, fin int) int {
	if inicio ==  fin {
		return inicio + 1
	}
	mitad := (inicio + fin) / 2

	if arr[mitad] == mitad+1 {
		return elementoFaltanteR(arr, mitad + 1, fin)
	}
	return elementoFaltanteR(arr, inicio, mitad)
}