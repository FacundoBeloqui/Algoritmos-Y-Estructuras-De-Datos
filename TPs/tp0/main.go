package main

import (
	"fmt"
	"tp0/ejercicios"
)

func main() {
	vector1 := []int{4, 2, 3, 4}
	maximo := ejercicios.Maximo(vector1)
	fmt.Printf("%d", maximo)
}
