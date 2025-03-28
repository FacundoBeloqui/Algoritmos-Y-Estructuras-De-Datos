package main

import (
	"TDAs/pila"
	"fmt"
)

func main() {
	pila1 := pila.CrearPilaDinamica[int]()
	pila1.Apilar(10)
	pila1.Apilar(20)
	pila1.Apilar(30)
	pila1.Desapilar()
	pila1.Desapilar()
	pila1.Desapilar()
	fmt.Println("El tope es", pila1.EstaVacia())
}
