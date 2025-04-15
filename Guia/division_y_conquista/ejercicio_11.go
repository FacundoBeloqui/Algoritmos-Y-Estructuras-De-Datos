package division_y_conquista

/*
(★★★) ♠♠ Se tiene un arreglo tal que [1, 1, 1, ..., 0, 0, ...] (es decir, unos seguidos de ceros). Se pide:

Una función de orden O(log⁡n)O(logn) que encuentre el índice del primer 0. Si no hay ningún 0 (solo hay unos), debe devolver -1.
Demostrar con el Teorema Maestro que la función es, en efecto, O(log⁡n)O(logn).

*/

func BuscarPrimerCero(arreglo []int, inicio int, fin int) int {
	if inicio == fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if arreglo[medio] == 0 && arreglo[medio-1] == 1 {
		return medio
	}
	if arreglo[medio] == 0 && arreglo[medio-1] == 0 {
		return BuscarPrimerCero(arreglo, inicio, medio)
	}
	return BuscarPrimerCero(arreglo, medio+1, fin)

}
