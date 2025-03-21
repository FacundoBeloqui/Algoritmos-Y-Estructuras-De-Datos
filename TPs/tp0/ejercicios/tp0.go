package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	} else {
		maximo := vector[0]
		posicion := 0
		for i := 1; i < len(vector); i++ {
			if vector[i] > maximo {
				maximo = vector[i]
				posicion = i
			}
		}
		return posicion
	}
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	for i := 0; i < len(vector1) && i < len(vector2); i++ {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}
	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector1) > len(vector2) {
		return 1
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := 0; i < len(vector); i++ {
		maxPosicion := Maximo(vector[:len(vector)-i])
		Swap(&vector[maxPosicion], &vector[len(vector)-i-1])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).

func sumaRecursiva(vector []int, indice int) int {
	if indice >= len(vector) {
		return 0
	}
	return vector[indice] + sumaRecursiva(vector, indice+1)
}

func Suma(vector []int) int {
	return sumaRecursiva(vector, 0)
}

func EsCadenaCapicuaRecursiva(cadena string, inicio int, fin int) bool {
	if inicio == len(cadena)/2 {
		return true
	}
	if cadena[inicio] != cadena[fin] {
		return false
	}

	return EsCadenaCapicuaRecursiva(cadena, inicio+1, fin-1)
}
func EsCadenaCapicua(cadena string) bool {
	return EsCadenaCapicuaRecursiva(cadena, 0, len(cadena)-1)
}
