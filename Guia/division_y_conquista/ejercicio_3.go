package division_y_conquista

import "fmt"

// Hacerle el seguimiento al siguiente algoritmo:

func imprimirDyC(m int) {
	if m < 4 {
		return
	}
	fmt.Println(m)
	imprimirDyC(m / 4)
	imprimirDyC(m - (m / 4))
}

// Indicar, utilizando el Teorema Maestro, la complejidad del ejercicio anterior.

// la ecuacion de recurrencia es T(n) = A*T(n/b)+O(n^c)

// en este caso, A = 2, B = 4/3, C = 0

// Por ende queda 2T(n/4)+O(1)

// Log 4/3(2) = 2,40 > C, por ende:
// la complejidad es = O(n^2)
