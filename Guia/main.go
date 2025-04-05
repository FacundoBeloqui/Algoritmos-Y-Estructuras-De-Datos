package main

import (
	"Guia/division_y_conquista"
	"fmt"
)

func main() {
	arreglo := []int{4, 5, 8, 2, 7}
	numeroMaximo := division_y_conquista.Maximo(arreglo)
	fmt.Println(numeroMaximo)
}
