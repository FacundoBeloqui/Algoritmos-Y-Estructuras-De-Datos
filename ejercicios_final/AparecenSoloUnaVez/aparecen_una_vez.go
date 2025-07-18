package main

import (
	"fmt"
	"tdas/diccionario"
)

//Implementar una función que reciba un arreglo de números y determine cuáles aparecen una única vez. Indicar y justificar la
//complejidad del algoritmo implementado.
func aparecenUnaVez(arr[] int) []int {
	dicc := diccionario.CrearHash[int, int]()
	destino := make([]int, 0)
	for _, elem := range arr {
		if !dicc.Pertenece(elem) {
			dicc.Guardar(elem, 0)
		}
		dicc.Guardar(elem, dicc.Obtener(elem)+1)
	}
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		elem, apariciones := iter.VerActual()
		if apariciones == 1 {
			destino = append(destino, elem)
		}
	}
	return destino
}

func main() {
	arreglo := []int {1,3, 2, 4 , 8,1 , 4, 7, 9, 8}
	fmt.Println(aparecenUnaVez(arreglo))
}