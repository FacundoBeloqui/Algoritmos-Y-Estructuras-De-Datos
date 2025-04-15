package division_y_conquista

/*
Implementar una función suma_total(arreglo []float) float que, por división y conquista, devuelva la suma
de todos los elementos. Indicar y justificar adecuadamente la complejidad de la función implementada.
*/

func SumaTotalRecursivo(arreglo []float64, inicio, fin int) float64 {
	if inicio == fin {
		return arreglo[inicio]
	}
	medio := (inicio + fin) / 2
	return SumaTotalRecursivo(arreglo, inicio, medio) + SumaTotalRecursivo(arreglo, medio+1, fin)
}
func SumaTotal(arreglo []float64) float64 {
	return SumaTotalRecursivo(arreglo, 0, len(arreglo)-1)
}
