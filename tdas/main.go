package main

import (
	"fmt"
	"tdas/cola"
)

func main() {

	cola1 := cola.CrearColaEnlazada[int]()
	cola1.Encolar(1)
	fmt.Println(cola1.EstaVacia())
	cola1.Desencolar()
	fmt.Println(cola1.EstaVacia())

}
