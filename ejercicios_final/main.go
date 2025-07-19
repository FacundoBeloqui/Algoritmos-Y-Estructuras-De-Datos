package main

import (
	ejerciciotopk "ejercicios_final/ejercicioTopk"
	"fmt"
)

func main() {
	arr := []int{1, 5, 3, 4, 2, 8}
	nuevo := ejerciciotopk.Multiplicar(arr, 3)
	fmt.Println(nuevo)
}