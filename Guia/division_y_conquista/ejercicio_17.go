package division_y_conquista

/*
(★★★★★) Tenemos un arreglo de tamaño 2n de la forma {C1, C2, C3, … Cn, D1, D2, D3, … Dn},
tal que la cantidad total de elementos del arreglo es potencia de 2 (por ende, n también lo es).
Implementar un algoritmo de División y Conquista que modifique el arreglo de tal forma que quede con la forma {C1, D1, C2, D2, C3, D3, …, Cn, Dn},
sin utilizar espacio adicional (obviando el utilizado por la recursividad). ¿Cual es la complejidad del algoritmo?
*/

func Alternar(arr []int) {
	AlternarRecursivo(arr, 0, len(arr)-1)
}

func AlternarRecursivo(arr []int, inicio, fin int) {
	if (fin-inicio+1)/2 <= 2 {
		return
	}
	medio := (inicio + fin) / 2
	n := (fin - inicio + 1) / 2
	m1 := inicio + n/2
	m2 := medio + 1

	for i := 0; i < n/2; i++ {
		arr[m1+1], arr[m2+1] = arr[m2+1], arr[m1+1]
	}

	AlternarRecursivo(arr, inicio, medio)
	AlternarRecursivo(arr, medio+1, fin)
}
