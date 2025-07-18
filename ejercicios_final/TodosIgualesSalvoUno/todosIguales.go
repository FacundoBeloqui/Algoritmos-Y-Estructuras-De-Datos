package main

import "fmt"

/*
Implementar un algoritmo que reciba un arreglo de n enteros (con n ≥ 3) en el que todos sus elementos son iguales
salvo 1, y determine (utilizando división y conquista) cual es dicho elemento no repetido. Indicar y justificar la
complejidad del algoritmo implementado.
*/
func todosIguales(arr[]int) int {
	if len(arr) <= 4 {
		var copia1 []int 
		var copia2 []int
		for _, e := range arr {
			if e == arr[0] {
				copia1 = append(copia1, e)
			} else {
				copia2 = append(copia2, e)
			}
		}
		if len(copia1) == 1 {
			return copia1[0]
		} else if len(copia2) == 1 {
			return copia2[0]
		}
		return -1
	}
	mitad := (len(arr)-1) / 2
	if  arr[mitad] != arr[mitad+1] && arr[mitad] != arr[mitad - 1] {
		return arr[mitad]
	}
	izq := todosIguales(arr[:mitad+1])
	der := todosIguales(arr[mitad:])
	if izq == -1 {
		return der
	} else {
		return izq
	}
	
}

func main() {
	var arr[]int  = []int{1,1,1,1,1,2,1,1,1}
	fmt.Println(todosIguales(arr))
}